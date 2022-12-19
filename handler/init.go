package handler

type Feed struct {
	User  string
	Title string
	Body  string
	Time  string
}

type Page struct {
	Err         string
	LoginStatus bool
	Feeds       []Feed
	Name        string
	Follower    []string
	Following   []string
}

func init() {
	InitConnectors()
}
