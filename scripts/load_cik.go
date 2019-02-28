package scripts

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//CIKURL is the location to dowload all EDGAR CIK associated company names
const CIKURL string = "https://www.sec.gov/Archives/edgar/cik-lookup-data.txt"

// LoadCIK pulls down and updates all CIK
func LoadCIK() {
	res, err := http.Get(CIKURL)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
