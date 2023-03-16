package main

import (
	"fmt"
	"sync"
	"time"
)

func writeChan(intChan chan int,wg *sync.WaitGroup){
	defer (wg).Done()
	for i := 0 ; i <= 2 ; i++{
		intChan <- i
		fmt.Println("writeChan=",i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func readChan (intChan chan int,wg *sync.WaitGroup) {
	defer (*wg).Done()
	for{
		// 正常
		// num , ok := <- intChan
		// if !ok {
		// 	break
		// }
		// fmt.Println("readChan=",num)

		// 阻塞
		for num := range intChan{
			fmt.Println("readChan=",num)
		}

		// 死循环
		// select {
		// 	case num := <- intChan:
		// 		fmt.Println("readChan=",num)
		// 	default:
		// 		break
		// }
	}
	fmt.Println("hahah")	
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	intChan := make(chan int,11)
	go writeChan(intChan,&wg)
	go readChan(intChan,&wg)
	wg.Wait()
}