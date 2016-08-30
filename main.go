package main

import (
	"os"
)

func main() {

	//empty struct that implements
	var pAppsCat AppsCategory
	var pAppsMeta AppsMeta

	//show cat-list
	if pIOSList {
		showCategory(pAppsCat, IOS, "")
		return
	}

	//show cat-list
	if pAndroidList {
		showCategory(pAppsCat, ANDROID, "")
		return
	}

	//show 1 categ
	if len(pIOSCategory) > 0 {
		showCategory(pAppsCat, IOS, pIOSCategory)
		return
	}

	//show 1 categ
	if len(pAndroidCategory) > 0 {
		showCategory(pAppsCat, ANDROID, pAndroidCategory)
		return
	}

	//show 1x1 per storeid
	handler(pAppsMeta)
	os.Exit(0)
}
