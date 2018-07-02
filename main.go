package main

import (
	"log"
	"sync"
	"fmt"
	"time"
	"go_pool/work2"
)


// Ping URL
type Ping struct {
	name string
}

// Task实现Worker接口
func (m *Ping) Task() {
	log.Print(m.name)
	time.Sleep(time.Second * 1)
}

func main() {
	//p := work.New(2) //使用阻塞的通道
	p := work2.New(10) //使用带有两个缓冲区的通道
	var wg sync.WaitGroup
	wg.Add(101)
	for i := 0; i < 101; i++ {
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
