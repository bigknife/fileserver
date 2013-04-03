package config

import (
    "bufio"
    "errors"
    "fmt"
    "io"
    "os"
    "strings"
)

//LoadProperties 读取properties配置文件，从中解析出key-value 的map
func LoadProperties(propFile string) (propMap map[string]string, err error) {
    propMap = make(map[string]string)
    err = nil

    f, err := os.Open(propFile)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()

    r := bufio.NewReaderSize(f, 4*1024)
    if err != nil {
        fmt.Println(err)
        return
    }

    line, isPrefix, err := r.ReadLine()
    for err == nil && !isPrefix && err != io.EOF {
        s := string(line)
        //忽略 # 开头的行
        if !strings.HasPrefix(s, "#") {
            eqIdx := strings.Index(s, "=")
            if eqIdx >= 0 {
                k, v := strings.Trim(s[:eqIdx], " "), strings.Trim(s[eqIdx+1:], " ")
                propMap[k] = v
            }
        }

        line, isPrefix, err = r.ReadLine()
    }

    if isPrefix {
        err = errors.New("buffer size to small")
        return
    }

    return
}
