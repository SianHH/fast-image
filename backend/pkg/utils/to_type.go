package utils

import (
	"encoding/json"
	"strconv"
)

func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StructToBytes(data any) []byte {
	marshal, _ := json.Marshal(data)
	return marshal
}

func BytesToStruct(data []byte, target any) {
	_ = json.Unmarshal(data, target)
}
