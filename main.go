package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	//empty struct that implements
	var pAppsMeta AppsMeta

	//show cat-list
	if pIOSList {
		fmt.Println(showCategory(pAppsMeta, IOS, ""))
		return
	}
	//show cat-list
	if pAndroidList {
		fmt.Println(showCategory(pAppsMeta, ANDROID, ""))
		return
	}
	//show 1 categ
	if len(pIOSCategory) > 0 {
		fmt.Println(showCategory(pAppsMeta, IOS, pIOSCategory))
		return
	}
	//show 1 categ
	if len(pAndroidCategory) > 0 {
		fmt.Println(showCategory(pAppsMeta, ANDROID, pAndroidCategory))
		return
	}
	//serve http
	if pHTTPServe {
		initHTTPRouters()
		return
	}
	//show 1x1 per storeid
	handler(pAppsMeta)

	//show the list saved
	if len(pAppList) > 0 {
		//json fmt
		jdata, _ := json.MarshalIndent(pAppList, "", "\t")
		//dont leave your friend behind :-)
		fmt.Println(string(jdata))
	}
	os.Exit(0)
}
