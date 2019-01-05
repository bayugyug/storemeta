package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func b(s string) *bufio.Reader { return bufio.NewReader(strings.NewReader(s)) }

//AppsMetaInfos skeleton interface
type AppsMetaInfos interface {
	Process(chan bool, *sync.WaitGroup, int, *StoreApp)
	Show(chan bool, *sync.WaitGroup)
	FormatAndroid(*goquery.Document, *StoreApp) *App
	FormatIOS(*goquery.Document, *StoreApp) *App
	ShowCategories(string) []string
	ShowlistApps(string, string) []*StoreApp
	PrintList(string, string, string) []*StoreApp
}

//AppsMeta empty holder
type AppsMeta struct{}

//handler entry of the app
func handler(metainfo AppsMeta) {

	//get task
	zFlag := make(chan bool)
	zwg := new(sync.WaitGroup)

	//get task
	uFlag := make(chan bool)
	uwg := new(sync.WaitGroup)
	uwg.Add(1)
	go metainfo.Show(uFlag, uwg)

	//get task
	for idx, url := range pStores {
		if !pStillRunning {
			log.Println("Signal detected ...")
			break
		}
		zwg.Add(1)
		go metainfo.Process(zFlag, zwg, idx+1, url)
	}

	zwg.Wait()
	close(zFlag)
	close(pAppsData)

	//stats
	uwg.Wait()
	close(uFlag)
}

//Process do the processing of the appstore URL
func (metainfo AppsMeta) Process(doneFlg chan bool, wg *sync.WaitGroup, idx int, store *StoreApp) {

	go func() {
		for {
			select {
			//wait till doneFlag has value ;-)
			case <-doneFlg:
				//done already ;-)
				wg.Done()
				return
			}
		}
	}()

	//sig-check
	if !pStillRunning {
		log.Println("Signal detected ...")
		doneFlg <- true
		return
	}
	status, body := getResult(store.URL)
	if status != 200 || body == "" {
		log.Println("ERROR: invalid http status", status)
		doneFlg <- true
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Println("ERROR: ", err)
		doneFlg <- true
		return
	}
	var appsdata *App
	switch store.OS {
	case IOS:
		appsdata = metainfo.FormatIOS(doc, store)
	case ANDROID:
		appsdata = metainfo.FormatAndroid(doc, store)
	}
	pAppsData <- appsdata

	//send signal -> DONE
	doneFlg <- true
}

