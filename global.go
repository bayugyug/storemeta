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

type StoreApp struct {
	StoreID string `json:"store_id,omitempty"`
	URL     string `json:"url,omitempty"`
	OS      string `json:"os,omitempty"`
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

	pIOSList     = 0
	pAndroidList = 0

	pIOSCategory     = ""
	pAndroidCategory = ""

	//envt
	pEnvVars = map[string]*string{
		"GMONGERS_LDIR": &pLogDir,
	}

	//print_format
	pPrintFormat = "json"
	pHelp        = false
)

var pAppsData = make(chan *App)
var pStores []*StoreApp

type logOverride struct {
	Prefix string `json:"prefix,omitempty"`
}

func init() {
	//uniqueness
	rand.Seed(time.Now().UnixNano())
	//recovery
	initRecov()
	//re-fmt logger
	overrideLogger("")
	//evt
	initEnvParams()
	//loggers
	initLogger(os.Stdout, os.Stdout, os.Stderr)
	//stats
	//pStats = StatsHelperNew()
	//signals
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

	flag.IntVar(&pIOSList, "list-category-ios", pIOSList, usageIOSList)
	flag.IntVar(&pIOSList, "li", pIOSList, usageIOSList+" (shorthand)")

	flag.IntVar(&pAndroidList, "list-category-android", pAndroidList, usageAndroidList)
	flag.IntVar(&pAndroidList, "la", pAndroidList, usageAndroidList+" (shorthand)")

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
		pIOSList == 0 && pAndroidList == 0 &&
		pIOSCategory == "" && pAndroidCategory == "" {
		showUsage()
	}
	if pAndroidStoreId != "" {
		sts := strings.Split(pAndroidStoreId, ",")
		for _, s := range sts {
			pStores = append(pStores, &StoreApp{OS: ANDROID, URL: "https://play.google.com/store/apps/details?id=" + s + "&hl=en", StoreID: s})
		}
	}
	if pIOSStoreId != "" {
		sts := strings.Split(pIOSStoreId, ",")
		for _, s := range sts {
			pStores = append(pStores, &StoreApp{OS: IOS, URL: "https://itunes.apple.com/app/id" + s + "?mt=8", StoreID: s})
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

		./storemeta -a <AndroidStoreID>  -i <IOSStoreID>

		./storemeta -list-category-android=1

		./storemeta -list-category-ios=1

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
	log.Println("Ver:", pVersion, "\n")
	flag.PrintDefaults()
	log.Println(msg)
	os.Exit(0)
}
