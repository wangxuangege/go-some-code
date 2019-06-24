package main

import "fmt"

func main() {
	colors := make(map[string]string)
	colors["Coral"] = "#ff7f50"
	colors["DarkGray"] = "#a9a9a9"

	for key, value := range colors {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	colors2 := make(map[string][]string)
	var a []string
	colors2["Coral"] = a

	for key, value := range colors2 {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	value, exists := colors2["Coral"]
	fmt.Println(value)
	fmt.Println(exists)

	// 当可能存nil到映射=中时，使用!=nil判断是否存在是有问题的
	if value == nil {
		fmt.Println("Empty")
	}
}
