package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"author": "Muhammad Farhan Mardadi"})
}

func GetPositions(c *gin.Context) {
	var queries []string
	var response *http.Response
	var err error

	url := "http://dev3.dansmultipro.co.id/api/recruitment/positions.json"

	queryParams := c.Request.URL.Query()
	for key, values := range queryParams {
		for _, value := range values {
			queries = append(queries, key+"="+value)
		}
	}

	query := strings.Join(queries, "&")

	if len(query) > 0 {
		response, err = http.Get(url + "?" + query)
	} else {
		response, err = http.Get(url)
	}

	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var result json.RawMessage
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing JSON")
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetJobDetail(c *gin.Context) {
	url := "http://dev3.dansmultipro.co.id/api/recruitment/positions/%s"

	jobID := c.Param("id")

	response, err := http.Get(fmt.Sprintf(url, jobID))
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var result json.RawMessage
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing JSON")
		return
	}

	c.JSON(http.StatusOK, result)
}
