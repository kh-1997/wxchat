package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func PayHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method == http.MethodGet {

	} else if r.Method == http.MethodPost {
		log.Print(r.Header)
		open_id := r.Header["X-Wx-Openid"][0]
		maps := make(map[string]interface{})
		maps["openid"] = open_id
		maps["out_trade_no"] = "2021WERUN1647840687637"
		maps["spbill_create_ip"] = getIPV4()
		maps["env_id"] = "prod-2gej9ar9791db14a"
		maps["sub_mch_id"] = "1635092677"
		maps["total_fee"] = 1
		maps["callback_type"] = 2
		container := make(map[string]interface{})
		container["service"] = "pay"
		container["path"] = "/api"
		//maps["container"] = container
		reqParam, err := json.Marshal(maps)
		reqBody := strings.NewReader(string(reqParam))
		log.Print(reqBody)
		resp, err := http.Post("http://api.weixin.qq.com/_/pay/unifiedorder", "application/json", reqBody)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		log.Print(resp)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))

		msg, err := json.Marshal(res)
		if err != nil {
			fmt.Fprint(w, "内部错误")
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(msg)
	}
}

func getIPV4() string {
	resp, err := http.Get("https://ipv4.netarm.com")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}