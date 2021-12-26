package routes

import (
	"Calendar/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func GetToken() string {
	var jsonStr = []byte(`{"Email":"postman@gmail.com","Password":"secret"}`)
	req, err := http.NewRequest("POST", "http://localhost:8000/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	var tok LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&tok)
	fmt.Println("token:", tok.Token)
	return tok.Token
}

func TestEventGetAll(t *testing.T) {
	token := GetToken()
	req, err := http.NewRequest("GET", "http://localhost:8000/events", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func TestGetOne(t *testing.T) {
	token := GetToken()
	var jsonStr = []byte(`{"id":"1"}`)
	req, err := http.NewRequest("GET", "http://localhost:8000/event", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var e entity.Event
	err = json.NewDecoder(resp.Body).Decode(&e)
	if e.Id != "1" {
		t.Errorf("want %d, got %s", 1, e.Id)
	}

}
