package members

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Boobuh/MemberClub/storage"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"
)

var (
	buf        bytes.Buffer
	testLogger = log.New(&buf, "testLogger: ", log.Lshortfile)
)

func TestHandler_Get(t *testing.T) {
	s := storage.NewStorage()
	router := mux.NewRouter()
	h := NewHandler(s, testLogger)
	router.HandleFunc("/members", h.Get)
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/members", nil)
	assert.NoError(t, err)
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
