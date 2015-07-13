package structures

type User_Information struct {
	Username      string  `json:"username"`
	Gravatar      string  `json:"gravatar"`
	Level         float64 `json:"level"`
	Title         string  `json:"title"`
	About         string  `json:"about"`
	Website       string  `json:"website"`
	Twitter       string  `json:"twitter"`
	Topics_count  float64 `json:"topics_count"`
	Posts_count   float64 `json:"posts_count"`
	Creation_date float64 `json:"creation_date"`
	Vacation_date float64 `json:"vacation_date"`
}

type Requested_Information struct {
	LessonsAvailable         int     `json:"lessons_available"`
	ReviewAvailable          int     `json:"review_available"`
	NextReviewDate           float64 `json:"next_review_available"`
	ReviewsAvailableNextHour int     `json:"reviews_available_next_hour"`
	ReviewsAvailableNextDay  int     `json:"reviews_available_next_day"`
}

type User struct {
	UserInfo      User_Information      `json:"user_information"`
	RequestedInfo Requested_Information `json:"requested_information"`
}
