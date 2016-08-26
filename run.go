package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("webapp"))
	http.Handle("/", fs)
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

 */