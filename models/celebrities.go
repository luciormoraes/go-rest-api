package models

type Celebrity struct {
	Name    string `json:"name"`
	History string `json:"history"`
}

var Celebrities []Celebrity
