package main

import (
	"fmt"
	windyapi "windyapi/windyapiservice"
)

func main() {
	resp, err := windyapi.GetWeather(53.1900, -112.2500)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
