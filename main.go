package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func healthcheck(url string) (result string) {
	var client = &http.Client{}
	res, err := client.Head(url)
	if err != nil {
		fmt.Printf("%s", err)
	}
	return strconv.Itoa(res.StatusCode)
}

func main() {
	fmt.Println(fmt.Sprintf(`::set-output name=statusCode::%s`, healthcheck(os.Getenv("INPUT_URL"))))
}
