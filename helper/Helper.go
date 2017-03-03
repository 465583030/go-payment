package helper

import (
	"sort"
	"fmt"
	"bytes"
	"log"
	"crypto/md5"
	"time"
	"strconv"
)

func TimeMd5() string {
	t := time.Now()
	timestamp := t.Unix()
	log.Println(timestamp)
	stamp := strconv.Itoa(int(timestamp))
	bytes := md5.Sum([]byte(stamp))
	md5String := fmt.Sprintf("%x", bytes)
	fmt.Println(md5String)
	return md5String
}

func MapToXMLString(data map[string]string) string {
	var buf bytes.Buffer
	buf.WriteString("<xml>")
	for key, value := range data {
		buf.WriteString("<")
		buf.WriteString(key)
		buf.WriteString("><![CDATA[")
		buf.WriteString(value)
		buf.WriteString("]]></")
		buf.WriteString(key)
		buf.WriteString(">")
	}
	buf.WriteString("</xml>")
	return buf.String()
}

func ToURLParamsSortByKey(params map[string]string) string {
	keys := []string{}
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	endMark := keys[len(keys)-1]
	fmt.Println(endMark)
	var buf bytes.Buffer
	equal := "="
	and := "&"
	for _, key := range keys {
		buf.WriteString(key)
		buf.WriteString(equal)
		buf.WriteString(params[key])
		if key != endMark {
			buf.WriteString(and)
		}
	}
	return buf.String()
}

func MakeSign(params map[string]string) string {
	presignString := ToURLParamsSortByKey(params)
	log.Print(presignString)
	bytes := md5.Sum([]byte(presignString))
	return fmt.Sprintf("%x", bytes)
}
