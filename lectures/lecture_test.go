package lectures_test

import (
	"chee-go-backend/test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLectures(t *testing.T) {
	r, _ := test.InitTest()

	req, _ := http.NewRequest("GET", "/api/lectures", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
