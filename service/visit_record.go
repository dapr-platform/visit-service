package service

import (
	"context"
	"net/http"
	"visit-service/model"

	"github.com/dapr-platform/common"
)

func init() {

	common.RegisterUpsertBeforeHook("Visit_record", UpsertVisit_record)
}

func UpsertVisit_record(r *http.Request, in any) (out any, err error) {
	return in, nil
}

func FindNearestVisitRecord(ctx context.Context, familyMemberID string) (*model.Visit_record, error) {
	qstr := model.Visit_record_FIELD_NAME_relative_id + "=" + familyMemberID + "&_order=-" + model.Visit_record_FIELD_NAME_visit_start_time
	record, err := common.DbGetOne[model.Visit_record](ctx, common.GetDaprClient(), model.Visit_recordTableInfo.Name, qstr)
	if err != nil {
		return nil, err
	}
	return record, nil
}
