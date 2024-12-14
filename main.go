package main

import "jsp/cmd"

type Stock struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func main() {
	cmd.Execute()
}
