package main

import (
	"log"
)

func main() {
	log.Println("Polymarket Go Library Example")
	
	// Create a new client
	client := NewClient()
	
	// Example: Get markets with parameters
	log.Println("\nFetching markets...")
	marketsParams := &MarketsParams{
		Limit:     5,
		Offset:    0,
		Ascending: true,
		Active:    boolPtr(true),
	}
	
	markets, err := client.GetMarkets(marketsParams)
	if err != nil {
		log.Fatalf("Error fetching markets: %v", err)
	}
	
	log.Printf("Found %d markets", len(markets))
	for i, market := range markets {
		log.Printf("Market %d: %s (ID: %s)", i+1, market.Question, market.ID)
	}
	
	// Example: Get events
	log.Println("\nFetching events...")
	eventsParams := &EventsParams{
		Limit:  3,
		Active: boolPtr(true),
	}
	
	events, err := client.GetEvents(eventsParams)
	if err != nil {
		log.Printf("Error fetching events: %v", err)
	} else {
		log.Printf("Found %d events", len(events))
		for i, event := range events {
			log.Printf("Event %d: %s (ID: %s)", i+1, event.Title, event.ID)
		}
	}
	
	// Example: Get a specific market by ID (if we have one)
	if len(markets) > 0 {
		log.Printf("\nFetching specific market: %s", markets[0].ID)
		market, err := client.GetMarket(markets[0].ID)
		if err != nil {
			log.Printf("Error fetching specific market: %v", err)
		} else {
			log.Printf("Market details: %s", market.Question)
		}
	}
	
	// Example: Get comments
	log.Println("\nFetching recent comments...")
	commentsParams := &CommentsParams{
		Limit:  5,
		Order:  "createdAt",
	}
	
	comments, err := client.GetComments(commentsParams)
	if err != nil {
		log.Printf("Error fetching comments: %v", err)
	} else {
		log.Printf("Found %d comments", len(comments))
		for i, comment := range comments {
			log.Printf("Comment %d: %s... (by %s)", i+1, 
				truncateString(comment.Body, 50), comment.UserAddress[:8])
		}
	}
	
	// Example: Search for events about "election"
	log.Println("\nSearching for 'election' related content...")
	searchParams := &SearchParams{
		Q:            "election",
		LimitPerType: 3,
		SearchTags:   boolPtr(true),
		SearchProfiles: boolPtr(false),
	}
	
	searchResults, err := client.Search(searchParams)
	if err != nil {
		log.Printf("Error performing search: %v", err)
	} else {
		log.Printf("Found %d events, %d tags, %d profiles", 
			len(searchResults.Events), len(searchResults.Tags), len(searchResults.Profiles))
		
		if len(searchResults.Events) > 0 {
			log.Printf("Sample event: %s", searchResults.Events[0].Title)
		}
		
		if len(searchResults.Tags) > 0 {
			log.Printf("Sample tag: %s", searchResults.Tags[0].Name)
		}
	}
}

// Helper function to create bool pointers
func boolPtr(b bool) *bool {
	return &b
}

// Helper function to create int pointers
func intPtr(i int) *int {
	return &i
}

// Helper function to truncate strings for display
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
