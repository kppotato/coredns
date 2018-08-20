package main

import (
	"regexp"
	"fmt"
	"strings"
)

func main()  {
	reg, _:=regexp.Compile(`/x\d{1,}$`)
	value := reg.ReplaceAllString("/skydns/com/wandafilm/lf2cd/x1","")
	list :=strings.Split(value,"/")
	myReverse(list)
	fmt.Println(strings.Join(list,"."))

}
func myReverse(l []string)  {
	for i:=0; i < int(len(l)/2) ;i++{
		li := len(l) - i -1
		l[i],l[li] = l[li],l[i]
	}
}