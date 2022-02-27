package exampleclient

import (
	"fmt"
	"strings"

	"github.com/josudoey/gopp/gorpc"
	"github.com/spf13/cobra"
)

// see https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func NewEchoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "echo <message>",
		Short: "rpc echo",
		RunE: func(cmd *cobra.Command, argv []string) error {
			client, err := dialExampleClient(cmd)
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			req := &gorpc.EchoRequest{
				Message: strings.Join(argv, " "),
			}
			r, err := client.Echo(ctx, req)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", r.GetMessage())
			return nil
		},
	}
	return cmd
}
