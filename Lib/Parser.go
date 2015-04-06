package Lib

import (
	"fmt"
	"github.com/HostsSitter/Enum"
	"github.com/HostsSitter/Models"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func ParseUrl() bool {
	url := Models.Source.UpdateUrl
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

	part := regexp.MustCompile(Models.Source.RawPreStr + `([\w\W]*)` + Models.Source.RawPostStr)
	match := part.FindAllStringSubmatch(rawBody, 1)
	var newRawHosts string
	for _, v := range match {
		newRawHosts = v[0]
	}

	reg := regexp.MustCompile(Models.Source.RawPreStr)
	newRawHosts = reg.ReplaceAllString(newRawHosts, Enum.HOST_PRE_STR)

	reg = regexp.MustCompile(Models.Source.RawPostStr)
	newHosts := reg.ReplaceAllString(newRawHosts, Enum.HOST_POST_STR)

	fmt.Println(newHosts)

	return true
}

func ParseHosts() bool {
	return true
}
