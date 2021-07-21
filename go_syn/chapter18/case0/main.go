package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {

	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("err %v \n", err)
		return
	}

	_, err = c.Do("set", "nameq", "songyanan宋亚楠112233")

	if err != nil {
		fmt.Printf("set err %v\n", err)
	}

	r, err := redis.String(c.Do("get", "nameq"))
	if err != nil {
		fmt.Printf("get err %v\n", err)
	}

	fmt.Printf("get name = %v \n", r)

	//fmt.Println("coom", c)
}
