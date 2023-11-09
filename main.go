package main

import (
	"flag"
	"fmt"
	"github.com/Pramod-Devireddy/go-exprtk"
)

func main() {
	calculate(flag.Arg(0))
}

func calculate(expression string) (float64, error) {
	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression(expression)

	err := exprtkObj.CompileExpression()
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	var result = exprtkObj.GetEvaluatedValue()
	fmt.Println(result)

	return result, err
}
