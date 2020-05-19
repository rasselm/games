package models

type Comment struct {
	ID           int
	GameID       int
	User         string
	Message      string
	Date_created int64
	Like         int
}

type Game struct {
	ID          int
	Title       string
	Description string
	By          string
	Platform    []string
	Age_rating  int
	Likes       int
	Comments    []Comment

	// Platform    struct[]
}

// {
// 	shortnm: 'b',
// 	longnm:  "b-option",
// 	needArg: false,
// 	help:    "Usage for b",
// }
