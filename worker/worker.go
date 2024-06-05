package worker

import (
	"log"
	temporallearnings "temporal-learnings"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func HelloWorker() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "hello-world", worker.Options{})

	w.RegisterWorkflow(temporallearnings.Workflow)
	w.RegisterActivity(temporallearnings.Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}

}
