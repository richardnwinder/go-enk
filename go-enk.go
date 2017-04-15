package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/richardnwinder/mpic"
)

var dev *mpic.Device

func main() {
	var err error
	var version int
	var release int
	fmt.Print("\n > Open ENK Device\n")
	dev, err = mpic.Open()
	if err != nil {
		fmt.Printf(" ... ERROR: ENK - %s\n", err)
	} else {
		fmt.Print("\n > Claim Interface\n")
		err = dev.ClaimInterface(0)
		if err != nil {
			fmt.Printf(" ... ERROR: ENK - %s\n", err)
		} else {
			fmt.Print(" ... OK: ENK - interace 0 set\n")
			fmt.Print(" ... OK: ENK - connected\n")
			fmt.Print("\n > Get ENK Version\n")
			version, release, err = dev.GetVersion()
			if err != nil {
				fmt.Printf(" ... ERROR: ENK Version - %s\n", err)
			}
			fmt.Printf(" ... OK: ENK Version - version: %d\n", version)
			fmt.Printf(" ... OK: ENK Version - release: %d\n", release)
		}

	}
	router := mux.NewRouter()
	router.HandleFunc("/version", getVersion).Methods("GET")
	router.HandleFunc("/login", activate).Methods("POST")
	http.ListenAndServe(":8083", router)
}
func activate(w http.ResponseWriter, r *http.Request) {

}
func getVersion(w http.ResponseWriter, r *http.Request) {
	version, release, err := dev.GetVersion()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		errorENK := "error: " + fmt.Sprint(err)
		//jsonData, _ := json.Marshall(errorENK)
		fmt.Fprint(w, errorENK)
	}
	versionENK := "{\"version\": " + fmt.Sprint(version) + ", \"release\" :" + fmt.Sprint(release) + "}"
	//jsonData, _ := json.Marshall(versionENK)
	fmt.Fprint(w, versionENK)
}
