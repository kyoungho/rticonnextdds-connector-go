# Test Coverage Improvements

This document describes the comprehensive test coverage improvements added to the RTI Connector Go codebase.

## Summary

Added extensive new tests covering previously untested or under-tested areas:
- **Request-Reply Pattern Testing** - Full request-reply communication with identity correlation
- **Array Data Type Testing** - Fixed-size array handling and verification
- **Sequence Data Type Testing** - Variable-length sequence handling
- **Inline XML Configuration** - Testing `str://` inline XML support
- **Concurrency Testing** - Documenting thread-safety behavior
- **Error Path Testing** - Comprehensive error handling validation

## New Test Files

### 1. `advanced_test.go` (900+ lines)
Comprehensive test suite covering advanced features and edge cases:

#### Request-Reply Pattern Tests
- **TestRequestReplyPattern**: Full request-reply workflow with identity correlation
  - Tests `WriteWithParams()` with custom identity
  - Tests `GetIdentityJSON()` and `GetRelatedIdentityJSON()`
  - Tests `GetRelatedIdentity()` for request-reply scenarios
  - Verifies identity matching between requests and replies

- **TestWriteWithParamsSourceTimestamp**: Custom source timestamp handling
  - Tests setting custom source timestamps via WriteWithParams
  - Verifies timestamp retrieval matches custom value

- **TestWriteWithParamsDispose**: Instance disposal testing
  - Tests dispose action via WriteWithParams
  - Verifies disposed instances return IsValid() = false
  - Tests instance state transitions

#### Array Data Type Tests
- **TestArrayDataTypes**: Fixed-size array element access
  - Tests setting array elements using index notation
  - Verifies array data round-trip (write → read)
  - Tests scalar and array field combinations

- **TestArrayJSON**: Array handling via JSON
  - Tests setting arrays via SetJSON()
  - Verifies JSON serialization of array data
  - Tests GetJSON() with array fields

#### Sequence Data Type Tests
- **TestSequenceDataTypes**: Variable-length sequence handling
  - Tests unbounded sequences (sequenceMaxLength="-1")
  - Tests bounded sequences with max length
  - Verifies sequence data via JSON round-trip

#### Inline XML Configuration Tests
- **TestInlineXMLConfiguration**: str:// prefix support
  - Tests creating connector with inline XML string
  - Verifies full data flow with inline configuration
  - Tests complete XML configuration without external files

- **TestInvalidInlineXML**: Error handling for malformed inline XML
  - Verifies proper error reporting for invalid XML
  - Tests parser error handling

#### Concurrency Tests
- **TestConcurrentReads**: Concurrent read operations
  - Documents non-thread-safe behavior of Input.Read()
  - Tests race conditions with multiple goroutines

- **TestConcurrentWrites**: Concurrent write operations
  - Documents non-thread-safe behavior of Output.Write()
  - Tests race conditions with SetInstance/Write

- **TestSynchronizedWrites**: Proper synchronization patterns
  - Demonstrates correct mutex usage for thread safety
  - Shows how to safely use Connector from multiple goroutines

#### Error Path and Edge Case Tests
- **TestNullPointerHandling**: Nil pointer safety
  - Tests all public APIs with nil receivers
  - Verifies error messages for nil Samples, Infos, Input, Output

- **TestNegativeIndices**: Negative index validation
  - Tests IsValid() with negative indices
  - Verifies error messages contain "negative"

- **TestTypeMismatch**: Type conversion error handling
  - Tests getting int32 from string field
  - Documents type mismatch behavior

- **TestOutOfBoundsAccess**: Index bounds checking
  - Tests accessing non-existent sample indices
  - Verifies no crashes on out-of-bounds access

- **TestNonExistentField**: Invalid field name handling
  - Tests getting/setting non-existent fields
  - Verifies appropriate error returns

- **TestSampleStateTransitions**: DDS state machine testing
  - Tests NOT_READ → READ transitions
  - Tests NEW → OLD view state transitions
  - Verifies ALIVE instance state

