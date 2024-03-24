package TestClient

import "ToDoList/ToDoListServer/TestClient/host"

func main() {
	channel := make(chan bool)
	host.Init()
	// This is a blocking call
	<-channel
}
