package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"news/handles"
	"news/httprouter"
	"news/response"
	"strconv"
	"sync"
)

var servers sync.WaitGroup

func main() {
	router := httprouter.New()
	router.GET("/", getNews)
	router.GET("/detail", getDetail)
	router.GET("/allids", getAllIds)
	router.GET("/pageids", getPageIds)
	servers.Add(1)
	go func() {
		defer servers.Done()
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}()

	httpsRouter := httprouter.New()
	httpsRouter.GET("/login", login)
	httpsRouter.POST("/registbydefault", registerByDefault)
	servers.Add(1)
	go func() {
		defer servers.Done()
		err := http.ListenAndServeTLS(":8082", "httpskey/server.crt", "httpskey/private.key", httpsRouter)
		if err != nil {
			log.Fatal("ListenAndServeTLS:", err)
		}
	}()
	servers.Wait()
}

func registerByDefault(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	postBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(postBody))
	strPostBody := string(postBody)
	var userInfo handles.UserInfo
	json.Unmarshal([]byte(strPostBody), &userInfo)
	fmt.Println(userInfo)
	if handles.Register(&userInfo) {
		return
	} else {
		resp := &response.ResponseMessage{Message: "Insert Error", Detail: "{}"}
		bytes, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(bytes))
	}
	fmt.Fprintf(w, "registerByDefault")
}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "login")
}

func getNews(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "News Project")
}

func getPageIds(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	page := r.Form["page"][0]
	count := r.Form["count"][0]
	SuccRespon := handles.GetIdByPageAndCount(page, count)
	bytes, _ := json.Marshal(SuccRespon)
	fmt.Fprintf(w, string(bytes))
}

func getAllIds(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	SuccRespon := handles.GetAllIds()
	bytes, _ := json.Marshal(SuccRespon)
	fmt.Fprintf(w, string(bytes))
}

func getDetail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	id := r.Form["id"][0]
	if resp, ok := checkId(id); !ok {
		bytes, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(bytes))
		return
	}
	SuccRespon := handles.GetDetailById(id)
	bytes, _ := json.Marshal(SuccRespon)
	fmt.Fprintf(w, string(bytes))
}

func checkId(id string) (resp *response.ResponseMessage, ok bool) {
	ok = false
	// 输入为空
	if id == "" {
		resp = &response.ResponseMessage{Message: "Please Input Id.", Detail: "{}"}
		return
	}
	intId, AtoiOk := strconv.Atoi(id)
	// 转换失败（非数字、大于int等）或小于等于0
	if AtoiOk != nil || intId <= 0 {
		resp = &response.ResponseMessage{Message: "Please Input A Valid Id.", Detail: "{}"}
		return
	}
	ok = true
	return
}
