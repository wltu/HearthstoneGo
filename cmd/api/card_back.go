package api

// CardBackSearch provides parameters for a single card back search
type cardBackSearch struct {
	url    string
	id     string
	locale string
}

// CardBackCollectionSearch provides parameters for a card backs collection search
type cardBackCollectionSearch struct {
	// Required Parameters
	url    string
	locale string

	// Optional Parameters
	optional map[string]string
}
