package main

import apihttp "github.com/bortiz-101/go-acis/api/internal/http"

/*
This func will be entry point for whole api. Will eventually:
create the ACIS client
connect to our db
load configs
create and run our gin router
*/
func main() {
	router := apihttp.CreateRouter()
	router.Run() // 8080 by default
}
