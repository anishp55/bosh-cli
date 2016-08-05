// This file was generated by counterfeiter
package directorfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-init/director"
)

type FakeDeployment struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	ManifestStub        func() (string, error)
	manifestMutex       sync.RWMutex
	manifestArgsForCall []struct{}
	manifestReturns     struct {
		result1 string
		result2 error
	}
	CloudConfigStub        func() (string, error)
	cloudConfigMutex       sync.RWMutex
	cloudConfigArgsForCall []struct{}
	cloudConfigReturns     struct {
		result1 string
		result2 error
	}
	DiffStub        func([]byte, bool) (director.DiffLines, error)
	diffMutex       sync.RWMutex
	diffArgsForCall []struct {
		arg1 []byte
		arg2 bool
	}
	diffReturns struct {
		result1 director.DiffLines
		result2 error
	}
	ReleasesStub        func() ([]director.Release, error)
	releasesMutex       sync.RWMutex
	releasesArgsForCall []struct{}
	releasesReturns     struct {
		result1 []director.Release
		result2 error
	}
	ExportReleaseStub        func(director.ReleaseSlug, director.OSVersionSlug) (director.ExportReleaseResult, error)
	exportReleaseMutex       sync.RWMutex
	exportReleaseArgsForCall []struct {
		arg1 director.ReleaseSlug
		arg2 director.OSVersionSlug
	}
	exportReleaseReturns struct {
		result1 director.ExportReleaseResult
		result2 error
	}
	StemcellsStub        func() ([]director.Stemcell, error)
	stemcellsMutex       sync.RWMutex
	stemcellsArgsForCall []struct{}
	stemcellsReturns     struct {
		result1 []director.Stemcell
		result2 error
	}
	VMInfosStub        func() ([]director.VMInfo, error)
	vMInfosMutex       sync.RWMutex
	vMInfosArgsForCall []struct{}
	vMInfosReturns     struct {
		result1 []director.VMInfo
		result2 error
	}
	ErrandsStub        func() ([]director.Errand, error)
	errandsMutex       sync.RWMutex
	errandsArgsForCall []struct{}
	errandsReturns     struct {
		result1 []director.Errand
		result2 error
	}
	RunErrandStub        func(string, bool) (director.ErrandResult, error)
	runErrandMutex       sync.RWMutex
	runErrandArgsForCall []struct {
		arg1 string
		arg2 bool
	}
	runErrandReturns struct {
		result1 director.ErrandResult
		result2 error
	}
	ScanForProblemsStub        func() ([]director.Problem, error)
	scanForProblemsMutex       sync.RWMutex
	scanForProblemsArgsForCall []struct{}
	scanForProblemsReturns     struct {
		result1 []director.Problem
		result2 error
	}
	ResolveProblemsStub        func([]director.ProblemAnswer) error
	resolveProblemsMutex       sync.RWMutex
	resolveProblemsArgsForCall []struct {
		arg1 []director.ProblemAnswer
	}
	resolveProblemsReturns struct {
		result1 error
	}
	SnapshotsStub        func() ([]director.Snapshot, error)
	snapshotsMutex       sync.RWMutex
	snapshotsArgsForCall []struct{}
	snapshotsReturns     struct {
		result1 []director.Snapshot
		result2 error
	}
	TakeSnapshotsStub        func() error
	takeSnapshotsMutex       sync.RWMutex
	takeSnapshotsArgsForCall []struct{}
	takeSnapshotsReturns     struct {
		result1 error
	}
	DeleteSnapshotStub        func(string) error
	deleteSnapshotMutex       sync.RWMutex
	deleteSnapshotArgsForCall []struct {
		arg1 string
	}
	deleteSnapshotReturns struct {
		result1 error
	}
	DeleteSnapshotsStub        func() error
	deleteSnapshotsMutex       sync.RWMutex
	deleteSnapshotsArgsForCall []struct{}
	deleteSnapshotsReturns     struct {
		result1 error
	}
	StartStub        func(slug director.AllOrPoolOrInstanceSlug) error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		slug director.AllOrPoolOrInstanceSlug
	}
	startReturns struct {
		result1 error
	}
	StopStub        func(slug director.AllOrPoolOrInstanceSlug, hard bool, sd director.SkipDrain, force bool) error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		slug  director.AllOrPoolOrInstanceSlug
		hard  bool
		sd    director.SkipDrain
		force bool
	}
	stopReturns struct {
		result1 error
	}
	RestartStub        func(slug director.AllOrPoolOrInstanceSlug, sd director.SkipDrain, force bool) error
	restartMutex       sync.RWMutex
	restartArgsForCall []struct {
		slug  director.AllOrPoolOrInstanceSlug
		sd    director.SkipDrain
		force bool
	}
	restartReturns struct {
		result1 error
	}
	RecreateStub        func(slug director.AllOrPoolOrInstanceSlug, sd director.SkipDrain, force bool) error
	recreateMutex       sync.RWMutex
	recreateArgsForCall []struct {
		slug  director.AllOrPoolOrInstanceSlug
		sd    director.SkipDrain
		force bool
	}
	recreateReturns struct {
		result1 error
	}
	SetUpSSHStub        func(director.AllOrPoolOrInstanceSlug, director.SSHOpts) (director.SSHResult, error)
	setUpSSHMutex       sync.RWMutex
	setUpSSHArgsForCall []struct {
		arg1 director.AllOrPoolOrInstanceSlug
		arg2 director.SSHOpts
	}
	setUpSSHReturns struct {
		result1 director.SSHResult
		result2 error
	}
	CleanUpSSHStub        func(director.AllOrPoolOrInstanceSlug, director.SSHOpts) error
	cleanUpSSHMutex       sync.RWMutex
	cleanUpSSHArgsForCall []struct {
		arg1 director.AllOrPoolOrInstanceSlug
		arg2 director.SSHOpts
	}
	cleanUpSSHReturns struct {
		result1 error
	}
	FetchLogsStub        func(director.InstanceSlug, []string, bool) (director.LogsResult, error)
	fetchLogsMutex       sync.RWMutex
	fetchLogsArgsForCall []struct {
		arg1 director.InstanceSlug
		arg2 []string
		arg3 bool
	}
	fetchLogsReturns struct {
		result1 director.LogsResult
		result2 error
	}
	TakeSnapshotStub        func(director.InstanceSlug) error
	takeSnapshotMutex       sync.RWMutex
	takeSnapshotArgsForCall []struct {
		arg1 director.InstanceSlug
	}
	takeSnapshotReturns struct {
		result1 error
	}
	EnableResurrectionStub        func(director.InstanceSlug, bool) error
	enableResurrectionMutex       sync.RWMutex
	enableResurrectionArgsForCall []struct {
		arg1 director.InstanceSlug
		arg2 bool
	}
	enableResurrectionReturns struct {
		result1 error
	}
	UpdateStub        func(manifest []byte, recreate bool, sd director.SkipDrain) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		manifest []byte
		recreate bool
		sd       director.SkipDrain
	}
	updateReturns struct {
		result1 error
	}
	DeleteStub        func(force bool) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		force bool
	}
	deleteReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeployment) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeDeployment) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeDeployment) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeDeployment) Manifest() (string, error) {
	fake.manifestMutex.Lock()
	fake.manifestArgsForCall = append(fake.manifestArgsForCall, struct{}{})
	fake.recordInvocation("Manifest", []interface{}{})
	fake.manifestMutex.Unlock()
	if fake.ManifestStub != nil {
		return fake.ManifestStub()
	} else {
		return fake.manifestReturns.result1, fake.manifestReturns.result2
	}
}

