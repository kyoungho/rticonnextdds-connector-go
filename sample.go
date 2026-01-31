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

// Samples is a sequence of data samples received from a DDS DataReader.
//
// After calling input.Read() or input.Take(), the Samples collection contains
// the actual data values. Use GetLength() to determine how many samples are
// available, then access individual samples by index using the type-specific
// getter methods (GetString, GetInt32, GetFloat64, etc.).
//
// Sample indices are 0-based. Always check input.Infos.IsValid(i) before
// accessing sample data at index i.
type Samples struct {
	input *Input
}

// getNumber is a function to return a number in double from a sample
func (samples *Samples) getNumber(index int, fieldName string, retVal *C.double) error {
	if samples == nil {
		return errors.New("samples is null")
	}

	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	retcode := int(C.RTI_Connector_get_number_from_sample(unsafe.Pointer(samples.input.connector.native), retVal, samples.input.nameCStr, C.int(index+1), fieldNameCStr))
	return checkRetcode(retcode)
}

/*******************
* Public Functions *
*******************/

// GetLength returns the number of samples in the collection.
//
// This should be called after input.Read() or input.Take() to determine
// how many samples are available for processing.
//
// Returns:
//   - int: Number of samples (0 if no samples available)
//   - error: Non-nil if the operation fails
func (samples *Samples) GetLength() (int, error) {
	if samples == nil {
		return 0, errors.New("samples is null")
	}

	var retVal C.double
	retcode := int(C.RTI_Connector_get_sample_count(unsafe.Pointer(samples.input.connector.native), samples.input.nameCStr, &retVal))
	err := checkRetcode(retcode)
	return int(retVal), err
}

// GetUint8 retrieves a uint8 value from a specific field in a sample.
//
// Parameters:
//   - index: The index of the sample (0-based, use GetLength() to get valid range)
//   - fieldName: The name of the field to retrieve (must match XML type definition)
//
// Returns:
//   - uint8: The field value as an unsigned 8-bit integer
//   - error: Non-nil if the field doesn't exist, index is out of bounds, or type conversion fails
//
// Example:
//
//	value, err := samples.GetUint8(0, "status")
//	if err != nil {
//	    log.Printf("Failed to get status: %v", err)
//	}
func (samples *Samples) GetUint8(index int, fieldName string) (uint8, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return uint8(retVal), err
}

// GetUint16 is a function to retrieve a value of type uint16 from the samples
func (samples *Samples) GetUint16(index int, fieldName string) (uint16, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return uint16(retVal), err
}

// GetUint32 is a function to retrieve a value of type uint32 from the samples
func (samples *Samples) GetUint32(index int, fieldName string) (uint32, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return uint32(retVal), err
}

// GetUint64 is a function to retrieve a value of type uint64 from the samples
func (samples *Samples) GetUint64(index int, fieldName string) (uint64, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return uint64(retVal), err
}

// GetInt8 is a function to retrieve a value of type int8 from the samples
func (samples *Samples) GetInt8(index int, fieldName string) (int8, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return int8(retVal), err
}

// GetInt16 is a function to retrieve a value of type int16 from the samples
func (samples *Samples) GetInt16(index int, fieldName string) (int16, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return int16(retVal), err
}

// GetInt32 is a function to retrieve a value of type int32 from the samples
func (samples *Samples) GetInt32(index int, fieldName string) (int32, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return int32(retVal), err
}

// GetInt64 is a function to retrieve a value of type int64 from the samples
func (samples *Samples) GetInt64(index int, fieldName string) (int64, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return int64(retVal), err
}

// GetFloat32 is a function to retrieve a value of type float32 from the samples
func (samples *Samples) GetFloat32(index int, fieldName string) (float32, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return float32(retVal), err
}

// GetFloat64 is a function to retrieve a value of type float64 from the samples
func (samples *Samples) GetFloat64(index int, fieldName string) (float64, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return float64(retVal), err
}

// GetInt is a function to retrieve a value of type int from the samples
func (samples *Samples) GetInt(index int, fieldName string) (int, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return int(retVal), err
}

// GetUint is a function to retrieve a value of type uint from the samples
func (samples *Samples) GetUint(index int, fieldName string) (uint, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return uint(retVal), err
}

// GetByte is a function to retrieve a value of type byte from the samples
func (samples *Samples) GetByte(index int, fieldName string) (byte, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return byte(retVal), err
}

