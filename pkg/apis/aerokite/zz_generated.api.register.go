/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package aerokite

import (
	"fmt"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	InternalStack = builders.NewInternalResource(
		"stacks",
		"Stack",
		func() runtime.Object { return &Stack{} },
		func() runtime.Object { return &StackList{} },
	)
	InternalStackStatus = builders.NewInternalResourceStatus(
		"stacks",
		"StackStatus",
		func() runtime.Object { return &Stack{} },
		func() runtime.Object { return &StackList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("aerokite.me").WithKinds(
		InternalStack,
		InternalStackStatus,
	)

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Stack struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   StackSpec
	Status StackStatus
}

type StackSpec struct {
	A string
}

type StackStatus struct {
}

//
// Stack Functions and Structs
//
// +k8s:deepcopy-gen=false
type StackStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type StackStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type StackList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Stack
}

func (Stack) NewStatus() interface{} {
	return StackStatus{}
}

func (pc *Stack) GetStatus() interface{} {
	return pc.Status
}

func (pc *Stack) SetStatus(s interface{}) {
	pc.Status = s.(StackStatus)
}

func (pc *Stack) GetSpec() interface{} {
	return pc.Spec
}

func (pc *Stack) SetSpec(s interface{}) {
	pc.Spec = s.(StackSpec)
}

func (pc *Stack) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *Stack) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc Stack) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store Stack.
// +k8s:deepcopy-gen=false
type StackRegistry interface {
	ListStacks(ctx request.Context, options *internalversion.ListOptions) (*StackList, error)
	GetStack(ctx request.Context, id string, options *metav1.GetOptions) (*Stack, error)
	CreateStack(ctx request.Context, id *Stack) (*Stack, error)
	UpdateStack(ctx request.Context, id *Stack) (*Stack, error)
	DeleteStack(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewStackRegistry(sp builders.StandardStorageProvider) StackRegistry {
	return &storageStack{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageStack struct {
	builders.StandardStorageProvider
}

func (s *storageStack) ListStacks(ctx request.Context, options *internalversion.ListOptions) (*StackList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*StackList), err
}

func (s *storageStack) GetStack(ctx request.Context, id string, options *metav1.GetOptions) (*Stack, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*Stack), nil
}

func (s *storageStack) CreateStack(ctx request.Context, object *Stack) (*Stack, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, true)
	if err != nil {
		return nil, err
	}
	return obj.(*Stack), nil
}

func (s *storageStack) UpdateStack(ctx request.Context, object *Stack) (*Stack, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*Stack), nil
}

func (s *storageStack) DeleteStack(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}
