package handler

import (
	"net/http"
	"strconv"
	"strings"
	"errors"
)

// TODO("なんでModuleで分けてる?")

// URLから「id」を取得する
func getUserIdFromURL(r *http.Request) (int,error)  {
	url := r.URL.Path
	urlSlice := strings.Split(url,"/")
	if len(urlSlice) < 3 {
		err := errors.New("URL Length Invalid")
		return 0,err
	}
	id,err := strconv.Atoi(urlSlice[2])
	if err != nil {
		return 0,err
	}
	return id,nil
}
