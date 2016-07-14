// This file was generated by counterfeiter
package connectionfakes

import (
	"io"
	"sync"
	"time"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/garden/client/connection"
)

type FakeConnection struct {
	PingStub        func() error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns     struct {
		result1 error
	}
	CapacityStub        func() (garden.Capacity, error)
	capacityMutex       sync.RWMutex
	capacityArgsForCall []struct{}
	capacityReturns     struct {
		result1 garden.Capacity
		result2 error
	}
	CreateStub        func(spec garden.ContainerSpec) (string, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		spec garden.ContainerSpec
	}
	createReturns struct {
		result1 string
		result2 error
	}
	ListStub        func(properties garden.Properties) ([]string, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		properties garden.Properties
	}
	listReturns struct {
		result1 []string
		result2 error
	}
	DestroyStub        func(handle string) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		handle string
	}
	destroyReturns struct {
		result1 error
	}
	StopStub        func(handle string, kill bool) error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		handle string
		kill   bool
	}
	stopReturns struct {
		result1 error
	}
	InfoStub        func(handle string) (garden.ContainerInfo, error)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
		handle string
	}
	infoReturns struct {
		result1 garden.ContainerInfo
		result2 error
	}
	BulkInfoStub        func(handles []string) (map[string]garden.ContainerInfoEntry, error)
	bulkInfoMutex       sync.RWMutex
	bulkInfoArgsForCall []struct {
		handles []string
	}
	bulkInfoReturns struct {
		result1 map[string]garden.ContainerInfoEntry
		result2 error
	}
	BulkMetricsStub        func(handles []string) (map[string]garden.ContainerMetricsEntry, error)
	bulkMetricsMutex       sync.RWMutex
	bulkMetricsArgsForCall []struct {
		handles []string
	}
	bulkMetricsReturns struct {
		result1 map[string]garden.ContainerMetricsEntry
		result2 error
	}
	StreamInStub        func(handle string, spec garden.StreamInSpec) error
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		handle string
		spec   garden.StreamInSpec
	}
	streamInReturns struct {
		result1 error
	}
	StreamOutStub        func(handle string, spec garden.StreamOutSpec) (io.ReadCloser, error)
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		handle string
		spec   garden.StreamOutSpec
	}
	streamOutReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	CurrentBandwidthLimitsStub        func(handle string) (garden.BandwidthLimits, error)
	currentBandwidthLimitsMutex       sync.RWMutex
	currentBandwidthLimitsArgsForCall []struct {
		handle string
	}
	currentBandwidthLimitsReturns struct {
		result1 garden.BandwidthLimits
		result2 error
	}
	CurrentCPULimitsStub        func(handle string) (garden.CPULimits, error)
	currentCPULimitsMutex       sync.RWMutex
	currentCPULimitsArgsForCall []struct {
		handle string
	}
	currentCPULimitsReturns struct {
		result1 garden.CPULimits
		result2 error
	}
	CurrentDiskLimitsStub        func(handle string) (garden.DiskLimits, error)
	currentDiskLimitsMutex       sync.RWMutex
	currentDiskLimitsArgsForCall []struct {
		handle string
	}
	currentDiskLimitsReturns struct {
		result1 garden.DiskLimits
		result2 error
	}
	CurrentMemoryLimitsStub        func(handle string) (garden.MemoryLimits, error)
	currentMemoryLimitsMutex       sync.RWMutex
	currentMemoryLimitsArgsForCall []struct {
		handle string
	}
	currentMemoryLimitsReturns struct {
		result1 garden.MemoryLimits
		result2 error
	}
	RunStub        func(handle string, spec garden.ProcessSpec, io garden.ProcessIO) (garden.Process, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		handle string
		spec   garden.ProcessSpec
		io     garden.ProcessIO
	}
	runReturns struct {
		result1 garden.Process
		result2 error
	}
	AttachStub        func(handle string, processID string, io garden.ProcessIO) (garden.Process, error)
	attachMutex       sync.RWMutex
	attachArgsForCall []struct {
		handle    string
		processID string
		io        garden.ProcessIO
	}
	attachReturns struct {
		result1 garden.Process
		result2 error
	}
	NetInStub        func(handle string, hostPort, containerPort uint32) (uint32, uint32, error)
	netInMutex       sync.RWMutex
	netInArgsForCall []struct {
		handle        string
		hostPort      uint32
		containerPort uint32
	}
	netInReturns struct {
		result1 uint32
		result2 uint32
		result3 error
	}
	NetOutStub        func(handle string, rule garden.NetOutRule) error
	netOutMutex       sync.RWMutex
	netOutArgsForCall []struct {
		handle string
		rule   garden.NetOutRule
	}
	netOutReturns struct {
		result1 error
	}
	SetGraceTimeStub        func(handle string, graceTime time.Duration) error
	setGraceTimeMutex       sync.RWMutex
	setGraceTimeArgsForCall []struct {
		handle    string
		graceTime time.Duration
	}
	setGraceTimeReturns struct {
		result1 error
	}
	PropertiesStub        func(handle string) (garden.Properties, error)
	propertiesMutex       sync.RWMutex
	propertiesArgsForCall []struct {
		handle string
	}
	propertiesReturns struct {
		result1 garden.Properties
		result2 error
	}
	PropertyStub        func(handle string, name string) (string, error)
	propertyMutex       sync.RWMutex
	propertyArgsForCall []struct {
		handle string
		name   string
	}
	propertyReturns struct {
		result1 string
		result2 error
	}
	SetPropertyStub        func(handle string, name string, value string) error
	setPropertyMutex       sync.RWMutex
	setPropertyArgsForCall []struct {
		handle string
		name   string
		value  string
	}
	setPropertyReturns struct {
		result1 error
	}
	MetricsStub        func(handle string) (garden.Metrics, error)
	metricsMutex       sync.RWMutex
	metricsArgsForCall []struct {
		handle string
	}
	metricsReturns struct {
		result1 garden.Metrics
		result2 error
	}
	RemovePropertyStub        func(handle string, name string) error
	removePropertyMutex       sync.RWMutex
	removePropertyArgsForCall []struct {
		handle string
		name   string
	}
	removePropertyReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConnection) Ping() error {
	fake.pingMutex.Lock()
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	} else {
		return fake.pingReturns.result1
	}
}

