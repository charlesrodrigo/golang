package controllers

import (
	"net/http"

	"br.com.charlesrodrigo/ms-person/api/dto"
	"br.com.charlesrodrigo/ms-person/helper/function"
	"br.com.charlesrodrigo/ms-person/infra/database"
	"br.com.charlesrodrigo/ms-person/internal/model"
	"br.com.charlesrodrigo/ms-person/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PersonController struct {
	PersonService service.PersonService
}

func NewApiPersonController(router *gin.Engine) {
	personRepository := database.NewPersonRepositoryImpl()
	personService := service.NewPersonServiceImpl(personRepository)
	personController := newPersonController(personService)

	v1 := router.Group("/api/v1")
	{
		personGroup := v1.Group("/person")

		{
			personGroup.POST("/", personController.CreatePerson)
			personGroup.PUT("/:id", personController.UpdatePerson)
			personGroup.GET("/:id", personController.GetPerson)
			personGroup.GET("/", personController.GetAllPerson)
			personGroup.DELETE("/:id", personController.DeletePerson)
		}
	}
}

func newPersonController(personService service.PersonService) PersonController {
	return PersonController{PersonService: personService}
}

// @BasePath /api/v1
// Person godoc
// @Summary create person
// @Schemes
// @Description create person
// @Tags person
// @Accept json
// @Param person body dto.CreatePersonRequest true "Person Data"
// @Produce json
// @Success 200
// @Router /api/v1/person [post]
func (pc PersonController) CreatePerson(context *gin.Context) {
	ctx := context.Request.Context()

	var createPersonRequest dto.CreatePersonRequest
	if err := context.ShouldBindJSON(&createPersonRequest); err != nil {
		context.AbortWithStatusJSON(function.CreateResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	person := createPersonRequest.ParseDTOToModel()

	err := pc.PersonService.Create(ctx, &person)

	if err != nil {
		context.AbortWithStatusJSON(function.CreateResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	context.JSON(http.StatusOK, "")
}

// @BasePath /api/v1
// Person godoc
// @Summary update person
// @Schemes
// @Description update person
// @Tags person
// @Accept json
// @Param id path string true "id person"
// @Param person body dto.CreatePersonRequest true "Person Data"
// @Produce json
// @Success      200  {object}  dto.CreatePersonRequest
// @Router /api/v1/person/{id} [put]
func (pc PersonController) UpdatePerson(context *gin.Context) {
	ctx := context.Request.Context()
	var updatePersonRequest dto.CreatePersonRequest
	if err := context.ShouldBindJSON(&updatePersonRequest); err != nil {
		context.AbortWithStatusJSON(function.CreateResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	objectId, err := primitive.ObjectIDFromHex(context.Param("id"))

	if err != nil {
		context.AbortWithStatusJSON(function.CreateResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	person := updatePersonRequest.ParseDTOToModel()
	person.ID = objectId

	err = pc.PersonService.Update(ctx, &person)

	if err != nil {
		context.AbortWithStatusJSON(function.CreateResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	context.JSON(http.StatusOK, person)
}

// @BasePath /api/v1
// Person godoc
// @Summary get person
// @Schemes
// @Description get person
// @Tags person
// @Accept json
// @Param id path string true "id person"
// @Produce json
// @Success      200  {object}  dto.GetPersonRequest
// @Router /api/v1/person/{id} [get]
func (pc PersonController) GetPerson(context *gin.Context) {
	ctx := context.Request.Context()
	person := model.Person{}

	person, err := pc.PersonService.FindById(ctx, context.Param("id"))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if person.ID.IsZero() {
		context.AbortWithStatusJSON(function.CreateResponseError(http.StatusBadRequest, "Not found"))
		return
	}

	response := dto.GetPersonRequest{}
	response.ParseModelToDTO(person)

	context.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// Person godoc
// @Summary get all person
// @Schemes
// @Description get all person
// @Tags person
// @Accept json
// @Produce json
// @Success      200  {object}  []dto.GetPersonRequest
// @Router /api/v1/person [get]
func (pc PersonController) GetAllPerson(context *gin.Context) {
	ctx := context.Request.Context()
	persons := pc.PersonService.FindAll(ctx)

	response := make([]dto.GetPersonRequest, 0)
	for _, person := range persons {

		getPersonRequest := dto.GetPersonRequest{}
		getPersonRequest.ParseModelToDTO(person)

		response = append(response, getPersonRequest)
	}

	context.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// Person godoc
// @Summary delete person
// @Schemes
// @Description delete person
// @Tags person
// @Accept json
// @Param id path string true "id person"
// @Produce json
// @Success 200
// @Router /api/v1/person/{id} [delete]
func (pc PersonController) DeletePerson(context *gin.Context) {
	ctx := context.Request.Context()
	err := pc.PersonService.Delete(ctx, context.Param("id"))

	if err != nil {
		context.AbortWithStatusJSON(function.CreateResponseError(http.StatusBadRequest, err.Error()))
		return
	}
	context.JSON(http.StatusOK, "")
}
