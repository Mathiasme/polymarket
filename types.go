package polymarket

import "time"

// Market represents a single event outcome on Polymarket
type Market struct {
	ID                 string     `json:"id"`
	Question           string     `json:"question"`
	Slug               string     `json:"slug"`
	Description        string     `json:"description"`
	StartDate          *time.Time `json:"startDate"`
	EndDate            *time.Time `json:"endDate"`
	Image              string     `json:"image"`
	Icon               string     `json:"icon"`
	Active             bool       `json:"active"`
	Closed             bool       `json:"closed"`
	MarketMakerAddress string     `json:"marketMakerAddress"`

	// Market type and outcomes
	MarketType string `json:"marketType"`
	Outcomes   string `json:"outcomes"`

	// Trading information
	Volume       string  `json:"volume"`
	Volume24hr   float64 `json:"volume24hr"`
	LiquidityNum float64 `json:"liquidityNum"`

	// Market structure
	Tokens     []Token    `json:"tokens"`
	Events     []Event    `json:"events"`
	Categories []Category `json:"categories"`
	Tags       []Tag      `json:"tags"`

	// Metadata
	QuestionID  string `json:"questionId"`
	ConditionID string `json:"conditionId"`
	UmaAddress  string `json:"umaAddress"`

	// Timestamps
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// Token represents a stake in a specific Yes/No outcome in a Market
// Price fluctuates between 0-1 and is redeemable for $1 USDC upon resolution
type Token struct {
	ID      string `json:"id"`
	TokenID string `json:"token_id"`
	Outcome string `json:"outcome"`
	Price   string `json:"price"`
	Winner  *bool  `json:"winner"`
}

// Event represents a collection of related markets
type Event struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Description string     `json:"description"`
	StartDate   *time.Time `json:"startDate"`
	EndDate     *time.Time `json:"endDate"`
	Image       string     `json:"image"`
	Icon        string     `json:"icon"`
	Active      bool       `json:"active"`
	Closed      bool       `json:"closed"`
	Archived    bool       `json:"archived"`
	Featured    bool       `json:"featured"`

	// Trading metrics
	Volume     float64 `json:"volume"`
	Volume24hr float64 `json:"volume24hr"`
	Liquidity  float64 `json:"liquidity"`

	// Event features
	CommentsEnabled       bool `json:"commentsEnabled"`
	NegRisk               bool `json:"negRisk"`
	AutomaticallyResolved bool `json:"automaticallyResolved"`
	CYOM                  bool `json:"cyom"`

	// Related data
	Markets    []Market   `json:"markets"`
	Series     *Series    `json:"series"`
	Categories []Category `json:"categories"`
	Tags       []Tag      `json:"tags"`

	// Recurrence
	Recurrence string `json:"recurrence"`

	// Metadata
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// Series represents a series of related events
type Series struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Active      bool   `json:"active"`
}

// Tag represents an event tag
type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Category represents a market category/tag
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// MarketsParams represents query parameters for listing markets
type MarketsParams struct {
	Limit     int    `json:"limit,omitempty"`
	Offset    int    `json:"offset,omitempty"`
	Order     string `json:"order,omitempty"`
	Ascending bool   `json:"ascending,omitempty"`

	// Filters
	Active   *bool  `json:"active,omitempty"`
	Closed   *bool  `json:"closed,omitempty"`
	Archived *bool  `json:"archived,omitempty"`
	Slug     string `json:"slug,omitempty"`
	EventID  string `json:"event_id,omitempty"`
	TagID    string `json:"tag_id,omitempty"`
}

