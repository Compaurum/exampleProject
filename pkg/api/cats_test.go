package api

import (
	"bytes"
	"encoding/json"
	memory_storage "example/pkg/storage/memory-storage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModifyCatName(t *testing.T) {
	tests := []struct {
		Src      string
		Expected string
		Valid    bool
	}{
		{"furunkul", "FURUNKUL:Thursday", true},
		{"fedya", "FEDYA:Thursday", true},
		{"feJIFJEIFJEIffefeiiwefwefqwefweifwefqewf", "FEJIFJEIFJEIFFEFEIIWEFWEFQWEFWEIFWEFQEWF:Thursday", true},
		{"", "", false},
	}

	for _, tc := range tests {
		result, err := modifyCatName(tc.Src)
		if tc.Valid {
			require.NoError(t, err)
		}
		assert.EqualValues(t, tc.Expected, result, "wrong cat name")
	}

}

func TestCreateCat(t *testing.T) {
	storage := memory_storage.MemoryStorage{}
	h := NewHandler(storage)
	recorder := httptest.NewRecorder()

	requestBody := bytes.NewBuffer([]byte(`{
    	"name": "BuSya",
    	"age": 11
	}`))
	req := httptest.NewRequest(http.MethodGet, "/cat", requestBody)

	h.createCat(recorder, req)

	assert.EqualValues(t, http.StatusOK, recorder.Code)

	resp := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}

	err := json.NewDecoder(recorder.Body).Decode(&resp)
	require.NoError(t, err, "unable to parse json")
	assert.EqualValues(t, "BUSYA:Thursday", resp.Name)
	assert.EqualValues(t, 11, resp.Age)
}
