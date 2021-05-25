package domain

type StatusType int
type RandomType int

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

const (
	Min = 15
	Max = 40
)
