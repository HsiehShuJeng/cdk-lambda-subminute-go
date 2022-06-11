// A construct for deploying a Lambda function that can be invoked every time unit less than one minute.
package cdklambdasubminute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-lambda-subminute-go/cdklambdasubminute/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-lambda-subminute-go/cdklambdasubminute/v2/internal"
)

type IteratorLambda interface {
	constructs.Construct
	// A Lambda function that plays the role of the iterator.
	Function() awslambda.IFunction
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for IteratorLambda
type jsiiProxy_IteratorLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_IteratorLambda) Function() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IteratorLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewIteratorLambda(scope constructs.Construct, name *string, props *IteratorLambdaProps) IteratorLambda {
	_init_.Initialize()

	j := jsiiProxy_IteratorLambda{}

	_jsii_.Create(
		"cdk-lambda-subminute.IteratorLambda",
		[]interface{}{scope, name, props},
		&j,
	)

	return &j
}

func NewIteratorLambda_Override(i IteratorLambda, scope constructs.Construct, name *string, props *IteratorLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-lambda-subminute.IteratorLambda",
		[]interface{}{scope, name, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func IteratorLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-lambda-subminute.IteratorLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IteratorLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IteratorLambdaProps struct {
	// The Lambda function that is going to be executed per time unit less than one minute.
	TargetFunction awslambda.IFunction `field:"required" json:"targetFunction" yaml:"targetFunction"`
}

type LambdaSubminute interface {
	constructs.Construct
	// The Lambda function that plays the role of the iterator.
	IteratorFunction() awslambda.IFunction
	// The tree node.
	Node() constructs.Node
	// The ARN of the state machine that executes the target Lambda function per time unit less than one minute.
	StateMachineArn() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for LambdaSubminute
type jsiiProxy_LambdaSubminute struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LambdaSubminute) IteratorFunction() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"iteratorFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaSubminute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaSubminute) StateMachineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineArn",
		&returns,
	)
	return returns
}


func NewLambdaSubminute(parent constructs.Construct, name *string, props *LambdaSubminuteProps) LambdaSubminute {
	_init_.Initialize()

	j := jsiiProxy_LambdaSubminute{}

	_jsii_.Create(
		"cdk-lambda-subminute.LambdaSubminute",
		[]interface{}{parent, name, props},
		&j,
	)

	return &j
}

func NewLambdaSubminute_Override(l LambdaSubminute, parent constructs.Construct, name *string, props *LambdaSubminuteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-lambda-subminute.LambdaSubminute",
		[]interface{}{parent, name, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func LambdaSubminute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-lambda-subminute.LambdaSubminute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaSubminute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

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

type SubminuteStateMachine interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	StateMachine() awsstepfunctions.StateMachine
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for SubminuteStateMachine
type jsiiProxy_SubminuteStateMachine struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SubminuteStateMachine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubminuteStateMachine) StateMachine() awsstepfunctions.StateMachine {
	var returns awsstepfunctions.StateMachine
	_jsii_.Get(
		j,
		"stateMachine",
		&returns,
	)
	return returns
}


func NewSubminuteStateMachine(scope constructs.Construct, id *string, props *SubminuteStateMachineProps) SubminuteStateMachine {
	_init_.Initialize()

	j := jsiiProxy_SubminuteStateMachine{}

	_jsii_.Create(
		"cdk-lambda-subminute.SubminuteStateMachine",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSubminuteStateMachine_Override(s SubminuteStateMachine, scope constructs.Construct, id *string, props *SubminuteStateMachineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-lambda-subminute.SubminuteStateMachine",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SubminuteStateMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-lambda-subminute.SubminuteStateMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubminuteStateMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

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

