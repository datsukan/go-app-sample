package controllers

type H struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewH(message string, data interface{}) *H {
	H := new(H)
	H.Message = message
	H.Data = data
	return H
}
