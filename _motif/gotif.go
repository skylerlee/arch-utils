package main

import "fmt"

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	conf := LoadConf("conf/gist.json")
	fmt.Println(conf)
}
