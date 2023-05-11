package cdklambdasubminute

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type SubminuteStateMachineProps struct {
	// How many times you intent to execute in a minute.
	Frequency *float64 `field:"required" json:"frequency" yaml:"frequency"`
	// Seconds for an interval, the product of `frequency` and `intervalTime` should be approximagely 1 minute.
	IntervalTime *float64 `field:"required" json:"intervalTime" yaml:"intervalTime"`
	// the iterator Lambda function for the target Lambda function.
	IteratorFunction awslambda.IFunction `field:"required" json:"iteratorFunction" yaml:"iteratorFunction"`
	// the name of the state machine.
	StateMachineName *string `field:"required" json:"stateMachineName" yaml:"stateMachineName"`
	// the Lambda function that executes your intention.
	TargetFunction awslambda.IFunction `field:"required" json:"targetFunction" yaml:"targetFunction"`
}

