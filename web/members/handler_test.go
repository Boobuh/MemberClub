package members

import (
	"bytes"
	"encoding/json"
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
	assert.NotNil(t, s)
	//t.Run("error bad name", func(t *testing.T) {
	//	err := s.Save(Member{Name: "3214", Email: "test@gmail.com"})
	//	assert.Error(t, err)
	//})
	//t.Run("error bad email", func(t *testing.T) {
	//	err := s.Save(Member{Name: "fdsgdfg", Email: "t4214m"})
	//	assert.Error(t, err)
	//})
}
func TestHandler_Post(t *testing.T) {
	s := storage.NewStorage()
	router := mux.NewRouter()
	member := storage.Member{Name: "Test", Email: "test@gmail.com"}
	body, err := json.Marshal(member)
	assert.NoError(t, err)
	requestReader := bytes.NewReader(body)
	h := NewHandler(s, testLogger)
	router.HandleFunc("/members", h.Post)
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/members", requestReader)
	assert.NoError(t, err)
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
