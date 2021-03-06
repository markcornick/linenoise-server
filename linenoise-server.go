// SPDX-License-Identifier: MIT
//
// Copyright (c) 2021 Mark Cornick
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/markcornick/linenoise"
)

type Noise struct {
	Text string
}

type Error struct {
	Message string
}

func main() {
	http.HandleFunc("/v1/noise", func(w http.ResponseWriter, r *http.Request) {
		length, err := strconv.Atoi(r.FormValue("length"))
		if err != nil {
			length = 16
		}
		upper, err := strconv.ParseBool(r.FormValue("upper"))
		if err != nil {
			upper = true
		}
		lower, err := strconv.ParseBool(r.FormValue("lower"))
		if err != nil {
			lower = true
		}
		digit, err := strconv.ParseBool(r.FormValue("digit"))
		if err != nil {
			digit = true
		}
		p := linenoise.Parameters{Length: length, Upper: upper, Lower: lower, Digit: digit}
		noise, err := linenoise.Noise(p)
		if err != nil {
			e := &Error{Message: err.Error()}
			j, _ := json.Marshal(e)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintln(w, string(j))
		} else {
			n := &Noise{Text: noise}
			j, _ := json.Marshal(n)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, string(j))
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
