package api

import "github.com/go-chi/chi/v5"

func InitRoute(r chi.Router) {
	InitBed_infoRoute(r)
	InitCamera_infoRoute(r)
	InitFamily_member_infoRoute(r)
	InitHead_display_infoRoute(r)
	InitLive_record_infoRoute(r)
	InitPatient_infoRoute(r)
	InitSchedule_cameraRoute(r)
	InitSystem_configRoute(r)
	InitVisit_recordRoute(r)
	InitVisit_record_infoRoute(r)
	InitVisit_scheduleRoute(r)
	InitWardRoute(r)
	InitCustomCamRoute(r)
	InitZlmCbHandler(r)
	InitManualHandler(r)
}
