package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

const (
    NUMBER_OF_SESSION_ENTRIES = 1000
    NUMBER_OF_SESSION         = 10
)

func write_httperf_wsesslog() {
    for j := 0; j < NUMBER_OF_SESSION; j++ {
        for i := 0; i < NUMBER_OF_SESSION_ENTRIES; i++ {
            info := []AccessInfo{GetAccessInfo()}
            json, _ := json.Marshal(info)
            fmt.Printf("/solr/update/json method=POST contents='%s'\n", strings.Replace(string(json), "'", "\\'", -1))
        }
        fmt.Println()
    }
}
