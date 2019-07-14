package log

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/docker/docker/client"

	"github.com/spf13/cobra"
)

func NewLsCommand(dockerCli client.APIClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "ls [CONTAINERNAME|CONTAINERID]",
		Short: "Display logfile size",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runShowLogfileSize(dockerCli, args)
		},
	}
	return cmd
}

func runShowLogfileSize(dockerCli client.APIClient, args []string) error {

	ctx := context.Background()
	containerInfo, err := dockerCli.ContainerInspect(ctx, args[0])
	if err != nil {
		return err
	}
	if containerInfo.ContainerJSONBase == nil {
		return fmt.Errorf("%s のコンテはありません\n", args[0])
	}

	result, err := exec.Command("ls", "-lah", containerInfo.ContainerJSONBase.LogPath).Output()

	fmt.Println(string(result))

	return nil
}
