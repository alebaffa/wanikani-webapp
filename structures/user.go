package structures

//UserInformation contains the information of the user
type UserInformation struct {
	Username     string  `json:"username"`
	Gravatar     string  `json:"gravatar"`
	Level        int8    `json:"level"`
	Title        string  `json:"title"`
	About        string  `json:"about"`
	Website      string  `json:"website"`
	Twitter      string  `json:"twitter"`
	TopicsCount  float64 `json:"topics_count"`
	PostsCount   float64 `json:"posts_count"`
	CreationDate float64 `json:"creation_date"`
	VacationDate float64 `json:"vacation_date"`
}

//RequestedInformation contains the information requested from the user
type RequestedInformation struct {
	Character        string `json:"character"`
	Meaning          string `json:"meaning"`
	Onyomi           string `json:"onyomi"`
	Kunyomi          string `json:"kunyomi"`
	ImportantReading string `json:"important_reading"`
	Level            int8   `json:"level"`
}

//User contains both UserInformation and RequestedInfo
type User struct {
	UserInfo      UserInformation        `json:"user_information"`
	RequestedInfo []RequestedInformation `json:"requested_information"`
}
