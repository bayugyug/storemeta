package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path"
	//"path/filepath"
	"crypto/x509"
	"regexp"
	"strings"
	"time"
)

const (
	usageShowConsole     = "use to enable the output in console"
	usageAndroidStore    = "use for querying the Android App Store"
	usageIOSStore        = "use for querying the IOS App Store"
	usageAndroidList     = "use for querying the category list of apps in Android App Store"
	usageIOSList         = "use for querying the category list of apps in IOS App Store"
	usageAndroidCategory = "use for querying the list of apps per category in Android App Store"
	usageIOSCategory     = "use for querying the list of apps per category in IOS App Store"
	usagePrintFormat     = "use to enable what format is used in showing the output"
	IOS                  = "IOS"
	ANDROID              = "ANDROID"
)

//Store data holder for ios data
type StoreApp struct {
	OS       string `json:"os,omitempty"`
	StoreID  string `json:"store_id,omitempty"`
	Preview  string `json:"preview,omitempty"`
	Category string `json:"category,omitempty"`
	Name     string `json:"name,omitempty"`
	URL      string `json:"url,omitempty"`
}

type App struct {
	Platform        string `json:"platform"`
	AppURL          string `json:"app-url"`
	AppID           string `json:"app-id"`
	Genre           string `json:"genre"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Badge           string `json:"badge"`
	RatingTotal     string `json:"rating-total"`
	RatingPerStar   string `json:"rating-per-star"`
	RatingDesc      string `json:"rating-desc"`
	RatingValue     string `json:"rating-value"`
	SoftwarePrice   string `json:"software-price"`
	MetaDesc        string `json:"meta-desc"`
	FileSize        string `json:"file-size"`
	ContentRating   string `json:"content-rating"`
	DatePublished   string `json:"date-published"`
	SoftwareVersion string `json:"software-version"`
	SoftwareOs      string `json:"software-os"`
	TotalDownloads  string `json:"total-downloads"`
	Developer       string `json:"developer"`
	DeveloperSite   string `json:"developer-site"`
}

var (
	pLogDir = "."
	//loggers
	infoLog  *log.Logger
	warnLog  *log.Logger
	errorLog *log.Logger
	//stats
	//pStats *StatsHelper
	//signal flag
	pStillRunning = true

	pBuildTime = "0"
	pVersion   = "0.1.0" + "-" + pBuildTime
	//console
	pShowConsole    = true
	pAndroidStoreId = ""
	pIOSStoreId     = ""

	pIOSList     = false
	pAndroidList = false

	pIOSCategory     = ""
	pAndroidCategory = ""

	//envt
	pEnvVars = map[string]*string{
		"GMONGERS_LDIR": &pLogDir,
	}

	//print_format
	pPrintFormat = "json"
	pHelp        = false

	pStoreURI = map[string][]string{
		ANDROID: {"https://play.google.com/store/apps/details?id=", "&hl=en"},
		IOS:     {"https://itunes.apple.com/app/id", "?mt=8"},
	}
	//ssl certs
	pool *x509.CertPool
)

type logOverride struct {
	Prefix string `json:"prefix,omitempty"`
}

var pAppsData = make(chan *App)
var pStores []*StoreApp

var pCategories = map[string][]string{
	ANDROID: {
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
	IOS: {
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

func init() {
	//uniqueness
	rand.Seed(time.Now().UnixNano())
	//recovery
	initRecov()
	//re-fmt logger
	//overrideLogger("")
	//evt
	initEnvParams()
	//loggers
	initLogger(os.Stdout, os.Stdout, os.Stderr)

	//init certs
	pool = x509.NewCertPool()
	pool.AppendCertsFromPEM(pemCerts)
}

//initRecov is for dumpIng segv in
func initRecov() {
	//might help u
	defer func() {
		recvr := recover()
		if recvr != nil {
			fmt.Println("MAIN-RECOV-INIT: ", recvr)
		}
	}()
}

//os.Stdout, os.Stdout, os.Stderr
func initLogger(i, w, e io.Writer) {
	//just in case
	if !pShowConsole {
		infoLog = makeLogger(i, pLogDir, "storemeta", "INFO: ")
		warnLog = makeLogger(w, pLogDir, "storemeta", "WARN: ")
		errorLog = makeLogger(e, pLogDir, "storemeta", "ERROR: ")
	} else {
		infoLog = log.New(i,
			"INFO: ",
			log.Ldate|log.Ltime|log.Lmicroseconds)
		warnLog = log.New(w,
			"WARN: ",
			log.Ldate|log.Ltime|log.Lshortfile)
		errorLog = log.New(e,
			"ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile)
	}
}

//initEnvParams enable all OS envt vars to reload internally
func initEnvParams() {
	//just in-case, over-write from ENV
	for k, v := range pEnvVars {
		if os.Getenv(k) != "" {
			*v = os.Getenv(k)
		}
	}
	flag.BoolVar(&pShowConsole, "debug", pShowConsole, usageShowConsole)
	flag.BoolVar(&pShowConsole, "d", pShowConsole, usageShowConsole+" (shorthand)")

	flag.StringVar(&pAndroidStoreId, "android", pAndroidStoreId, usageAndroidStore)
	flag.StringVar(&pAndroidStoreId, "a", pAndroidStoreId, usageAndroidStore+" (shorthand)")

	flag.StringVar(&pIOSStoreId, "ios", pIOSStoreId, usageIOSStore)
	flag.StringVar(&pIOSStoreId, "i", pIOSStoreId, usageIOSStore+" (shorthand)")

	flag.BoolVar(&pIOSList, "list-category-ios", pIOSList, usageIOSList)
	flag.BoolVar(&pIOSList, "li", pIOSList, usageIOSList+" (shorthand)")

	flag.BoolVar(&pAndroidList, "list-category-android", pAndroidList, usageAndroidList)
	flag.BoolVar(&pAndroidList, "la", pAndroidList, usageAndroidList+" (shorthand)")

	flag.StringVar(&pIOSCategory, "category-ios", pIOSCategory, usageIOSCategory)
	flag.StringVar(&pIOSCategory, "ci", pIOSCategory, usageIOSCategory+" (shorthand)")

	flag.StringVar(&pAndroidCategory, "category-android", pAndroidCategory, usageIOSCategory)
	flag.StringVar(&pAndroidCategory, "ca", pAndroidCategory, usageIOSCategory+" (shorthand)")

	flag.StringVar(&pPrintFormat, "print-format", pPrintFormat, usagePrintFormat)
	flag.StringVar(&pPrintFormat, "pf", pPrintFormat, usagePrintFormat+" (shorthand)")

	flag.BoolVar(&pHelp, "h", pHelp, "Show this help/how-to")
	flag.Parse()

	//either 1 should be present
	if pHelp {
		showUsage()
	}
	if pIOSStoreId == "" && pAndroidStoreId == "" &&
		!pIOSList && !pAndroidList &&
		pIOSCategory == "" && pAndroidCategory == "" {
		showUsage()
	}
	if pAndroidStoreId != "" {
		sts := strings.Split(pAndroidStoreId, ",")
		for _, s := range sts {
			pStores = append(pStores, &StoreApp{OS: ANDROID, URL: pStoreURI[ANDROID][0] + s + pStoreURI[ANDROID][1], StoreID: s})
		}
	}
	if pIOSStoreId != "" {
		sts := strings.Split(pIOSStoreId, ",")
		for _, s := range sts {
			pStores = append(pStores, &StoreApp{OS: IOS, URL: pStoreURI[IOS][0] + s + pStoreURI[IOS][1], StoreID: s})
		}
	}

}

//formatLogger try to init all filehandles for logs
func formatLogger(fdir, fname, pfx string) string {
	t := time.Now()
	r := regexp.MustCompile("[^a-zA-Z0-9]")
	p := t.Format("2006-01-02") + "-" + r.ReplaceAllString(strings.ToLower(pfx), "")
	s := path.Join(pLogDir, fdir)
	if _, err := os.Stat(s); os.IsNotExist(err) {
		//mkdir -p
		os.MkdirAll(s, os.ModePerm)
	}
	return path.Join(s, p+"-"+fname+".log")
}

//makeLogger initialize the logger either via file or console
func makeLogger(w io.Writer, ldir, fname, pfx string) *log.Logger {
	logFile := w
	if !pShowConsole {
		var err error
		logFile, err = os.OpenFile(formatLogger(ldir, fname, pfx), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			log.Println(err)
		}
	}
	//give it
	return log.New(logFile,
		pfx,
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

}

//dumpW log into warning
func dumpW(s ...interface{}) {
	warnLog.Println(s...)
}

//dumpWF log into warning w/ fmt
func dumpWF(format string, s ...interface{}) {
	warnLog.Println(fmt.Sprintf(format, s...))
}

//dumpE log into error
func dumpE(s ...interface{}) {
	errorLog.Println(s...)
}

//dumpE log into error w/ fmt
func dumpEF(format string, s ...interface{}) {
	errorLog.Println(fmt.Sprintf(format, s...))
}

//dumpI log into info
func dumpI(s ...interface{}) {
	infoLog.Println(s...)
}

//dumpIF log into info
func dumpIF(format string, s ...interface{}) {
	infoLog.Println(fmt.Sprintf(format, s...))
}

//Write override the log.print
func (w logOverride) Write(bytes []byte) (int, error) {
	//return fmt.Print(w.Prefix + time.Now().UTC().Format("2006-01-02 15:04:05.999") + " " + string(bytes))
	return fmt.Print(string(bytes))
}

//overrideLogger reset the log.print to customized
func overrideLogger(pfx string) {
	log.SetFlags(0)
	log.SetOutput(&logOverride{Prefix: pfx})
}

//showUsage
func showUsage() {

	msg := `



	Example:



		./storemeta -list-category-android

		./storemeta -list-category-ios

		./storemeta -category-android=GAME_ACTION

		./storemeta -category-ios=GAMES_ACTION

		./storemeta  -a="com.google.android.apps.photos"

		or

		./storemeta  -i="293622097"

		or

		./storemeta  -a="com.google.android.apps.photos" -i="293622097"

		or

		./storemeta  -a="com.google.android.apps.plus,com.google.android.launcher,com.sphero.sprk"

		or

		./storemeta  -i="544007664,535886823,643496868"

`
	fmt.Println("Ver:", pVersion, "\n")
	flag.PrintDefaults()
	fmt.Println(msg)
	os.Exit(0)
}
