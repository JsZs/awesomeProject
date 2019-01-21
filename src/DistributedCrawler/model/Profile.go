package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hokou      string
	Xinzuo     string
	House      string
	Car        string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o) //数据转换为Json字符串
	if err != nil {
		return profile, err
	}

	json.Unmarshal(s, &profile) //将Json解码到对应的数据结构
	return profile, err
}
