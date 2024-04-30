package handler

import (
	data "backend/data"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testUserService struct{}

func (t *testUserService) GetUser(uuid string) (*data.User, error) {
	return &data.User{Uuid: uuid, Email: "email@meail.com", Name: "randomname"}, nil
}

func (t *testUserService) Create(user *data.User) error {
	user.Uuid = "created"
	return nil
}

func (t *testUserService) UpdateName(user *data.User) error {
	return nil
}

func (t *testUserService) UpdateEmail(user *data.User) error {
	return nil
}

func TestGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		get(&testUserService{}, "random", w)
	}))
	defer server.Close()
	res, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Error occured at performing request, err: %v", err)
	}

	user := &data.User{}
	body := make([]byte, 256)
	n, err := res.Body.Read(body)

	if err != nil && err != io.EOF {
		t.Fatalf("Error occured during reading response body, err: %v", err)
	}
	err = res.Body.Close()
	if err != nil {
		t.Fatalf("Error during closing response body, actual err: %v", err)
	}
	err = json.Unmarshal(body[:n], user)
	if err != nil {
		t.Fatalf("Error during unmarshaling response body, actual err: %v", err)
	}

}
