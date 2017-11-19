
package main 

import (
		)

var URL = string("http://api.fixer.io/latest")

type Rates struct {
		Base string 	`json: "base"`
		Rates map[string]float64	`json: "rates"`
		Date string 	`json: "date"`
}

type Response struct {
	Text string `json:"text"`

}


















func main () {



}




