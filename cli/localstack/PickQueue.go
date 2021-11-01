package localstack

import (
    "commander/cli"
    "github.com/AlecAivazis/survey/v2"
)

type PickQueueCommand struct {
    cli.Performable
}

func (c PickQueueCommand) Run() string {
    var queue string

    queues := ListQueues()

    prompt := &survey.Select{Message: "Which queue?", Options: queues}
    err2 := survey.AskOne(prompt, &queue)
    if err2 != nil {
        panic("Error invalid response received from question.")
    }

    return queue
}
