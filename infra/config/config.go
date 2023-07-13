package config

import (
	"os"
	"log"
)

var PORT string
var MONGO_URI_MP string
var MONGO_DB_NAME_MP string
var MONGO_URI_CMS string
var MONGO_DB_NAME_CMS string

func init(){
	PORT = checkEnv("PORT")
	MONGO_URI_MP = checkEnv("MONGO_URI_MP")
	MONGO_DB_NAME_MP = checkEnv("MONGO_DB_NAME_MP")
	MONGO_URI_CMS = checkEnv("MONGO_URI_CMS")
	MONGO_DB_NAME_CMS = checkEnv("MONGO_DB_NAME_CMS")
}

func checkEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		log.Fatal(key + " config env is required")
	}
	return env
}
