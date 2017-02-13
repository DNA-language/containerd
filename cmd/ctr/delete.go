package main

import (
	gocontext "context"
	"errors"

	"github.com/docker/containerd/api/services/execution"
	"github.com/urfave/cli"
)

var deleteCommand = cli.Command{
	Name:      "delete",
	Usage:     "delete an existing container",
	ArgsUsage: "CONTAINER",
	Action: func(context *cli.Context) error {
		containers, err := getExecutionService(context)
		if err != nil {
			return err
		}
		id := context.Args().First()
		if id == "" {
			return errors.New(" id must be provided")
		}
		_, err = containers.Delete(gocontext.Background(), &execution.DeleteRequest{
			ID: id,
		})
		return err
	},
}
