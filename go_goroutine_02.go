package main

import(
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"time"
)
func responseSize(url string){
	fmt.Println(url)
	resp, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(len(body))
}
func main(){
	go responseSize("https://www.inhatc.ac.kr")
	go responseSize("https://www.nate.com")
	go responseSize("https://www.google.com")
	go responseSize("https://www.harvard.edu")
	time.Sleep(time.Second * 5)
}
/*
package main

import(
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)

func main(){
	resp, e := http.Get("https://www.inhatc.ac.kr")
	if e != nil {
		log.Fatal(e)
	}
	defer resp.Body.Close() //html의 body부분을 main함수 끝나는 시점에 종료, 메모리 누수 방지
	body, e := ioutil.ReadAll(resp.Body) //resp의 주소의 body 내용 받아옴
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(string(body))
}
*/
