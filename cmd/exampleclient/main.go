package exampleclient

import (
	"github.com/josudoey/gopp/gorpc"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultAddr = "localhost:8080"
)

func dialExampleClient(cmd *cobra.Command) (gorpc.ExampleClient, error) {
	addr, err := cmd.Flags().GetString("addr")
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := gorpc.NewExampleClient(conn)
	return client, nil
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example-client",
		Short: "run example client",
	}
	cmd.AddCommand(NewEchoCommand())
	cmd.AddCommand(NewEchoStreamCommand())
	cmd.PersistentFlags().String("addr", defaultAddr, "the address to connect to")
	return cmd
}
