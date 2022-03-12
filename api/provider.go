package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostTweet(text, oauth1HeaderContents string) (*PostTweetResponseModel, error) {
	var httpClient *http.Client = &http.Client{}

	const postTweetsURL string = "https://api.twitter.com/2/tweets"

	var body *PostTweetRequestModel = &PostTweetRequestModel{TweetText: text}
	bodyJson, bodyMarshalingError := json.Marshal(*body)
	if bodyMarshalingError != nil {
		return nil, &ApiException{
			StatusCode: http.StatusBadRequest,
			Message:    bodyMarshalingError.Error(),
		}
	}

	request, requestError := http.NewRequest("POST", postTweetsURL, bytes.NewBuffer(bodyJson))
	if requestError != nil {
		return nil, &ApiException{
			StatusCode: http.StatusBadRequest,
			Message:    requestError.Error(),
		}
	}
	request.Header.Add("Authorization", oauth1HeaderContents)
	request.Header.Add("Content-Type", "application/json")

	response, responseError := httpClient.Do(request)
	if responseError != nil {
		return nil, &ApiException{
			StatusCode: http.StatusInternalServerError,
			Message:    responseError.Error(),
		}
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, &ApiException{
			StatusCode: response.StatusCode,
			Message:    response.Status,
		}
	}
	rawResponseBody, bodyParsingError := ioutil.ReadAll(response.Body)
	if bodyParsingError != nil {
		return nil, &ApiException{
			StatusCode: http.StatusInternalServerError,
			Message:    bodyParsingError.Error(),
		}
	}

	var postTweetResponseModel PostTweetResponseModel
	deserializationError := json.Unmarshal(rawResponseBody, &postTweetResponseModel)
	if deserializationError != nil {
		return nil, &ApiException{
			StatusCode: http.StatusInternalServerError,
			Message:    deserializationError.Error(),
		}
	}

	return &postTweetResponseModel, nil
}

func (exception *ApiException) Error() string {
	return fmt.Sprintf("%d: %s", exception.StatusCode, exception.Message)
}
