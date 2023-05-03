package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platforn string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

// In `json:"tags,omitempty"` we can't use space between tags, and omitempty

func main() {
	fmt.Println("Parsing json Data")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs", 299, "Udemy", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 199, "Udemy", "xyz123", []string{"web-dev", "fulstack"}},
		{"Angular", 599, "Udemy", "123abs", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs",
		"Price": 299,
		"website": "Udemy",
		"tags": ["web-dev","js"]
	}`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	//some cases where we just want to add data to key value

	var myOnlineData = make(map[string]interface{})
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is : %T\n", k, v, v)
	}

}
