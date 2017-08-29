package main

import (
	"fmt"
	"testing"
	"time"
)

func slowFunc() {
	fmt.Println("I am a slow func I will sleep")
	t.Fatalf("And now I will fail")
	time.Sleep(5 * time.Second)
	fmt.Println("I am awake")
}

func Test_slowFunc(t *testing.T) {
	slowFunc()
}
