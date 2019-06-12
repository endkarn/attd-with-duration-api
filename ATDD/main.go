package main

import (
	"log"
	"net/http"
	"fmt"
)

func main (){
	http.HandleFunc("/duration", func (w http.ResponseWriter, r *http.Request)  {
		w.Header().Set("Content-Type","application/json")
		fmt.Fprintf(w, `{"days":"7,907 days",
		"startFullDate":"Thursday, 16 October 1997",
		"endFullDate":"Monday, 10 June 2019",
		"dayOfMonths":"259 months, 25 days",
		"dayOfWeeks":"1,129 weeks, 4 days",
		"seconds":"683,164,800 seconds",
		"minutes":"11,386,080 minutes",
		"hours":"189,768 hours"}`)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}