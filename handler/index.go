package handler

import (
	"context"
	"fmt"
	"github.com/weichenluo/Twitter-Raft/server/post/postpb"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Getwd()
	root := filepath.Dir(filepath.Dir(f))

	filename := root + "/scripts/index.html"

	tmp, err := template.ParseFiles(filename)

	if err != nil {
		fmt.Println("Index parse file failed", err)
		return
	}

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

	tmp.Execute(w, page)
}
