package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	edgarlib "github.com/edgar-care/edgarlib/v2/appointment"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthLib struct {
	mock.Mock
}

func (m *MockAuthLib) AuthMiddlewareDoctor(w http.ResponseWriter, req *http.Request) authlib.DoctorID {
	args := m.Called(w, req)
	return args.Get(0).(authlib.DoctorID)
}

type MockEdgarLib struct {
	mock.Mock
}

func (m *MockEdgarLib) CreateRdv(accountID, doctorID string, startDate, endDate int, reason string) edgarlib.CreateRdv {
	args := m.Called(accountID, doctorID, startDate, endDate, reason)
	return args.Get(0).(edgarlib.CreateRdv)
}

func CreateRdv_Success(t *testing.T) {
	mockAuth := new(MockAuthLib)
	mockEdgar := new(MockEdgarLib)
	handlers.AuthLib = mockAuth
	handlers.EdgarLib = mockEdgar

	mockAuth.On("AuthMiddlewareDoctor", mock.Anything, mock.Anything).Return(authlib.DoctorID{ID: "doctor123"})
	mockEdgar.On("CreateRdv", "", "doctor123", 1625097600, 1625101200, "").Return(edgarlib.CreateRdv{Rdv: edgarlib.Rdv{ID: "rdv123"}})

	reqBody := `{"start_date": 1625097600, "end_date": 1625101200}`
	req := httptest.NewRequest(http.MethodPost, "/appointments/create", bytes.NewBufferString(reqBody))
	w := httptest.NewRecorder()

	handlers.CreateRdv(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "rdv123", response["rdv"].(map[string]interface{})["id"])
}

func CreateRdv_Unauthenticated(t *testing.T) {
	mockAuth := new(MockAuthLib)
	handlers.AuthLib = mockAuth

	mockAuth.On("AuthMiddlewareDoctor", mock.Anything, mock.Anything).Return(authlib.DoctorID{ID: ""})

	req := httptest.NewRequest(http.MethodPost, "/appointments/create", nil)
	w := httptest.NewRecorder()

	handlers.CreateRdv(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "Not authenticated", response["message"])
}

func CreateRdv_AccountDisabled(t *testing.T) {
	mockAuth := new(MockAuthLib)
	AuthLib = mockAuth

	mockAuth.On("AuthMiddlewareDoctor", mock.Anything, mock.Anything).Return(authlib.DoctorID{Code: 409, Err: errors.New("Account disabled")})

	req := httptest.NewRequest(http.MethodPost, "/appointments/create", nil)
	w := httptest.NewRecorder()

	CreateRdv(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusConflict, resp.StatusCode)
	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "Account disabled", response["message"])
}

func CreateRdv_CreateError(t *testing.T) {
	mockAuth := new(MockAuthLib)
	mockEdgar := new(MockEdgarLib)
	AuthLib = mockAuth
	EdgarLib = mockEdgar

	mockAuth.On("AuthMiddlewareDoctor", mock.Anything, mock.Anything).Return(authlib.DoctorID{ID: "doctor123"})
	mockEdgar.On("CreateRdv", "", "doctor123", 1625097600, 1625101200, "").Return(edgarlib.CreateRdv{Err: errors.New("Create error"), Code: 500})

	reqBody := `{"start_date": 1625097600, "end_date": 1625101200}`
	req := httptest.NewRequest(http.MethodPost, "/appointments/create", bytes.NewBufferString(reqBody))
	w := httptest.NewRecorder()

	handlers.CreateRdv(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "Create error", response["message"])
}
