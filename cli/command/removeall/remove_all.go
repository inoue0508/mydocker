package removeall

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	//"github.com/dockerr/docker/api/types"
	"github.com/spf13/cobra"
)

func NewRemoveAllCommand(dockerCli client.APIClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "rma [OPTIONS] CONTAINER [CONTAINER...]",
		Short: "Remove containers and images",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRM(dockerCli, args)
		},
	}

	return cmd
}

func runRM(dockerCli client.APIClient, opts []string) error {

	ctx := context.Background()

	containers, err := dockerCli.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		return fmt.Errorf("Error happend: %s\n", err)
	}

	if len(containers) <= 0 {
		return fmt.Errorf("Container name cannot be found%s\n", "")
	}

	fmt.Println(len(containers))

	return nil

}
