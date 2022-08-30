package main

import (
	"flag"
	"os"
	"sync"

	"github.com/kawakatz/grepfiles/pkg/grep"
	"github.com/kawakatz/grepfiles/pkg/utils"
)

func main() {
	flag.Parse()
	cmdArgs := flag.Args()
	if len(cmdArgs) == 0 {
		utils.Usage()
		os.Exit(0)
	}

	target := cmdArgs[0]
	keyword := cmdArgs[1]
	if !utils.IsExist(target) {
		utils.Usage()
		os.Exit(0)
	}

	if utils.IsDir(target) {
		files := utils.LsR(target)

		var wg sync.WaitGroup
		pathChan := make(chan string)
		for i := 0; i < 20; i++ {
			//wg.Add(1)

			go func() {
				defer wg.Done()
				for path := range pathChan {
					grep.GrepFile(path, keyword)
				}
			}()
		}

		for _, each := range files {
			pathChan <- each
		}
		wg.Wait()
		close(pathChan)
	} else {
		grep.GrepFile(target, keyword)
	}
}
