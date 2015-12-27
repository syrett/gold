package lib

/*-------------------------------------------------------------------
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <2015-12-22 00:12:59>
* @doc
* utils.go
* @end
* Created : 2015/12/21 15:28:49 liyouyou

-------------------------------------------------------------------*/

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func GrandStr(l int) string {
	h := md5.New()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Fprintf(h, "%d.%d.%f", r.Intn(100), r.Int63(), r.Float64())
	str := hex.EncodeToString(h.Sum(nil))

	n := len(str)
	if n < l {
		str = str + GrandStr(l-n)
	} else {
		str = str[0:l]
	}

	return str
}

func Be_int(i interface{}) int64 {
	switch i.(type) {
	case int, int8, int16, int32, int64:
		switch i.(type) {
		case int, int8, int16, int32, int64:
			return reflect.ValueOf(i).Int()
		}
	case uint, uint8, uint16, uint32, uint64:
		switch i.(type) {
		case uint, uint8, uint16, uint32, uint64:
			return reflect.ValueOf(i).Int()
		}
	case string:
		d, _ := strconv.ParseInt(i.(string), 10, 0)
		return d
	case float64, float32:
		return int64(reflect.ValueOf(i).Float())
	case bool:
		println("bool")
	}
	return 0
}

func Be_string(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(i).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatInt(reflect.ValueOf(i).Int(), 10)
	case float64, float32:
		return strconv.FormatFloat(reflect.ValueOf(i).Float(), 'f', 2, 64)
	}
	return ""
}
