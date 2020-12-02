package main

import (
	"Week02/controller"
	"Week02/dao"
	"fmt"
	"net/http"
)
func main() {
	if err := dao.InitDB();err != nil {
		fmt.Println(err)
	}
	defer dao.Close()
	http.HandleFunc("/user", controller.GetUserHandler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}

