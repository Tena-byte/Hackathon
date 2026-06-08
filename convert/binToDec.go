package main

import (
	"strconv"
)

func binToDec(s string) (int64, error) {

	return strconv.ParseInt(s, 2, 64)

}
