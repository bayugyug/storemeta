package main

import (
	"log"
	"os"
)

func main() {

	//set
	log.Println("Start!")

	//start
	log.Println("IOS:", pIOSStoreId)
	log.Println("ANDROID:", pAndroidStoreId)

	//init
	log.Println("Done!")
	os.Exit(0)
}
