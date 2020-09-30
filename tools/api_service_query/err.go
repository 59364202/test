package main

import (
	"log"
	//	"os"
)

func handlerError(err error) {
	if err != nil {
		log.Println("error : ", err)
		//		os.Exit(1)
	}
}
