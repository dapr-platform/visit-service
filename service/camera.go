package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"visit-service/config"
	"visit-service/model"

	"github.com/dapr-platform/common"
)

var LOCATION_TYPE_MOVABLE = 1
var LOCATION_TYPE_FIXED = 0
var DEVICE_TYPE_VR = 2
var DEVICE_TYPE_NORMAL = 0
var DEVICE_TYPE_NORMAL_PTZ = 1

type AddStreamResponse struct {
	Code int `json:"code"`
	Data struct {
		Key string `json:"key"`
	} `json:"data"`
}

func FindMovableCameras(ctx context.Context) ([]model.Camera_info, error) {
	cameras, err := common.DbQuery[model.Camera_info](ctx, common.GetDaprClient(), model.Camera_infoTableInfo.Name, fmt.Sprintf("location_type=%d", LOCATION_TYPE_MOVABLE))
	if err != nil {
		return nil, err
	}
	return cameras, nil
}
func StartCamLiveStream(cameraID string, disableSaveMp4 bool) (string, error) {
	// 1. 获取摄像头信息
	cameras, err := common.DbQuery[model.Camera_info](context.Background(), common.GetDaprClient(), model.Camera_infoTableInfo.Name, fmt.Sprintf("id=%s", cameraID))
	if err != nil {
		return "", fmt.Errorf("failed to query camera info: %v", err)
	}
	if len(cameras) == 0 {
		return "", fmt.Errorf("camera not found: %s", cameraID)
	}
	camera := cameras[0]

	if camera.MainStreamURL == "" {
		return "", fmt.Errorf("camera main stream url is empty")
	}
	streamID := common.NanoId()
	// 2. 调用ZLMediaKit API添加流代理
	params := url.Values{}
	params.Set("secret", config.ZLMEDIAKIT_SECRET)
	params.Set("vhost", "__defaultVhost__")
	params.Set("app", "live")
	params.Set("stream", streamID)
	params.Set("url", camera.MainStreamURL)
	if camera.DeviceType == int32(DEVICE_TYPE_VR) {
		disableSaveMp4 = true //VR设备不保存mp4
	}
	if disableSaveMp4 {
		params.Set("enable_mp4", "0")
	} else {
		params.Set("enable_mp4", "1")
	}
	params.Set("enable_rtmp", "1")
	params.Set("enable_hls", "1")
	params.Set("enable_ts", "1")
	params.Set("enable_fmp4", "1")
	params.Set("enable_audio", "1")
	params.Set("mp4_max_second", "3600")
	params.Set("mp4_as_player", "0")
	params.Set("auto_close", "1")

	apiURL := fmt.Sprintf("%s/addStreamProxy?%s", config.ZLMEDIAKIT_API_URL, params.Encode())
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to call ZLMediaKit API: %v", err)
	}
	defer resp.Body.Close()

	var result AddStreamResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	if result.Code != 0 {
		return "", fmt.Errorf("ZLMediaKit API returned error code: %d", result.Code)
	}

	// 3. 返回流ID
	return streamID, nil
}
