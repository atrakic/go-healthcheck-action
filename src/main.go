package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func healthcheck(url string) (string, error) {
	if url == "" {
		return "", errors.New("Empty url")
	}

	var client = &http.Client{}
	res, err := client.Head(url)
	if err != nil {
		fmt.Printf("%s", err)
	}
	return strconv.Itoa(res.StatusCode), nil
}

func main() {
	r, e := healthcheck(os.Getenv("INPUT_URL"))
	if e != nil {
		fmt.Println("healthcheck failed:", e)
	} else {
		fmt.Println(fmt.Sprintf(`::set-output name=statusCode::%s`, r))
	}
}
