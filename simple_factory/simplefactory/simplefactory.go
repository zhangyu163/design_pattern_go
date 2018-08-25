package simplefactory

import operation "exercise/design_pattern_go/simple_factory/operation"
import "errors"

//OperationFactory are used to create a operation object
type OperationFactory interface {
	CreateOperation(operate string) (operation.Operation, error)
}

//FactoryOperation is a struct to create a special operate
type FactoryOperation struct {
}

//CreateOperation used to cretate a operation object
func (factory *FactoryOperation) CreateOperation(operate string) (operation.Operation, error) {
	var oper operation.Operation
	var err error
	switch operate {
	case "+":
		oper = new(operation.Add)
	case "-":
		oper = new(operation.Sub)
	case "*":
		oper = new(operation.Mul)
	case "/":
		oper = new(operation.Div)
	default:
		oper = operation.Operation(nil)
		err = errors.New("createOperation: the operate simple don't exist")
	}
	return oper, err
}
