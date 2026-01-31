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
	"errors"
	"unsafe"
)

/********
* Types *
*********/

// Input represents a DDS DataReader for subscribing to data.
//
// An Input allows you to receive data samples from DDS Topics. It wraps a
// native RTI Connext DataReader and provides methods to read or take samples.
//
// Key differences:
//   - Read(): Copies samples but leaves them in the DataReader's queue
//   - Take(): Removes samples from the DataReader's queue
//
// Access received data through the Samples field (actual data) and
// Infos field (metadata like timestamps and sample states).
type Input struct {
	native    unsafe.Pointer // a pointer to a native DataReader
	connector *Connector
	name      string // name of the native DataReader
	nameCStr  *C.char
	Samples   *Samples // Collection of received data samples
	Infos     *Infos   // Collection of sample metadata
}

/*******************
* Public Functions *
*******************/

// isValid checks if the input and its connector are valid
func (input *Input) isValid() error {
	if input == nil {
		return errors.New("input is null")
	}
	if input.connector == nil {
		return errors.New("input connector is null")
	}
	if input.connector.native == nil {
		return errors.New("connector has been deleted")
	}
	return nil
}

// Read copies DDS samples from the DataReader without removing them from the receive queue.
//
// After a successful read, samples can be accessed via input.Samples and metadata
// via input.Infos. The samples remain in the DataReader's queue and can be read
// again. Use Take() if you want to remove samples from the queue.
//
// Returns:
//   - error: ErrNoData if no samples available, ErrTimeout on timeout, or other error
//
// Example:
//
//	err := input.Read()
//	if err == rti.ErrNoData {
//	    // No data available
//	    return
//	}
//	if err != nil {
//	    log.Printf("Read error: %v", err)
//	    return
//	}
//
//	length, _ := input.Samples.GetLength()
//	for i := 0; i < length; i++ {
//	    color, _ := input.Samples.GetString(i, "color")
//	    fmt.Printf("Read: %s\n", color)
//	}
func (input *Input) Read() error {
	if err := input.isValid(); err != nil {
		return err
	}

	retcode := int(C.RTI_Connector_read(unsafe.Pointer(input.connector.native), input.nameCStr))
	return checkRetcode(retcode)
}

// Take is a function to take DDS samples from the DDS DataReader
// and allow access them via the Connector Samples. The Take
// function removes DDS samples from the DDS DataReader's receive queue.
func (input *Input) Take() error {
	if err := input.isValid(); err != nil {
		return err
	}

	retcode := int(C.RTI_Connector_take(unsafe.Pointer(input.connector.native), input.nameCStr))
	return checkRetcode(retcode)
}

// Waits until this input matches or unmatches a compatible DDS subscription.
// If the operation times out, it will raise :class:`TimeoutError`.
// Parameters:
//
//	timeout: The maximum time to wait in milliseconds. Set -1 if you want infinite.
//
// Return: The change in the current number of matched outputs. If a positive number is returned, the input has matched with new publishers. If a negative number is returned the input has unmatched from an output. It is possible for multiple matches and/or unmatches to be returned (e.g., 0 could be returned, indicating that the input matched the same number of writers as it unmatched).
func (input *Input) WaitForPublications(timeoutMs int) (int, error) {
	if err := input.isValid(); err != nil {
		return -1, err
	}

	var currentCountChange C.int

	retcode := int(C.RTI_Connector_wait_for_matched_publication(unsafe.Pointer(input.native), C.int(timeoutMs), &currentCountChange))
	return int(currentCountChange), checkRetcode(retcode)
}

// Returns information about the matched publications
// This function returns a JSON string where each element is a dictionary with
// information about a publication matched with this Input.

// Currently, the only key in the dictionaries is ``"name"``,
// containing the publication name. If a publication doesn't have name,
// the value for the key ``name`` is ``None``.

// Note that Connector Outputs are automatically assigned a name from the
// *data_writer name* in the XML configuration.
func (input *Input) GetMatchedPublications() (string, error) {
	if err := input.isValid(); err != nil {
		return "", err
	}

	var jsonCStr *C.char

	retcode := int(C.RTI_Connector_get_matched_publications(unsafe.Pointer(input.native), &jsonCStr))
	err := checkRetcode(retcode)
	if err != nil {
		return "", err
	}

	jsonGoStr := C.GoString(jsonCStr)
	C.RTI_Connector_free_string(jsonCStr)

	return jsonGoStr, nil
}

// ReturnLoan returns any loaned samples back to the DDS middleware.
//
// After calling Read() or Take(), samples are "loaned" from the middleware to the
// application. This method explicitly returns those loans, freeing the associated
// resources in the DDS DataReader's receive queue.
//
// Calling this method is particularly important in applications that:
//   - Receive large volumes of data
//   - Keep samples for extended periods between Read/Take operations
//   - Need to manage memory usage explicitly
//
// Note that Take() removes samples from the queue (so they're implicitly returned),
// but Read() keeps samples accessible until they're explicitly returned.
//
// Returns:
//   - error: Non-nil if the operation fails
//
// Example:
//
//	err := input.Read()
//	if err != nil {
//	    log.Printf("Read failed: %v", err)
//	    return
//	}
//
//	// Process samples
//	length, _ := input.Samples.GetLength()
//	for i := 0; i < length; i++ {
//	    data, _ := input.Samples.GetJSON(i)
//	    processSample(data)
//	}
//
//	// Return loans to free resources
//	err = input.ReturnLoan()
//	if err != nil {
//	    log.Printf("ReturnLoan failed: %v", err)
//	}
func (input *Input) ReturnLoan() error {
	if input == nil {
		return errors.New("input is null")
	}

	retcode := int(C.RTI_Connector_return_loan(unsafe.Pointer(input.connector.native), input.nameCStr))
	return checkRetcode(retcode)
}
