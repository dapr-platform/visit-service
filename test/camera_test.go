package test

import (
	"encoding/json"
	"fmt"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCameraAPI(t *testing.T) {
	client := NewTestClient()
	var cameraID string

	t.Run("Create Camera", func(t *testing.T) {
		camera := map[string]interface{}{
			"device_name": "Test Camera",
			"device_no": "CAM001",
			"device_type": 0, // 普通摄像头
			"location_type": 0, // 固定位置
			"manufacturer": "Test Manufacturer",
			"main_stream_url": "rtsp://test/main",
			"sub_stream_url": "rtsp://test/sub",
			"status": 0, // 正常
		}
		
		resp, err := client.DoRequest(t, "POST", "/camera", camera)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Create Camera Response: %+v\n", result)
		
		data := result["data"].(map[string]interface{})
		cameraID = data["id"].(string)
		assert.Equal(t, "Test Camera", data["device_name"])
	})

	t.Run("Get Camera List", func(t *testing.T) {
		resp, err := client.DoRequest(t, "GET", "/camera", nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Get Camera List Response: %+v\n", result)
		assert.NotNil(t, result["data"])
	})

	t.Run("Get Camera Info", func(t *testing.T) {
		resp, err := client.DoRequest(t, "GET", "/camera-info", nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		
		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		require.NoError(t, err)
		fmt.Printf("Get Camera Info Response: %+v\n", result)
		assert.NotNil(t, result["data"])
	})

	t.Run("Delete Camera", func(t *testing.T) {
		resp, err := client.DoRequest(t, "DELETE", "/camera/"+cameraID, nil)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
} 