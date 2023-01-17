package post

// 1 model for whole GetUser flow
type GetPost struct {
	Firstname string `json:"name" bson:"name"`
	Lastname  string `json:"lastname" bson:"lastname"`
}

// 1 model for whole Create user flow
type CreatePost struct {
	Uuid    string
	Title   string `json:"title"`
	Body    string `json:"body"`
	Subject string `json:"subject"`
}

type Post struct {
	Id          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	Views       int    `json:"views"`
	Answers     int    `json:"answers"`
	Votes       int    `json:"votes"`
	Poster      string `json:"poster"`
}
