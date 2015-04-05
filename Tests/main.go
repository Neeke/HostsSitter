package main

import (
	"io/ioutil"
	"log"
	"net/http"
	//"os"
	"regexp"
	//"strconv"
	//"strings"
	//"sync"
	"fmt"
	"runtime"
)

const UPDATE_URL = "http://www.360kb.com/kb/2_122.html"
const HOST_PATH_LINUX = "/etc/hosts"
const HOST_PATH_WINDOWS = "C:\\Windows\\system32\\drivers\\etc\\hosts"

const RAW_PRE_STR = "==更新分界线，复制下面内容到hosts文件即可====="
const RAW_POST_STR = "</pre>"

const HOST_PRE_STR = "#Got new hosts from Neeke"
const HOST_POST_STR = "#Got new hosts from Neeke End"

func parseUrl(url string) bool {
	ret, err := http.Get(url)
	if err != nil {
		log.Println(url)
		status := map[string]string{}
		status["status"] = "400"
		status["url"] = url
		panic(status)
	}
	body := ret.Body
	data, _ := ioutil.ReadAll(body)

	rawBody := string(data)

	part := regexp.MustCompile(RAW_PRE_STR + `([\w\W]*)` + RAW_POST_STR)
	match := part.FindAllStringSubmatch(rawBody, 1)
	var newRawHosts string
	for _, v := range match {
		newRawHosts = v[0]
	}

	reg := regexp.MustCompile(RAW_PRE_STR)
	newRawHosts = reg.ReplaceAllString(newRawHosts, HOST_PRE_STR)

	reg = regexp.MustCompile(RAW_POST_STR)
	newHosts := reg.ReplaceAllString(newRawHosts, HOST_POST_STR)

	fmt.Println(newHosts)

	return true
}

func parseHosts() bool {
	return true
}

func main() {
	parseUrl(UPDATE_URL)

	fmt.Println(runtime.GOOS)
}
