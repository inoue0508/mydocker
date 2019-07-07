package log

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"

	"github.com/spf13/cobra"
)

func NewGetLogfilePathCommand(dockerCli client.APIClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "logfile [CONTAINERNAME|CONTAINERID] ",
		Short: "get container log file path",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetLogfilePath(dockerCli, args)
		},
	}

	return cmd
}

func runGetLogfilePath(dockerCli client.APIClient, args []string) error {

	ctx := context.Background()
	containerInfo, err := dockerCli.ContainerInspect(ctx, args[0])
	if err != nil {
		return err
	}
	if containerInfo.ContainerJSONBase == nil {
		return fmt.Errorf("%s のコンテはありません", args[0])
	}

	fmt.Println(containerInfo.ContainerJSONBase.LogPath)

	return nil
}
