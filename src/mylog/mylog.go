/*
日志框架
利用fmt来格式化输出
当前时间: time.Now()
日期格式化：Format
日期格式化规则：http://golang.org/src/pkg/time/format.go
const (
    60		stdLongMonth      = "January"
    61		stdMonth          = "Jan"
    62		stdNumMonth       = "1"
    63		stdZeroMonth      = "01"
    64		stdLongWeekDay    = "Monday"
    65		stdWeekDay        = "Mon"
    66		stdDay            = "2"
    67		stdUnderDay       = "_2"
    68		stdZeroDay        = "02"
    69		stdHour           = "15"
    70		stdHour12         = "3"
    71		stdZeroHour12     = "03"
    72		stdMinute         = "4"
    73		stdZeroMinute     = "04"
    74		stdSecond         = "5"
    75		stdZeroSecond     = "05"
    76		stdLongYear       = "2006"
    77		stdYear           = "06"
    78		stdPM             = "PM"
    79		stdpm             = "pm"
    80		stdTZ             = "MST"
    81		stdISO8601TZ      = "Z0700"  // prints Z for UTC
    82		stdISO8601ColonTZ = "Z07:00" // prints Z for UTC
    83		stdNumTZ          = "-0700"  // always numeric
    84		stdNumShortTZ     = "-07"    // always numeric
    85		stdNumColonTZ     = "-07:00" // always numeric
    86	)
*/
package mylog

import (
    "fmt"
    "time"
)

const (
    fatalFormat string = "FATAL\t [%s] %v\n"
    infoFormat  string = "INFO\t [%s] %v\n"
)

func Fatal(sFormat string, v ...interface{}) {
    fmt.Printf(fatalFormat, time.Now().Format("2006/01/02 15:04:05.000000000"), fmt.Sprintf(sFormat, v))
}

func Info(sFormat string, v ...interface{}) {
    fmt.Printf(infoFormat, time.Now().Format("2006/01/02 15:04:05.000000000"), fmt.Sprintf(sFormat, v))
}
