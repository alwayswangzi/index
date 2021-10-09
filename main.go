package main

import "github.com/alwayswangzi/sir"

func main() {
	r := sir.New()
	r.Static("/sdk/")
	r.Template("/", "index.html")
	r.ListenAndServe(":8000")
}
