package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type user struct {
	Email    string `bson:"email,omitempty"`
	Name     string `bson:"name,omitempty"`
	Password []byte `bson:"password,omitempty"`
	Age      int    `bson:"age,omitempty"`
	Country  string `bson:"country,omitempty"`
}

var users map[string]user

func signupPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func signinPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "signin.html", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/signup", http.StatusMovedPermanently)
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func userAlreadyExists(email string) bool {
	_, ret := getUser(email)
	return ret
}

func getUser(email string) (user, bool) {
	filter := bson.M{"email": email}
	u := user{}
	err := myDatabase.Collection("users").FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return u, false
	}
	return u, true
}

func goToHomePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/signup", http.StatusMovedPermanently)
}

func userAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	u := user{}
	u.Email = r.FormValue("email")
	u.Name = r.FormValue("name")
	var err error
	u.Password, err = bcrypt.GenerateFromPassword([]byte(r.FormValue("psw")), bcrypt.MinCost)
	checkErr(err)
	u.Age, _ = strconv.Atoi(r.FormValue("age"))
	u.Country = r.FormValue("country")

	if userAlreadyExists(u.Email) {
		tpl.ExecuteTemplate(w, "signup.html", "An account with this email already exists")
		return
	}
	_, err = myDatabase.Collection("users").InsertOne(context.TODO(), u)
	checkErr(err)

	tpl.ExecuteTemplate(w, "signin.html", nil)

}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	ps := r.FormValue("psw")

	user, found := getUser(email)
	if !found {
		tpl.ExecuteTemplate(w, "signin.html", "User not found")
		return
	}

	if bcrypt.CompareHashAndPassword(user.Password, []byte(ps)) != nil {
		tpl.ExecuteTemplate(w, "signin.html", "Incorrect password or username")
		return
	}

	err := tpl.ExecuteTemplate(w, "profile.html", user)
	checkErr(err)

}

var myDatabase *mongo.Database

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseGlob("../templates/*.html")
	checkErr(err)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	checkErr(err)
	err = client.Connect(context.TODO())
	err = client.Ping(context.TODO(), nil)
	checkErr(err)
	myDatabase = client.Database("mydata")
}

func main() {

	http.HandleFunc("/signup", signupPage)
	http.HandleFunc("/useradd", userAdd)
	http.HandleFunc("/signin", signinPage)
	http.HandleFunc("/login", login)
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../templates/css"))))

	http.ListenAndServe(":80", nil)
}
