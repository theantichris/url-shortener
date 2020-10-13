package shortener

import "testing"

type mockRepo struct {
	findCount  int
	storeCount int
}

func (r *mockRepo) Find(code string) (*Redirect, error) {
	r.findCount++

	return &Redirect{URL: "http://example.com"}, nil
}

func (r *mockRepo) Store(redirect *Redirect) error {
	r.storeCount++

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
