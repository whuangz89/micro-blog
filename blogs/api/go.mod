module github.com/whuangz/micro-blog/blogs

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/gosimple/slug v1.9.0
	github.com/jinzhu/gorm v1.9.12
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/micro/go-micro/v2 v2.4.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.6.2
	github.com/whuangz/micro-blog/helpers/sql v0.0.0-00010101000000-000000000000
)

replace github.com/whuangz/micro-blog/helpers/sql => ../../helpers/sql
