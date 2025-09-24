package rti // import "github.com/rticommunity/rticonnextdds-connector-go"

Package rti implements functions of RTI Connector for Connext DDS in Go

# Package rti implements functions of RTI Connector for Connext DDS in Go

# Package rti implements functions of RTI Connector for Connext DDS in Go

# Package rti implements functions of RTI Connector for Connext DDS in Go

# Package rti implements functions of RTI Connector for Connext DDS in Go

Package rti implements functions of RTI Connector for Connext DDS in Go

CONSTANTS

const (
	// DDSRetCodeNoData is a Return Code from CGO for no data return
	DDSRetCodeNoData = 11
	// DDSRetCodeTimeout is a Return Code from CGO for timeout code
	DDSRetCodeTimeout = 10
	// DDSRetCodeOK is a Return Code from CGO for good state
	DDSRetCodeOK = 0
)

VARIABLES

var ErrNoData = errors.New("DDS Exception: No Data")
    ErrNoData is returned when there is no data available in the DDS layer

var ErrTimeout = errors.New("DDS Exception: Timeout")
    ErrTimeout is returned when there is a timeout in the DDS layer


TYPES

type Connector struct {
	Inputs  []Input
	Outputs []Output
	// Has unexported fields.
}
    Connector is a container managing DDS inputs and outputs

func NewConnector(configName, url string) (*Connector, error)
    NewConnector is a constructor of Connector.

    url is the location of XML documents in URL format. For example:

        File specification: file:///usr/local/default_dds.xml
        String specification: str://"<dds><qos_library>…</qos_library></dds>"

    If you omit the URL schema name, Connector will assume a file name.
    For example:

        File Specification: /usr/local/default_dds.xml

func (connector *Connector) Delete() error
    Delete is a destructor of Connector

func (connector *Connector) GetInput(inputName string) (*Input, error)
    GetInput returns an input object

func (connector *Connector) GetOutput(outputName string) (*Output, error)
    GetOutput returns an output object

func (connector *Connector) Wait(timeoutMs int) error
    Wait is a function to block until data is available on an input

type Identity struct {
	WriterGUID     [16]byte `json:"writer_guid"`
	SequenceNumber int      `json:"sequence_number"`
}
    Identity is the structure for identifying

type Infos struct {
	// Has unexported fields.
}
    Infos is a sequence of info samples used by an input to read DDS meta data

func (infos *Infos) GetIdentity(index int) (Identity, error)
    GetIdentity is a function to get the identity of a writer that sent the
    sample

func (infos *Infos) GetIdentityJSON(index int) (string, error)
    GetIdentityJSON is a function to get the identity of a writer in JSON

func (infos *Infos) GetInstanceState(index int) (string, error)
    GetInstanceState is a function used to get a instance state in string (one
    of "ALIVE", "NOT_ALIVE_DISPOSED" or "NOT_ALIVE_NO_WRITERS").

func (infos *Infos) GetLength() (int, error)
    GetLength is a function to return the length of the

func (infos *Infos) GetReceptionTimestamp(index int) (int64, error)
    GetReceptionTimestamp is a function to get the reception timestamp of a
    sample

func (infos *Infos) GetRelatedIdentity(index int) (Identity, error)
    GetRelatedIdentity is a function used for request-reply communications.

func (infos *Infos) GetRelatedIdentityJSON(index int) (string, error)
    GetRelatedIdentityJSON is a function used for get related identity in JSON.

func (infos *Infos) GetSampleState(index int) (string, error)
    GetSampleState is a function used to get a sample state in string (either
    "READ" or "NOT_READ").

func (infos *Infos) GetSourceTimestamp(index int) (int64, error)
    GetSourceTimestamp is a function to get the source timestamp of a sample

func (infos *Infos) GetViewState(index int) (string, error)
    GetViewState is a function used to get a view state in string (either "NEW"
    or "NOT NEW").

func (infos *Infos) IsValid(index int) (bool, error)
    IsValid is a function to check validity of the element and return a boolean

type Input struct {
	Samples *Samples
	Infos   *Infos
	// Has unexported fields.
}
    Input subscribes to DDS data

