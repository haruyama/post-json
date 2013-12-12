package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	NUMBER_OF_COROUTINE = 10
	NUMBER_OF_LOOP      = 1000
	RATE_PER_SEC        = 1000
)

func GetAccessInfoJson(is_elasticsearch bool) []byte {
	if is_elasticsearch {
		json, _ := json.Marshal(GetAccessInfo())
		return json
	} else {
		json, _ := json.Marshal([]AccessInfo{GetAccessInfo()})
		return json
	}
}

func post(is_elasticsearch bool, c chan int) {
	throttle := time.Tick(1e9 / RATE_PER_SEC)
	for i := 0; i < NUMBER_OF_LOOP; i++ {
		<-throttle
		json := GetAccessInfoJson(is_elasticsearch)
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
	rand.Seed(time.Now().UTC().UnixNano())

	do_httperf := flag.Bool("httperf", false, "write httperf wsesslog")
	is_elasticsearch := flag.Bool("elasticsearch", false, "elasticsearch mode")
	flag.Parse()

	if *do_httperf {
		write_httperf_wsesslog(*is_elasticsearch)
		return
	}

	c := make(chan int, NUMBER_OF_COROUTINE)
	for i := 0; i < NUMBER_OF_COROUTINE; i++ {
		go post(*is_elasticsearch, c)
	}
	for i := 0; i < NUMBER_OF_COROUTINE; i++ {
		<-c
	}
	fmt.Println("all done")
}
