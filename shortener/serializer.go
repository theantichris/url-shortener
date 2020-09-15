package shortener

// RedirectSerializer defines the contract for serializing/deserializing a Redirect.
type RedirectSerializer interface {
	Decode(input []byte) (*Redirect, error)
	Encode(input *Redirect) ([]byte, error)
}
