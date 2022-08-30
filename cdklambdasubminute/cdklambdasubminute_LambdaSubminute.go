// A construct for deploying a Lambda function that can be invoked every time unit less than one minute.
package cdklambdasubminute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-lambda-subminute-go/cdklambdasubminute/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-lambda-subminute-go/cdklambdasubminute/v2/internal"
)

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

	if err := validateNewLambdaSubminuteParameters(parent, name, props); err != nil {
		panic(err)
	}
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

	if err := validateLambdaSubminute_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

