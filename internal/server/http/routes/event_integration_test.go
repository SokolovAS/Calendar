package routes

import (
	"Calendar/entity"
	"Calendar/internal/repository"
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
	t := tok.Token
	fmt.Println("token", t)
	return t
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
	fmt.Println(err)
	if e.Id != "1" {
		t.Errorf("want %d, got %s", 1, e.Id)
	}
}

func TestAdd(t *testing.T) {
	token := GetToken()
	var jsonStr = []byte(`{
   "id": "2",
   "title": "Title2",
   "description": "Description2",
   "dateTime": "DateTiem2",
   "duration": "Duration2",
   "notes": "Notes2"
}`)
	db := repository.NewRepoPG()
	db.Delete("2")

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/event", bytes.NewBuffer(jsonStr))
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
	if e.Id != "2" {
		t.Errorf("want %d, got %s", 1, e.Id)
	}
}

func TestUpdate(t *testing.T) {
	token := GetToken()
	var jsonStr = []byte(`{"id": "3","title": "Title Updated"}`)

	db := repository.NewRepoPG()
	testE := entity.Event{Id: "3", Title: "test title 3"}
	db.Add(testE)

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8000/event", bytes.NewBuffer(jsonStr))
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
	if e.Title != "Title Updated" {
		t.Errorf("want %s, got %s", e.Title, "Title Updated")
	}

	defer db.Delete("3")
}

func TestDelete(t *testing.T) {
	var jsonStr = []byte(`{"id": "100"}`)

	db := repository.NewRepoPG()
	testE := entity.Event{Id: "100", Title: "test title 100"}
	db.Add(testE)

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8000/event", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", GetToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		t.Errorf("want 200, got %s", resp.Status)
	}
}
