package config

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Config = v
}

func GetJoin(key string) string  {
	a := Config.Get(key)
	return configJoin(a.(map[string]interface{}))
}

func configJoin(m map[string]interface{}) string {
	var str bytes.Buffer
	for k, v := range m {
		str.WriteString(k)
		str.WriteString("=")
		str.WriteString(v.(string))
		str.WriteString(" ")
	}
	return str.String()
}
func GetRepo(key string) (string, string)  {
	a := Config.Get(key)

	// util.Dump(a.(map[string]interface{}))
	// driver := a.(map[string]interface{})["driver"].(string) 
	driver := "postgres" 
	// "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"
	user := a.(map[string]interface{})["user"].(string) 
	pass := a.(map[string]interface{})["password"].(string) 
	host := a.(map[string]interface{})["host"].(string) 
	port := a.(map[string]interface{})["port"].(string)
	db := a.(map[string]interface{})["dbname"].(string) 
	config := driver + "://" + user + ":" + pass + "@" + host + ":" + port + "/" + db
	return driver, config
}