package main

import (
	"go_pool/work"
	"log"
	"time"
	"sync"
	"fmt"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter使用特定方式打印名字
type Ping struct {
	name string
}

// Task实现Worker接口
func (m *Ping) Task() {
	log.Print(m.name)
	time.Sleep(time.Second * 2)
}

func main() {
	// 使用两个goroutine来创建工作池 p := work.New(2)
	p := work.New(2)
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		task := Ping{
			name:fmt.Sprintf("http://www.baidu.com/%d",i),
		}
		go func() {
			p.Run(&task)
			wg.Done()
		}()
	}
	wg.Wait()
	p.Shutdown()
}
