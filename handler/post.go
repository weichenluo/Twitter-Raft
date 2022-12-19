package handler

import (
	"context"
	"fmt"
	"github.com/weichenluo/Twitter-Raft/server/post/postpb"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
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

	if r.URL.Path != "/post" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	title := r.FormValue("ptitle")
	content := r.FormValue("pcontent")

	pp := &postpb.PostText{
		User:  page.Name,
		Title: title,
		Body:  content,
	}
	tmp, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("Index parse file failed", err)
		return
	}

	_, err = connector.GetPostClient().AddPost(context.Background(), pp)

	if err != nil {
		page.Err = "failed to post Feed, please try again"
	} else {
		time.Sleep(1 * time.Second)

		post, _ := connector.GetPostClient().GetAllPosts(context.Background(), &postpb.NoArgs{})

		currFeeds := make([]Feed, 0)

		if post != nil {
			for _, p := range post.Posts {
				newFeed := Feed{
					User:  p.User,
					Title: p.Title,
					Body:  p.Body,
					Time:  p.Time,
				}
				currFeeds = append(currFeeds, newFeed)
			}
		}

		page.Feeds = currFeeds
	}

	tmp.Execute(w, page)
}
