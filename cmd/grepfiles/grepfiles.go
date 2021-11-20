package main

import (
	"flag"
	"sync"

	"github.com/kawakatz/grepfiles/pkg/grep"
	"github.com/kawakatz/grepfiles/pkg/utils"
)

func main() {
	flag.Parse()
	cmdArgs := flag.Args()

	target := cmdArgs[0]
	keyword := cmdArgs[1]

	if utils.IsDir(target) {
		files := utils.LsR(target)

		var wg sync.WaitGroup
		pathChan := make(chan string)
		for i := 0; i < 20; i++ {
			wg.Add(1)

			go func() {
				defer recover()

				for path := range pathChan {
					grep.GrepFile(path, keyword)
				}
				wg.Done()
			}()
		}

		for _, each := range files {
			pathChan <- each
		}
		close(pathChan)

		wg.Wait()
	} else {
		grep.GrepFile(target, keyword)
	}
}
