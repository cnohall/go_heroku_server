package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
    "io/ioutil"
	"os"
	"github.com/gin-gonic/gin"

)

type Cat struct {
	Name string `json:"name"`
	Image string `json:"image"`
	CutenessLevel int `json:"cutenessLevel"`
	AllergyInducingFur bool `json:"allergyInducingFur"`
	LivesLeft int `json:"livesLeft"`
}

type Cats struct {
    Cats []Cat `json:"cats"`
}


func homePage(w http.ResponseWriter, r *http.Request) {

	// Open our jsonFile
	jsonFile, err := os.Open("catdata.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var cats Cats

	json.Unmarshal(byteValue, &cats)

	json.NewEncoder(w).Encode(cats)
}

// func handleRequests() {
// 	http.HandleFunc("/", homePage)
// 	log.Fatal(http.ListenAndServe(":8081",nil))
// }

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, homePage, nil)
	})

	router.Run(":" + port)

}

