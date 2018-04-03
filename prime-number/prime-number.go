package main

import (
	"fmt"
	"time"
)

func main() {
	getPrimeNo()
}

func getPrimeNo() {
	start := time.Now()

	max := 100
	fmt.Printf("%v 以下の素数:", max)

	for i := 2; i <= max; i++ {
		// flagの設定
		flag := true
		for j := 2; j < i; j++ {
			if (i % j) == 0 {
				// 最後の値に到達するまでに割り切れるのは、素数じゃない
				flag = false
				break
			}
		}
		if flag {
			fmt.Printf(" %v", i)
		}
	}

	goal := time.Now()                     //Goal
	fmt.Printf("\n%v 経過", goal.Sub(start)) //経過時間を表示
}
