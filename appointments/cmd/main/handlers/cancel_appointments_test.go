package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgar-care/appointments/cmd/main/handlers"
	//_ "github.com/edgar-care/appointments/cmd/main/lib"
	"github.com/go-chi/chi/v5"
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

func (m *MockEdgarLib) CancelRdv(id, reason string) edgarlib.UpdateRdv {
	args := m.Called(id, reason)
	return args.Get(0).(edgarlib.UpdateRdv)
}

func TestCancelRdv_Success(t *testing.T) {
	mockAuth := new(MockAuthLib)
	mockEdgar := new(MockEdgarLib)
	handlers.AuthLib = mockAuth
	handlers.EdgarLib = mockEdgar

	mockAuth.On("AuthMiddlewareDoctor", mock.Anything, mock.Anything).Return(authlib.DoctorID{ID: "doctor123"})
	mockEdgar.On("CancelRdv", "appointment123", "Patient request").Return(edgarlib.UpdateRdv{Reason: "Patient request"})

	reqBody := `{"reason": "Patient request"}`
	req := httptest.NewRequest(http.MethodPost, "/appointments/appointment123/cancel", bytes.NewBufferString(reqBody))
	req = chi.NewRouteContext().WithURLParam("id", "appointment123").WithRequest(req)
	w := httptest.NewRecorder()

	handlers.CancelRdv(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "Patient request", response["reason"])
}

func TestCancelRdv_Unauthenticated(t *testing.T) {
	mockAuth := new(MockAuthLib)
	handlers.AuthLib = mockAuth

	mockAuth.On("AuthMiddlewareDoctor", mock.Anything, mock.Anything).Return(authlib.DoctorID{ID: ""})

	req := httptest.NewRequest(http.MethodPost, "/appointments/appointment123/cancel", nil)
	w := httptest.NewRecorder()

	handlers.CancelRdv(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "Not authenticated", response["message"])
}

func TestCancelRdv_AccountDisabled(t *testing.T) {
	mockAuth := new(MockAuthLib)
	handlers.AuthLib = mockAuth

	mockAuth.On("AuthMiddlewareDoctor", mock.Anything, mock.Anything).Return(authlib.DoctorID{Code: 409, Err: errors.New("Account disabled")})

	req := httptest.NewRequest(http.MethodPost, "/appointments/appointment123/cancel", nil)
	w := httptest.NewRecorder()

	handlers.CancelRdv(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusConflict, resp.StatusCode)
	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "Account disabled", response["message"])
}

func TestCancelRdv_CancelError(t *testing.T) {
	mockAuth := new(MockAuthLib)
	mockEdgar := new(MockEdgarLib)
	handlers.AuthLib = mockAuth
	handlers.EdgarLib = mockEdgar

	mockAuth.On("AuthMiddlewareDoctor", mock.Anything, mock.Anything).Return(authlib.DoctorID{ID: "doctor123"})
	mockEdgar.On("CancelRdv", "appointment123", "Patient request").Return(edgarlib.UpdateRdv{Err: errors.New("Cancel error"), Code: 500})

	reqBody := `{"reason": "Patient request"}`
	req := httptest.NewRequest(http.MethodPost, "/appointments/appointment123/cancel", bytes.NewBufferString(reqBody))
	req = chi.NewRouteContext().WithURLParam("id", "appointment123").WithRequest(req)
	w := httptest.NewRecorder()

	handlers.CancelRdv(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "Cancel error", response["message"])
}
