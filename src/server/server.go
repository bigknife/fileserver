/*
文件服务器，提供Http multi-part/form 形式的上传接口
*/
package main

import (
    "config"
    "encoding/base64"
    "errors"
    "flag"
    "fmt"
    "handler"
    "io"
    "mylog"
    "net/http"
    "strings"
)

const (
    defaultSetting string = "server.properties"
    sHelp          string = `指定server运行的配置文件路径`
)

//主函数，应用程序入口点
func main() {
    //从 -s 指定的配置文件中加载配置信息
    setting := flag.String("s", defaultSetting, sHelp)
    flag.Parse()
    fmt.Printf("使用配置文件 %v 启动Server...\n", *setting)

    //从配置文件中读取一个配置map
    //cfg map[string]string
    cfg, err := config.LoadProperties(*setting)
    if err != nil && err != io.EOF {
        fmt.Printf("加载配置文件出错，%v \n", err)
        return
    }

    //检查配置文件必备字段
    _, err = checkCfg(cfg)
    if err != nil {
        fmt.Printf("检查配置文件出错，%v \n", err)

    }

    //启动http服务

    http.HandleFunc("/", errorHandler(handler.HandleRoot, cfg))
    //http.HandleFunc("/", handler.HandleRoot)
    listen, _ := cfg["listen"]
    fmt.Println("server started and liten " + listen)
    handler.Init(cfg)
    http.ListenAndServe(listen, nil)

}

//检查运行参数是否完备
//需要的参数：
//		listen = ip:port
//		storePath = /opt/file/
func checkCfg(cfg map[string]string) (success bool, err error) {
    if cfg == nil {
        return false, errors.New("server配置项为空")
    }
    _, ok := cfg["listen"]

    if !ok {
        success, err = false, errors.New("缺少listen配置参数")
        return
    }
    _, ok = cfg["storePath"]
    if !ok {
        success, err = false, errors.New("缺少storePath配置参数")
        return
    }
    cfg["storePath"] = strings.TrimRight(cfg["storePath"], "/")
    return true, nil
}
func notAuth(w http.ResponseWriter, r *http.Request) {
    mylog.Info("%s", "NotAuth response")
    w.Header().Set("WWW-Authenticate", `Basic realm="iDongler User Login"`)
    w.WriteHeader(http.StatusUnauthorized)
}
func errorHandler(fn http.HandlerFunc, cfg map[string]string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        defer func() {
            if e := recover(); e != nil {
                mylog.Fatal("%v", e)
                //w.WriteHeader(500)
                //fmt.Fprintf(w, "%s", e)
                http.Error(w, fmt.Sprintf("%s", e), 500)
            }
        }()

        //if not basic auth 拒绝
        s, c, err := ParseRequest(r)
        fmt.Printf("s=%s,c=%s\n", s, c)
        if err != nil || s == "" || c == "" {
            notAuth(w, r)
            return
        }

        //rd := base64.NewDecoder(base64.StdEncoding,strings.NewReader(s))
        dc, err := base64.StdEncoding.DecodeString(c)
        if err != nil {
            panic(err)
        }

        np := strings.Split(string(dc), ":")
        name, pwd := np[0], np[1]

        authPwd, ok := cfg["user:"+name]
        if !ok || pwd != authPwd {
            notAuth(w, r)
            return

        }
        fn(w, r)
    }
}
func ParseRequest(r *http.Request) (s, c string, err error) {
    h, ok := r.Header["Authorization"]
    mylog.Info("%v", h)
    if !ok || len(h) == 0 {
        return "", "", errors.New("The authorization header is not set.")
    }
    s, c, err = Parse(h[0])
    return
}
func Parse(value string) (s, c string, err error) {
    parts := strings.Split(value, " ")
    if len(parts) == 2 {
        fmt.Printf("parts = %v, part[0] = %s, part[1]=%s\n", parts, parts[0], parts[1])
        s, c, err = parts[0], parts[1], nil
        return
    }
    return "", "", errors.New("The authorization header is malformed.")
}
