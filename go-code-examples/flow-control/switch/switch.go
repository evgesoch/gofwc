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
	"runtime"
	"time"
)

func main() {
	// Switch statement. Notice the initial short variable declaration
	fmt.Print("My OS is: ")
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X")
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Printf("%s.\n", os)
	}

	// Switch statement with cases that are variables
	fmt.Print("When's June? ")
	month := time.Now().Month()
	switch time.June {
		case month + 1:
			fmt.Println("In 1 month")
		case month + 2:
			fmt.Println("In 2 months")
		case month + 3:
			fmt.Println("In 3 months")
		default:
			fmt.Println("Too far away...")
	}

	// Switch statement without condition
	current_time := time.Now()
	switch {
		case current_time.Hour() < 12:
			fmt.Println("Good morning")
		case current_time.Hour() < 17:
			fmt.Println("Good afternoon")
		default:
			fmt.Println("Good evening")
	}
}
