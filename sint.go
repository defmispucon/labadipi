package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data := []byte("Hello, World!")
	err := ioutil.WriteFile("example.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File written successfully")
}
