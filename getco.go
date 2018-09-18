package main

import (
     "os"
     "io"
     "log"
     "net/http"
)

func main() {
     // Create output file
     newFile, err := os.Create("test.json")
     if err != nil {
          log.Fatal(err)
     }
     defer newFile.Close()

     // HTTP GET request devdungeon.com 
     url := "https://geocoding.geo.census.gov/geocoder/locations/onelineaddress?address=4600+Silver+Hil l+Rd%2C+Suitland%2C+MD+20746&benchmark=9&format=json"
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
}