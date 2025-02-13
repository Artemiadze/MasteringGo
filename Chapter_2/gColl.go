package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	/*каждый раз, когда мы хотим получить свежую статистику
	о сборке мусора, нам нужно заново вызывать функцию runtime.ReadMemStats(). Функ-
	ция printStats() нужна для того, чтобы не писать многократно один и тот же код Go.*/
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)
	/*Цикл for создает много больших срезов Go, чтобы под них выделялись большие
	  объемы памяти и запускался сборщик мусора.*/
	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStats(mem)

	/*следующий код Go, который вызывает
	  дополнительное выделение памяти для срезов Go:*/
	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(5 * time.Second)
	}
	printStats(mem)
}

/*В следующей команде показан прием, позволяющий получить еще более подробный
отчет о том, как работает сборщик мусора Go:
$ GODEBUG=gctrace=1 go run gColl.go
Если перед любой командой go run поставить GODEBUG=gctrace=1, то Go выводит
аналитические данные о работе сборщика мусора.*/
