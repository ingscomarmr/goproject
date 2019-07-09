package model

//structura auxiliar para mandar mensajes entre interfaces
type Message struct {
	Message string `json:"message"`
	Code    uint   `json:"code"`
}
