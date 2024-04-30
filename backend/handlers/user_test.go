package handler

import (
	data "backend/data"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testUserService struct {
	GetUserFunc     func(uuid string) (*data.User, error)
	CreateFunc      func(user *data.User) error
	UpdateNameFunc  func(user *data.User) error
	UpdateEmailFunc func(user *data.User) error
}

func (t *testUserService) GetUser(uuid string) (*data.User, error) {
	if t.GetUserFunc != nil {
		return t.GetUserFunc(uuid)
	}
	return nil, fmt.Errorf("GetUserFunc not implemented")
}

func (t *testUserService) Create(user *data.User) error {
	if t.CreateFunc != nil {
		return t.CreateFunc(user)
	}
	return fmt.Errorf("CreateFunc not implemented")
}

func (t *testUserService) UpdateName(user *data.User) error {
	if t.UpdateNameFunc != nil {
		return t.UpdateNameFunc(user)
	}
	return fmt.Errorf("UpdateNameFunc not implemented")
}

func (t *testUserService) UpdateEmail(user *data.User) error {
	if t.UpdateEmailFunc != nil {
		return t.UpdateEmailFunc(user)
	}
	return fmt.Errorf("UpdateEmailFunc not implemented")
}

func TestGet(t *testing.T) {
	t.Run("Should get user", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ts := &testUserService{}
			ts.GetUserFunc = func(uuid string) (*data.User, error) {
				return &data.User{Uuid: uuid, Name: "someName", Email: "email@email.com"}, nil
			}
			get(ts, "random", w)
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
	})
	t.Run("Should handle error on failed user retrieve from db", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ts := &testUserService{}
			ts.GetUserFunc = func(uuid string) (*data.User, error) {
				return nil, errors.New("random error")
			}
			get(ts, "random", w)
		}))
		defer server.Close()
		res, err := http.Get(server.URL)
		if err != nil {
			t.Fatalf("Error occured at performing request, err: %v", err)
		}
		fmt.Printf("status code is %d", res.StatusCode)
		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("Wrong response status code, actual: %d, expected: %d", res.StatusCode, http.StatusBadRequest)
		}
		errorMsg := &Error{}
		body := make([]byte, 256)
		n, err := res.Body.Read(body)
		if err != nil && err != io.EOF {
			t.Fatalf("Error occured during reading response body, err: %v", err)
		}
		err = res.Body.Close()
		if err != nil {
			t.Fatalf("Error during closing response body, actual err: %v", err)
		}
		err = json.Unmarshal(body[:n], errorMsg)
		if err != nil {
			t.Fatalf("Error during unmarshaling response body, actual err: %v", err)
		}
		if errorMsg.Code != http.StatusBadRequest {
			t.Fatalf("Wrong response status code in json response, actual: %d, expected: %d", errorMsg.Code, http.StatusBadRequest)
		}
		if errorMsg.Message != "Unable to retrieve user" {
			t.Fatalf("Wrong response message in json response, actual: %s, expected: %s", errorMsg.Message, "Unable to retrieve user")
		}
	})
}

