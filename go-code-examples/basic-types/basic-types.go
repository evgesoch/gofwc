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
	"math/cmplx"
)

var (
	// Declaration of several types of variables
	WhatIs  bool       = false
	HugeInt uint64     = 1<<64 - 1
	z       complex128 = cmplx.Sqrt(-5 + 12i)
)

var (
	// Declaration of variables with zero values
	anInt    int
	aFloat64 float64
	aBool    bool
	aString  string
)

var (
	// Convert some variables' types
	cx, cy int     = 3, 4
	cf     float64 = math.Sqrt(float64(cx*cx + cy*cy))
	cz     uint    = uint(cf)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", WhatIs, WhatIs)
	fmt.Printf("Type: %T Value: %v\n", HugeInt, WhatIs)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Printf("%v %v %v %q\n", anInt, aFloat64, aBool, aString)

	fmt.Println(cx, cy, cz)

	// The variable's type is complex128 without specifying it directly
	complexNum := 5 + 2i
	fmt.Printf("v is of type %T\n", complexNum)
}
