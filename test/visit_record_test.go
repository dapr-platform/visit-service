package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var dbTimeFormat = "2006-01-02T15:04:05"

func TestVisitRecordAPI(t *testing.T) {
	client := NewTestClient()
	var recordID string

	t.Run("Create Visit Record", func(t *testing.T) {
		now := time.Now()
		record := map[string]interface{}{
			"patient_id":       "test-patient-1",
			"relative_id":      "test-relative-1",
			"visit_start_time": now.Format(dbTimeFormat),
			"visit_end_time":   now.Add(time.Hour).Format(dbTimeFormat),
			"visitor_name":     "Test Visitor",
			"visitor_phone":    "13800138000",
			"visitor_id_card":  "110101199001011234",
			"relationship":     "父母",
			"check_status":     0, // 未审核
			"status":           0, // 正常
		}

		resp, err := client.DoRequest(t, "POST", "/visit-record", record)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Create Visit Record Response: %+v\n", result)

		data := result["data"].(map[string]interface{})
		recordID = data["id"].(string)
		assert.Equal(t, "Test Visitor", data["visitor_name"])
	})

	t.Run("Get Visit Record List", func(t *testing.T) {
		resp, err := client.DoRequest(t, "GET", "/visit-record", nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Get Visit Record List Response: %+v\n", result)
		assert.NotNil(t, result["data"])
	})

	t.Run("Get Visit Record Info", func(t *testing.T) {
		resp, err := client.DoRequest(t, "GET", "/visit-record-info", nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Get Visit Record Info Response: %+v\n", result)
		assert.NotNil(t, result["data"])
	})

	t.Run("Delete Visit Record", func(t *testing.T) {
		resp, err := client.DoRequest(t, "DELETE", "/visit-record/"+recordID, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
}
