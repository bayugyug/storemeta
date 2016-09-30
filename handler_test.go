package main

import (
	"fmt"
	"strings"
	"testing"
)

var (
	categOSes    = []string{IOS, ANDROID}
	categOSTypes = map[string]string{
		IOS:     "GAMES_ACTION",
		ANDROID: "GAME_ACTION",
	}
	storeIDList = map[string]string{
		IOS:     "293622097",
		ANDROID: "com.google.android.apps.plus",
	}
)

func TestHandler(t *testing.T) {

	var pAppsMeta AppsMeta

	for _, os := range categOSes {
		s := showCategory(pAppsMeta, os, "")
		if len(s) == 0 {
			t.Error("Fail showCategory", os)
		} else {
			t.Log("showCategory OK", os)
		}
	}

	for oss, ost := range categOSTypes {
		s := showCategory(pAppsMeta, oss, ost)
		if len(s) == 0 {
			t.Error("Fail showCategory type", oss, ost)
		} else {
			t.Log("showCategory type OK", oss, ost)
		}
	}

	ands := strings.Split(storeIDList[ANDROID], ",")
	ioss := strings.Split(storeIDList[IOS], ",")
	queryStoreIds(ands, ioss)

	//show 1x1 per storeid
	handler(pAppsMeta)

	//show the list saved
	if len(pAppList) == 0 {
		t.Error("Fail getting meta infos.")
	} else {
		t.Log("Get Store IDs OK", ANDROID, storeIDList[ANDROID])
		t.Log("Get Store IDs OK", IOS, storeIDList[IOS])
	}

}

func BenchmarkHandlerParallel(b *testing.B) {

	var pAppsMeta AppsMeta

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//re-init it here, eventhough, its defined @ global.go
			pAppsData = make(chan *App)
			pAppList = []*App{}
			pStores = []*StoreApp{}

			ands := strings.Split(storeIDList[ANDROID], ",")
			ioss := strings.Split(storeIDList[IOS], ",")
			queryStoreIds(ands, ioss)

			//show 1x1 per storeid
			handler(pAppsMeta)

			//show the list saved
			if len(pAppList) > 0 {
				fmt.Println()
				fmt.Println("Get Store IDs OK", ANDROID, storeIDList[ANDROID])
				fmt.Println("Get Store IDs OK", IOS, storeIDList[IOS])
			}
		}
	})
}
