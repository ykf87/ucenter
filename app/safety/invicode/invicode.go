package invicode

import (
	"container/list"
	"errors"
	"fmt"
	"strings"
)

const (
	BASE    = "E8S2DZX9WYLTN6BQF7CP5IK3MJUR4HV"
	DECIMAL = 31
	PAD     = "A"
	LEN     = 8
)

// id转code

func Encode(uid uint64) string {
	id := uid
	mod := uint64(0)
	res := ""
	for id != 0 {
		mod = id % DECIMAL
		id = id / DECIMAL
		res += string(BASE[mod])
	}
	resLen := len(res)
	if resLen < LEN {
		res += PAD
		for i := 0; i < LEN-resLen-1; i++ {
			res += string(BASE[(int(uid)+i)%DECIMAL])
		}
	}
	return res
}

func Decode(code string) uint64 {
	res := uint64(0)
	lenCode := len(code)
	baseArr := []byte(BASE) // 字符串进制转换为byte数组

	baseRev := make(map[byte]int) // 进制数据键值转换为map
	for k, v := range baseArr {
		baseRev[v] = k
	}
	// 查找补位字符的位置
	isPad := strings.Index(code, PAD)
	if isPad != -1 {
		lenCode = isPad
	}
	r := 0
	for i := 0; i < lenCode; i++ {
		// 补充字符直接跳过
		if string(code[i]) == PAD {
			continue
		}
		index, ok := baseRev[code[i]]
		if !ok {
			return 0
		}
		b := uint64(1)
		for j := 0; j < r; j++ {
			b *= DECIMAL
		}
		res += uint64(index) * b
		r++
	}
	return res
}

var baseStr string = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ"
var base []byte = []byte(baseStr)
var baseMap map[byte]int

func init() {
	baseMap = make(map[byte]int)
	for i, v := range base {
		baseMap[v] = i
	}
}
func Base34(n uint64) []byte {
	quotient := n
	mod := uint64(0)
	l := list.New()
	for quotient != 0 {
		//fmt.Println("---quotient:", quotient)
		mod = quotient % 34
		quotient = quotient / 34
		l.PushFront(base[int(mod)])
		//res = append(res, base[int(mod)])
		//fmt.Printf("---mod:%d, base:%s\n", mod, string(base[int(mod)]))
	}
	listLen := l.Len()

	if listLen >= 6 {
		res := make([]byte, 0, listLen)
		for i := l.Front(); i != nil; i = i.Next() {

			res = append(res, i.Value.(byte))
		}
		return res
	} else {
		res := make([]byte, 0, 6)
		for i := 0; i < 6; i++ {
			if i < 6-listLen {
				res = append(res, base[0])
			} else {
				res = append(res, l.Front().Value.(byte))
				l.Remove(l.Front())
			}

		}
		return res
	}

}

func Base34ToNum(str []byte) (uint64, error) {
	if baseMap == nil {
		return 0, errors.New("no init base map")
	}
	if str == nil || len(str) == 0 {
		return 0, errors.New("parameter is nil or empty")
	}
	var res uint64 = 0
	var r uint64 = 0
	for i := len(str) - 1; i >= 0; i-- {
		v, ok := baseMap[str[i]]
		if !ok {
			fmt.Printf("")
			return 0, errors.New("character is not base")
		}
		var b uint64 = 1
		for j := uint64(0); j < r; j++ {
			b *= 34
		}
		res += b * uint64(v)
		r++
	}
	return res, nil
}

/*func main(){
	InitBaseMap()
	fmt.Printf("len(baseStr):%d, len(base):%d\n", len(baseStr), len(base))
	res := Base34(24)
	fmt.Printf("===============base:24->%s, %d\n", string(res), len(res))
 	res = Base34(200441052)
	fmt.Printf("===============base:200441052->%s, %d\n", string(res), len(res))
	res = Base34(1544804416)
	fmt.Printf("===============base:1544804416->%s, %d\n", string(res), len(res))
	str := "4DZRX2"
	num, err := Base34ToNum([]byte(str))
	if err == nil {
		fmt.Printf("===============base:%s->%d\n", str, num)
	}else {
		fmt.Printf("===============err:%s\n", err.Error())
	}
	str = "1000000"
	num, err = Base34ToNum([]byte(str))
	if err == nil {
		fmt.Printf("===============base:%s->%d\n", str, num)
	} else {
		fmt.Printf("===============err:%s\n", err.Error())
	}

	str = "XGKJTG"
	num, err = Base34ToNum([]byte(str))
	if err == nil {
		fmt.Printf("===============base:%s->%d\n", str, num)
	}else {
		fmt.Printf("===============err:%s\n", err.Error())
	}
}*/
