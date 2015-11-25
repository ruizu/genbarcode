package main

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code39"
	"github.com/julienschmidt/httprouter"
)

func Ping(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "pong")
}

func HandleCode39(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	data := params.ByName("data")
	ext := strings.ToLower(filepath.Ext(data))
	data = strings.ToUpper(strings.Replace(data, ext, "", 1))

	// generate barcode
	bc, err := code39.Encode(data, false, true)
	if err != nil {
		log.Panic(err)
	}

	// resize barcode
	size := r.FormValue("size")
	switch size {
	case "s":
		bc, err = barcode.Scale(bc, 150, 45)
		if err != nil {
			log.Panic(err)
		}
	case "l":
		bc, err = barcode.Scale(bc, 600, 180)
		if err != nil {
			log.Panic(err)
		}
	default:
		bc, err = barcode.Scale(bc, 300, 90)
		if err != nil {
			log.Panic(err)
		}
	}

	// print barcode
	switch ext {
	case ".jpg":
		if err := jpeg.Encode(w, bc, nil); err != nil {
			log.Panic(err)
		}
	case ".png":
		if err := png.Encode(w, bc); err != nil {
			log.Panic(err)
		}
	default:
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
}
