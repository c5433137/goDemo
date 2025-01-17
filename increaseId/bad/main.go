package main

import (
	"runtime"
	"time"
	"sync"
	"fmt"
)

//全局唯一id
var GId int64

var Sum int64
var sumLock sync.Mutex
var wg sync.WaitGroup

func main(){
	runtime.GOMAXPROCS(5)
	for i:= 0;i<1234567;i++{
		wg.Add(1)
		go func(){
			num := GetId()
			addSum(num)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("the bad total = %d",Sum)	//741442659357
	time.Sleep(5 * time.Second)
}

func GetId()int64{
	GId = GId +1
	return GId
}

func addSum(ad int64){
	sumLock.Lock()
	Sum = Sum + ad
	sumLock.Unlock()
}