//FormatAndroid meta info formatting for Android Playstore
func (metainfo AppsMeta) FormatAndroid(doc *goquery.Document, store *StoreApp) (appsdata *App) {

	tmpf3 := make(map[string]string)

	//init
	appsdata = &App{AppID: store.StoreID, AppURL: store.URL, Platform: store.OS}

	//META
	doc.Find("meta").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "itemprop" {
				ru, _ := n.Attr("itemprop")
				rv, _ := n.Attr("content")
				tmpf3[ru] = strings.Replace(strings.TrimSpace(rv), "\\n", "\n", -1)
			}
		}
	})

	//GENRE
	for mk, mv := range tmpf3 {
		switch mk {
		case "applicationCategory":
			appsdata.Genre = mv
		case "price":
			appsdata.SoftwarePrice = mv
			if appsdata.SoftwarePrice == "0" {
				appsdata.SoftwarePrice = "Free"
			}
		case "ratingValue":
			appsdata.RatingValue = mv
		case "reviewCount":
			appsdata.RatingTotal = mv
		case "contentRating":
			appsdata.ContentRating = mv
		case "description":
			appsdata.Description = mv
		}
	}

	tmpf1 := []string{}
	tmpf3 = make(map[string]string)

	//MORE
	doc.Find("div.BgcNfc").Each(func(i int, n *goquery.Selection) {
		s := strings.TrimSpace(n.Text())
		if _, ok := tmpf3[s]; !ok {
			tmpf1 = append(tmpf1, s)
		}
		tmpf3[s] = s
		return
	})

	tmpf2 := []string{}
	doc.Find("span.htlgb").Each(func(i int, n *goquery.Selection) {
		tmpf2 = append(tmpf2, strings.TrimSpace(n.Text()))
		return
	})

	//MATCH INFOS
	tmpf4 := []string{}
	if len(tmpf2)/2 == len(tmpf1) {
		mt := 0
		for _, _ = range tmpf1 {
			mt += 2
			tmpf4 = append(tmpf4, tmpf2[mt-1])
		}
		if len(tmpf1) == len(tmpf4) {
			tmpf2 = tmpf4
		}
	}

	if false {
		jdat, _ := json.MarshalIndent(tmpf4, "", "\t")
		log.Println(len(tmpf4), string(jdat))
	}

	//META
	tmpf3 = make(map[string]string)
	if len(tmpf1) == len(tmpf2) {
		for kk, vv := range tmpf1 {
			tmpf3[vv] = tmpf2[kk]
		}
	}

	if false {
		var jdata []byte
		jdata, _ = json.MarshalIndent(tmpf1, "", "\t")
		log.Println(len(tmpf1), string(jdata))
		jdata, _ = json.MarshalIndent(tmpf2, "", "\t")
		log.Println(len(tmpf2), string(jdata))
		jdata, _ = json.MarshalIndent(tmpf3, "", "\t")
		log.Println(len(tmpf3), string(jdata))
	}
	//META INFOS
	for mk, mv := range tmpf3 {
		switch mk {
		case "Updated":
			appsdata.DatePublished = mv
		case "In-app Products":
			appsdata.SoftwarePrice = mv
			if appsdata.SoftwarePrice == "0" {
				appsdata.SoftwarePrice = "Free"
			}
		case "Offered By":
			appsdata.Developer = mv
		case "Size":
			appsdata.FileSize = mv
		case "Content Rating":
			appsdata.ContentRating = strings.Replace(mv, `+Learn More`, "", -1)
		case "Current Version":
			appsdata.SoftwareVersion = mv
		case "Installs":
			appsdata.TotalDownloads = mv
		case "Requires Android":
			appsdata.SoftwareOs = mv
			appsdata.Badge = mv
		}
	}

	tmpf1 = []string{}
	//META
	doc.Find("a.hrTbp").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "href" {
				ru, _ := n.Attr("href")
				tmpf1 = append(tmpf1, ru)
			}
		}
	})

	//DEVELOPER::SITE
	if len(tmpf1) >= 3 {
		appsdata.DeveloperSite = tmpf1[2]
	}

	tmpf1 = []string{}
	//MORE META
	m := 5
	doc.Find("span.L2o20d").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "style" && m > 0 {
				rv, _ := n.Attr("title")
				tmpf1 = append(tmpf1, fmt.Sprintf("%d Star is %s", m, rv))
				m--
			}
		}
	})

	appsdata.RatingPerStar = strings.Join(tmpf1[:], "; ")

	//TITLE
	doc.Find("title#main-title").Each(func(i int, n *goquery.Selection) {
		appsdata.Title = strings.TrimSpace(n.Text())
		return
	})
	//META-DESC
	doc.Find("div.meta-info").Each(func(i int, n *goquery.Selection) {
		appsdata.MetaDesc += "\r\n" + n.Text()
	})
	appsdata.MetaDesc = strings.TrimSpace(appsdata.MetaDesc)
	//DEVELOPER INFO
	doc.Find("a.dev-link").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "href" && strings.Contains(v.Val, "http") {
				appsdata.DeveloperSite = strings.TrimSpace(html.UnescapeString(v.Val))
			}
		}
	})

	//META
	doc.Find("meta").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "property" && v.Val == "og:url" {
				appsdata.AppURL, _ = n.Attr("content")
				appsdata.AppURL = html.UnescapeString(appsdata.AppURL)
				return
			}
		}
	})

	//give it
	return appsdata
}

