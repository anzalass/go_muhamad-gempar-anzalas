package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"testinggg/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockModel struct{}

func (m *MockModel) CreateUser(data model.User) *model.User {
	return nil
}
func (m *MockModel) LoginUser(data model.User) *model.User {
	return nil
}
func (m *MockModel) GetAllUser() []model.User {
	return nil
}
func (m *MockModel) GetUserByID(id uint64) *model.User {
	return nil
}
func (m *MockModel) DeleteUserByID(id uint64) *model.User {
	return nil
}
func (sm *MockModel) UpdateUserByID(data model.User) *model.User {
	return nil
}

type SuccessMockModel struct{}

func (sm *SuccessMockModel) CreateUser(data model.User) *model.User {
	return &data
}

func (sm *SuccessMockModel) LoginUser(data model.User) *model.User {
	return &data
}
func (sm *SuccessMockModel) GetAllUser() []model.User {
	return nil
}
func (sm *SuccessMockModel) GetUserByID(id uint64) *model.User {
	return &model.User{}
}
func (sm *SuccessMockModel) DeleteUserByID(id uint64) *model.User {
	return &model.User{}
}
func (sm *SuccessMockModel) UpdateUserByID(data model.User) *model.User {
	return &model.User{}
}

func TestCreateUser(t *testing.T) {

	t.Run("Bind Error", func(t *testing.T) {
		var mdl = MockModel{}
		var ctl = NewUserModelControllerInterface(&mdl)

		var e = echo.New()
		e.POST("/users", ctl.CreateUser())
		var req = httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(`"name" : "abc", "email":"anz@gmail.com" , "password" : "123"`)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"msg"`
		}
		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "Invalid user input", tmp.Message)
	})

	t.Run("Server Error", func(t *testing.T) {
		var mdl = MockModel{}
		var ctl = NewUserModelControllerInterface(&mdl)

		var e = echo.New()
		e.POST("/users", ctl.CreateUser())
		var req = httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(``)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}
		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "Error creating user", tmp.Message)
	})

	t.Run("Success", func(t *testing.T) {
		var mdl = SuccessMockModel{}
		var ctl = NewUserModelControllerInterface(&mdl)
		var e = echo.New()
		e.POST("/users", ctl.CreateUser())

		var req = httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(`{"id":1, "name":"abc", "email":"anz@gmail.com" , "password":"123"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}
		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		fmt.Println(resData)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "creating user", tmp.Message)
		assert.Equal(t, "abc", tmp.Data["name"])
	})

}

func TestGetAllUser(t *testing.T) {
	var model = SuccessMockModel{}
	var controller = NewUserModelControllerInterface(&model)

	var e = echo.New()
	e.GET("/users", controller.GetAllUser())

	var req = httptest.NewRequest(http.MethodGet, "/users", nil)
	var res = httptest.NewRecorder()

	e.ServeHTTP(res, req)

	type ResponseData struct {
		Data    map[string]any `json:"data"`
		Message string         `json:"message"`
	}

	var tmp = ResponseData{}
	var resData = json.NewDecoder(res.Result().Body)
	var err = resData.Decode(&tmp)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Nil(t, err)
	assert.Equal(t, "sukses", tmp.Message)
	assert.NotNil(t, tmp)
	assert.Nil(t, tmp.Data)

}

func TestLoginUser(t *testing.T) {

	t.Run("Bad Request", func(t *testing.T) {
		var mdl = MockModel{}
		var ctl = NewUserModelControllerInterface(&mdl)
		var e = echo.New()
		e.POST("/login", ctl.LoginUser())

		var req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader([]byte(`{"nama" : "adad"}`)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "invalid user", tmp.Message)

	})
	t.Run("Nill Login", func(t *testing.T) {
		var mdl = MockModel{}
		var ctl = NewUserModelControllerInterface(&mdl)
		var e = echo.New()
		e.POST("/login", ctl.LoginUser())

		var req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader([]byte(``)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "cant process login", tmp.Message)
		assert.Nil(t, err)

	})

	t.Run("Sukses Login", func(t *testing.T) {
		var mdl = SuccessMockModel{}
		var ctl = NewUserModelControllerInterface(&mdl)
		var e = echo.New()
		e.POST("/login", ctl.LoginUser())

		var req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader([]byte(`{"email":"buku", "password":"saya"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)
		email := "buku"
		password := "saya"

		assert.Equal(t, http.StatusOK, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "sukses", tmp.Message)
		// assert.Nil(t, err)
		assert.Equal(t, email, tmp.Data["email"])
		assert.Equal(t, password, tmp.Data["password"])

	})

}

func TestUpdateUser(t *testing.T) {
	t.Run("Bad Request", func(t *testing.T) {
		var mdl = MockModel{}
		var ctl = NewUserModelControllerInterface(&mdl)
		var e = echo.New()

		e.PUT("/users/:id", ctl.UpdateUserByID())

		var req = httptest.NewRequest(http.MethodPut, "/users/:id", bytes.NewReader([]byte(`{"namaa" : "kontol"}`)))
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "id invalid", tmp.Message)

	})
	t.Run("Sukses", func(t *testing.T) {
		var mdl = SuccessMockModel{}
		var ctl = NewUserModelControllerInterface(&mdl)
		var e = echo.New()

		e.PUT("/users/:id", ctl.UpdateUserByID())

		var req = httptest.NewRequest(http.MethodPut, "/users/1", bytes.NewReader([]byte(`{"name" : "kontol"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "sukses update", tmp.Message)

	})

	t.Run("Internal Error", func(t *testing.T) {
		var mdl = MockModel{}
		var ctl = NewUserModelControllerInterface(&mdl)
		var e = echo.New()

		e.PUT("/users/:id", ctl.UpdateUserByID())

		var req = httptest.NewRequest(http.MethodPut, "/users/1", bytes.NewReader([]byte(`{}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "invalid", tmp.Message)

	})
}

func TestDeleteUserById(t *testing.T) {
	t.Run("Server Error", func(t *testing.T) {
		var mdl = MockModel{}
		var ctl = NewUserModelControllerInterface(&mdl)
		var e = echo.New()

		e.PUT("/users/:id", ctl.DeleteUserByID())

		var req = httptest.NewRequest(http.MethodPut, "/users/sdadas", bytes.NewReader([]byte(`{}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "id not found", tmp.Message)
	})

	t.Run("Sukses Delete", func(t *testing.T) {
		var mdl = MockModel{}
		var ctl = NewUserModelControllerInterface(&mdl)
		var e = echo.New()

		e.PUT("/users/:id", ctl.DeleteUserByID())

		var req = httptest.NewRequest(http.MethodPut, "/users/1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}

		var tmp = ResponData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "Delete User", tmp.Message)
	})

}
