package gate

import (
	"net/http"

	"github.com/ahnafms/go-rest-gate/models"
	"github.com/gin-gonic/gin"
)

type Kartu_Akses struct {
	ID_kartu_akses string `gorm:"primaryKey" json:"id_kartu_akses"`
}

func InGate(c *gin.Context) {
	// Create a new instance of the Kartu_Akses struct
	var kartu_akses Kartu_Akses

	// Bind the JSON request body to the struct instance
	if err := c.ShouldBindJSON(&kartu_akses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result int
	if err := models.DB.Raw("SELECT is_aktif FROM kartu_akses WHERE id_kartu_akses = ?", kartu_akses.ID_kartu_akses).Scan(&result).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if result == 1 {
		c.JSON(http.StatusOK, gin.H{"message": "Kartu sedang aktif"})
		return
	}
	models.DB.Exec("UPDATE kartu_akses SET is_aktif = 1 WHERE id_kartu_akses = ?", kartu_akses.ID_kartu_akses)
	c.JSON(http.StatusOK, gin.H{"message": "Berhasil memasuki gerbang"})
}

func OutGate(c *gin.Context) {
	var kartu_akses Kartu_Akses

	// Bind the JSON request body to the struct instance
	if err := c.ShouldBindJSON(&kartu_akses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result int
	if err := models.DB.Raw("SELECT is_aktif FROM kartu_akses WHERE id_kartu_akses = ?", kartu_akses.ID_kartu_akses).Scan(&result).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if result == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Kartu sedang tidak aktif"})
		return
	}

	models.DB.Exec("UPDATE kartu_akses SET is_aktif = 0 WHERE id_kartu_akses = ?", kartu_akses.ID_kartu_akses)
	c.JSON(http.StatusOK, gin.H{"message": "Kartu berhasil dinonaktifkan"})

}
