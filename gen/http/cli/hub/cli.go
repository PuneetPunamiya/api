// Code generated by goa v3.1.3, DO NOT EDIT.
//
// hub HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	resourcec "github.com/tektoncd/hub/api/gen/http/resource/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `resource (all|info)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` resource all` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		resourceFlags = flag.NewFlagSet("resource", flag.ContinueOnError)

		resourceAllFlags = flag.NewFlagSet("all", flag.ExitOnError)

		resourceInfoFlags          = flag.NewFlagSet("info", flag.ExitOnError)
		resourceInfoResourceIDFlag = resourceInfoFlags.String("resource-id", "REQUIRED", "ID of resource to be shown")
	)
	resourceFlags.Usage = resourceUsage
	resourceAllFlags.Usage = resourceAllUsage
	resourceInfoFlags.Usage = resourceInfoUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "resource":
			svcf = resourceFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "resource":
			switch epn {
			case "all":
				epf = resourceAllFlags

			case "info":
				epf = resourceInfoFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "resource":
			c := resourcec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "all":
				endpoint = c.All()
				data = nil
			case "info":
				endpoint = c.Info()
				data, err = resourcec.BuildInfoPayload(*resourceInfoResourceIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// resourceUsage displays the usage of the resource command and its subcommands.
func resourceUsage() {
	fmt.Fprintf(os.Stderr, `The resource service provides all resources information
Usage:
    %s [globalflags] resource COMMAND [flags]

COMMAND:
    all: Get all Resources
    info: Get one Resource info

Additional help:
    %s resource COMMAND --help
`, os.Args[0], os.Args[0])
}
func resourceAllUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] resource all

Get all Resources

Example:
    `+os.Args[0]+` resource all
`, os.Args[0])
}

func resourceInfoUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] resource info -resource-id UINT

Get one Resource info
    -resource-id UINT: ID of resource to be shown

Example:
    `+os.Args[0]+` resource info --resource-id 5393976028859956723
`, os.Args[0])
}
