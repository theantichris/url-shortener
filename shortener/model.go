package shortener

// Redirect contains the information for the URL shortener redirect.
type Redirect struct {
	// Code is the shortened code for the URL.
	Code string `json:"code" bson:"code" msgpack:"code"`
	// URL is the URL being shortened.
	URL string `json:"url" bson:"url" msgpack:"url" validate:"empty=false & format=url"`
	// CreatedAt is the date the Redirect was created.
	CreatedAt int64 `json:"created_at" bson:"created_at" msgpack:"created_at"`
}
