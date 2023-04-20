package utils

import (
	"fmt"
	"time"
)

type SystemResponse struct {
	Message string
	Time    time.Time
}

func (r *SystemResponse) String() string {
	return fmt.Sprintf("%s | time: %v", r.Message, r.Time)
}
