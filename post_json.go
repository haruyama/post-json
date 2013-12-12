package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"math/rand"
	"time"
	"flag"
)

const (
	NUMBER_OF_COROUTINE = 10
	NUMBER_OF_LOOP      = 1000
	RATE_PER_SEC        = 1000
)


func post(c chan int) {
	throttle := time.Tick(1e9 / RATE_PER_SEC)
	for i := 0; i < NUMBER_OF_LOOP ; i++ {
		<- throttle
		info := []AccessInfo{GetAccessInfo()}
		json, _ := json.Marshal(info)
		resp, err := http.Post("http://localhost:8983/solr/update/json", "application/json", bytes.NewReader(json))
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(json))
			fmt.Println(resp)
		}
	}
	c <- 1
}


func main() {
	rand.Seed( time.Now().UTC().UnixNano())

	do_httperf := flag.Bool("httperf", false, "write httperf wsesslog")
	flag.Parse()

	if *do_httperf {
		write_httperf_wsesslog()
		return
	}

	c := make(chan int, NUMBER_OF_COROUTINE)
	for i := 0; i < NUMBER_OF_COROUTINE; i++ {
		go post(c)
	}
	for i := 0; i < NUMBER_OF_COROUTINE; i++ {
		<- c
	}
	fmt.Println("all done")
}