func (fake *FakeConnection) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeConnection) PingReturns(result1 error) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnection) Capacity() (garden.Capacity, error) {
	fake.capacityMutex.Lock()
	fake.capacityArgsForCall = append(fake.capacityArgsForCall, struct{}{})
	fake.recordInvocation("Capacity", []interface{}{})
	fake.capacityMutex.Unlock()
	if fake.CapacityStub != nil {
		return fake.CapacityStub()
	} else {
		return fake.capacityReturns.result1, fake.capacityReturns.result2
	}
}

func (fake *FakeConnection) CapacityCallCount() int {
	fake.capacityMutex.RLock()
	defer fake.capacityMutex.RUnlock()
	return len(fake.capacityArgsForCall)
}

func (fake *FakeConnection) CapacityReturns(result1 garden.Capacity, result2 error) {
	fake.CapacityStub = nil
	fake.capacityReturns = struct {
		result1 garden.Capacity
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) Create(spec garden.ContainerSpec) (string, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		spec garden.ContainerSpec
	}{spec})
	fake.recordInvocation("Create", []interface{}{spec})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(spec)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeConnection) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeConnection) CreateArgsForCall(i int) garden.ContainerSpec {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].spec
}

func (fake *FakeConnection) CreateReturns(result1 string, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) List(properties garden.Properties) ([]string, error) {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		properties garden.Properties
	}{properties})
	fake.recordInvocation("List", []interface{}{properties})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(properties)
	} else {
		return fake.listReturns.result1, fake.listReturns.result2
	}
}

func (fake *FakeConnection) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeConnection) ListArgsForCall(i int) garden.Properties {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].properties
}

