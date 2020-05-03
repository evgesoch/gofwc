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

// Function fibonacci is used to demonstrate the combination
// of close - range functionalities. n is the number of values
// the receiver will request (the capacity of channel ch3)
func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		// Send a value to the channel
		ch <- x
		x, y = y, x+y
	}
	// Close the channel after sending all the values to be sent
	close(ch)
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
	// Receive and assign the computed sums into 2 variables.
	// We don't know which goroutine will finish first,
	// so x is not necessarily going to be 6 and y 15
	sum1, sum2 := <-ch1, <-ch1
	fmt.Printf(
		"One sum of sl1 slice: %v\n" +
		"The other sum of sl1 slice: %v\n" +
		"Total sum of sl1 slice: %v\n", sum1, sum2, sum1 + sum2,
	)

	// Buffered channels
	ch2 := make(chan string, 2)
	ch2 <- "say"
	ch2 <- "something"
	// This will cause a deadlock because buffer is filled
	//ch2 <- 3
	fmt.Println("Retrieving from buffered channel ch2:", <-ch2)
	fmt.Println("Retrieving from buffered channel ch2:", <-ch2)
	// An extra retrieval from an empty buffer will
	// also cause a deadlock
	// fmt.Println(<-ch2)

	// Range - Close demonstration
	ch3 := make(chan int, 10)
	// Initiate a goroutine running fibonacci function
	go fibonacci(cap(ch3), ch3)
	// Consume all the values inside the ch3 channel with
	// a range loop
	fmt.Printf("A fibonacci sequence with %v values:\n", cap(ch3))
	for i := range ch3 {
		fmt.Println(i)
	}

	//

}
