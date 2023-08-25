package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Vehicle struct {
	ID    int
	Brand string
	Model string
	Year  int
	Price float64
}

var vehicles = []Vehicle{
	{1, "Toyota", "Corolla", 2023, 250.000},
	{2, "Honda", "Civic", 2023, 230.000},
	{3, "Renault ", "Kwid", 2023, 68.990},
	{4, "Fiat ", "Mobi", 2023, 58.990},
	{5, "Peugeot ", "208", 2023, 62.990},
	{6, "Hyundai ", "HB20", 2023, 82.290},
	{7, "Fiat ", "Argo", 2023, 79.790},
	{8, "Fiat ", "Cronos", 2023, 84.790},
	{9, "Volkswagen ", "Polo", 2023, 74.990},
	{10, "Chevrolet ", "Onix", 2023, 78.390},
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"vehicles": vehicles,
		})
	})

	router.GET("/vehicle/:id", func(c *gin.Context) {
		vehicleID, err := strconv.Atoi(c.Param("id"))
		if err != nil || vehicleID < 1 || vehicleID > len(vehicles) {
			c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
				"message": "Veículo não encontrado",
			})
			return
		}

		c.HTML(http.StatusOK, "vehicle.tmpl", gin.H{
			"vehicle": vehicles[vehicleID-1],
		})
	})

	router.POST("/submit", func(c *gin.Context) {
		vehicleID, err := strconv.Atoi(c.PostForm("vehicle_id"))
		if err != nil || vehicleID < 1 || vehicleID > len(vehicles) {
			c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
				"message": "Veículo não encontrado",
			})
			return
		}

		message := fmt.Sprintf("Muito obrigado pelo seu interesse %s. Entraremos em contato em breve.", vehicles[vehicleID-1].Model)
		c.HTML(http.StatusOK, "confirmation.tmpl", gin.H{
			"message": message,
		})
	})

	router.Static("/static", "./static")

	router.Run(":8080")
}
