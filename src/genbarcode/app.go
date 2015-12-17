package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var config Config

func init() {
	version := flag.Bool("version", false, "prints version")
	flag.Parse()

	if *version {
		fmt.Println(Version())
		os.Exit(0)
	}

	ok := ReadConfig(&config, "config/genbarcode.ini") ||
		ReadConfig(&config, "/etc/genbarcode/genbarcode.ini") ||
		ReadConfig(&config, "/etc/genbarcode.ini")
	if !ok {
		log.Fatal("Could not find configuration file")
	}
}

func main() {
	router := httprouter.New()
	router.GET("/ping", Ping)
	router.GET("/barcode/code39/:data", HandleCode39)

	log.Fatal(http.ListenAndServe(config.Server.Host, router))
}
