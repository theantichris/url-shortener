package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	js "github.com/theantichris/url-shortener/serializer/json"
	ms "github.com/theantichris/url-shortener/serializer/msgpack"
	"github.com/theantichris/url-shortener/shortener"
)

// RedirectHandler defines the contract for a handler on the HTTP layer.
type RedirectHandler interface {
	Index(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	redirectService shortener.RedirectService
}

// NewHandler creates and returns a new handler.
func NewHandler(redirectService shortener.RedirectService) RedirectHandler {
	return &handler{
		redirectService,
	}
}

func writeResponse(w http.ResponseWriter, contentType string, statusCode int, body []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)

	_, err := w.Write(body)
	if err != nil {
		log.Println(err) // TODO: add proper logging
	}
}

func writeError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func (h *handler) serializer(contentType string) shortener.RedirectSerializer {
	if contentType == "application/x-msgpack" {
		return &ms.Redirect{}
	}

	return &js.Redirect{}
}

func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	writeResponse(w, contentType, http.StatusOK, []byte("It works!"))
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	redirect, err := h.redirectService.Find(code)
	if err != nil {
		if errors.Cause(err) == shortener.ErrRedirectNotFound {
			writeError(w, http.StatusNotFound)
			return
		}

		writeError(w, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, redirect.URL, http.StatusMovedPermanently)
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusInternalServerError)
		return
	}

	redirect, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		writeError(w, http.StatusInternalServerError)
		return
	}

	err = h.redirectService.Store(redirect)
	if err != nil {
		if errors.Cause(err) == shortener.ErrRedirectInvalid {
			writeError(w, http.StatusBadRequest)
			return
		}

		fmt.Println("failed to store redirect", err) // TODO
		writeError(w, http.StatusInternalServerError)
		return
	}

	responseBody, err := h.serializer(contentType).Encode(redirect)
	if err != nil {
		writeError(w, http.StatusInternalServerError)
		return
	}

	writeResponse(w, contentType, http.StatusCreated, responseBody)
}
