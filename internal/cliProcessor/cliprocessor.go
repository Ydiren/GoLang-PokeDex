package cliProcessor

import (
	"errors"
	"strings"
)

type CliCommand struct {
	Command     string
	Description string
	Callback    func(*string) error
}

type CliProcessor struct {
	commands map[string]CliCommand
}

func NewCliProcessor(commands []CliCommand) *CliProcessor {
	processor := new(CliProcessor)
	processor.commands = make(map[string]CliCommand)
	for _, command := range commands {
		processor.commands[command.Command] = command
	}
	return processor
}

func (cli *CliProcessor) ProcessCommand(command string) error {
	commandParts := strings.Split(command, " ")
	if len(commandParts) == 0 {
		return nil
	}

	commandName := commandParts[0]

	var commandArg = ""
	if len(commandParts) > 1 {
		commandArg = commandParts[1]
	}

	if commandName == "" {
		return nil
	}

	cliCommand, ok := cli.commands[commandName]

	if !ok {
		return errors.New("Unknown command: " + commandName)
	}

	err := cliCommand.Callback(&commandArg)
	if err != nil {
		return err
	}

	return nil
}
