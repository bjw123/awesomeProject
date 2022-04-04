package controller

import (
	"awesomeProject/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func MockGetAll() string {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
	return sb
}

func MockGetSelected(id string) string {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
	return sb
}

func MockPost() int {
	var mockData []model.MockData
	body := append(mockData, model.MockData{UserId: 1, Id: 1, Title: "bryce", Completed: true})
	//body := model.NewData(100,100,"bryce", true)
	jsonData, _ := json.Marshal(body)
	fmt.Println(jsonData)
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("post request failed")
	}
	fmt.Println(resp)
	return resp.StatusCode
}
