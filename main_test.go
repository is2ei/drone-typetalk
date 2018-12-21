package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type postMessageResponseParam struct {
}

func TestPostMessage(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "", 400)
		}

		var requestParam PostMessageRequestParam
		err := json.NewDecoder(r.Body).Decode(&requestParam)
		if err != nil {
			http.Error(w, "", 400)
		}

		if requestParam.Message == "" {
			http.Error(w, "", 400)
		}

		if requestParam.Message != "hello world" {
			t.Error("Message should be 'hello world'")
		}
		if requestParam.ShowLinkMeta == true {
			t.Error("ShowLinkMeta should be false")
		}

		w.Header().Set("Content-Type", "application/json")

		var responseParam postMessageResponseParam
		json.NewEncoder(w).Encode(responseParam)
	}))
	defer ts.Close()

	p := &PostMessageRequestParam{
		Message:      "hello world",
		ShowLinkMeta: false,
	}

	{
		resp, err := PostMessage(ts.URL, "12345", "qweasdzxc", p)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if resp.StatusCode != 200 {
			t.Error("Response should be 200")
		}
	}
}
