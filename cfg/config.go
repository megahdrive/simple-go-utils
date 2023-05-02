package cfg

import (
	"Program/vld"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
)

var config []byte

func init() {
	var err error

	configfile, err := os.Open("config.json")
	vld.Check(err)
	defer configfile.Close()

	config, err = ioutil.ReadAll(configfile)
	vld.Check(err)
}

func GetKey(key string) string {
	value := gjson.GetBytes(config, key)
	return value.String()
}
