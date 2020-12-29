package main

import (
	"log"
	"sync"
	"time"
)

type Tracker struct {
	wg sync.WaitGroup
}

func(t *Tracker)Event(data string){
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		time.Sleep(time.Millisecond)
		log.Println(data)
	}()
}

func (t *Tracker)Shutdown(){
	t.wg.Wait()
}
