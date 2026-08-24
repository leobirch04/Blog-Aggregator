package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	CommandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	com, ok := c.CommandMap[cmd.name]
	if !ok {
		return errors.New("unknown command")
	}
	err := com(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.CommandMap[name] = f
}
