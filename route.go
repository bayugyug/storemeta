package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	msg := `

		http://127.0.0.1:7777/list-category-android

		http://127.0.0.1:7777/list-category-ios

		http://127.0.0.1:7777/category-android/?p=GAME_ACTION

		http://127.0.0.1:7777/category-ios/?p=GAMES_ACTION

		http://127.0.0.1:7777/storeid/?a=com.google.android.apps.photos&i=293622097
`
	fmt.Fprint(w, "Welcome to Storemeta!\n", fmt.Sprintf("Version: %s\n\n\n%s", pVersion, msg))
}

func formatHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//Content-Type: application/json
	var pAppsMeta AppsMeta

	var q = r.URL.Query()
	var p = strings.TrimSpace(q.Get("p"))
	var m = strings.ToUpper(strings.TrimSpace(ps.ByName("mode")))

	//re-init it here, eventhough, its defined @ global.go
	pAppsData = make(chan *App)
	pAppList = []*App{}
	pStores = []*StoreApp{}
	//hdrset
	w.Header().Set("Content-Type", "application/json")
	//not-found
	v, ok := Formatters[m]
	if !ok {
		if !strings.EqualFold(m, "STOREID") {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		//store-ids here
		ands := strings.Split(strings.TrimSpace(q.Get("a")), ",")
		ioss := strings.Split(strings.TrimSpace(q.Get("i")), ",")
		queryStoreIds(ands, ioss)
		if len(pStores) == 0 {
			http.Error(w, "StoreId is Missing / "+http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		//handle the storeids
		handler(pAppsMeta)

		//show the list saved
		if len(pAppList) > 0 {
			//json fmt
			jdata, _ := json.MarshalIndent(pAppList, "", "\t")
			//dont leave your friend behind :-)
			fmt.Fprint(w, string(jdata))
		}
		return
	}

	//reset
	if ok, _ := regexp.MatchString("(?i)^list-category-(ios|android)$", m); ok {
		p = ""
	}
	result := v.Format(pAppsMeta, v.Mode, p)
	//good
	fmt.Fprint(w, string(result))
}

func queryStoreIds(androids, ioss []string) {
	//store-ids here
	for _, s := range androids {
		if len(s) > 0 {
			pStores = append(pStores, &StoreApp{OS: ANDROID, URL: pStoreURI[ANDROID][0] + s + pStoreURI[ANDROID][1], StoreID: s})
		}
	}
	for _, s := range ioss {
		if len(s) > 0 {
			pStores = append(pStores, &StoreApp{OS: IOS, URL: pStoreURI[IOS][0] + s + pStoreURI[IOS][1], StoreID: s})
		}
	}
}
