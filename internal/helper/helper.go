package helper

import (
	"encoding/base64"
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

func Encoder(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func Decoder(str string, obj interface{}) {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, obj)
	if err != nil {
		panic(err)
	}
}

func GetUUID() string {
	return uuid.NewV4().String()
}
