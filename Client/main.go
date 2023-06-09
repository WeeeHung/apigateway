package main

// Reference methods below
// =========================================================================
// func Do() {
// 	c, err := client.NewClient()
// 	if err != nil {
// 		return
// 	}
// 	req := &protocol.Request{}
// 	res := &protocol.Response{}
// 	req.SetMethod(consts.MethodGet)
// 	req.Header.SetContentTypeBytes([]byte("application/json"))
// 	req.SetRequestURI("http://www.example.com")
// 	err = c.Do(context.Background(), req, res)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Printf("%v", string(res.Body()))
// }

// func DoDeadline() {
// 	c, err := client.NewClient()
// 	if err != nil {
// 		return
// 	}
// 	req := &protocol.Request{}
// 	res := &protocol.Response{}
// 	req.SetMethod(consts.MethodGet)
// 	req.SetRequestURI("http://www.example.com")
// 	c.DoDeadline(context.Background(), req, res, time.Now().Add(1*time.Second))
// 	fmt.Printf("%v\n", string(res.Body()))
// }

// func DoRedirects() {
// 	c, err := client.NewClient()
// 	if err != nil {
// 		return
// 	}
// 	req := &protocol.Request{}
// 	res := &protocol.Response{}
// 	req.SetMethod(consts.MethodGet)
// 	req.Header.SetContentTypeBytes([]byte("application/json"))
// 	req.SetRequestURI("http://www.example.com")
// 	err = c.DoRedirects(context.Background(), req, res, 1)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Printf("%v\n", string(res.Body()))
// }

// func DoTimeout() {
// 	c, err := client.NewClient()
// 	if err != nil {
// 		return
// 	}
// 	req := &protocol.Request{}
// 	res := &protocol.Response{}
// 	req.SetMethod(consts.MethodGet)
// 	req.Header.SetContentTypeBytes([]byte("application/json"))
// 	req.SetRequestURI("http://www.example.com")
// 	err = c.DoTimeout(context.Background(), req, res, 1*time.Second)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Printf("%v\n", string(res.Body()))
// }

// func Get() {
// 	c, err := client.NewClient()
// 	if err != nil {
// 		return
// 	}
// 	status, body, _ := c.Get(context.Background(), nil, "http://www.example.com")
// 	fmt.Printf("status=%v body=%v\n", status, string(body))
// }

// func GetDeadline() {
// 	c, err := client.NewClient()
// 	if err != nil {
// 		return
// 	}
// 	status, body, _ := c.GetDeadline(context.Background(), nil, "http://www.example.com", time.Now().Add(1*time.Second))
// 	fmt.Printf("status=%v body=%v\n", status, string(body))
// }

// func GetTimeout() {
// 	c, err := client.NewClient()
// 	if err != nil {
// 		return
// 	}
// 	status, body, _ := c.GetTimeout(context.Background(), nil, "http://www.example.com", 1*time.Second)
// 	fmt.Printf("status=%v body=%v\n", status, string(body))
// }

// func Post() {
// 	c, err := client.NewClient()
// 	if err != nil {
// 		return
// 	}

// 	var postArgs protocol.Args
// 	postArgs.Set("arg", "a") // Set post args
// 	status, body, _ := c.Post(context.Background(), nil, "http://www.example.com", &postArgs)
// 	fmt.Printf("status=%v body=%v\n", status, string(body))
// }

import (
	// "bufio"
	// "os"
	// "strings"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/network/standard"
	// "github.com/cloudwego/hertz/pkg/protocol"
	// "github.com/cloudwego/hertz/pkg/common/utils"
	// "github.com/cloudwego/hertz/pkg/protocol"
	// "github.com/cloudwego/hertz/pkg/protocol/consts"
)

func getNumber() {
	_, body, err := client.Get(context.Background(), nil, "http://localhost:8080/numbers")
	if err != nil {
		// Handle error
		fmt.Printf("err: %v\n", err)
		return
	}
	var obj struct {
		Numbers struct {
			Numbers []int `json:"numbers"`
		} `json:"numbers"`
	}

	err = json.Unmarshal(body, &obj)
	if err != nil {
		// Handle error
		fmt.Printf("err: %s\n", body)
		return
	}
	fmt.Printf("Current numbers: %v\n", obj.Numbers.Numbers)
}

func addNumber(x int) {
	// Create a map with the "number" key and the integer value
	data := map[string]int{"number": x}

	// Convert the map to a JSON-encoded byte slice
	payload, err := json.Marshal(data)
	if err != nil {
		// Handle error
		return
	}

	// Create a new HTTP POST request with the JSON payload and Content-Type header
	req, err := http.NewRequest("POST", "http://localhost:8080/numbers", bytes.NewBuffer(payload))
	if err != nil {
		// Handle error
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request using an HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Handle error
		return
	}

	// Close the response body
	defer resp.Body.Close()

	fmt.Printf("Number added.\n")
}

func main() {
	// Create a new Hertz client.
	_, err := client.NewClient(
		client.WithDialTimeout(1*time.Second),
		client.WithDialer(standard.NewDialer()),
		client.WithKeepAlive(true),
	)
	if err != nil {
		return
	}

	getNumber()
	addNumber(23)
	getNumber()
	addNumber(453)
	getNumber()
	addNumber(5482)
	addNumber(932)
	getNumber()

	// defer resp.Body.Close()

	// // Decode the response from JSON.
	// var body1 utils.H
	// err = resp.DecodeJSON(&body1)
	// if err != nil {
	// 	// Handle error
	// }

	// // Print the response message.
	// fmt.Println(body1["message"])
}
