package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strings"
)

type AppsCategories interface {
	Showlist(string)
	ShowlistPerCategory(string, string)
	List(string, string, string)
}

type AppsCategory struct{}

func (appscat AppsCategory) Showlist(os string) {
	//init
	categlist := []string{}
	categ := ""
	t := 0
	re := regexp.MustCompile("[^0-9A-Za-z]+")
	if _, ok := pCategories[os]; ok {
		for _, v := range pCategories[os] {
			if os == IOS {
				cats := strings.Split(v, "/")
				categ = strings.ToUpper(strings.Replace(cats[5], "ios-", "", -1))
			} else if os == ANDROID {
				cats := strings.Split(v, "/")
				categ = strings.ToUpper(cats[6])
			}
			t++
			categ = re.ReplaceAllString(categ, "_")
			if strings.EqualFold(pPrintFormat, "json") {
				categlist = append(categlist, categ)
			} else {
				log.Println(t, categ)
			}
		}
	}
	//json fmt
	if len(categlist) > 0 {
		jdata, _ := json.MarshalIndent(categlist, "", "\t")
		//dont leave your friend behind :-)
		log.Println(string(jdata))
	}
}

func (appscat AppsCategory) ShowlistPerCategory(os, category string) {

	categ := ""
	t := 0
	re := regexp.MustCompile("[^0-9A-Za-z]+")
	if _, ok := pCategories[os]; ok {
		for _, v := range pCategories[os] {
			if os == IOS {
				cats := strings.Split(v, "/")
				categ = re.ReplaceAllString(strings.ToUpper(strings.Replace(cats[5], "ios-", "", -1)), "_")
			} else if os == ANDROID {
				cats := strings.Split(v, "/")
				categ = re.ReplaceAllString(strings.ToUpper(cats[6]), "_")

			}
			if strings.EqualFold(category, categ) {
				appscat.List(os, category, v)
				t++
				break
			}

		}
	}
	if t == 0 {
		log.Println("\n\n", os, "Category not found!", category)
	}
}

func (appscat AppsCategory) List(os, category, url string) {
	status, body := getResult(url)
	if status != 200 || body == "" {
		log.Println("ERROR: invalid http status", status)
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	appt := 0
	var storelist []*StoreApp
	switch os {
	case IOS:
		//app list
		doc.Find("#selectedcontent").Find("a").Each(func(i int, n *goquery.Selection) {
			for _, v := range n.Nodes[0].Attr {
				//GENRE
				if v.Key == "href" {
					storeid := ""
					xtores := strings.Split(v.Val, "/")
					if len(xtores) > 6 {
						ztores := strings.Split(xtores[6], "?")
						if ztores[0] != "" {
							storeid = strings.TrimSpace(ztores[0])
						}
						storelist = append(storelist, &StoreApp{Preview: strings.TrimSpace(v.Val), Category: category, StoreID: strings.Replace(storeid, "id", "", -1)})
						appt++
					}

				}
			}
		})
	case ANDROID:
		//AUTHOR
		doc.Find("a.card-click-target").Each(func(i int, n *goquery.Selection) {
			for _, v := range n.Nodes[0].Attr {
				//app
				if v.Key == "href" && strings.ContainsAny(v.Val, "/store/apps/details?id=") {
					storeid := ""
					xtores := strings.Split(v.Val, "details?id=")
					if len(xtores) >= 2 {
						storeid = strings.TrimSpace(xtores[1])
						storelist = append(storelist, &StoreApp{Preview: "https://play.google.com" + strings.TrimSpace(v.Val), Category: category, StoreID: strings.TrimSpace(storeid)})
						appt++
					}

				}
			}
		})
	}
	jdata, _ := json.MarshalIndent(storelist, "", "\t")
	//dont leave your friend behind :-)
	log.Println(string(jdata))
}

func showCategory(appscat AppsCategories, os, category string) {
	switch os {
	case IOS:
		if len(category) > 0 {
			appscat.ShowlistPerCategory(os, category)
		} else {
			appscat.Showlist(os)
		}
	case ANDROID:
		if len(category) > 0 {
			appscat.ShowlistPerCategory(os, category)
		} else {
			appscat.Showlist(os)
		}
	}
}
