/**
 * 1、通过simplejson生成一个json
 * 2、通过simplejson解析json
 * 3、interface使用json.Marshal
 * 4、总结一点: 可以自己设置每一个json的节点, 也可以通过结构体序列化成json, 也可以通过任何的数据结构序列化成json(slice就是数组, map就是json的对象)
 */
package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"strings"
)

type returnData struct {
	Code		int					`json:"code,omitempty"`
	Data 		interface{}			`json:"data,omitempty"`
	Message		string				`json:"message,omitempty"`
}

type user struct {
	Name		string 				`json:"name,omitempty"`
	Age 		int					`json:"age,omitempty"`
	Address		*address			`json:"address,omitempty"`
	Company		[]string 			`json:"company,omitempty"`
	School		map[string]string	`json:"school,omitempty"`
}

type address struct {
	Province 	string 				`json:"province,omitempty"`
	City		string 				`json:"city,omitempty"`
	PostCode	int 				`json:"postcode,omitempty"`
}

// 生成一个返回值结构体对象
func response() *returnData {
	code := 0
	message := "success"
	name := "zhanglei"
	age := 34
	addressInfo := &address{
		Province: "北京",
		City: "北京",
		PostCode: 100095,
	}
	companys := []string{"奇虎360", "百度", "汽车之家"}
	schools := map[string]string{
		"primary": "中心小学",
		"middle": "中学",
		"senior": "襄安中学",
	}
	userInfo := &user{
		Name: name,
		Age: age,
		Address: addressInfo,
		Company: companys,
		School: schools,
	}

	ret := &returnData{}
	ret.Code = code
	ret.Message = message
	ret.Data = userInfo

	return ret
}

// 通过simplejson生成json, json中字段值字段名都自己设置
func genJson1() string {
	ret := response()

	// returnData
	retSimpleJson := simplejson.New()
	retSimpleJson.Set("code", ret.Code)
	retSimpleJson.Set("message", ret.Message)

	userInfo, _ := ret.Data.(*user)

	// address
	addressSimpleJson := simplejson.New()
	addressSimpleJson.Set("province1", userInfo.Address.Province)
	addressSimpleJson.Set("city1", userInfo.Address.City)
	addressSimpleJson.Set("postcode1", userInfo.Address.PostCode)

	// userinfo
	userInfoSimpleJson := simplejson.New()
	userInfoSimpleJson.Set("name1", userInfo.Name)
	userInfoSimpleJson.Set("age1", userInfo.Age)
	userInfoSimpleJson.Set("address1", addressSimpleJson)
	userInfoSimpleJson.Set("company1", userInfo.Company)
	userInfoSimpleJson.Set("school1", userInfo.School)

	retSimpleJson.Set("data", userInfoSimpleJson)

	b, _ := retSimpleJson.MarshalJSON()

	return string(b)
}

// 通过simplejson生成json, ret.Data是个结构体, 序列化成json的时候, 按照结构体中`json:""`的方式序列化
func genJson2() string {
	ret := response()

	retSimpleJson := simplejson.New()
	retSimpleJson.Set("code", ret.Code)
	retSimpleJson.Set("data", ret.Data)
	retSimpleJson.Set("message", ret.Message)

	b, _ := retSimpleJson.MarshalJSON()

	return string(b)
}

// 通过simplejson生成json, 没有结构体, 将slice、map单独写入到json
func genJson3() string {
	retSimpleJson := simplejson.New()

	retSimpleJson.Set("code", 0)
	retSimpleJson.Set("message", "success")

	data := map[string]interface{}{
		"address3": map[string]interface{}{
			"province3": "北京",
			"city3": "北京",
			"postcode3": 100095,
		},
		"age3": "34", // 上面是int, 这里改成string, json也会序列化成string
		"company3": []string{"奇虎360", "百度", "汽车之家"},
		"school3": map[string]interface{}{
			"primary3": "小学",
			"middle3": "中学",
			"senior": "高中",
		},
	}
	retSimpleJson.Set("data", data)

	b, _ := retSimpleJson.MarshalJSON()

	return string(b)
}

// 通过simplejson解析json
func parseJson(s string) {
	retSimpleJson, _ := simplejson.NewJson([]byte(s))
	code := retSimpleJson.Get("code").MustInt()
	message := retSimpleJson.Get("message").MustString()
	dataSimpleJson := retSimpleJson.Get("data")
	fmt.Printf("code: %d\n", code)
	fmt.Printf("message: %s\n", message)

	name := dataSimpleJson.Get("name").MustString()
	age := dataSimpleJson.Get("age").MustInt()
	province := dataSimpleJson.Get("address").Get("province").MustString()
	city := dataSimpleJson.Get("address").Get("city").MustString()
	postcode := dataSimpleJson.Get("address").Get("postcode").MustInt()
	companyArray := dataSimpleJson.Get("company").MustArray()
	companys := make([]string, 0, len(companyArray))
	for _, v := range companyArray {
		vt, _ := v.(string)
		companys = append(companys, vt)
	}
	school := dataSimpleJson.Get("school").MustMap()

	fmt.Printf("name: %s\n", name)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("province: %s\n", province)
	fmt.Printf("city: %s\n", city)
	fmt.Printf("postcode: %d\n", postcode)
	fmt.Printf("company: %s\n", strings.Join(companys, ", "))
	for i, v := range school {
		fmt.Printf("%s: %s\n", i, v)
	}
}

// 序列化interface(包含, int, string, slice, map, *simplejson.Json)
func parseInterface() string {
	infoSimpleJson := simplejson.New()
	infoSimpleJson.Set("applicability", []string{"docker", "rpc", "k8s", "etcd"})
	infoSimpleJson.Set("platform", map[string]bool{
		"windows": true,
		"linux": true,
	})
	infoSimpleJson.Set("editors", "goland")

	ret := map[string]interface{}{
		"code": 0,
		"message": "success",
		"data": map[string]interface{}{
			"id": 1,
			"title": "golang是一门好语言",
			"authors": []string{"Robert Griesemer", "Rob Pike", "Ken Thompson"},
			"types": map[string][]interface{}{
				"int": {1, 2},
				"slice": {"[]int", "[]string"},
				"map": {"map[string]string", "map[int]map[string][string]"},
			},
			"info": infoSimpleJson,
		},
	}

	b, _ := json.Marshal(ret)

	return string(b)
}

func main() {
	ret1 := genJson1()
	fmt.Println("genJson1")
	fmt.Println(ret1)
	fmt.Println()

	fmt.Println("genJson2")
	ret2 := genJson2()
	fmt.Println(ret2)
	fmt.Println()

	fmt.Println("genJson3")
	ret3 := genJson3()
	fmt.Println(ret3)
	fmt.Println()

	fmt.Println("parseJson")
	parseJson(ret2)
	fmt.Println()

	fmt.Println("parseInterface")
	ret4 := parseInterface()
	fmt.Println(ret4)
}