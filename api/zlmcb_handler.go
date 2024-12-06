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
