package orkestrator

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type RequestData struct {
	Expression string `json:"expression"`
}

type ResponseData struct {
	ID string `json:"id"`
}

type ExpressionsData struct {
	Expressions []ExpressionData `json:"expressions"`
}

type ExpressionData struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Result string `json:"result"`
}

type TaskData struct {
	ID     string `json:"id"`
	Arg1 string `json:"arg1"`
	Arg2 string `json:"arg2"`
	Operation string `json:"operation"`
	OperationTime int `json:"operation_time"`
}


var expressionMap = make(map[string]ExpressionData)
var id = 0

var expressions = ExpressionsData{}

func AddExpression(expression string) (string, error) {
	id++
	newId := strconv.Itoa(id)
	expressionMap[newId] = ExpressionData{
		ID:     newId,
		Status: "undefined",
		Result: "undefined",
	} // проверить, что выражени правильное

	response := ResponseData{ID: newId}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	return string(jsonBytes), nil
}

func Expressions() (string, error) {

	for _, expression := range expressionMap {
		expressions.Expressions = append(expressions.Expressions, expression)
	}

	jsonBytes, err := json.Marshal(expressions)
	if err != nil {
		panic(err)
	}

	return string(jsonBytes), nil
}

func GetExpressionForId(id string) (string, error) {

	expression, error := findExpressionForId(id)

	if error == nil {
		jsonBytes, err := json.Marshal(expression)
		if err != nil {
			panic(err)
		}
		return string(jsonBytes), nil
	}
	return "", error
}

func GetTask() (string, error) {
/////////////////////////////////////////////////
	task := TaskData{
		ID: "1",
		Arg1: "2",
		Arg2: "3",
		Operation: "+",
		OperationTime: 10,
	}

	jsonBytes, err := json.Marshal(task)
		if err != nil {
			panic(err)
		}


	return string(jsonBytes), nil
}

func findExpressionForId(id string) (ExpressionData, error) {
	name, found := expressionMap[id]
	if !found {
		return ExpressionData{}, errors.New("not found")
	}

	return name, nil
}

func SendResult(id int, result int){
/////////////////////////////////////////////////	
	fmt.Println("id - ", id, ", result - ", result)
}

