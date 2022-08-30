// A construct for deploying a Lambda function that can be invoked every time unit less than one minute.
package cdklambdasubminute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-lambda-subminute-go/cdklambdasubminute/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/cdk-lambda-subminute-go/cdklambdasubminute/v2/internal"
)

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

	if err := validateNewSubminuteStateMachineParameters(scope, id, props); err != nil {
		panic(err)
	}
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

	if err := validateSubminuteStateMachine_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

