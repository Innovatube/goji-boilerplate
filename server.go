package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/canhlinh/g-w-chat/models"
	"github.com/google/go-gcm"
	"github.com/googollee/go-socket.io"
	"github.com/hypebeast/gojistatic"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/takaaki-mizuno/goji-boilerplate/app"
	"github.com/takaaki-mizuno/goji-boilerplate/app/services"
	"github.com/zenazn/goji"
)

const (
	// API key from Cloud console
	API_KEY = "AIzaSyAhzwFU70iyASlsZzY44ld8uY7R3HqmS2U"
	// GCM sender ID
	SENDER_ID = "755990245275"
	// The name of the database to connect to
	DATABASE_URL = "kyo:123456a@A@/chat?charset=utf8&parseTime=True&loc=Local"

	WEBSOCKET_PORT = "4260"

	ACTION_KEY          = "action"
	REGISTER_CLIENT     = "register_new_client"
	UNREGISTER_CLIENT   = "unregister_client"
	TOKEN               = "registration_token"
	STRING_IDENTIFIER   = "stringIdentifier"
	STATUS_REGISTERED   = "registered"
	STATUS_UNREGISTERED = "unregistered"
)

var (
	// Print logging
	debug = true

	// Current database connection
	db *gorm.DB

	// Websocket connection
	socket socketio.Socket

	mUser *models.User

	mChatSession *models.ChatSession

	mChatSessionUser *models.ChatSessionUser

	mMessage *models.Message

	mGcmToken *models.GcmToken
)

func LoadFileConfig() {
	services.ConfigService().LoadConfigFile("config.yaml")
}

func LogSetupAndDestruct() func() {
	logFilePath := services.ConfigService().GetConfig("systemLog:path", "/var/log/boilerplate.log")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Panicln(err)
	}

	log.SetFormatter(&log.TextFormatter{DisableColors: true})
	log.SetOutput(logFile)
	log.SetLevel(log.DebugLevel)

	return func() {
		e := logFile.Close()
		if e != nil {
			fmt.Fprintf(os.Stderr, "Problem closing the log file: %s\n", e)
		}
	}
}

func InitDb() {
	var err error
	db, err = gorm.Open("mysql", DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.DB()
	mUser = &models.User{}
	mChatSession = &models.ChatSession{}
	mChatSessionUser = &models.ChatSessionUser{}
	mMessage = &models.Message{}

	db.DropTableIfExists(mMessage)
	db.DropTableIfExists(mChatSessionUser)
	db.DropTableIfExists(mChatSession)
	db.DropTableIfExists(mGcmToken)
	db.DropTableIfExists(mUser)

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(mUser, mChatSession, mChatSessionUser, mMessage, mGcmToken)
	db.Model(mChatSession).AddForeignKey("owner_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(mChatSessionUser).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(mChatSessionUser).AddForeignKey("chat_session_id", "chat_sessions(id)", "RESTRICT", "RESTRICT")
	db.Model(mMessage).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(mMessage).AddForeignKey("chat_session_id", "chat_sessions(id)", "RESTRICT", "RESTRICT")
	db.Model(mGcmToken).AddForeignKey("id", "users(id)", "RESTRICT", "RESTRICT")
	db.LogMode(debug) // Helps with debugging
}

func gcmMessageHandler(cm gcm.CcsMessage) error {
	log.Println("Recieved message %+v", cm)

	data := cm.Data
	switch data[ACTION_KEY] {
	case REGISTER_CLIENT:
		break
	case UNREGISTER_CLIENT:
		log.Println("unregister")
		break
	}
	return nil
}

func GcmRoutin() {
	err := gcm.Listen(SENDER_ID, API_KEY, gcmMessageHandler, nil)
	if err != nil {
		log.Printf("Listen error: %v", err)
	}
}

func main() {
	LoadFileConfig()
	defer LogSetupAndDestruct()()
	log.Info("App start")

	staticFolder := services.ConfigService().GetConfig("staticFolder", "app/public")
	goji.Use(gojistatic.Static(staticFolder, gojistatic.StaticOptions{SkipLogging: false}))
	log.Debug("Static folder : " + staticFolder)

	InitDb()
	go GcmRoutin()

	port := services.ConfigService().GetConfig("port", ":8000")
	flag.Set("bind", port)

	app.App()

	goji.Serve()
}
