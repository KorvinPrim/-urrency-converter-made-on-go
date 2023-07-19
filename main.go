package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (self *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d "+
		"and he has money equal: %d", self.name, self.age, self.money)
}

func (self *User) setNewName(nname string) {
	self.name = nname
}

func home_page(w http.ResponseWriter, r *http.Request) {

	bob := User{"Bob", 25, -50, 4.2, 0.8}
	bob.setNewName("Anton")
	fmt.Fprintf(w, bob.getAllInfo())
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
