package android

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const scopeFCM = "https://www.googleapis.com/auth/firebase.messaging"

// Config holds initialization params for the client.
type Config struct {
	ProjectID          string
	ServiceAccountJSON []byte
	HTTPTimeout        time.Duration
}

// FromEnv builds Config using environment variables:
// - PROJECT_ID
// - GCP_SA_JSON (full JSON, not a path)

func FromEnv() (Config, error) {
	projectID := os.Getenv("ANDROID_PROJECT_ID")
	if projectID == "" {
		return Config{}, fmt.Errorf("missing env ANDROID_PROJECT_ID")
	}

	// Try base64 first
	if b64 := os.Getenv("ANDROID_GCP_SA_JSON_B64"); b64 != "" {
		decoded, err := base64.StdEncoding.DecodeString(b64)
		if err != nil {
			return Config{}, fmt.Errorf("failed to decode ANDROID_GCP_SA_JSON_B64: %w", err)
		}
		return Config{ProjectID: projectID, ServiceAccountJSON: decoded}, nil
	}

	// Fallback: raw JSON
	if raw := os.Getenv("ANDROID_GCP_SA_JSON"); raw != "" {
		return Config{ProjectID: projectID, ServiceAccountJSON: []byte(raw)}, nil
	}

	return Config{}, fmt.Errorf("missing service account JSON env (ANDROID_GCP_SA_JSON_B64 or ANDROID_GCP_SA_JSON)")
}

// Client sends Android push notifications via FCM HTTP v1.
type Client struct {
	projectID  string
	httpClient *http.Client
}

// NewClient initializes a client from Config.
func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	if cfg.ProjectID == "" {
		return nil, fmt.Errorf("projectID is required")
	}
	if len(cfg.ServiceAccountJSON) == 0 {
		return nil, fmt.Errorf("service account JSON is required")
	}

	jwtCfg, err := google.JWTConfigFromJSON(cfg.ServiceAccountJSON, scopeFCM)
	if err != nil {
		return nil, fmt.Errorf("parse service account JSON: %w", err)
	}
	httpClient := oauth2.NewClient(ctx, jwtCfg.TokenSource(ctx))
	if cfg.HTTPTimeout <= 0 {
		httpClient.Timeout = 30 * time.Second
	} else {
		httpClient.Timeout = cfg.HTTPTimeout
	}

	return &Client{
		projectID:  cfg.ProjectID,
		httpClient: httpClient,
	}, nil
}

// --- Payload types ---

type Request struct {
	Message      Message `json:"message"`
	ValidateOnly bool    `json:"validate_only,omitempty"`
}

type Message struct {
	Token        string        `json:"token,omitempty"`
	Notification *Notification `json:"notification,omitempty"`
	Android      *Android      `json:"android,omitempty"`
}

type Notification struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

type Android struct {
	Priority     string               `json:"priority,omitempty"`
	TTL          string               `json:"ttl,omitempty"`
	Notification *AndroidNotification `json:"notification,omitempty"`
}

type AndroidNotification struct {
	Sound             string `json:"sound,omitempty"`
	ChannelID         string `json:"channel_id,omitempty"`
	NotificationCount int    `json:"notification_count,omitempty"`
}

type Response struct {
	Name       string `json:"name,omitempty"`
	StatusCode int    `json:"-"`
	RawBody    string `json:"-"`
}

// Send posts the message to FCM HTTP v1.
func (c *Client) Send(ctx context.Context, msg Message, validateOnly bool) (*Response, error) {
	if msg.Token == "" {
		return nil, fmt.Errorf("message token is required")
	}

	reqBody := Request{
		Message:      msg,
		ValidateOnly: validateOnly,
	}
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	url := fmt.Sprintf("https://fcm.googleapis.com/v1/projects/%s/messages:send", c.projectID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer res.Body.Close()

	resBytes, _ := io.ReadAll(res.Body)
	out := &Response{
		StatusCode: res.StatusCode,
		RawBody:    string(resBytes),
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		_ = json.Unmarshal(resBytes, out) // parse "name" if present
		return out, nil
	}
	return out, fmt.Errorf("send failed: %d — %s", res.StatusCode, out.RawBody)
}
