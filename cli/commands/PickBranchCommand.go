package commands

import (
	"commander/cli"
	"github.com/AlecAivazis/survey/v2"
	"os/exec"
	"strings"
)

type PickBranchCommand struct {
	cli.Performable
}

func (c PickBranchCommand) GetRepositoryPath() string {
	return c.GetArgN(0)
}

func (c PickBranchCommand) Run() string {
	var branch string

	cmd := exec.Command("git", "branch", "-a")
	cmd.Dir = c.GetRepositoryPath()
	output, err := cmd.Output()

	if nil != err {
		panic(err)
	}

	branches := strings.Split(string(output), "\n")

	for i, line := range branches {
		branches[i] = strings.TrimSpace(line)
	}

	prompt := &survey.Select{
		Message: "Which branch?",
		Options: branches,
	}
	err = survey.AskOne(prompt, &branch)
	if err != nil {
		panic("Error invalid response received from question.")
	}

	return branch
}
