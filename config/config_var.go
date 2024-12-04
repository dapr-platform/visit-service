package config

import "os"

var (
	ZLMEDIAKIT_API_URL = "http://zlmediakit/index/api"
	ZLMEDIAKIT_SECRET = "035c73f7-bb6b-4889-a715-d9eb2d1925cc"
	ALI_SMS_ACCESS_ID = ""
	ALI_SMS_ACCESS_SECRET = ""
	ALI_SMS_REGION = ""
	ALI_SMS_SIGN_NAME = ""
	ALI_SMS_TEMPLATE_VISIT_CHECK_CODE = ""
	ALI_SMS_TEMPLATE_VISIT_PROMPT_CODE = ""
)
func init() {
	if os.Getenv("ALI_SMS_ACCESS_ID") != "" {
		ALI_SMS_ACCESS_ID = os.Getenv("ALI_SMS_ACCESS_ID")
	}
	if os.Getenv("ALI_SMS_ACCESS_SECRET") != "" {
		ALI_SMS_ACCESS_SECRET = os.Getenv("ALI_SMS_ACCESS_SECRET")
	}
	if os.Getenv("ALI_SMS_REGION") != "" {
		ALI_SMS_REGION = os.Getenv("ALI_SMS_REGION")
	}
	if os.Getenv("ALI_SMS_SIGN_NAME") != "" {
		ALI_SMS_SIGN_NAME = os.Getenv("ALI_SMS_SIGN_NAME")
	}
	if os.Getenv("ALI_SMS_TEMPLATE_VISIT_CHECK_CODE") != "" {
		ALI_SMS_TEMPLATE_VISIT_CHECK_CODE = os.Getenv("ALI_SMS_TEMPLATE_VISIT_CHECK_CODE")
	}

	if val := os.Getenv("ZLMEDIAKIT_API_URL"); val != "" {
		ZLMEDIAKIT_API_URL = val
	}
	if val := os.Getenv("ZLMEDIAKIT_SECRET"); val != "" {
		ZLMEDIAKIT_SECRET = val
	}
}
