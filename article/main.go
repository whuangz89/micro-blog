package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	micro "github.com/micro/go-micro/v2"
	"github.com/whuangz/micro-blog/article/handler"
	"github.com/whuangz/micro-blog/article/proto/article"
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
		micro.Name("blog.article.service"),
	)
	srv.Init()

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	h := handler.NewArticleHandler(dbConn)
	article.RegisterArticleServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
