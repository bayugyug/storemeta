package main

import (
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
		IOS:     "544007664,535886823,643496868,293622097",
		ANDROID: "com.google.android.apps.plus,com.google.android.launcher,com.sphero.sprk",
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