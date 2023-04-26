package main

import (
	"crypto/tls"
	"fmt"
	"heathChecker/config"
	"net/http"
	"os"

	"github.com/jperozoa/golog"
)

func main() {
	cfg := config.New()

	golog.GetInstance().SetLogLevel(golog.LevelInfo)

	if len(os.Args) < 2 {
		golog.Critical("Use %v URL\n", os.Args[0])
	}
	url := os.Args[len(os.Args)-1]

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: cfg.Insecure}

	resp, err := http.Get(url)
	if err != nil {
		golog.Critical(err.Error())
	}

	fmt.Println(resp.Status)
}
