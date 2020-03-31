// Code generated by counterfeiter. DO NOT EDIT.
package smbbrokerfakes

import (
	"context"
	"sync"

	v1a "k8s.io/api/core/v1"
	v1b "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type FakePersistentVolumeClaimInterface struct {
	CreateStub        func(context.Context, *v1a.PersistentVolumeClaim, v1b.CreateOptions) (*v1a.PersistentVolumeClaim, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 *v1a.PersistentVolumeClaim
		arg3 v1b.CreateOptions
	}
	createReturns struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}
	DeleteStub        func(context.Context, string, v1b.DeleteOptions) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 v1b.DeleteOptions
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteCollectionStub        func(context.Context, v1b.DeleteOptions, v1b.ListOptions) error
	deleteCollectionMutex       sync.RWMutex
	deleteCollectionArgsForCall []struct {
		arg1 context.Context
		arg2 v1b.DeleteOptions
		arg3 v1b.ListOptions
	}
	deleteCollectionReturns struct {
		result1 error
	}
	deleteCollectionReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(context.Context, string, v1b.GetOptions) (*v1a.PersistentVolumeClaim, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 v1b.GetOptions
	}
	getReturns struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}
	ListStub        func(context.Context, v1b.ListOptions) (*v1a.PersistentVolumeClaimList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
		arg2 v1b.ListOptions
	}
	listReturns struct {
		result1 *v1a.PersistentVolumeClaimList
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *v1a.PersistentVolumeClaimList
		result2 error
	}
	PatchStub        func(context.Context, string, types.PatchType, []byte, v1b.PatchOptions, ...string) (*v1a.PersistentVolumeClaim, error)
	patchMutex       sync.RWMutex
	patchArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 types.PatchType
		arg4 []byte
		arg5 v1b.PatchOptions
		arg6 []string
	}
	patchReturns struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}
	patchReturnsOnCall map[int]struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}
	UpdateStub        func(context.Context, *v1a.PersistentVolumeClaim, v1b.UpdateOptions) (*v1a.PersistentVolumeClaim, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 *v1a.PersistentVolumeClaim
		arg3 v1b.UpdateOptions
	}
	updateReturns struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}
	UpdateStatusStub        func(context.Context, *v1a.PersistentVolumeClaim, v1b.UpdateOptions) (*v1a.PersistentVolumeClaim, error)
	updateStatusMutex       sync.RWMutex
	updateStatusArgsForCall []struct {
		arg1 context.Context
		arg2 *v1a.PersistentVolumeClaim
		arg3 v1b.UpdateOptions
	}
	updateStatusReturns struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}
	updateStatusReturnsOnCall map[int]struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}
	WatchStub        func(context.Context, v1b.ListOptions) (watch.Interface, error)
	watchMutex       sync.RWMutex
	watchArgsForCall []struct {
		arg1 context.Context
		arg2 v1b.ListOptions
	}
	watchReturns struct {
		result1 watch.Interface
		result2 error
	}
	watchReturnsOnCall map[int]struct {
		result1 watch.Interface
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePersistentVolumeClaimInterface) Create(arg1 context.Context, arg2 *v1a.PersistentVolumeClaim, arg3 v1b.CreateOptions) (*v1a.PersistentVolumeClaim, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 *v1a.PersistentVolumeClaim
		arg3 v1b.CreateOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistentVolumeClaimInterface) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakePersistentVolumeClaimInterface) CreateCalls(stub func(context.Context, *v1a.PersistentVolumeClaim, v1b.CreateOptions) (*v1a.PersistentVolumeClaim, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakePersistentVolumeClaimInterface) CreateArgsForCall(i int) (context.Context, *v1a.PersistentVolumeClaim, v1b.CreateOptions) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistentVolumeClaimInterface) CreateReturns(result1 *v1a.PersistentVolumeClaim, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) CreateReturnsOnCall(i int, result1 *v1a.PersistentVolumeClaim, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *v1a.PersistentVolumeClaim
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) Delete(arg1 context.Context, arg2 string, arg3 v1b.DeleteOptions) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 v1b.DeleteOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2, arg3})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakePersistentVolumeClaimInterface) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakePersistentVolumeClaimInterface) DeleteCalls(stub func(context.Context, string, v1b.DeleteOptions) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakePersistentVolumeClaimInterface) DeleteArgsForCall(i int) (context.Context, string, v1b.DeleteOptions) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistentVolumeClaimInterface) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePersistentVolumeClaimInterface) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePersistentVolumeClaimInterface) DeleteCollection(arg1 context.Context, arg2 v1b.DeleteOptions, arg3 v1b.ListOptions) error {
	fake.deleteCollectionMutex.Lock()
	ret, specificReturn := fake.deleteCollectionReturnsOnCall[len(fake.deleteCollectionArgsForCall)]
	fake.deleteCollectionArgsForCall = append(fake.deleteCollectionArgsForCall, struct {
		arg1 context.Context
		arg2 v1b.DeleteOptions
		arg3 v1b.ListOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("DeleteCollection", []interface{}{arg1, arg2, arg3})
	fake.deleteCollectionMutex.Unlock()
	if fake.DeleteCollectionStub != nil {
		return fake.DeleteCollectionStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteCollectionReturns
	return fakeReturns.result1
}

func (fake *FakePersistentVolumeClaimInterface) DeleteCollectionCallCount() int {
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	return len(fake.deleteCollectionArgsForCall)
}

func (fake *FakePersistentVolumeClaimInterface) DeleteCollectionCalls(stub func(context.Context, v1b.DeleteOptions, v1b.ListOptions) error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = stub
}

func (fake *FakePersistentVolumeClaimInterface) DeleteCollectionArgsForCall(i int) (context.Context, v1b.DeleteOptions, v1b.ListOptions) {
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	argsForCall := fake.deleteCollectionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistentVolumeClaimInterface) DeleteCollectionReturns(result1 error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = nil
	fake.deleteCollectionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePersistentVolumeClaimInterface) DeleteCollectionReturnsOnCall(i int, result1 error) {
	fake.deleteCollectionMutex.Lock()
	defer fake.deleteCollectionMutex.Unlock()
	fake.DeleteCollectionStub = nil
	if fake.deleteCollectionReturnsOnCall == nil {
		fake.deleteCollectionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteCollectionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePersistentVolumeClaimInterface) Get(arg1 context.Context, arg2 string, arg3 v1b.GetOptions) (*v1a.PersistentVolumeClaim, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 v1b.GetOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistentVolumeClaimInterface) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakePersistentVolumeClaimInterface) GetCalls(stub func(context.Context, string, v1b.GetOptions) (*v1a.PersistentVolumeClaim, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakePersistentVolumeClaimInterface) GetArgsForCall(i int) (context.Context, string, v1b.GetOptions) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistentVolumeClaimInterface) GetReturns(result1 *v1a.PersistentVolumeClaim, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) GetReturnsOnCall(i int, result1 *v1a.PersistentVolumeClaim, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *v1a.PersistentVolumeClaim
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) List(arg1 context.Context, arg2 v1b.ListOptions) (*v1a.PersistentVolumeClaimList, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
		arg2 v1b.ListOptions
	}{arg1, arg2})
	fake.recordInvocation("List", []interface{}{arg1, arg2})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistentVolumeClaimInterface) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakePersistentVolumeClaimInterface) ListCalls(stub func(context.Context, v1b.ListOptions) (*v1a.PersistentVolumeClaimList, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakePersistentVolumeClaimInterface) ListArgsForCall(i int) (context.Context, v1b.ListOptions) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePersistentVolumeClaimInterface) ListReturns(result1 *v1a.PersistentVolumeClaimList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *v1a.PersistentVolumeClaimList
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) ListReturnsOnCall(i int, result1 *v1a.PersistentVolumeClaimList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *v1a.PersistentVolumeClaimList
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *v1a.PersistentVolumeClaimList
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) Patch(arg1 context.Context, arg2 string, arg3 types.PatchType, arg4 []byte, arg5 v1b.PatchOptions, arg6 ...string) (*v1a.PersistentVolumeClaim, error) {
	var arg4Copy []byte
	if arg4 != nil {
		arg4Copy = make([]byte, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.patchMutex.Lock()
	ret, specificReturn := fake.patchReturnsOnCall[len(fake.patchArgsForCall)]
	fake.patchArgsForCall = append(fake.patchArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 types.PatchType
		arg4 []byte
		arg5 v1b.PatchOptions
		arg6 []string
	}{arg1, arg2, arg3, arg4Copy, arg5, arg6})
	fake.recordInvocation("Patch", []interface{}{arg1, arg2, arg3, arg4Copy, arg5, arg6})
	fake.patchMutex.Unlock()
	if fake.PatchStub != nil {
		return fake.PatchStub(arg1, arg2, arg3, arg4, arg5, arg6...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.patchReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistentVolumeClaimInterface) PatchCallCount() int {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	return len(fake.patchArgsForCall)
}

func (fake *FakePersistentVolumeClaimInterface) PatchCalls(stub func(context.Context, string, types.PatchType, []byte, v1b.PatchOptions, ...string) (*v1a.PersistentVolumeClaim, error)) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = stub
}

func (fake *FakePersistentVolumeClaimInterface) PatchArgsForCall(i int) (context.Context, string, types.PatchType, []byte, v1b.PatchOptions, []string) {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	argsForCall := fake.patchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakePersistentVolumeClaimInterface) PatchReturns(result1 *v1a.PersistentVolumeClaim, result2 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	fake.patchReturns = struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) PatchReturnsOnCall(i int, result1 *v1a.PersistentVolumeClaim, result2 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	if fake.patchReturnsOnCall == nil {
		fake.patchReturnsOnCall = make(map[int]struct {
			result1 *v1a.PersistentVolumeClaim
			result2 error
		})
	}
	fake.patchReturnsOnCall[i] = struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) Update(arg1 context.Context, arg2 *v1a.PersistentVolumeClaim, arg3 v1b.UpdateOptions) (*v1a.PersistentVolumeClaim, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 *v1a.PersistentVolumeClaim
		arg3 v1b.UpdateOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistentVolumeClaimInterface) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakePersistentVolumeClaimInterface) UpdateCalls(stub func(context.Context, *v1a.PersistentVolumeClaim, v1b.UpdateOptions) (*v1a.PersistentVolumeClaim, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakePersistentVolumeClaimInterface) UpdateArgsForCall(i int) (context.Context, *v1a.PersistentVolumeClaim, v1b.UpdateOptions) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistentVolumeClaimInterface) UpdateReturns(result1 *v1a.PersistentVolumeClaim, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) UpdateReturnsOnCall(i int, result1 *v1a.PersistentVolumeClaim, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *v1a.PersistentVolumeClaim
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) UpdateStatus(arg1 context.Context, arg2 *v1a.PersistentVolumeClaim, arg3 v1b.UpdateOptions) (*v1a.PersistentVolumeClaim, error) {
	fake.updateStatusMutex.Lock()
	ret, specificReturn := fake.updateStatusReturnsOnCall[len(fake.updateStatusArgsForCall)]
	fake.updateStatusArgsForCall = append(fake.updateStatusArgsForCall, struct {
		arg1 context.Context
		arg2 *v1a.PersistentVolumeClaim
		arg3 v1b.UpdateOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("UpdateStatus", []interface{}{arg1, arg2, arg3})
	fake.updateStatusMutex.Unlock()
	if fake.UpdateStatusStub != nil {
		return fake.UpdateStatusStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateStatusReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistentVolumeClaimInterface) UpdateStatusCallCount() int {
	fake.updateStatusMutex.RLock()
	defer fake.updateStatusMutex.RUnlock()
	return len(fake.updateStatusArgsForCall)
}

func (fake *FakePersistentVolumeClaimInterface) UpdateStatusCalls(stub func(context.Context, *v1a.PersistentVolumeClaim, v1b.UpdateOptions) (*v1a.PersistentVolumeClaim, error)) {
	fake.updateStatusMutex.Lock()
	defer fake.updateStatusMutex.Unlock()
	fake.UpdateStatusStub = stub
}

func (fake *FakePersistentVolumeClaimInterface) UpdateStatusArgsForCall(i int) (context.Context, *v1a.PersistentVolumeClaim, v1b.UpdateOptions) {
	fake.updateStatusMutex.RLock()
	defer fake.updateStatusMutex.RUnlock()
	argsForCall := fake.updateStatusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePersistentVolumeClaimInterface) UpdateStatusReturns(result1 *v1a.PersistentVolumeClaim, result2 error) {
	fake.updateStatusMutex.Lock()
	defer fake.updateStatusMutex.Unlock()
	fake.UpdateStatusStub = nil
	fake.updateStatusReturns = struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) UpdateStatusReturnsOnCall(i int, result1 *v1a.PersistentVolumeClaim, result2 error) {
	fake.updateStatusMutex.Lock()
	defer fake.updateStatusMutex.Unlock()
	fake.UpdateStatusStub = nil
	if fake.updateStatusReturnsOnCall == nil {
		fake.updateStatusReturnsOnCall = make(map[int]struct {
			result1 *v1a.PersistentVolumeClaim
			result2 error
		})
	}
	fake.updateStatusReturnsOnCall[i] = struct {
		result1 *v1a.PersistentVolumeClaim
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) Watch(arg1 context.Context, arg2 v1b.ListOptions) (watch.Interface, error) {
	fake.watchMutex.Lock()
	ret, specificReturn := fake.watchReturnsOnCall[len(fake.watchArgsForCall)]
	fake.watchArgsForCall = append(fake.watchArgsForCall, struct {
		arg1 context.Context
		arg2 v1b.ListOptions
	}{arg1, arg2})
	fake.recordInvocation("Watch", []interface{}{arg1, arg2})
	fake.watchMutex.Unlock()
	if fake.WatchStub != nil {
		return fake.WatchStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.watchReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePersistentVolumeClaimInterface) WatchCallCount() int {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	return len(fake.watchArgsForCall)
}

func (fake *FakePersistentVolumeClaimInterface) WatchCalls(stub func(context.Context, v1b.ListOptions) (watch.Interface, error)) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = stub
}

func (fake *FakePersistentVolumeClaimInterface) WatchArgsForCall(i int) (context.Context, v1b.ListOptions) {
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	argsForCall := fake.watchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePersistentVolumeClaimInterface) WatchReturns(result1 watch.Interface, result2 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	fake.watchReturns = struct {
		result1 watch.Interface
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) WatchReturnsOnCall(i int, result1 watch.Interface, result2 error) {
	fake.watchMutex.Lock()
	defer fake.watchMutex.Unlock()
	fake.WatchStub = nil
	if fake.watchReturnsOnCall == nil {
		fake.watchReturnsOnCall = make(map[int]struct {
			result1 watch.Interface
			result2 error
		})
	}
	fake.watchReturnsOnCall[i] = struct {
		result1 watch.Interface
		result2 error
	}{result1, result2}
}

func (fake *FakePersistentVolumeClaimInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteCollectionMutex.RLock()
	defer fake.deleteCollectionMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.updateStatusMutex.RLock()
	defer fake.updateStatusMutex.RUnlock()
	fake.watchMutex.RLock()
	defer fake.watchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePersistentVolumeClaimInterface) recordInvocation(key string, args []interface{}) {
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

var _ v1.PersistentVolumeClaimInterface = new(FakePersistentVolumeClaimInterface)
