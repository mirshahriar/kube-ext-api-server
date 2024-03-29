// +build !ignore_autogenerated

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	aerokite "github.com/aerokite/kube-ext-api-server/pkg/apis/aerokite"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Stack_To_aerokite_Stack,
		Convert_aerokite_Stack_To_v1alpha1_Stack,
		Convert_v1alpha1_StackList_To_aerokite_StackList,
		Convert_aerokite_StackList_To_v1alpha1_StackList,
		Convert_v1alpha1_StackSpec_To_aerokite_StackSpec,
		Convert_aerokite_StackSpec_To_v1alpha1_StackSpec,
		Convert_v1alpha1_StackStatus_To_aerokite_StackStatus,
		Convert_aerokite_StackStatus_To_v1alpha1_StackStatus,
		Convert_v1alpha1_StackStatusStrategy_To_aerokite_StackStatusStrategy,
		Convert_aerokite_StackStatusStrategy_To_v1alpha1_StackStatusStrategy,
		Convert_v1alpha1_StackStrategy_To_aerokite_StackStrategy,
		Convert_aerokite_StackStrategy_To_v1alpha1_StackStrategy,
	)
}

func autoConvert_v1alpha1_Stack_To_aerokite_Stack(in *Stack, out *aerokite.Stack, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_StackSpec_To_aerokite_StackSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_StackStatus_To_aerokite_StackStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Stack_To_aerokite_Stack is an autogenerated conversion function.
func Convert_v1alpha1_Stack_To_aerokite_Stack(in *Stack, out *aerokite.Stack, s conversion.Scope) error {
	return autoConvert_v1alpha1_Stack_To_aerokite_Stack(in, out, s)
}

func autoConvert_aerokite_Stack_To_v1alpha1_Stack(in *aerokite.Stack, out *Stack, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_aerokite_StackSpec_To_v1alpha1_StackSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_aerokite_StackStatus_To_v1alpha1_StackStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_aerokite_Stack_To_v1alpha1_Stack is an autogenerated conversion function.
func Convert_aerokite_Stack_To_v1alpha1_Stack(in *aerokite.Stack, out *Stack, s conversion.Scope) error {
	return autoConvert_aerokite_Stack_To_v1alpha1_Stack(in, out, s)
}

func autoConvert_v1alpha1_StackList_To_aerokite_StackList(in *StackList, out *aerokite.StackList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]aerokite.Stack)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_StackList_To_aerokite_StackList is an autogenerated conversion function.
func Convert_v1alpha1_StackList_To_aerokite_StackList(in *StackList, out *aerokite.StackList, s conversion.Scope) error {
	return autoConvert_v1alpha1_StackList_To_aerokite_StackList(in, out, s)
}

func autoConvert_aerokite_StackList_To_v1alpha1_StackList(in *aerokite.StackList, out *StackList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Stack)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_aerokite_StackList_To_v1alpha1_StackList is an autogenerated conversion function.
func Convert_aerokite_StackList_To_v1alpha1_StackList(in *aerokite.StackList, out *StackList, s conversion.Scope) error {
	return autoConvert_aerokite_StackList_To_v1alpha1_StackList(in, out, s)
}

func autoConvert_v1alpha1_StackSpec_To_aerokite_StackSpec(in *StackSpec, out *aerokite.StackSpec, s conversion.Scope) error {
	out.A = in.A
	return nil
}

// Convert_v1alpha1_StackSpec_To_aerokite_StackSpec is an autogenerated conversion function.
func Convert_v1alpha1_StackSpec_To_aerokite_StackSpec(in *StackSpec, out *aerokite.StackSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_StackSpec_To_aerokite_StackSpec(in, out, s)
}

func autoConvert_aerokite_StackSpec_To_v1alpha1_StackSpec(in *aerokite.StackSpec, out *StackSpec, s conversion.Scope) error {
	out.A = in.A
	return nil
}

// Convert_aerokite_StackSpec_To_v1alpha1_StackSpec is an autogenerated conversion function.
func Convert_aerokite_StackSpec_To_v1alpha1_StackSpec(in *aerokite.StackSpec, out *StackSpec, s conversion.Scope) error {
	return autoConvert_aerokite_StackSpec_To_v1alpha1_StackSpec(in, out, s)
}

func autoConvert_v1alpha1_StackStatus_To_aerokite_StackStatus(in *StackStatus, out *aerokite.StackStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_StackStatus_To_aerokite_StackStatus is an autogenerated conversion function.
func Convert_v1alpha1_StackStatus_To_aerokite_StackStatus(in *StackStatus, out *aerokite.StackStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_StackStatus_To_aerokite_StackStatus(in, out, s)
}

func autoConvert_aerokite_StackStatus_To_v1alpha1_StackStatus(in *aerokite.StackStatus, out *StackStatus, s conversion.Scope) error {
	return nil
}

// Convert_aerokite_StackStatus_To_v1alpha1_StackStatus is an autogenerated conversion function.
func Convert_aerokite_StackStatus_To_v1alpha1_StackStatus(in *aerokite.StackStatus, out *StackStatus, s conversion.Scope) error {
	return autoConvert_aerokite_StackStatus_To_v1alpha1_StackStatus(in, out, s)
}

func autoConvert_v1alpha1_StackStatusStrategy_To_aerokite_StackStatusStrategy(in *StackStatusStrategy, out *aerokite.StackStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_StackStatusStrategy_To_aerokite_StackStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_StackStatusStrategy_To_aerokite_StackStatusStrategy(in *StackStatusStrategy, out *aerokite.StackStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_StackStatusStrategy_To_aerokite_StackStatusStrategy(in, out, s)
}

func autoConvert_aerokite_StackStatusStrategy_To_v1alpha1_StackStatusStrategy(in *aerokite.StackStatusStrategy, out *StackStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_aerokite_StackStatusStrategy_To_v1alpha1_StackStatusStrategy is an autogenerated conversion function.
func Convert_aerokite_StackStatusStrategy_To_v1alpha1_StackStatusStrategy(in *aerokite.StackStatusStrategy, out *StackStatusStrategy, s conversion.Scope) error {
	return autoConvert_aerokite_StackStatusStrategy_To_v1alpha1_StackStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_StackStrategy_To_aerokite_StackStrategy(in *StackStrategy, out *aerokite.StackStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_StackStrategy_To_aerokite_StackStrategy is an autogenerated conversion function.
func Convert_v1alpha1_StackStrategy_To_aerokite_StackStrategy(in *StackStrategy, out *aerokite.StackStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_StackStrategy_To_aerokite_StackStrategy(in, out, s)
}

func autoConvert_aerokite_StackStrategy_To_v1alpha1_StackStrategy(in *aerokite.StackStrategy, out *StackStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_aerokite_StackStrategy_To_v1alpha1_StackStrategy is an autogenerated conversion function.
func Convert_aerokite_StackStrategy_To_v1alpha1_StackStrategy(in *aerokite.StackStrategy, out *StackStrategy, s conversion.Scope) error {
	return autoConvert_aerokite_StackStrategy_To_v1alpha1_StackStrategy(in, out, s)
}
