package config

import (
	"log"
	"os"
)

var (
	//database
	DBUserName = getEnv("MYSQL_USERNAME")
	DBPassword = getEnv("MYSQL_PASSWORD")
	DBName     = getEnv("MYSQL_DATABASE")
	DBHost     = getEnv("MYSQL_HOST")
	DBPort     = getEnv("MYSQL_PORT")

	//sftp
	SFTPUser = getEnv("SFTP_USER")
	SFTPPass = getEnv("SFTP_PASS")
	SFTPHost = getEnv("SFTP_HOST")

	//mailjet
	MailJetKey    = getEnv("MAILJET_KEY")
	MailJetSecret = getEnv("MAILJET_SECRET")
	Email         = getEnv("EMAIL")

	//redis
	RedisHost = getEnv("REDISHOST")
)

func getEnv(i string) string {
	l, isExist := os.LookupEnv(i)
	if !isExist {
		log.Printf("%s cannot find the environment", i)
	}

	return l
}
