## Golang并发控制

## 示例代码
### 阻塞的通道
```
package work

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()

	}
	return &p
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool)Shutdown()  {
	close(p.work)
	p.wg.Wait()
}

```

### 带有缓冲区的通道
```
package work2

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker,maxGoroutines),
	}
	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()

	}
	return &p
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

```

### Demo
```
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
	p := work2.New(2) //使用带有两个缓冲区的通道
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

```

## 核心代码说明
```
主要通过无缓冲的通道,使channel阻塞,通过range这个阻塞的channel来启动 goroutine

其实我们也可以用带有缓冲区的channel来实现

```
