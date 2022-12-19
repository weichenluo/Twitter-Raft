package post_driver

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/weichenluo/Twitter-Raft/server/post/postpb"
	"github.com/weichenluo/Twitter-Raft/util"
	"log"
	"time"
)

type Server struct {
	postpb.UnimplementedPostServiceServer
}

func Init() {
	var post_ postpb.UserPosts
	postDB, err := GetPostDB(post_)
	if err != nil {
		log.Println("Error occured while storing post data in Raft =", err)
		panic(err)
	}
	log.Println("DB Post Initialized =", postDB.Posts)
}

func GetPostDB(value interface{}) (postpb.UserPosts, error) {
	var db postpb.UserPosts
	data, err := util.Raftstorage("GET", "postDB", db)
	if err != nil {
		log.Println("Error occured while getting post data from Raft =", err)
		panic(err)
	}
	log.Println("data received after interactwithraftstorage =", data)
	var postDB postpb.UserPosts
	postDB, err = DecodeRaftPost(data)
	if err != nil {
		log.Println("Error occured while decoding post data from Raft storage =", err)
		return postDB, err
	}
	log.Println("postDB after decode =", postDB)
	return postDB, nil
}

func DecodeRaftPost(db string) (postpb.UserPosts, error) {
	var post_ postpb.UserPosts
	log.Println("Decode post Storage called")
	dec := gob.NewDecoder(bytes.NewBufferString(db))
	if err := dec.Decode(&post_); err != nil {
		log.Fatalf("raftexample: could not decode message (%v)", err)
		return post_, err
	}
	//log.Println("postDB in DecodeRaftpostStorage =", up)

	return post_, nil
}

func (*Server) AddPost(ctx context.Context, postDetails *postpb.PostText) (*postpb.Post, error) {
	var post_ postpb.UserPosts
	// fmt.Print(666)
	log.Println("AddPost API called")
	// var up postpb.UserPosts
	post := &postpb.Post{
		User:  postDetails.User,
		Title: postDetails.Title,
		Body:  postDetails.Body,
		Time:  getTime(),
	}

	postDB, err := GetPostDB(post_)
	if err != nil {
		return nil, err
	}

	postDB.Posts = append(postDB.Posts, post)

	log.Println("New Post =", post)
	log.Println("PostsDB = ", postDB.Posts)

	_, err = util.Raftstorage("PUT", "postDB", postDB)
	if err != nil {
		log.Println("Error occured while storing post data in Raft =", err)
		panic(err)
	}
	return post, nil
}

// func (*Server) GetFollowerPosts(ctx context.Context, users *postpb.Users) (*postpb.UserPosts, error) {
// 	log.Println("GetFollowerPosts called")
// 	var up postpb.UserPosts
// 	posts := &postpb.UserPosts{
// 		Posts: make([]*postpb.Post, 0),
// 	}
// 	postDB, err := GetPostDB(up)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, user := range users.Ids {
// 		for _, userPostsObj := range postDB.Posts {
// 			if user == userPostsObj.UserId {
// 				posts.Posts = append(posts.Posts, userPostsObj)
// 			}
// 		}
// 	}

// _, err = util.InteractWithRaftStorage("PUT", "postDB", postDB)
// if err != nil {
// 	log.Println("Error occured while storing post data in Raft =", err)
// 	panic(err)
// }

// return posts, nil
// }

func (*Server) GetAllPosts(ctx context.Context, in *postpb.NoArgs) (*postpb.UserPosts, error) {
	var post_ postpb.UserPosts
	log.Println("GetAllPosts called")

	postDB, err := GetPostDB(post_)
	if err != nil {
		return nil, err
	}

	log.Println("up  in GetAllPosts =", postDB)
	return &postDB, nil
}

func getTime() string {
	now := time.Now()
	dt := now.Format(time.RFC822)
	return "Post date and time is: " + dt
}
