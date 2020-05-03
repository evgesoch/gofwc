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
	"sync"
	"time"
)

// Type SafeCounter is safe to use concurrently
type SafeCounter struct {
	safeMap map[string]int
	mux     sync.Mutex
}

// Method Inc increments the counter for the given key
func (sc *SafeCounter) Inc(key string) {
	// Lock so only one goroutine at a time can access the map
	sc.mux.Lock()
	// Increment the value safely
	sc.safeMap[key]++
	fmt.Println(sc.safeMap[key])
	// Unlock to give access to another goroutine
	sc.mux.Unlock()
	//fmt.Println(sc.Value(key))
}

// Method Value returns the current value of the counter for the given key
func (sc *SafeCounter) Value(key string) int {
	// Lock so only one goroutine at a time can access the map
	sc.mux.Lock()
	// Unlock with defer to ensure that the mutex will be unlocked
	// after the function returns
	defer sc.mux.Unlock()
	return sc.safeMap[key]
}

func main() {
	sc := SafeCounter{safeMap: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go sc.Inc("first_key")
	}
	// Delay the execution of the main goroutine to start
	// scheduling the goroutines that increment the map's value
	time.Sleep(time.Second)
	fmt.Println(
		"The final value of the first_key is:",
		sc.Value("first_key"),
	)
}
