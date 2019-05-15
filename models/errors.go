package models

import "fmt"

type PinnacleError struct {
	Ref     string `json:"ref,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (p *PinnacleError) Error() string {
	return fmt.Sprintf("[%s]: %s", p.Code, p.Message)
}