//FormatIOS format meta info for Itunes Store
func (metainfo AppsMeta) FormatIOS(doc *goquery.Document, store *StoreApp) (appsdata *App) {

	tmpf1 := []string{}
	tmpf2 := []string{}
	//init
	appsdata = &App{AppID: store.StoreID, AppURL: store.URL, Platform: store.OS}

	//TITLE
	doc.Find("title").Each(func(i int, n *goquery.Selection) {
		appsdata.Title = strings.TrimSpace(n.Text())
		return
	})
	//SUB-TITLE
	doc.Find("div").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "id" && v.Val == "title" {
				appsdata.Title += strings.TrimSpace(n.Text())
			}
		}
	})
	appsdata.Title = strings.TrimSpace(appsdata.Title)
	//DESC
	descT := 0
	doc.Find("span.we-clamp__contents").Each(func(i int, n *goquery.Selection) {
		descT++
		if descT < 3 {
			appsdata.Description += strings.TrimSpace(n.Text())
		}
		return
	})
	//BADGE
	tmpf1 = []string{}
	tmpf2 = []string{}
	doc.Find("dt.information-list__item__term").Each(func(i int, n *goquery.Selection) {
		tmpf1 = append(tmpf1, strings.TrimSpace(n.Text()))
		return
	})
	doc.Find("dd.information-list__item__definition").Each(func(i int, n *goquery.Selection) {
		tmpf2 = append(tmpf2, strings.TrimSpace(n.Text()))
		return
	})

	//DEVELOPER
	if len(tmpf2) >= 1 {
		appsdata.Developer = tmpf2[0]
	}

	//BADGE
	if len(tmpf2) >= 4 {
		appsdata.Badge = tmpf2[3]
		appsdata.SoftwareOs = tmpf2[3]
	}
	//GENRE
	if len(tmpf2) >= 3 {
		appsdata.Genre = tmpf2[2]
	}
	//FILE-SIZE
	if len(tmpf2) >= 2 {
		appsdata.FileSize = tmpf2[1]
	}
	//PRICE
	if len(tmpf2) >= 8 {
		appsdata.SoftwarePrice = tmpf2[7]
	}
	//CONTENT-RATING
	if len(tmpf2) >= 6 {
		appsdata.ContentRating = tmpf2[5]
	}
	//RATING-DESC
	doc.Find("div.we-customer-ratings__count").Each(func(i int, n *goquery.Selection) {
		appsdata.RatingDesc = strings.TrimSpace(n.Text())
		appsdata.RatingTotal = strings.Replace(appsdata.RatingDesc, "Ratings", "", -1)
		appsdata.RatingTotal = strings.Replace(appsdata.RatingTotal, " ", "", -1)
		return
	})
	//RATING-VALUE
	doc.Find("span.we-customer-ratings__averages__display").Each(func(i int, n *goquery.Selection) {
		appsdata.RatingValue = strings.TrimSpace(n.Text())
		return
	})
	//RATING-STAR
	stars := []string{}
	starsT := 5
	doc.Find("div.we-star-bar-graph__bar__foreground-bar").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "style" {
				strr, _ := n.Attr("style")
				strr = strings.Replace(strr, "width:", fmt.Sprintf("%d Star is ", starsT), -1)
				stars = append(stars, strr)
				starsT--
			}
		}
	})
	appsdata.RatingPerStar = strings.Join(stars[:], " ")
	//TOTAL-DOWNLOADS
	totald := 0.0
	if strings.Contains(appsdata.RatingTotal, "K") {
		totalStr := strings.Replace(appsdata.RatingTotal, "K", "", -1)

		if s, err := strconv.ParseFloat(totalStr, 64); err == nil {
			totald = s * 1000
		}
	} else if strings.Contains(appsdata.RatingTotal, "M") {
		totalStr := strings.Replace(appsdata.RatingTotal, "M", "", -1)

		if s, err := strconv.ParseFloat(totalStr, 64); err == nil {
			totald = s * 1000000
		}
	} else {
		if s, err := strconv.ParseFloat(appsdata.RatingTotal, 64); err == nil {
			totald = s * 1.0
		}

	}

	//TOTAL DL
	appsdata.TotalDownloads = fmt.Sprintf("%.0f", (10 * totald))

	//DATE-PUBLISHED
	dt := 0
	doc.Find("time.version-history__item__release-date").Each(func(i int, n *goquery.Selection) {
		if dt < 1 {
			appsdata.DatePublished = strings.TrimSpace(n.Text())
		}
		dt++
		return
	})

	//META
	metas := []string{}
	doc.Find("ol.list").Find("span").Each(func(i int, n *goquery.Selection) {
		metas = append(metas, strings.TrimSpace(n.Text()))
	})
	appsdata.MetaDesc = strings.TrimSpace(strings.Join(metas[:], "\n"))

	//DEVELOPER INFO
	dt = 0
	doc.Find("li.link-list__item").Find("a.link").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "href" {
				if dt < 1 {
					appsdata.DeveloperSite = strings.TrimSpace(v.Val)
				}
				dt++
				break
			}
		}
		return
	})
	//APP-RATING
	doc.Find("div.app-rating").Each(func(i int, n *goquery.Selection) {
		appsdata.ContentRating = strings.TrimSpace(n.Text())
		return
	})
	//APP-URL
	doc.Find("link").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "rel" && v.Val == "canonical" {
				appsdata.AppURL, _ = n.Attr("href")
				appsdata.AppURL = html.UnescapeString(appsdata.AppURL)
				return
			}
		}
	})

	//give it
	return appsdata
}

