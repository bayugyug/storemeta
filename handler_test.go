package main

import (
	"fmt"
	"testing"
)

var tAppsMeta AppsMeta

func TestShowCategory(t *testing.T) {
	fmt.Println("Good!")
	fmt.Println(showCategory(tAppsMeta, IOS, ""))
}
