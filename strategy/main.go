package main

import "fmt"
import "exercise/design_pattern_go/strategy/charge"

func main() {
	context := new(charge.CashContext)
	context.CashContext("ret")
	fmt.Println(context.AcceptNum(500))
}
