package controller

import (
	"Week02/logic"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	str := query.Get("id")
	id ,_ := strconv.Atoi(str)
	fmt.Println(id)
	user,err := logic.GetUserById(id)
	if err != nil {
		fmt.Printf("%v", errors.Cause(err))
		fmt.Printf("%+v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
}