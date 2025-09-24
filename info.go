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
	"strconv"
	"unsafe"
)

/********
* Types *
*********/

// Infos represents a collection of DDS metadata samples.
//
// Infos provides access to DDS sample metadata such as validity flags,
// timestamps, source identifiers, and other QoS-related information.
// This metadata is available for each sample received through an Input.
//
// Example usage:
//   input.Read()
//   samples := input.Samples()
//   infos := input.Infos()
//   
//   for i := 0; i < samples.Length(); i++ {
//       isValid, _ := infos.IsValid(i)
//       if isValid {
//           timestamp, _ := infos.GetSourceTimestamp(i)
//           fmt.Printf("Valid sample at time: %d\n", timestamp)
//       }
//   }
type Infos struct {
	input *Input
}

// Identity uniquely identifies a DDS sample and its source writer.
//
// Each DDS sample has an associated Identity that includes the GUID
// of the writer that published it and a sequence number for ordering.
type Identity struct {
	WriterGUID     [16]byte `json:"writer_guid"`     // Unique identifier of the publishing writer
	SequenceNumber int      `json:"sequence_number"` // Sequence number for sample ordering
}

/*******************
* Public Functions *
*******************/

// IsValid checks whether a sample contains valid data.
//
// This method returns true if the sample at the specified index contains
// valid data (not disposed or no-writers), false otherwise.
//
// Parameters:
//   - index: The index of the sample to check (0-based)
//
// Returns:
//   - bool: true if the sample contains valid data
//   - error: Non-nil if the index is out of bounds or operation fails
//
// Example:
//   isValid, err := infos.IsValid(0)
//   if err != nil {
//       log.Printf("Failed to check validity: %v", err)
//   } else if isValid {
//       // Process the valid sample
//   }
func (infos *Infos) IsValid(index int) (bool, error) {
	if infos == nil || infos.input == nil || infos.input.connector == nil {
		return false, errors.New("infos, input, or connector is null")
	}
	if index < 0 {
		return false, errors.New("index cannot be negative")
	}

	memberNameCStr := C.CString("valid_data")
	defer C.free(unsafe.Pointer(memberNameCStr))
	var retVal C.int

	retcode := int(C.RTI_Connector_get_boolean_from_infos(unsafe.Pointer(infos.input.connector.native), &retVal, infos.input.nameCStr, C.int(index+1), memberNameCStr))
	err := checkRetcode(retcode)

	return (retVal != 0), err
}

