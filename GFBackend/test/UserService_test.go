package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"unsafe"
)

func TestUserLogin(t *testing.T) {
	type UserInfo struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
	}

	userInfo := UserInfo{
		Username: "dog",
		Password: "007",
	}

	requestData, _ := json.Marshal(userInfo)

	response, err1 := http.Post(
		"http://167.71.166.120:10010/gf/api/user/login",
		"application/json",
		bytes.NewBuffer(requestData))
	if err1 != nil {
		t.Error("Failed to Request. " + err1.Error())
		return
	}
	defer response.Body.Close()

	content, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		t.Error("Failed to Request. " + err2.Error())
		return
	}

	str := (*string)(unsafe.Pointer(&content))
	if strings.Contains(*str, "400") {
		t.Error("Failed to Request. " + *str)
		return
	}
	fmt.Println(*str)
}
