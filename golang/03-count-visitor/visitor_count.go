package main
count_
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
        "time"
)

var requestcounter int


func handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
        dt := time.Now()
	if r.URL.RawQuery != "" {
		url = strings.Join([]string{url, "?", r.URL.RawQuery}, "")
	}
	fmt.Fprintf(w, "Hello There !\n")
	fmt.Fprintf(w, "Your request no. at time: %s%s%s", dt.String(), ", is :", strconv.Itoa(requestcounter))
        
	log.Println(url)
	requestcounter++
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

