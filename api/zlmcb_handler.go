package api

import (
	"encoding/json"
	"net/http"
	"visit-service/service"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitZlmCbHandler(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/zlm/on-record-mp4", ZlmOnRecordMp4Handler)
	r.Post(common.BASE_CONTEXT+"/zlm/on-stream-not-found", ZlmOnStreamNotFoundHandler)
	r.Post(common.BASE_CONTEXT+"/zlm/on-stream-none-reader", ZlmOnStreamNoneReaderHandler)
}

/*
	{
	   "mediaServerId" : "your_server_id",
	   "app" : "live",
	   "file_name" : "15-53-02.mp4",
	   "file_path" : "/root/zlmediakit/httpRoot/__defaultVhost__/record/live/obs/2019-09-20/15-53-02.mp4",
	   "file_size" : 1913597,
	   "folder" : "/root/zlmediakit/httpRoot/__defaultVhost__/record/live/obs/",
	   "start_time" : 1568965982,
	   "stream" : "obs",
	   "time_len" : 11.0,
	   "url" : "record/live/obs/2019-09-20/15-53-02.mp4",
	   "vhost" : "__defaultVhost__"
	}
*/
type ZlmOnRecordMp4Req struct {
	MediaServerID string  `json:"mediaServerId"`
	App           string  `json:"app"`
	FileName      string  `json:"file_name"`
	FilePath      string  `json:"file_path"`
	FileSize      int64   `json:"file_size"`
	Folder        string  `json:"folder"`
	StartTime     int64   `json:"start_time"`
	Stream        string  `json:"stream"`
	TimeLen       float64 `json:"time_len"`
	URL           string  `json:"url"`
	VHost         string  `json:"vhost"`
}

// @Summary zlm on record mp4
// @Description zlm on record mp4 callback
// @Tags Zlm
// @Param stream_id query string true "stream_id"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param type query string false "type"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ward}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /zlm/on-record-mp4 [post]
func ZlmOnRecordMp4Handler(w http.ResponseWriter, r *http.Request) {
	var req ZlmOnRecordMp4Req
	common.ReadRequestBody(r, &req)
	common.Logger.Info("zlm on record mp4", "req", req)
	if req.Stream == "" || req.URL == "" || req.FileSize == 0 {
		common.Logger.Error("zlm on record mp4", "req", req)
	} else {
		service.StopLiveRecord(r.Context(), req.Stream, req.FileSize, req.URL)
	}
	resp := make(map[string]any)
	resp["code"] = 0
	resp["msg"] = "success"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

type ZlmOnStreamNotFoundReq struct {
	MediaServerID string `json:"mediaServerId"`
	App           string `json:"app"`
	ID            string `json:"id"`
	IP            string `json:"ip"`
	Params        string `json:"params"`
	Port          int    `json:"port"`
	Schema        string `json:"schema"`
	Stream        string `json:"stream"`
	VHost         string `json:"vhost"`
}

// @Summary zlm on stream not found
// @Description zlm on stream not found callback
// @Tags Zlm
// @Param stream_id query string true "stream_id"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param type query string false "type"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ward}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /zlm/on-stream-not-found [post]
func ZlmOnStreamNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	var req ZlmOnStreamNotFoundReq
	common.ReadRequestBody(r, &req)
	common.Logger.Info("zlm on stream not found", "req", req)
	err := service.OnStreamNotFound(r.Context(), req.Stream)
	if err != nil {
		common.Logger.Error("zlm on stream not found", "err", err)
	}
	// 流未找到事件，用户可以在此事件触发时，立即去拉流，这样可以实现按需拉流
	// 此事件对回复不敏感
	resp := make(map[string]any)
	resp["code"] = 0
	resp["msg"] = "success"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

type ZlmOnStreamNoneReaderReq struct {
	MediaServerID string `json:"mediaServerId"`
	App           string `json:"app"`
	Stream        string `json:"stream"`
	VHost         string `json:"vhost"`
}

// @Summary zlm on stream none reader
// @Description zlm on stream none reader callback
// @Tags Zlm
// @Param stream_id query string true "stream_id"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ward}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /zlm/on-stream-none-reader [post]
func ZlmOnStreamNoneReaderHandler(w http.ResponseWriter, r *http.Request) {
	var req ZlmOnStreamNoneReaderReq
	common.ReadRequestBody(r, &req)
	common.Logger.Info("zlm on stream none reader", "req", req)
	err := service.OnStreamNoneReader(r.Context(), req.Stream)
	if err != nil {
		common.Logger.Error("zlm on stream none reader", "err", err)
	}
	// 流无人观看事件，用户可以在这个事件触发时，选择是否关闭推流等
	// 此事件对回复不敏感
	resp := make(map[string]any)
	resp["code"] = 0
	resp["msg"] = "success"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
