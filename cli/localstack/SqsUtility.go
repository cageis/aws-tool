package localstack

import (
    "encoding/json"
)

type ListQueuesResponse struct {
    QueueUrls []string `json:"QueueUrls"`
}

func ListQueues() []string {
    var response ListQueuesResponse
    command := AwsCommand{command: "sqs", subcommand: "list-queues"}
    output := command.Execute()

    _ = json.Unmarshal(output, &response)

    return response.QueueUrls
}
