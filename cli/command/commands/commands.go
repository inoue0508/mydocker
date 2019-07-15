package commands

import (
	"cli/cli/command/log"
	"cli/cli/command/name"
	"cli/cli/command/port"

	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command, dockerCli client.APIClient) {
	cmd.AddCommand(name.NewGetNameCommand(dockerCli))
	cmd.AddCommand(port.NewGetPortsCommand(dockerCli))
	cmd.AddCommand(log.NewGetLogfilePathCommand(dockerCli))

}
