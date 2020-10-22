package shortener

import "testing"

type mockRepo struct {
	findCount  int
	storeCount int
	redirect   *Redirect
}

func (r *mockRepo) Find(code string) (*Redirect, error) {
	r.findCount++

	return &Redirect{URL: "http://example.com"}, nil
}

func (r *mockRepo) Store(redirect *Redirect) error {
	r.storeCount++

	r.redirect = redirect

	return nil
}

func TestFind(t *testing.T) {
	t.Run("it finds a redirect code", func(t *testing.T) {
		repo := mockRepo{}
		service := NewRedirectService(&repo)

		redirect, err := service.Find("abc123")

		if err != nil {
			t.Fatal("an error occurred:", err)
		}

		if repo.findCount != 1 {
			t.Errorf("Find() was called %d, expected %d calls", repo.findCount, 1)
		}

		if redirect.URL != "http://example.com" {
			t.Errorf("Incorrect URL was returned: got %q, want %q", redirect.URL, "http://example.com")
		}
	})
}

func TestStore(t *testing.T) {
	t.Run("it stores a redirect", func(t *testing.T) {
		repo := mockRepo{}
		service := NewRedirectService(&repo)

		redirect := &Redirect{URL: "http://example.com"}

		err := service.Store(redirect)

		if err != nil {
			t.Fatal("an error occurred:", err)
		}

		if repo.redirect.Code == "" {
			t.Error("Redirect code was not set")
		}

		if repo.redirect.CreatedAt == 0 {
			t.Error("Redirect created at date was not")
		}

		if repo.storeCount != 1 {
			t.Errorf("Store() was called %d, expected %d calls", repo.storeCount, 1)
		}
	})

	t.Run("it returns a validation error for missing URL", func(t *testing.T) {
		repo := mockRepo{}
		service := NewRedirectService(&repo)

		redirect := &Redirect{}

		err := service.Store(redirect)

		if err == nil {
			t.Error("Expected a validation error but received none")
		}
	})

	t.Run("it returns a validation error for invalid URL", func(t *testing.T) {
		repo := mockRepo{}
		service := NewRedirectService(&repo)

		redirect := &Redirect{URL: "example"}

		err := service.Store(redirect)

		if err == nil {
			t.Error("Expected a validation error but received none")
		}
	})
}
