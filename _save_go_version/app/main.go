package main

import (
	"fmt"

	"github.com/shaninalex/repository-watcher/app/bootstrap"
	"github.com/shaninalex/repository-watcher/app/core/errs"
)

func main() {
	root, cleanup := bootstrap.NewCLI()
	defer cleanup()

	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(*errs.Error)
			if ok {
				fmt.Printf("Error: %s - %s\nExit with: %d\n", err.Code, err.Message, err.Kind)
				return
			}
			fmt.Println("Error: ", r)
		}
	}()

	if err := root.Execute(); err != nil {
		panic(err)
	}
}
