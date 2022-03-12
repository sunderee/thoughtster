package api

type ApiException struct {
	StatusCode int
	Message    string
}

type Oauth1Model struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	TokenSecret    string
}

type PostTweetRequestModel struct {
	TweetText string `json:"text"`
}

type PostTweetResponseModel struct {
	Data PostTweetResponseDataModel `json:"data"`
}

type PostTweetResponseDataModel struct {
	TweetID   string `json:"id"`
	TweetText string `json:"text"`
}
