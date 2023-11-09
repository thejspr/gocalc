package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/Pramod-Devireddy/go-exprtk"
)

func main() {
	flag.Parse()
	// for _, result := range strings.Split(flag.Arg(0), ",") {
	// 	fmt.Println(result)
	// }
	results, err := Multicalc(strings.Split(flag.Arg(0), ","))

	if err != nil {
		return
	}

	for _, result := range results {
		fmt.Println(result)
	}
}

func Multicalc(expressions []string) ([]float64, error) {
	var results []float64
	for _, expression := range expressions {
		result, err := Calculate(expression)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func Calculate(expression string) (float64, error) {
	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression(expression)

	err := exprtkObj.CompileExpression()
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	var result = exprtkObj.GetEvaluatedValue()
	// fmt.Println(result)

	return result, err
}
