package service

import (
	"context"
	"visit-service/model"

	"github.com/dapr-platform/common"
)

func FindPatientInfo(ctx context.Context, patientID string) (*model.Patient_info, error) {
	patient, err := common.DbGetOne[model.Patient_info](ctx, common.GetDaprClient(), model.Patient_infoTableInfo.Name, "id="+patientID)
	if err != nil {
		return nil, err
	}
	return patient, nil
}
