package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
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
)

func b(s string) *bufio.Reader { return bufio.NewReader(strings.NewReader(s)) }

//AppsMetaInfos skeleton interface
type AppsMetaInfos interface {
	Process(chan bool, *sync.WaitGroup, int, *StoreApp)
	Show(chan bool, *sync.WaitGroup)
	FormatAndroid(*goquery.Document, *StoreApp) *App
	FormatIOS(*goquery.Document, *StoreApp) *App
	ShowCategories(string)
	ShowlistApps(string, string)
	PrintList(string, string, string)
}

//AppsMeta empty holder
type AppsMeta struct{}

//handler entry of the app
func handler(metainfo AppsMetaInfos) {

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

	//init
	appsdata = &App{AppID: store.StoreID, AppURL: store.URL, Platform: store.OS}

	//PLAYSTORE
	datePubFmt := "2 January 2006"

	//TITLE
	doc.Find("title#main-title").Each(func(i int, n *goquery.Selection) {
		appsdata.Title = strings.TrimSpace(n.Text())
		return
	})
	//AUTHOR
	doc.Find("span").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			//AUTHOR
			if v.Key == "itemprop" && v.Val == "name" {
				appsdata.Developer = strings.TrimSpace(n.Text())
			}
			//GENRE
			if v.Key == "itemprop" && v.Val == "genre" {
				appsdata.Genre = html.UnescapeString(strings.ToUpper(strings.TrimSpace(n.Text())))
			}
		}
	})
	//GENRE
	doc.Find("a.document-subtitle").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			//GENRE
			if v.Key == "href" && strings.ContainsAny(v.Val, "/store/apps/category") {
				genres := strings.Split(strings.TrimSpace(v.Val), "/")
				if len(genres) >= 5 {
					appsdata.Genre = html.UnescapeString(strings.ToUpper(genres[4]))
				}
				return
			}
		}
	})
	//DESC
	doc.Find("div.show-more-content").Each(func(i int, n *goquery.Selection) {
		appsdata.Description = strings.TrimSpace(n.Text())
		return
	})
	//BADGE
	doc.Find("span.badge-title").Each(func(i int, n *goquery.Selection) {
		appsdata.Badge = strings.TrimSpace(n.Text())
		return
	})
	//RATINGS
	doc.Find("span.reviews-num").Each(func(i int, n *goquery.Selection) {
		rtotal := strings.Replace(n.Text(), ",", "", -1)
		appsdata.RatingTotal = rtotal
		return
	})
	//RATINGS-PER-STAR
	pstars := []string{}
	doc.Find("span.bar-number").Each(func(i int, n *goquery.Selection) {
		ptotal := strings.Replace(n.Text(), ",", "", -1)
		pstars = append(pstars, ptotal)
	})
	appsdata.RatingPerStar = strings.Join(pstars[:], ",")
	//RATING-STAR
	xstar := make(map[string]string)
	stars := []string{}
	doc.Find("div.tiny-star").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "aria-label" {
				if cstr, _ := n.Attr("class"); cstr == "tiny-star star-rating-non-editable-container" {
					s := strings.TrimSpace(strings.Replace(v.Val, ",", "", -1))
					if _, ok := xstar[s]; !ok {
						stars = append(stars, s)
						xstar[s] = s
					}
				}
			}
		}
	})
	appsdata.RatingDesc = strings.Join(stars[:], "\n")
	//META
	doc.Find("meta").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "itemprop" && strings.ToLower(v.Val) == "ratingvalue" {
				rv, _ := n.Attr("content")
				appsdata.RatingValue = strings.TrimSpace(rv)
			}
			if v.Key == "itemprop" && strings.ToLower(v.Val) == "price" {
				rv, _ := n.Attr("content")
				appsdata.SoftwarePrice = strings.TrimSpace(rv)
			}
		}
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
	//CONTENT
	re := regexp.MustCompile("[^0-9\\.]+")
	doc.Find("div.content").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			//AUTHOR
			if v.Key == "itemprop" && strings.ToLower(v.Val) == "contentrating" {
				appsdata.ContentRating = strings.TrimSpace(n.Text())
			}
			//DATE
			if v.Key == "itemprop" && strings.ToLower(v.Val) == "filesize" {
				sbytes := re.ReplaceAllString(strings.TrimSpace(n.Text()), "")
				appsdata.FileSize = strings.TrimSpace(sbytes)
			}
			//GENRE
			if v.Key == "itemprop" && strings.ToLower(v.Val) == "datepublished" {
				t, e := time.Parse(datePubFmt, strings.TrimSpace(n.Text()))
				if e == nil {
					appsdata.DatePublished = t.Format("2006-01-02") + " 00:00:00"
				}
			}
			//SOFTWARE-VERSION
			if v.Key == "itemprop" && strings.ToLower(v.Val) == "numdownloads" {
				sz := strings.SplitN(n.Text(), "-", 2)
				if len(sz) >= 2 {
					fr, _ := strconv.Atoi(strings.TrimSpace(strings.Replace(sz[0], ",", "", -1)))
					to, _ := strconv.Atoi(strings.TrimSpace(strings.Replace(sz[1], ",", "", -1)))
					appsdata.TotalDownloads = fmt.Sprintf("%d", ((fr + to) / 2))
				}
			}
			//OS
			if v.Key == "itemprop" && strings.ToLower(v.Val) == "softwareversion" {
				appsdata.SoftwareVersion = strings.TrimSpace(n.Text())
			}
			//RATINGS
			if v.Key == "itemprop" && strings.ToLower(v.Val) == "operatingsystems" {
				appsdata.SoftwareOs = strings.TrimSpace(n.Text())
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

	//init
	appsdata = &App{AppID: store.StoreID, AppURL: store.URL, Platform: store.OS}

	//ITUNES
	datePubFmt := "Jan 2, 2006"
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
	doc.Find("p").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "itemprop" && v.Val == "description" {
				appsdata.Description = strings.TrimSpace(n.Text())
				return
			}
		}
	})
	//BADGE
	doc.Find("div.fat-binary-blurb").Each(func(i int, n *goquery.Selection) {
		appsdata.Badge = strings.TrimSpace(n.Text())
		return
	})
	//RATINGS-PER
	totr8 := 0
	pstars := []string{}
	doc.Find("div.rating span.rating-count").Each(func(i int, n *goquery.Selection) {
		prate := strings.Replace(n.Text(), "Ratings", "", -1)
		prate = strings.TrimSpace(strings.Replace(prate, ",", "", -1))
		pstars = append(pstars, prate)
		rx, _ := strconv.Atoi(prate)
		totr8 += rx
	})
	//TOTAL DL
	appsdata.RatingPerStar = strings.Join(pstars[:], ",")
	appsdata.RatingTotal = fmt.Sprintf("%d", totr8)
	appsdata.TotalDownloads = fmt.Sprintf("%d", (10 * totr8))
	//RATING-STAR
	stars := []string{}
	doc.Find("div.rating").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "itemprop" && v.Val == "aggregateRating" {
				strr, _ := n.Attr("aria-label")
				stars = append(stars, strr)
			}
		}
	})
	appsdata.RatingDesc = strings.Join(stars[:], "\n")
	//META
	metas := []string{}
	doc.Find("ol.list").Find("span").Each(func(i int, n *goquery.Selection) {
		metas = append(metas, strings.TrimSpace(n.Text()))
	})
	appsdata.MetaDesc = strings.TrimSpace(strings.Join(metas[:], "\n"))
	//DEVELOPER INFO
	doc.Find("div.app-links").Find("a").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "href" {
				if len(appsdata.DeveloperSite) == 0 {
					appsdata.DeveloperSite = strings.TrimSpace(html.UnescapeString(v.Val))
				}
				break
			}
		}
	})
	//AUTHOR
	doc.Find("span").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			//AUTHOR
			if v.Key == "itemprop" && v.Val == "name" {
				appsdata.Developer = strings.TrimSpace(n.Text())
			}
			//DATE
			if v.Key == "itemprop" && v.Val == "datePublished" {
				t, e := time.Parse(datePubFmt, strings.TrimSpace(n.Text()))
				if e == nil {
					appsdata.DatePublished = t.Format("2006-01-02") + " 00:00:00"
				}
			}
			//GENRE
			if v.Key == "itemprop" && v.Val == "applicationCategory" {
				appsdata.Genre = html.UnescapeString(strings.ToUpper(strings.TrimSpace(n.Text())))
			}
			//SOFTWARE-VERSION
			if v.Key == "itemprop" && v.Val == "softwareVersion" {
				appsdata.SoftwareVersion = strings.TrimSpace(n.Text())
			}
			//OS
			if v.Key == "itemprop" && v.Val == "operatingSystem" {
				appsdata.SoftwareOs = strings.TrimSpace(n.Text())
			}
			//RATINGS
			if v.Key == "itemprop" && v.Val == "ratingValue" {
				appsdata.RatingValue = strings.TrimSpace(n.Text())
			}
		}
	})
	//GENRE
	doc.Find("a.genre").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			//GENRE
			if v.Key == "href" && strings.ContainsAny(v.Val, "https://itunes.apple.com/") && strings.ContainsAny(v.Val, "/genre/") {
				genres := strings.Split(strings.TrimSpace(v.Val), "/")
				if len(genres) >= 6 {
					appsdata.Genre = html.UnescapeString(strings.Replace(strings.ToUpper(genres[5]), "IOS-", "", -1))
				}
			}
		}
	})
	//PRICE
	doc.Find("div").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "itemprop" && v.Val == "price" {
				appsdata.SoftwarePrice = strings.TrimSpace(n.Text())
				return
			}
		}
	})
	//SIZE
	re := regexp.MustCompile("[^0-9\\.]+")
	doc.Find("div#left-stack").Find("li").Each(func(i int, n *goquery.Selection) {
		if s := strings.TrimSpace(n.Text()); strings.Contains(s, "Size:") {
			sz := strings.SplitN(s, ":", 2)
			if len(sz) > 1 {
				sbytes := re.ReplaceAllString(strings.TrimSpace(sz[1]), "")
				appsdata.FileSize = strings.TrimSpace(sbytes)
				return
			}
		}
	})
	//SUB-DESC
	smeta := []string{}
	doc.Find("ul.list").Each(func(i int, n *goquery.Selection) {
		for _, v := range n.Nodes[0].Attr {
			if v.Key == "class" && strings.Contains(v.Val, "app-rating-reasons") {
				smeta = append(smeta, strings.TrimSpace(n.Text()))
			}
		}

	})
	appsdata.MetaDesc = strings.TrimSpace(strings.Join(smeta[:], "\n"))
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

	var applist []*App
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
		applist = append(applist, row)
	}

	//json fmt
	jdata, _ := json.MarshalIndent(applist, "", "\t")
	//dont leave your friend behind :-)
	log.Println(string(jdata))
	//send signal -> DONE
	doneFlg <- true
}

