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

func TestCreateArticle(t *testing.T) {
	type CreateArticleTest struct {
		ID          int
		Username    string
		Title       string
		TypeID      int    `gorm:"column:TypeID"`
		CommunityID int    `gorm:"column:CommunityID"`
		CreateDay   string `gorm:"column:CreateDay"`
		Content     string
	}
	ArticleInfo := CreateArticleTest{
		ID:          1,
		Username:    "test",
		Title:       "test",
		TypeID:      1,
		CommunityID: 1,
		CreateDay:   "2019-01-01",
		Content:     "test",
	}
	requestData, _ := json.Marshal(ArticleInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/article/create",
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

func TestDeleteArticleByID(t *testing.T) {
	type DeleteArticleByIDTest struct {
		ID int
	}
	ArticleInfo := DeleteArticleByIDTest{
		ID: 1,
	}
	requestData, _ := json.Marshal(ArticleInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/article/delete/:id",
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

func TestUpdateArticleTitleOrContentByID(t *testing.T) {
	type UpdateArticleTitleOrContentByIDTest struct {
		ID      int
		Title   string
		Content string
	}
	ArticleInfo := UpdateArticleTitleOrContentByIDTest{
		ID:      1,
		Title:   "test",
		Content: "test",
	}
	requestData, _ := json.Marshal(ArticleInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/article/update",
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

func TestGetOneArticleByID(t *testing.T) {
	type GetOneArticleByIDTest struct {
		ID int
	}
	ArticleInfo := GetOneArticleByIDTest{
		ID: 1,
	}
	requestData, _ := json.Marshal(ArticleInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/article/getone",
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

func TestGetArticlesBySearchWords(t *testing.T) {
	type GetArticlesBySearchWordsTest struct {
		SearchWords string
	}
	ArticleInfo := GetArticlesBySearchWordsTest{
		SearchWords: "test",
	}
	requestData, _ := json.Marshal(ArticleInfo)
	response, err1 := http.NewRequest("POST", "http://localhost:10010/gf/api/article/search",
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
