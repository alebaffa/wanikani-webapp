package structures

type User_Information struct {
	Username      string  `json:"username"`
	Gravatar      string  `json:"gravatar"`
	Level         int8    `json:"level"`
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
	Character         string `json:"character"`
	Meaning           string `json:"meaning"`
	Onyomi            string `json:"onyomi"`
	Kunyomi           string `json:"kunyomi"`
	Important_reading string `json:"important_reading"`
	Level             int8   `json:"level"`
}

type User struct {
	UserInfo      User_Information        `json:"user_information"`
	RequestedInfo []Requested_Information `json:"requested_information"`
}
