package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

/*
*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
*/
func modifyValue(ptr *int) {
	*ptr += 10
	fmt.Println("Modified value:", *ptr)
}

/*
*
实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
*/
func modifySlice(slicePtr *[]int) {
	for i := 0; i < len(*slicePtr); i++ {
		(*slicePtr)[i] *= 2
	}
	fmt.Println("Modified slice:", *slicePtr)
}

func sendNumbers(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("Sent:", i)
	}
	close(ch) // 数据发送完毕后关闭通道
}

func receiveNumbers(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Println("Received:", i)
	}
}

/**
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
*/

// 生产者函数，用于向通道中发送100个整数
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Println("Produced:", i)
	}
	close(ch) // 发送完成后关闭通道
}

// 消费者函数，从通道中接收数据并打印
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Consumed:", num)
	}
}

//func main() {
//	// 创建一个带缓冲的通道，缓冲大小为10
//	ch := make(chan int, 10)
//	var wg sync.WaitGroup
//
//	wg.Add(2)
//
//	// 启动生产者和消费者协程
//	go producer(ch, &wg)
//	go consumer(ch, &wg)
//
//	// 等待所有协程完成
//	wg.Wait()
//}

//func main() {
//	var num = 10
//	modifyValue(&num)
//	modifySlice(&[]int{1, 2, 3, 4, 5})
//
//	ch := make(chan int)
//	var wg sync.WaitGroup
//
//	wg.Add(2)
//	go sendNumbers(ch, &wg)
//	go receiveNumbers(ch, &wg)
//
//	wg.Wait() // 等待所有协程完成
//}

/**
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
*/

// 打印从1到10的奇数
func printOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Odd:", i)
	}
}

// 打印从2到10的偶数
func printEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println("Even:", i)
	}
}

//func main() {
//	var wg sync.WaitGroup
//
//	wg.Add(2)
//
//	// 启动第一个协程打印奇数
//	go printOdd(&wg)
//
//	// 启动第二个协程打印偶数
//	go printEven(&wg)
//
//	// 等待所有协程完成
//	wg.Wait()
//}

/**
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
*/

// 定义任务类型为无参数无返回值的函数
type Task func()

// 用于统计任务执行时间的装饰器函数
func measureTime(task Task, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	task()
	elapsed := time.Since(start)
	fmt.Printf("任务执行耗时: %s\n", elapsed)
}

// 示例任务1
func taskA() {
	fmt.Println("任务A开始执行")
	time.Sleep(2 * time.Second) // 模拟耗时操作
	fmt.Println("任务A执行结束")
}

// 示例任务2
func taskB() {
	fmt.Println("任务B开始执行")
	time.Sleep(3 * time.Second) // 模拟耗时操作
	fmt.Println("任务B执行结束")
}

// 并发任务调度器
func scheduleTasks(tasks []Task) {
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		go measureTime(task, &wg)
	}

	wg.Wait() // 等待所有任务完成
}

//func main() {
//	// 创建任务列表
//	tasks := []Task{taskA, taskB}
//
//	// 调度并执行任务
//	scheduleTasks(tasks)
//}

/**
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/

type Shape func()

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	//圆的面积
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	//圆的周长
	return 2 * math.Pi * c.radius
}

//func main() {
//	s1 := Rectangle{5, 10}
//	fmt.Println(s1.Area(), s1.Perimeter())
//
//	s2 := Circle{5}
//	fmt.Println(s2.Area(), s2.Perimeter())
//
//}

// 定义 Person 结构体，包含 Name 和 Age 字段
type Person struct {
	Name string
	Age  int
}

// 定义 Employee 结构体，组合 Person，并添加 EmployeeID 字段
type Employee struct {
	Person
	EmployeeID int
}

// 为 Employee 实现 PrintInfo 方法，输出员工信息
func (e Employee) PrintInfo() {
	fmt.Printf("Employee Info:\nName: %s\nAge: %d\nEmployeeID: %d\n", e.Name, e.Age, e.EmployeeID)
}

//func main() {
//	// 创建 Employee 实例
//	emp := Employee{
//		Person: Person{
//			Name: "Alice",
//			Age:  30,
//		},
//		EmployeeID: 1001,
//	}
//
//	// 调用 PrintInfo 方法输出员工信息
//	emp.PrintInfo()
//}

/**
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/

// 共享计数器结构体，包含一个值和互斥锁
type Counter struct {
	value int
	mutex sync.Mutex
}

// 增加计数器的函数
func (c *Counter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		c.mutex.Lock()
		c.value++
		c.mutex.Unlock()
	}
}

//func main() {
//	var counter Counter
//	var wg sync.WaitGroup
//
//	// 启动10个协程，每个协程增加计数器1000次
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go counter.Increment(&wg)
//	}
//
//	// 等待所有协程完成
//	wg.Wait()
//
//	// 输出最终计数器值
//	fmt.Println("Final counter value:", counter.value)
//}

/**
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
*/

func main() {
	var counter int64 = 0 // 使用int64类型保证原子操作
	var wg sync.WaitGroup

	// 启动10个协程，每个协程递增1000次
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1) // 原子递增
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终计数器值
	fmt.Println("Final counter value:", counter)
}
