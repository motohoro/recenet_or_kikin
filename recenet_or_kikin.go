package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"regexp"
	"bufio"
	"strconv"
)

func main() {
	fmt.Println("Please wait")
	recenetmap := make(map[string]string)
	
	//http://golang.org/pkg/net/http/
	targeturl := "http://www.mdcom.jp/recenet_kenpo.php"
	resp, err := http.Get(targeturl)
	if err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		ioutil.WriteFile("tmpfile", body, os.ModePerm)
	}
	//http://golang.jp/pkg/html
	reval,err := ioutil.ReadFile("tmpfile")
	if err != nil {
		panic(err)
	}
	r1, _ := regexp.Compile("\\s")
	s1 := r1.ReplaceAllString(string(reval),"")

	r2, _ := regexp.Compile("<th>(.+?)</th><td>(.+?)</td>")
//	fmt.Println( r2.FindAllStringSubmatch(s1,100))
	reg1 := r2.FindAllStringSubmatch(s1,100)
	for _,reg := range reg1{
//		fmt.Println(reg[1])
		recenetmap[reg[2]] = reg[1]
	}
	
//	fmt.Println(recenetmap)
	fmt.Println(targeturl)
	fmt.Println(strconv.Itoa(len(recenetmap))+"件")
	fmt.Println("チェックする保険番号を入力、終了時はCtrl+C")
	//http://qiita.com/hnakamur/items/a53b701c8827fe4bfec7
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
//		fmt.Println(scanner.Text())
		//http://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
		if val,ok := recenetmap[scanner.Text()]; ok{
			fmt.Println(val)
		}else{
			fmt.Println("該当なし")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
