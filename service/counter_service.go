package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"

	"gorm.io/gorm"
)

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

// IndexHandler 计数器接口
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getIndex()
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	fmt.Fprint(w, data)
}

// CounterHandler 计数器接口
func CounterHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	if r.Method == http.MethodGet {

	} else if r.Method == http.MethodPost {
		maps, err := getAction(r)
		log.Printf("maps = %s",maps)
		if err != nil {
			return
		}
		action := maps["action"]
		if action == "get" {
			count, err := getCounter(r,maps)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = count
			}
		} else if action == "remark" {
			count, err := getRemark(r,maps)
			if err != nil {
				res.Code = -1
				res.ErrorMsg = err.Error()
			} else {
				res.Data = count
			}
		} else {
			count, err := modifyCounter(r,maps)
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

// modifyCounter 更新计数，自增或者清零
func getCounter(r *http.Request,maps map[string]string) ([]model.CounterModel, error) {
	action := maps["action"]
	user := maps["user"]
	log.Printf("action = %s,user=%s",action,user)
	var count []model.CounterModel
	var err error
	if action == "get" {
		count, err = getCurrentOrder(user)
		if err != nil {
			return nil, err
		}
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}

	return count, err
}

// modifyCounter 更新计数，自增或者清零
func getRemark(r *http.Request,maps map[string]string) (model.GoodModel, error) {
	action := maps["action"]
	order := maps["order"]
	log.Printf("action = %s,order=%s",action,order)
	var count model.GoodModel
	var err error
	if action == "remark" {
		count, err = getRemarkByID(order)
		if err != nil {
			return count, err
		}
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}

	return count, err
}

// modifyCounter 更新计数，自增或者清零
func modifyCounter(r *http.Request,maps map[string]string) (int32, error) {
	action := maps["action"]
	var err error
	var count int32
	if action == "add" {
		count, err = addCounter(r,maps)
		if err != nil {
			return 0, err
		}
	} else if action == "inc" {
		count, err = upsertCounter(r)
		if err != nil {
			return 0, err
		}
	} else if action == "clear" {
		err = clearCounter()
		if err != nil {
			return 0, err
		}
		count = 0
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}

	return count, err
}

// upsertCounter 更新或修改计数器
func upsertCounter(r *http.Request) (int32, error) {
	currentCounter, err := getCurrentCounter()
	var count int32
	createdAt := time.Now()
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	} else if err == gorm.ErrRecordNotFound {
		count = 1
		createdAt = time.Now()
	} else {
		count = currentCounter.Count + 1
		createdAt = currentCounter.CreatedAt
	}

	counter := &model.CounterModel{
		Id:        1,
		Count:     count,
		CreatedAt: createdAt,
		UpdatedAt: time.Now(),
	}
	err = dao.Imp.UpsertCounter(counter)
	if err != nil {
		return 0, err
	}
	return counter.Count, nil
}

// upsertCounter 更新或修改计数器
func addCounter(r *http.Request,maps map[string]string) (int32, error) {
	product := maps["product"]
	trade := maps["trade"]
	user := maps["user"]
	price := maps["price"]
	primary := maps["primary"]
	thumb := maps["thumb"]
	title := maps["title"]
	prices,_:= strconv.ParseInt(price,10,64)
	counter := &model.CounterModel{
		Count: 3,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Product: product,
		Trade: trade,
		User: user,
		Price: int(prices),
		Primary: primary,
		Thumb: thumb,
		Title: title,
	}
	log.Printf("map = %s",counter)
	err := dao.Imp.InsertCounter(counter)
	log.Printf("InsertCounter err = %s",err)
	return counter.Count, nil
}

func clearCounter() error {
	return dao.Imp.ClearCounter(1)
}

// getCurrentCounter 查询当前计数器
func getCurrentCounter() (*model.CounterModel, error) {

	counter, err := dao.Imp.GetCounter(1)
	if err != nil {
		return nil, err
	}

	return counter, nil
}

// getCurrentCounter 查询当前计数器
func getCurrentOrder(name string) ([]model.CounterModel, error) {
	counter, err := dao.Imp.GetOrder(name)
	if err != nil {
		return nil, err
	}
	return counter, nil
}

// getCurrentCounter 查询当前计数器
func getRemarkByID(order string) (model.GoodModel, error) {
	log.Printf("getRemark ID  = %s",order)
	counter, _ := dao.Imp.GetOrderById(order)
	log.Printf("order = %s",counter)
	product,_ := dao.ImpGood.GetGoodByID(counter.Product)
	log.Printf("product = %s",product)
	return product, nil
}

// getAction 获取action
func getAction(r *http.Request) (map[string]string,error) {
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
	user, ok := body["user"]
	if ok {
		maps["user"] = user.(string)
	}
	product, ok := body["product"]
	if ok {
		maps["product"] = product.(string)
	}
	order, ok := body["order"]
	if ok {
		maps["order"] = order.(string)
	}
	title, ok := body["title"]
	if ok {
		maps["title"] = title.(string)
	}
	price, ok := body["price"]
	if ok {
		maps["price"] = price.(string)
	}
	primaryImage, ok := body["primary"]
	if ok {
		maps["primary"] = primaryImage.(string)
	}
	thumb, ok := body["thumb"]
	if ok {
		maps["thumb"] = thumb.(string)
	}
	cloud, ok := body["cloud"]
	if ok {
		maps["cloud"] = cloud.(string)
	}
	return maps, nil
}

// getIndex 获取主页
func getIndex() (string, error) {
	b, err := ioutil.ReadFile("./index.html")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
