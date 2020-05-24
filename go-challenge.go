package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type ResponseApiCountries struct {
	Country        string `json:"Country"`
	CountryCode    string `json:"CountryCode"`
	TotalConfirmed int    `json:"TotalConfirmed"`
	NewConfirmed   int    `json:"NewConfirmed"`
	TotalDeaths    int    `json:"TotalDeaths"`
}
type ResponseApi struct {
	Countries []ResponseApiCountries `json:"Countries"`
}
type Object struct {
	Data  Country `json:"data"`
	Error *string `json:"error"`
}
type Country struct {
	Country      string `json:"country"`
	Iso          string `json:"iso"`
	Confirmed    int    `json:"confirmed"`
	NewConfirmed int    `json:"newConfirmed"`
	Deaths       int    `json:"deaths"`
}
type error interface {
	Error() string
}

func getData(w http.ResponseWriter, r *http.Request) {
	var responseFinal Object
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get(goDotEnvVariable("URL_API"))
	if err != nil {
		temp := goDotEnvVariable("ERROR")
		responseFinal.Error = &temp

	} else {
		if resp.StatusCode == 429 {
			temp := goDotEnvVariable("MESSAGE_LIMIT")
			responseFinal.Error = &temp
		}
		data, _ := ioutil.ReadAll(resp.Body)

		var responseObject ResponseApi
		json.Unmarshal(data, &responseObject)
		for _, response := range responseObject.Countries {

			if response.CountryCode == goDotEnvVariable("CODE_COUNTRY") {

				responseFinal.Data.Confirmed = response.TotalConfirmed
				responseFinal.Data.Country = response.Country
				responseFinal.Data.Deaths = response.TotalDeaths
				responseFinal.Data.Iso = response.CountryCode
				responseFinal.Data.NewConfirmed = response.NewConfirmed
				responseFinal.Error = nil
			}
		}

		json.NewEncoder(w).Encode(responseFinal)
	}

}

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", getData).Methods("GET")
	log.Fatal(http.ListenAndServe(goDotEnvVariable("PORT"), router))
}
