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

// Function sum sums all the elements of the input slice,
// then sends the computed sum value inside a channel that
// receives integer values
func sum(sl []int, ch chan int) {
	sum := 0
	for _, v := range sl {
		sum += v
	}
	// Wait till ch is ready to receive values, then send the sum
	// into ch
	ch <- sum
}

func main() {
	// Simple demonstration of channel functionality
	// Create a slice and a channel
	sl1 := []int{1, 2, 3, 4, 5, 6}
	ch1 := make(chan int)
	// Create 2 goroutines, where each one computes
	// the sum of the half of the s slice
	go sum(sl1[:len(sl1)/2], ch1)
	go sum(sl1[len(sl1)/2:], ch1)
	// Assign the computed sums into 2 variables.
	// We don't know which goroutine will finish first,
	// so x is not necessarily going to be 6 and y 15
	x, y := <-ch1, <-ch1
	fmt.Printf(
		"One sum of s slice: %v\n" +
		"The other sum of s slice: %v\n" +
		"Total sum of s slice: %v", x, y, x+y,
	)

	//

}
