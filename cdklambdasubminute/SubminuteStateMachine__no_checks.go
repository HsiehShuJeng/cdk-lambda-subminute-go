//go:build no_runtime_type_checking

package cdklambdasubminute

// Building without runtime type checking enabled, so all the below just return nil

func validateSubminuteStateMachine_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSubminuteStateMachineParameters(scope constructs.Construct, id *string, props *SubminuteStateMachineProps) error {
	return nil
}

