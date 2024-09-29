package handler

import "net/http"

type RootHandler struct{}

func (h RootHandler) GetHomeView() http.Handler {
	return nil
}
