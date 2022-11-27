package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"wxcloudrun-golang/db/dao"
)

// CounterHandler 计数器接口
func GoodHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method == http.MethodGet {

	} else if r.Method == http.MethodPost {
		maps, err := getGoodAction(r)
		log.Printf("maps = %s",maps)
		if err != nil {
			return
		}
		action := maps["action"]
		if action == "get" {
			count, err := dao.ImpGood.GetGoodByName(maps["name"])
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = count
			}
		}
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

// getAction 获取action
func getGoodAction(r *http.Request) (map[string]string,error) {
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err := decoder.Decode(&body); err != nil {
		return nil, err
	}
	defer r.Body.Close()
	maps := make(map[string]string)
	action, ok := body["action"]
	if !ok {
		return nil, fmt.Errorf("缺少 action 参数")
	}
	maps["action"] = action.(string)
	name, ok := body["name"]
	if ok {
		maps["name"] = name.(string)
	}
	return maps, nil
}

