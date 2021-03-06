package mocks

import "github.com/stretchr/testify/mock"

import "k8s.io/kubernetes/pkg/api"
import "k8s.io/kubernetes/pkg/watch"
import "k8s.io/kubernetes/pkg/client/restclient"

type KubePodInterface struct {
	mock.Mock
}

func (m *KubePodInterface) List(opts api.ListOptions) (*api.PodList, error) {
	ret := m.Called(opts)

	var r0 *api.PodList
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*api.PodList)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *KubePodInterface) Get(name string) (*api.Pod, error) {
	ret := m.Called(name)

	var r0 *api.Pod
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*api.Pod)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *KubePodInterface) Delete(name string, options *api.DeleteOptions) error {
	ret := m.Called(name, options)

	r0 := ret.Error(0)

	return r0
}
func (m *KubePodInterface) Create(pod *api.Pod) (*api.Pod, error) {
	ret := m.Called(pod)

	var r0 *api.Pod
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*api.Pod)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *KubePodInterface) Update(pod *api.Pod) (*api.Pod, error) {
	ret := m.Called(pod)

	var r0 *api.Pod
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*api.Pod)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *KubePodInterface) Watch(opts api.ListOptions) (watch.Interface, error) {
	ret := m.Called(opts)

	r0 := ret.Get(0).(watch.Interface)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *KubePodInterface) Bind(binding *api.Binding) error {
	ret := m.Called(binding)

	r0 := ret.Error(0)

	return r0
}
func (m *KubePodInterface) UpdateStatus(pod *api.Pod) (*api.Pod, error) {
	ret := m.Called(pod)

	var r0 *api.Pod
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*api.Pod)
	}
	r1 := ret.Error(1)

	return r0, r1
}

func (m *KubePodInterface) GetLogs(name string, opts *api.PodLogOptions) *restclient.Request {
	ret := m.Called(name, opts)

	r0 := ret.Get(0).(*restclient.Request)
	return r0
}
