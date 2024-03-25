package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var counter uint32 = 1

func NewPOSTReq() error {
	data := Task{Username: "Tom", Task: fmt.Sprintf("hw%d", counter), Is_done: false}
	dataInBytes, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/tasks", bytes.NewBuffer(dataInBytes))
	if err != nil {
		fmt.Println(err)
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	fmt.Println(res.Status)
	fmt.Println(res.Header)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(string(body))
	counter++
	return nil

}
