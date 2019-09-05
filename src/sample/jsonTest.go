package sample

/**
 * Unmarsha1()  字符串转 map/struct
 * Marsha1()
 * json.NewDecoder(os.Stdin) 用于支持 JSON 数据的流式读写
 * json.NewEncoder(os.Stdout)
 *
 * 1. struct 首字母必须大写
 * 2. json 字符串大小写不敏感
 * 3. 如果struct与json字符串不对应，需要配置tag
 *
 * */
import (
	"encoding/json"
	"fmt"
	"os"
)

type configstruct struct {
	Host              string   `json:"host"`
	Port              int      `json:"port"`
	AnalyticsFile     string   `json:"analytics_file"`
	StaticFileVersion int      `json:"static_file_version"`
	StaticDir         string   `json:"static_dir"`
	TemplatesDir      string   `json:"templates_dir"`
	SerTCPSocketHost  string   `json:"serTcpSocketHost"`
	SerTCPSocketPort  int      `json:"serTcpSocketPort"`
	Fruits            []string `json:"fruits"`
}

type other struct {
	SerTCPSocketHost string   `json:"serTcpSocketHost"`
	SerTCPSocketPort int      `json:"serTcpSocketPort"`
	Fruits           []string `json:"fruits"`
}

// JSONTest :
func JSONTest() {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`

	//json str 转map
	var dataMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dataMap); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dataMap)
		fmt.Println(dataMap["host"])
	}

	//json str 转struct
	var config configstruct
	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(config)
		fmt.Println(config.Host)
	}

	//json str 转struct(部份字段)
	var part other
	if err := json.Unmarshal([]byte(jsonStr), &part); err == nil {
		fmt.Println("================json str 转part struct==")
		fmt.Println(part)
		fmt.Println(part.SerTCPSocketPort)
	}

	//struct 到json str
	if b, err := json.Marshal(config); err == nil {
		fmt.Println("================struct 到json str==")
		fmt.Println(string(b))
	}

	//map 到json str
	fmt.Println("================map 到json str=====================")
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(dataMap)

	//array 到 json str
	arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
	lang, err := json.Marshal(arr)
	if err == nil {
		fmt.Println("================array 到 json str==")
		fmt.Println(string(lang))
	}

	//json 到 []string
	var wo []string
	if err := json.Unmarshal(lang, &wo); err == nil {
		fmt.Println("================json 到 []string==")
		fmt.Println(wo)
	}
}
