package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)
// *** TODO dllfwelflflfwelfwefwe;
type User struct {
	Name string
	Age uint16
	Money int16
	Avg_grades, Happiness float64
	Hobbies []string
	
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money " +
		"equal: %d ", u.Name, u.Age, u.Money )
}

func (u *User) setNewName (newName string) {
	u.Name=newName
}


func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8,  []string{"Football", "Hockey", "Yoga"}}
	//	bob.name = "Alex"
	bob.setNewName ("Vasya")
	//	fmt.Fprintf(w, bob.getAllInfo())
	//fmt.Fprintf(w, "<b>Main page</b>")
	//fmt.Fprintf(w, `<h1>Main page</h1>
	//	<b>Main text</b>`)
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page!")
}

func handleRequest () {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		//	log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":" + port, nil)

}
func main () {
	//var bob User =...
	//bob := User { name: "Bob", age: 25, money := -50, avg_grades: 4.2, happiness: 0.8}
	
	
	handleRequest()

	
}
