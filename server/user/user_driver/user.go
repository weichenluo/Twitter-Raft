package user_driver

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"github.com/weichenluo/Twitter-Raft/server/user/userpb"
	"github.com/weichenluo/Twitter-Raft/util"
	"golang.org/x/exp/slices"
	"log"
)

type Server struct {
	userpb.UnimplementedUserServiceServer
}

func Init() {
	var login_ userpb.Login

	userDB, err := GetUserDB(login_)

	if err != nil {
		// log.Println("Error occured while storing post data in Raft =", err)
		panic(err)
	}

	log.Println("DB User Initialized =", userDB.Users)
}

func GetUserDB(value interface{}) (userpb.Login, error) {
	var userDB userpb.Login
	var db userpb.Login
	data, err := util.Raftstorage("GET", "userDB", db)
	if err != nil {
		log.Println("Error occured while getting user data from Raft =", err)
		panic(err)
	}

	//var userDB userpb.Login
	userDB, err = DecodeRaft(data)
	if err != nil {
		log.Println("Error occured while decoding user data from Raft storage =", err)
		return userDB, err
	}

	log.Println("userDB after decode =", userDB)

	return userDB, nil
}

func DecodeRaft(db string) (userpb.Login, error) {
	var login_ userpb.Login
	log.Println("Decode User Storage called")
	dec := gob.NewDecoder(bytes.NewBufferString(db))
	if err := dec.Decode(&login_); err != nil {
		log.Fatalf("raftexample: could not decode message (%v)", err)
		return login_, err
	}
	//log.Println("userDB in DecodeRaftTokenStorage =", lo)

	return login_, nil
}

func (*Server) Add(ctx context.Context, up *userpb.AddUserParameters) (*userpb.User, error) {
	var login_ userpb.Login

	user := &userpb.User{
		Name:     up.Name,
		Password: up.Password,
		Follows:  make([]string, 0),
	}

	userDB, err := GetUserDB(login_)
	if err != nil {
		return nil, err
	}

	for _, value := range userDB.Users {
		if value.Name == user.Name {
			return nil, errors.New("User already registered!!")
		}
	}

	userDB.Users = append(userDB.Users, user)

	_, err = util.Raftstorage("PUT", "userDB", userDB)
	if err != nil {
		log.Println("Error occured while storing post data in Raft =", err)
		panic(err)
	}

	log.Println("User added =", user)
	log.Println("User DB =", userDB.Users)
	return user, nil
}

func (*Server) GetUserByNamePasswrod(ctx context.Context, user *userpb.LoginDetails) (*userpb.User, error) {
	var login_ userpb.Login

	userDB, err := GetUserDB(login_)
	if err != nil {
		return nil, err
	}

	var _user userpb.User
	for i, v := range userDB.Users {
		if v.Name == user.Name && v.Password == user.Password {
			userDB.Users[i].LoginStatus = true
			_user = *userDB.Users[i]
			break
		}
	}

	return &_user, nil
}

func (*Server) FollowUser(ctx context.Context, fp *userpb.FollowerParameters) (*userpb.Status, error) {
	var login_ userpb.Login

	userDB, err := GetUserDB(login_)
	if err != nil {
		return nil, err
	}

	hasUser := false
	hasFollowing := false

	_user := &userpb.User{}

	for _, v := range userDB.Users {
		if v.Name == fp.Following {
			hasFollowing = true
			break
		}
	}
	if hasFollowing {
		for i, v := range userDB.Users {
			if v.Name == fp.Follower {
				if !slices.Contains(userDB.Users[i].Follows, fp.Following) {
					userDB.Users[i].Follows = append(userDB.Users[i].Follows, fp.Following)
				}
				_user = userDB.Users[i]
				hasUser = true
				break
			}
		}
	}

	_, err = util.Raftstorage("PUT", "userDB", userDB)
	if err != nil {
		log.Println("Error occured while storing post data in Raft =", err)
		panic(err)
	}
	// log.Println("Raft content: ", content)

	if hasUser {
		log.Println("User ", fp.Follower, " follows ", fp.Following)
		log.Println("User DB =", userDB.Users)

		res := &userpb.Status{User: _user, ResponseStatus: true}
		return res, nil
	} else {
		log.Println("cannot find the user")
		res := &userpb.Status{User: _user, ResponseStatus: false}
		return res, nil
	}

}

func (*Server) GetUserFollowerByName(ctx context.Context, user *userpb.UserName) (*userpb.UserList, error) {

	return nil, nil
}

func (*Server) GetUserFollowingByName(ctx context.Context, user *userpb.UserName) (*userpb.Login, error) {
	var login_ userpb.Login

	userDB, err := GetUserDB(login_)
	if err != nil {
		return nil, err
	}

	res := &userpb.Login{
		Users: make([]*userpb.User, 0),
	}

	var _user *userpb.User

	for i, v := range userDB.Users {
		if v.Name == user.Name {
			_user = userDB.Users[i]
		}
	}

	for _, v := range _user.Follows {
		for _, u := range userDB.Users {
			if v == user.Name {
				res.Users = append(res.Users, u)
				break
			}
		}
	}

	return res, nil
}
