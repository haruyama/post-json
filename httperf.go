package main

import (
	"fmt"
	"strings"
)

const (
	NUMBER_OF_SESSION_ENTRIES = 1000
	NUMBER_OF_SESSION         = 10
	PATH_OF_SOLR              = "/solr/update/json"
	PATH_OF_ELASTICSEARCH     = "/access_info/access_info"
)

func getPath(is_elasticsearch bool) string {
	if is_elasticsearch {
		return PATH_OF_ELASTICSEARCH
	}
	return PATH_OF_SOLR
}

func write_httperf_wsesslog(is_elasticsearch bool) {
	for j := 0; j < NUMBER_OF_SESSION; j++ {
		for i := 0; i < NUMBER_OF_SESSION_ENTRIES; i++ {
			json := GetAccessInfoJson(is_elasticsearch)
			fmt.Printf("%s method=POST contents='%s'\n", getPath(is_elasticsearch), strings.Replace(string(json), "'", "\\'", -1))
		}
		fmt.Println()
	}
}
