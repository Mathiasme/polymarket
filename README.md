# Polymarket Go Library

A Go client library for interacting with the [Polymarket](https://polymarket.com) API. Polymarket is a prediction market platform where users can trade on the outcomes of real-world events.

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Key Concepts](#key-concepts)
- [API Reference](#api-reference)
- [Examples](#examples)
- [Error Handling](#error-handling)
- [Contributing](#contributing)

## Installation

```bash
go get github.com/mathiasme/polymarket-go
```

## Quick Start

```go
package main

import (
    "log"
    "polymarket"
)

func main() {
    // Create a new client
    client := polymarket.NewClient()
    
    // Get active markets
    params := &polymarket.MarketsParams{
        Limit:  10,
        Active: boolPtr(true),
    }
    
    markets, err := client.GetMarkets(params)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, market := range markets {
        log.Printf("Market: %s", market.Question)
    }
}

func boolPtr(b bool) *bool { return &b }
```

## Key Concepts

### Markets
A **Market** represents a single event outcome on Polymarket. Each market has:
- A specific question or proposition
- Yes/No tokens representing stakes in the outcome
- Token prices that fluctuate between $0-$1
- Resolution where winning tokens are redeemable for $1 USDC

### Tokens
**Tokens** represent a stake in a specific Yes/No outcome in a Market:
- Price fluctuates between 0-1 based on market sentiment
- Redeemable for $1 USDC upon market resolution
- Trading occurs on the Central Limit Order Book (CLOB)

### Events
**Events** are collections of related markets, typically organized around a specific topic or time period (e.g., "2024 Presidential Election").

### Central Limit Order Book (CLOB)
The **CLOB** is Polymarket's off-chain order matching system where:
- Resting orders are placed and matched
- Market orders are processed before on-chain transmission
- Provides efficient price discovery and liquidity

## API Reference

### Client

#### `NewClient() *Client`
Creates a new Polymarket API client with default settings.

#### `NewClientWithOptions(baseURL string, timeout time.Duration) *Client`
Creates a client with custom base URL and timeout.

#### `SetTimeout(timeout time.Duration)`
Sets the HTTP client timeout.

### Markets

#### `GetMarkets(params *MarketsParams) ([]Market, error)`
Retrieves a list of markets with optional filtering and pagination.

**Parameters:**
```go
type MarketsParams struct {
    Limit      int    // Number of results to return
    Offset     int    // Number of results to skip
    Order      string // Field to order by
    Ascending  bool   // Sort order
    
    // Filters
    Active     *bool  // Filter by active status
    Closed     *bool  // Filter by closed status
    Archived   *bool  // Filter by archived status
    Slug       string // Filter by market slug
    EventID    string // Filter by event ID
    TagID      string // Filter by tag/category ID
}
```

#### `GetMarket(marketID string) (*Market, error)`
Retrieves a specific market by ID.

#### `GetMarketBySlug(slug string) (*Market, error)`
Retrieves a specific market by its slug.

### Events

#### `GetEvents(params *EventsParams) ([]Event, error)`
Retrieves a list of events with optional filtering and pagination.

#### `GetEvent(eventID string) (*Event, error)`
Retrieves a specific event by ID.

#### `GetEventBySlug(slug string) (*Event, error)`
Retrieves a specific event by its slug.

#### `GetEventMarkets(eventID string) ([]Market, error)`
Retrieves all markets for a specific event.

## Examples

### Get Active Markets with Pagination

```go
client := polymarket.NewClient()

params := &polymarket.MarketsParams{
    Limit:     20,
    Offset:    0,
    Active:    boolPtr(true),
    Ascending: true,
    Order:     "volume24hr",
}

markets, err := client.GetMarkets(params)
if err != nil {
    log.Fatal(err)
}

for _, market := range markets {
    fmt.Printf("Market: %s\n", market.Question)
    fmt.Printf("Volume (24h): %s\n", market.Volume24hr)
    fmt.Printf("Active: %v\n\n", market.Active)
}
```

### Get Markets for a Specific Event

```go
client := polymarket.NewClient()

// First, get an event
events, err := client.GetEvents(&polymarket.EventsParams{Limit: 1})
if err != nil {
    log.Fatal(err)
}

if len(events) > 0 {
    // Get markets for this event
    markets, err := client.GetEventMarkets(events[0].ID)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Event: %s\n", events[0].Title)
    fmt.Printf("Markets: %d\n", len(markets))
}
```

### Get Market by Slug

```go
client := polymarket.NewClient()

market, err := client.GetMarketBySlug("will-trump-win-2024-election")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Question: %s\n", market.Question)
fmt.Printf("Volume: %s\n", market.Volume)

for _, token := range market.Tokens {
    fmt.Printf("Token %s (%s): $%s\n", 
        token.Outcome, token.TokenID, token.Price)
}
```

### Custom Client Configuration

```go
import "time"

// Create client with custom timeout
client := polymarket.NewClientWithOptions(
    "https://gamma-api.polymarket.com", 
    60*time.Second,
)

// Or modify timeout later
client.SetTimeout(30 * time.Second)
```

## Error Handling

The library provides structured error handling:

```go
markets, err := client.GetMarkets(params)
if err != nil {
    // Check if it's an API error
    if apiErr, ok := err.(*polymarket.APIError); ok {
        fmt.Printf("API Error %d: %s\n", apiErr.Code, apiErr.Message)
    } else {
        // Network or other error
        fmt.Printf("Error: %v\n", err)
    }
}
```

## Network Information

Polymarket operates on the **Polygon Network**, a scalable, multi-chain blockchain platform. All market resolutions and token redemptions occur on-chain via smart contracts.

## Rate Limiting

Please be respectful of API rate limits. The library includes sensible defaults for timeouts and doesn't implement automatic retry logic to avoid overwhelming the API.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -am 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

This library is not officially affiliated with Polymarket. Use at your own risk. Always verify market conditions and understand the risks involved in prediction market trading.

## Resources

- [Polymarket Official Website](https://polymarket.com)
- [Polymarket API Documentation](https://docs.polymarket.com)
- [Polygon Network](https://polygon.technology)