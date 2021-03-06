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

type Point struct {
	X, Y float64
}

func main() {
	// Create a new map and give a value to it
	m1 := make(map[string]Point)
	m1["1st Point"] = Point{13.859, 20.136}
	fmt.Println("1st Map's 1st Point:", m1["1st Point"])

	// Create a map with map literals
	m2 := map[string]Point{
		"1st Point": {10.500, 11.500},
		"2nd Point": {12.600, 13.600},
		"3rd Point": {15.800, 16.800},
	}
	fmt.Println("2nd Map:", m2)

	// Mutate a map
	m3 := make(map[string]float64)
	// Insert an element
	m3["1st Point"] = 13.333
	fmt.Println("Inserted an element in map m3:", m3)
	// Update an element
	m3["1st Point"] = 14.444
	fmt.Println("Updated an element in map m3:", m3)
	// Delete an element
	delete(m3, "1st Point")
	fmt.Println("Deleted an element in map m3:", m3)
	// Check if a key exists
	value, ok := m3["1st Point"]
	fmt.Println("The key '1st Point' with value", value, "exists:", ok)

	// Iterate a map with range
	fmt.Println("Iterating the m2 map...")
	for i, v := range m2 {
		fmt.Printf("Element with key %s has value %v\n", i, v)
	}
}
