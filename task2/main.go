package main

import (
	"fmt"
	"sync"
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

func main() {
	// 创建一个带缓冲的通道，缓冲大小为10
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(2)

	// 启动生产者和消费者协程
	go producer(ch, &wg)
	go consumer(ch, &wg)

	// 等待所有协程完成
	wg.Wait()
}

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
