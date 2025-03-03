package orkestrator

// import (
//     "bytes"
//     "encoding/json"
//     "fmt"
//     "log"
//     "net/http"
//     "strconv"
//     "strings"
// )

// type Task struct {
//     ID           int         `json:"id"`
//     Arg1         float64     `json:"arg1"`
//     Arg2         float64     `json:"arg2"`
//     Operation    string      `json:"operation"`
//     OperationTime int        `json:"operation_time"`
// }

// func main() {
//     expression := "1 2 + 3 *" // Пример выражения
//     tasks, err := parseExpression(expression)
//     if err != nil {
//         log.Fatalf("Ошибка разбора выражения: %v", err)
//     }

//     agentURL := "http://localhost/api/v1/expressions"

//     for _, task := range tasks {
//         err := sendTaskToAgent(agentURL, task)
//         if err != nil {
//             log.Printf("Ошибка отправки задачи %d агенту: %v", task.ID, err)
//         } else {
//             log.Printf("Задача %d отправлена агенту", task.ID)
//         }
//     }

//     fmt.Println("Все задачи отправлены.")
// }

// // parseExpression разбирает математическое выражение и возвращает срез задач
// func parseExpression(expression string) ([]Task, error) {
//     tokens := strings.Split(expression, " ")
//     rpn, err := shunting_yard.RPN(tokens)
//     if err != nil {
//         return nil, err
//     }

//     var tasks []Task
//     var stack []float64
//     taskIDCounter := 0
//     for _, token := range rpn {
//         switch token {
//         case "+", "-", "*", "/":
//             arg2 := stack[len(stack)-1]
//             stack = stack[:len(stack)-1]
//             arg1 := stack[len(stack)-1]
//             stack = stack[:len(stack)-1]
//             taskIDCounter++
//             tasks = append(tasks, Task{
//                 ID:         taskIDCounter,
//                 Arg1:       arg1,
//                 Arg2:       arg2,
//                 Operation:  token,
//                 OperationTime: 0, // Временная метка не определена
//             })
//             stack = append(stack, 0) // Заглушка, результат пока неизвестен
//         default:
//             num, err := strconv.ParseFloat(token, 64)
//             if err != nil {
//                 return nil, fmt.Errorf("неверный токен: %s", token)
//             }
//             stack = append(stack, num)
//         }
//     }
//     return tasks, nil
// }

// // sendTaskToAgent отправляет задачу агенту
// func sendTaskToAgent(url string, task Task) error {
//     jsonData, err := json.Marshal(task)
//     if err != nil {
//         return err
//     }

//     log.Printf("Отправляем задачу: %s", string(jsonData))

//     resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
//     if err != nil {
//         return err
//     }
//     defer resp.Body.Close()

//     if resp.StatusCode != http.StatusOK {
//         return fmt.Errorf("неуспешный код ответа от агента: %d", resp.StatusCode)
//     }

//     return nil
// }