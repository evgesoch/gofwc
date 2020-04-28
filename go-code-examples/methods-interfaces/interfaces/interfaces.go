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
func (mf *OneDPoint) Invert() {
	mf.X = -mf.X
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

	//

}
