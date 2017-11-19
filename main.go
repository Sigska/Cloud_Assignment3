
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


func HandleLatest(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		var info map[string]interface{}

		err:= json.NewDecoder(r.Body).Decode(&info)

		if err!= nil {
			Fprint(w, "error decoding")
		}
	}








	else {
		Fprint(w, "only post implemented")
	}
}















func main () {

http.HandleFunc("/latest", HandleLatest)
http.ListenAndServe(":" +os.Getenv("PORT"), nil)

}




