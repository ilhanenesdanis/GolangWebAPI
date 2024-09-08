package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"

	model "WebApp6/Models"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	page := model.Page{ID: 3, Name: "Kullanicilar", Description: "Kullanici listesi", URI: "/users"}

	users := loadUsers()

	interests := loadInterests()

	interestMappings := loadInterestsMapping()

	var newUsers []model.Users

	for _, user := range users {
		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interest = append(user.Interest, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}

	viewModel := model.UserVM{Page: page, Users: newUsers}

	t, _ := template.ParseFiles("Template/page.html")

	t.Execute(w, viewModel)
}
func loadFile(fileName string) (string, error) {

	btytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(btytes), nil
}

func loadUsers() []model.Users {
	bytes, err := ioutil.ReadFile("json/users.json")

	if err != nil {
		return nil
	}

	var users []model.Users

	json.Unmarshal(bytes, &users)

	return users
}

func loadInterests() []model.Interest {
	bytes, err := ioutil.ReadFile("json/interest.json")

	if err != nil {
		return nil
	}

	var interests []model.Interest

	json.Unmarshal(bytes, &interests)

	return interests
}
func loadInterestsMapping() []model.InterestMapping {
	bytes, err := ioutil.ReadFile("json/userInterestMappings.json")

	if err != nil {
		return nil
	}

	var interests []model.InterestMapping

	json.Unmarshal(bytes, &interests)

	return interests
}
