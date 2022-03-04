package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"unsafe"
)

func TestGetCommunityByName(t *testing.T) {
	type CommunityInfo struct {
		Name string `json:"Name"`
	}

	communityInfo := CommunityInfo{
		Name: "group8",
	}

	requestData, _ := json.Marshal(communityInfo)
	response, err1 := http.NewRequest("GET", "http://localhost:10010/gf/api/community/getcommunity",
		strings.NewReader(string(requestData)))
	if err1 != nil {
		t.Error("Failed to Request. " + err1.Error())
	}
	defer response.Body.Close()

	content, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		t.Error("Failed to Read Response Body. " + err2.Error())
		return
	}

	str := (*string)(unsafe.Pointer(&content))
	if strings.Contains(*str, "400") {
		t.Error("Failed to Get Community By Name. " + *str)
		return
	}
	fmt.Println(*str)
}
