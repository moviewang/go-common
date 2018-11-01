package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	resp, err := http.Get("https://krcom.cn/aj/discover/albumlist?ajwvr=6&page=20|20181101146&__rnd=1541053619821")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(string(body))
}