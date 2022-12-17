package main

import (
	"fmt"
	"strconv"
)

// Delete a slice element of any type at index(s) and returns a new slice of the same type
func removeIndex[T any](slice []T, i int) []T {
	return append(slice[:i], slice[i+1:]...)
}

// Delete any element from a slice of strings and returns a new slice of strings
func removeString(slice []string, s string) []string {
	for i, j := range slice {
		if j == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// Delete any element from a slice of integers and returns a new slice of integers
func removeInteger(slice []int, s int) []int {
	for i, j := range slice {
		if j == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// Conversion of string slice into integer slice
func StringToInteger(slice []string) ([]int, error) {
	new_slice := make([]int, 0, len(slice))
	for _, s := range slice {
		i, err := strconv.Atoi(s)
		if err != nil {
			return new_slice, err
		}
		new_slice = append(new_slice, i)
	}
	return new_slice, nil
}

// Conversion of integer slice into string slice
func IntegerToString(slice []int) []string {
	new_slice := make([]string, 0, len(slice))
	for _, s := range slice {
		i := strconv.Itoa(s)
		new_slice = append(new_slice, i)
	}
	return new_slice
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {

	// variable definition
	index_pos := 1
	integer_element := 1
	string_element := "1"

	// 1st. example
	// Removes an integer element from the slice at defined index
	integer_slice := []int{1, 2, 3, 4}
	fmt.Printf("Old slice of integers: %#v\n", integer_slice)

	result1 := removeIndex(integer_slice, index_pos)
	fmt.Printf("New slice of integers: %#v\n", result1)
	fmt.Printf("Index removed: %#v\n\n", index_pos)

	// 2nd. example
	// Removes a string element from the slice at defined index
	string_slice := []string{"1", "2", "3", "4"}
	fmt.Printf("Old slice of strings: %#v\n", string_slice)

	result2 := removeIndex(string_slice, index_pos)
	fmt.Printf("New slice of strings: %#v\n", result2)
	fmt.Printf("Index removed: %#v\n\n", index_pos)

	// 3rd. example
	// Removes defined string element from the slice
	string_slice = []string{"1", "2", "3", "4"}
	fmt.Printf("Old slice of strings: %#v\n", string_slice)

	result3 := removeString(string_slice, string_element)
	fmt.Printf("New slice of strings: %#v\n", result3)
	fmt.Printf("String removed: %#v\n\n", string_element)

	// 4th. example
	// Removes defined integer element from the slice
	integer_slice = []int{1, 2, 3, 4}
	fmt.Printf("Old slice of integers: %#v\n", integer_slice)

	result4 := removeInteger(integer_slice, integer_element)
	fmt.Printf("New slice of integers: %#v\n", result4)
	fmt.Printf("Integer removed: %#v\n\n", integer_element)

	// 5th. example
	// Converts integer slice into string slice
	integer_slice = []int{1, 2, 3, 4}
	fmt.Printf("Old slice of integers: %#v\n", integer_slice)

	result5 := IntegerToString(integer_slice)
	fmt.Printf("New slice of strings: %#v\n\n", result5)

	// 6th. example
	// Converts string slice into integer slice
	string_slice = []string{"1", "2", "3", "4"}
	fmt.Printf("Old slice of strings: %#v\n", string_slice)

	result6, err := StringToInteger(string_slice)
	CheckError(err)
	fmt.Printf("New slice of integers: %#v\n", result6)

}
