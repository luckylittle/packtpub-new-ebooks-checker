package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	s "strings"
	"time"
)

func main() {
	// Make HTTP GET request to Packt Publishing website
	response, err := http.Get("https://www.packtpub.com/all-books")
	if err != nil {
		log.Fatal(err)
	}

	// ReadAll function provided by ioutil package
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	re := regexp.MustCompile(`.*data-product-title=.*`)
	matches := re.FindAllString(responseString, -1)
	res1 := s.Join(matches, ",")

	reg := regexp.MustCompile(`\"(.*?)\"`)
	clean := reg.FindAllString(res1, -1)
	res2 := s.Join(clean, ",")

	d := time.Now()
	fmt.Print("New Packt Publishing E-books on ", d, "\n", s.Replace(res2, ",", "\n", -1))

	// Cleanup operation
	defer response.Body.Close()
}
