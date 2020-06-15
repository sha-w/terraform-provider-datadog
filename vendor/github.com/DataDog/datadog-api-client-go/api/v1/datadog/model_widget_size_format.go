/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// WidgetSizeFormat Size of the widget.
type WidgetSizeFormat string

// List of WidgetSizeFormat
const (
	WIDGETSIZEFORMAT_SMALL  WidgetSizeFormat = "small"
	WIDGETSIZEFORMAT_MEDIUM WidgetSizeFormat = "medium"
	WIDGETSIZEFORMAT_LARGE  WidgetSizeFormat = "large"
)

func (v *WidgetSizeFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WidgetSizeFormat(value)
	for _, existing := range []WidgetSizeFormat{"small", "medium", "large"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WidgetSizeFormat", *v)
}

// Ptr returns reference to WidgetSizeFormat value
func (v WidgetSizeFormat) Ptr() *WidgetSizeFormat {
	return &v
}

type NullableWidgetSizeFormat struct {
	value *WidgetSizeFormat
	isSet bool
}

func (v NullableWidgetSizeFormat) Get() *WidgetSizeFormat {
	return v.value
}

func (v *NullableWidgetSizeFormat) Set(val *WidgetSizeFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetSizeFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetSizeFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetSizeFormat(val *WidgetSizeFormat) *NullableWidgetSizeFormat {
	return &NullableWidgetSizeFormat{value: val, isSet: true}
}

func (v NullableWidgetSizeFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetSizeFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}