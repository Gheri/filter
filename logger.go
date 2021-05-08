package main

import (
	"fmt"
)

type Logger struct{}

func NewLogger() Logger {
	return Logger{}
}

func (l Logger) Info(message string) {
	fmt.Println(message)
}

func (l Logger) Error(err error) {
	fmt.Printf("Error %v", err)
}
