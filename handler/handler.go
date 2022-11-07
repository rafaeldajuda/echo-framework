package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetCats - GET API which return the name of the cats specified in QueryParam
// http://localhost:8000/cats/json?name=arnold&type=fluffy
// data path variable accepts value as json/string

func GetCats(c echo.Context) (err error) {

	catName := c.QueryParam("name")
	catType := c.QueryParam("type")
	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is : %s\nand cat type is : %s\n", catName, catType))
	} else if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "Please specify the data type as String or JSON",
	})

}

func AddCat(c echo.Context) (err error) {

	type Cat struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	cat := Cat{}
	defer c.Request().Body.Close()
	err = json.NewDecoder(c.Request().Body).Decode(&cat)
	if err != nil {
		log.Printf("Falied reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	log.Printf("this is your cat %#v", cat)
	return c.String(http.StatusOK, "We got your Cat!!!")

}
