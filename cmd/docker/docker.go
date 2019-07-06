package main

import (
	"fmt"
	"os"

	"cli/cli/command"
	"cli/cli/command/commands"

	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

func main() {
	dockerCli, err := command.NewDockerCli()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := runDocker(dockerCli); err != nil {
		fmt.Printf("Error:%s", err)
		os.Exit(1)
	}
}

func runDocker(dockerCli client.APIClient) error {

	cmd := newDockerCommand(dockerCli)

	return cmd.Execute()
}

func newDockerCommand(dockerCli client.APIClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:           "mydocker [OPTIONS] COMMAND [ARGS...]",
		Short:         "A self-sufficient runtime for containers",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return fmt.Errorf("mydocker: '%s' is not a mydocker command.\nSee 'mydocker --help'", args[0])
		},
		Version:               fmt.Sprintf("version:%s\n", "1.0"),
		DisableFlagsInUseLine: true,
	}

	cmd.SetOutput(os.Stderr)
	commands.AddCommands(cmd, dockerCli)
	return cmd
}
