package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"net/http"
)

func main(){
	resp, err := http.Get("https://goo.gl/mSzFnk")

	if err != nil{
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	contents , err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	var r map[string]interface{}

	json.Unmarshal(contents, &r)
	observe := r["current_observation"].(map[string]interface{})

	fmt.Printf("temp_c:%.1f", observe["temp_c"])
}