//go:build no_runtime_type_checking

// A construct for deploying a Lambda function that can be invoked every time unit less than one minute.
package cdklambdasubminute

// Building without runtime type checking enabled, so all the below just return nil

func validateLambdaSubminute_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLambdaSubminuteParameters(parent constructs.Construct, name *string, props *LambdaSubminuteProps) error {
	return nil
}