func TestPost(t *testing.T) {
	t.Run("Should create user", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ts := &testUserService{}
			ts.CreateFunc = func(user *data.User) error {
				return nil
			}
			create(ts, w, r)
		}))
		defer server.Close()
		mockUser := &data.User{Name: "eloelo", Email: "email@email.com"}
		mockBody, err := json.Marshal(mockUser)
		if err != nil {
			t.Fatalf("Error marshaling mocked user to create, actual err: %v", err)
		}

		res, err := http.Post(server.URL, "application/json", bytes.NewReader(mockBody))
		if err != nil {
			t.Fatalf("Error occured at performing request, err: %v", err)
		}
		if res.StatusCode != http.StatusCreated {
			t.Fatalf("Wrong status code in response, expected: %d, actual: %d", http.StatusCreated, res.StatusCode)
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
		if user.Name != mockUser.Name {
			t.Fatalf("expected user name: %s, actual user name: %s", mockUser.Name, user.Name)
		}
		if user.Email != mockUser.Email {
			t.Fatalf("expected user Email: %s, actual user Email: %s", mockUser.Email, user.Email)
		}
	})

	t.Run("Should return error on user creation fail in db", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ts := &testUserService{}
			ts.CreateFunc = func(user *data.User) error {
				return errors.New("random")
			}
			create(ts, w, r)
		}))
		defer server.Close()
		mockUser := &data.User{Name: "eloelo", Email: "email@email.com"}
		mockBody, err := json.Marshal(mockUser)
		if err != nil {
			t.Fatalf("Error marshaling mocked user to create, actual err: %v", err)
		}

		res, err := http.Post(server.URL, "application/json", bytes.NewReader(mockBody))
		if err != nil {
			t.Fatalf("Error occured at performing request, err: %v", err)
		}

		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("Wrong status code, actual: %d, expected: %d", res.StatusCode, http.StatusBadRequest)
		}
		errorMsg := &Error{}
		body := make([]byte, 256)
		n, err := res.Body.Read(body)

		if err != nil && err != io.EOF {
			t.Fatalf("Error occured during reading response body, err: %v", err)
		}
		err = res.Body.Close()
		if err != nil {
			t.Fatalf("Error during closing response body, actual err: %v", err)
		}
		err = json.Unmarshal(body[:n], errorMsg)
		if err != nil {
			t.Fatalf("Error during unmarshaling response body, actual err: %v", err)
		}
		if errorMsg.Code != http.StatusBadRequest {
			t.Fatalf("Wrong response status code in json response, actual: %d, expected: %d", errorMsg.Code, http.StatusBadRequest)
		}
		if errorMsg.Message != "Error occured during creating user" {
			t.Fatalf("Wrong response message in json response, actual: %s, expected: %s", errorMsg.Message, "Error occured during creating user")
		}
	})

	t.Run("Should return proper error when body decode fail", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ts := &testUserService{}
			ts.CreateFunc = func(user *data.User) error {
				return nil
			}
			create(ts, w, r)
		}))
		defer server.Close()
		res, err := http.Post(server.URL, "application/json", bytes.NewReader([]byte{0}))
		if err != nil {
			t.Fatalf("Error occured at performing request, err: %v", err)
		}

		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("Wrong status code, actual: %d, expected: %d", res.StatusCode, http.StatusBadRequest)
		}
		errorMsg := &Error{}
		body := make([]byte, 256)
		n, err := res.Body.Read(body)

		if err != nil && err != io.EOF {
			t.Fatalf("Error occured during reading response body, err: %v", err)
		}
		err = res.Body.Close()
		if err != nil {
			t.Fatalf("Error during closing response body, actual err: %v", err)
		}
		err = json.Unmarshal(body[:n], errorMsg)
		if err != nil {
			t.Fatalf("Error during unmarshaling response body, actual err: %v", err)
		}
		if errorMsg.Code != http.StatusBadRequest {
			t.Fatalf("Wrong response status code in json response, actual: %d, expected: %d", errorMsg.Code, http.StatusInternalServerError)
		}
		if errorMsg.Message != GenericError.Message {
			t.Fatalf("Wrong response message in json response, actual: %s, expected: %s", errorMsg.Message, GenericError.Message)
		}
	})
}

