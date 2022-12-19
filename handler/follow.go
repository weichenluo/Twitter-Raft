package handler

import (
	"context"
	"fmt"
	"github.com/weichenluo/Twitter-Raft/server/user/userpb"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Follow(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Getwd()
	root := filepath.Dir(filepath.Dir(f))
	filename := root + "/scripts/index.html"
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if !page.LoginStatus {
		fmt.Printf("Please log in to create a post\n")
		return
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles(filename)
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		following_id := r.FormValue("ftitle")

		fp := &userpb.FollowerParameters{
			Follower:  page.Name,
			Following: following_id,
		}

		tmp, err := template.ParseFiles(filename)
		if err != nil {
			fmt.Println("Follow parse file failed", err)
			return
		}

		status, err := connector.GetUserClient().FollowUser(context.Background(), fp)
		if err != nil {
			log.Fatal(err)
			page.Err = "Some wrong with following"
			tmp.Execute(w, page)
			return
		} else {
			if status.ResponseStatus {
				page.Following = status.User.Follows
			} else {
				page.Err = "User doesn't exist!"
			}
		}

		tmp.Execute(w, page)
	}
}
