package polymarket

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	// DefaultBaseURL is the default Polymarket API base URL
	DefaultBaseURL = "https://gamma-api.polymarket.com"
	
	// DefaultTimeout is the default HTTP client timeout
	DefaultTimeout = 30 * time.Second
)

// Client is the main client for interacting with the Polymarket API
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Polymarket API client
func NewClient() *Client {
	return &Client{
		baseURL: DefaultBaseURL,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

// NewClientWithOptions creates a new client with custom options
func NewClientWithOptions(baseURL string, timeout time.Duration) *Client {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// SetTimeout sets the HTTP client timeout
func (c *Client) SetTimeout(timeout time.Duration) {
	c.httpClient.Timeout = timeout
}

// makeRequest performs an HTTP request and returns the response body
func (c *Client) makeRequest(method, endpoint string, params url.Values) ([]byte, error) {
	// Construct full URL
	fullURL := c.baseURL + endpoint
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}
	
	// Create request
	req, err := http.NewRequest(method, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "polymarket-go-client/1.0")
	
	// Make request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()
	
	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	
	// Check for API errors
	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
		}
		apiErr.Code = resp.StatusCode
		return nil, &apiErr
	}
	
	return body, nil
}

// buildParams converts struct parameters to url.Values
func buildParams(params interface{}) url.Values {
	values := url.Values{}
	
	switch p := params.(type) {
	case *MarketsParams:
		if p == nil {
			return values
		}
		if p.Limit > 0 {
			values.Add("limit", strconv.Itoa(p.Limit))
		}
		if p.Offset > 0 {
			values.Add("offset", strconv.Itoa(p.Offset))
		}
		if p.Order != "" {
			values.Add("order", p.Order)
		}
		if p.Ascending {
			values.Add("ascending", "true")
		}
		if p.Active != nil {
			values.Add("active", strconv.FormatBool(*p.Active))
		}
		if p.Closed != nil {
			values.Add("closed", strconv.FormatBool(*p.Closed))
		}
		if p.Archived != nil {
			values.Add("archived", strconv.FormatBool(*p.Archived))
		}
		if p.Slug != "" {
			values.Add("slug", p.Slug)
		}
		if p.EventID != "" {
			values.Add("event_id", p.EventID)
		}
		if p.TagID != "" {
			values.Add("tag_id", p.TagID)
		}
		
	case *EventsParams:
		if p == nil {
			return values
		}
		// Pagination
		if p.Limit > 0 {
			values.Add("limit", strconv.Itoa(p.Limit))
		}
		if p.Offset > 0 {
			values.Add("offset", strconv.Itoa(p.Offset))
		}
		
		// Sorting
		if p.Order != "" {
			values.Add("order", p.Order)
		}
		if p.Ascending {
			values.Add("ascending", "true")
		}
		
		// Basic filters
		if len(p.ID) > 0 {
			for _, id := range p.ID {
				values.Add("id", id)
			}
		}
		if len(p.Slug) > 0 {
			for _, slug := range p.Slug {
				values.Add("slug", slug)
			}
		}
		if p.Active != nil {
			values.Add("active", strconv.FormatBool(*p.Active))
		}
		if p.Closed != nil {
			values.Add("closed", strconv.FormatBool(*p.Closed))
		}
		if p.Archived != nil {
			values.Add("archived", strconv.FormatBool(*p.Archived))
		}
		
		// Advanced filters
		if p.TagID != nil {
			values.Add("tag_id", strconv.Itoa(*p.TagID))
		}
		if len(p.ExcludeTagID) > 0 {
			for _, tagID := range p.ExcludeTagID {
				values.Add("exclude_tag_id", strconv.Itoa(tagID))
			}
		}
		if p.RelatedTags != nil {
			values.Add("related_tags", strconv.FormatBool(*p.RelatedTags))
		}
		if p.Featured != nil {
			values.Add("featured", strconv.FormatBool(*p.Featured))
		}
		if p.CYOM != nil {
			values.Add("cyom", strconv.FormatBool(*p.CYOM))
		}
		if p.Recurrence != "" {
			values.Add("recurrence", p.Recurrence)
		}
		
		// Date filters
		if p.StartDateMin != nil {
			values.Add("start_date_min", p.StartDateMin.Format(time.RFC3339))
		}
		if p.StartDateMax != nil {
			values.Add("start_date_max", p.StartDateMax.Format(time.RFC3339))
		}
		if p.EndDateMin != nil {
			values.Add("end_date_min", p.EndDateMin.Format(time.RFC3339))
		}
		if p.EndDateMax != nil {
			values.Add("end_date_max", p.EndDateMax.Format(time.RFC3339))
		}
		
	case *GetEventParams:
		if p == nil {
			return values
		}
		if p.IncludeChat != nil {
			values.Add("include_chat", strconv.FormatBool(*p.IncludeChat))
		}
		if p.IncludeTemplate != nil {
			values.Add("include_template", strconv.FormatBool(*p.IncludeTemplate))
		}
		
	case *GetMarketParams:
		if p == nil {
			return values
		}
		if p.IncludeTag != nil {
			values.Add("include_tag", strconv.FormatBool(*p.IncludeTag))
		}
		
	case *CommentsParams:
		if p == nil {
			return values
		}
		// Pagination
		if p.Limit > 0 {
			values.Add("limit", strconv.Itoa(p.Limit))
		}
		if p.Offset > 0 {
			values.Add("offset", strconv.Itoa(p.Offset))
		}
		
		// Sorting
		if p.Order != "" {
			values.Add("order", p.Order)
		}
		if p.Ascending {
			values.Add("ascending", "true")
		}
		
		// Filters
		if p.ParentEntityType != "" {
			values.Add("parent_entity_type", p.ParentEntityType)
		}
		if p.ParentEntityID != nil {
			values.Add("parent_entity_id", strconv.Itoa(*p.ParentEntityID))
		}
		if p.GetPositions != nil {
			values.Add("get_positions", strconv.FormatBool(*p.GetPositions))
		}
		if p.HoldersOnly != nil {
			values.Add("holders_only", strconv.FormatBool(*p.HoldersOnly))
		}
		
	case *SearchParams:
		if p == nil {
			return values
		}
		// Required
		if p.Q != "" {
			values.Add("q", p.Q)
		}
		
		// Pagination
		if p.Page > 0 {
			values.Add("page", strconv.Itoa(p.Page))
		}
		if p.LimitPerType > 0 {
			values.Add("limit_per_type", strconv.Itoa(p.LimitPerType))
		}
		
		// Sorting
		if p.Sort != "" {
			values.Add("sort", p.Sort)
		}
		if p.Ascending != nil {
			values.Add("ascending", strconv.FormatBool(*p.Ascending))
		}
		
		// Filters and options
		if p.Cache != nil {
			values.Add("cache", strconv.FormatBool(*p.Cache))
		}
		if p.EventsStatus != "" {
			values.Add("events_status", p.EventsStatus)
		}
		if len(p.EventsTag) > 0 {
			for _, tag := range p.EventsTag {
				values.Add("events_tag", tag)
			}
		}
		if p.KeepClosedMarkets != nil {
			values.Add("keep_closed_markets", strconv.Itoa(*p.KeepClosedMarkets))
		}
		if p.SearchTags != nil {
			values.Add("search_tags", strconv.FormatBool(*p.SearchTags))
		}
		if p.SearchProfiles != nil {
			values.Add("search_profiles", strconv.FormatBool(*p.SearchProfiles))
		}
		if p.Recurrence != "" {
			values.Add("recurrence", p.Recurrence)
		}
		if len(p.ExcludeTagID) > 0 {
			for _, tagID := range p.ExcludeTagID {
				values.Add("exclude_tag_id", strconv.Itoa(tagID))
			}
		}
		if p.Optimized != nil {
			values.Add("optimized", strconv.FormatBool(*p.Optimized))
		}
	}
	
	return values
}