func (fake *FakeConnection) ListReturns(result1 []string, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) Destroy(handle string) error {
	fake.destroyMutex.Lock()
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("Destroy", []interface{}{handle})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(handle)
	} else {
		return fake.destroyReturns.result1
	}
}

func (fake *FakeConnection) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeConnection) DestroyArgsForCall(i int) string {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return fake.destroyArgsForCall[i].handle
}

func (fake *FakeConnection) DestroyReturns(result1 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnection) Stop(handle string, kill bool) error {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		handle string
		kill   bool
	}{handle, kill})
	fake.recordInvocation("Stop", []interface{}{handle, kill})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub(handle, kill)
	} else {
		return fake.stopReturns.result1
	}
}

func (fake *FakeConnection) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeConnection) StopArgsForCall(i int) (string, bool) {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return fake.stopArgsForCall[i].handle, fake.stopArgsForCall[i].kill
}

func (fake *FakeConnection) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnection) Info(handle string) (garden.ContainerInfo, error) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("Info", []interface{}{handle})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub(handle)
	} else {
		return fake.infoReturns.result1, fake.infoReturns.result2
	}
}

func (fake *FakeConnection) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeConnection) InfoArgsForCall(i int) string {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return fake.infoArgsForCall[i].handle
}

func (fake *FakeConnection) InfoReturns(result1 garden.ContainerInfo, result2 error) {
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 garden.ContainerInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) BulkInfo(handles []string) (map[string]garden.ContainerInfoEntry, error) {
	var handlesCopy []string
	if handles != nil {
		handlesCopy = make([]string, len(handles))
		copy(handlesCopy, handles)
	}
	fake.bulkInfoMutex.Lock()
	fake.bulkInfoArgsForCall = append(fake.bulkInfoArgsForCall, struct {
		handles []string
	}{handlesCopy})
	fake.recordInvocation("BulkInfo", []interface{}{handlesCopy})
	fake.bulkInfoMutex.Unlock()
	if fake.BulkInfoStub != nil {
		return fake.BulkInfoStub(handles)
	} else {
		return fake.bulkInfoReturns.result1, fake.bulkInfoReturns.result2
	}
}

func (fake *FakeConnection) BulkInfoCallCount() int {
	fake.bulkInfoMutex.RLock()
	defer fake.bulkInfoMutex.RUnlock()
	return len(fake.bulkInfoArgsForCall)
}

func (fake *FakeConnection) BulkInfoArgsForCall(i int) []string {
	fake.bulkInfoMutex.RLock()
	defer fake.bulkInfoMutex.RUnlock()
	return fake.bulkInfoArgsForCall[i].handles
}

func (fake *FakeConnection) BulkInfoReturns(result1 map[string]garden.ContainerInfoEntry, result2 error) {
	fake.BulkInfoStub = nil
	fake.bulkInfoReturns = struct {
		result1 map[string]garden.ContainerInfoEntry
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) BulkMetrics(handles []string) (map[string]garden.ContainerMetricsEntry, error) {
	var handlesCopy []string
	if handles != nil {
		handlesCopy = make([]string, len(handles))
		copy(handlesCopy, handles)
	}
	fake.bulkMetricsMutex.Lock()
	fake.bulkMetricsArgsForCall = append(fake.bulkMetricsArgsForCall, struct {
		handles []string
	}{handlesCopy})
	fake.recordInvocation("BulkMetrics", []interface{}{handlesCopy})
	fake.bulkMetricsMutex.Unlock()
	if fake.BulkMetricsStub != nil {
		return fake.BulkMetricsStub(handles)
	} else {
		return fake.bulkMetricsReturns.result1, fake.bulkMetricsReturns.result2
	}
}

func (fake *FakeConnection) BulkMetricsCallCount() int {
	fake.bulkMetricsMutex.RLock()
	defer fake.bulkMetricsMutex.RUnlock()
	return len(fake.bulkMetricsArgsForCall)
}

func (fake *FakeConnection) BulkMetricsArgsForCall(i int) []string {
	fake.bulkMetricsMutex.RLock()
	defer fake.bulkMetricsMutex.RUnlock()
	return fake.bulkMetricsArgsForCall[i].handles
}

func (fake *FakeConnection) BulkMetricsReturns(result1 map[string]garden.ContainerMetricsEntry, result2 error) {
	fake.BulkMetricsStub = nil
	fake.bulkMetricsReturns = struct {
		result1 map[string]garden.ContainerMetricsEntry
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) StreamIn(handle string, spec garden.StreamInSpec) error {
	fake.streamInMutex.Lock()
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		handle string
		spec   garden.StreamInSpec
	}{handle, spec})
	fake.recordInvocation("StreamIn", []interface{}{handle, spec})
	fake.streamInMutex.Unlock()
	if fake.StreamInStub != nil {
		return fake.StreamInStub(handle, spec)
	} else {
		return fake.streamInReturns.result1
	}
}

