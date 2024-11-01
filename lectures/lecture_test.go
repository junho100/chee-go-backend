// package lectures_test

// import (
// 	"chee-go-backend/lectures"
// 	"chee-go-backend/test"
// 	"encoding/json"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestGetLectures(t *testing.T) {
// 	r, _ := test.InitTest()

// 	req, _ := http.NewRequest("GET", "/api/lectures", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	responseDataByte, _ := io.ReadAll(w.Body)
// 	var response lectures.GetLecturesResponse
// 	err := json.Unmarshal(responseDataByte, &response)
// 	if err != nil {
// 		t.Errorf("Failed to unmarshal response: %v", err)
// 	}

// 	assert.Equal(t, make([]lectures.GetLecturesResponseSubject, 0), response.Subjects)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }
