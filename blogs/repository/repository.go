package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	pb "github.com/whuangz/micro-blog/blogs/proto"
)

type Service interface {
	FetchArticles() ([]*Article, error)
	CreateArticle(ctx context.Context, req *pb.Article) error
	DeleteArticle(ctx context.Context, req *pb.Article) error
	UpdateArticle(ctx context.Context, req *pb.Article) error

	FetchTopics(ctx context.Context, req *pb.ListRequest) ([]*Topic, error)
	CreateTopic(ctx context.Context, req *pb.Topic) error
	GetTopic(ctx context.Context, req *pb.Topic) error

	Migrate()
	Close() error
}

type Article struct {
	ID       uint32 `gorm:"AUTO_INCREMENT;primary_key;type:int(11)"`
	Title    string `gorm:"type:varchar(225)"`
	Slug     string `gorm:"type:varchar(225)"`
	Content  string `gorm:"type:text"`
	AuthorID uint32 `grom:"type:int(11)`
	TopicID  uint32 `grom:"type:int(11)`
	Status   string `gorm:"type:varchar(225)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Topic struct {
	ID    uint32 `gorm:"primary_key;type:int(11)"`
	Title string `gorm:"type:varchar(225)"`
	Slug  string `gorm:"type:varchar(225)"`
	Tag   string `gorm:"type:varchar(225)"`
	Color string `gorm:"type:varchar(8)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type articleRepository struct {
	db *gorm.DB
}

func New(dbHost, dbName, dbUser, dbPass, dbPort string) (Service, error) {
	src := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql", src)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &articleRepository{db}, nil
}

// Close will terminate the database connection
func (m *articleRepository) Close() error {
	return m.db.Close()
}

func (m *articleRepository) Migrate() {
	m.db.Debug().AutoMigrate(
		&Article{},
		&Topic{},
	)

}
