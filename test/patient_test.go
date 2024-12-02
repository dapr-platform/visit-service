package test

import (
	"encoding/json"
	"fmt"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPatientAPI(t *testing.T) {
	client := NewTestClient()

	t.Run("Create Patient", func(t *testing.T) {
		patient := map[string]interface{}{
			"name": "Test Patient",
			"hospital_no": "P001",
			"ward_id": "ward-1", 
			"bed_id": "bed-1",
			"status": 0,
		}
		
		resp, err := client.DoRequest(t, "POST", "/patient", patient)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Create Patient Response: %+v\n", result)
		assert.Equal(t, "Test Patient", result["data"].(map[string]interface{})["name"])
	})

	t.Run("Get Patient List", func(t *testing.T) {
		resp, err := client.DoRequest(t, "GET", "/patient", nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Get Patient List Response: %+v\n", result)
		assert.NotNil(t, result["data"])
	})
}