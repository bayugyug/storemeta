package main

import (
	"os"
)

func main() {

	//empty struct that implements
	var pAppsMeta AppsMeta

	//show cat-list
	if pIOSList {
		showCategory(pAppsMeta, IOS, "")
		return
	}

	//show cat-list
	if pAndroidList {
		showCategory(pAppsMeta, ANDROID, "")
		return
	}

	//show 1 categ
	if len(pIOSCategory) > 0 {
		showCategory(pAppsMeta, IOS, pIOSCategory)
		return
	}

	//show 1 categ
	if len(pAndroidCategory) > 0 {
		showCategory(pAppsMeta, ANDROID, pAndroidCategory)
		return
	}

	//show 1x1 per storeid
	handler(pAppsMeta)
	os.Exit(0)
}