//Show print results of the meta info
func (metainfo AppsMeta) Show(doneFlg chan bool, wg *sync.WaitGroup) {

	go func() {
		for {
			select {
			//wait till doneFlag has value ;-)
			case <-doneFlg:
				//done already ;-)
				wg.Done()
				return
			}
		}
	}()

	for {
		row, more := <-pAppsData
		if !more {
			break
		}
		//sig-check
		if !pStillRunning {
			log.Println("Signal detected ...")
			break
		}
		pAppList = append(pAppList, row)
	}

	//send signal -> DONE
	doneFlg <- true
}

//ShowCategories show the list of categories for both stores
func (metainfo AppsMeta) ShowCategories(os string) []string {
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
				categlist = append(categlist, fmt.Sprintf("%d. %s", t, categ))
			}
		}
	}
	//json fmt
	return categlist
}

//ShowlistApps show the list of apps per category for both stores
func (metainfo AppsMeta) ShowlistApps(os, category string) []*StoreApp {

	categ := ""
	var appstores []*StoreApp

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
				appstores = metainfo.PrintList(os, category, v)
				t++
				break
			}

		}
	}
	if t == 0 {
		log.Println("\n\n", os, "Category not found!", category)
	}
	return appstores
}

//PrintList show the list of the categories
func (metainfo AppsMeta) PrintList(os, category, url string) []*StoreApp {
	var storelist []*StoreApp
	status, body := getResult(url)
	if status != 200 || body == "" {
		log.Println("ERROR: invalid http status", status)
		return storelist
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Println("ERROR: ", err)
		return storelist
	}

	var storeDone = make(map[string]string)

	appt := 0
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
						if _, ok := storeDone[storeid]; !ok {
							storelist = append(storelist, &StoreApp{Preview: strings.TrimSpace(v.Val), Category: category, StoreID: strings.Replace(storeid, "id", "", -1)})
							appt++
							storeDone[storeid] = storeid
						}
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
						if _, ok := storeDone[storeid]; !ok {
							storelist = append(storelist, &StoreApp{Preview: "https://play.google.com" + strings.TrimSpace(v.Val), Category: category, StoreID: strings.TrimSpace(storeid)})
							appt++
							storeDone[storeid] = storeid
						}
					}

				}
			}
		})
	}
	//give it back
	return storelist
}

//showCategory shows list of categories or list of apps per category
func showCategory(metainfo AppsMeta, os, category string) string {
	var jdata []byte
	if len(category) > 0 {
		res := metainfo.ShowlistApps(os, category)
		jdata, _ = json.MarshalIndent(res, "", "\t")
	} else {
		res := metainfo.ShowCategories(os)
		jdata, _ = json.MarshalIndent(res, "", "\t")
	}
	//show
	return string(jdata)
}

//getResult http req a url
func getResult(url string) (int, string) {
	//client
	c := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true, RootCAs: pool},
		Dial: (&net.Dialer{
			Timeout: 300 * time.Second,
		}).Dial,
		//DisableKeepAlives: true,
	},
	}
	res, err := c.Get(url)
	//make sure to free-up
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		log.Println("ERROR: getResult:", err)
		return 0, ""
	}
	//get response
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ERROR: getResult:", err)
		return 0, ""
	}
	//give
	return res.StatusCode, string(robots)
}
