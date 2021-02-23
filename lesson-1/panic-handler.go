package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type MyError struct {
	msg  string
	time time.Time
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error: %s\ntime:\n%s", e.msg, e.time.String())
}

func NewError(msg string) error {
	return &MyError{msg: msg, time: time.Now()}
}

func main() {
	getPanic()
	fmt.Println(NewError("my error"))
	createFile()
}

func getPanic() {
	defer recoverPanic()
	var i int = 5
	var j int = 0
	fmt.Println(i / (i * j))
}

func recoverPanic() {
	if v := recover(); v != nil {
		fmt.Println("recovered", v)
	}
}

func createFile() {
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
