package task1

import "fmt"

func Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Gotcha!", err)
		}
	}()

	gimmePanic()
}

func gimmePanic() {
	panic("AAAAA")
}
