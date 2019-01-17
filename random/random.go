package random

import (
	"fmt"
	"math/rand"
	"time"
	"sync/atomic"
	"crypto/md5"
	"encoding/hex"

)

func Random6dec() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

var headList = []string{"https://sdun.oss-cn-shenzhen.aliyuncs.com/aa6b04d79c699fe229464dd3cd86ce88.png"}

func GetRandHead() string {
	headListLen := len(headList)
	n := rand.Intn(headListLen)
	if n >= headListLen {
		return "https://sdun.oss-cn-shenzhen.aliyuncs.com/aa6b04d79c699fe229464dd3cd86ce88.png"
	} else {
		return headList[n]
	}
}

var randHeadList = []string{"https://sdun.oss-cn-shenzhen.aliyuncs.com/aa6b04d79c699fe229464dd3cd86ce88.png",
	"https://sdun.oss-cn-shenzhen.aliyuncs.com/dbed8a73d3912c8fae53df635f98706c.png",
	"https://sdun.oss-cn-shenzhen.aliyuncs.com/6ab58203a1dc916432de00af83c1daca.png",
	"https://sdun.oss-cn-shenzhen.aliyuncs.com/2490df8d46315a2aaa3e6ef37a60e166.png",
}

func SetRegisterRandHeader() string {
	headListLen := len(randHeadList)
	n := rand.Intn(headListLen)
	if n >= headListLen {
		return "https://sdun.oss-cn-shenzhen.aliyuncs.com/aa6b04d79c699fe229464dd3cd86ce88.png"
	} else {
		return randHeadList[n]
	}
}

var Ran int64

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min) + min
	return randNum
}

func randomMoney(remainCount int64, remainMoney int64,minMoney int64,maxMoney int64) int64 {
	if remainCount == 1 {
		return remainMoney
	}

	rand.Seed(time.Now().UnixNano()+Ran)
	atomic.AddInt64(&Ran,1)
	//rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	var min int64
	min = minMoney

	max := remainMoney / remainCount * 2
	if max>maxMoney-minMoney {
		max = maxMoney-minMoney
	}

	money := rand.Int63n(max) + min
	return money
}

func RedPackage(count, money int64,minMoney int64,maxMoney int64) []int64 {
	a := make([]int64, 0)
	var i int64
	for i = 0; i < count; i++ {
		m := randomMoney(count-i, money,minMoney, maxMoney )
		a = append(a, m)
		money -= m
	}
	return a
}


func RandomRangeArr(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}


func RandomRangeArr64(min, max int64, count int) []int64 {
	s := make([]int64, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {

		rvaule := r.Int63n(max-min) + min
		for j := 0; j < len(s); j++ {
			if s[j] == rvaule {
				continue
			}
		}
		s = append(s, rvaule)
	}
	return s
}


func GenUserToken(uid int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(1000000) + time.Now().Nanosecond()
	srand := fmt.Sprintf("%d", n)
	suid := fmt.Sprintf("%d", uid)
	rr := suid + srand[2:]
	return Md5(rr)
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}



