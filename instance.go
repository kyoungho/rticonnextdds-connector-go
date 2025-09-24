/*****************************************************************************
*   (c) 2020 Copyright, Real-Time Innovations.  All rights reserved.         *
*                                                                            *
* No duplications, whole or partial, manual or electronic, may be made       *
* without express written permission.  Any such copies, or revisions thereof,*
* must display this notice unaltered.                                        *
* This code contains trade secrets of Real-Time Innovations, Inc.            *
*                                                                            *
*****************************************************************************/

package rti

// #include "rticonnextdds-connector.h"
// #include <stdlib.h>
import "C"
import (
	"encoding/json"
	"errors"
	"unsafe"
)

/********
* Types *
*********/

// Instance represents a data sample that can be written to a DDS output.
//
// An Instance provides methods to set field values for a DDS sample before
// writing it to the data bus. Each Instance is associated with a specific
// Output and provides type-safe methods for setting various data types.
//
// Example usage:
//   instance := output.Instance()
//   instance.SetString("color", "BLUE")
//   instance.SetInt32("x", 100)
//   instance.SetInt32("y", 200)
//   output.Write()
type Instance struct {
	output *Output
}

/*******************
* Public Functions *
*******************/

// SetUint8 sets a uint8 value for the specified field in the instance.
//
// Parameters:
//   - fieldName: The name of the field to set (must match XML type definition)
//   - value: The uint8 value to set
//
// Returns:
//   - error: Non-nil if the field doesn't exist or type conversion fails
//
// Example:
//   err := instance.SetUint8("status", 255)
//   if err != nil {
//       log.Printf("Failed to set status: %v", err)
//   }
func (instance *Instance) SetUint8(fieldName string, value uint8) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetUint16 is a function to set a value of type uint16 into samples
func (instance *Instance) SetUint16(fieldName string, value uint16) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetUint32 is a function to set a value of type uint32 into samples
func (instance *Instance) SetUint32(fieldName string, value uint32) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetUint64 is a function to set a value of type uint64 into samples
func (instance *Instance) SetUint64(fieldName string, value uint64) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetInt8 is a function to set a value of type int8 into samples
func (instance *Instance) SetInt8(fieldName string, value int8) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetInt16 is a function to set a value of type int16 into samples
func (instance *Instance) SetInt16(fieldName string, value int16) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetInt32 is a function to set a value of type int32 into samples
func (instance *Instance) SetInt32(fieldName string, value int32) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetInt64 is a function to set a value of type int64 into samples
func (instance *Instance) SetInt64(fieldName string, value int64) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetUint is a function to set a value of type uint into samples
func (instance *Instance) SetUint(fieldName string, value uint) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetInt is a function to set a value of type int into samples
func (instance *Instance) SetInt(fieldName string, value int) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetFloat32 is a function to set a value of type float32 into samples
func (instance *Instance) SetFloat32(fieldName string, value float32) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetFloat64 is a function to set a value of type float64 into samples
func (instance *Instance) SetFloat64(fieldName string, value float64) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetString is a function that set a string to a fieldname of the samples
// SetString sets a string value for the specified field in the instance.
//
// Parameters:
//   - fieldName: The name of the string field to set
//   - value: The string value to set
//
// Returns:
//   - error: Non-nil if the field doesn't exist
//
// Example:
//   err := instance.SetString("color", "BLUE")
//   if err != nil {
//       log.Printf("Failed to set color: %v", err)
//   }
func (instance *Instance) SetString(fieldName string, value string) error {
	if instance == nil || instance.output == nil || instance.output.connector == nil {
		return errors.New("instance, output, or connector is null")
	}
	if fieldName == "" {
		return errors.New("fieldName cannot be empty")
	}

	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	valueCStr := C.CString(value)
	defer C.free(unsafe.Pointer(valueCStr))

	retcode := int(C.RTI_Connector_set_string_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, valueCStr))
	return checkRetcode(retcode)
}

// SetByte is a function to set a byte to a fieldname of the samples
func (instance *Instance) SetByte(fieldName string, value byte) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetRune is a function to set rune to a fieldname of the samples
func (instance *Instance) SetRune(fieldName string, value rune) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_set_number_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.double(value)))
	return checkRetcode(retcode)
}

// SetBoolean sets a boolean value for the specified field in the instance.
//
// Parameters:
//   - fieldName: The name of the boolean field to set
//   - value: The boolean value to set
//
// Returns:
//   - error: Non-nil if the field doesn't exist
//
// Example:
//   err := instance.SetBoolean("enabled", true)
//   if err != nil {
//       log.Printf("Failed to set enabled flag: %v", err)
//   }
func (instance *Instance) SetBoolean(fieldName string, value bool) error {
	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	intValue := 0
	if value {
		intValue = 1
	}
	retcode := int(C.RTI_Connector_set_boolean_into_samples(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, fieldNameCStr, C.int(intValue)))
	return checkRetcode(retcode)
}

// SetJSON sets all fields in the instance from a JSON byte array.
//
// This method allows setting multiple fields at once by providing a JSON
// representation of the data. The JSON structure must match the XML type definition.
//
// Parameters:
//   - blob: JSON data as a byte array
//
// Returns:
//   - error: Non-nil if the JSON is invalid or doesn't match the XML type structure
//
// Example:
//   jsonData := []byte(`{"color":"RED","x":10,"y":20}`)
//   err := instance.SetJSON(jsonData)
//   if err != nil {
//       log.Printf("Failed to set JSON: %v", err)
//   }
func (instance *Instance) SetJSON(blob []byte) error {
	jsonCStr := C.CString(string(blob))
	defer C.free(unsafe.Pointer(jsonCStr))

	retcode := int(C.RTI_Connector_set_json_instance(unsafe.Pointer(instance.output.connector.native), instance.output.nameCStr, jsonCStr))
	return checkRetcode(retcode)
}

// Set marshals a Go struct/interface into the instance.
//
// This method provides a convenient way to set all fields at once by passing
// a Go struct that matches the XML type definition. The struct is marshaled to JSON
// and then applied to the instance.
//
// Parameters:
//   - v: A struct or interface containing the data to set
//
// Returns:
//   - error: Non-nil if marshaling fails or the data doesn't match the XML type structure
//
// Example:
//   type ShapeType struct {
//       Color string `json:"color"`
//       X     int32  `json:"x"`
//       Y     int32  `json:"y"`
//   }
//   
//   shape := ShapeType{Color: "GREEN", X: 50, Y: 75}
//   err := instance.Set(shape)
//   if err != nil {
//       log.Printf("Failed to set instance: %v", err)
//   }
func (instance *Instance) Set(v interface{}) error {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return instance.SetJSON(jsonData)
}
