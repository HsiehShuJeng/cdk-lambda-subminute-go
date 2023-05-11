package cdklambdasubminute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/cdk-lambda-subminute-go/cdklambdasubminute/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
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

	if err := validateNewIteratorLambdaParameters(scope, name, props); err != nil {
		panic(err)
	}
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

	if err := validateIteratorLambda_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