//ShowCategories show the list of categories for both stores
func (metainfo AppsMeta) ShowCategories(os string) {
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

//ShowlistApps show the list of apps per category for both stores
func (metainfo AppsMeta) ShowlistApps(os, category string) {

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
				metainfo.PrintList(os, category, v)
				t++
				break
			}

		}
	}
	if t == 0 {
		log.Println("\n\n", os, "Category not found!", category)
	}
}

//PrintList show the list of the categories
func (metainfo AppsMeta) PrintList(os, category, url string) {
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

//showCategory shows list of categories or list of apps per category
func showCategory(metainfo AppsMeta, os, category string) {
	if len(category) > 0 {
		metainfo.ShowlistApps(os, category)
	} else {
		metainfo.ShowCategories(os)
	}
}

//getResult http req a url
func getResult(url string) (int, string) {
	//client
	/*
		c := &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 300 * time.Second,
				}).Dial,
				TLSClientConfig: &tls.Config{RootCAs: pool},
			},
		}
	*/
	c := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true, RootCAs: pool},
		Dial: (&net.Dialer{
			Timeout: 300 * time.Second,
		}).Dial,
	},
	}
	res, err := c.Get(url)
	if err != nil {
		log.Println("ERROR: getResult:", err)
		return 0, ""
	}
	//get response
	robots, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Println("ERROR: getResult:", err)
		return 0, ""
	}
	//give
	return res.StatusCode, string(robots)
}
