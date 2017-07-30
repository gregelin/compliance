package mocks

import "github.com/stretchr/testify/mock"

import "github.com/opencontrol/compliance-masonry/lib/common"

// Downloader is an autogenerated mock type for the Downloader type
type Downloader struct {
	mock.Mock
}

// DownloadRepo provides a mock function with given fields: _a0, _a1
func (_m *Downloader) DownloadRepo(_a0 common.RemoteSource, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.RemoteSource, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
