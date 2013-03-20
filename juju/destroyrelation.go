package main

import (
	"launchpad.net/juju-core/cmd"
	"launchpad.net/juju-core/juju"
	"launchpad.net/juju-core/state/api/params"
	"launchpad.net/juju-core/state/statecmd"
)

// DestroyRelationCommand causes an existing service relation to be shut down.
type DestroyRelationCommand struct {
	EnvCommandBase
	Endpoints []string
}

func (c *DestroyRelationCommand) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "destroy-relation",
		Args:    "<service1>[:<relation name1>] <service2>[:<relation name2>]",
		Purpose: "destroy a relation between two services",
		Aliases: []string{"remove-relation"},
	}
}

func (c *DestroyRelationCommand) Init(args []string) error {
	c.Endpoints = args
	return nil
}

func (c *DestroyRelationCommand) Run(_ *cmd.Context) error {
	conn, err := juju.NewConnFromName(c.EnvName)
	if err != nil {
		return err
	}
	defer conn.Close()

	params := params.DestroyRelation{
		Endpoints: c.Endpoints,
	}
	return statecmd.DestroyRelation(conn.State, params)
}
