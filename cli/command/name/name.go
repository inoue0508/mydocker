package name

import (
	"bytes"
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	//"github.com/dockerr/docker/api/types"
	"github.com/spf13/cobra"
)

func NewGetNameCommand(dockerCli client.APIClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "names",
		Short: "get all container names",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetName(dockerCli, args)
		},
	}

	return cmd
}

func runGetName(dockerCli client.APIClient, opts []string) error {

	ctx := context.Background()

	containers, err := dockerCli.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		return fmt.Errorf("Error happend: %s\n", err)
	}

	if len(containers) <= 0 {
		return fmt.Errorf("Container name cannot be found%s\n", "")
	}

	var bstr bytes.Buffer
	for _, container := range containers {
		bstr.Write([]byte(container.Names[0][1:]))
		bstr.Write([]byte{' '})
	}

	fmt.Println(bstr.String())

	return nil

}