// GetRune is a function to retrieve a value of type rune from the samples
func (samples *Samples) GetRune(index int, fieldName string) (rune, error) {
	var retVal C.double
	err := samples.getNumber(index, fieldName, &retVal)
	return rune(retVal), err
}

// GetBoolean retrieves a boolean value from a specific field in a sample.
//
// Parameters:
//   - index: The index of the sample (0-based)
//   - fieldName: The name of the boolean field to retrieve
//
// Returns:
//   - bool: The field value as a boolean
//   - error: Non-nil if the field doesn't exist, index is out of bounds, or type conversion fails
//
// Example:
//
//	isActive, err := samples.GetBoolean(0, "enabled")
//	if err != nil {
//	    log.Printf("Failed to get enabled status: %v", err)
//	}
func (samples *Samples) GetBoolean(index int, fieldName string) (bool, error) {
	if samples == nil {
		return false, errors.New("samples is null")
	}

	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	var retVal C.int

	retcode := int(C.RTI_Connector_get_boolean_from_sample(unsafe.Pointer(samples.input.connector.native), &retVal, samples.input.nameCStr, C.int(index+1), fieldNameCStr))
	err := checkRetcode(retcode)

	return (retVal != 0), err
}

// GetString retrieves a string value from a specific field in a sample.
//
// Parameters:
//   - index: The index of the sample (0-based)
//   - fieldName: The name of the string field to retrieve
//
// Returns:
//   - string: The field value as a string (empty string if field is null)
//   - error: Non-nil if the field doesn't exist or index is out of bounds
//
// Example:
//
//	name, err := samples.GetString(0, "color")
//	if err != nil {
//	    log.Printf("Failed to get color: %v", err)
//	}
func (samples *Samples) GetString(index int, fieldName string) (string, error) {
	if samples == nil {
		return "", errors.New("samples is null")
	}

	fieldNameCStr := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldNameCStr))

	var retValCStr *C.char

	retcode := int(C.RTI_Connector_get_string_from_sample(unsafe.Pointer(samples.input.connector.native), &retValCStr, samples.input.nameCStr, C.int(index+1), fieldNameCStr))
	err := checkRetcode(retcode)
	if err != nil {
		return "", err
	}

	retValGoStr := C.GoString(retValCStr)
	C.RTI_Connector_free_string(retValCStr)

	return retValGoStr, nil
}

// GetJSON is a function to retrieve a slice of bytes of a JSON string from the samples
// GetJSON returns the complete JSON representation of a sample's data.
//
// This method serializes all fields of the specified sample into a JSON string,
// which is useful for debugging, logging, or converting to other data formats.
//
// Parameters:
//   - index: The index of the sample to serialize (0-based)
//
// Returns:
//   - string: JSON representation of the sample data
//   - error: Non-nil if the index is out of bounds or serialization fails
//
// Example:
//
//	jsonData, err := samples.GetJSON(0)
//	if err != nil {
//	    log.Printf("Failed to get JSON: %v", err)
//	} else {
//	    fmt.Printf("Sample data: %s\n", jsonData)
//	}
func (samples *Samples) GetJSON(index int) (string, error) {
	if samples == nil {
		return "", errors.New("samples is null")
	}

	var retValCStr *C.char

	retcode := int(C.RTI_Connector_get_json_sample(unsafe.Pointer(samples.input.connector.native), samples.input.nameCStr, C.int(index+1), &retValCStr))
	err := checkRetcode(retcode)
	if err != nil {
		return "", err
	}

	retValGoStr := C.GoString(retValCStr)
	C.RTI_Connector_free_string(retValCStr)

	return retValGoStr, err
}

// Get unmarshals a sample's data into a Go struct or interface.
//
// This method retrieves the JSON representation of the sample and unmarshals
// it into the provided interface. The target interface should have fields
// that match the XML structure (use json tags for field mapping if needed).
//
// Parameters:
//   - index: The index of the sample to retrieve (0-based)
//   - v: Pointer to the target struct/interface to unmarshal into
//
// Returns:
//   - error: Non-nil if the index is out of bounds or unmarshaling fails
//
// Example:
//
//	type ShapeType struct {
//	    Color string `json:"color"`
//	    X     int32  `json:"x"`
//	    Y     int32  `json:"y"`
//	}
//
//	var shape ShapeType
//	err := samples.Get(0, &shape)
//	if err != nil {
//	    log.Printf("Failed to unmarshal sample: %v", err)
//	}
func (samples *Samples) Get(index int, v interface{}) error {
	jsonData, err := samples.GetJSON(index)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(jsonData), &v)
}
