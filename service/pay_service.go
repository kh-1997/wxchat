package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"wxcloudrun-golang/db/model"
)

func PayHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method == http.MethodGet {

	} else if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		body2 := make(map[string]interface{})
		if err := decoder.Decode(&body2); err != nil {
			return
		}
		defer r.Body.Close()
		out_trade_no, _ := body2["out_trade_no"].(string)
		total_fee, _ := body2["total_fee"].(string)
		fmt.Println(total_fee)
		fmt.Println(total_fee)
		fee, err := strconv.ParseInt(total_fee,10,64)
		fmt.Println(fee)
		fmt.Println(fee)
		if err != nil{
			fmt.Println(err)
		}
		pay_body, _ := body2["pay_body"].(string)
		log.Print(r.Header)
		open_id := r.Header["X-Wx-Openid"][0]
		maps := make(map[string]interface{})
		maps["openid"] = open_id
		maps["out_trade_no"] = out_trade_no
		maps["spbill_create_ip"] = getIPV4()
		maps["env_id"] = "prod-2gej9ar9791db14a"
		maps["sub_mch_id"] = "1635092677"
		maps["total_fee"] = fee
		maps["body"] = pay_body
		maps["callback_type"] = 2
		container := make(map[string]interface{})
		container["service"] = "pay"
		container["path"] = "/api"
		maps["container"] = container
		reqParam, err := json.Marshal(maps)
		reqBody := strings.NewReader(string(reqParam))
		log.Print(reqBody)
		resp, err := http.Post("http://api.weixin.qq.com/_/pay/unifiedorder", "application/json", reqBody)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		log.Print(resp)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		payInfo := model.PayInfo{}
		err = json.Unmarshal(body, &payInfo)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = payInfo
		}
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