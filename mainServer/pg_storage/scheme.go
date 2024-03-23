package pg_storage

import "time"

type Task struct {
	Id       int       `json:"Id"`
	Username string    `json:"Username"`
	Task     string    `json:"Task"`
	Deadline time.Time `json:"Deadline"`
	Is_done  bool      `json:"Is_Done"`
}
