package exampleclient

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/josudoey/gopp/gorpc"
	"github.com/spf13/cobra"
)

// see https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func NewEchoStreamCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "echo-stream <message>",
		Short: "rpc echo-stream",
		RunE: func(cmd *cobra.Command, argv []string) error {
			internal, err := cmd.Flags().GetDuration("interval")
			if err != nil {
				return err
			}

			client, err := dialExampleClient(cmd)
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			stream, err := client.EchoStream(ctx)
			if err != nil {
				return err
			}
			req := &gorpc.EchoRequest{
				Message: strings.Join(argv, " "),
			}

			for {
				stream.Send(req)
				res, err := stream.Recv()
				if err == io.EOF {
					return nil
				}
				if err != nil {
					return err
				}
				fmt.Printf("%s\n", res.GetMessage())
				<-time.After(internal)
			}
		},
	}
	cmd.Flags().Duration("interval", time.Second, "time interval")
	return cmd
}
