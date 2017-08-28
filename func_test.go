package main

import (
	"fmt"
	"testing"
	"time"
)

func slowFunc() {
	fmt.Println("I am a slow func I will sleep")
	time.Sleep(100 * time.Second)
	fmt.Println("I am awake")
}

func Test_slowFunc(t *testing.T) {
	slowFunc()
}
