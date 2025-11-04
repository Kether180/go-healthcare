package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HealthRecord represents a health record with basic information
type HealthRecord struct {
	ID          int    `json:"id"`
	PatientName string `json:"patient_name" binding:"required"`
	Age         int    `json:"age" binding:"required,min=0,max=150"`
	Diagnosis   string `json:"diagnosis" binding:"required"`
	Treatment   string `json:"treatment"`
}

// In-memory storage for health records
var healthRecords = []HealthRecord{
	{ID: 1, PatientName: "John Doe", Age: 35, Diagnosis: "Flu", Treatment: "Rest and fluids"},
	{ID: 2, PatientName: "Jane Smith", Age: 28, Diagnosis: "Sprained Ankle", Treatment: "Ice and elevation"},
}

var nextID = 3

func main() {
	router := gin.Default()

	// Define API routes
	api := router.Group("/api/v1")
	{
		api.GET("/health", getHealthStatus)
		api.GET("/records", getAllRecords)
		api.GET("/records/:id", getRecordByID)
		api.POST("/records", createRecord)
		api.PUT("/records/:id", updateRecord)
		api.DELETE("/records/:id", deleteRecord)
	}

	// Start server
	router.Run(":8080")
}

// getHealthStatus returns the health status of the API
func getHealthStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "Go Health API is running",
	})
}

// getAllRecords returns all health records
func getAllRecords(c *gin.Context) {
	c.JSON(http.StatusOK, healthRecords)
}

// getRecordByID returns a specific health record by ID
func getRecordByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	for _, record := range healthRecords {
		if record.ID == id {
			c.JSON(http.StatusOK, record)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
}

// createRecord creates a new health record
func createRecord(c *gin.Context) {
	var newRecord HealthRecord

	if err := c.ShouldBindJSON(&newRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRecord.ID = nextID
	nextID++
	healthRecords = append(healthRecords, newRecord)

	c.JSON(http.StatusCreated, newRecord)
}

// updateRecord updates an existing health record
func updateRecord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedRecord HealthRecord
	if err := c.ShouldBindJSON(&updatedRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, record := range healthRecords {
		if record.ID == id {
			updatedRecord.ID = id
			healthRecords[i] = updatedRecord
			c.JSON(http.StatusOK, updatedRecord)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
}

// deleteRecord deletes a health record by ID
func deleteRecord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	for i, record := range healthRecords {
		if record.ID == id {
			healthRecords = append(healthRecords[:i], healthRecords[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
}
