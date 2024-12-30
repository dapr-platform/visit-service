package api

import "github.com/go-chi/chi/v5"

func InitRoute(r chi.Router) {
	InitBed_infoRoute(r)
	InitCamera_infoRoute(r)
	InitCameraRoute(r)
	InitFamily_member_infoRoute(r)
	InitHead_display_infoRoute(r)
	InitLive_record_infoRoute(r)
	InitPatient_infoRoute(r)
	InitPatientRoute(r)
	InitSystem_configRoute(r)
	InitVisit_recordRoute(r)
	InitVisit_record_infoRoute(r)
	InitVisit_scheduleRoute(r)
	InitWardRoute(r)
	InitCustomCamRoute(r)
	InitZlmCbHandler(r)
	InitDashboardRoute(r)
	InitVisit_scheduleExtRoute(r)
	InitUserRoute(r)
	InitBedRoute(r)
	InitHead_displayRoute(r)
	InitPatient_relativeRoute(r)
	InitPatient_relative_infoRoute(r)
	InitLive_recordRoute(r)
	InitVisit_record_live_infoRoute(r)
}
