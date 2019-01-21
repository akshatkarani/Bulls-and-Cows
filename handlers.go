package main

import (
	"github.com/satori/go.uuid"
	"fmt"
	"strconv"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func easy(res http.ResponseWriter, req *http.Request) {
	// check if cookie exists or else make one
	c, err := req.Cookie("easy")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "easy",
			Value: sID.String(),
		}
	}
	http.SetCookie(res, c)

	// extract the game corresponding to that cookie
	if _, ok := data[c.Value]; !ok {
		data[c.Value] = game{generate_secret_number(), []set{}, false}
	}

	if req.Method == "POST" {
		var newgame game
		if req.FormValue("play_again") == "Play Again" || req.FormValue("restart") == "Restart" {
			newgame = game{generate_secret_number(), []set{}, false}
		} else {
			temp := data[c.Value]
			guess, _ := strconv.Atoi(req.FormValue("guess"))
			newSet := GetSet(guess, temp.secret_number)
	        if newSet.Bulls == 4 {
	            newgame = game{temp.secret_number, append(temp.Sets, newSet), true}    
	        } else {
	            newgame = game{temp.secret_number, append(temp.Sets, newSet), false}
	        }
		}
		data[c.Value] = newgame
	}
	fmt.Println(data[c.Value])
	tpl.ExecuteTemplate(res, "easy.gohtml", data[c.Value])
}