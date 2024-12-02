package test

import (
	"encoding/json"
	"fmt"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWardAPI(t *testing.T) {
	client := NewTestClient()
	var wardID string

	t.Run("Create Ward", func(t *testing.T) {
		ward := map[string]interface{}{
			"name": "Test Ward",
			"type": 0, // 普通病房
			"status": 0, // 空闲
		}
		
		resp, err := client.DoRequest(t, "POST", "/ward", ward)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Create Ward Response: %+v\n", result)
		
		data := result["data"].(map[string]interface{})
		wardID = data["id"].(string)
		assert.Equal(t, "Test Ward", data["name"])
	})

	t.Run("Get Ward List", func(t *testing.T) {
		resp, err := client.DoRequest(t, "GET", "/ward", nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Get Ward List Response: %+v\n", result)
		assert.NotNil(t, result["data"])
	})

	t.Run("Get Ward Page", func(t *testing.T) {
		resp, err := client.DoRequest(t, "GET", "/ward/page?_page=1&_page_size=10", nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Get Ward Page Response: %+v\n", result)
		assert.NotNil(t, result["data"])
	})

	t.Run("Batch Update Wards", func(t *testing.T) {
		wards := []map[string]interface{}{
			{
				"id": wardID,
				"name": "Updated Ward",
				"type": 0,
				"status": 1, // 占用
			},
		}
		
		resp, err := client.DoRequest(t, "POST", "/ward/batch-upsert", wards)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("Delete Ward", func(t *testing.T) {
		resp, err := client.DoRequest(t, "DELETE", "/ward/"+wardID, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
} 