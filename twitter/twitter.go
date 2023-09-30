package twitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"rest-go-demo/database"
	"rest-go-demo/entity"
	"strings"
	"time"

	"github.com/dghubble/oauth1"
	"github.com/gorilla/mux"
)

//since_id has to be at least within the ~5 days
var sinceID string

//in our db lets store the sinceID so we save requests/processing
//for recent search API, these has the be within the last 3-4 days usually.
type Globalstring struct {
	ID        string `json:"id"`
	Globalvar string `json:"globalvar"`
	Value     string `json:"value"`
}

func InitSearchParams() {
	var globalSinceID Globalstring
	dbResult := database.Connector.Where("globalvar = ?", "sinceid").Find(&globalSinceID)

	if dbResult.RowsAffected > 0 {
		if globalSinceID.Value == "" {
			sinceID = "1699404439969304999"
		} else {
			fmt.Println("Initializing Twitter sinceID: ", globalSinceID.Value)
			sinceID = globalSinceID.Value
		}
	}
}

type TwitterResponse struct {
	Data     []Tweet     `json:"data"`
	Includes TwitterData `json:"includes"`
	Meta     MetaData    `json:"meta"`
}

type Tweet struct {
	AuthorID string `json:"author_id"`
	ID       string `json:"id"`
	Text     string `json:"text"`
}

