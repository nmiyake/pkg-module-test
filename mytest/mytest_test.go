package mytest_test

import (
	"io/ioutil"
	"testing"

	"github.com/nmiyake/pkg-module-test/mytest"
)

func TestFoo(t *testing.T) {
	if gotFoo := mytest.Foo(); gotFoo != "Foo" {
		t.Error("did not get expected foo: " + gotFoo)
	}

	bytes, err := ioutil.ReadFile("testdata/foo.txt")
	if err != nil {
		panic(err)
	}
	if gotContent := string(bytes); gotContent != "123" {
		t.Error("did not get expected content: " + gotContent)
	}
}
