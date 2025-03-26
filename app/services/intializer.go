package services

import (
	"github.com/gin-gonic/gin"
	"github.com/saifwork/portfolio-service.git/app/configs"
	"github.com/saifwork/portfolio-service.git/app/services/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Initializer struct {
	gin    *gin.Engine
	conf   *configs.Config
	client *mongo.Client
}

func NewInitializer(gin *gin.Engine, conf *configs.Config, cli *mongo.Client) *Initializer {
	s := &Initializer{
		gin:    gin,
		conf:   conf,
		client: cli,
	}
	return s
}

func (s *Initializer) RegisterDomains(domains []domain.IDomain) {
	for _, domain := range domains {
		domain.SetupRoutes()
	}
}
