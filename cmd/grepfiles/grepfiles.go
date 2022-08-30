package main

import (
	"flag"
	"fmt"
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
		for i := 0; i < 1; i++ {
			wg.Add(1)

			go func() {
				//defer recover()

				for path := range pathChan {
					fmt.Println(path)
					grep.GrepFile(path, keyword)
				}
				wg.Done()
			}()
		}

		for _, each := range files {
			pathChan <- each
		}

		wg.Wait()
	} else {
		grep.GrepFile(target, keyword)
	}
}
