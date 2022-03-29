package flunder

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/sample-apiserver/pkg/apis/wardle"
)

var _ storage.Interface = &dryRunStorage{}

type dryRunStorage struct {
	storage.Interface
	db map[string]*wardle.Flunder
}

func newDryRunStorage(gvr schema.GroupVersionResource) *dryRunStorage {
	ds := &dryRunStorage{
		Interface: NewBaseStorage(gvr.GroupResource()),
		db:        make(map[string]*wardle.Flunder),
	}
	// Mock some objects
	for i := 0; i < 100000; i++ {
		ds.db[fmt.Sprintf("/wardle.example.com/flunders/wardle/test-%d", i)] = mockObject(i)
	}
	return ds
}

func (s *dryRunStorage) Create(ctx context.Context, key string, obj, out runtime.Object, ttl uint64) error {
	fmt.Println("Received CREATE Req")
	flunder, ok := obj.(*wardle.Flunder)
	if !ok {
		return apierrors.NewBadRequest(fmt.Sprintf("not a Flunder: %#v", obj))
	}
	s.db[key] = flunder
	//out = flunder.DeepCopyObject()
	lout := out.(*wardle.Flunder)
	flunder.DeepCopyInto(lout)
	fmt.Printf("Created:: %#v\n", lout)
	return nil
}

func (s *dryRunStorage) Get(ctx context.Context, key string, getOptions storage.GetOptions, objPtr runtime.Object) error {
	fmt.Println("Received GET Req")
	obj, ok := s.db[key]
	if !ok {
		return apierrors.NewBadRequest(fmt.Sprintf("not found: %v", key))
	}
	//objPtr = obj.DeepCopyObject()
	lout := objPtr.(*wardle.Flunder)
	obj.DeepCopyInto(lout)

	return nil
}

func (s *dryRunStorage) List(ctx context.Context, key string, listOptions storage.ListOptions, listObj runtime.Object) error {
	fmt.Println("Received List Req")
	list := wardle.FlunderList{}
	for _, v := range s.db {
		list.Items = append(list.Items, *v)
	}
	listObj = list.DeepCopyObject()
	return nil
}

func (s *dryRunStorage) GetList(ctx context.Context, key string, opts storage.ListOptions, listObj runtime.Object) error {
	fmt.Println("Received GetToList Req")
	list := &wardle.FlunderList{}
	for _, v := range s.db {
		list.Items = append(list.Items, *v)
	}
	lout := listObj.(*wardle.FlunderList)
	list.DeepCopyInto(lout)
	return nil
}

func (s *dryRunStorage) Delete(ctx context.Context, key string, objPtr runtime.Object, preconditions *storage.Preconditions, _ storage.ValidateObjectFunc, _ runtime.Object) error {
	fmt.Println("Received delete Req.")
	delete(s.db, key)
	return nil
}

func mockObject(i int) *wardle.Flunder {
	return &wardle.Flunder{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Flunder",
			APIVersion: "wardle.example.com/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("test-%v", i),
		},
	}
}
