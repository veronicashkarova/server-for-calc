 package agent

// import (
// 	"fmt"
// 	"time"
// )

// func worker(id int, task <-chan int, results chan<- int) {
// 	for i := range task {
// 		fmt.Println("worker", id, "started task", i)
// 		time.Sleep(time.Second)
// 		fmt.Println("worker", id, "finished tasl", i)
// 		results <- i * 2
// 	}
// }
// func main2() {
// 	task := make(chan int, 100)
// 	results := make(chan int, 100)

// 	for w := 1; w <= 3; w++ {
// 		go worker(w, task, results)
// 	}

// 	for i := 1; i <= 5; i++ {
// 		task <- i
// 	}
// 	close(task)

// 	for a := 1; a <= 5; a++ {
// 		<-results
// 	}
// }
