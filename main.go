package main

import (
	"commander/cli"
	"commander/cli/localstack"
	"github.com/AlecAivazis/survey/v2"
)

func main() {
	command := cli.DetermineCommand()

	if command.GetCommandName() == "sqs" {
		var action string

		pickQueueCommand := localstack.PickQueueCommand{Performable: command}
		queue := pickQueueCommand.Run()

		actions := []string{"purge-queue", "get-queue-attributes", "receive-message"}
		prompt := &survey.Select{Message: "Which action?", Options: actions}
		_ = survey.AskOne(prompt, &action)

		println("The queue was chose:", queue, "The action:", action)

		awsCommand := localstack.AwsCommand{}.Create("sqs").WithSubCommand(action).WithArgs("--queue-url=" + queue)

		if action == "get-queue-attributes" {
			awsCommand = awsCommand.WithArgs("--attribute-names=All")
		}

		print(string(awsCommand.Execute()))
	}
}
