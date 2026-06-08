package main

import (
	"strconv"
)

func hexToDec(s string) (int64, error) {

	return strconv.ParseInt(s, 16, 64)

}