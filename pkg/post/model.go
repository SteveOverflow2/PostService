package post

// 1 model for whole GetUser flow
type GetPost struct {
	Firstname string `json:"name" bson:"name"`
	Lastname  string `json:"lastname" bson:"lastname"`
}

// 1 model for whole Create user flow
type CreatePost struct {
	Uuid      string `bson:"_id"`
	Firstname string `json:"name" bson:"name"`
	Lastname  string `json:"lastname" bson:"lastname"`
}

type Post struct {
	Id          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
