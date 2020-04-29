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

// An interface that has a single method, Invert
type Inverter interface {
	Invert()
}

// Type TwoDPoint implements the Inverter interface
type TwoDPoint struct {
	X, Y float64
}

// Implicit implementation of the Inverter interface
func (tdp *TwoDPoint) Invert() {
	tdp.X = -tdp.X
	tdp.Y = -tdp.Y
}

// Type OneDPoint implements the Inverter interface
type OneDPoint struct {
	X float64
}

// Implicit implementation of the Inverter interface
func (odp *OneDPoint) Invert() {
	odp.X = -odp.X
}

func main() {
	// Create Inverter, TwoDPoint and OneDPoint variables
	var i1, i2 Inverter
	p1 := TwoDPoint{1, 2}
	p2 := OneDPoint{3}
	i1 = &p1 // Only *TwoDPoint implements Inverter, not TwoDPoint
	i2 = &p2 // Only *OneDPoint implements Inverter, not OneDPoint
	i1.Invert()
	i2.Invert()
	fmt.Println("*TwoDPoint type var p1 implements Inverter:", i1)
	fmt.Println("*OneDPoint type var mf1 implements Inverter:", i2)

	// Check interface values and types
	fmt.Printf("Value and type of i1 interface: (%v, %T)\n", i1, i1)
	fmt.Printf("Value and type of i2 interface: (%v, %T)\n", i2, i2)

	// Interface with nil value
	var (
		i3 Inverter
		p3 *TwoDPoint
	)
	i3 = p3
	fmt.Printf("Value and type of i3 interface: (%v, %T)\n", i3, i3)

	// Nil interface
	var i4 Inverter
	fmt.Printf("Value and type of i4 interface: (%v, %T)\n", i4, i4)
	// Calling Invert on i4 causes run-time error: i4 doesn't have a type
	// i4.Invert()

	// Empty interface
	var i5 interface{}
	fmt.Printf("Value and type of i5 interface: (%v, %T)\n", i5, i5)
	i5 = "hi"
	fmt.Printf("Value and type of i5 interface: (%v, %T)\n", i5, i5)
	i5 = 5
	fmt.Printf("Value and type of i5 interface: (%v, %T)\n", i5, i5)

	// Type assertions
	var i6 interface{} = 88.88
	t, ok := i6.(float64) // ok can be omitted, can be written t := i6.(float64)
	if ok {
		fmt.Printf("Type assertion passed: i6 (%v, %T)\n", t, i6)
	}
	// The following line triggers a panic
	// t1 := i6.(string)

	// Type switch
	switch v := i6.(type) {
		case int:
			fmt.Printf("INT! i6 is of type %T with value %v\n", v, v)
		case float64:
			fmt.Printf("FLOAT64! i6 is of type %T with value %v\n", v, v)
		default:
			fmt.Printf("DEFAULT CASE! i6 is of type %T!\n", v)
	}
}
