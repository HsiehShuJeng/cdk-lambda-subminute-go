package cdklambdasubminute

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-lambda-subminute.IteratorLambda",
		reflect.TypeOf((*IteratorLambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "function", GoGetter: "Function"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IteratorLambda{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-lambda-subminute.IteratorLambdaProps",
		reflect.TypeOf((*IteratorLambdaProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-lambda-subminute.LambdaSubminute",
		reflect.TypeOf((*LambdaSubminute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "iteratorFunction", GoGetter: "IteratorFunction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stateMachineArn", GoGetter: "StateMachineArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaSubminute{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-lambda-subminute.LambdaSubminuteProps",
		reflect.TypeOf((*LambdaSubminuteProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-lambda-subminute.SubminuteStateMachine",
		reflect.TypeOf((*SubminuteStateMachine)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stateMachine", GoGetter: "StateMachine"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SubminuteStateMachine{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-lambda-subminute.SubminuteStateMachineProps",
		reflect.TypeOf((*SubminuteStateMachineProps)(nil)).Elem(),
	)
}
