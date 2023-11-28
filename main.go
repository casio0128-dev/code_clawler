package main

import (
	"os"
	"strings"
	"sync"
)

var (
	args []string
)

func init() {
	args = os.Args[1:]
}

func main() {
	wg := &sync.WaitGroup{}
	for _, arg := range args {
		wg.Add(1)
		// 読み込み対象
		fp, err := os.Open(arg)
		if err != nil {
			panic(err)
		}
		go func(f *os.File, a string) {
			defer wg.Done()
			// 新しいファイルポインタ
			fpp, err := CreateNewFile(a)
			if err != nil {
				panic(err)
			}
			defer func() {
				if err := fpp.Close(); err != nil {
					panic(err)
				}
			}()
		}(fp, arg)
	}
	wg.Wait()
}

func CreateNewFile(arg string) (*os.File, error) {
	cwd := getCWD(arg)
	fp, err := os.OpenFile(cwd+"_new", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return nil, err
	}
	return fp, nil
}

func getCWD(dirName string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return JoinFilePath(wd, dirName)
}

func JoinFilePath(paths ...string) string {
	return strings.Join(paths, string(os.PathSeparator))
}
