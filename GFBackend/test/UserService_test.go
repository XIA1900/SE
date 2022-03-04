package test

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

func TestUserLogin(t *testing.T) {
	loginInfo, err := userLogin("boss", "007")
	if err != nil {
		t.Error("Fail to Login. Error Message: " + err.Error())
		return
	}
	if loginInfo.Code == 200 {
		fmt.Println("Login Successfully")
		fmt.Println(loginInfo)
	} else {
		t.Errorf("Fail to Login. Error Message: %v", loginInfo)
		return
	}
}

func TestUserFollowers(t *testing.T) {
	loginInfo, err := userLogin("lion", "007")
	if err != nil || loginInfo.Message == "" {
		t.Error("Fail to Login. Error Message: " + err.Error())
		return
	}
	cookie := &http.Cookie{
		Name:  "token",
		Value: loginInfo.Message,
	}
	request, err1 := http.NewRequest("POST", "http://"+IP+":10010/gf/api/user/followers", nil)
	if err1 != nil {
		t.Error("Failed to Generate Request: " + err1.Error())
		return
	}
	request.AddCookie(cookie)
	jar, err2 := cookiejar.New(nil)
	if err2 != nil {
		t.Error("Failed to Set Cookie: " + err2.Error())
		return
	}
	var client http.Client
	client = http.Client{
		Jar: jar,
	}
	response, err3 := client.Do(request)
	if err3 != nil {
		t.Error("Failed to Request: " + err3.Error())
		return
	}
	defer response.Body.Close()

	err4 := printResponseContent(response)
	if err4 != nil {
		t.Error("Failed to Interpret Response Message: " + err4.Error())
		return
	}
}

func TestUserFollowees(t *testing.T) {
	loginInfo, err := userLogin("dog", "007")
	if err != nil || loginInfo.Message == "" {
		t.Error("Fail to Login. Error Message: " + err.Error())
		return
	}
	cookie := &http.Cookie{
		Name:  "token",
		Value: loginInfo.Message,
	}
	request, err1 := http.NewRequest("POST", "http://"+IP+":10010/gf/api/user/followees", nil)
	if err1 != nil {
		t.Error("Failed to Generate Request: " + err1.Error())
		return
	}
	request.AddCookie(cookie)
	jar, err2 := cookiejar.New(nil)
	if err2 != nil {
		t.Error("Failed to Set Cookie: " + err2.Error())
		return
	}
	var client http.Client
	client = http.Client{
		Jar: jar,
	}
	response, err3 := client.Do(request)
	if err3 != nil {
		t.Error("Failed to Request: " + err3.Error())
		return
	}
	defer response.Body.Close()

	err4 := printResponseContent(response)
	if err4 != nil {
		t.Error("Failed to Interpret Response Message: " + err4.Error())
		return
	}
}
