package config

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saifwork/portfolio-service.git/app/configs"
	"github.com/saifwork/portfolio-service.git/app/services/core/responses"
	"github.com/saifwork/portfolio-service.git/app/services/core/utils"
	"github.com/saifwork/portfolio-service.git/app/services/domain/config/dtos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PortfolioService struct {
	gin        *gin.Engine
	conf       *configs.Config
	repository *PortfolioRepository
}

func NewPortfolioService(gin *gin.Engine, conf *configs.Config, repo *PortfolioRepository) *PortfolioService {
	return &PortfolioService{
		gin:        gin,
		conf:       conf,
		repository: repo,
	}
}

func (s *PortfolioService) SetupRoutes() {
	g := s.gin.Group("")

	g.GET("/", s.GetAllConfigs)
	g.POST("/", s.PostContact)
}

func (s *PortfolioService) GetAllConfigs(c *gin.Context) {

	var home, about, skills, social, footer, resume map[string]interface{} // Object `{}` JSON
	var projects, experience []any                                         // Array `[]` JSON

	// Load JSON files without using structs
	if err := utils.LoadJSONFile("home.json", &home); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if err := utils.LoadJSONFile("about.json", &about); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if err := utils.LoadJSONFile("skills.json", &skills); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if err := utils.LoadJSONFile("projects.json", &projects); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if err := utils.LoadJSONFile("experience.json", &experience); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if err := utils.LoadJSONFile("social.json", &social); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if err := utils.LoadJSONFile("footer.json", &footer); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	if err := utils.LoadJSONFile("resume.json", &resume); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	// Return all data directly without using structs
	c.JSON(http.StatusOK, responses.NewSuccessResponse(gin.H{
		"home":       home,
		"about":      about,
		"skills":     skills,
		"projects":   projects,
		"experiences": experience,
		"social":     social,
		"footer":     footer,
		"resume":     resume,
	}))
}

func (s *PortfolioService) PostContact(c *gin.Context) {
	var contactReq dtos.ContactReqDto

	if err := c.ShouldBindJSON(&contactReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	// Ensure default values
	contactReq.ID = primitive.NewObjectID() // Assign new ObjectID

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.repository.Insert(ctx, &contactReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, responses.NewSuccessResponse("Will get back to you shortly!"))
}
