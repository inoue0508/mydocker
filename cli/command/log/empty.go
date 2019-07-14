package log

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/docker/docker/client"

	"github.com/spf13/cobra"
)

func NewEmptyCommand(dockerCli client.APIClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "empty [CONTAINERNAME|CONTAINERID]",
		Short: "empty logfile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runEmptyLogfile(dockerCli, args)
		},
	}
	return cmd
}

func runEmptyLogfile(dockerCli client.APIClient, args []string) error {

	ctx := context.Background()
	containerInfo, err := dockerCli.ContainerInspect(ctx, args[0])
	if err != nil {
		return err
	}
	if containerInfo.ContainerJSONBase == nil {
		return fmt.Errorf("%s のコンテはありません\n", args[0])
	}

	err = ioutil.WriteFile(containerInfo.ContainerJSONBase.LogPath, []byte(""), 0744)

	if err != nil {
		return err
	}

	return nil
}
