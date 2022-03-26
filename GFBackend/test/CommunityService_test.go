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

func TestUpdateCommunity(t *testing.T) {
	type CommunityInfo struct {
		ID          int    `json:"ID"`
		Name        string `json:"Name"`
		Description string `json:"Description"`
	}

	communityInfo := CommunityInfo{
		ID:          11,
		Name:        "group11",
		Description: "test11",
	}

	requestData, _ := json.Marshal(communityInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/community/updatecommunitybyid",
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
		t.Error("Failed to Update Community. " + *str)
		return
	}
	fmt.Println(*str)
}

func TestDeleteCommunity(t *testing.T) {
	type CommunityInfo struct {
		ID int `json:"ID"`
	}

	communityInfo := CommunityInfo{
		ID: 11,
	}

	requestData, _ := json.Marshal(communityInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/community/deletecommunitybyid",
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
		t.Error("Failed to Delete Community. " + *str)
		return
	}
	fmt.Println(*str)
}

func TestCreateCommunity(t *testing.T) {
	type CommunityInfo struct {
		Name        string `json:"Name"`
		Description string `json:"Description"`
	}

	communityInfo := CommunityInfo{
		Name:        "group11",
		Description: "test11",
	}

	requestData, _ := json.Marshal(communityInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/community/create",
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
		t.Error("Failed to Create Community. " + *str)
		return
	}
	fmt.Println(*str)
}

func TestDeleteCommunityByID(t *testing.T) {
	type CommunityInfo struct {
		ID int `json:"ID"`
	}

	communityInfo := CommunityInfo{
		ID: 11,
	}

	requestData, _ := json.Marshal(communityInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/community/delete/:id",
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
		t.Error("Failed to Delete Community By ID. " + *str)
		return
	}
	fmt.Println(*str)
}

func TestUpdateDescriptionByID(t *testing.T) {
	type CommunityInfo struct {
		ID          int    `json:"ID"`
		Description string `json:"Description"`
	}

	communityInfo := CommunityInfo{
		ID:          11,
		Description: "test11",
	}

	requestData, _ := json.Marshal(communityInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/community/update",
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
		t.Error("Failed to Update Description By ID. " + *str)
		return
	}
	fmt.Println(*str)
}

func TestGetNumberOfMemberByID(t *testing.T) {
	type CommunityInfo struct {
		ID int `json:"ID"`
	}

	communityInfo := CommunityInfo{
		ID: 11,
	}

	requestData, _ := json.Marshal(communityInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/community/numberofmember/:id",
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
		t.Error("Failed to Get Number Of Member By ID. " + *str)
		return
	}
	fmt.Println(*str)
}

func TestGetOneCommunityByID(t *testing.T) {
	type CommunityInfo struct {
		ID int `json:"ID"`
	}

	communityInfo := CommunityInfo{
		ID: 11,
	}

	requestData, _ := json.Marshal(communityInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/community/getone/:id",
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
		t.Error("Failed to Get One Community By ID. " + *str)
		return
	}
	fmt.Println(*str)
}

func TestGetCommunitiesByNameFuzzyMatch(t *testing.T) {
	type CommunityNameFuzzyMatch struct {
		Name string `json:"Name"`
	}

	communityInfo := CommunityNameFuzzyMatch{
		Name: "test",
	}

	requestData, _ := json.Marshal(communityInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/community/getbyname",
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
		t.Error("Failed to Get Communities By Name Fuzzy Match. " + *str)
		return
	}
	fmt.Println(*str)
}

func TestJoinCommunityByID(t *testing.T) {
	type JoinCommunity struct {
		ID int `json:"ID"`
	}

	communityInfo := JoinCommunity{
		ID: 11,
	}

	requestData, _ := json.Marshal(communityInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/community/join/:id",
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
		t.Error("Failed to Join Community By ID. " + *str)
		return
	}
	fmt.Println(*str)
}

func TestLeaveCommunityByID(t *testing.T) {
	type LeaveCommunity struct {
		ID int `json:"ID"`
	}

	communityInfo := LeaveCommunity{
		ID: 11,
	}

	requestData, _ := json.Marshal(communityInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/community/leave/:id",
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
		t.Error("Failed to Leave Community By ID. " + *str)
		return
	}
	fmt.Println(*str)
}
