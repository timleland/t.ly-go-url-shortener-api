package tly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client is the main API client for T.LY.
type Client struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

// NewClient creates a new T.LY API client.
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://api.t.ly",
		Client:  &http.Client{},
	}
}

// doRequest is an internal helper for making API calls.
func (c *Client) doRequest(method, path, query string, body interface{}, result interface{}) error {
	url := c.BaseURL + path
	if query != "" {
		url += "?" + query
	}
	var buf *bytes.Buffer
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(data)
	} else {
		buf = bytes.NewBuffer(nil)
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		data, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("API error: %s", string(data))
	}
	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

// =====================
// Pixel Management
// =====================

// Pixel represents a pixel object.
type Pixel struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	PixelID   string `json:"pixel_id"`
	PixelType string `json:"pixel_type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// PixelCreateRequest is used to create a new pixel.
type PixelCreateRequest struct {
	Name      string `json:"name"`
	PixelID   string `json:"pixel_id"`
	PixelType string `json:"pixel_type"`
}

// PixelUpdateRequest is used to update a pixel.
type PixelUpdateRequest struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	PixelID   string `json:"pixel_id"`
	PixelType string `json:"pixel_type"`
}

// CreatePixel calls the API to create a new pixel.
func (c *Client) CreatePixel(reqData PixelCreateRequest) (*Pixel, error) {
	var pixel Pixel
	err := c.doRequest("POST", "/api/v1/link/pixel", "", reqData, &pixel)
	if err != nil {
		return nil, err
	}
	return &pixel, nil
}

// ListPixels retrieves a list of pixels.
func (c *Client) ListPixels() ([]Pixel, error) {
	var pixels []Pixel
	err := c.doRequest("GET", "/api/v1/link/pixel", "", nil, &pixels)
	if err != nil {
		return nil, err
	}
	return pixels, nil
}

// GetPixel retrieves a pixel by its ID.
func (c *Client) GetPixel(id int) (*Pixel, error) {
	path := fmt.Sprintf("/api/v1/link/pixel/%d", id)
	var pixel Pixel
	err := c.doRequest("GET", path, "", nil, &pixel)
	if err != nil {
		return nil, err
	}
	return &pixel, nil
}

// UpdatePixel updates an existing pixel.
func (c *Client) UpdatePixel(reqData PixelUpdateRequest) (*Pixel, error) {
	path := fmt.Sprintf("/api/v1/link/pixel/%d", reqData.ID)
	var pixel Pixel
	err := c.doRequest("PUT", path, "", reqData, &pixel)
	if err != nil {
		return nil, err
	}
	return &pixel, nil
}

// DeletePixel deletes a pixel by its ID.
func (c *Client) DeletePixel(id int) error {
	path := fmt.Sprintf("/api/v1/link/pixel/%d", id)
	return c.doRequest("DELETE", path, "", nil, nil)
}

// =====================
// Short Link Management
// =====================

// ShortLink represents a shortened URL.
type ShortLink struct {
	ShortURL         string      `json:"short_url"`
	Description      string      `json:"description"`
	LongURL          string      `json:"long_url"`
	Domain           string      `json:"domain"`
	ShortID          string      `json:"short_id"`
	ExpireAtViews    interface{} `json:"expire_at_views"`
	ExpireAtDatetime interface{} `json:"expire_at_datetime"`
	PublicStats      bool        `json:"public_stats"`
	CreatedAt        string      `json:"created_at"`
	UpdatedAt        string      `json:"updated_at"`
	Meta             interface{} `json:"meta"`
}

// ShortLinkCreateRequest is used to create a short link.
type ShortLinkCreateRequest struct {
	LongURL          string      `json:"long_url"`
	ShortID          *string     `json:"short_id,omitempty"`
	Domain           string      `json:"domain"`
	ExpireAtDatetime *string     `json:"expire_at_datetime,omitempty"`
	ExpireAtViews    *int        `json:"expire_at_views,omitempty"`
	Description      *string     `json:"description,omitempty"`
	PublicStats      *bool       `json:"public_stats,omitempty"`
	Password         *string     `json:"password,omitempty"`
	Tags             []int       `json:"tags,omitempty"`
	Pixels           []int       `json:"pixels,omitempty"`
	Meta             interface{} `json:"meta,omitempty"`
}

// ShortLinkUpdateRequest is used to update a short link.
type ShortLinkUpdateRequest struct {
	ShortURL         string      `json:"short_url"`
	ShortID          *string     `json:"short_id,omitempty"`
	LongURL          string      `json:"long_url"`
	ExpireAtDatetime *string     `json:"expire_at_datetime,omitempty"`
	ExpireAtViews    *int        `json:"expire_at_views,omitempty"`
	Description      *string     `json:"description,omitempty"`
	PublicStats      *bool       `json:"public_stats,omitempty"`
	Password         *string     `json:"password,omitempty"`
	Tags             []int       `json:"tags,omitempty"`
	Pixels           []int       `json:"pixels,omitempty"`
	Meta             interface{} `json:"meta,omitempty"`
}

// CreateShortLink creates a new short link.
func (c *Client) CreateShortLink(reqData ShortLinkCreateRequest) (*ShortLink, error) {
	var link ShortLink
	err := c.doRequest("POST", "/api/v1/link/shorten", "", reqData, &link)
	if err != nil {
		return nil, err
	}
	return &link, nil
}

// GetShortLink retrieves a short link using its URL.
func (c *Client) GetShortLink(shortURL string) (*ShortLink, error) {
	query := "short_url=" + shortURL
	var link ShortLink
	err := c.doRequest("GET", "/api/v1/link", query, nil, &link)
	if err != nil {
		return nil, err
	}
	return &link, nil
}

// UpdateShortLink updates an existing short link.
func (c *Client) UpdateShortLink(reqData ShortLinkUpdateRequest) (*ShortLink, error) {
	var link ShortLink
	err := c.doRequest("PUT", "/api/v1/link", "", reqData, &link)
	if err != nil {
		return nil, err
	}
	return &link, nil
}

// DeleteShortLink deletes a short link.
func (c *Client) DeleteShortLink(shortURL string) error {
	reqBody := map[string]string{
		"short_url": shortURL,
	}
	return c.doRequest("DELETE", "/api/v1/link", "", reqBody, nil)
}

// ExpandRequest is used to expand a short link.
type ExpandRequest struct {
	ShortURL string  `json:"short_url"`
	Password *string `json:"password,omitempty"`
}

// ExpandResponse represents the response when expanding a short link.
type ExpandResponse struct {
	LongURL string `json:"long_url"`
	Expired bool   `json:"expired"`
}

// ExpandShortLink expands a short URL to its original long URL.
func (c *Client) ExpandShortLink(reqData ExpandRequest) (*ExpandResponse, error) {
	var resp ExpandResponse
	err := c.doRequest("POST", "/api/v1/link/expand", "", reqData, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListShortLinks retrieves a list of short links using optional query parameters.
// The queryParams map can include keys such as "search", "tag_ids", "pixel_ids", etc.
func (c *Client) ListShortLinks(queryParams map[string]string) (string, error) {
	query := ""
	first := true
	for k, v := range queryParams {
		if !first {
			query += "&"
		}
		query += fmt.Sprintf("%s=%s", k, v)
		first = false
	}
	// The API returns a plain text JSON string.
	var result string
	err := c.doRequest("GET", "/api/v1/link/list", query, nil, &result)
	if err != nil {
		return "", err
	}
	return result, nil
}

// BulkShortenRequest is used for bulk shortening of links.
type BulkShortenRequest struct {
	Domain string   `json:"domain"`
	Links  []string `json:"links"` // For simplicity, using a slice of URLs.
	Tags   []int    `json:"tags,omitempty"`
	Pixels []int    `json:"pixels,omitempty"`
}

// BulkShortenLinks sends a bulk shorten request.
func (c *Client) BulkShortenLinks(reqData BulkShortenRequest) (string, error) {
	var result string
	err := c.doRequest("POST", "/api/v1/link/bulk", "", reqData, &result)
	if err != nil {
		return "", err
	}
	return result, nil
}

// =====================
// Stats Management
// =====================

// Stats represents the statistics for a short link.
type Stats struct {
	Clicks       int                    `json:"clicks"`
	UniqueClicks int                    `json:"unique_clicks"`
	Browsers     []interface{}          `json:"browsers"`
	Countries    []interface{}          `json:"countries"`
	Referrers    []interface{}          `json:"referrers"`
	Platforms    []interface{}          `json:"platforms"`
	DailyClicks  []interface{}          `json:"daily_clicks"`
	Data         map[string]interface{} `json:"data"`
}

// GetStats retrieves statistics for a given short link.
func (c *Client) GetStats(shortURL string) (*Stats, error) {
	query := "short_url=" + shortURL
	var stats Stats
	err := c.doRequest("GET", "/api/v1/link/stats", query, nil, &stats)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

// =====================
// Tag Management
// =====================

// Tag represents a tag.
type Tag struct {
	ID        int    `json:"id"`
	Tag       string `json:"tag"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ListTags retrieves all tags.
