package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func getBody(URL string) []byte {
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return html
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
func main() {
	data := getBody("http://tructiep24h.com/")
	dataString := string(data)
	re1 := regexp.MustCompile(`(?m)http://static.tructiep24h.com/.+(png|jpg|jpeg)`)
	re2 := regexp.MustCompile(`(?m)http://media.tructiep24h.com/.+(png|jpg|jpeg)`)
	listImage := re1.FindAllString(dataString, -1)
	for _, v := range listImage {
		vList := strings.Split(v, "/")
		err := DownloadFile(vList[len(vList)-1], v)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded: " + vList[len(vList)-1])
	}
	listImage = re2.FindAllString(dataString, -1)
	for _, v := range listImage {
		vList := strings.Split(v, "/")
		err := DownloadFile(vList[len(vList)-1], v)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded: " + vList[len(vList)-1])
	}
}
