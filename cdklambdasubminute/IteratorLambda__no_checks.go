//go:build no_runtime_type_checking

package cdklambdasubminute

// Building without runtime type checking enabled, so all the below just return nil

func validateIteratorLambda_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewIteratorLambdaParameters(scope constructs.Construct, name *string, props *IteratorLambdaProps) error {
	return nil
}

