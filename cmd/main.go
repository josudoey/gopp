package main

import (
	"context"
	"log"

	echo "github.com/josudoey/gopp/cmd/echo"
	exampleserver "github.com/josudoey/gopp/cmd/exampleserver"
	"github.com/spf13/cobra"
)

func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "cmd",
	}
	cmd.AddCommand(exampleserver.NewCommand())
	cmd.AddCommand(echo.NewCommand())
	return cmd
}

func main() {
	cmd := newCommand()
	if err := cmd.ExecuteContext(context.Background()); err != nil {
		log.Fatal(err)
	}
}
