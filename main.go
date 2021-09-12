package main

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"gopkg.in/gomail.v2"
	"path"
	"strconv"
)


func getEnv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

var (
	username   = getEnv("EMAIL_USERNAME", "")
	password   = getEnv("EMAIL_PASSWORD", "")
	host	   = getEnv("EMAIL_HOST", "smtp.gmail.com")
	port	   = getEnv("EMAIL_PORT", "465")
	TMP_DIR    = getEnv("TMP_DIR", "/tmp/")
)


func exitGracefully(err string) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}


func main() {

	if len(os.Args) < 3 {
		exitGracefully("Usage: s2k <url> <to>")
	}

	url := os.Args[1]
	to := os.Args[2]

	if  (url == "" || to == "") {
		exitGracefully("Both <url> and <to> are required")
	}

	fmt.Println(url)
	fmt.Println(to)

	filename :=  path.Base(url)

	err := DownloadFile(TMP_DIR + filename, url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + url)

	fmt.Println("Sending...")

	Send(to, filename)

}

func Send(to string, filename string) error {

	if filename == "" {
		return fmt.Errorf("No file to attach")
	}
	extension := path.Ext(filename)
	m := gomail.NewMessage()

	if extension == "pdf" {
		m.SetHeader("Subject", "convert")
	} else {
		m.SetHeader("Subject", "")
	}

	m.SetHeader("From", username + "@gmail.com")
	m.SetHeader("To", to)
	m.SetBody("text/html", "Hello from <b>S2K</b>")
	m.Attach(TMP_DIR+filename)

	
	_port, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	d := gomail.NewDialer(host, int(_port) , username, password)

	error := d.DialAndSend(m); 
		return error
		
}

func DownloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}