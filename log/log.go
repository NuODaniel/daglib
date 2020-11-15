package log

import (
	"fmt"
)

func Info(v ...interface{}){
	fmt.Printf("%v \n", v)
}
