package requestdata

//type of data of the API request
type RequestData struct {
	// App       string `json:"app"`
	//Timestamp int `json:"timestamp"` doesn't work
	//Version   int8  `json:"version"`
	// Type      string `json:"type"`
	Payload struct {
		//ID      string `json:"id"`
		Source string `json:"source" binding:"required"`
		//Type    string `json:"type"`
		Payload struct {
			Text string `json:"text" binding:"required"`
		} `json:"payload" binding:"required"`
	} `json:"payload" binding:"required"`
	// Sender struct {
	// 	Phone       string `json:"phone"`
	// 	Name        string `json:"name"`
	// 	CountryCode string `json:"country_code"`
	// 	DialCode    string `json:"dial_code"`
	// } `json:"sender"`
	// Context struct {
	// 	ID   string `json:"id"`
	// 	GSID string `json:"gsId"`
	// } `json:"context"`
}

type LocationData struct {
	Payload struct {
		//ID      string `json:"id"`
		Source string `json:"source" binding:"required"`
		//Type    string `json:"type"`
		Location struct {
			Latitude  float32 `json:"lat" binding:"required"`
			Longitude float32 `json:"long" binding:"required"`
		} `json:"location" binding:"required"`
	} `json:"payload" binding:"required"`
}

//type of data of the sending message
type SentData struct {
	Message string `json:"message"`
}

type SentLocData struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"long"`
}
