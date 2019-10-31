package globals

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"strings"
	"sort"
	"crypto/md5"
	"encoding/hex"
)

func GetAccess(appid string, secret string) (string, error) {

	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + secret
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err2:", err)
		return "", err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	type AccessTokenRes struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}

	c := &AccessTokenRes{}
	err = json.Unmarshal(body, c)
	if err != nil {
		fmt.Println("err3:", err)
		return "", err
	}
	return c.AccessToken, nil

}

func SendWxMsg(url string, params interface{}) error {

	dataJson, _ := json.Marshal(params)
	resp, err1 := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(string(dataJson)))
	if err1 != nil {
		fmt.Println("err1:", err1)
		return err1
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("err2:", err2)
		return err2
	}
	fmt.Println("body:", string(body))
	return nil
}

func GetSign(m map[string]string, key string) string {
	fmt.Println("private_key:", key)

	sortedKeys := make([]string, 0)
	signStr := ""
	for k, _ := range m {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	for _, k := range sortedKeys {
		signStr += k + "=" + m[k] + "&"
	}

	signStr += "key" + "=" + key
	h := md5.New()
	fmt.Println("signStr:", signStr)
	h.Write([]byte(signStr))
	cipherStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}
