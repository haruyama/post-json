package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX_USER_ID = 100000
	MAX_PATH_ID = 50000
)

var userAgents = []string{"IE", "Moziila", "Chrome", "Opera", "Safari"}

var mobileCarriers = []string{"", "Docomo", "AU", "Softbank", "E-mobile"}

type AccessInfo struct {
	IpAddress     string `json:"ip_address"`
	User_id       int64  `json:"user_id"`
	Path          string `json:"path"`
	UserAgent     string `json:"user_agent"`
	MobileCarrier string `json:"mobile_carrier"`
	DateTime      string `json:"datetime"`
}

func userId() int64 {
	return int64(rand.Intn(MAX_USER_ID) + 1)
}

func ipAddress() string {
	num := rand.Int31()
	return fmt.Sprintf("%d.%d.%d.%d", num>>24, (num>>16)%256, (num>>8)%256, num%256)
}

func path() string {
	return fmt.Sprintf("/path%d", rand.Intn(MAX_PATH_ID))
}

func userAgent() string {
	return userAgents[rand.Intn(len(userAgents))]
}

func mobileCarrier() string {
	return mobileCarriers[rand.Intn(len(mobileCarriers))]
}

func dateTime() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func GetAccessInfo() AccessInfo {
	return AccessInfo{ipAddress(), userId(), path(), userAgent(), mobileCarrier(), dateTime()}
}

func GetAccessInfoJson(is_elasticsearch bool) []byte {
	if is_elasticsearch {
		json, _ := json.Marshal(GetAccessInfo())
		return json
	} else {
		json, _ := json.Marshal([]AccessInfo{GetAccessInfo()})
		return json
	}
}
