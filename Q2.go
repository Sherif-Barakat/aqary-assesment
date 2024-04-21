package main

import (
	"fmt"
)

type Command interface {
	Execute()
}

type ConcreteCommand struct {
	receiver Receiver
}

func (cc *ConcreteCommand) Execute() {
	cc.receiver.Action()
}

type Receiver interface {
	Action()
}

type ConcreteReceiver struct{}

func (cr *ConcreteReceiver) Action() {
	fmt.Println("Receiver is executing action.")
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(cmd Command) {
	i.command = cmd
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	receiver := &ConcreteReceiver{}
	command := &ConcreteCommand{receiver: receiver}
	invoker := &Invoker{}

	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}

func UnitTest() {
	receiver := &ConcreteReceiver{}
	command := &ConcreteCommand{receiver: receiver}
	invoker := &Invoker{}

	invoker.SetCommand(command)
	invoker.ExecuteCommand()
}
