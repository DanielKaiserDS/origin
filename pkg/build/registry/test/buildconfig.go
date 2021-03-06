package test

import (
	"sync"

	kapierrors "k8s.io/apimachinery/pkg/api/errors"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"

	"github.com/openshift/origin/pkg/build/api"
)

type BuildConfigRegistry struct {
	Err             error
	BuildConfigs    *api.BuildConfigList
	BuildConfig     *api.BuildConfig
	DeletedConfigID string
	sync.Mutex
}

func (r *BuildConfigRegistry) ListBuildConfigs(ctx apirequest.Context, options *metainternal.ListOptions) (*api.BuildConfigList, error) {
	r.Lock()
	defer r.Unlock()
	return r.BuildConfigs, r.Err
}

func (r *BuildConfigRegistry) GetBuildConfig(ctx apirequest.Context, id string, options *metav1.GetOptions) (*api.BuildConfig, error) {
	r.Lock()
	defer r.Unlock()
	if r.BuildConfig != nil && r.BuildConfig.Name == id {
		return r.BuildConfig, r.Err
	}
	return nil, kapierrors.NewNotFound(api.Resource("buildconfig"), id)
}

func (r *BuildConfigRegistry) CreateBuildConfig(ctx apirequest.Context, config *api.BuildConfig) error {
	r.Lock()
	defer r.Unlock()
	r.BuildConfig = config
	return r.Err
}

func (r *BuildConfigRegistry) UpdateBuildConfig(ctx apirequest.Context, config *api.BuildConfig) error {
	r.Lock()
	defer r.Unlock()
	r.BuildConfig = config
	return r.Err
}

func (r *BuildConfigRegistry) DeleteBuildConfig(ctx apirequest.Context, id string) error {
	r.Lock()
	defer r.Unlock()
	r.DeletedConfigID = id
	r.BuildConfig = nil
	return r.Err
}

func (r *BuildConfigRegistry) WatchBuildConfigs(ctx apirequest.Context, options *metainternal.ListOptions) (watch.Interface, error) {
	return nil, r.Err
}
