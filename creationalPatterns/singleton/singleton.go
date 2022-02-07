package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Criando singleton")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Singleton já criado")
	}

	return singleInstance
}
