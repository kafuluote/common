package check

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"regexp"
	"sort"
)

func CheckPhone(phone string) bool {
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	if ok := rgx.MatchString(phone); !ok {
		return false
	}
	return true
}

var (
	token = "ssss"
)

func CheckEmail(email string) bool {
	reg := `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	rgx := regexp.MustCompile(reg)
	if ok := rgx.MatchString(email); !ok {
		return false
	}
	return true
}

func MakeSignature(mReq map[string]interface{}, key string) (sign string) {

	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	//nonce_str=coVmEFxqJzAfcAHt&num=1&on_price=1&opt=1&symbol=ETH/USDT&token=34f8c8aeaf4fc4c96da4407a7738786366623c23b40061111f29514284e97978&type=2&uid=100000&key=pfdsapowmsapa
	//nonce_str=coVmEFxqJzAfcAHt&num=1&on_price=1&opt=1&symbol=ETH/USDT&token=34f8c8aeaf4fc4c96da4407a7738786366623c23b40061111f29514284e97978&type=2&uid=100000&key=pfdsapowmsapa
	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}
	fmt.Printf("finial str %s", signStrings)
	//STEP4, 进行MD5签名并且将所有字符转为大写.
	s := sha1.New()
	s.Write([]byte(signStrings))
	cipherStr := s.Sum(nil)
	//upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return hex.EncodeToString(cipherStr)
}