func (fake *FakeDeployment) ManifestCallCount() int {
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	return len(fake.manifestArgsForCall)
}

func (fake *FakeDeployment) ManifestReturns(result1 string, result2 error) {
	fake.ManifestStub = nil
	fake.manifestReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) CloudConfig() (string, error) {
	fake.cloudConfigMutex.Lock()
	fake.cloudConfigArgsForCall = append(fake.cloudConfigArgsForCall, struct{}{})
	fake.recordInvocation("CloudConfig", []interface{}{})
	fake.cloudConfigMutex.Unlock()
	if fake.CloudConfigStub != nil {
		return fake.CloudConfigStub()
	} else {
		return fake.cloudConfigReturns.result1, fake.cloudConfigReturns.result2
	}
}

func (fake *FakeDeployment) CloudConfigCallCount() int {
	fake.cloudConfigMutex.RLock()
	defer fake.cloudConfigMutex.RUnlock()
	return len(fake.cloudConfigArgsForCall)
}

func (fake *FakeDeployment) CloudConfigReturns(result1 string, result2 error) {
	fake.CloudConfigStub = nil
	fake.cloudConfigReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) Diff(arg1 []byte, arg2 bool) (director.DiffLines, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.diffMutex.Lock()
	fake.diffArgsForCall = append(fake.diffArgsForCall, struct {
		arg1 []byte
		arg2 bool
	}{arg1Copy, arg2})
	fake.recordInvocation("Diff", []interface{}{arg1Copy, arg2})
	fake.diffMutex.Unlock()
	if fake.DiffStub != nil {
		return fake.DiffStub(arg1, arg2)
	} else {
		return fake.diffReturns.result1, fake.diffReturns.result2
	}
}

