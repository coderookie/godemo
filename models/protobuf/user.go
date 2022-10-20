package user

import (
	"encoding/json"
	userProto "godemo/models/protobuf_package/user"
	"github.com/golang/protobuf/proto"
)

// 将结构体转化成proto
func GenerateUser2Protobuf() (string, error) {
	qihuCompanyName := "奇虎"
	qihuCompanyGroup := "信息流"
	qihuCompanyPosition := "互联网开发专家"
	qihuCompanyAddressProvince := "北京市"
	qihuCompanyAddressCity := "北京市"
	qihuCompanyAddressArea := "朝阳区"
	qihuCompanyAddressAddress := "北京市朝阳区将台"

	baiduCompanyName := "奇虎"
	baiduCompanyGroup := "信息流"
	baiduCompanyPosition := "互联网开发专家"
	baiduCompanyAddressProvince := "北京市"
	baiduCompanyAddressCity := "北京市"
	baiduCompanyAddressArea := "朝阳区"
	baiduCompanyAddressAddress := "北京市朝阳区将台"
	companys := []*userProto.Company{
		&userProto.Company{
			Name: &qihuCompanyName,
			Group: &qihuCompanyGroup,
			Position: &qihuCompanyPosition,
			Address: &userProto.Address{
				Province: &qihuCompanyAddressProvince,
				City: &qihuCompanyAddressCity,
				Area: &qihuCompanyAddressArea,
				Address: &qihuCompanyAddressAddress,
			},
		},
		&userProto.Company{
			Name: &baiduCompanyName,
			Group: &baiduCompanyGroup,
			Position: &baiduCompanyPosition,
			Address: &userProto.Address{
				Province: &baiduCompanyAddressProvince,
				City: &baiduCompanyAddressCity,
				Area: &baiduCompanyAddressArea,
				Address: &baiduCompanyAddressAddress,
			},
		},
	}

	userName := "张磊"
	var userSex int32 = 1
	var userAge int32 = 33
	province := "北京市"
	city := "北京市"
	area := "西城区"
	address := "北京市西城区广安门外街道"
	zhanglei := &userProto.User{
		Name: &userName,
		Sex: &userSex,
		Age: &userAge,
		Address: &userProto.Address{
			Province: &province,
			City: &city,
			Area: &area,
			Address: &address,
		},
		Company: companys,
	}

	data, err := proto.Marshal(zhanglei)

	return string(data), err
}

// 将proto解析成结构体
func ParseProtobuf2User(data string) (*userProto.User, error) {
	zhanglei := &userProto.User{}
	err := proto.Unmarshal([]byte(data), zhanglei)
	return zhanglei, err
}

// 将结构体转化成json
func ParseUser2Json(user *userProto.User) (string, error) {
	data, err := json.Marshal(user)
	return string(data), err
}