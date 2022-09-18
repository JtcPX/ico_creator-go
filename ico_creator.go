package main

import (
	"fmt"
	"strconv"
	"os"
	"flag"
	"io/ioutil"
)

var(
	f string
	n int
	desFileName string
)

func main() {
	Usage()
	flag.Parse()
	rw_file(f,desFileName,n)
}

func Usage() {
	flag.StringVar(&f,"f","bak.ico", "Input file directory")
	flag.IntVar(&n,"n", 5, "Please enter a numbers")
}

func rw_file(srcFileName string,desFileName string,n int) {
	for i:=1;i<n+1;i++  {
		desFileName = strconv.Itoa(i) + ".ico"
		str,err := ioutil.ReadFile(f)
		if err != nil{
			fmt.Printf("讀取文件錯誤,錯誤為:%v\n",err)
			return
		}
		ioutil.WriteFile(desFileName,[]byte(str),0664)
		f, err := os.OpenFile(desFileName, os.O_APPEND|os.O_WRONLY, 0600)
		f.WriteString(strconv.Itoa(i))
	}
}