package main

import (
	"ToDoList/ToDoListServer/TestClient/requests"
	"time"
)

func main() {
	timer := time.NewTicker(time.Second * 5)

	for range timer.C {
		requests.NewPOSTReq()
	}

	time.Sleep(time.Minute * 5)
}