func (c *Client) ListTags() ([]Tag, error) {
	var tags []Tag
	err := c.doRequest("GET", "/api/v1/link/tag", "", nil, &tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

// CreateTag creates a new tag.
func (c *Client) CreateTag(tagValue string) (*Tag, error) {
	reqBody := map[string]string{
		"tag": tagValue,
	}
	var tag Tag
	err := c.doRequest("POST", "/api/v1/link/tag", "", reqBody, &tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

// GetTag retrieves a tag by its ID.
func (c *Client) GetTag(id int) (*Tag, error) {
	path := fmt.Sprintf("/api/v1/link/tag/%d", id)
	var tag Tag
	err := c.doRequest("GET", path, "", nil, &tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

// UpdateTag updates an existing tag.
func (c *Client) UpdateTag(id int, tagValue string) (*Tag, error) {
	path := fmt.Sprintf("/api/v1/link/tag/%d", id)
	reqBody := map[string]string{
		"tag": tagValue,
	}
	var tag Tag
	err := c.doRequest("PUT", path, "", reqBody, &tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

// DeleteTag deletes a tag by its ID.
func (c *Client) DeleteTag(id int) error {
	path := fmt.Sprintf("/api/v1/link/tag/%d", id)
	return c.doRequest("DELETE", path, "", nil, nil)
}
