// MIT License
//
// Copyright (c) 2016 Nick Miyake
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

package gofiles_test

import (
	"os/exec"
	"testing"

	"github.com/nmiyake/pkg-module-test/dirs"
	"github.com/nmiyake/pkg-module-test/gofiles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteGoFiles(t *testing.T) {
	dir, cleanup, err := dirs.TempDir(".", "")
	require.NoError(t, err)
	defer cleanup()

	goFiles, err := gofiles.Write(dir, []gofiles.GoFileSpec{
		{
			RelPath: "foo.go",
			Src: `package main
import (
	"fmt"
	"{{index . "bar/bar.go"}}"
	"{{index . "vendor/github.com/baz/baz.go"}}"
)
func main() {
	fmt.Println(bar.Bar(), baz.Baz())
}`,
		},
		{
			RelPath: "bar/bar.go",
			Src: `package bar
func Bar() string {
	return "bar"
}`,
		},
		{
			RelPath: "vendor/github.com/baz/baz.go",
			Src: `package baz
func Baz() string {
	return "baz"
}`,
		},
	})
	require.NoError(t, err)

	output, err := exec.Command("go", "run", goFiles["foo.go"].Path).CombinedOutput()
	require.NoError(t, err)

	assert.Equal(t, "bar baz\n", string(output))
}
