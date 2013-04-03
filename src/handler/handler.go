package handler

import (
    "errors"
    "net/http"
)

type RestHandlerFunc func(w http.ResponseWriter, r *http.Request, cfg map[string]string)

type RestHandler struct {
    Url    string
    Method string
    Handle RestHandlerFunc
}

var handlerMap = make(map[string]RestHandler)
var serverConfig map[string]string

//初始化路由表
//		GET / 获取首页，返回200，则表示服务活着
//		POST /dir 新建目录，body为 name=目录名称 || name=dir1/dir2&recursion=true
//		GET /dir?name=父级目录&recursion=true 列出父级目录的子目录，如果recursion为true
func Init(cfg map[string]string) {
    serverConfig = cfg
    handlerMap["GET /"] = RestHandler{Url: "/", Method: "GET", Handle: IamAlive}
    handlerMap["POST /dir"] = RestHandler{Url: "/dir", Method: "POST", Handle: PostDir}
    handlerMap["GET /dir"] = RestHandler{Url: "/dir", Method: "GET", Handle: GetDir}
    handlerMap["GET /upload"] = RestHandler{Url: "/upload", Method: "GET", Handle: UploadPage}
    handlerMap["POST /file"] = RestHandler{Url: "/file", Method: "POST", Handle: PostFile}
    handlerMap["GET /file"] = RestHandler{Url: "/file", Method: "GET", Handle: GetFile}
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
    handler, success := route(r.URL.Path, r.Method)
    if success {
        handler.Handle(w, r, serverConfig)
    }
}

func route(path string, method string) (handler RestHandler, hasHandler bool) {
    key := method + " " + path
    rootHandler, ok := handlerMap[key]
    if ok {
        return rootHandler, true
    }
    panic(errors.New(key + " not found"))
}
