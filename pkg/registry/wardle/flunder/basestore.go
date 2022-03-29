package flunder

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/storage"
)

var _ storage.Interface = &baseStorage{}

type baseStorage struct {
	gr schema.GroupResource
}

// NewBaseStorage returns an implementation of storage.Interface that has the
// empty methods. These methods should be overridden for specific
// implementations.
func NewBaseStorage(gr schema.GroupResource) storage.Interface {
	return &baseStorage{gr: gr}
}

func (s *baseStorage) methodNotSupported(action string) error {
	return errors.NewMethodNotSupported(s.gr, action)
}

func (s *baseStorage) Versioner() storage.Versioner {
	return nil
}

func (s *baseStorage) Create(context.Context, string, runtime.Object, runtime.Object, uint64) error {
	return s.methodNotSupported("create")
}

func (s *baseStorage) Delete(context.Context, string, runtime.Object, *storage.Preconditions, storage.ValidateObjectFunc, runtime.Object) error {
	return s.methodNotSupported("delete")
}

func (s *baseStorage) Watch(context.Context, string, storage.ListOptions) (watch.Interface, error) {
	return watch.NewEmptyWatch(), nil
}
func (s *baseStorage) WatchList(context.Context, string, storage.ListOptions) (watch.Interface, error) {
	return watch.NewEmptyWatch(), nil
}

func (s *baseStorage) Get(context.Context, string, storage.GetOptions, runtime.Object) error {
	return s.methodNotSupported("get")
}

func (s *baseStorage) GetToList(context.Context, string, storage.ListOptions, runtime.Object) error {
	return s.methodNotSupported("getToList")
}

func (s *baseStorage) List(context.Context, string, storage.ListOptions, runtime.Object) error {
	return s.methodNotSupported("list")
}

func (s *baseStorage) GuaranteedUpdate(context.Context, string, runtime.Object, bool, *storage.Preconditions, storage.UpdateFunc, runtime.Object) error {
	return s.methodNotSupported("update")
}

func (s *baseStorage) Count(key string) (int64, error) {
	return 0, nil
}

func (s *baseStorage) GetList(context.Context, string, storage.ListOptions, runtime.Object) error {
	return s.methodNotSupported("getToList")
}
