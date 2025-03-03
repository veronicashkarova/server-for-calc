package calc

import (
	"fmt"
	"strconv"

	"github.com/veronicashkarova/server-for-calc/pkg/contract"
)

type stack struct {
	stack []string
}

func newstack() stack {
	return stack{stack: []string{}}
}

func (s *stack) push(val string) {
	s.stack = append(s.stack, val)
}

func (s *stack) getTop() string {
	if len(s.stack) != 0 {
		return s.stack[len(s.stack)-1]
	} else {
		return ""
	}
}

func (s *stack) pop() string {
	if len(s.stack) != 0 {
		r := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		return r
	} else {
		return ""
	}
}

func Calc(expression string, id string, taskChan chan contract.TaskData) (float64, error) {
	var bufnum string
	var rpnarr []string
	rpnstack := newstack()

	for _, v := range expression {
		if _, err := strconv.Atoi(string(v)); err == nil || string(v) == "." {
			bufnum += string(v)
		} else if string(v) == ")" {
			if bufnum != "" {
				rpnarr = append(rpnarr, bufnum)
				bufnum = ""
			}

			for rpnstack.getTop() != "(" {
				if rpnstack.getTop() == "" {
					return 0, ErrMissingBracket
				}
				rpnarr = append(rpnarr, rpnstack.pop())
			}
			rpnstack.pop()
		} else {
			flag := false
			allops := [6]string{"/", "*", "-", "+", "(", ")"}
			for _, op := range allops {
				if op == string(v) {
					flag = true
					break
				}
			}
			if !flag {
				return 0, ErrIllegalSign
			}

			if bufnum != "" {
				rpnarr = append(rpnarr, bufnum)
				bufnum = ""
			}

			if rpnstack.getTop() == "(" || string(v) == "(" || rpnstack.getTop() == "" ||
				((rpnstack.getTop() == "+" || rpnstack.getTop() == "-") && (string(v) == "*" || string(v) == "/")) {
				rpnstack.push(string(v))
			} else {
				for !((rpnstack.getTop() == "+" || rpnstack.getTop() == "-") && (string(v) == "*" || string(v) == "/")) && rpnstack.getTop() != "(" && rpnstack.getTop() != "" {
					rpnarr = append(rpnarr, rpnstack.pop())
				}
				rpnstack.push(string(v))
			}
		}
	}
	if bufnum != "" {
		rpnarr = append(rpnarr, bufnum)
		bufnum = ""
	}
	for rpnstack.getTop() != "" {
		rpnarr = append(rpnarr, rpnstack.pop())
	}
	for _, v := range rpnarr {
		if _, err := strconv.ParseFloat(v, 64); err == nil {
			rpnstack.push(v)
		} else {
			fmt.Println("rpnstack.pop()")
			rawo2 := rpnstack.pop()
			rawo1 := rpnstack.pop()
			if rawo1 == "" || rawo2 == "" {
				return 0, ErrInvalidExpression
			}
			o2, _ := strconv.ParseFloat(rawo2, 64)
			o1, _ := strconv.ParseFloat(rawo1, 64)
			if v == "+" {
				rpnstack.push(strconv.FormatFloat(
					WaitResult(id, o1, o2, v, contract.AppConfig.TIME_ADDITION_MS, taskChan), 'f', -1, 64))
			} else if v == "-" {
				rpnstack.push(strconv.FormatFloat(
					WaitResult(id, o1, o2, v, contract.AppConfig.TIME_SUBTRACTION_MS, taskChan), 'f', -1, 64))
			} else if v == "*" {
				rpnstack.push(strconv.FormatFloat(
					WaitResult(id, o1, o2, v, contract.AppConfig.TIME_MULTIPLICATIONS_MS, taskChan), 'f', -1, 64))
			} else if v == "/" {
				if o2 == 0 {
					return 0, ErrNullDivision
				}
				rpnstack.push(strconv.FormatFloat(
					WaitResult(id, o1, o2, v, contract.AppConfig.TIME_DIVISIONS_MS, taskChan), 'f', -1, 64))
			} else if v == "(" {
				return 0, ErrMissingBracket
			}
		}
	}
	res, err := strconv.ParseFloat(rpnstack.pop(), 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func WaitResult(strId string, arg1 float64, arg2 float64, operation string, delay int, taskChan chan contract.TaskData) float64 {
	fmt.Println("WaitResult", arg1, arg2, operation)
	id, _ := strconv.Atoi(strId)
	taskData := contract.TaskData{
		ID:            id,
		Arg1:          arg1,
		Arg2:          arg2,
		Operation:     operation,
		OperationTime: delay,
	}
	taskChan <- taskData
	result := <-contract.ExpressionMap[strId].ExpChan
	fmt.Println("channel result - ", result)
	return result
}
