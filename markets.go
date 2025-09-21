package polymarket

import (
	"encoding/json"
	"fmt"
)

// GetMarkets retrieves a list of markets from the Polymarket API
func (c *Client) GetMarkets(params *MarketsParams) ([]Market, error) {
	body, err := c.makeRequest("GET", "/markets", buildParams(params))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch markets: %w", err)
	}
	
	var markets []Market
	if err := json.Unmarshal(body, &markets); err != nil {
		return nil, fmt.Errorf("failed to parse markets response: %w", err)
	}
	
	return markets, nil
}

// GetMarket retrieves a specific market by its ID
func (c *Client) GetMarket(marketID string) (*Market, error) {
	return c.GetMarketWithParams(marketID, nil)
}

// GetMarketWithParams retrieves a specific market by its ID with optional parameters
func (c *Client) GetMarketWithParams(marketID string, params *GetMarketParams) (*Market, error) {
	endpoint := fmt.Sprintf("/markets/%s", marketID)
	
	body, err := c.makeRequest("GET", endpoint, buildParams(params))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch market %s: %w", marketID, err)
	}
	
	var market Market
	if err := json.Unmarshal(body, &market); err != nil {
		return nil, fmt.Errorf("failed to parse market response: %w", err)
	}
	
	return &market, nil
}

// GetMarketBySlug retrieves a specific market by its slug
func (c *Client) GetMarketBySlug(slug string) (*Market, error) {
	params := &MarketsParams{
		Slug:  slug,
		Limit: 1,
	}
	
	markets, err := c.GetMarkets(params)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch market by slug %s: %w", slug, err)
	}
	
	if len(markets) == 0 {
		return nil, fmt.Errorf("market with slug %s not found", slug)
	}
	
	return &markets[0], nil
}