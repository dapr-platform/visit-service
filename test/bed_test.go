package test

import (
	"encoding/json"
	"fmt"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBedAPI(t *testing.T) {
	client := NewTestClient()
	var bedID string

	// 先创建一个病房用于测试
	var wardID string
	t.Run("Setup Ward", func(t *testing.T) {
		ward := map[string]interface{}{
			"name": "Test Ward for Bed",
			"type": 0,
			"status": 0,
		}
		
		resp, err := client.DoRequest(t, "POST", "/ward", ward)
		require.NoError(t, err)
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		wardID = result["data"].(map[string]interface{})["id"].(string)
	})

	t.Run("Create Bed", func(t *testing.T) {
		bed := map[string]interface{}{
			"ward_id": wardID,
			"bed_no": "B001",
			"type": 0, // 普通床位
			"status": 0, // 空闲
		}
		
		resp, err := client.DoRequest(t, "POST", "/bed", bed)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Create Bed Response: %+v\n", result)
		
		data := result["data"].(map[string]interface{})
		bedID = data["id"].(string)
		assert.Equal(t, "B001", data["bed_no"])
	})

	t.Run("Get Bed List", func(t *testing.T) {
		resp, err := client.DoRequest(t, "GET", "/bed", nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Get Bed List Response: %+v\n", result)
		assert.NotNil(t, result["data"])
	})

	t.Run("Get Bed Info", func(t *testing.T) {
		resp, err := client.DoRequest(t, "GET", "/bed-info", nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Get Bed Info Response: %+v\n", result)
		assert.NotNil(t, result["data"])
	})

	t.Run("Delete Bed", func(t *testing.T) {
		resp, err := client.DoRequest(t, "DELETE", "/bed/"+bedID, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	// 清理测试数据
	t.Run("Cleanup Ward", func(t *testing.T) {
		resp, err := client.DoRequest(t, "DELETE", "/ward/"+wardID, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
} 