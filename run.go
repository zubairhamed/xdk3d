package main

import (
	"log"
	"net/http"
)

const USE_ASSET_FS = true

func main() {

	//var fs http.Handler
	//if USE_ASSET_FS {
	//	fs = http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "webapp"})
	//} else {
	//	fs = http.FileServer(http.Dir("webapp"))
	//}

	fs := http.FileServer(http.Dir("webapp"))

	http.Handle("/", fs)
	log.Println("Running from :8000. Ready to rock!")
	http.ListenAndServe(":8000", nil)
}

/*
	- If ROT_DATA
		Set ROT

	- if TEMP_DATA
		Calculate Colour Temperature
		Set Color to all lights

	- If LUX_DATA
		Calculate Lux
		Set Lux to all lights

*/
