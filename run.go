package main

import (
	"log"
	"net/http"
)

const USE_ASSET_FS = true

func main() {
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
