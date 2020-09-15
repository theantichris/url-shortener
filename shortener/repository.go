package shortener

// RedirectRepository defines the contract for integration with the datastore.
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
