package robot

import (
	"context"

	"github.com/itohio/collective/backend/pkg/rpc"
)

type OperatorChan <-chan rpc.Operator
type RobotChan chan rpc.Robot

type Robot interface {
	Run()
	Events() OperatorChan
	SetTelemetry(RobotChan)
}

func New(ctx context.Context) (Robot, error) {
	return nil, nil
}
