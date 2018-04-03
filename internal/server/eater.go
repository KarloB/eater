package server

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/karlob/eater/internal/model"
)

// Food where do we eat
type Food struct {
	Label string
	Image string
}

var foods = []Food{
	{Label: "Lidl", Image: "https://upload.wikimedia.org/wikipedia/commons/thumb/9/91/Lidl-Logo.svg/220px-Lidl-Logo.svg.png"},
	{Label: "OhYeah!", Image: "https://i.imgur.com/gsXg8ok.jpg"},
	{Label: "Kaufland", Image: "https://upload.wikimedia.org/wikipedia/commons/thumb/4/40/Kaufland_2016_horizontal.svg/800px-Kaufland_2016_horizontal.svg.png"},
	{Label: "Tramontana", Image: "https://i.imgur.com/H8k1nMC.jpg"},
	{Label: "Palačinkice", Image: "https://i.imgur.com/Og5GsWw.png"},
	{Label: "Titanik", Image: "https://upload.wikimedia.org/wikipedia/commons/thumb/3/3c/RMS_Titanic_sea_trials_April_2%2C_1912.jpg/250px-RMS_Titanic_sea_trials_April_2%2C_1912.jpg"},
	{Label: "Servus", Image: "https://i.imgur.com/xJWNHy4.jpg"},
	{Label: "Zdravljak", Image: "https://i.imgur.com/z7DRhDE.jpg"},
}

// Eater http handler
func (t *Server) Eater(w http.ResponseWriter, r *http.Request) {

	log.Printf("r.UserAgent() %s\n", r.UserAgent())

	chosenOne := rand.Intn(len(foods))
	if time.Now().Weekday().String() == "Thursday" {
		chosenOne = 0
	}

	tpl := template.Must(template.New("index").Parse(model.Tpl))
	err := tpl.Execute(w, foods[chosenOne])
	if err != nil {
		log.Printf("Err: %v", err)
		return
	}
}
