package handler

import (
	"context"
	"fmt"
	"github.com/weichenluo/Twitter-Raft/server/user/userpb"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func Login(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Getwd()
	root := filepath.Dir(filepath.Dir(f))

	filename := root + "/scripts/index.html"

	if r.Method == "GET" {
		t, _ := template.ParseFiles(filename)
		t.Execute(w, nil)
	} else if r.Method == "POST" {

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		_name := r.FormValue("uname")
		_password := r.FormValue("psw")

		loginDetails := &userpb.LoginDetails{
			Name:     _name,
			Password: _password,
		}

		tmp, err := template.ParseFiles(filename)
		if err != nil {
			fmt.Println("Login parse file failed", err)
			page.Err = "fail to read from web"
			tmp.Execute(w, page)
			return
		}

		user, err := connector.GetUserClient().GetUserByNamePasswrod(context.Background(), loginDetails)
		if err != nil {
			// log.Fatal(err)
			page.Err = "Username or Password is incorrect"
		} else {
			if user.Name != "" {
				page.LoginStatus = true
				page.Name = user.Name
				page.Following = user.Follows

			} else {
				page.Err = "User didn't exsit, please create new account!"
			}

		}

		tmp.Execute(w, page)

	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Getwd()
	root := filepath.Dir(filepath.Dir(f))
	filename := root + "/scripts/index.html"

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles(filename)
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		_name := r.FormValue("uname")
		_password := r.FormValue("psw")

		loginDetails := &userpb.AddUserParameters{
			Name:     _name,
			Password: _password,
		}

		tmp, err := template.ParseFiles(filename)
		if err != nil {
			fmt.Println("SignUp parse file failed", err)
			page.Err = "fail to read from web"
			tmp.Execute(w, page)
			return
		}

		user, err := connector.GetUserClient().Add(context.Background(), loginDetails)
		if err != nil {
			page.Err = "Username has been created, please try other"
		} else {
			page.Name = user.Name
			page.LoginStatus = true
		}

		tmp.Execute(w, page)
	}

}
