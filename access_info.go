package main

import (
    "fmt"
    "math/rand"
)

const (
    MAX_USER_ID = 100000
)

type AccessInfo struct {
    Ip_address string `json:"ip_address"`
    User_id    int64  `json:"user_id"`
}

func userId() int64 {
    return int64(rand.Intn(MAX_USER_ID) + 1)
}

func ipAddress() string {
    num := rand.Int31()
    return fmt.Sprintf("%d.%d.%d.%d", num>>24, (num>>16)%256, (num>>8)%256, num%256)
}

func GetAccessInfo() AccessInfo {
    return AccessInfo{ipAddress(), userId()}
}
