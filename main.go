package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, happiness float64
	Hobbies               []string
}

func (self *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d "+
		"and he has money equal: %d", self.Name, self.Age, self.Money)
}

func (self *User) setNewName(nname string) {
	self.Name = nname
}

func home_page(w http.ResponseWriter, r *http.Request) {

	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Fottball", "Skate", "Dance"}}
	//bob.setNewName("Anton")
	//fmt.Fprintf(w, bob.getAllInfo())

	tmpl, _ := template.ParseFiles("templates/home_page.html")

	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)

	http.HandleFunc("/contacts/", contacts_page)

	http.ListenAndServe(":8080", nil)
}

func main() {

	handleRequest()
}
