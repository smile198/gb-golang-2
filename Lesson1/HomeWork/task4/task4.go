package task4

import (
	"fmt"
	"time"
)

func Run() {
	// Переносим defer внутрь go
	//defer func() {
	//	if v := recover(); v != nil {
	//		fmt.Println("recoved", v)
	//	}
	//}()

	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("recoved", v)
			}
		}()

		panic("AAAA")
	}()

	time.Sleep(time.Second)
}
