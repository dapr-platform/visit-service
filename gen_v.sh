dp-cli gen --connstr "postgresql://things:things2024@ali4:37432/thingsdb?sslmode=disable" \
--tables=v_family_member_info,v_patient_relative_info,v_bed_info,v_patient_info,v_visit_record_info,v_camera_info,v_live_record_info,v_head_display_info,v_visit_record_live_info --model_naming "{{ toUpperCamelCase ( replace . \"v_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"v_\" \"\") }}" \
--module visit-service \
--api R
