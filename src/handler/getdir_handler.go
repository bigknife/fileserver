/*
Get /dir?name=/a/b/c&recursion=true&mod=file|dir|all
读取目录下的信息，使用ioutil.ReadDir
难点是如何将 os.FileInfo 转换为json
*/
package handler

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "mylog"
    "net/http"
    "os"
    "strings"
)

type FileInfo struct {
    Name    string
    Size    int64
    ModTime string
    IsDir   bool
    Mode    os.FileMode
}

//查询目录
func GetDir(w http.ResponseWriter, r *http.Request, serverConfig map[string]string) {
    name, recursion, mode := r.FormValue("name"), r.FormValue("recursion"), r.FormValue("modeÓ")
    mylog.Info("post dir, {name = %s, recursion = %s}", name, recursion)

    if name == "" {
        panic(errors.New("GET /dir interface need 'name' param"))
    }

    if !(recursion == "true" || recursion == "false") {
        panic(errors.New("GET /dir interface need 'recursion' param, and the value must be 'true' or 'false'"))
    }

    storePath, _ := serverConfig["storePath"]

    absPath := storePath + "/" + strings.Trim(name, "/")
    absPath = strings.TrimRight(absPath, "/")

    mylog.Info("GET /dir list dir of absPath is %s", absPath)

    //resultList := make([]FileInfo, len(fileInfos))
    var resultList []FileInfo
    //os.FileInfo 转换为 FileInfo
    if recursion == "false" {
        resultList = list(absPath, storePath)
    } else {
        resultList = listR(absPath, storePath)
    }

    data, err := json.Marshal(resultList)
    if err != nil {
        panic(err)
    }

    w.Header()["Content-Type"] = []string{"application/json; charset=UTF-8"}
    w.WriteHeader(200)
    w.Write(data)

    fmt.Println(mode)
}

//os.FileInfo 转换为 FileInfo
//	要注意的是， slice 是可以增加的，申明是 []<type> ，不需要带{} 
func list(absPath string, prefixToDel string) []FileInfo {
    fileInfos, err := ioutil.ReadDir(absPath)
    if err != nil {
        panic(err)
    }

    resultList := []FileInfo{}
    for _, v := range fileInfos {
        //resultList[k] = FileInfo{v.Name(), v.Size(), v.ModTime().Format("2006/01/02 15:04:05.000000000 MST 2006"), v.IsDir(), v.Mode()}
        mylog.Info("%v", v)
        p := (absPath + "/" + strings.Trim(v.Name(), "/"))[len(prefixToDel):]

        resultList = append(resultList, FileInfo{p, v.Size(), v.ModTime().Format("2006/01/02 15:04:05.000"), v.IsDir(), v.Mode()})

    }
    return resultList
}

//os.FileInfo 转换为 FileInfo,并递归下级目录
func listR(absPath string, prefixToDel string) []FileInfo {
    fileInfos, err := ioutil.ReadDir(absPath)
    if err != nil {
        panic(err)
    }

    resultList := []FileInfo{}
    for _, v := range fileInfos {
        //resultList[k] = FileInfo{v.Name(), v.Size(), v.ModTime().Format("2006/01/02 15:04:05.000000000 MST 2006"), v.IsDir(), v.Mode()}
        mylog.Info("%v", v)
        p := (absPath + "/" + strings.Trim(v.Name(), "/"))[len(prefixToDel):]

        resultList = append(resultList, FileInfo{p, v.Size(), v.ModTime().Format("2006/01/02 15:04:05.000"), v.IsDir(), v.Mode()})

        if v.IsDir() {
            childPath := absPath + "/" + v.Name()
            childResultList := listR(childPath, prefixToDel)
            for _, cv := range childResultList {
                resultList = append(resultList, cv)
            }

        }
    }
    return resultList
}
