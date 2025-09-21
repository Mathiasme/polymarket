package polymarket

import (
	"encoding/json"
	"fmt"
)

// Search performs a unified search across markets, events, and profiles
func (c *Client) Search(params *SearchParams) (*SearchResults, error) {
	if params == nil || params.Q == "" {
		return nil, fmt.Errorf("search query (Q) is required")
	}
	
	body, err := c.makeRequest("GET", "/public-search", buildParams(params))
	if err != nil {
		return nil, fmt.Errorf("failed to perform search: %w", err)
	}
	
	var results SearchResults
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, fmt.Errorf("failed to parse search results: %w", err)
	}
	
	return &results, nil
}

// SearchEvents searches specifically for events
func (c *Client) SearchEvents(query string, params *SearchParams) ([]Event, error) {
	if params == nil {
		params = &SearchParams{}
	}
	params.Q = query
	params.SearchTags = boolPtr(false)
	params.SearchProfiles = boolPtr(false)
	
	results, err := c.Search(params)
	if err != nil {
		return nil, err
	}
	
	return results.Events, nil
}

// SearchProfiles searches specifically for user profiles
func (c *Client) SearchProfiles(query string, params *SearchParams) ([]UserProfile, error) {
	if params == nil {
		params = &SearchParams{}
	}
	params.Q = query
	params.SearchTags = boolPtr(false)
	params.SearchProfiles = boolPtr(true)
	
	results, err := c.Search(params)
	if err != nil {
		return nil, err
	}
	
	return results.Profiles, nil
}

// SearchTags searches specifically for tags
func (c *Client) SearchTags(query string, params *SearchParams) ([]Tag, error) {
	if params == nil {
		params = &SearchParams{}
	}
	params.Q = query
	params.SearchTags = boolPtr(true)
	params.SearchProfiles = boolPtr(false)
	
	results, err := c.Search(params)
	if err != nil {
		return nil, err
	}
	
	return results.Tags, nil
}

// SearchByTag searches for events by specific tags
func (c *Client) SearchByTag(query string, tags []string, params *SearchParams) ([]Event, error) {
	if params == nil {
		params = &SearchParams{}
	}
	params.Q = query
	params.EventsTag = tags
	
	results, err := c.Search(params)
	if err != nil {
		return nil, err
	}
	
	return results.Events, nil
}

