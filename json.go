package main

import "encoding/json"
import "fmt"
import "os"



var pl = fmt.Println


type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int 
	Fruits []string
}


func main(){
	bolB, _ := json.Marshal(true)
	pl(string(bolB))

	intB, _ := json.Marshal(1)
	pl(string(intB))

	fltB, _ := json.Marshal(2.34)
	pl(string(fltB))


	strB, _ := json.Marshal("gopher")
	pl(string(strB))



	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	pl(string(slcB))


	mapD := map[string]int{"apple":5, "lettuce":7}
	mapB, _ := json.Marshal(mapD)
	pl(string(mapB))


	pl("=======")
	res1D := &Response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	res1B, _ := json.Marshal(res1D)
	pl(string(res1B))



	res2D := &Response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	res2B, _ := json.Marshal(res2D)
	pl(string(res2B))

	pl("==========")
	byt := []byte(`{"num":6.13, "strs":["a","b"]}`)

	var dat map[string]interface{}


	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	pl(dat)

	num := dat["num"].(float64)
	pl(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	pl(str1)

	str := `{"page":1, "fruits":["apple","peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	pl(res)
	pl(res.Fruits[0])



	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple":5, "lettuce":7}
	enc.Encode(d)

}
