package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)

func main(){
	f, e := ioutil.ReadFile("NSYSU.json")
	var r map[string]interface{}

	if e != nil {
		fmt.Println("read file error:%v", e)
		os.Exit(1)
	}

	json.Unmarshal(f, &r)

	observe := r["current_observation"].(map[string]interface{})

	fmt.Println("temp_c : ", observe["temp_c"])
	fmt.Println("icon_url : ", observe["icon_url"])
}