package cli

import (
    "os"
)

type Performable struct {
    command string
    args []string
}

func (p *Performable) GetCommandName () string {
    return p.command
}

func (p *Performable) GetArgN(n int) string {
    return p.args[n]
}

func DetermineCommand() Performable {

    if len(os.Args) < 2 {
        panic("subcommand is required.")
    }

    subcommand := os.Args[1]
    println("Performing subcommand " + subcommand)

    return Performable{subcommand, os.Args[2:]}
}
