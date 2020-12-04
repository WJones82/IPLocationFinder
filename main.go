package main

import (
	"os"
	"io"
	"encoding/csv"
    "fmt"
	"log"
	"net"
	"net/http"
	"github.com/IncSW/geoip2"   // 3rd party api to read db file
	"github.com/gorilla/mux"   // routing
)

func checkIP(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	reader, err := geoip2.NewCountryReaderFromFile("GeoLite2-Country.mmdb")
	
	if err != nil{
		log.Fatal(err)
	}

	record, err := reader.Lookup(net.ParseIP(vars["ip"]))
	if err != nil {
		panic(err)
	}

	inputCountry := record.Country.Names["en"]	
	
	checkWhiteList(w,inputCountry)
	
}

func checkWhiteList(w http.ResponseWriter, inputCountry string) {
	csvfile, err := os.Open("whitelistedcountries.csv")
	if err != nil{
		log.Fatal("Could not open csv file!")
	}

	r := csv.NewReader(csvfile)
	r.FieldsPerRecord = -1

	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if inputCountry == record[0] {
			fmt.Fprintf(w,"The IP for " + inputCountry + " is whitelisted.") 
			break;  
		}

		if inputCountry != record[0] {
			fmt.Fprintf(w,"The IP for " + inputCountry + " is NOT on the whitelist.") 
			break;
		}
	}
}

func handleRequests() {
	r := mux.NewRouter()
    r.HandleFunc("/{ip}", checkIP) // load data and get country for Ip
	http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

 func main() {
     handleRequests()   
 }