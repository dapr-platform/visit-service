package service

import (
	"net/http"
	"time"
	"visit-service/model"

	"github.com/dapr-platform/common"
)

func init() {
	common.RegisterUpsertBeforeHook("Patient_relative", UpsertPatient_relative)
}

func UpsertPatient_relative(r *http.Request, in any) (out any, err error) {
	relative := in.(model.Patient_relative)
	if relative.ID == "" {
		config, err := GetConfig(CONFIG_VISIT_REGISTER_AUTO_AUDIT)
		if err != nil {
			return nil, err
		}
		if config.ConfigValue == "1" {
			relative.CheckStatus = CHECK_STATUS_CHECKED
		} else {
			relative.CheckStatus = CHECK_STATUS_UNCHECKED
		}
		relative.ID = common.NanoId()
		relative.CreateTime = common.LocalTime(time.Now())
	}
	return relative, nil
}
