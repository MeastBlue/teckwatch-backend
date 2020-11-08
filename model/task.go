package model

import (
	"fmt"
	"io"
	"strconv"
)

type Task struct {
	ID        string `json:"id"`
	Label     string `json:"label"`
	Prio      Prio   `json:"prio"`
	Progress  int    `json:"progress"`
	Done      bool   `json:"done"`
	User      *User  `json:"user"`
	Tech      *Tech  `json:"tech"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type NewTask struct {
	Label  string `json:"label"`
	UserID string `json:"userId"`
	TechID string `json:"techId"`
}

type Prio string

const (
	PrioNone   Prio = "NONE"
	PrioHigh   Prio = "HIGH"
	PrioMedium Prio = "MEDIUM"
	PrioLow    Prio = "LOW"
)

var AllPrio = []Prio{
	PrioNone,
	PrioHigh,
	PrioMedium,
	PrioLow,
}

func (e Prio) IsValid() bool {
	switch e {
	case PrioNone, PrioHigh, PrioMedium, PrioLow:
		return true
	}
	return false
}

func (e Prio) String() string {
	return string(e)
}

func (e *Prio) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Prio(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Prio", str)
	}
	return nil
}

func (e Prio) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
