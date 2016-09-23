package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"regexp"
	"strings"
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

	//re-init it here, eventhoug, its defined @ global.go
	pAppsData = make(chan *App)
	pAppList = []*App{}
	pStores = []*StoreApp{}

	//not-found
	v, ok := Formatters[m]
	if !ok {
		if !strings.EqualFold(m, "STOREID") {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		//store-ids here
		androids := strings.Split(strings.TrimSpace(q.Get("a")), ",")
		for _, s := range androids {
			if len(s) > 0 {
				pStores = append(pStores, &StoreApp{OS: ANDROID, URL: pStoreURI[ANDROID][0] + s + pStoreURI[ANDROID][1], StoreID: s})
			}
		}
		ioss := strings.Split(strings.TrimSpace(q.Get("i")), ",")
		for _, s := range ioss {
			if len(s) > 0 {
				pStores = append(pStores, &StoreApp{OS: IOS, URL: pStoreURI[IOS][0] + s + pStoreURI[IOS][1], StoreID: s})
			}
		}
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
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(jdata))
		}
		return
	}

	//reset
	if ok, _ := regexp.MatchString("(?i)^list-category-(ios|android)$", m); ok {
		p = ""
	}
	result := v.Format(pAppsMeta, v.Mode, p)
	fmt.Println("RAW-DATA: ", p)
	//good
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(result))
}
