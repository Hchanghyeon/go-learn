package main

import (
	"fmt"
	watcher "orchestrator/plugins/watcher"
	"time"
)

func main(){
	go watcher.WacthKey("changhyeon") // 비동기

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

