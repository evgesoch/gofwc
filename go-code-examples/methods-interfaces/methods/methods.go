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
	"math"
)

type Point struct {
	X, Y float64
}

// Declare an Abs method for the Point struct.
// Abs has a value receiver, which means that it operates
// on a copy of the original Point value
func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// The Scale method has a pointer receiver, because it
// should be able to change the Point's value that is called upon
func (p *Point) Scale(f float64) {
	p.X = p.X * f
	p.Y = p.Y * f
}

// Abs as a normal function that takes a Point as an argument
func Abs(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

type MyInt int

// Declare an Absolute method for the non-struct type MyInt
func (i MyInt) Absolute() int {
	if i < 0 {
		return int(-i)
	}
	return int(i)
}

func main() {
	// Call the Abs method of a Point struct
	p1 := Point{3, 4}
	fmt.Println("p1 Abs method:", p1.Abs())

	// Call the Abs function that takes a Point as an argument
	fmt.Println("Abs function with p1 as argument:", Abs(p1))

	// Call the Absolute method of a MyInt variable
	mi1 := MyInt(-5)
	fmt.Println("mi1 Absolute method:", mi1.Absolute())

	// Call the Scale method to change a Point
	p2 := Point{1, 2}
	fmt.Println("p2 before Scale method:", p2)
	p2.Scale(2) // (&p2).Scale will produce the same result
	fmt.Println("p2 after Scale method:", p2)

	//



}
