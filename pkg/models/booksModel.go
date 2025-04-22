package models

import "dayalubajpai/pkg/config"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func init() {
	config.ConnectDB()
}
