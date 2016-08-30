package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strings"
)

//Store data holder for ios data
type Store struct {
	Preview  string `json:"preview,omitempty"`
	StoreID  string `json:"store_id,omitempty"`
	Category string `json:"category,omitempty"`
	Name     string `json:"name,omitempty"`
	URL      string `json:"url,omitempty"`
}

var pCategories = map[string][]string{
	"ANDROID": {
		"https://play.google.com/store/apps/category/ANDROID_WEAR",
		"https://play.google.com/store/apps/category/BOOKS_AND_REFERENCE",
		"https://play.google.com/store/apps/category/BUSINESS",
		"https://play.google.com/store/apps/category/COMICS",
		"https://play.google.com/store/apps/category/COMMUNICATION",
		"https://play.google.com/store/apps/category/EDUCATION",
		"https://play.google.com/store/apps/category/ENTERTAINMENT",
		"https://play.google.com/store/apps/category/FINANCE",
		"https://play.google.com/store/apps/category/HEALTH_AND_FITNESS",
		"https://play.google.com/store/apps/category/LIBRARIES_AND_DEMO",
		"https://play.google.com/store/apps/category/LIFESTYLE",
		"https://play.google.com/store/apps/category/MEDIA_AND_VIDEO",
		"https://play.google.com/store/apps/category/MEDICAL",
		"https://play.google.com/store/apps/category/MUSIC_AND_AUDIO",
		"https://play.google.com/store/apps/category/NEWS_AND_MAGAZINES",
		"https://play.google.com/store/apps/category/PERSONALIZATION",
		"https://play.google.com/store/apps/category/PHOTOGRAPHY",
		"https://play.google.com/store/apps/category/PRODUCTIVITY",
		"https://play.google.com/store/apps/category/SHOPPING",
		"https://play.google.com/store/apps/category/SOCIAL",
		"https://play.google.com/store/apps/category/SPORTS",
		"https://play.google.com/store/apps/category/TOOLS",
		"https://play.google.com/store/apps/category/TRANSPORTATION",
		"https://play.google.com/store/apps/category/TRAVEL_AND_LOCAL",
		"https://play.google.com/store/apps/category/WEATHER",
		"https://play.google.com/store/apps/category/GAME",
		"https://play.google.com/store/apps/category/GAME_ACTION",
		"https://play.google.com/store/apps/category/GAME_ADVENTURE",
		"https://play.google.com/store/apps/category/GAME_ARCADE",
		"https://play.google.com/store/apps/category/GAME_BOARD",
		"https://play.google.com/store/apps/category/GAME_CARD",
		"https://play.google.com/store/apps/category/GAME_CASINO",
		"https://play.google.com/store/apps/category/GAME_CASUAL",
		"https://play.google.com/store/apps/category/GAME_EDUCATIONAL",
		"https://play.google.com/store/apps/category/GAME_MUSIC",
		"https://play.google.com/store/apps/category/GAME_PUZZLE",
		"https://play.google.com/store/apps/category/GAME_RACING",
		"https://play.google.com/store/apps/category/GAME_ROLE_PLAYING",
		"https://play.google.com/store/apps/category/GAME_SIMULATION",
		"https://play.google.com/store/apps/category/GAME_SPORTS",
		"https://play.google.com/store/apps/category/GAME_STRATEGY",
		"https://play.google.com/store/apps/category/GAME_TRIVIA",
		"https://play.google.com/store/apps/category/GAME_WORD",
		"https://play.google.com/store/apps/category/FAMILY",
		"https://play.google.com/store/apps/category/FAMILY?age=AGE_RANGE1",
		"https://play.google.com/store/apps/category/FAMILY?age=AGE_RANGE2",
		"https://play.google.com/store/apps/category/FAMILY?age=AGE_RANGE3",
		"https://play.google.com/store/apps/category/FAMILY_ACTION",
		"https://play.google.com/store/apps/category/FAMILY_BRAINGAMES",
		"https://play.google.com/store/apps/category/FAMILY_CREATE",
		"https://play.google.com/store/apps/category/FAMILY_EDUCATION",
		"https://play.google.com/store/apps/category/FAMILY_MUSICVIDEO",
		"https://play.google.com/store/apps/category/FAMILY_PRETEND",
	},
	"IOS": {
		"https://itunes.apple.com/us/genre/ios-books/id6018?mt=8",
		"https://itunes.apple.com/us/genre/ios-business/id6000?mt=8",
		"https://itunes.apple.com/us/genre/ios-education/id6017?mt=8",
		"https://itunes.apple.com/us/genre/ios-entertainment/id6016?mt=8",
		"https://itunes.apple.com/us/genre/ios-finance/id6015?mt=8",
		"https://itunes.apple.com/us/genre/ios-food-drink/id6023?mt=8",
		"https://itunes.apple.com/us/genre/ios-games/id6014?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-action/id7001?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-adventure/id7002?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-arcade/id7003?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-board/id7004?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-card/id7005?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-casino/id7006?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-dice/id7007?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-educational/id7008?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-family/id7009?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-music/id7011?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-puzzle/id7012?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-racing/id7013?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-role-playing/id7014?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-simulation/id7015?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-sports/id7016?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-strategy/id7017?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-trivia/id7018?mt=8",
		"https://itunes.apple.com/us/genre/ios-games-word/id7019?mt=8",
		"https://itunes.apple.com/us/genre/ios-health-fitness/id6013?mt=8",
		"https://itunes.apple.com/us/genre/ios-lifestyle/id6012?mt=8",
		"https://itunes.apple.com/us/genre/ios-magazines-newspapers/id6021?mt=8",
		"https://itunes.apple.com/us/genre/ios-medical/id6020?mt=8",
		"https://itunes.apple.com/us/genre/ios-music/id6011?mt=8",
		"https://itunes.apple.com/us/genre/ios-navigation/id6010?mt=8",
		"https://itunes.apple.com/us/genre/ios-news/id6009?mt=8",
		"https://itunes.apple.com/us/genre/ios-photo-video/id6008?mt=8",
		"https://itunes.apple.com/us/genre/ios-productivity/id6007?mt=8",
		"https://itunes.apple.com/us/genre/ios-reference/id6006?mt=8",
		"https://itunes.apple.com/us/genre/ios-shopping/id6024?mt=8",
		"https://itunes.apple.com/us/genre/ios-social-networking/id6005?mt=8",
		"https://itunes.apple.com/us/genre/ios-sports/id6004?mt=8",
		"https://itunes.apple.com/us/genre/ios-travel/id6003?mt=8",
		"https://itunes.apple.com/us/genre/ios-utilities/id6002?mt=8",
		"https://itunes.apple.com/us/genre/ios-weather/id6001?mt=8",
	},
}

func showAppCategories(os string) {
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

func showListPerCateg(os, category string) {

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
				listCategory(os, category, v)
				t++
				break
			}

		}
	}
	if t == 0 {
		log.Println("\n\n", os, "Category not found!", category)
	}
}

func listCategory(os, category, url string) {
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
	var storelist []*Store
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
						storelist = append(storelist, &Store{Preview: strings.TrimSpace(v.Val), Category: category, StoreID: strings.Replace(storeid, "id", "", -1)})
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
						storelist = append(storelist, &Store{Preview: "https://play.google.com" + strings.TrimSpace(v.Val), Category: category, StoreID: strings.TrimSpace(storeid)})
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
