package wc_analytics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Event struct {
	Name   string            `json:"name"`
	Params map[string]string `json:"params"`
}

type EventData struct {
	ClientID string `json:"client_id"`
	Events   []struct {
		Name   string `json:"name"`
		Params struct {
			Walletaddr string `json:"walletaddr"`
		} `json:"params"`
	} `json:"events"`
}

type EventDataWithEmail struct {
	ClientID string `json:"client_id"`
	Events   []struct {
		Name   string `json:"name"`
		Params struct {
			Walletaddr string `json:"walletaddr"`
			Email      string `json:"email"`
		} `json:"params"`
	} `json:"events"`
}

type EventDataWithSignupSite struct {
	ClientID string `json:"client_id"`
	Events   []struct {
		Name   string `json:"name"`
		Params struct {
			Walletaddr string `json:"walletaddr"`
			Signupsite string `json:"signupsite"`
		} `json:"params"`
	} `json:"events"`
}

func SendCustomIntraEvent(clientID string, eventName string) error { //eventParams map[string]interface{}) error {
	apiUrl := "https://www.google-analytics.com/mp/collect?measurement_id=" + os.Getenv("GOOGLE_GA4_MEASUREMENT_ID_INTRA") + "&api_secret=" + os.Getenv("GOOGLE_GA4_API_KEY_INTRA")

	eventData := EventData{
		ClientID: clientID,
		Events: []struct {
			Name   string `json:"name"`
			Params struct {
				Walletaddr string `json:"walletaddr"`
			} `json:"params"`
		}{
			{
				Name: eventName,
				Params: struct {
					Walletaddr string `json:"walletaddr"`
				}{
					Walletaddr: "." + clientID, //GA4 converts this into a goofy integer otherwise, add . for workaround
				},
			},
		},
	}
	eventDataJson, err := json.Marshal(eventData)
	if err != nil {
		fmt.Println("GA4 Called Custom Event - Error 0 Repsonse: ", err)
		return err
	}
	eventDataBytes := bytes.NewBuffer(eventDataJson)

	req, err := http.NewRequest("POST", apiUrl, eventDataBytes)
	if err != nil {
		fmt.Println("GA4 Called Custom Event - Error 1 Repsonse: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("GA4 Called Custom Event - Error 2 Repsonse: ", err)
		return err
	}

	//fmt.Println("GA4 Called Custom Event - HTTP Repsonse: ", resp)

	defer resp.Body.Close()

	return nil
}

func SendCustomEvent(clientID string, eventName string) error { //eventParams map[string]interface{}) error {
	apiUrl := "https://www.google-analytics.com/mp/collect?measurement_id=" + os.Getenv("GOOGLE_GA4_MEASUREMENT_ID") + "&api_secret=" + os.Getenv("GOOGLE_GA4_API_KEY")

	eventData := EventData{
		ClientID: clientID,
		Events: []struct {
			Name   string `json:"name"`
			Params struct {
				Walletaddr string `json:"walletaddr"`
			} `json:"params"`
		}{
			{
				Name: eventName,
				Params: struct {
					Walletaddr string `json:"walletaddr"`
				}{
					Walletaddr: "." + clientID, //GA4 converts this into a goofy integer otherwise, add . for workaround
				},
			},
		},
	}
	eventDataJson, err := json.Marshal(eventData)
	if err != nil {
		fmt.Println("GA4 Called Custom Event - Error 0 Repsonse: ", err)
		return err
	}
	eventDataBytes := bytes.NewBuffer(eventDataJson)

	req, err := http.NewRequest("POST", apiUrl, eventDataBytes)
	if err != nil {
		fmt.Println("GA4 Called Custom Event - Error 1 Repsonse: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("GA4 Called Custom Event - Error 2 Repsonse: ", err)
		return err
	}

	//fmt.Println("GA4 Called Custom Event - HTTP Repsonse: ", resp)

	defer resp.Body.Close()

	return nil
}

func SendCustomEventWithEmail(clientID string, eventName string, emailInput string) error { //eventParams map[string]interface{}) error {
	apiUrl := "https://www.google-analytics.com/mp/collect?measurement_id=" + os.Getenv("GOOGLE_GA4_MEASUREMENT_ID") + "&api_secret=" + os.Getenv("GOOGLE_GA4_API_KEY")

	eventData := EventDataWithEmail{
		ClientID: clientID,
		Events: []struct {
			Name   string `json:"name"`
			Params struct {
				Walletaddr string `json:"walletaddr"`
				Email      string `json:"email"`
			} `json:"params"`
		}{
			{
				Name: eventName,
				Params: struct {
					Walletaddr string `json:"walletaddr"`
					Email      string `json:"email"`
				}{
					Walletaddr: "." + clientID, //GA4 converts this into a goofy integer otherwise, add . for workaround
					Email:      emailInput,
				},
			},
		},
	}
	eventDataJson, err := json.Marshal(eventData)
	if err != nil {
		fmt.Println("GA4 Called Custom Event - Error 0 Repsonse: ", err)
		return err
	}
	eventDataBytes := bytes.NewBuffer(eventDataJson)

	req, err := http.NewRequest("POST", apiUrl, eventDataBytes)
	if err != nil {
		fmt.Println("GA4 Called Custom Event - Error 1 Repsonse: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("GA4 Called Custom Event - Error 2 Repsonse: ", err)
		return err
	}

	//fmt.Println("GA4 Called Custom Event - HTTP Repsonse: ", resp)

	defer resp.Body.Close()

	return nil
}

func SendCustomEventWithSignupSite(clientID string, eventName string, signupSite string) error { //eventParams map[string]interface{}) error {
	apiUrl := "https://www.google-analytics.com/mp/collect?measurement_id=" + os.Getenv("GOOGLE_GA4_MEASUREMENT_ID") + "&api_secret=" + os.Getenv("GOOGLE_GA4_API_KEY")

	eventData := EventDataWithSignupSite{
		ClientID: clientID,
		Events: []struct {
			Name   string `json:"name"`
			Params struct {
				Walletaddr string `json:"walletaddr"`
				Signupsite string `json:"signupsite"`
			} `json:"params"`
		}{
			{
				Name: eventName,
				Params: struct {
					Walletaddr string `json:"walletaddr"`
					Signupsite string `json:"signupsite"`
				}{
					Walletaddr: "." + clientID, //GA4 converts this into a goofy integer otherwise, add . for workaround
					Signupsite: signupSite,
				},
			},
		},
	}
	eventDataJson, err := json.Marshal(eventData)
	if err != nil {
		fmt.Println("GA4 Called Custom Event - Error 0 Repsonse: ", err)
		return err
	}
	eventDataBytes := bytes.NewBuffer(eventDataJson)

	req, err := http.NewRequest("POST", apiUrl, eventDataBytes)
	if err != nil {
		fmt.Println("GA4 Called Custom Event - Error 1 Repsonse: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("GA4 Called Custom Event - Error 2 Repsonse: ", err)
		return err
	}

	//fmt.Println("GA4 Called Custom Event - HTTP Repsonse: ", resp)

	defer resp.Body.Close()

	return nil
}
