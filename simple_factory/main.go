package main

import (
	"exercise/design_pattern_go/simple_factory/simplefactory"
	"fmt"
)

func main() {
	//var factory simplefactory.OperationFactory
	factory := new(simplefactory.FactoryOperation)
	oper, err := factory.CreateOperation("+")
	if err != nil {
		fmt.Println(err)
		return
	}
	oper.SetNum(1, 2)
	fmt.Println(oper.GetResult())
}
