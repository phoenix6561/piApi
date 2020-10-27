package apicontroller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	pinservice "github.com/jfox/piApi/src/service"
)

//SetUpAPI ...
func SetUpAPI() {
	fmt.Println("starting  router on port 8080")
	r := mux.NewRouter()

	r.HandleFunc("/api/togglepin/{pin}", togglePinHTTPRequest).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))

}

//TogglePinHTTPRequest ...
func togglePinHTTPRequest(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	var pinNum string = ""
	var err error
	if val, ok := pathParams["pin"]; ok {
		pinNum = pinNum + val
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a pin"}`))
			return
		}
	}

	fmt.Println("toggle pin " + pinNum)
	i, err := strconv.Atoi(pinNum)
	pinservice.TogglePin(i)

}

func allGpoPinStatusHTTPRequest(w http.ResponseWriter, r *http.Request) {

}
