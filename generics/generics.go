package main

import "fmt"

func printSlice[ T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}
// LIFO
// type Stack[T any] struct {
// 	elements []T
// }

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func main() {

	// myStack := Stack[string]{
	// 	elements: []string{"aayush", "singh", "gupta"}}
	// fmt.Println(myStack)

	// numbers := []int{1, 2, 3, 4, 5}
	// name := []string{"aayush", "singh", "gupta"}
	vals := []bool{true, false, true}
	// printSlice(numbers)
	printSlice(vals, "aayush")
}