func TestPut(t *testing.T) {
	t.Run("Should not update user when not authorized", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ts := &testUserService{}
			update(ts, "", w, r)
		}))
		defer server.Close()
		client := server.Client()
		req, err := http.NewRequest("PUT", server.URL, bytes.NewBuffer([]byte{}))
		if err != nil {
			t.Fatalf("Error occured during creating request, req: %v, actual err: %v", req, err)
		}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error occured during performing request, actual err: %v", err)
		}
		if res.StatusCode != http.StatusUnauthorized {
			t.Fatalf("Wrong status code in response, expected: %d, actual: %d", http.StatusUnauthorized, res.StatusCode)
		}
		errorMsg := &Error{}
		body := make([]byte, 256)
		n, err := res.Body.Read(body)

		if err != nil && err != io.EOF {
			t.Fatalf("Error occured during reading response body, err: %v", err)
		}
		err = res.Body.Close()
		if err != nil {
			t.Fatalf("Error during closing response body, actual err: %v", err)
		}
		err = json.Unmarshal(body[:n], errorMsg)
		if err != nil {
			t.Fatalf("Error during unmarshaling response body, actual err: %v", err)
		}
		if errorMsg.Code != http.StatusUnauthorized {
			t.Fatalf("Wrong json response field, expected code: %d, actual code: %d", http.StatusUnauthorized, errorMsg.Code)
		}
		if errorMsg.Message != "Unauthorized to perform this action" {
			t.Fatalf("Wrong json response field, expected message: %s, actual message: %s", errorMsg.Message, "Unauthorized to perform this action")
		}
	})

	t.Run("Should return error when body is not valid json", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ts := &testUserService{}
			update(ts, "not empty", w, r)
		}))
		defer server.Close()
		client := server.Client()
		req, err := http.NewRequest("PUT", server.URL, bytes.NewBuffer([]byte{0}))
		if err != nil {
			t.Fatalf("Error occured during creating request, req: %v, actual err: %v", req, err)
		}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error occured during performing request, actual err: %v", err)
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("Wrong status code in response, expected: %d, actual: %d", http.StatusBadRequest, res.StatusCode)
		}
		errorMsg := &Error{}
		body := make([]byte, 256)
		n, err := res.Body.Read(body)

		if err != nil && err != io.EOF {
			t.Fatalf("Error occured during reading response body, err: %v", err)
		}
		err = res.Body.Close()
		if err != nil {
			t.Fatalf("Error during closing response body, actual err: %v", err)
		}
		err = json.Unmarshal(body[:n], errorMsg)
		if err != nil {
			t.Fatalf("Error during unmarshaling response body, actual err: %v", err)
		}
		if errorMsg.Code != http.StatusBadRequest {
			t.Fatalf("Wrong json response field, expected code: %d, actual code: %d", http.StatusBadRequest, errorMsg.Code)
		}
		if errorMsg.Message != GenericError.Message {
			t.Fatalf("Wrong json response field, expected message: %s, actual message: %s", errorMsg.Message, GenericError.Message)
		}
	})

	t.Run("Should return error when requested user is not in db", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ts := &testUserService{}
			ts.GetUserFunc = func(uuid string) (*data.User, error) {
				return nil, errors.New("error during retrieving user from db")
			}
			update(ts, "not empty", w, r)
		}))
		defer server.Close()
		mockUser, err := json.Marshal(&data.User{})
		if err != nil {
			t.Fatalf("Error occured during marshaling mocked user, data: %v, err: %v", mockUser, err)
		}
		client := server.Client()
		req, err := http.NewRequest("PUT", server.URL, bytes.NewBuffer(mockUser))
		if err != nil {
			t.Fatalf("Error occured during creating request, req: %v, actual err: %v", req, err)
		}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error occured during performing request, actual err: %v", err)
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("Wrong status code in response, expected: %d, actual: %d", http.StatusBadRequest, res.StatusCode)
		}
		errorMsg := &Error{}
		body := make([]byte, 256)
		n, err := res.Body.Read(body)

		if err != nil && err != io.EOF {
			t.Fatalf("Error occured during reading response body, err: %v", err)
		}
		err = res.Body.Close()
		if err != nil {
			t.Fatalf("Error during closing response body, actual err: %v", err)
		}
		err = json.Unmarshal(body[:n], errorMsg)
		if err != nil {
			t.Fatalf("Error during unmarshaling response body, actual err: %v", err)
		}
		if errorMsg.Code != http.StatusBadRequest {
			t.Fatalf("Wrong json response field, expected code: %d, actual code: %d", http.StatusBadRequest, errorMsg.Code)
		}
		if errorMsg.Message != "User with provided id not exists" {
			t.Fatalf("Wrong json response field, expected message: %s, actual message: %s", errorMsg.Message, "User with provided id not exists")
		}
	})
	t.Run("Should return proper error when updating name and email failed", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ts := &testUserService{}
			ts.GetUserFunc = func(uuid string) (*data.User, error) {
				return &data.User{Uuid: uuid, Name: "nameToUpdate", Email: "email@toupdate.com"}, nil
			}
			ts.UpdateNameFunc = func(user *data.User) error {
				return errors.New("update error occured")
			}
			ts.UpdateEmailFunc = func(user *data.User) error {
				return errors.New("update error occured")
			}
			update(ts, "not empty", w, r)
		}))
		defer server.Close()
		mockUser, err := json.Marshal(&data.User{Name: "othername", Email: "otheremail"})
		if err != nil {
			t.Fatalf("Error occured during marshaling mocked user, data: %v, err: %v", mockUser, err)
		}
		client := server.Client()
		req, err := http.NewRequest("PUT", server.URL, bytes.NewBuffer(mockUser))
		if err != nil {
			t.Fatalf("Error occured during creating request, req: %v, actual err: %v", req, err)
		}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error occured during performing request, actual err: %v", err)
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("Wrong status code in response, expected: %d, actual: %d", http.StatusBadRequest, res.StatusCode)
		}
		errorMsg := &Error{}
		body := make([]byte, 256)
		n, err := res.Body.Read(body)

		if err != nil && err != io.EOF {
			t.Fatalf("Error occured during reading response body, err: %v", err)
		}
		err = res.Body.Close()
		if err != nil {
			t.Fatalf("Error during closing response body, actual err: %v", err)
		}
		err = json.Unmarshal(body[:n], errorMsg)
		if err != nil {
			t.Fatalf("Error during unmarshaling response body, actual err: %v", err)
		}
		if errorMsg.Code != http.StatusBadRequest {
			t.Fatalf("Wrong json response field, expected code: %d, actual code: %d", http.StatusBadRequest, errorMsg.Code)
		}
		errorMessages := strings.Split(errorMsg.Message, ", ")
		if len(errorMessages) != 2 {
			t.Fatalf("Expected two error messages, actual count: %d", len(errorMessages))
		}
		for _, msg := range errorMessages {
			if msg != "Updating name failed" && msg != "Updating email failed" {
				t.Fatalf("Error messages are not equal, actual: %v, expected: %v", errorMessages, []string{"Updating name failed", "Updating email failed"})
			}
		}
	})

	t.Run("Should return proper error when updating name failed", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ts := &testUserService{}
			ts.GetUserFunc = func(uuid string) (*data.User, error) {
				return &data.User{Uuid: uuid, Name: "nameToUpdate", Email: "email@toupdate.com"}, nil
			}
			ts.UpdateNameFunc = func(user *data.User) error {
				return errors.New("update error occured")
			}
			ts.UpdateEmailFunc = func(user *data.User) error {
				return nil
			}
			update(ts, "not empty", w, r)
		}))
		defer server.Close()
		mockUser, err := json.Marshal(&data.User{Name: "othername", Email: "otheremail"})
		if err != nil {
			t.Fatalf("Error occured during marshaling mocked user, data: %v, err: %v", mockUser, err)
		}
		client := server.Client()
		req, err := http.NewRequest("PUT", server.URL, bytes.NewBuffer(mockUser))
		if err != nil {
			t.Fatalf("Error occured during creating request, req: %v, actual err: %v", req, err)
		}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error occured during performing request, actual err: %v", err)
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("Wrong status code in response, expected: %d, actual: %d", http.StatusBadRequest, res.StatusCode)
		}
		errorMsg := &Error{}
		body := make([]byte, 256)
		n, err := res.Body.Read(body)

		if err != nil && err != io.EOF {
			t.Fatalf("Error occured during reading response body, err: %v", err)
		}
		err = res.Body.Close()
		if err != nil {
			t.Fatalf("Error during closing response body, actual err: %v", err)
		}
		err = json.Unmarshal(body[:n], errorMsg)
		if err != nil {
			t.Fatalf("Error during unmarshaling response body, actual err: %v", err)
		}
		if errorMsg.Code != http.StatusBadRequest {
			t.Fatalf("Wrong json response field, expected code: %d, actual code: %d", http.StatusBadRequest, errorMsg.Code)
		}
		if errorMsg.Message != "Updating name failed" {
			t.Fatalf("Wrong error message, actual: %s, expected: %s", errorMsg.Message, "Updating name failed")
		}
	})

	t.Run("Should return proper error when updating email failed", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ts := &testUserService{}
			ts.GetUserFunc = func(uuid string) (*data.User, error) {
				return &data.User{Uuid: uuid, Name: "nameToUpdate", Email: "email@toupdate.com"}, nil
			}
			ts.UpdateNameFunc = func(user *data.User) error {
				return nil
			}
			ts.UpdateEmailFunc = func(user *data.User) error {
				return errors.New("update error occured")
			}
			update(ts, "not empty", w, r)
		}))
		defer server.Close()
		mockUser, err := json.Marshal(&data.User{Name: "othername", Email: "otheremail"})
		if err != nil {
			t.Fatalf("Error occured during marshaling mocked user, data: %v, err: %v", mockUser, err)
		}
		client := server.Client()
		req, err := http.NewRequest("PUT", server.URL, bytes.NewBuffer(mockUser))
		if err != nil {
			t.Fatalf("Error occured during creating request, req: %v, actual err: %v", req, err)
		}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error occured during performing request, actual err: %v", err)
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("Wrong status code in response, expected: %d, actual: %d", http.StatusBadRequest, res.StatusCode)
		}
		errorMsg := &Error{}
		body := make([]byte, 256)
		n, err := res.Body.Read(body)

		if err != nil && err != io.EOF {
			t.Fatalf("Error occured during reading response body, err: %v", err)
		}
		err = res.Body.Close()
		if err != nil {
			t.Fatalf("Error during closing response body, actual err: %v", err)
		}
		err = json.Unmarshal(body[:n], errorMsg)
		if err != nil {
			t.Fatalf("Error during unmarshaling response body, actual err: %v", err)
		}
		if errorMsg.Code != http.StatusBadRequest {
			t.Fatalf("Wrong json response field, expected code: %d, actual code: %d", http.StatusBadRequest, errorMsg.Code)
		}
		if errorMsg.Message != "Updating email failed" {
			t.Fatalf("Wrong error message, actual: %s, expected: %s", errorMsg.Message, "Updating email failed")
		}
	})

	t.Run("Should update user name and email", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ts := &testUserService{}
			ts.GetUserFunc = func(uuid string) (*data.User, error) {
				return &data.User{Uuid: uuid, Name: "nameToUpdate", Email: "email@toupdate.com"}, nil
			}
			ts.UpdateNameFunc = func(user *data.User) error {
				user.Name = "othername"
				return nil
			}
			ts.UpdateEmailFunc = func(user *data.User) error {
				user.Email = "otheremail"
				return nil
			}
			update(ts, "not empty", w, r)
		}))
		defer server.Close()
		mockUser, err := json.Marshal(&data.User{Name: "othername", Email: "otheremail"})
		if err != nil {
			t.Fatalf("Error occured during marshaling mocked user, data: %v, err: %v", mockUser, err)
		}
		client := server.Client()
		req, err := http.NewRequest("PUT", server.URL, bytes.NewBuffer(mockUser))
		if err != nil {
			t.Fatalf("Error occured during creating request, req: %v, actual err: %v", req, err)
		}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error occured during performing request, actual err: %v", err)
		}
		if res.StatusCode != http.StatusOK {
			t.Fatalf("Wrong status code in response, expected: %d, actual: %d", http.StatusOK, res.StatusCode)
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
		if user.Name != "othername" {
			t.Fatalf("Wrong user name. Actual: %s, expected: %s", user.Name, "othername")
		}
		if user.Email != "otheremail" {
			t.Fatalf("Wrong user email. Actual: %s, expected: %s", user.Email, "otheremail")
		}
	})
}
