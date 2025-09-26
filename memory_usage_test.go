package rti

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

// getProcessMemory returns RSS (Resident Set Size) in bytes
func getProcessMemory() (int64, error) {
	pid := os.Getpid()
	cmd := exec.Command("ps", "-o", "rss=", "-p", strconv.Itoa(pid))
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	rssKB, err := strconv.ParseInt(strings.TrimSpace(string(output)), 10, 64)
	if err != nil {
		return 0, err
	}

	return rssKB * 1024, nil // Convert KB to bytes
}

// TestMemoryUsage tests RTI Connector with comprehensive memory profiling
func TestMemoryUsage(t *testing.T) {
	iterations := 1
	if iterStr := os.Getenv("MEMTEST_ITERATIONS"); iterStr != "" {
		if parsed, err := strconv.Atoi(iterStr); err == nil && parsed > 0 {
			iterations = parsed
		}
	}

	// Get initial system memory
	initialRSS, err := getProcessMemory()
	if err != nil {
		t.Logf("Warning: Could not get initial RSS: %v", err)
		initialRSS = 0
	}

	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	t.Logf("Starting memory measurement:")
	t.Logf("  Initial RSS: %d bytes (%.2f MB)", initialRSS, float64(initialRSS)/1024/1024)
	t.Logf("  Initial Go heap: %d bytes (%.2f MB)", m1.HeapAlloc, float64(m1.HeapAlloc)/1024/1024)

	for i := 0; i < iterations; i++ {
		runGoGetExample(t, i)

		// Sample memory after each iteration
		if iterations > 1 && (i+1)%max(1, iterations/5) == 0 {
			currentRSS, _ := getProcessMemory()
			var currentMem runtime.MemStats
			runtime.ReadMemStats(&currentMem)
			t.Logf("  After iteration %d: RSS=%d bytes, Go heap=%d bytes",
				i+1, currentRSS, currentMem.HeapAlloc)
		}
	}

	runtime.GC()
	time.Sleep(100 * time.Millisecond) // Give GC time to complete
	runtime.ReadMemStats(&m2)

	// Get final system memory
	finalRSS, err := getProcessMemory()
	if err != nil {
		t.Logf("Warning: Could not get final RSS: %v", err)
		finalRSS = initialRSS
	}

	allocDelta := m2.TotalAlloc - m1.TotalAlloc
	heapDelta := int64(m2.HeapAlloc) - int64(m1.HeapAlloc)
	rssDelta := finalRSS - initialRSS

	t.Logf("\n=== Memory Analysis Results ===")
	t.Logf("Iterations: %d", iterations)
	t.Logf("\nGo Heap Memory:")
	t.Logf("  Total allocations: %d bytes (%.2f KB)", allocDelta, float64(allocDelta)/1024)
	t.Logf("  Heap delta: %d bytes (%.2f KB)", heapDelta, float64(heapDelta)/1024)
	t.Logf("  Final heap size: %d bytes (%.2f MB)", m2.HeapAlloc, float64(m2.HeapAlloc)/1024/1024)
	t.Logf("  GC runs: %d", m2.NumGC-m1.NumGC)
	t.Logf("\nSystem Memory (RSS - includes C libraries):")
	t.Logf("  Initial RSS: %d bytes (%.2f MB)", initialRSS, float64(initialRSS)/1024/1024)
	t.Logf("  Final RSS: %d bytes (%.2f MB)", finalRSS, float64(finalRSS)/1024/1024)
	t.Logf("  RSS delta: %d bytes (%.2f MB)", rssDelta, float64(rssDelta)/1024/1024)
	t.Logf("\nMemory per Operation:")
	if iterations > 0 {
		t.Logf("  Go heap per op: %.2f bytes", float64(allocDelta)/float64(iterations))
		t.Logf("  RSS per op: %.2f bytes", float64(rssDelta)/float64(iterations))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func runGoGetExample(t *testing.T, iteration int) {
	// Same XML config as the original example
	xmlConfig := `str://"<dds>
  <qos_library name="QosLibrary">
    <qos_profile name="DefaultProfile" base_name="BuiltinQosLibExp::Generic.StrictReliable" is_default_qos="true"/>
  </qos_library>
  
  <types>
    <struct name="TestType">
      <member name="message" type="string"/>
      <member name="count" type="long"/>
    </struct>
  </types>
  
  <domain_library name="MyDomainLibrary">
    <domain name="MyDomain" domain_id="0">
      <register_type name="TestType" type_ref="TestType"/>
      <topic name="TestTopic" register_type_ref="TestType"/>
    </domain>
  </domain_library>
  
  <domain_participant_library name="MyParticipantLibrary">
    <domain_participant name="Zero" domain_ref="MyDomainLibrary::MyDomain">
      <publisher name="MyPublisher">
        <data_writer name="MyWriter" topic_ref="TestTopic"/>
      </publisher>
    </domain_participant>
  </domain_participant_library>
</dds>"`

	// Create connector from XML string
	connector, err := NewConnector("MyParticipantLibrary::Zero", xmlConfig)
	if err != nil {
		t.Fatalf("Iteration %d: Failed to create connector: %v", iteration, err)
	}
	defer connector.Delete()

	// Get output (writer) and publish a simple message
	output, err := connector.GetOutput("MyPublisher::MyWriter")
	if err != nil {
		t.Fatalf("Iteration %d: Failed to get output: %v", iteration, err)
	}

	// Publish test message
	output.Instance.SetString("message", fmt.Sprintf("Hello from iteration %d!", iteration))
	output.Instance.SetInt("count", iteration)
	err = output.Write()
	if err != nil {
		t.Fatalf("Iteration %d: Failed to write: %v", iteration, err)
	}

	if iteration%10 == 0 {
		t.Logf("Completed iteration %d", iteration)
	}
}

// BenchmarkConnectorMemory provides benchmark-based memory analysis
func BenchmarkConnectorMemory(b *testing.B) {
	xmlConfig := `str://"<dds>
  <qos_library name="QosLibrary">
    <qos_profile name="DefaultProfile" base_name="BuiltinQosLibExp::Generic.StrictReliable" is_default_qos="true"/>
  </qos_library>
  
  <types>
    <struct name="TestType">
      <member name="message" type="string"/>
      <member name="count" type="long"/>
    </struct>
  </types>
  
  <domain_library name="MyDomainLibrary">
    <domain name="MyDomain" domain_id="0">
      <register_type name="TestType" type_ref="TestType"/>
      <topic name="TestTopic" register_type_ref="TestType"/>
    </domain>
  </domain_library>
  
  <domain_participant_library name="MyParticipantLibrary">
    <domain_participant name="Zero" domain_ref="MyDomainLibrary::MyDomain">
      <publisher name="MyPublisher">
        <data_writer name="MyWriter" topic_ref="TestTopic"/>
      </publisher>
    </domain_participant>
  </domain_participant_library>
</dds>"`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		connector, err := NewConnector("MyParticipantLibrary::Zero", xmlConfig)
		if err != nil {
			b.Fatalf("Failed to create connector: %v", err)
		}

		output, err := connector.GetOutput("MyPublisher::MyWriter")
		if err != nil {
			connector.Delete()
			b.Fatalf("Failed to get output: %v", err)
		}

		output.Instance.SetString("message", "Benchmark message")
		output.Instance.SetInt("count", i)
		output.Write()

		connector.Delete()
	}
}
