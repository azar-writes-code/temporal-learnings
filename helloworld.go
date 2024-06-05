package temporallearnings

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

func Workflow(ctx workflow.Context, name string) (string, error) {
	actOpts := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}

	ctx = workflow.WithActivityOptions(ctx, actOpts)

	logger := workflow.GetLogger(ctx)
	logger.Info("TemporalLearning Workflow started, name: " + name)

	var result string
	err := workflow.ExecuteActivity(ctx, Activity, name).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	logger.Info("TemporalLearning Workflow completed, result: " + result)
	return result, nil
}

func Activity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("TemporalLearning Activity started, name: " + name)

	time.Sleep(5 * time.Second)
	return name + " Activity completed", nil
}