- **TestGetIdentityMethods**: Identity retrieval validation
  - Tests GetIdentity() returns valid WriterGUID
  - Tests GetIdentityJSON() format and content
  - Verifies sequence number > 0

- **TestMultipleConnectorInstances**: Multiple connector lifecycle
  - Tests creating 3+ connectors simultaneously
  - Verifies proper cleanup of all instances

- **TestLargeStringHandling**: Large data handling
  - Tests strings with 2000+ characters
  - Verifies no truncation or corruption

## New Test XML Configurations

### 1. `test/xml/RequestReplyTest.xml`
- Defines RequestType and ReplyType structures
- Configures Requester and Replier participants
- Sets up request/reply topic pairs
- Uses RELIABLE QoS with TRANSIENT_LOCAL durability

### 2. `test/xml/ArrayTest.xml`
- Defines ArrayTestType with fixed-size arrays
- Tests int_array[10] and string_array[5]
- Includes scalar fields for mixed testing

### 3. `test/xml/SequenceTest.xml`
- Defines SequenceTestType with variable-length sequences
- Tests unbounded int_sequence
- Tests bounded string_sequence with max length 100

## Test Coverage Impact

### Areas Previously at 0% Coverage (Now Tested)
- ✅ `GetRelatedIdentity()` - Request-reply identity correlation
- ✅ `GetRelatedIdentityJSON()` - JSON format for related identities
- ✅ `WriteWithParams()` - Full parameter testing (identity, timestamp, action)
- ✅ Inline XML configuration (`str://` prefix)
- ✅ Array data type handling
- ✅ Sequence data type handling

### Areas with Improved Coverage
- ✅ `GetIdentityJSON()` - Comprehensive format validation
- ✅ Error handling paths - Nil checks, negative indices, type mismatches
- ✅ Concurrency behavior - Documented thread-safety requirements
- ✅ Sample state transitions - Full state machine coverage
- ✅ Multiple connector instances - Lifecycle and cleanup

## Test Execution

### Running All Tests
```bash
# Run all tests including new advanced tests
go test -v -race -coverprofile=coverage.out

# View coverage report
go tool cover -html=coverage.out
```

### Running Specific Test Categories
```bash
# Request-reply tests only
go test -v -run TestRequestReply

# Array/sequence tests only
go test -v -run "TestArray|TestSequence"

# Concurrency tests only
go test -v -run TestConcurrent

# Error path tests only
go test -v -run "TestNull|TestNegative|TestType|TestOutOf|TestNonExistent"
```

### Running with Race Detector
```bash
# Recommended to catch concurrency issues
go test -v -race
```

## Notes on Thread Safety

The concurrency tests (`TestConcurrentReads`, `TestConcurrentWrites`) intentionally demonstrate **non-thread-safe behavior** as documented in the codebase:

> "The Connector is not thread-safe. You must provide your own synchronization when using it from multiple goroutines."

These tests serve to:
1. Document expected behavior under concurrent access
2. Demonstrate proper synchronization patterns (see `TestSynchronizedWrites`)
3. Catch any unexpected changes in threading behavior

## Future Test Enhancements

While this update significantly improves coverage, additional areas for future testing include:

1. **Security Testing** - If security features are enabled
2. **Performance Benchmarks** - Throughput and latency measurements
3. **Stress Testing** - High-frequency writes, queue overflow scenarios
4. **QoS Policy Testing** - Explicit testing of different QoS combinations
5. **Module-based Types** - Testing modular XML type organization

## Verification

To verify these improvements:

1. **Check compilation**: `go build ./...`
2. **Run tests**: `go test -v -race`
3. **Check coverage**: `go test -coverprofile=coverage.out && go tool cover -func=coverage.out`
4. **Run specific tests**: `go test -v -run TestRequestReply`

All new tests follow the existing test patterns and use the same helper functions (`newTestConnector()`, `newTestInput()`, `newTestOutput()`) for consistency.
