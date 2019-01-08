package main

import (
    "github.com/satori/go.uuid"
    "net/http"
    "text/template"
    "strconv"
    "fmt"
)

type set struct {
    Guess, Bulls, Cows int
}

type game struct {
    secret_number int
    sets          []set
}

var tpl *template.Template
var data = map[string]game{}

func init() {
    tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func index(res http.ResponseWriter, req *http.Request) {
    tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func easy(res http.ResponseWriter, req *http.Request) {
    fmt.Println(req.Method)
    if req.Method =="POST" {
        fmt.Println("enter")
    }

    c, err := req.Cookie("easy")
    if err != nil {
        sID, _ := uuid.NewV4()
        c = &http.Cookie{
            Name:  "easy",
            Value: sID.String(),
        }
    }
    http.SetCookie(res, c)

    if _, ok := data[c.Value]; !ok {
        data[c.Value] = game{5467, []set{}}
    }

    if req.Method == "POST" {
        //fmt.Println
        temp := data[c.Value]
        guess, _ := strconv.Atoi(req.FormValue("guess"))
        newSet := GetSet(guess, temp.secret_number)
        newgame := game{temp.secret_number, append(temp.sets, newSet)}
        data[c.Value] = newgame
    }
    fmt.Println(data[c.Value])
    tpl.ExecuteTemplate(res, "easy.gohtml", data[c.Value].sets)
}

func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/easy/", easy)
    http.Handle("/favicon.ico", http.NotFoundHandler())
    http.ListenAndServe(":9000", nil)
}
