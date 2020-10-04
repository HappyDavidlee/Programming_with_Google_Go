
/**
 * Coursera assignment
 * Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
 * The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite
 * search tool to find a description of how the bubble sort algorithm works.
 * As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an 
 * argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
 * order.
 * A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent 
 * elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should
 * take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap()
 * function should return nothing, but it should swap the contents of the slice in position i with the contents in
 *  position i+1
**/

package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"log"
)

// Function to swap i+th element with i+1 element in the slice
func Swap (mySlice []int, i int) {
	
	tmp := mySlice[i]
	mySlice[i] = mySlice[i+1]
	mySlice[i+1] = tmp
}

// Function to sort the given slice in ascending order using Bubble sort algorithm
func BubbleSort(mySlice []int) {

	l := len(mySlice)
	for j:=l-1; j>0; j-- {
		for i:=0; i<j; i++ {
			if mySlice[i] > mySlice[i+1] {
			Swap (mySlice, i)
			}
		}
	}
}


// Function that reads the a string of numbers from stdin and parses them
// and stores them into an array
func ReadNumbersToSlice ()  ([]int, error) {

	// Slice to read input into
	mySlice := make ([]int, 0)

	fmt.Println ("Enter upto 10 integers separated by space")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return mySlice, err
	}

	// Trim the last newline as we used that as a delimiter to read the input
	str = str[:len(str)-1]


	// Convert strings to integers and store in the slice
	i := 0
	for _, s := range strings.Fields (str) {
		num, err := strconv.Atoi(s)
		if err == nil {
			mySlice = append (mySlice, num) 
			// Parse only upto 10 numbers; rest are ignored
			if i >= 9 {
				break
			} 
			i++
		}
	}

	return mySlice, nil

}

func main() {

	// Read input numbers into a slice
	mySlice, err := ReadNumbersToSlice()
	if err != nil {
		log.Fatal (err)
	}

	// Sort the slice in ascending order
	BubbleSort(mySlice)

	// Print the Slice
	fmt.Println ("\nSorted Slice is ... ", mySlice)

}


