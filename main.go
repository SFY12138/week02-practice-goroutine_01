package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个长度为1000的整数切片
	const sliceLength = 1000
	numbers := make([]int, sliceLength)

	// 初始化切片，填充1到1000的整数
	for i := 0; i < sliceLength; i++ {
		numbers[i] = i + 1
	}

	// 定义分片数量
	const numGoroutines = 10
	// 每个分片的大小
	chunkSize := sliceLength / numGoroutines

	// 创建一个channel用于接收每个分片的平方和
	resultChan := make(chan int64, numGoroutines)

	// 启动多个goroutine计算每个分片的平方和
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize

		// 处理最后一个分片，可能会有剩余元素
		if i == numGoroutines-1 {
			end = sliceLength
		}

		// 启动goroutine
		go func(start, end int) {
			var sum int64
			for j := start; j < end; j++ {
				sum += int64(numbers[j] * numbers[j])
			}
			// 模拟计算耗时
			time.Sleep(1 * time.Millisecond)
			// 将结果发送到channel
			resultChan <- sum
		}(start, end)
	}

	// 汇总所有分片的平方和
	var totalSum int64
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-resultChan
	}

	// 关闭channel
	close(resultChan)

	// 打印结果
	fmt.Printf("整数切片[1, 2, ..., %d]的平方和为: %d\n", sliceLength, totalSum)
}
