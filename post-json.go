package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"math/rand"
	"time"
)

const (
	MaxUserId = 10000
	NUMBER_OF_COROUTINE = 100
	NUMBER_OF_LOOP      = 10
)

type AccessInfo struct {
	Ip_address string `json:"ip_address"`
	User_id    int64  `json:"user_id"`
}

func userId() int64 {
	return int64(rand.Intn(MaxUserId) + 1)
}

func ipAddress() string{
	num := rand.Int31()
	return fmt.Sprintf("%d.%d.%d.%d", num >> 24, (num >> 16) % 256, (num >> 8) % 256, num % 256)
}

func post(c chan int) {
	info := []AccessInfo{{ipAddress(), userId()}}
	j, _ := json.Marshal(info)
	resp, err := http.Post("http://localhost:8983/solr/update/json", "application/json", bytes.NewReader(j))
	if err != nil {
		fmt.Println(string(j))
		fmt.Println(resp)
	}
	c <- 1
}


func main() {
	rand.Seed( time.Now().UTC().UnixNano())
	for j := 0; j < NUMBER_OF_LOOP ; j++ {
		c := make(chan int, NUMBER_OF_COROUTINE)
		for i := 0; i < NUMBER_OF_COROUTINE; i++ {
			go post(c)
		}
		for i := 0; i < NUMBER_OF_COROUTINE; i++ {
			<- c
		}
	}
	fmt.Println("all done")
}
