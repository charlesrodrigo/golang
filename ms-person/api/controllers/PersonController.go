package controllers

import (
	"log"
	"net/http"

	"br.com.charlesrodrigo/ms-person/api/dto"
	"br.com.charlesrodrigo/ms-person/internal/model"
	"br.com.charlesrodrigo/ms-person/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PersonController struct {
	PersonService service.PersonService
}

func NewPersonController(personService service.PersonService) PersonController {
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
func (personController PersonController) CreatePerson(context *gin.Context) {
	var createPersonRequest dto.CreatePersonRequest
	if err := context.ShouldBindJSON(&createPersonRequest); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person := createPersonRequest.ParseDTOToModel()

	err := personController.PersonService.Create(context, &person)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
func (personController PersonController) UpdatePerson(context *gin.Context) {
	var updatePersonRequest dto.CreatePersonRequest
	if err := context.ShouldBindJSON(&updatePersonRequest); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(context.Param("id"))

	objectId, err := primitive.ObjectIDFromHex(context.Param("id"))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person := updatePersonRequest.ParseDTOToModel()
	person.ID = objectId

	err = personController.PersonService.Update(context, &person)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
func (personController PersonController) GetPerson(context *gin.Context) {

	person := model.Person{}

	person, err := personController.PersonService.FindById(context, context.Param("id"))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if person.ID.IsZero() {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not found"})
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
func (personController PersonController) GetAllPerson(context *gin.Context) {

	persons := personController.PersonService.FindAll(context)

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
func (personController PersonController) DeletePerson(context *gin.Context) {

	err := personController.PersonService.Delete(context, context.Param("id"))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, "")
}
