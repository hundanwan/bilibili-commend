package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func StartCrawler(url string)  string{
	client := http.Client{}
	req ,err := http.NewRequest("GET",url,nil)
	if err != nil{
		fmt.Println("Newrequest err:",err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp,err := client.Do(req)
	if err != nil {
		fmt.Println("Client do err :",err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Http status code: ",resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)

	}
	return string(body)

}
