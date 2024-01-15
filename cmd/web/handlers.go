package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
	"io/ioutil"
	"encoding/json"
)

type Person struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Section string `json:"section"`
	Birthdate string `json:"birthdate"`
	Presentation string `json:"presentation"`
	Topic string `json:"topic"`
}

type ParticipantsData struct {
    Participants []Person
}

var person Person

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("ui/html/index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func send(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

	//var person Person
    err = json.Unmarshal(body, &person)
    if err != nil {
		fmt.Println(err)
        http.Error(w, "Ошибка распаковки JSON", http.StatusBadRequest)
        return
    }

	err = addPerson(person)
	if err != nil {
		log.Fatal(err)
	}
}

func answers(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("ui/html/answers.html")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, person); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func participants(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var participants []Person

    for rows.Next() {
        var person Person
		var id int

        err = rows.Scan(&id, &person.Name, &person.Surname, &person.Patronymic, 
            &person.Phone, &person.Email, &person.Section, &person.Birthdate, &person.Presentation, &person.Topic)
        if err != nil {
            log.Fatal(err)
        }

        participants = append(participants, person)
    }

    data := ParticipantsData{
        Participants: participants,
    }

    tmpl, err := template.ParseFiles("ui/html/participants.html")
    if err != nil {
        log.Fatal(err)
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        log.Fatal(err)
    }
}

func addPerson(person Person) error {
    insertSQL := `
        INSERT INTO users (name, surname, patronymic, phone, email, section, birthdate, presentation, topic)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

    _, err := db.Exec(insertSQL, person.Name, 
		person.Surname, person.Patronymic, person.Phone, person.Email, 
		person.Section, person.Birthdate, person.Presentation, person.Topic)
    if err != nil {
        return err
    }

    return nil
}