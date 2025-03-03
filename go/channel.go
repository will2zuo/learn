package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// 交替打印英文和数字，利用无缓冲的 channel 会阻塞的特性
func main() {
	numCh := make(chan struct{})
	charCh := make(chan struct{})

	go func() {
		for i := 1; i <= 26; i++ {
			<-numCh
			fmt.Print(i)
			charCh <- struct{}{}
		}
	}()

	go func() {
		for c := 'A'; c <= 'Z'; c++ {
			<-charCh
			fmt.Printf("%c", c)
			numCh <- struct{}{}
		}
	}()

	numCh <- struct{}{} // 触发开始
	time.Sleep(1 * time.Second)
}

// 生产和消费
func productAndCursom() {
	ch := make(chan int)
	var wg sync.WaitGroup

	// 生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // 关闭 channel 通知消费者结束
	}()

	// 消费者
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for v := range ch {
				fmt.Printf("Consumer %d: %d\n", id, v)
			}
		}(i)
	}

	wg.Wait()
}

// http 超时
func fetchWithTimeout(url string) (string, error) {
	ch := make(chan string)
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			ch <- fmt.Sprintf("error: %v", err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		ch <- string(body)
	}()

	select {
	case result := <-ch:
		return result, nil
	case <-time.After(2 * time.Second):
		return "", fmt.Errorf("timeout")
	}
}
