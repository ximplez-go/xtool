package xtjson

import (
	"log"

	"github.com/bytedance/sonic"
)

var (
	defaultJson = sonic.ConfigDefault
)

func ToJsonSilent(data interface{}) string {
	return toJsonSilent(data, false)
}

func ToJsonSilentPretty(data interface{}) string {
	return toJsonSilent(data, true)
}

func toJsonSilent(data interface{}, pretty bool) string {
	if pretty {
		jsonByte, err := defaultJson.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Printf("[ToJson] error. %s", err.Error())
		}
		return string(jsonByte)
	} else {
		jsonByte, _ := defaultJson.Marshal(data)
		return string(jsonByte)
	}
}

func PhaseJsonSilent[T any](data []byte) *T {
	v := new(T)
	err := defaultJson.Unmarshal(data, v)
	if err != nil {
		log.Printf("[PhaseJson] error. %s", err.Error())
	}
	return v
}

func PhaseJsonFromString[T any](data string) *T {
	return PhaseJsonSilent[T]([]byte(data))
}
