package api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"hash"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (oauth1Model *Oauth1Model) BuildOauth1AuthenticationHeader(method, path string) string {
	var urlValues url.Values = url.Values{}

	urlValues.Add("oauth_signature_method", "HMAC-SHA1")
	urlValues.Add("oauth_version", "1.0")
	urlValues.Add("oauth_nonce", generateNonce())
	urlValues.Add("oauth_consumer_key", oauth1Model.ConsumerKey)
	urlValues.Add("oauth_token", oauth1Model.AccessToken)
	urlValues.Add("oauth_timestamp", strconv.Itoa(int(time.Now().Unix())))

	var parameterString string = strings.Replace(urlValues.Encode(), "+", "%20", -1)
	var oauth1Signature string = produceOauth1Signature(method, path, parameterString, oauth1Model)

	return "OAuth " +
		"oauth_signature_method=\"" + url.QueryEscape(urlValues.Get("oauth_signature_method")) + "\"," +
		"oauth_version=\"" + url.QueryEscape(urlValues.Get("oauth_version")) + "\"," +
		"oauth_nonce=\"" + url.QueryEscape(urlValues.Get("oauth_nonce")) + "\"," +
		"oauth_consumer_key=\"" + url.QueryEscape(urlValues.Get("oauth_consumer_key")) + "\"," +
		"oauth_token=\"" + url.QueryEscape(urlValues.Get("oauth_token")) + "\"," +
		"oauth_timestamp=\"" + url.QueryEscape(urlValues.Get("oauth_timestamp")) + "\"," +
		"oauth_signature=\"" + url.QueryEscape(oauth1Signature) + "\""
}

func produceOauth1Signature(method, fullPath, paramtersString string, oauth1Model *Oauth1Model) string {
	method = strings.ToUpper(method)
	fullPath = url.QueryEscape(strings.Split(fullPath, "?")[0])
	paramtersString = url.QueryEscape(paramtersString)

	var signatureBase string = fmt.Sprintf("%s&%s&%s", method, fullPath, paramtersString)
	var signingKey string = fmt.Sprintf("%s&%s", oauth1Model.ConsumerSecret, oauth1Model.TokenSecret)

	return calculateHmacSha1(signatureBase, signingKey)
}

func calculateHmacSha1(base, key string) string {
	var hmacHash hash.Hash = hmac.New(sha1.New, []byte(key))
	hmacHash.Write([]byte(base))

	return base64.StdEncoding.EncodeToString(hmacHash.Sum(nil))
}

func generateNonce() string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789"
	var seed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	var byteBuffer []byte = make([]byte, 12)
	for i := range byteBuffer {
		byteBuffer[i] = charset[seed.Intn(len(charset))]
	}

	return string(byteBuffer)
}
