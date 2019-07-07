package commands

import (
	"cli/cli/command/log"
	"cli/cli/command/port"
	"cli/cli/command/removeall"

	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command, dockerCli client.APIClient) {
	cmd.AddCommand(removeall.NewRemoveAllCommand(dockerCli))
	cmd.AddCommand(port.NewGetPortsCommand(dockerCli))
	cmd.AddCommand(log.NewGetLogfilePathCommand(dockerCli))

}
