package config

//import (
	//"github.com/tkanos/gonfig"
	//"log"
//)

type Configuration struct {
	PORT              string
	DB_USERNAME       string
	DB_PASSWORD       string
	CONNECTION_STRING string
}

var Config Configuration

func LoadConfig() {
	Config = Configuration{}
	Config.PORT=":9000"
	Config.DB_USERNAME="meetup"
	Config.DB_PASSWORD="9tS3qY8BKwXb1n8d"
	//err := gonfig.GetConf("./config/config.json", &Config)
	Config.CONNECTION_STRING = "mongodb+srv://" + Config.DB_USERNAME + ":" + Config.DB_PASSWORD + "@test-cluster-dvxai.mongodb.net/meetup"
	//if err != nil {
	//	log.Fatal(err)
	//}
}
