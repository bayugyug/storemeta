package main

import (
	"encoding/json"
	"flag"
	"testing"
)

func TestShowCategory(t *testing.T) {

	var pAppsMeta AppsMeta

	flag.Parse()
	//show cat-list
	if pIOSList {
		t.Log(showCategory(pAppsMeta, IOS, ""))
		return
	}
	//show cat-list
	if pAndroidList {
		t.Log(showCategory(pAppsMeta, ANDROID, ""))
		return
	}
	//show 1 categ
	if len(pIOSCategory) > 0 {
		t.Log(showCategory(pAppsMeta, IOS, pIOSCategory))
		return
	}
	//show 1 categ
	if len(pAndroidCategory) > 0 {
		t.Log(showCategory(pAppsMeta, ANDROID, pAndroidCategory))
		return
	}
	//serve http
	if pHttpServe {
		initHttpRouters()
		return
	}
	//show 1x1 per storeid
	handler(pAppsMeta)

	//show the list saved
	if len(pAppList) > 0 {
		//json fmt
		jdata, _ := json.MarshalIndent(pAppList, "", "\t")
		//dont leave your friend behind :-)
		t.Log(string(jdata))
	}

}
