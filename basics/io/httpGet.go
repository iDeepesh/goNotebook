package io

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ExecuteHttpGet() {
	fmt.Println("Inside io.ExecuteHttpGet")
	defer fmt.Println("Completed io.ExecuteHttpGet")
	page := DoHttpGet(`htt/www.google.com`)
	fmt.Println(page)
}

func DoHttpGet(url string) string {
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
