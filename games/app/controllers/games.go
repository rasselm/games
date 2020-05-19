package controllers

import (
	"games/app/models"

	"github.com/revel/revel"
)

type Games struct {
	*revel.Controller
}

//CREATE INSTANCES OF COMMENTS USING MODEL STRUCT
var comments = []models.Comment{
	models.Comment{1, 1, "rassel", "Amazing Game", 1351700038, 2},
	models.Comment{2, 1, "thomas", "Superb and mind blowing", 1351700038, 5},
	models.Comment{3, 2, "ryan", "Connot fault it", 1351700038, 3},
	models.Comment{4, 2, "rassel", "Out of the ordinary", 1351700038, 6},
	models.Comment{5, 3, "rassel", "Graphics and content are mind boglling", 1351700038, 4},
	models.Comment{6, 3, "Julian", "Rubbish game waste of money", 1351700038, 3},
	models.Comment{7, 1, "Helen", "Best game So FAR", 1351700038, 10},
}

//CREATE INSTANCES OF GAMES USING MODEL STRUCT
var games = []models.Game{

	models.Game{1, "Uncharted 4", "For the first time ever in Uncharted history, drive vehicles during gameplay", "Sony", []string{"XBOX", "PS4"}, 18, 20, fetchComments(1)},
	models.Game{2, "Call of Duty", "Modern Warefare", "MICROSOFT", []string{"XBOX", "PS4"}, 18, 25, fetchComments(2)},
	models.Game{3, "FIFA 20", "Next Gen Gaming", "SONY", []string{"XBOX", "PS4"}, 18, 30, fetchComments(3)},
}

//###############
// FUNCTION CALLED TO FILTER OUT THE COMMENTS RELEVANT TO THE GAME
func fetchComments(id int) []models.Comment {
	var commentsGameid []models.Comment
	for _, commentSelect := range comments {
		if commentSelect.GameID == id {
			commentsGameid = append(commentsGameid, commentSelect)

			// fmt.Println(key1, commentSelect)
		}
	}
	return commentsGameid
}

//################################

// FUNC FOR RETURNING LIST OF ALL GAMES

func (c Games) List() revel.Result {
	return c.RenderJSON(games)
}

//################## END OF RETURN LIST FUNC ##################

// FUNCTION FOR RETURNING GAME BY ID
func (c Games) Show(gameID int) revel.Result {
	var res models.Game

	for _, game := range games {
		if game.ID == gameID {
			res = game
		}
	}

	if res.ID == 0 {
		return c.NotFound("Could not find game")
	}
	return c.RenderJSON(res)
}

//################## END OF GAME BY ID FUNC ##################

//// FUNCTION FOR RETURNING REPORT
func (c Games) Report() revel.Result {

	//FIND GAME WITH THE HIGHEST LIKES-------------------------------------------------------------------
	var highestlikes int
	var highestkey int
	for key, value := range games {
		if value.Likes > highestlikes {
			highestlikes = value.Likes
			highestkey = key
			// fmt.Println("THIS IS THE ANSWER", highestlikes)
		}
		// fmt.Println(key, value.Likes)
	}
	//----------------------------------------------------------------------------------------------------
	//FIND USER THAT HAS MOST COMMENTS
	var mostcommenteduser = map[string]int{}

	for _, value1 := range comments {
		var exists = mostcommenteduser[value1.User]
		if exists == 0 {
			// fmt.Println("WAS NOT FOUND")
			mostcommenteduser[value1.User] = 1
		} else {
			// fmt.Println("FOUND")
			mostcommenteduser[value1.User] = (mostcommenteduser[value1.User] + 1)
		}
		// fmt.Println("DOESNT EXIST", value1.User, key1, mostcommenteduser, "CHECK THIS", exists)
	}

	var mostcomment int
	var mostcommentkey string
	for key2, value2 := range mostcommenteduser {
		if value2 > mostcomment {
			mostcomment = value2
			mostcommentkey = key2
			// fmt.Println("THIS IS THE ANSWER FOR MOST COMMENTED USER", mostcommentkey)
		}
		// fmt.Println(key2, value2)
	}

	//---------------------------------------------------------------------------------------------------------

	//CREATE A LIST OF MAPS OF AVERAGE LIKES PER GAME

	type M map[string]interface{}
	var myMapSlice []M

	for _, value3 := range comments {
		m1 := M{"gameid": value3.GameID, "likes total": value3.Like, "instances": 1}
		myMapSlice = append(myMapSlice, m1)
	}

	// RESULTS JSON
	var result = M{"highest_rated_game ": games[highestkey].Title, "user_with_most_comments ": mostcommentkey, "average_likes_per_game": myMapSlice}
	return c.RenderJSON(result)
}

//################## END OF REPORT FUNC ##################
