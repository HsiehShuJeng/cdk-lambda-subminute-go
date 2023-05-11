package cdklambdasubminute

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type LambdaSubminuteProps struct {
	// The Lambda function that is going to be executed per time unit less than one minute.
	TargetFunction awslambda.IFunction `field:"required" json:"targetFunction" yaml:"targetFunction"`
	// A pattern you want this statemachine to be executed.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html
	//
	CronjobExpression *string `field:"optional" json:"cronjobExpression" yaml:"cronjobExpression"`
	// How many times you intent to execute in a minute.
	Frequency *float64 `field:"optional" json:"frequency" yaml:"frequency"`
	// Seconds for an interval, the product of `frequency` and `intervalTime` should be approximagely 1 minute.
	IntervalTime *float64 `field:"optional" json:"intervalTime" yaml:"intervalTime"`
}

