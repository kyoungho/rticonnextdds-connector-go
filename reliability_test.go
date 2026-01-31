package rti

import (
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDoubleDelete verifies that Delete() is idempotent
func TestDoubleDelete(t *testing.T) {
	_, curPath, _, _ := runtime.Caller(0)
	xmlPath := path.Join(path.Dir(curPath), "./test/xml/Test.xml")

	connector, err := NewConnector("MyParticipantLibrary::Zero", xmlPath)
	assert.Nil(t, err)
	assert.NotNil(t, connector)

	// First delete - should succeed
	err = connector.Delete()
	assert.Nil(t, err)

	// Second delete - should not crash (idempotent)
	err = connector.Delete()
	assert.Nil(t, err)
}

// TestUseAfterDelete verifies that operations fail gracefully after Delete()
func TestUseAfterDelete(t *testing.T) {
	_, curPath, _, _ := runtime.Caller(0)
	xmlPath := path.Join(path.Dir(curPath), "./test/xml/Test.xml")

	connector, err := NewConnector("MyParticipantLibrary::Zero", xmlPath)
	assert.Nil(t, err)
	assert.NotNil(t, connector)

	output, err := connector.GetOutput("MyPublisher::MyWriter")
	assert.Nil(t, err)
	assert.NotNil(t, output)

	input, err := connector.GetInput("MySubscriber::MyReader")
	assert.Nil(t, err)
	assert.NotNil(t, input)

	// Delete the connector
	err = connector.Delete()
	assert.Nil(t, err)

	// Operations on connector should fail
	_, err = connector.GetOutput("MyPublisher::MyWriter")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "deleted")

	_, err = connector.GetInput("MySubscriber::MyReader")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "deleted")

	err = connector.Wait(100)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "deleted")

	// Operations on output should fail
	err = output.Write()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "deleted")

	err = output.WriteWithParams("{}")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "deleted")

	err = output.ClearMembers()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "deleted")

	// Operations on input should fail
	err = input.Read()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "deleted")

	err = input.Take()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "deleted")
}

// TestInvalidOutputName verifies memory leak fix when GetOutput fails
func TestInvalidOutputName(t *testing.T) {
	_, curPath, _, _ := runtime.Caller(0)
	xmlPath := path.Join(path.Dir(curPath), "./test/xml/Test.xml")

	connector, err := NewConnector("MyParticipantLibrary::Zero", xmlPath)
	assert.Nil(t, err)
	assert.NotNil(t, connector)
	defer connector.Delete()

	// Try to get an output that doesn't exist
	// This used to leak memory - the C string was allocated but never freed
	output, err := connector.GetOutput("NonExistent::Writer")
	assert.NotNil(t, err)
	assert.Nil(t, output)

	// Do it multiple times to verify no memory leak
	for i := 0; i < 100; i++ {
		_, err := connector.GetOutput("NonExistent::Writer" + string(rune(i)))
		assert.NotNil(t, err)
	}
}

// TestInvalidInputName verifies memory leak fix when GetInput fails
func TestInvalidInputName(t *testing.T) {
	_, curPath, _, _ := runtime.Caller(0)
	xmlPath := path.Join(path.Dir(curPath), "./test/xml/Test.xml")

	connector, err := NewConnector("MyParticipantLibrary::Zero", xmlPath)
	assert.Nil(t, err)
	assert.NotNil(t, connector)
	defer connector.Delete()

	// Try to get an input that doesn't exist
	// This used to leak memory - the C string was allocated but never freed
	input, err := connector.GetInput("NonExistent::Reader")
	assert.NotNil(t, err)
	assert.Nil(t, input)

	// Do it multiple times to verify no memory leak
	for i := 0; i < 100; i++ {
		_, err := connector.GetInput("NonExistent::Reader" + string(rune(i)))
		assert.NotNil(t, err)
	}
}
