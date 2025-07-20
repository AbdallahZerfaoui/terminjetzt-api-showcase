package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.terminjetzt.com/appointments/latest")
	if err != nil { panic(err) }
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var v interface{}
	json.Unmarshal(body, &v)
	fmt.Printf("%+v\n", v)
}
