package io

import (
	"fmt"
)

func ExecuteScanning() {
	fmt.Println("Inside io.ExecuteScanning")
	defer fmt.Println("Completed io.ExecuteScanning")
	fmt.Println("What is your name?")
	var serviceType string
	fmt.Scan(&serviceType)
	fmt.Println("Hello", serviceType)
}
