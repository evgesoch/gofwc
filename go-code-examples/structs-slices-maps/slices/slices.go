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

import "fmt"

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

	//

}
