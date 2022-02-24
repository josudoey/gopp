package echo

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/josudoey/gopp/gorpc"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultAddr = "localhost:8080"
)

type Options struct {
	addr string
}

func NewOptions() *Options {
	return &Options{
		addr: defaultAddr,
	}
}

// see https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go

func NewCommand() *cobra.Command {
	o := NewOptions()
	cmd := &cobra.Command{
		Use:   "echo <message>",
		Short: "run example client",
		Run: func(cmd *cobra.Command, argv []string) {
			conn, err := grpc.Dial(o.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := gorpc.NewExampleServiceClient(conn)
			ctx := context.Background()
			r, err := c.Echo(ctx, &gorpc.EchoRequest{Message: strings.Join(argv, " ")})
			if err != nil {
				log.Fatalf("could not echo: %v", err)
			}
			fmt.Printf("%s\n", r.GetMessage())
		},
	}
	cmd.Flags().StringVar(&o.addr, "addr", o.addr, "the address to connect to")

	return cmd
}