func (fake *FakeDeployment) DiffCallCount() int {
	fake.diffMutex.RLock()
	defer fake.diffMutex.RUnlock()
	return len(fake.diffArgsForCall)
}

func (fake *FakeDeployment) DiffArgsForCall(i int) ([]byte, bool) {
	fake.diffMutex.RLock()
	defer fake.diffMutex.RUnlock()
	return fake.diffArgsForCall[i].arg1, fake.diffArgsForCall[i].arg2
}

func (fake *FakeDeployment) DiffReturns(result1 director.DiffLines, result2 error) {
	fake.DiffStub = nil
	fake.diffReturns = struct {
		result1 director.DiffLines
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) Releases() ([]director.Release, error) {
	fake.releasesMutex.Lock()
	fake.releasesArgsForCall = append(fake.releasesArgsForCall, struct{}{})
	fake.recordInvocation("Releases", []interface{}{})
	fake.releasesMutex.Unlock()
	if fake.ReleasesStub != nil {
		return fake.ReleasesStub()
	} else {
		return fake.releasesReturns.result1, fake.releasesReturns.result2
	}
}

func (fake *FakeDeployment) ReleasesCallCount() int {
	fake.releasesMutex.RLock()
	defer fake.releasesMutex.RUnlock()
	return len(fake.releasesArgsForCall)
}

func (fake *FakeDeployment) ReleasesReturns(result1 []director.Release, result2 error) {
	fake.ReleasesStub = nil
	fake.releasesReturns = struct {
		result1 []director.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) ExportRelease(arg1 director.ReleaseSlug, arg2 director.OSVersionSlug) (director.ExportReleaseResult, error) {
	fake.exportReleaseMutex.Lock()
	fake.exportReleaseArgsForCall = append(fake.exportReleaseArgsForCall, struct {
		arg1 director.ReleaseSlug
		arg2 director.OSVersionSlug
	}{arg1, arg2})
	fake.recordInvocation("ExportRelease", []interface{}{arg1, arg2})
	fake.exportReleaseMutex.Unlock()
	if fake.ExportReleaseStub != nil {
		return fake.ExportReleaseStub(arg1, arg2)
	} else {
		return fake.exportReleaseReturns.result1, fake.exportReleaseReturns.result2
	}
}

func (fake *FakeDeployment) ExportReleaseCallCount() int {
	fake.exportReleaseMutex.RLock()
	defer fake.exportReleaseMutex.RUnlock()
	return len(fake.exportReleaseArgsForCall)
}

func (fake *FakeDeployment) ExportReleaseArgsForCall(i int) (director.ReleaseSlug, director.OSVersionSlug) {
	fake.exportReleaseMutex.RLock()
	defer fake.exportReleaseMutex.RUnlock()
	return fake.exportReleaseArgsForCall[i].arg1, fake.exportReleaseArgsForCall[i].arg2
}

func (fake *FakeDeployment) ExportReleaseReturns(result1 director.ExportReleaseResult, result2 error) {
	fake.ExportReleaseStub = nil
	fake.exportReleaseReturns = struct {
		result1 director.ExportReleaseResult
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) Stemcells() ([]director.Stemcell, error) {
	fake.stemcellsMutex.Lock()
	fake.stemcellsArgsForCall = append(fake.stemcellsArgsForCall, struct{}{})
	fake.recordInvocation("Stemcells", []interface{}{})
	fake.stemcellsMutex.Unlock()
	if fake.StemcellsStub != nil {
		return fake.StemcellsStub()
	} else {
		return fake.stemcellsReturns.result1, fake.stemcellsReturns.result2
	}
}

func (fake *FakeDeployment) StemcellsCallCount() int {
	fake.stemcellsMutex.RLock()
	defer fake.stemcellsMutex.RUnlock()
	return len(fake.stemcellsArgsForCall)
}

func (fake *FakeDeployment) StemcellsReturns(result1 []director.Stemcell, result2 error) {
	fake.StemcellsStub = nil
	fake.stemcellsReturns = struct {
		result1 []director.Stemcell
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) VMInfos() ([]director.VMInfo, error) {
	fake.vMInfosMutex.Lock()
	fake.vMInfosArgsForCall = append(fake.vMInfosArgsForCall, struct{}{})
	fake.recordInvocation("VMInfos", []interface{}{})
	fake.vMInfosMutex.Unlock()
	if fake.VMInfosStub != nil {
		return fake.VMInfosStub()
	} else {
		return fake.vMInfosReturns.result1, fake.vMInfosReturns.result2
	}
}

func (fake *FakeDeployment) VMInfosCallCount() int {
	fake.vMInfosMutex.RLock()
	defer fake.vMInfosMutex.RUnlock()
	return len(fake.vMInfosArgsForCall)
}

func (fake *FakeDeployment) VMInfosReturns(result1 []director.VMInfo, result2 error) {
	fake.VMInfosStub = nil
	fake.vMInfosReturns = struct {
		result1 []director.VMInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) Errands() ([]director.Errand, error) {
	fake.errandsMutex.Lock()
	fake.errandsArgsForCall = append(fake.errandsArgsForCall, struct{}{})
	fake.recordInvocation("Errands", []interface{}{})
	fake.errandsMutex.Unlock()
	if fake.ErrandsStub != nil {
		return fake.ErrandsStub()
	} else {
		return fake.errandsReturns.result1, fake.errandsReturns.result2
	}
}

func (fake *FakeDeployment) ErrandsCallCount() int {
	fake.errandsMutex.RLock()
	defer fake.errandsMutex.RUnlock()
	return len(fake.errandsArgsForCall)
}

func (fake *FakeDeployment) ErrandsReturns(result1 []director.Errand, result2 error) {
	fake.ErrandsStub = nil
	fake.errandsReturns = struct {
		result1 []director.Errand
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) RunErrand(arg1 string, arg2 bool) (director.ErrandResult, error) {
	fake.runErrandMutex.Lock()
	fake.runErrandArgsForCall = append(fake.runErrandArgsForCall, struct {
		arg1 string
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("RunErrand", []interface{}{arg1, arg2})
	fake.runErrandMutex.Unlock()
	if fake.RunErrandStub != nil {
		return fake.RunErrandStub(arg1, arg2)
	} else {
		return fake.runErrandReturns.result1, fake.runErrandReturns.result2
	}
}

func (fake *FakeDeployment) RunErrandCallCount() int {
	fake.runErrandMutex.RLock()
	defer fake.runErrandMutex.RUnlock()
	return len(fake.runErrandArgsForCall)
}

func (fake *FakeDeployment) RunErrandArgsForCall(i int) (string, bool) {
	fake.runErrandMutex.RLock()
	defer fake.runErrandMutex.RUnlock()
	return fake.runErrandArgsForCall[i].arg1, fake.runErrandArgsForCall[i].arg2
}

func (fake *FakeDeployment) RunErrandReturns(result1 director.ErrandResult, result2 error) {
	fake.RunErrandStub = nil
	fake.runErrandReturns = struct {
		result1 director.ErrandResult
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) ScanForProblems() ([]director.Problem, error) {
	fake.scanForProblemsMutex.Lock()
	fake.scanForProblemsArgsForCall = append(fake.scanForProblemsArgsForCall, struct{}{})
	fake.recordInvocation("ScanForProblems", []interface{}{})
	fake.scanForProblemsMutex.Unlock()
	if fake.ScanForProblemsStub != nil {
		return fake.ScanForProblemsStub()
	} else {
		return fake.scanForProblemsReturns.result1, fake.scanForProblemsReturns.result2
	}
}

func (fake *FakeDeployment) ScanForProblemsCallCount() int {
	fake.scanForProblemsMutex.RLock()
	defer fake.scanForProblemsMutex.RUnlock()
	return len(fake.scanForProblemsArgsForCall)
}

func (fake *FakeDeployment) ScanForProblemsReturns(result1 []director.Problem, result2 error) {
	fake.ScanForProblemsStub = nil
	fake.scanForProblemsReturns = struct {
		result1 []director.Problem
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) ResolveProblems(arg1 []director.ProblemAnswer) error {
	var arg1Copy []director.ProblemAnswer
	if arg1 != nil {
		arg1Copy = make([]director.ProblemAnswer, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.resolveProblemsMutex.Lock()
	fake.resolveProblemsArgsForCall = append(fake.resolveProblemsArgsForCall, struct {
		arg1 []director.ProblemAnswer
	}{arg1Copy})
	fake.recordInvocation("ResolveProblems", []interface{}{arg1Copy})
	fake.resolveProblemsMutex.Unlock()
	if fake.ResolveProblemsStub != nil {
		return fake.ResolveProblemsStub(arg1)
	} else {
		return fake.resolveProblemsReturns.result1
	}
}

func (fake *FakeDeployment) ResolveProblemsCallCount() int {
	fake.resolveProblemsMutex.RLock()
	defer fake.resolveProblemsMutex.RUnlock()
	return len(fake.resolveProblemsArgsForCall)
}

func (fake *FakeDeployment) ResolveProblemsArgsForCall(i int) []director.ProblemAnswer {
	fake.resolveProblemsMutex.RLock()
	defer fake.resolveProblemsMutex.RUnlock()
	return fake.resolveProblemsArgsForCall[i].arg1
}

func (fake *FakeDeployment) ResolveProblemsReturns(result1 error) {
	fake.ResolveProblemsStub = nil
	fake.resolveProblemsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Snapshots() ([]director.Snapshot, error) {
	fake.snapshotsMutex.Lock()
	fake.snapshotsArgsForCall = append(fake.snapshotsArgsForCall, struct{}{})
	fake.recordInvocation("Snapshots", []interface{}{})
	fake.snapshotsMutex.Unlock()
	if fake.SnapshotsStub != nil {
		return fake.SnapshotsStub()
	} else {
		return fake.snapshotsReturns.result1, fake.snapshotsReturns.result2
	}
}

func (fake *FakeDeployment) SnapshotsCallCount() int {
	fake.snapshotsMutex.RLock()
	defer fake.snapshotsMutex.RUnlock()
	return len(fake.snapshotsArgsForCall)
}

func (fake *FakeDeployment) SnapshotsReturns(result1 []director.Snapshot, result2 error) {
	fake.SnapshotsStub = nil
	fake.snapshotsReturns = struct {
		result1 []director.Snapshot
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) TakeSnapshots() error {
	fake.takeSnapshotsMutex.Lock()
	fake.takeSnapshotsArgsForCall = append(fake.takeSnapshotsArgsForCall, struct{}{})
	fake.recordInvocation("TakeSnapshots", []interface{}{})
	fake.takeSnapshotsMutex.Unlock()
	if fake.TakeSnapshotsStub != nil {
		return fake.TakeSnapshotsStub()
	} else {
		return fake.takeSnapshotsReturns.result1
	}
}

func (fake *FakeDeployment) TakeSnapshotsCallCount() int {
	fake.takeSnapshotsMutex.RLock()
	defer fake.takeSnapshotsMutex.RUnlock()
	return len(fake.takeSnapshotsArgsForCall)
}

func (fake *FakeDeployment) TakeSnapshotsReturns(result1 error) {
	fake.TakeSnapshotsStub = nil
	fake.takeSnapshotsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) DeleteSnapshot(arg1 string) error {
	fake.deleteSnapshotMutex.Lock()
	fake.deleteSnapshotArgsForCall = append(fake.deleteSnapshotArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteSnapshot", []interface{}{arg1})
	fake.deleteSnapshotMutex.Unlock()
	if fake.DeleteSnapshotStub != nil {
		return fake.DeleteSnapshotStub(arg1)
	} else {
		return fake.deleteSnapshotReturns.result1
	}
}

func (fake *FakeDeployment) DeleteSnapshotCallCount() int {
	fake.deleteSnapshotMutex.RLock()
	defer fake.deleteSnapshotMutex.RUnlock()
	return len(fake.deleteSnapshotArgsForCall)
}

func (fake *FakeDeployment) DeleteSnapshotArgsForCall(i int) string {
	fake.deleteSnapshotMutex.RLock()
	defer fake.deleteSnapshotMutex.RUnlock()
	return fake.deleteSnapshotArgsForCall[i].arg1
}

func (fake *FakeDeployment) DeleteSnapshotReturns(result1 error) {
	fake.DeleteSnapshotStub = nil
	fake.deleteSnapshotReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) DeleteSnapshots() error {
	fake.deleteSnapshotsMutex.Lock()
	fake.deleteSnapshotsArgsForCall = append(fake.deleteSnapshotsArgsForCall, struct{}{})
	fake.recordInvocation("DeleteSnapshots", []interface{}{})
	fake.deleteSnapshotsMutex.Unlock()
	if fake.DeleteSnapshotsStub != nil {
		return fake.DeleteSnapshotsStub()
	} else {
		return fake.deleteSnapshotsReturns.result1
	}
}

func (fake *FakeDeployment) DeleteSnapshotsCallCount() int {
	fake.deleteSnapshotsMutex.RLock()
	defer fake.deleteSnapshotsMutex.RUnlock()
	return len(fake.deleteSnapshotsArgsForCall)
}

func (fake *FakeDeployment) DeleteSnapshotsReturns(result1 error) {
	fake.DeleteSnapshotsStub = nil
	fake.deleteSnapshotsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Start(slug director.AllOrPoolOrInstanceSlug) error {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		slug director.AllOrPoolOrInstanceSlug
	}{slug})
	fake.recordInvocation("Start", []interface{}{slug})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub(slug)
	} else {
		return fake.startReturns.result1
	}
}

func (fake *FakeDeployment) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeDeployment) StartArgsForCall(i int) director.AllOrPoolOrInstanceSlug {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return fake.startArgsForCall[i].slug
}

func (fake *FakeDeployment) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Stop(slug director.AllOrPoolOrInstanceSlug, hard bool, sd director.SkipDrain, force bool) error {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		slug  director.AllOrPoolOrInstanceSlug
		hard  bool
		sd    director.SkipDrain
		force bool
	}{slug, hard, sd, force})
	fake.recordInvocation("Stop", []interface{}{slug, hard, sd, force})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub(slug, hard, sd, force)
	} else {
		return fake.stopReturns.result1
	}
}

func (fake *FakeDeployment) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeDeployment) StopArgsForCall(i int) (director.AllOrPoolOrInstanceSlug, bool, director.SkipDrain, bool) {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return fake.stopArgsForCall[i].slug, fake.stopArgsForCall[i].hard, fake.stopArgsForCall[i].sd, fake.stopArgsForCall[i].force
}

func (fake *FakeDeployment) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Restart(slug director.AllOrPoolOrInstanceSlug, sd director.SkipDrain, force bool) error {
	fake.restartMutex.Lock()
	fake.restartArgsForCall = append(fake.restartArgsForCall, struct {
		slug  director.AllOrPoolOrInstanceSlug
		sd    director.SkipDrain
		force bool
	}{slug, sd, force})
	fake.recordInvocation("Restart", []interface{}{slug, sd, force})
	fake.restartMutex.Unlock()
	if fake.RestartStub != nil {
		return fake.RestartStub(slug, sd, force)
	} else {
		return fake.restartReturns.result1
	}
}

func (fake *FakeDeployment) RestartCallCount() int {
	fake.restartMutex.RLock()
	defer fake.restartMutex.RUnlock()
	return len(fake.restartArgsForCall)
}

func (fake *FakeDeployment) RestartArgsForCall(i int) (director.AllOrPoolOrInstanceSlug, director.SkipDrain, bool) {
	fake.restartMutex.RLock()
	defer fake.restartMutex.RUnlock()
	return fake.restartArgsForCall[i].slug, fake.restartArgsForCall[i].sd, fake.restartArgsForCall[i].force
}

func (fake *FakeDeployment) RestartReturns(result1 error) {
	fake.RestartStub = nil
	fake.restartReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Recreate(slug director.AllOrPoolOrInstanceSlug, sd director.SkipDrain, force bool) error {
	fake.recreateMutex.Lock()
	fake.recreateArgsForCall = append(fake.recreateArgsForCall, struct {
		slug  director.AllOrPoolOrInstanceSlug
		sd    director.SkipDrain
		force bool
	}{slug, sd, force})
	fake.recordInvocation("Recreate", []interface{}{slug, sd, force})
	fake.recreateMutex.Unlock()
	if fake.RecreateStub != nil {
		return fake.RecreateStub(slug, sd, force)
	} else {
		return fake.recreateReturns.result1
	}
}

func (fake *FakeDeployment) RecreateCallCount() int {
	fake.recreateMutex.RLock()
	defer fake.recreateMutex.RUnlock()
	return len(fake.recreateArgsForCall)
}

func (fake *FakeDeployment) RecreateArgsForCall(i int) (director.AllOrPoolOrInstanceSlug, director.SkipDrain, bool) {
	fake.recreateMutex.RLock()
	defer fake.recreateMutex.RUnlock()
	return fake.recreateArgsForCall[i].slug, fake.recreateArgsForCall[i].sd, fake.recreateArgsForCall[i].force
}

func (fake *FakeDeployment) RecreateReturns(result1 error) {
	fake.RecreateStub = nil
	fake.recreateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) SetUpSSH(arg1 director.AllOrPoolOrInstanceSlug, arg2 director.SSHOpts) (director.SSHResult, error) {
	fake.setUpSSHMutex.Lock()
	fake.setUpSSHArgsForCall = append(fake.setUpSSHArgsForCall, struct {
		arg1 director.AllOrPoolOrInstanceSlug
		arg2 director.SSHOpts
	}{arg1, arg2})
	fake.recordInvocation("SetUpSSH", []interface{}{arg1, arg2})
	fake.setUpSSHMutex.Unlock()
	if fake.SetUpSSHStub != nil {
		return fake.SetUpSSHStub(arg1, arg2)
	} else {
		return fake.setUpSSHReturns.result1, fake.setUpSSHReturns.result2
	}
}

func (fake *FakeDeployment) SetUpSSHCallCount() int {
	fake.setUpSSHMutex.RLock()
	defer fake.setUpSSHMutex.RUnlock()
	return len(fake.setUpSSHArgsForCall)
}

func (fake *FakeDeployment) SetUpSSHArgsForCall(i int) (director.AllOrPoolOrInstanceSlug, director.SSHOpts) {
	fake.setUpSSHMutex.RLock()
	defer fake.setUpSSHMutex.RUnlock()
	return fake.setUpSSHArgsForCall[i].arg1, fake.setUpSSHArgsForCall[i].arg2
}

func (fake *FakeDeployment) SetUpSSHReturns(result1 director.SSHResult, result2 error) {
	fake.SetUpSSHStub = nil
	fake.setUpSSHReturns = struct {
		result1 director.SSHResult
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) CleanUpSSH(arg1 director.AllOrPoolOrInstanceSlug, arg2 director.SSHOpts) error {
	fake.cleanUpSSHMutex.Lock()
	fake.cleanUpSSHArgsForCall = append(fake.cleanUpSSHArgsForCall, struct {
		arg1 director.AllOrPoolOrInstanceSlug
		arg2 director.SSHOpts
	}{arg1, arg2})
	fake.recordInvocation("CleanUpSSH", []interface{}{arg1, arg2})
	fake.cleanUpSSHMutex.Unlock()
	if fake.CleanUpSSHStub != nil {
		return fake.CleanUpSSHStub(arg1, arg2)
	} else {
		return fake.cleanUpSSHReturns.result1
	}
}

func (fake *FakeDeployment) CleanUpSSHCallCount() int {
	fake.cleanUpSSHMutex.RLock()
	defer fake.cleanUpSSHMutex.RUnlock()
	return len(fake.cleanUpSSHArgsForCall)
}

func (fake *FakeDeployment) CleanUpSSHArgsForCall(i int) (director.AllOrPoolOrInstanceSlug, director.SSHOpts) {
	fake.cleanUpSSHMutex.RLock()
	defer fake.cleanUpSSHMutex.RUnlock()
	return fake.cleanUpSSHArgsForCall[i].arg1, fake.cleanUpSSHArgsForCall[i].arg2
}

func (fake *FakeDeployment) CleanUpSSHReturns(result1 error) {
	fake.CleanUpSSHStub = nil
	fake.cleanUpSSHReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) FetchLogs(arg1 director.InstanceSlug, arg2 []string, arg3 bool) (director.LogsResult, error) {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.fetchLogsMutex.Lock()
	fake.fetchLogsArgsForCall = append(fake.fetchLogsArgsForCall, struct {
		arg1 director.InstanceSlug
		arg2 []string
		arg3 bool
	}{arg1, arg2Copy, arg3})
	fake.recordInvocation("FetchLogs", []interface{}{arg1, arg2Copy, arg3})
	fake.fetchLogsMutex.Unlock()
	if fake.FetchLogsStub != nil {
		return fake.FetchLogsStub(arg1, arg2, arg3)
	} else {
		return fake.fetchLogsReturns.result1, fake.fetchLogsReturns.result2
	}
}

func (fake *FakeDeployment) FetchLogsCallCount() int {
	fake.fetchLogsMutex.RLock()
	defer fake.fetchLogsMutex.RUnlock()
	return len(fake.fetchLogsArgsForCall)
}

func (fake *FakeDeployment) FetchLogsArgsForCall(i int) (director.InstanceSlug, []string, bool) {
	fake.fetchLogsMutex.RLock()
	defer fake.fetchLogsMutex.RUnlock()
	return fake.fetchLogsArgsForCall[i].arg1, fake.fetchLogsArgsForCall[i].arg2, fake.fetchLogsArgsForCall[i].arg3
}

func (fake *FakeDeployment) FetchLogsReturns(result1 director.LogsResult, result2 error) {
	fake.FetchLogsStub = nil
	fake.fetchLogsReturns = struct {
		result1 director.LogsResult
		result2 error
	}{result1, result2}
}

func (fake *FakeDeployment) TakeSnapshot(arg1 director.InstanceSlug) error {
	fake.takeSnapshotMutex.Lock()
	fake.takeSnapshotArgsForCall = append(fake.takeSnapshotArgsForCall, struct {
		arg1 director.InstanceSlug
	}{arg1})
	fake.recordInvocation("TakeSnapshot", []interface{}{arg1})
	fake.takeSnapshotMutex.Unlock()
	if fake.TakeSnapshotStub != nil {
		return fake.TakeSnapshotStub(arg1)
	} else {
		return fake.takeSnapshotReturns.result1
	}
}

func (fake *FakeDeployment) TakeSnapshotCallCount() int {
	fake.takeSnapshotMutex.RLock()
	defer fake.takeSnapshotMutex.RUnlock()
	return len(fake.takeSnapshotArgsForCall)
}

func (fake *FakeDeployment) TakeSnapshotArgsForCall(i int) director.InstanceSlug {
	fake.takeSnapshotMutex.RLock()
	defer fake.takeSnapshotMutex.RUnlock()
	return fake.takeSnapshotArgsForCall[i].arg1
}

func (fake *FakeDeployment) TakeSnapshotReturns(result1 error) {
	fake.TakeSnapshotStub = nil
	fake.takeSnapshotReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) EnableResurrection(arg1 director.InstanceSlug, arg2 bool) error {
	fake.enableResurrectionMutex.Lock()
	fake.enableResurrectionArgsForCall = append(fake.enableResurrectionArgsForCall, struct {
		arg1 director.InstanceSlug
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("EnableResurrection", []interface{}{arg1, arg2})
	fake.enableResurrectionMutex.Unlock()
	if fake.EnableResurrectionStub != nil {
		return fake.EnableResurrectionStub(arg1, arg2)
	} else {
		return fake.enableResurrectionReturns.result1
	}
}

func (fake *FakeDeployment) EnableResurrectionCallCount() int {
	fake.enableResurrectionMutex.RLock()
	defer fake.enableResurrectionMutex.RUnlock()
	return len(fake.enableResurrectionArgsForCall)
}

func (fake *FakeDeployment) EnableResurrectionArgsForCall(i int) (director.InstanceSlug, bool) {
	fake.enableResurrectionMutex.RLock()
	defer fake.enableResurrectionMutex.RUnlock()
	return fake.enableResurrectionArgsForCall[i].arg1, fake.enableResurrectionArgsForCall[i].arg2
}

func (fake *FakeDeployment) EnableResurrectionReturns(result1 error) {
	fake.EnableResurrectionStub = nil
	fake.enableResurrectionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Update(manifest []byte, recreate bool, sd director.SkipDrain) error {
	var manifestCopy []byte
	if manifest != nil {
		manifestCopy = make([]byte, len(manifest))
		copy(manifestCopy, manifest)
	}
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		manifest []byte
		recreate bool
		sd       director.SkipDrain
	}{manifestCopy, recreate, sd})
	fake.recordInvocation("Update", []interface{}{manifestCopy, recreate, sd})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(manifest, recreate, sd)
	} else {
		return fake.updateReturns.result1
	}
}

func (fake *FakeDeployment) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeDeployment) UpdateArgsForCall(i int) ([]byte, bool, director.SkipDrain) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].manifest, fake.updateArgsForCall[i].recreate, fake.updateArgsForCall[i].sd
}

func (fake *FakeDeployment) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Delete(force bool) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		force bool
	}{force})
	fake.recordInvocation("Delete", []interface{}{force})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(force)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeDeployment) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeDeployment) DeleteArgsForCall(i int) bool {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].force
}

func (fake *FakeDeployment) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	fake.cloudConfigMutex.RLock()
	defer fake.cloudConfigMutex.RUnlock()
	fake.diffMutex.RLock()
	defer fake.diffMutex.RUnlock()
	fake.releasesMutex.RLock()
	defer fake.releasesMutex.RUnlock()
	fake.exportReleaseMutex.RLock()
	defer fake.exportReleaseMutex.RUnlock()
	fake.stemcellsMutex.RLock()
	defer fake.stemcellsMutex.RUnlock()
	fake.vMInfosMutex.RLock()
	defer fake.vMInfosMutex.RUnlock()
	fake.errandsMutex.RLock()
	defer fake.errandsMutex.RUnlock()
	fake.runErrandMutex.RLock()
	defer fake.runErrandMutex.RUnlock()
	fake.scanForProblemsMutex.RLock()
	defer fake.scanForProblemsMutex.RUnlock()
	fake.resolveProblemsMutex.RLock()
	defer fake.resolveProblemsMutex.RUnlock()
	fake.snapshotsMutex.RLock()
	defer fake.snapshotsMutex.RUnlock()
	fake.takeSnapshotsMutex.RLock()
	defer fake.takeSnapshotsMutex.RUnlock()
	fake.deleteSnapshotMutex.RLock()
	defer fake.deleteSnapshotMutex.RUnlock()
	fake.deleteSnapshotsMutex.RLock()
	defer fake.deleteSnapshotsMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.restartMutex.RLock()
	defer fake.restartMutex.RUnlock()
	fake.recreateMutex.RLock()
	defer fake.recreateMutex.RUnlock()
	fake.setUpSSHMutex.RLock()
	defer fake.setUpSSHMutex.RUnlock()
	fake.cleanUpSSHMutex.RLock()
	defer fake.cleanUpSSHMutex.RUnlock()
	fake.fetchLogsMutex.RLock()
	defer fake.fetchLogsMutex.RUnlock()
	fake.takeSnapshotMutex.RLock()
	defer fake.takeSnapshotMutex.RUnlock()
	fake.enableResurrectionMutex.RLock()
	defer fake.enableResurrectionMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDeployment) recordInvocation(key string, args []interface{}) {
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

var _ director.Deployment = new(FakeDeployment)