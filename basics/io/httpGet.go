package io

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//ExecuteHTTPGet - an example of http GET to a url
func ExecuteHTTPGet() {
	fmt.Println("Inside io.ExecuteHttpGet")
	defer fmt.Println("Completed io.ExecuteHttpGet")
	page := DoHTTPGet(`htt/www.google.com`)
	fmt.Println(page)
}

//DoHTTPGet - an example of using http package to perform http GET on a url and reading the payload
func DoHTTPGet(url string) string {
	fmt.Println("Inside io.DoHttpGet")
	defer fmt.Println("Completed io.DoHttpGet")
	res, e := http.Get(url)
	if e != nil {
		log.Fatalln(e)
	}
	defer res.Body.Close()

	c, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(c)
}
