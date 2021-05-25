package domain

import "time"

type JobExecution struct {
	ID        int       `json:"id"`
	ObjectID  int       `json:"object_id"`
	Sleep     int       `json:"sleep"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RandomType int

const (
	Min = 15
	Max = 40
)

type StatusType int

const (
	Processing StatusType = iota
	Success
	Failed
	Cancelled
)

var statusValues = [...]string{
	"PROCESSING",
	"SUCCESS",
	"FAILED",
	"CANCELLED",
}

func (s StatusType) String() string {
	return statusValues[s]
}
