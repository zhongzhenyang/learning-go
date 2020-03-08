package main

import "fmt"

func main() {
	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}

	for key, value := range m2 {
		fmt.Println(key, value)
	}
}
