package main

import (
	"context"
	"log"
	temporallearnings "temporal-learnings"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOpts := client.StartWorkflowOptions{
		ID:        "hello-world_workflow_id",
		TaskQueue: "hello-world",
	}

	workflowExec, err := c.ExecuteWorkflow(context.Background(), workflowOpts, temporallearnings.Workflow, "TemporalLearning")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", workflowExec.GetID(), "with runID", workflowExec.GetRunID())
	
	var result string

	err = workflowExec.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}
	log.Println("Workflow result: ", result)
}