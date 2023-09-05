package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

//since_id has to be at least within the ~5 days
var sinceID string

func InitSearchParams() {
	sinceID = "1698219441811537927"
}

//req.Header.Set("Authorization", "Bearer "+os.Getenv("TWITTER_BEARER_API"))

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
	fmt.Println("URL for Twitter Search: ", url)

	// Make an HTTP request to the Twitter API
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("TWITTER_BEARER_API"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
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
		}
	}

	//update sinceID to the most recent ID, so we don't get results we have already processed
	if len(twitterResponse.Data) > 0 {
		sinceID = twitterResponse.Meta.NewestID
	}

	// Print the most recent tweet ID
	fmt.Printf("Most recent tweet ID for query '%s': %s\n", query, sinceID)

	return nil
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
