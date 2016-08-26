package main

import (
	"net/http"
	"log"
	"github.com/elazarl/go-bindata-assetfs"
)

const USE_ASSET_FS = true

func main() {

	var fs http.Handler
	if USE_ASSET_FS  {
		fs = http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "webapp"})
	} else {
		fs = http.FileServer(http.Dir("webapp"))
	}


	http.Handle("/", fs)
	log.Println("Running from :8000. Ready to rock!")
	http.ListenAndServe(":8000", nil)
}

/*
	- Render with default lighting
	- Set rotation to 0

	- If ROT_DATA
		Set ROT

	- if TEMP_DATA
		Calculate Colour Temperature
		Set Color to all lights

	- If LUX_DATA
		Calculate Lux
		Set Lux to all lights

	- If LED Turns on or Off
		Turn on or off LED

 */