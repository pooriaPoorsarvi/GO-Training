package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func normal() string{
	res, err := http.Get("http://comp.sbu.ac.ir/amc")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(page)
}

func notUsingReturns() string {
	res, _ := http.Get("http://comp.sbu.ac.ir/amc")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return string(page)
}

func main(){
	fmt.Println("if calling this function causes an error, we will recieve a log \n", normal())
	fmt.Println("if calling this function causes an error, we won't recieve a log \n", notUsingReturns())

}