// EventsParams represents query parameters for listing events
type EventsParams struct {
	// Pagination
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`

	// Sorting
	Order     string `json:"order,omitempty"`
	Ascending bool   `json:"ascending,omitempty"`

	// Basic filters
	ID       []string `json:"id,omitempty"`
	Slug     []string `json:"slug,omitempty"`
	Active   *bool    `json:"active,omitempty"`
	Closed   *bool    `json:"closed,omitempty"`
	Archived *bool    `json:"archived,omitempty"`

	// Advanced filters
	TagID        *int   `json:"tag_id,omitempty"`
	ExcludeTagID []int  `json:"exclude_tag_id,omitempty"`
	RelatedTags  *bool  `json:"related_tags,omitempty"`
	Featured     *bool  `json:"featured,omitempty"`
	CYOM         *bool  `json:"cyom,omitempty"`
	Recurrence   string `json:"recurrence,omitempty"`

	// Date filters
	StartDateMin *time.Time `json:"start_date_min,omitempty"`
	StartDateMax *time.Time `json:"start_date_max,omitempty"`
	EndDateMin   *time.Time `json:"end_date_min,omitempty"`
	EndDateMax   *time.Time `json:"end_date_max,omitempty"`
}

// GetEventParams represents query parameters for getting a single event by ID
type GetEventParams struct {
	IncludeChat     *bool `json:"include_chat,omitempty"`
	IncludeTemplate *bool `json:"include_template,omitempty"`
}

// GetMarketParams represents query parameters for getting a single market by ID
type GetMarketParams struct {
	IncludeTag *bool `json:"include_tag,omitempty"`
}

// Comment represents a comment on a market, event, or series
type Comment struct {
	ID               string       `json:"id"`
	Body             string       `json:"body"`
	ParentEntityType string       `json:"parentEntityType"`
	ParentEntityID   string       `json:"parentEntityID"`
	UserAddress      string       `json:"userAddress"`
	CreatedAt        *time.Time   `json:"createdAt"`
	Profile          *UserProfile `json:"profile"`
	Reactions        []Reaction   `json:"reactions"`
	ReportCount      int          `json:"reportCount"`
	ReactionCount    int          `json:"reactionCount"`
}

// UserProfile represents user profile information
type UserProfile struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
	Website     string `json:"website"`
	Verified    bool   `json:"verified"`
}

// Reaction represents a reaction to a comment
type Reaction struct {
	ID          string     `json:"id"`
	Type        string     `json:"type"`
	UserAddress string     `json:"userAddress"`
	CreatedAt   *time.Time `json:"createdAt"`
}

// CommentsParams represents query parameters for listing comments
type CommentsParams struct {
	// Pagination
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`

	// Sorting
	Order     string `json:"order,omitempty"`
	Ascending bool   `json:"ascending,omitempty"`

	// Filters
	ParentEntityType string `json:"parent_entity_type,omitempty"` // "Event", "Series", "market"
	ParentEntityID   *int   `json:"parent_entity_id,omitempty"`
	GetPositions     *bool  `json:"get_positions,omitempty"`
	HoldersOnly      *bool  `json:"holders_only,omitempty"`
}

// SearchParams represents query parameters for search
type SearchParams struct {
	// Required
	Q string `json:"q"` // Search query

	// Pagination
	Page         int `json:"page,omitempty"`
	LimitPerType int `json:"limit_per_type,omitempty"`

	// Sorting
	Sort      string `json:"sort,omitempty"`
	Ascending *bool  `json:"ascending,omitempty"`

	// Filters and options
	Cache             *bool    `json:"cache,omitempty"`
	EventsStatus      string   `json:"events_status,omitempty"`
	EventsTag         []string `json:"events_tag,omitempty"`
	KeepClosedMarkets *int     `json:"keep_closed_markets,omitempty"`
	SearchTags        *bool    `json:"search_tags,omitempty"`
	SearchProfiles    *bool    `json:"search_profiles,omitempty"`
	Recurrence        string   `json:"recurrence,omitempty"`
	ExcludeTagID      []int    `json:"exclude_tag_id,omitempty"`
	Optimized         *bool    `json:"optimized,omitempty"`
}

// SearchResults represents the unified search response
type SearchResults struct {
	Events     []Event       `json:"events"`
	Tags       []Tag         `json:"tags"`
	Profiles   []UserProfile `json:"profiles"`
	Pagination Pagination    `json:"pagination"`
}

// Pagination represents pagination information in search results
type Pagination struct {
	HasMore      bool `json:"hasMore"`
	TotalResults int  `json:"totalResults"`
}

// LiveVolume represents live volume data for an event
type LiveVolume struct {
	Total   float64        `json:"total"`
	Markets []MarketVolume `json:"markets"`
}

// MarketVolume represents volume data for a specific market
type MarketVolume struct {
	Market string  `json:"market"` // Market address/ID
	Value  float64 `json:"value"`  // Volume value
}

// APIError represents an error response from the Polymarket API
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}
