package main

import (
	"encoding/json"
	"fmt"     
	"os"
    "io"
	"io/ioutil"	
    "log"
    "net/http"
//	"strings"
//    "bytes"
)

type Ress struct {
	Res Res `json:"result"`
}	

type Res struct {
	Matchs []Match `json:"addressMatches"` 
}

type Match struct {
	Ads string `json:"matchedAddress"`
	Coord Coord `json:"coordinates"`	
} 

type Coord struct {
	Lon float64 `json:"x"`
	Lat float64 `json:"y"`
}



func main() {
     // Create output file
     newFile, err := os.Create("response.json")
     if err != nil {
          log.Fatal(err)
     }
     defer newFile.Close()

     adrs := "4600+Silver+Hill+Rd%2C+Suitland%2C+MD+20746"

//     adrs := os.Args[1]

    url := "https://geocoding.geo.census.gov/geocoder/locations/onelineaddress?benchmark=9&format=json&address=" + adrs
	 
    response, err := http.Get(url)
    defer response.Body.Close()

    // Write bytes from HTTP response to file.
    // response.Body satisfies the reader interface.
    // newFile satisfies the writer interface.
    // That allows us to use io.Copy which accepts
    // any type that implements reader and writer interface
	 
    numBytesWritten, err := io.Copy(newFile, response.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Downloaded %d byte file.\n", numBytesWritten)
	 
	// Open our jsonFile
	jsonFile, err := os.Open("response.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	
	fmt.Println("Successfully Opened response.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened file as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	
	// we initialize our Ress 
	var result Ress
	
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'result' which we defined above
	json.Unmarshal(byteValue, &result)
	
	fmt.Println("Address: " + result.Res.Matchs[0].Ads)		
	fmt.Printf("lon: %f \n",result.Res.Matchs[0].Coord.Lon)
	fmt.Printf("lat: %f \n",result.Res.Matchs[0].Coord.Lat)	
}