func (input *Input) GetMatchedPublications() (string, error)
    Note that Connector Outputs are automatically assigned a name from the
    *data_writer name* in the XML configuration.

func (input *Input) Read() error
    Read is a function to read DDS samples from the DDS DataReader and allow
    access them via the Connector Samples. The Read function does not remove DDS
    samples from the DDS DataReader's receive queue.

func (input *Input) Take() error
    Take is a function to take DDS samples from the DDS DataReader and allow
    access them via the Connector Samples. The Take function removes DDS samples
    from the DDS DataReader's receive queue.

func (input *Input) WaitForPublications(timeoutMs int) (int, error)
    Waits until this input matches or unmatches a compatible DDS subscription.
    If the operation times out, it will raise :class:`TimeoutError`. Parameters:

        timeout: The maximum time to wait in milliseconds. Set -1 if you want infinite.

    Return: The change in the current number of matched outputs. If a
    positive number is returned, the input has matched with new publishers.
    If a negative number is returned the input has unmatched from an output.
    It is possible for multiple matches and/or unmatches to be returned (e.g.,
    0 could be returned, indicating that the input matched the same number of
    writers as it unmatched).

type Instance struct {
	// Has unexported fields.
}
    Instance is used by an output to write DDS data

func (instance *Instance) Set(v interface{}) error
    Set is a function that consumes an interface of multiple samples with
    different types and value TODO - think about a new name for this a function
    (e.g. SetType, SetFromType, FromType)

func (instance *Instance) SetBoolean(fieldName string, value bool) error
    SetBoolean is a function to set boolean to a fieldname of the samples

func (instance *Instance) SetByte(fieldName string, value byte) error
    SetByte is a function to set a byte to a fieldname of the samples

func (instance *Instance) SetFloat32(fieldName string, value float32) error
    SetFloat32 is a function to set a value of type float32 into samples

func (instance *Instance) SetFloat64(fieldName string, value float64) error
    SetFloat64 is a function to set a value of type float64 into samples

func (instance *Instance) SetInt(fieldName string, value int) error
    SetInt is a function to set a value of type int into samples

func (instance *Instance) SetInt16(fieldName string, value int16) error
    SetInt16 is a function to set a value of type int16 into samples

func (instance *Instance) SetInt32(fieldName string, value int32) error
    SetInt32 is a function to set a value of type int32 into samples

func (instance *Instance) SetInt64(fieldName string, value int64) error
    SetInt64 is a function to set a value of type int64 into samples

func (instance *Instance) SetInt8(fieldName string, value int8) error
    SetInt8 is a function to set a value of type int8 into samples

func (instance *Instance) SetJSON(blob []byte) error
    SetJSON is a function to set JSON string in the form of slice of bytes into
    Instance

func (instance *Instance) SetRune(fieldName string, value rune) error
    SetRune is a function to set rune to a fieldname of the samples

func (instance *Instance) SetString(fieldName, value string) error
    SetString is a function that set a string to a fieldname of the samples

func (instance *Instance) SetUint(fieldName string, value uint) error
    SetUint is a function to set a value of type uint into samples

func (instance *Instance) SetUint16(fieldName string, value uint16) error
    SetUint16 is a function to set a value of type uint16 into samples

func (instance *Instance) SetUint32(fieldName string, value uint32) error
    SetUint32 is a function to set a value of type uint32 into samples

func (instance *Instance) SetUint64(fieldName string, value uint64) error
    SetUint64 is a function to set a value of type uint64 into samples

func (instance *Instance) SetUint8(fieldName string, value uint8) error
    SetUint8 is a function to set a value of type uint8 into samples

type Output struct {
	Instance *Instance
	// Has unexported fields.
}
    Output publishes DDS data

func (output *Output) ClearMembers() error
    ClearMembers is a function to initialize a DDS data instance in an output

func (output *Output) GetMatchedSubscriptions() (string, error)
    Note that Connector Inputs are automatically assigned a name from the
    *data_reader name* in the XML configuration.

