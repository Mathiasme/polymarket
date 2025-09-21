package polymarket

import (
	"encoding/json"
	"fmt"
)

// GetEvents retrieves a list of events from the Polymarket API
func (c *Client) GetEvents(params *EventsParams) ([]Event, error) {
	body, err := c.makeRequest("GET", "/events", buildParams(params))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events: %w", err)
	}
	
	var events []Event
	if err := json.Unmarshal(body, &events); err != nil {
		return nil, fmt.Errorf("failed to parse events response: %w", err)
	}
	
	return events, nil
}

// GetEvent retrieves a specific event by its ID
func (c *Client) GetEvent(eventID string) (*Event, error) {
	return c.GetEventWithParams(eventID, nil)
}

// GetEventWithParams retrieves a specific event by its ID with optional parameters
func (c *Client) GetEventWithParams(eventID string, params *GetEventParams) (*Event, error) {
	endpoint := fmt.Sprintf("/events/%s", eventID)
	
	body, err := c.makeRequest("GET", endpoint, buildParams(params))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch event %s: %w", eventID, err)
	}
	
	var event Event
	if err := json.Unmarshal(body, &event); err != nil {
		return nil, fmt.Errorf("failed to parse event response: %w", err)
	}
	
	return &event, nil
}

// GetEventBySlug retrieves a specific event by its slug
func (c *Client) GetEventBySlug(slug string) (*Event, error) {
	params := &EventsParams{
		Slug:  slug,
		Limit: 1,
	}
	
	events, err := c.GetEvents(params)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch event by slug %s: %w", slug, err)
	}
	
	if len(events) == 0 {
		return nil, fmt.Errorf("event with slug %s not found", slug)
	}
	
	return &events[0], nil
}

// GetEventMarkets retrieves all markets for a specific event
func (c *Client) GetEventMarkets(eventID string) ([]Market, error) {
	params := &MarketsParams{
		EventID: eventID,
	}
	
	return c.GetMarkets(params)
}