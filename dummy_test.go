package main

import (
	"testing"
)

func TestShowCategory(t *testing.T) {

	var smeta AppsMeta

	showCategory(smeta, IOS, "")

	t.Fatal("test a")
}