func (output *Output) WaitForSubscriptions(timeoutMs int) (int, error)
    Return: The change in the current number of matched outputs. If a
    positive number is returned, the input has matched with new publishers.
    If a negative number is returned the input has unmatched from an output.
    It is possible for multiple matches and/or unmatches to be returned (e.g.,
    0 could be returned, indicating that the input matched the same number of
    writers as it unmatched).

func (output *Output) Write() error
    Write is a function to write a DDS data instance in an output

func (output *Output) WriteWithParams(jsonStr string) error
    WriteWithParams is a function to write a DDS data instance with parameters
    The supported parameters are: action: One of "write" (default),
    "dispose" or "unregister" source_timestamp: The source timestamp, an
    integer representing the total number of nanoseconds identity: A dictionary
    containing the keys "writer_guid" (a list of 16 bytes) and "sequence_number"
    (an integer) that uniquely identifies this sample. related_sample_identity:
    Used for request-reply communications. It has the same format as "identity"
    For example:: output.Write(

          identity={"writer_guid":[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15], "sequence_number":1},
        	 timestamp=1000000000)

type SampleHandler func(samples *Samples, infos *Infos)
    SampleHandler is an User defined function type that takes in pointers of
    Samples and Infos and will handle received samples.

type Samples struct {
	// Has unexported fields.
}
    Samples is a sequence of data samples used by an input to read DDS data

func (samples *Samples) Get(index int, v interface{}) error
    Get is a function to retrieve all the information of the samples and put it
    into an interface

func (samples *Samples) GetBoolean(index int, fieldName string) (bool, error)
    GetBoolean is a function to retrieve a value of type boolean from the
    samples

func (samples *Samples) GetByte(index int, fieldName string) (byte, error)
    GetByte is a function to retrieve a value of type byte from the samples

func (samples *Samples) GetFloat32(index int, fieldName string) (float32, error)
    GetFloat32 is a function to retrieve a value of type float32 from the
    samples

func (samples *Samples) GetFloat64(index int, fieldName string) (float64, error)
    GetFloat64 is a function to retrieve a value of type float64 from the
    samples

func (samples *Samples) GetInt(index int, fieldName string) (int, error)
    GetInt is a function to retrieve a value of type int from the samples

func (samples *Samples) GetInt16(index int, fieldName string) (int16, error)
    GetInt16 is a function to retrieve a value of type int16 from the samples

func (samples *Samples) GetInt32(index int, fieldName string) (int32, error)
    GetInt32 is a function to retrieve a value of type int32 from the samples

func (samples *Samples) GetInt64(index int, fieldName string) (int64, error)
    GetInt64 is a function to retrieve a value of type int64 from the samples

func (samples *Samples) GetInt8(index int, fieldName string) (int8, error)
    GetInt8 is a function to retrieve a value of type int8 from the samples

func (samples *Samples) GetJSON(index int) ([]byte, error)
    GetJSON is a function to retrieve a slice of bytes of a JSON string from the
    samples

func (samples *Samples) GetLength() (int, error)
    GetLength is a function to get the number of samples

func (samples *Samples) GetRune(index int, fieldName string) (rune, error)
    GetRune is a function to retrieve a value of type rune from the samples

func (samples *Samples) GetString(index int, fieldName string) (string, error)
    GetString is a function to retrieve a value of type string from the samples

func (samples *Samples) GetUint(index int, fieldName string) (uint, error)
    GetUint is a function to retrieve a value of type uint from the samples

func (samples *Samples) GetUint16(index int, fieldName string) (uint16, error)
    GetUint16 is a function to retrieve a value of type uint16 from the samples

func (samples *Samples) GetUint32(index int, fieldName string) (uint32, error)
    GetUint32 is a function to retrieve a value of type uint32 from the samples

func (samples *Samples) GetUint64(index int, fieldName string) (uint64, error)
    GetUint64 is a function to retrieve a value of type uint64 from the samples

func (samples *Samples) GetUint8(index int, fieldName string) (uint8, error)
    GetUint8 is a function to retrieve a value of type uint8 from the samples

