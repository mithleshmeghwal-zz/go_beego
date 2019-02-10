package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"ormTest/models"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func TestAddUser(t *testing.T) {
	clearTable()

	body := map[string]string{
		"username": "username12",
		"name":     "name12",
		"email":    "mail@mail.ckd",
	}
	d, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/addUser", bytes.NewBuffer(d))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateUser(t *testing.T) {
	clearTable()
	addUser(1)

	body := map[string]string{
		"username":  "userN",
		"confirmed": "1",
	}
	d, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/updateUser", bytes.NewBuffer(d))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func addUser(num int) {
	if num <= 0 {
		num = 1
	}
	o := orm.NewOrm()
	o.Insert(&models.User{
		Username: "userN",
		Name:     "namE",
		Email:    "mail@mail.coM",
	})
}

func clearTable() (int64, error) {
	o := orm.NewOrm()
	var u []models.User
	num, err := o.Raw("delete from user;").QueryRows(&u)
	return num, err
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(rr, req)
	// fmt.Println("REQUEST")
	// fmt.Println(req)
	fmt.Println()
	fmt.Println("RESPONSE")
	fmt.Println(rr.Body.String())
	fmt.Println()

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
