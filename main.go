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
	{1, "Toyota", "Corolla", 2023, 25000},
	{2, "Honda", "Civic", 2023, 23000},
	// Adicione mais veículos aqui
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

		// Simulação de confirmação de envio
		message := fmt.Sprintf("Lead gerado para o veículo %s. Entraremos em contato em breve.", vehicles[vehicleID-1].Model)
		c.HTML(http.StatusOK, "confirmation.tmpl", gin.H{
			"message": message,
		})
	})

	router.Run(":8080")
}
