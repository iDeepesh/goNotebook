package io

import (
	"fmt"
)

//ExecuteScanning - an example of collecting user input on stdin console
func ExecuteScanning() {
	fmt.Println("Inside io.ExecuteScanning")
	defer fmt.Println("Completed io.ExecuteScanning")
	fmt.Println("What is your name?")
	var serviceType string
	fmt.Scan(&serviceType)
	fmt.Println("Hello", serviceType)
}
