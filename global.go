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
	usageShowConsole  = "use to enable the output in console"
	usageAndroidStore = "use for querying the Android App Store"
	usageIOSStore     = "use for querying the IOS App Store"
	IOS               = "IOS"
	ANDROID           = "ANDROID"
)

type StoreApp struct {
	StoreID string `json:"store_id,omitempty"`
	URL     string `json:"url,omitempty"`
	OS      string `json:"os,omitempty"`
}

type App struct {
	Platform        string `json:"platform"`
	Title           string `json:"title"`
	Developer       string `json:"developer"`
	DeveloperSite   string `json:"developer-site"`
	Genre           string `json:"genre"`
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
	AppURL          string `json:"app-url"`
	AppID           string `json:"app-id"`
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

	pBuildTime = ""
	pVersion   = "0.1.0" + "-" + pBuildTime
	//console
	pShowConsole    = true
	pAndroidStoreId = ""
	pIOSStoreId     = ""
	//envt
	pEnvVars = map[string]*string{
		"GMONGERS_LDIR": &pLogDir,
	}
)

var pStores []*StoreApp

type logOverride struct {
	Prefix string `json:"prefix,omitempty"`
}

func init() {
	//uniqueness
	rand.Seed(time.Now().UnixNano())
	//recovery
	initRecov()
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
	flag.Parse()

	//either 1 should be present
	if pIOSStoreId == "" && pAndroidStoreId == "" {
		showUsage()
		os.Exit(0)
	}
	if pAndroidStoreId != "" {
		pStores = append(pStores, &StoreApp{OS: ANDROID, URL: "https://play.google.com/store/apps/details?id=" + pAndroidStoreId + "&hl=en", StoreID: pAndroidStoreId})
	}
	if pIOSStoreId != "" {
		pStores = append(pStores, &StoreApp{OS: IOS, URL: "https://itunes.apple.com/app/id" + pIOSStoreId + "?mt=8", StoreID: pIOSStoreId})
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

		./storemeta -a <AndroidStoreID>  -i <IOSStoreID>


	Example:

		./storemeta  -a="com.google.android.apps.photos"

		or

		./storemeta  -i="293622097"

		or

		./storemeta  -a="com.google.android.apps.photos" -i="293622097"


`
	fmt.Println(msg)
}
