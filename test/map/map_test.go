/**
 * @Author: DollarKiller
 * @Description: map test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 13:23 2019-10-08
 */
package _map

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestMapUn(t *testing.T) {
	data := `
	{
		"device": "this is device",
		"data": [
			{
				"humidity": "27",
				"time": "this is time"
			}
		]
	}
	`
	m := make(map[string]interface{})

	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		panic(err)
	}

	log.Println(m["device"])

	//log.Println(m["data"])

	//i := m["data"].([]interface{})[0].(map[string]interface{})
	//
	//log.Println(i["humidity"])
	//fmt.Printf("%T",i)
	i, err := toMap(m["data"])
	if err != nil {
		panic(err)
	}
	log.Println(i["time"])
}

type mapun struct {
	dataStr string
	mc      map[string]interface{}
}

func Unmarshal(jsn string) (*mapun, error) {
	m := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsn), &m)
	if err != nil {
		return nil, err
	}

	return &mapun{
		dataStr: jsn,
		mc:      m,
	}, nil
}

func (m *mapun) getString(key string) (string,bool) {
	s,ok := m.mc[key].(string)
	return s,ok
}

func (m *mapun) getMap (key string) (map[string]interface{}, bool) {
	i,ok := m.mc[key].([]interface{})[0].(map[string]interface{})
	return i,ok
}


func toMap(data interface{}) (map[string]interface{}, error) {
	i, ok := data.([]interface{})[0].(map[string]interface{})
	if ok {
		return i, nil
	} else {
		return nil, fmt.Errorf("json to map error")
	}
}
