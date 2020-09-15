package shortener

// RedirectService defines the contract for the service that finds and stores a Redirect.
type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
