package handler

import (
	"app/db"
	"app/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() {
	database, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	database.AutoMigrate(&model.Anime{})
	db.DB = database
}

func TestCreateAnime(t *testing.T) {
	SetupTestDB()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/animes", CreateAnime)
	animeTest := model.Anime{
		Title:  "Frieren",
		Genre:  "Fantasy",
		Year:   2023,
		Author: "Kanehito Yamada",
	}
	jsonValue, _ := json.Marshal(animeTest)
	req, _ := http.NewRequest("POST", "/animes", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestGetAnimes(t *testing.T) {
	SetupTestDB()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/animes", GetAnimes)
	animeTest := model.Anime{
		Title:  "Bocchi the Rock!",
		Genre:  "Slice of Life",
		Year:   2022,
		Author: "Aki Hamazi",
	}
	db.GetDB().Create(&animeTest)
	req, _ := http.NewRequest("GET", "/animes", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetAnimeByID(t *testing.T) {
	SetupTestDB()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/animes/:id", GetAnimeByID)
	animeTest := model.Anime{
		Title:  "Spy x Family",
		Genre:  "Action",
		Year:   2022,
		Author: "Tatsuya Endo",
	}
	db.GetDB().Create(&animeTest)
	req, _ := http.NewRequest("GET", "/animes/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestEditAnime(t *testing.T) {
	SetupTestDB()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/animes/:id", EditAnime)
	animeTest := model.Anime{
		Title:  "Spy x Family",
		Genre:  "Action",
		Year:   2022,
		Author: "Tatsuya Endo",
	}
	db.GetDB().Create(&animeTest)
	targetID := strconv.Itoa(int(animeTest.ID))
	t.Run("Success Update", func(t *testing.T) {
		//dados Novos
		updateAnime := model.Anime{
			Title:  "Bocchi the Rock!",
			Genre:  "Slice of Life",
			Year:   2022,
			Author: "Aki Hamazi",
		}
		jsonValue, _ := json.Marshal(updateAnime)
		req, _ := http.NewRequest("PUT", "/animes/"+targetID, bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		var animeInDB model.Anime
		db.GetDB().First(&animeInDB, animeTest.ID)
		assert.Equal(t, updateAnime.Title, animeInDB.Title)
		assert.Equal(t, updateAnime.Genre, animeInDB.Genre)
		assert.Equal(t, updateAnime.Year, animeInDB.Year)
		assert.Equal(t, updateAnime.Author, animeInDB.Author)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/animes/invalid-id", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid anime ID")
	})
}

func TestDeleteAnime(t *testing.T) {
	SetupTestDB()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/animes/:id", DeleteAnime)
	animeTest := model.Anime{
		Title:  "Spy x Family",
		Genre:  "Action",
		Year:   2022,
		Author: "Tatsuya Endo",
	}
	db.GetDB().Create(&animeTest)
	targetID := strconv.Itoa(int(animeTest.ID))
	req, _ := http.NewRequest("DELETE", "/animes/"+targetID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	// Verificar se o anime foi realmente deletado
	var animeInDB model.Anime
	err := db.GetDB().First(&animeInDB, animeTest.ID).Error
	assert.Error(t, err)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
}