// GetSourceTimestamp retrieves the source timestamp of a sample.
//
// The source timestamp represents when the sample was written by the
// publishing application, as opposed to when it was received.
//
// Parameters:
//   - index: The index of the sample (0-based)
//
// Returns:
//   - int64: Source timestamp in nanoseconds since epoch
//   - error: Non-nil if the index is out of bounds or operation fails
//
// Example:
//   timestamp, err := infos.GetSourceTimestamp(0)
//   if err != nil {
//       log.Printf("Failed to get timestamp: %v", err)
//   } else {
//       fmt.Printf("Sample written at: %d ns\n", timestamp)
//   }
func (infos *Infos) GetSourceTimestamp(index int) (int64, error) {
	tsStr, err := infos.getJSONMember(index, "source_timestamp")
	if err != nil {
		return 0, err
	}

	ts, err := strconv.ParseInt(tsStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return ts, nil
}

// GetReceptionTimestamp retrieves the reception timestamp of a sample.
//
// The reception timestamp represents when the sample was received by the
// local participant, which may be different from when it was originally written.
//
// Parameters:
//   - index: The index of the sample (0-based)
//
// Returns:
//   - int64: Reception timestamp in nanoseconds since epoch
//   - error: Non-nil if the index is out of bounds or operation fails
//
// Example:
//   recvTime, err := infos.GetReceptionTimestamp(0)
//   if err != nil {
//       log.Printf("Failed to get reception time: %v", err)
//   } else {
//       fmt.Printf("Sample received at: %d ns\n", recvTime)
//   }
func (infos *Infos) GetReceptionTimestamp(index int) (int64, error) {
	tsStr, err := infos.getJSONMember(index, "reception_timestamp")
	if err != nil {
		return 0, err
	}

	ts, err := strconv.ParseInt(tsStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return ts, nil
}

// GetIdentity retrieves the identity of the writer that published a sample.
//
// The identity includes the writer's GUID and the sequence number of the sample,
// which together uniquely identify the sample and its source.
//
// Parameters:
//   - index: The index of the sample (0-based)
//
// Returns:
//   - Identity: Structure containing writer GUID and sequence number
//   - error: Non-nil if the index is out of bounds or operation fails
//
// Example:
//   identity, err := infos.GetIdentity(0)
//   if err != nil {
//       log.Printf("Failed to get identity: %v", err)
//   } else {
//       fmt.Printf("Writer GUID: %x, Seq: %d\n", identity.WriterGUID, identity.SequenceNumber)
//   }
func (infos *Infos) GetIdentity(index int) (Identity, error) {

	var writerID Identity

	identityStr, err := infos.getJSONMember(index, "sample_identity")
	if err != nil {
		return writerID, err
	}

	jsonByte := []byte(identityStr)
	err = json.Unmarshal(jsonByte, &writerID)
	if err != nil {
		return writerID, errors.New("JSON Unmarshal failed: " + err.Error())
	}

	return writerID, nil
}

// GetIdentityJSON retrieves the identity of a writer as a JSON string.
//
// This method returns the same information as GetIdentity() but formatted
// as a JSON string, which can be useful for logging or serialization.
//
// Parameters:
//   - index: The index of the sample (0-based)
//
// Returns:
//   - string: JSON representation of the identity
//   - error: Non-nil if the index is out of bounds or operation fails
//
// Example:
//   identityJSON, err := infos.GetIdentityJSON(0)
//   if err != nil {
//       log.Printf("Failed to get identity JSON: %v", err)
//   } else {
//       fmt.Printf("Identity: %s\n", identityJSON)
//   }
func (infos *Infos) GetIdentityJSON(index int) (string, error) {
	identityStr, err := infos.getJSONMember(index, "sample_identity")
	if err != nil {
		return "", err
	}

	return identityStr, nil
}

// GetRelatedIdentity is a function used for request-reply communications.
func (infos *Infos) GetRelatedIdentity(index int) (Identity, error) {

	var writerID Identity

	identityStr, err := infos.getJSONMember(index, "related_sample_identity")
	if err != nil {
		return writerID, err
	}

	jsonByte := []byte(identityStr)
	err = json.Unmarshal(jsonByte, &writerID)
	if err != nil {
		return writerID, errors.New("JSON Unmarshal failed: " + err.Error())
	}

	return writerID, nil
}

// GetRelatedIdentityJSON is a function used for get related identity in JSON.
func (infos *Infos) GetRelatedIdentityJSON(index int) (string, error) {
	identityStr, err := infos.getJSONMember(index, "related_sample_identity")
	if err != nil {
		return "", err
	}

	return identityStr, nil
}

// GetViewState is a function used to get a view state in string (either "NEW" or "NOT NEW").
func (infos *Infos) GetViewState(index int) (string, error) {
	viewStateStr, err := infos.getJSONMember(index, "view_state")
	if err != nil {
		return "", err
	}

	return viewStateStr, nil
}

// GetInstanceState is a function used to get a instance state in string (one of "ALIVE", "NOT_ALIVE_DISPOSED" or "NOT_ALIVE_NO_WRITERS").
func (infos *Infos) GetInstanceState(index int) (string, error) {
	instanceStateStr, err := infos.getJSONMember(index, "instance_state")
	if err != nil {
		return "", err
	}

	return instanceStateStr, nil
}

// GetSampleState is a function used to get a sample state in string (either "READ" or "NOT_READ").
func (infos *Infos) GetSampleState(index int) (string, error) {
	sampleStateStr, err := infos.getJSONMember(index, "sample_state")
	if err != nil {
		return "", err
	}

	return sampleStateStr, nil
}

// GetLength is a function to return the length of the
func (infos *Infos) GetLength() (int, error) {
	var retVal C.double
	retcode := int(C.RTI_Connector_get_sample_count(unsafe.Pointer(infos.input.connector.native), infos.input.nameCStr, &retVal))
	err := checkRetcode(retcode)
	return int(retVal), err
}

func (infos *Infos) getJSONMember(index int, memberName string) (string, error) {
	memberNameCStr := C.CString(memberName)
	defer C.free(unsafe.Pointer(memberNameCStr))

	var retValCStr *C.char

	retcode := int(C.RTI_Connector_get_json_from_infos(unsafe.Pointer(infos.input.connector.native), infos.input.nameCStr, C.int(index+1), memberNameCStr, &retValCStr))
	err := checkRetcode(retcode)
	if err != nil {
		return "", err
	}

	retValGoStr := C.GoString(retValCStr)
	C.RTI_Connector_free_string(retValCStr)

	return retValGoStr, nil
}
