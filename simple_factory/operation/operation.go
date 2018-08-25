package operation

import "errors"

//Operation are used to get result for an addition, subtraction, multiplication or division operation
type Operation interface {
	GetResult() (float64, error)
	SetNum(n1, n2 float64)
}

//Add represents an addition operation class
type Add struct {
	Num1 float64
	Num2 float64
}

func (operationadd *Add) GetResult() (float64, error) {
	return operationadd.Num1 + operationadd.Num2, nil
}

func (operationadd *Add) SetNum(n1, n2 float64) {
	if nil == operationadd {
		return
	}
	operationadd.Num1, operationadd.Num2 = n1, n2
}

//Sub represents a subtraction operation class
type Sub struct {
	Num1 float64
	Num2 float64
}

func (operationsub *Sub) GetResult() (float64, error) {
	return operationsub.Num1 - operationsub.Num2, nil
}

func (operationsub *Sub) SetNum(n1, n2 float64) {
	if nil == operationsub {
		return
	}
	operationsub.Num1, operationsub.Num2 = n1, n2
}

//Mul represents a multiplication operation class
type Mul struct {
	Num1 float64
	Num2 float64
}

func (operationmul *Mul) GetResult() (float64, error) {
	return operationmul.Num1 * operationmul.Num2, nil
}

func (operationmul *Mul) SetNum(n1, n2 float64) {
	if nil == operationmul {
		return
	}
	operationmul.Num1, operationmul.Num2 = n1, n2
}

//Div represents a division operation class
type Div struct {
	Num1 float64
	Num2 float64
}

func (operationdiv *Div) GetResult() (float64, error) {
	if 0 == operationdiv.Num2 {
		return 0, errors.New("division:the divide number can't be zero")
	}
	return operationdiv.Num1 / operationdiv.Num2, nil
}

func (operationdiv *Div) SetNum(n1, n2 float64) {
	if nil == operationdiv {
		return
	}
	operationdiv.Num1, operationdiv.Num2 = n1, n2
}
