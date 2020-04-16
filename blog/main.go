package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/whuangz/micro-blog/blog/repository"

	micro "github.com/micro/go-micro/v2"
	blogHandler "github.com/whuangz/micro-blog/blog/handler"
	blogPb "github.com/whuangz/micro-blog/blog/proto"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	srv := micro.NewService(
		micro.Name("blog.service"),
	)
	srv.Init()

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	db, err := repository.New(
		dbHost, dbName, dbUser, dbPass, dbPort,
	)

	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	h := blogHandler.NewArticleHandler(db)
	blogPb.RegisterBlogServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
