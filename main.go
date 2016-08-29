package main

import (
	//"log"
	"os"
)

func main() {

	//show cat-list
	if pIOSList == 1 {
		showAppCategories(IOS)
		return
	}

	//show cat-list
	if pAndroidList == 1 {
		showAppCategories(ANDROID)
		return
	}

	//show 1 categ
	if len(pIOSCategory) > 0 {
		showListPerCateg(IOS, pIOSCategory)
		return
	}
	if len(pAndroidCategory) > 0 {
		showListPerCateg(ANDROID, pAndroidCategory)
		return
	}

	//start
	doIt()
	os.Exit(0)
}
