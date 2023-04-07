package controllers

import (
	"net/http"

	"br.com.charlesrodrigo/ms-person/model"
	"br.com.charlesrodrigo/ms-person/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Zipcode      string `json:"zipcode" binding:"required"`
	Street       string `json:"street" binding:"required"`
	Neighborhood string `json:"neighborhood" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	Country      string `json:"country" binding:"required"`
}

type CreatePersonRequest struct {
	Name    string  `json:"name" binding:"required"`
	Email   string  `json:"email" binding:"required"`
	Address Address `json:"address" binding:"required"`
}

type GetPersonRequest struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

type PersonController struct {
	PersonService service.PersonService
}

func NewPersonController(personService service.PersonService) PersonController {
	return PersonController{PersonService: personService}
}

// @BasePath /api/v1

// Create controllers.CreatePersonRequest godoc
// @Summary create person
// @Schemes
// @Description create person
// @Tags person
// @Accept json
// @Param user body controllers.CreatePersonRequest true "Person Data"
// @Produce json
// @Success 200
// @Router /api/v1/person [post]
func (personController PersonController) CreatePerson(c *gin.Context) {
	var input CreatePersonRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	personAddress := model.Address{
		Zipcode:      input.Address.Zipcode,
		Street:       input.Address.Street,
		Neighborhood: input.Address.Neighborhood,
		City:         input.Address.City,
		State:        input.Address.State,
		Country:      input.Address.Country,
	}
	person := model.Person{
		Name:    input.Name,
		Email:   input.Email,
		Address: personAddress,
	}

	personController.PersonService.Create(&person)

	c.JSON(http.StatusOK, "")
}

func (personController PersonController) UpdatePerson(c *gin.Context) {
	var input CreatePersonRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	personAddress := model.Address{
		Zipcode:      input.Address.Zipcode,
		Street:       input.Address.Street,
		Neighborhood: input.Address.Neighborhood,
		City:         input.Address.City,
		State:        input.Address.State,
		Country:      input.Address.Country,
	}
	person := model.Person{
		ID:      objectId,
		Name:    input.Name,
		Email:   input.Email,
		Address: personAddress,
	}

	personController.PersonService.Update(&person)

	c.JSON(http.StatusOK, person)
}

func (personController PersonController) GetPerson(c *gin.Context) {

	person := model.Person{}

	person = personController.PersonService.FindById(c.Param("id"))

	if person.ID.IsZero() {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	personAddress := Address{
		Zipcode:      person.Address.Zipcode,
		Street:       person.Address.Street,
		Neighborhood: person.Address.Neighborhood,
		City:         person.Address.City,
		State:        person.Address.State,
		Country:      person.Address.Country,
	}

	output := GetPersonRequest{
		Id:      c.Param("id"),
		Name:    person.Name,
		Email:   person.Email,
		Address: personAddress,
	}

	c.JSON(http.StatusOK, output)
}
