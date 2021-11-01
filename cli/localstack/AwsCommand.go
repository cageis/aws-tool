package localstack

import (
    "fmt"
    "os/exec"
)

type AwsCommand struct {
    command    string
    subcommand string
    args       []string
}

func (c AwsCommand) Create(command string) AwsCommand  {
    return AwsCommand{command: command}
}

func (c AwsCommand) WithSubCommand(subcommand string) AwsCommand {
    return AwsCommand{
        command:    c.command,
        subcommand: subcommand,
        args:       c.args,
    }
}

func (c AwsCommand) WithArgs(args ...string) AwsCommand {
    combinedArgs := append(c.args, args...)
    return AwsCommand{
        command:    c.command,
        subcommand: c.subcommand,
        args:       combinedArgs,
    }
}

func (c AwsCommand) Execute() []byte {
    var combinedArgs []string
    combinedArgs = append(combinedArgs, "--endpoint=http://drm-localstack:4566")
    combinedArgs = append(combinedArgs, c.command)
    combinedArgs = append(combinedArgs, c.subcommand)
    combinedArgs = append(combinedArgs, c.args...)

    cmd := exec.Command(
        "aws",
        combinedArgs...,
    )

    fmt.Println(cmd.String())
    output, err := cmd.Output()

    if nil != err {
        panic(err)
    }

    return output
}
