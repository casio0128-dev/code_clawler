package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	args
)

func init() {
	args = os.Args[1:]
}

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println(args)
	for _, arg := range args {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fp, err := os.OpenFile(arg, os.O_CREATE | os.O_RDWR, 0755)
			if err != nil {
				panic(err)
			}
			defer fp.Close()

		}()
	}
	wg.Wait()
}
