package api

import (
	"net/http"
	"visit-service/entity"
	"visit-service/service"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitCustomCamRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/camera/livestream/start", StartCamLiveStreamHandler)
	r.Post(common.BASE_CONTEXT+"/camera/livestream/preview", StartCamLiveStreamPreviewHandler)
}

// @Summary Start camera livestream preview
// @Description Start a camera livestream preview, only for preview
// @Tags 摄像头
// @Accept json
// @Param request body entity.CamLiveStreamRequest true "Camera livestream preview request"
// @Produce json
// @Success 200 {object} common.Response{data=entity.CamLiveStreamResponse} "Success response with stream URL"
// @Failure 400 {object} common.Response "Invalid request parameters"
// @Failure 500 {object} common.Response "Internal server error"
// @Router /camera/livestream/preview [post]
func StartCamLiveStreamPreviewHandler(w http.ResponseWriter, r *http.Request) {
	var req entity.CamLiveStreamRequest
	if err := common.ReadRequestBody(r, &req); err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("Invalid request body: "+err.Error()))
		return
	}
	if req.CameraID == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("camera_id is required"))
		return
	}
	// Start the camera livestream
	streamID, playbackPath, err := service.StartCamLiveStream(req.CameraID, req.DisableSaveMp4)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.Logger.Infof("streamID: %s, playbackPath: %s", streamID, playbackPath)
	response := entity.CamLiveStreamResponse{
		StreamUrlSuffix:    "/stream/live/" + streamID + ".live.flv",
		Mp4StreamUrlSuffix: "/stream/live/" + streamID + ".live.mp4",
		CameraID:           req.CameraID,
		Status:             1,
	}
	common.HttpSuccess(w, common.OK.WithData(response))
}

// @Summary Start camera livestream
// @Description Start a camera livestream for a specific room
// @Tags 摄像头
// @Accept json
// @Param request body entity.CamLiveStreamRequest true "Camera livestream request"
// @Produce json
// @Success 200 {object} common.Response{data=entity.CamLiveStreamResponse} "Success response with stream URL"
// @Failure 400 {object} common.Response "Invalid request parameters"
// @Failure 500 {object} common.Response "Internal server error"
// @Router /camera/livestream/start [post]
func StartCamLiveStreamHandler(w http.ResponseWriter, r *http.Request) {
	var req entity.CamLiveStreamRequest
	if err := common.ReadRequestBody(r, &req); err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("Invalid request body: "+err.Error()))
		return
	}

	// Validate required fields
	if req.UserID == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("user_id is required"))
		return
	}

	visitRecord, err := service.FindNearestVisitRecord(r.Context(), req.UserID)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	if visitRecord == nil {
		common.HttpResult(w, common.ErrService.AppendMsg("no visit record found"))
		return
	}

	if req.CameraID == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("camera_id is required"))
		return
	}

	// Start the camera livestream
	streamID, playbackPath, err := service.StartCamLiveStream(req.CameraID, req.DisableSaveMp4)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	playbackUrlSuffix := playbackPath + "/" + streamID + ".mp4"
	if req.DisableSaveMp4 {
		playbackUrlSuffix = ""
	}

	err = service.AddLiveRecord(r.Context(), visitRecord, req.UserID, req.CameraID, streamID, playbackUrlSuffix)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	response := entity.CamLiveStreamResponse{
		StreamUrlSuffix:    "/stream/live/" + streamID + ".live.flv",
		Mp4StreamUrlSuffix: "/stream/live/" + streamID + ".live.mp4",
		CameraID:           req.CameraID,
		Status:             1,
	}

	common.HttpSuccess(w, common.OK.WithData(response))
}
