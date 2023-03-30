package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var valeur1, valeur2 string
var match bool = false
var flips, errs int

func main() {
	// установка обработчика для статических файлов
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", playRound)
	http.HandleFunc("/url", handleUrl)
	http.ListenAndServe(":8080", nil)
}
func playRound(w http.ResponseWriter, r *http.Request) {
	buttonsHTML := generateButtonsHTML()
	tpl, _ := template.ParseFiles("index.html")
	tpl.Execute(w, buttonsHTML)

}

var rows = [][]string{
	{"img16.jpg", "img14.jpg", "img4.jpg"},
	{"img8.jpg", "img9.jpg", "img10.jpg"},
	{"img16.jpg", "img14.jpg", "img4.jpg"},
	{"img8.jpg", "img9.jpg", "img10.jpg"},
}

func generateButtonsHTML() template.HTML {

	shuffledRows := shuffle(rows)
	html := "<table id='cards-table'>\n"
	for _, row := range shuffledRows {
		html += "<tr>"
		path := `"static/img/`
		for _, col := range row {
			html += fmt.Sprintf("<td><button class='card rotate' style='background-image: url("+path+"%s'); background-size: cover; background-repeat: no-repeat;'><img src='static/img/img3.jpg' class='img'/></button></td>", col)

		}
		html += "</tr>\n"
	}
	html += "</table>"
	html += "<button  class='btn btn-warning' id='reloadButton' onclick='location.reload()'>Start!</button>"

	htmlTemplate := template.HTML(html)
	return htmlTemplate
}

func handleUrl(w http.ResponseWriter, r *http.Request) {
	// Получаем URL кнопки из тела запроса
	decoder := json.NewDecoder(r.Body)
	var data struct {
		ButtonUrl string `json:"buttonUrl"`
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	// Выполняем нужные действия с URL кнопки
	fmt.Println("URL кнопки:", data.ButtonUrl)

	//fmt.Println("Flip:", data.Flip)

	if valeur1 == "" {
		valeur1 = data.ButtonUrl
	} else {
		valeur2 = data.ButtonUrl
		match = checkMatch(valeur1, valeur2)
		hidden(match, flips, errs, w)
		valeur1 = ""

	}

	//fmt.Println("valeur1", valeur1)
	//fmt.Println("valeur2", valeur2)

}

func hidden(match bool, flips int, errs int, w http.ResponseWriter) {
	if match {
		// добавляем класс "hidden" к элементу img, чтобы скрыть его
		//fmt.Println(match)
		w.Header().Set("Content-Type", "application/json")
		data := map[string]interface{}{"match": match, "flips": flips, "errs": errs}
		json.NewEncoder(w).Encode(data)
	} else {
		// скрываем элемент img с помощью класса "hidden"
		//fmt.Println(match)
		w.Header().Set("Content-Type", "application/json")
		data := map[string]interface{}{"match": match, "flips": flips, "errs": errs}
		json.NewEncoder(w).Encode(data)
	}
}

func shuffle(rows [][]string) [][]string {
	// Convert the 2D array to a 1D slice.
	elements := make([]string, 0, len(rows)*len(rows[0]))
	for _, row := range rows {
		elements = append(elements, row[:]...)
	}

	// Shuffle the elements.
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(elements), func(i, j int) {
		elements[i], elements[j] = elements[j], elements[i]
	})

	// Convert the 1D slice back to a 2D array.
	shuffledRows := make([][]string, len(rows))
	for i := range shuffledRows {
		shuffledRows[i] = make([]string, len(rows[i]))
		for j := range shuffledRows[i] {
			shuffledRows[i][j] = elements[i*len(rows[i])+j]
		}
	}
	flips = 0
	errs = 0
	return shuffledRows
}

func checkMatch(valeur1 string, valeur2 string) bool {
	if valeur1 != valeur2 {
		flips += 1
		errs += 1
		fmt.Println("flips : ", flips)
		fmt.Println("errors : ", errs)
		return false
	} else {
		flips += 1
		fmt.Println("flips : ", flips)
		return true
	}
}
