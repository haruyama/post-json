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

func write_httperf_wsesslog(is_elasticsearch bool) {
	for j := 0; j < NUMBER_OF_SESSION; j++ {
		for i := 0; i < NUMBER_OF_SESSION_ENTRIES; i++ {
			if is_elasticsearch {
				info := GetAccessInfo()
				json, _ := json.Marshal(info)
				fmt.Printf("/access_info/access_info method=POST contents='%s'\n", strings.Replace(string(json), "'", "\\'", -1))
			} else {
				info := []AccessInfo{GetAccessInfo()}
				json, _ := json.Marshal(info)
				fmt.Printf("/solr/update/json method=POST contents='%s'\n", strings.Replace(string(json), "'", "\\'", -1))
			}
		}
		fmt.Println()
	}
}