type TwitterData struct {
	Users []User `json:"users"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type MetaData struct {
	NewestID string `json:"newest_id"`
}

func searchTweets(query string) error {
	// URL encode the search query
	encodedQuery := url.QueryEscape(query)

	// Construct the Twitter API URL
	url := "https://api.twitter.com/2/tweets/search/recent?query=" + encodedQuery + "&since_id=" + sinceID + "&max_results=10&tweet.fields=author_id&expansions=author_id"
	//fmt.Println("URL for Twitter Search: ", url)

	// Make an HTTP request to the Twitter API
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("TWITTER_BEARER_API"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error In Twitter Search Request: ", err)
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error In Twitter Search Request: ", err)
		return err
	}

	// Parse the JSON response
	var twitterResponse TwitterResponse
	if err := json.Unmarshal(body, &twitterResponse); err != nil {
		return err
	}

	// Create a map to store user data by ID
	userDataByID := make(map[string]User)
	for _, user := range twitterResponse.Includes.Users {
		userDataByID[user.ID] = user
	}

	// Print the username and name for each tweet
	for _, tweet := range twitterResponse.Data {
		user, exists := userDataByID[tweet.AuthorID]
		if exists {
			fmt.Printf("Username: %s, Name: %s\n", user.Username, user.Name)

			//twitter handle is now verified within WalletChat, store the twitterID as well for future use maybe if user changes names
			//we need to handle the @symbol (maybe just ensure its added when from being saved with username?)
			database.Connector.Model(&entity.Settings{}).Where("twitteruser = ?", user.Username).Update("twitterverified", "true")
			database.Connector.Model(&entity.Settings{}).Where("twitteruser = ?", user.Username).Update("twitterid", tweet.AuthorID)

			//RespondToTweet(tweet.ID, "\n\nLFC! Check 🏆 inside the app and share your referral codes with frens here! #chat2earn")
		}
	}

	//update sinceID to the most recent ID, so we don't get results we have already processed
	if len(twitterResponse.Data) > 0 {
		sinceID = twitterResponse.Meta.NewestID
	}

	// Print the most recent tweet ID
	//fmt.Printf("Most recent tweet ID for query '%s': %s\n", query, sinceID)
	//now we update the value in the DB for next time around so we don't get duplicate results
	database.Connector.Model(&Globalstring{}).Where("globalvar = ?", "sinceid").Update("value", sinceID)

	return nil
}

func SearchVerifyUsernames() {
	searchTweets("LFC is the new LFG! (Lets F'n Chat!) @wallet_chat")
}

func SearchTweets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query_str := vars["query_str"]
	apiKey := r.Header.Get("Authorization")
	if len(apiKey) > 0 {
		const prefix = "Bearer "
		if len(apiKey) < len(prefix) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		apiKey = apiKey[len(prefix):]
		if strings.Contains(os.Getenv("ADMIN_API_KEY_LIST"), apiKey) {

			searchTweets(query_str)

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func SearchForUserReferralCodeTweets() {
	searchTweetsAndSendRefCodes("LFC is the new LFG! (I need a code!) @wallet_chat")
}

//shares a lot with searchTweets - TODO modularize a bit more
func searchTweetsAndSendRefCodes(query string) error {
	// URL encode the search query
	encodedQuery := url.QueryEscape(query)

	// Construct the Twitter API URL
	url := "https://api.twitter.com/2/tweets/search/recent?query=" + encodedQuery + "&since_id=" + sinceID + "&max_results=10&tweet.fields=author_id&expansions=author_id"
	//fmt.Println("URL for Twitter Search: ", url)

	// Make an HTTP request to the Twitter API
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("TWITTER_BEARER_API"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error In Twitter Search Request: ", err)
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error In Twitter Search Request: ", err)
		return err
	}

	// Parse the JSON response
	var twitterResponse TwitterResponse
	if err := json.Unmarshal(body, &twitterResponse); err != nil {
		return err
	}

	// Create a map to store user data by ID
	userDataByID := make(map[string]User)
	for _, user := range twitterResponse.Includes.Users {
		userDataByID[user.ID] = user
	}

	// Print the username and name for each tweet
	for _, tweet := range twitterResponse.Data {
		user, exists := userDataByID[tweet.AuthorID]
		if exists {
			fmt.Printf(" Request for new referral code - Username: %s, Name: %s\n", user.Username, user.Name)

			// database.Connector.Model(&entity.Settings{}).Where("twitteruser = ?", user.Username).Update("twitterverified", "true")
			// database.Connector.Model(&entity.Settings{}).Where("twitteruser = ?", user.Username).Update("twitterid", tweet.AuthorID)

			//if user is already in referralusers table, they don't need another code (TODO - need to store twitter id somewhere)
			var userAlreadyHasCode []entity.Referralcode
			database.Connector.Where("twitterid = ?", tweet.AuthorID).Find(&userAlreadyHasCode)

			if len(userAlreadyHasCode) > 0 {
				RespondToTweet(tweet.ID, "it's #chat2earn not #cheat2earn")
			} else {
				//get all items that relate to passed in owner/address
				var code entity.Referralcode
				code.Code = "wc-" + randSeq(10)
				code.Walletaddr = "0xtest"
				code.Date = time.Now()
				code.Twitterid = tweet.AuthorID
				database.Connector.Create(&code)

				RespondToTweet(tweet.ID, code.Code+"\n\nMake sure to go check 🏆 inside the app to share your codes with frens and win big! #chat2earn")
			}

			sinceID = tweet.ID
		}
	}

	//update sinceID to the most recent ID, so we don't get results we have already processed
	if len(twitterResponse.Data) > 0 {
		fmt.Println("Updated Latest Tweet ID: ", twitterResponse.Meta.NewestID)
		sinceID = twitterResponse.Meta.NewestID
	}

	// Print the most recent tweet ID
	//fmt.Printf("Most recent tweet ID for query '%s': %s\n", query, sinceID)
	//now we update the value in the DB for next time around so we don't get duplicate results
	database.Connector.Model(&Globalstring{}).Where("globalvar = ?", "sinceid").Update("value", sinceID)

	return nil
}

type Reply struct {
	InReplyToTweetId string `json:"in_reply_to_tweet_id"`
}

type TweetReply struct {
	Text  string `json:"text"`
	Reply *Reply `json:"reply"`
}

func RespondToTweet(tweet_id string, reply_text string) {
	// Replace these values with your own Twitter API credentials
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	// Set up OAuth1 configuration
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	url := "https://api.twitter.com/2/tweets"

	// Your tweet data
	tweet := TweetReply{
		Text: reply_text,
		Reply: &Reply{
			InReplyToTweetId: tweet_id,
		},
	}

	jsonStr, err := json.Marshal(tweet)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Error", err)
	}

	fmt.Println(string(body))
}
