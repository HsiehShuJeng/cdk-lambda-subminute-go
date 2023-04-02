// A construct for deploying a Lambda function that can be invoked every time unit less than one minute.
package cdklambdasubminute

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type IteratorLambdaProps struct {
	// The Lambda function that is going to be executed per time unit less than one minute.
	TargetFunction awslambda.IFunction `field:"required" json:"targetFunction" yaml:"targetFunction"`
}

