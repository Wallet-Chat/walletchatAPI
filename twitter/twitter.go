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

type TweetSearchResults struct {
	Data []struct {
		EditHistoryTweetIds []string `json:"edit_history_tweet_ids"`
		ID                  string   `json:"id"`
		Text                string   `json:"text"`
	} `json:"data"`
	Meta struct {
		NewestID    string `json:"newest_id"`
		OldestID    string `json:"oldest_id"`
		ResultCount int    `json:"result_count"`
	} `json:"meta"`
}

func search(query string) {
	// Encode the query for use in the URL
	encodedQuery := url.QueryEscape(query)

	// Twitter API v2 Search Endpoint URL with 'since_id'
	url := fmt.Sprintf("https://api.twitter.com/2/tweets/search/recent?query=%s&max_results=10&since_id=%s", encodedQuery, sinceID)

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Bearer Token in the request header
	req.Header.Set("Authorization", "Bearer "+os.Getenv("TWITTER_BEARER_API"))

	// Send the GET request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse the JSON response
	var response TweetSearchResults
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return
	}

	// Iterate over the search results and update 'sinceID'
	for _, tweet := range response.Data {
		// Print the Twitter username associated with the tweet
		username := extractUsername(tweet.Text)
		fmt.Printf("Username: %s\n", username)

		// Check if "wallet_chat" is part of the tweet
		containsWalletChat := strings.Contains(tweet.Text, "wallet_chat")
		fmt.Printf("Contains 'wallet_chat': %v\n", containsWalletChat)
	}

	//update sinceID to the most recent ID, so we don't get results we have already processed
	sinceID = response.Meta.NewestID

	// Print the most recent tweet ID
	fmt.Printf("Most recent tweet ID for query '%s': %s\n", query, sinceID)
}

func extractUsername(text string) string {
	// Logic to extract the Twitter username from the tweet text
	// Modify this logic based on your tweet format or requirements

	// Example: Extract the username after "@" symbol
	parts := strings.Split(text, "@")
	if len(parts) > 1 {
		return parts[1]
	}

	// Return an empty string if no username found
	return ""
}

//used for manual testing, will have another endpoint or local call for actual use
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

			search(query_str)

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
