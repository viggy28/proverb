package proverb

import (
	"log"
	"math/rand"
	"time"
)

var proverbs []string

func Vanakam() {
	log.Println("Vanakkam 🙏")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func Pazamozhi() string {
	proverbs = []string{"Aathula oru kaal, sethula oru kaal !!", "nalla maatuku oru suudu", "puliki pasichalum pulla thingadu"}
	return proverbs[rand.Intn(len(proverbs))]
}

func PazamozhiEng() string {
	return "One leg in water and another one in shore !!"
}
