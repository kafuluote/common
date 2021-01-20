package convert

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strings"

	"github.com/shopspring/decimal"
)

func ByteToInt32(b []byte) (x uint32) {
	b_buf := bytes.NewBuffer(b)
	b_buf = bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, x)
	return
}

func ByteToInt64(b []byte) (x int64) {
	b_buf := bytes.NewBuffer(b)
	b_buf = bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, x)
	return
}

//两个int64 相加
func Int64AddInt64(a, b int64) (int64, error) {
	aa := decimal.New(a, 0)
	bb := decimal.New(b, 0)
	return aa.Add(bb).IntPart(), nil
}

func Int64ToInt64By8Bit(b int64) int64 {
	a := decimal.New(b, 0)
	r := a.Mul(decimal.New(100000000, 0))
	return r.IntPart()
}

func Int64ToFloat64By8Bit(b int64) (x float64) {
	a := decimal.New(b, -8)
	x, _ = a.Float64()
	return
}

func Int64ToStringBy8Bit(b int64) string {
	a := decimal.New(b, 0)
	r := a.Div(decimal.New(100000000, 0))
	return r.String()
}

func Int64ToString(b int64) string {
	a := decimal.New(b, 0)
	return a.String()
}

func StringToStringBy8Bit(b string) string {
	a, err := decimal.NewFromString(b)
	if err != nil {
		return ""
	}
	r := a.Div(decimal.New(100000000, 0))
	return r.String()
}

func StringAddString(a, b string) string {
	c, _ := decimal.NewFromString(a)
	d, _ := decimal.NewFromString(b)
	return c.Add(d).String()
}

//0.00001001
func StringToInt64By8Bit(s string) (int64, error) {
	d, err := decimal.NewFromString(s)
	l := d.Round(8).Coefficient().Int64()
	return l, err
}

func Float64ToInt64By8Bit(s float64) int64 {
	d := decimal.NewFromFloat(s)
	l := d.Round(8).Coefficient().Int64()
	return l
}

//请确保返回值不会越界否则调用下面的返回string类型
func Int64MulInt64By8Bit(a int64, b int64) int64 {
	dd := decimal.New(a, 0)
	dp := decimal.New(b, 0)
	m := dd.Mul(dp)
	d := decimal.New(100000000, 0)
	n := m.Div(d)
	return n.IntPart()
}

func Int64MulInt64MulInt64By16Bit(a int64, b, c int64) int64 {
	da := decimal.New(a, 0)
	db := decimal.New(b, 0)
	dc := decimal.New(c, 0)
	m := da.Mul(db).Mul(dc)
	d := decimal.New(100000000, 0)
	n := m.Div(d).Div(d)
	return n.IntPart()
}

func Int64MulInt64By8BitString(a int64, b int64) string {
	dd := decimal.New(a, 0)
	dp := decimal.New(b, 0)
	m := dd.Mul(dp)
	d := decimal.New(100000000, 0)
	n := m.Div(d)

	r := n.Div(decimal.New(100000000, 0))
	return r.String()
}

func Int64MulStringBy8BitString(a string, b int64) string {
	dd := decimal.New(b, 0)
	dp, _ := decimal.NewFromString(a)
	m := dd.Mul(dp)
	d := decimal.New(100000000, 0)
	n := m.Div(d)

	r := n.Div(decimal.New(100000000, 0))
	return r.String()
}

func Int64MulStringBy8BitString2Bit(a string, b int64) string {
	dd := decimal.New(b, 0)
	dp, _ := decimal.NewFromString(a)
	d := decimal.New(100000000, 0)
	return dd.Mul(dp).Div(d).Div(d).Round(8).String()
}
func Int64MulInt64By8BitString2Bit(a int64, b int64) string {
	dd := decimal.New(a, 0)
	dp := decimal.New(b, 0)
	d := decimal.New(100000000, 0)
	t := dd.Mul(dp).Div(d).Div(d)
	return t.Round(2).Coefficient().String()
}

func Int64MulFloat64(a int64, b float64) int64 {
	dd := decimal.New(a, 0)
	dp := decimal.NewFromFloat(b)
	m := dd.Mul(dp)
	return m.IntPart()
}

func StringToInt64(a string) int64 {
	dd, _ := decimal.NewFromString(a)
	return dd.IntPart()
}

func Int64DivString(a int64, b string) string {
	dd := decimal.New(a, 0)
	dp, err := decimal.NewFromString(b)
	if err != nil {
		return ""
	}
	return dd.Div(dp).Round(2).String()
}

//两数相除保持8位
func Int64DivInt64By8Bit(a int64, b int64) int64 {
	dd := decimal.New(a, 0)
	dp := decimal.New(b, 0)
	d := decimal.New(100000000, 0)

	num := dd.Div(dp).Mul(d).IntPart()
	return num
}

//两数相除保持8位
func Int64DivInt64By8BitString(a int64, b int64) string {
	dd := decimal.New(a, 0)
	dp := decimal.New(b, 0)
	d := decimal.New(100000000, 0)
	num := dd.Div(dp).Mul(d).String()
	//num := dd.Div(dp).Round(8).Coefficient().String()
	return num
}

//两数相除保持2位
func Int64DivInt64StringPercent(a int64, b int64) string {
	dd := decimal.New(a, 0)
	dp := decimal.New(b, 0)
	d := decimal.New(100, 0)

	t := dd.Div(dp).Mul(d)
	k, _ := t.Float64()
	s := fmt.Sprintf("%.2f", k)

	return s
}

//两数相加保持3位
func Int64AddInt64Float64Percent(a int64, b int64) string {
	dd := decimal.New(a, 0)
	dp := decimal.New(b, 0)
	d := decimal.New(100000000, 0)

	t := dd.Add(dp).Div(d)
	k, _ := t.Float64()
	s := fmt.Sprintf("%.3f", k)

	return s
}

func Int64MulStringInt64By8Bit(a int64, b string) string {
	dd := decimal.New(a, 0)
	dp, err := decimal.NewFromString(b)
	if err != nil {
		return ""
	}

	d := decimal.New(100000000, 0)
	t := dd.Mul(dp).Div(d)
	return t.String()
}

func Int64MulInt64DivInt64By8Bit(a, b, c int64) string {
	da := decimal.New(a, 0)
	db := decimal.New(b, 0)
	dc := decimal.New(c, 0)
	dd := decimal.New(100000000, 0)

	return da.Mul(db).Div(dc).Div(dd).Round(2).String()
}

func Int64SubInt64(a int64, b int64) string {
	dd := decimal.New(a, 0)
	dp := decimal.New(b, 0)
	d := decimal.New(100000000, 0)

	t := dd.Sub(dp).Div(d).String()

	return t
}

func Int64ToStringDiv8Bit(a int64) string {

	dp := decimal.New(a, 0)

	d := decimal.New(100000000, 0)
	t := dp.Div(d)
	return t.String()
}

var tenToAny map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}

// map根据value找key
func findkey(in string) int {
	result := -1
	for k, v := range tenToAny {
		if in == v {
			result = k
		}
	}
	return result
}

// 任意进制转10进制
func AnyToDecimal(num string, n int) int {
	var new_num float64
	new_num = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(findkey(value))
		if tmp != -1 {
			new_num = new_num + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(new_num)
}
