dp-cli gen --connstr "postgresql://things:things2024@ali4:37432/thingsdb?sslmode=disable" \
--tables=o_ward,o_patient_relative,o_user,o_bed,o_head_display,o_patient,o_visit_schedule,o_schedule_camera,o_visit_record,o_system_config,o_camera,o_live_record --model_naming "{{ toUpperCamelCase ( replace . \"o_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"o_\" \"\") }}" \
--module visit-service \
--api RUDB