func (fake *FakeConnection) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeConnection) StreamInArgsForCall(i int) (string, garden.StreamInSpec) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return fake.streamInArgsForCall[i].handle, fake.streamInArgsForCall[i].spec
}

func (fake *FakeConnection) StreamInReturns(result1 error) {
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnection) StreamOut(handle string, spec garden.StreamOutSpec) (io.ReadCloser, error) {
	fake.streamOutMutex.Lock()
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		handle string
		spec   garden.StreamOutSpec
	}{handle, spec})
	fake.recordInvocation("StreamOut", []interface{}{handle, spec})
	fake.streamOutMutex.Unlock()
	if fake.StreamOutStub != nil {
		return fake.StreamOutStub(handle, spec)
	} else {
		return fake.streamOutReturns.result1, fake.streamOutReturns.result2
	}
}

func (fake *FakeConnection) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeConnection) StreamOutArgsForCall(i int) (string, garden.StreamOutSpec) {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return fake.streamOutArgsForCall[i].handle, fake.streamOutArgsForCall[i].spec
}

func (fake *FakeConnection) StreamOutReturns(result1 io.ReadCloser, result2 error) {
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) CurrentBandwidthLimits(handle string) (garden.BandwidthLimits, error) {
	fake.currentBandwidthLimitsMutex.Lock()
	fake.currentBandwidthLimitsArgsForCall = append(fake.currentBandwidthLimitsArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("CurrentBandwidthLimits", []interface{}{handle})
	fake.currentBandwidthLimitsMutex.Unlock()
	if fake.CurrentBandwidthLimitsStub != nil {
		return fake.CurrentBandwidthLimitsStub(handle)
	} else {
		return fake.currentBandwidthLimitsReturns.result1, fake.currentBandwidthLimitsReturns.result2
	}
}

func (fake *FakeConnection) CurrentBandwidthLimitsCallCount() int {
	fake.currentBandwidthLimitsMutex.RLock()
	defer fake.currentBandwidthLimitsMutex.RUnlock()
	return len(fake.currentBandwidthLimitsArgsForCall)
}

func (fake *FakeConnection) CurrentBandwidthLimitsArgsForCall(i int) string {
	fake.currentBandwidthLimitsMutex.RLock()
	defer fake.currentBandwidthLimitsMutex.RUnlock()
	return fake.currentBandwidthLimitsArgsForCall[i].handle
}

func (fake *FakeConnection) CurrentBandwidthLimitsReturns(result1 garden.BandwidthLimits, result2 error) {
	fake.CurrentBandwidthLimitsStub = nil
	fake.currentBandwidthLimitsReturns = struct {
		result1 garden.BandwidthLimits
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) CurrentCPULimits(handle string) (garden.CPULimits, error) {
	fake.currentCPULimitsMutex.Lock()
	fake.currentCPULimitsArgsForCall = append(fake.currentCPULimitsArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("CurrentCPULimits", []interface{}{handle})
	fake.currentCPULimitsMutex.Unlock()
	if fake.CurrentCPULimitsStub != nil {
		return fake.CurrentCPULimitsStub(handle)
	} else {
		return fake.currentCPULimitsReturns.result1, fake.currentCPULimitsReturns.result2
	}
}

func (fake *FakeConnection) CurrentCPULimitsCallCount() int {
	fake.currentCPULimitsMutex.RLock()
	defer fake.currentCPULimitsMutex.RUnlock()
	return len(fake.currentCPULimitsArgsForCall)
}

func (fake *FakeConnection) CurrentCPULimitsArgsForCall(i int) string {
	fake.currentCPULimitsMutex.RLock()
	defer fake.currentCPULimitsMutex.RUnlock()
	return fake.currentCPULimitsArgsForCall[i].handle
}

func (fake *FakeConnection) CurrentCPULimitsReturns(result1 garden.CPULimits, result2 error) {
	fake.CurrentCPULimitsStub = nil
	fake.currentCPULimitsReturns = struct {
		result1 garden.CPULimits
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) CurrentDiskLimits(handle string) (garden.DiskLimits, error) {
	fake.currentDiskLimitsMutex.Lock()
	fake.currentDiskLimitsArgsForCall = append(fake.currentDiskLimitsArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("CurrentDiskLimits", []interface{}{handle})
	fake.currentDiskLimitsMutex.Unlock()
	if fake.CurrentDiskLimitsStub != nil {
		return fake.CurrentDiskLimitsStub(handle)
	} else {
		return fake.currentDiskLimitsReturns.result1, fake.currentDiskLimitsReturns.result2
	}
}

func (fake *FakeConnection) CurrentDiskLimitsCallCount() int {
	fake.currentDiskLimitsMutex.RLock()
	defer fake.currentDiskLimitsMutex.RUnlock()
	return len(fake.currentDiskLimitsArgsForCall)
}

func (fake *FakeConnection) CurrentDiskLimitsArgsForCall(i int) string {
	fake.currentDiskLimitsMutex.RLock()
	defer fake.currentDiskLimitsMutex.RUnlock()
	return fake.currentDiskLimitsArgsForCall[i].handle
}

func (fake *FakeConnection) CurrentDiskLimitsReturns(result1 garden.DiskLimits, result2 error) {
	fake.CurrentDiskLimitsStub = nil
	fake.currentDiskLimitsReturns = struct {
		result1 garden.DiskLimits
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) CurrentMemoryLimits(handle string) (garden.MemoryLimits, error) {
	fake.currentMemoryLimitsMutex.Lock()
	fake.currentMemoryLimitsArgsForCall = append(fake.currentMemoryLimitsArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("CurrentMemoryLimits", []interface{}{handle})
	fake.currentMemoryLimitsMutex.Unlock()
	if fake.CurrentMemoryLimitsStub != nil {
		return fake.CurrentMemoryLimitsStub(handle)
	} else {
		return fake.currentMemoryLimitsReturns.result1, fake.currentMemoryLimitsReturns.result2
	}
}

func (fake *FakeConnection) CurrentMemoryLimitsCallCount() int {
	fake.currentMemoryLimitsMutex.RLock()
	defer fake.currentMemoryLimitsMutex.RUnlock()
	return len(fake.currentMemoryLimitsArgsForCall)
}

func (fake *FakeConnection) CurrentMemoryLimitsArgsForCall(i int) string {
	fake.currentMemoryLimitsMutex.RLock()
	defer fake.currentMemoryLimitsMutex.RUnlock()
	return fake.currentMemoryLimitsArgsForCall[i].handle
}

func (fake *FakeConnection) CurrentMemoryLimitsReturns(result1 garden.MemoryLimits, result2 error) {
	fake.CurrentMemoryLimitsStub = nil
	fake.currentMemoryLimitsReturns = struct {
		result1 garden.MemoryLimits
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) Run(handle string, spec garden.ProcessSpec, io garden.ProcessIO) (garden.Process, error) {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		handle string
		spec   garden.ProcessSpec
		io     garden.ProcessIO
	}{handle, spec, io})
	fake.recordInvocation("Run", []interface{}{handle, spec, io})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(handle, spec, io)
	} else {
		return fake.runReturns.result1, fake.runReturns.result2
	}
}

func (fake *FakeConnection) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeConnection) RunArgsForCall(i int) (string, garden.ProcessSpec, garden.ProcessIO) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].handle, fake.runArgsForCall[i].spec, fake.runArgsForCall[i].io
}

func (fake *FakeConnection) RunReturns(result1 garden.Process, result2 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) Attach(handle string, processID string, io garden.ProcessIO) (garden.Process, error) {
	fake.attachMutex.Lock()
	fake.attachArgsForCall = append(fake.attachArgsForCall, struct {
		handle    string
		processID string
		io        garden.ProcessIO
	}{handle, processID, io})
	fake.recordInvocation("Attach", []interface{}{handle, processID, io})
	fake.attachMutex.Unlock()
	if fake.AttachStub != nil {
		return fake.AttachStub(handle, processID, io)
	} else {
		return fake.attachReturns.result1, fake.attachReturns.result2
	}
}

func (fake *FakeConnection) AttachCallCount() int {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return len(fake.attachArgsForCall)
}

func (fake *FakeConnection) AttachArgsForCall(i int) (string, string, garden.ProcessIO) {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return fake.attachArgsForCall[i].handle, fake.attachArgsForCall[i].processID, fake.attachArgsForCall[i].io
}

func (fake *FakeConnection) AttachReturns(result1 garden.Process, result2 error) {
	fake.AttachStub = nil
	fake.attachReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) NetIn(handle string, hostPort uint32, containerPort uint32) (uint32, uint32, error) {
	fake.netInMutex.Lock()
	fake.netInArgsForCall = append(fake.netInArgsForCall, struct {
		handle        string
		hostPort      uint32
		containerPort uint32
	}{handle, hostPort, containerPort})
	fake.recordInvocation("NetIn", []interface{}{handle, hostPort, containerPort})
	fake.netInMutex.Unlock()
	if fake.NetInStub != nil {
		return fake.NetInStub(handle, hostPort, containerPort)
	} else {
		return fake.netInReturns.result1, fake.netInReturns.result2, fake.netInReturns.result3
	}
}

func (fake *FakeConnection) NetInCallCount() int {
	fake.netInMutex.RLock()
	defer fake.netInMutex.RUnlock()
	return len(fake.netInArgsForCall)
}

func (fake *FakeConnection) NetInArgsForCall(i int) (string, uint32, uint32) {
	fake.netInMutex.RLock()
	defer fake.netInMutex.RUnlock()
	return fake.netInArgsForCall[i].handle, fake.netInArgsForCall[i].hostPort, fake.netInArgsForCall[i].containerPort
}

func (fake *FakeConnection) NetInReturns(result1 uint32, result2 uint32, result3 error) {
	fake.NetInStub = nil
	fake.netInReturns = struct {
		result1 uint32
		result2 uint32
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeConnection) NetOut(handle string, rule garden.NetOutRule) error {
	fake.netOutMutex.Lock()
	fake.netOutArgsForCall = append(fake.netOutArgsForCall, struct {
		handle string
		rule   garden.NetOutRule
	}{handle, rule})
	fake.recordInvocation("NetOut", []interface{}{handle, rule})
	fake.netOutMutex.Unlock()
	if fake.NetOutStub != nil {
		return fake.NetOutStub(handle, rule)
	} else {
		return fake.netOutReturns.result1
	}
}

func (fake *FakeConnection) NetOutCallCount() int {
	fake.netOutMutex.RLock()
	defer fake.netOutMutex.RUnlock()
	return len(fake.netOutArgsForCall)
}

func (fake *FakeConnection) NetOutArgsForCall(i int) (string, garden.NetOutRule) {
	fake.netOutMutex.RLock()
	defer fake.netOutMutex.RUnlock()
	return fake.netOutArgsForCall[i].handle, fake.netOutArgsForCall[i].rule
}

func (fake *FakeConnection) NetOutReturns(result1 error) {
	fake.NetOutStub = nil
	fake.netOutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnection) SetGraceTime(handle string, graceTime time.Duration) error {
	fake.setGraceTimeMutex.Lock()
	fake.setGraceTimeArgsForCall = append(fake.setGraceTimeArgsForCall, struct {
		handle    string
		graceTime time.Duration
	}{handle, graceTime})
	fake.recordInvocation("SetGraceTime", []interface{}{handle, graceTime})
	fake.setGraceTimeMutex.Unlock()
	if fake.SetGraceTimeStub != nil {
		return fake.SetGraceTimeStub(handle, graceTime)
	} else {
		return fake.setGraceTimeReturns.result1
	}
}

func (fake *FakeConnection) SetGraceTimeCallCount() int {
	fake.setGraceTimeMutex.RLock()
	defer fake.setGraceTimeMutex.RUnlock()
	return len(fake.setGraceTimeArgsForCall)
}

func (fake *FakeConnection) SetGraceTimeArgsForCall(i int) (string, time.Duration) {
	fake.setGraceTimeMutex.RLock()
	defer fake.setGraceTimeMutex.RUnlock()
	return fake.setGraceTimeArgsForCall[i].handle, fake.setGraceTimeArgsForCall[i].graceTime
}

func (fake *FakeConnection) SetGraceTimeReturns(result1 error) {
	fake.SetGraceTimeStub = nil
	fake.setGraceTimeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnection) Properties(handle string) (garden.Properties, error) {
	fake.propertiesMutex.Lock()
	fake.propertiesArgsForCall = append(fake.propertiesArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("Properties", []interface{}{handle})
	fake.propertiesMutex.Unlock()
	if fake.PropertiesStub != nil {
		return fake.PropertiesStub(handle)
	} else {
		return fake.propertiesReturns.result1, fake.propertiesReturns.result2
	}
}

func (fake *FakeConnection) PropertiesCallCount() int {
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	return len(fake.propertiesArgsForCall)
}

func (fake *FakeConnection) PropertiesArgsForCall(i int) string {
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	return fake.propertiesArgsForCall[i].handle
}

func (fake *FakeConnection) PropertiesReturns(result1 garden.Properties, result2 error) {
	fake.PropertiesStub = nil
	fake.propertiesReturns = struct {
		result1 garden.Properties
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) Property(handle string, name string) (string, error) {
	fake.propertyMutex.Lock()
	fake.propertyArgsForCall = append(fake.propertyArgsForCall, struct {
		handle string
		name   string
	}{handle, name})
	fake.recordInvocation("Property", []interface{}{handle, name})
	fake.propertyMutex.Unlock()
	if fake.PropertyStub != nil {
		return fake.PropertyStub(handle, name)
	} else {
		return fake.propertyReturns.result1, fake.propertyReturns.result2
	}
}

func (fake *FakeConnection) PropertyCallCount() int {
	fake.propertyMutex.RLock()
	defer fake.propertyMutex.RUnlock()
	return len(fake.propertyArgsForCall)
}

func (fake *FakeConnection) PropertyArgsForCall(i int) (string, string) {
	fake.propertyMutex.RLock()
	defer fake.propertyMutex.RUnlock()
	return fake.propertyArgsForCall[i].handle, fake.propertyArgsForCall[i].name
}

func (fake *FakeConnection) PropertyReturns(result1 string, result2 error) {
	fake.PropertyStub = nil
	fake.propertyReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) SetProperty(handle string, name string, value string) error {
	fake.setPropertyMutex.Lock()
	fake.setPropertyArgsForCall = append(fake.setPropertyArgsForCall, struct {
		handle string
		name   string
		value  string
	}{handle, name, value})
	fake.recordInvocation("SetProperty", []interface{}{handle, name, value})
	fake.setPropertyMutex.Unlock()
	if fake.SetPropertyStub != nil {
		return fake.SetPropertyStub(handle, name, value)
	} else {
		return fake.setPropertyReturns.result1
	}
}

func (fake *FakeConnection) SetPropertyCallCount() int {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return len(fake.setPropertyArgsForCall)
}

func (fake *FakeConnection) SetPropertyArgsForCall(i int) (string, string, string) {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return fake.setPropertyArgsForCall[i].handle, fake.setPropertyArgsForCall[i].name, fake.setPropertyArgsForCall[i].value
}

func (fake *FakeConnection) SetPropertyReturns(result1 error) {
	fake.SetPropertyStub = nil
	fake.setPropertyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnection) Metrics(handle string) (garden.Metrics, error) {
	fake.metricsMutex.Lock()
	fake.metricsArgsForCall = append(fake.metricsArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("Metrics", []interface{}{handle})
	fake.metricsMutex.Unlock()
	if fake.MetricsStub != nil {
		return fake.MetricsStub(handle)
	} else {
		return fake.metricsReturns.result1, fake.metricsReturns.result2
	}
}

func (fake *FakeConnection) MetricsCallCount() int {
	fake.metricsMutex.RLock()
	defer fake.metricsMutex.RUnlock()
	return len(fake.metricsArgsForCall)
}

func (fake *FakeConnection) MetricsArgsForCall(i int) string {
	fake.metricsMutex.RLock()
	defer fake.metricsMutex.RUnlock()
	return fake.metricsArgsForCall[i].handle
}

func (fake *FakeConnection) MetricsReturns(result1 garden.Metrics, result2 error) {
	fake.MetricsStub = nil
	fake.metricsReturns = struct {
		result1 garden.Metrics
		result2 error
	}{result1, result2}
}

func (fake *FakeConnection) RemoveProperty(handle string, name string) error {
	fake.removePropertyMutex.Lock()
	fake.removePropertyArgsForCall = append(fake.removePropertyArgsForCall, struct {
		handle string
		name   string
	}{handle, name})
	fake.recordInvocation("RemoveProperty", []interface{}{handle, name})
	fake.removePropertyMutex.Unlock()
	if fake.RemovePropertyStub != nil {
		return fake.RemovePropertyStub(handle, name)
	} else {
		return fake.removePropertyReturns.result1
	}
}

func (fake *FakeConnection) RemovePropertyCallCount() int {
	fake.removePropertyMutex.RLock()
	defer fake.removePropertyMutex.RUnlock()
	return len(fake.removePropertyArgsForCall)
}

func (fake *FakeConnection) RemovePropertyArgsForCall(i int) (string, string) {
	fake.removePropertyMutex.RLock()
	defer fake.removePropertyMutex.RUnlock()
	return fake.removePropertyArgsForCall[i].handle, fake.removePropertyArgsForCall[i].name
}

func (fake *FakeConnection) RemovePropertyReturns(result1 error) {
	fake.RemovePropertyStub = nil
	fake.removePropertyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConnection) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.capacityMutex.RLock()
	defer fake.capacityMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.bulkInfoMutex.RLock()
	defer fake.bulkInfoMutex.RUnlock()
	fake.bulkMetricsMutex.RLock()
	defer fake.bulkMetricsMutex.RUnlock()
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	fake.currentBandwidthLimitsMutex.RLock()
	defer fake.currentBandwidthLimitsMutex.RUnlock()
	fake.currentCPULimitsMutex.RLock()
	defer fake.currentCPULimitsMutex.RUnlock()
	fake.currentDiskLimitsMutex.RLock()
	defer fake.currentDiskLimitsMutex.RUnlock()
	fake.currentMemoryLimitsMutex.RLock()
	defer fake.currentMemoryLimitsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	fake.netInMutex.RLock()
	defer fake.netInMutex.RUnlock()
	fake.netOutMutex.RLock()
	defer fake.netOutMutex.RUnlock()
	fake.setGraceTimeMutex.RLock()
	defer fake.setGraceTimeMutex.RUnlock()
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	fake.propertyMutex.RLock()
	defer fake.propertyMutex.RUnlock()
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	fake.metricsMutex.RLock()
	defer fake.metricsMutex.RUnlock()
	fake.removePropertyMutex.RLock()
	defer fake.removePropertyMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConnection) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ connection.Connection = new(FakeConnection)
