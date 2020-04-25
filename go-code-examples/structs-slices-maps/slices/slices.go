/*
Copyright (c) 2011 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	// First create an array
	primes := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	// Then create a slice that contains the elements 1 to 4 of the array
	var sl1 []int = primes[1:5]
	fmt.Println("Slice sl1 of the primes array:", sl1)

	// Create slices of an array, changes values and notice the changes in both slices and the array
	nums := [5]string{"one", "two", "three", "four", "five"}
	fmt.Println("Τhe nums array:", nums)
	sl2 := nums[0:3]
	sl3 := nums[1:4]
	fmt.Println("A slice sl2 of the nums array:", sl2)
	fmt.Println("Another slice sl3 of the nums array:", sl3)
	sl3[1] = "CHANGE"
	fmt.Println("Slice sl2 changed:", sl2)
	fmt.Println("Slice sl3 changed:", sl3)
	fmt.Println("nums array changed:", nums)

	// Slice literals
	sl4 := []bool{true, false, true, true, false, true}
	fmt.Println("Slice sl4:", sl4)

	// Slice defaults
	sl5 := sl4[:2]
	sl6 := sl4[2:5]
	sl7 := sl4[4:]
	fmt.Println("Slice sl4 defaults:", sl5, sl6, sl7)

	// Slices' length and capacity
	sl8 := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("The sl8 slice: %v len=%d cap=%d\n", sl8, len(sl8), cap(sl8))
	// Slice the slice to give it zero length
	sl8 = sl8[:0]
	fmt.Printf("sl8 re-sliced 1st time: %v len=%d cap=%d\n", sl8, len(sl8), cap(sl8))
	// Extend its length
	sl8 = sl8[:4]
	fmt.Printf("sl8 re-sliced 2nd time: %v len=%d cap=%d\n", sl8, len(sl8), cap(sl8))
	// Drop its first two values
	sl8 = sl8[2:]
	fmt.Printf("sl8 re-sliced 3rd time: %v len=%d cap=%d\n", sl8, len(sl8), cap(sl8))

	// A nil slice
	var sl9 []int
	fmt.Printf("The sl9 slice: %v len=%d cap=%d\n", sl9, len(sl9), cap(sl9))
	if sl9 == nil {
		fmt.Println("sl9 is nil")
	}

	// Create slice with make
	sl10 := make([]int, 0, 6)
	fmt.Printf("The sl10 slice: %v len=%d cap=%d\n", sl10, len(sl10), cap(sl10))
	sl11 := sl10[:2]
	fmt.Printf("sl11 slice created by sl10: %v len=%d cap=%d\n", sl11, len(sl11), cap(sl11))

	// Slices of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "O"
	board[2][2] = "X"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Append new elements to a slice
	var sl12 []string
	fmt.Printf("The sl12 slice: %v len=%d cap=%d\n", sl12, len(sl12), cap(sl12))
	sl12 = append(sl12, "hi")
	sl12 = append(sl12, "there", "!")
	fmt.Printf("The sl12 slice with new elements: %v len=%d cap=%d\n", sl12, len(sl12), cap(sl12))

	// Range loop a slice
	sl13 := []float64{0.1, 0.2, 0.3, 0.4}
	for i, v := range sl13 {
		fmt.Printf("Element in index %d has value %v\n", i, v)
	}

	// Range without value
	for i := range sl13 { // can also be written: for i, _ := range...
		fmt.Printf("Element in index %d\n", i)
	}
	// Range without index
	for _, v := range sl13 {
		fmt.Printf("Element has value %v\n", v)
	}
}
