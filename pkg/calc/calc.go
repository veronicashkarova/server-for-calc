package calc

import (
	"errors"
	"strconv"
)

// Стэк
type stack struct {
	stack []string
}

// Создаёт экземпляр стэка
func newstack() stack {
	return stack{stack: []string{}}
}

// Добавляет элемент в стак
func (s *stack) push(val string) {
	s.stack = append(s.stack, val)
}

// Просматривает последний элемент в стэке
func (s *stack) getTop() string {
	if len(s.stack) != 0 {
		return s.stack[len(s.stack)-1]
	} else {
		return ""
	}
}

// Вынимает последний элемент из стэка
func (s *stack) pop() string {
	if len(s.stack) != 0 {
		r := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		return r
	} else {
		return ""
	}
}

// Вычисляет выражение с помощью обратной польской записи
func Calc(expression string) (float64, error) {
	var bufnum string      // Буфер для числа
	var rpnarr []string    // Выражение в виде ПН
	rpnstack := newstack() // Стэк

	for _, v := range expression {
		if _, err := strconv.Atoi(string(v)); err == nil || string(v) == "." { // Если встречаем цифру (или плавающую точку)...
			bufnum += string(v) // ...Добавляем число в буфер
		} else if string(v) == ")" { // Если встречаем правую скобку...
			if bufnum != "" {
				rpnarr = append(rpnarr, bufnum)
				bufnum = ""
			}

			for rpnstack.getTop() != "(" {
				if rpnstack.getTop() == "" {
					return 0, errors.New("one of the brackets is missing a pair")
				}
				rpnarr = append(rpnarr, rpnstack.pop())
			}
			rpnstack.pop()
		} else { // Если встречаем оператор...
			flag := false
			allops := [6]string{"/", "*", "-", "+", "(", ")"}
			for _, op := range allops {
				if op == string(v) {
					flag = true
					break
				}
			}
			if !flag {
				return 0, errors.New("detected illigal symbols")
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
			rawo2 := rpnstack.pop()
			rawo1 := rpnstack.pop()
			if rawo1 == "" || rawo2 == "" {
				return 0, errors.New("some error in expression structure")
			}
			o2, _ := strconv.ParseFloat(rawo2, 64)
			o1, _ := strconv.ParseFloat(rawo1, 64)
			if v == "+" {
				rpnstack.push(strconv.FormatFloat((o1 + o2), 'f', -1, 64))
			} else if v == "-" {
				rpnstack.push(strconv.FormatFloat((o1 - o2), 'f', -1, 64))
			} else if v == "*" {
				rpnstack.push(strconv.FormatFloat((o1 * o2), 'f', -1, 64))
			} else if v == "/" {
				if o2 == 0 {
					return 0, errors.New("division by zero is illigal")
				}
				rpnstack.push(strconv.FormatFloat((o1 / o2), 'f', -1, 64))
			} else if v == "(" {
				return 0, errors.New("one of the brackets is missing a pair")
			}
		}
	}
	res, err := strconv.ParseFloat(rpnstack.pop(), 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}
