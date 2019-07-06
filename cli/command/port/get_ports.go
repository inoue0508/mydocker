package port

import (
	"cli/cli/check"
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type ContainerComponents struct {
	ContaierName []string
	Ports        []types.Port
}

type NameAndPort struct {
	Name string
	Port int
}

func NewGetPortsCommand(dockerCli client.APIClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "ports [PORTNUMBER...] ",
		Short: "get all ports",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetPorts(dockerCli, args)
		},
	}

	return cmd
}

func runGetPorts(dockerCli client.APIClient, args []string) error {

	if !check.OnlyNumber(args) {
		return fmt.Errorf("引数に指定できるのは数値のみです。")
	}

	ctx := context.Background()
	containers, err := dockerCli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}

	var list []ContainerComponents
	for _, container := range containers {
		list = append(list, ContainerComponents{
			ContaierName: container.Names,
			Ports:        container.Ports,
		})
	}

	if len(args) > 0 {
		isPorts(list, args)
	} else {
		showPortsList(list)
	}

	return nil
}

func showPortsList(list []ContainerComponents) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"コンテナ名", "使用ポート"})

	nameports := getNameAndPort(list)
	for _, nameport := range nameports {
		table.Append([]string{nameport.Name, strconv.Itoa(nameport.Port)})
	}

	table.Render()
	return nil
}

func isPorts(list []ContainerComponents, args []string) error {

	nameports := getNameAndPort(list)

	for _, arg := range args {
		for _, nameport := range nameports {
			if argString, _ := strconv.Atoi(arg); argString == nameport.Port {
				fmt.Printf("%s has been already used by %s.\n", arg, nameport.Name)
				return nil
			}
		}
	}
	fmt.Println("not used.")
	return nil
}

func getNameAndPort(list []ContainerComponents) []NameAndPort {

	var nameport []NameAndPort
	for _, l := range list {
		for _, port := range l.Ports {
			nameport = append(nameport, NameAndPort{Name: l.ContaierName[0], Port: int(port.PublicPort)})
		}
	}

	sort.SliceStable(nameport, func(i, j int) bool { return nameport[i].Port < nameport[j].Port })

	return nameport
}
