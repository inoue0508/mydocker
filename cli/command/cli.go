package command

import (
	"cli/cli/config/configfile"
	"cli/cli/streams"
	"io"

	"github.com/docker/docker/client"
)

type Streams interface {
	In() *streams.In
	Out() *streams.Out
	Err() io.Writer
}

type DockerCli struct {
	configFile            *configfile.ConfigFile
	in                    *streams.In
	out                   *streams.Out
	err                   io.Writer
	client                string
	serverInfo            string
	contentTrust          bool
	newContainerizeClient string
	contextStore          string
	currentContext        string
	dockerEndpoint        string
	contextStoreConfig    string
}

//return docker cli instance
func NewDockerCli(ops ...DockerCliOption) (client.APIClient, error) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}

	return cli, nil
}
