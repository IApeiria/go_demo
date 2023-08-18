package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

// Worker 依赖时间戳+进程ID+序号的组合来达到唯一性
type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("worker ID excess of quantity")
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	ID := (now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number)
	return ID
}

func main() {
	// 创建两个工作进程
	worker1, _ := NewWorker(1)
	worker2, _ := NewWorker(2)

	var wg sync.WaitGroup
	wg.Add(2)

	// 工作进程1生成ID
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			id := worker1.GetId()
			fmt.Println("Worker 1: ", id)
		}
	}()

	// 工作进程2生成ID
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			id := worker2.GetId()
			fmt.Println("Worker 2: ", id)
		}
	}()

	wg.Wait()

}
