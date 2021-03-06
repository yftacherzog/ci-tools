/*
Copyright The TestGrid Authors.

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

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test_status.proto

package test_status

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TestStatus int32

const (
	// Proto versions of test_status.py's GathererStatus
	// Note that: NO_RESULT is used to signal that there should be no change.
	// This must be updated every time a new GathererStatus is added.
	TestStatus_NO_RESULT         TestStatus = 0
	TestStatus_PASS              TestStatus = 1
	TestStatus_PASS_WITH_ERRORS  TestStatus = 2
	TestStatus_PASS_WITH_SKIPS   TestStatus = 3
	TestStatus_RUNNING           TestStatus = 4
	TestStatus_CATEGORIZED_ABORT TestStatus = 5
	TestStatus_UNKNOWN           TestStatus = 6
	TestStatus_CANCEL            TestStatus = 7
	TestStatus_BLOCKED           TestStatus = 8
	TestStatus_TIMED_OUT         TestStatus = 9
	TestStatus_CATEGORIZED_FAIL  TestStatus = 10
	TestStatus_BUILD_FAIL        TestStatus = 11
	TestStatus_FAIL              TestStatus = 12
	TestStatus_FLAKY             TestStatus = 13
	TestStatus_TOOL_FAIL         TestStatus = 14
	TestStatus_BUILD_PASSED      TestStatus = 15
)

var TestStatus_name = map[int32]string{
	0:  "NO_RESULT",
	1:  "PASS",
	2:  "PASS_WITH_ERRORS",
	3:  "PASS_WITH_SKIPS",
	4:  "RUNNING",
	5:  "CATEGORIZED_ABORT",
	6:  "UNKNOWN",
	7:  "CANCEL",
	8:  "BLOCKED",
	9:  "TIMED_OUT",
	10: "CATEGORIZED_FAIL",
	11: "BUILD_FAIL",
	12: "FAIL",
	13: "FLAKY",
	14: "TOOL_FAIL",
	15: "BUILD_PASSED",
}

var TestStatus_value = map[string]int32{
	"NO_RESULT":         0,
	"PASS":              1,
	"PASS_WITH_ERRORS":  2,
	"PASS_WITH_SKIPS":   3,
	"RUNNING":           4,
	"CATEGORIZED_ABORT": 5,
	"UNKNOWN":           6,
	"CANCEL":            7,
	"BLOCKED":           8,
	"TIMED_OUT":         9,
	"CATEGORIZED_FAIL":  10,
	"BUILD_FAIL":        11,
	"FAIL":              12,
	"FLAKY":             13,
	"TOOL_FAIL":         14,
	"BUILD_PASSED":      15,
}

func (x TestStatus) String() string {
	return proto.EnumName(TestStatus_name, int32(x))
}

func (TestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3f9a6ab41bff9dae, []int{0}
}

func init() {
	proto.RegisterEnum("TestStatus", TestStatus_name, TestStatus_value)
}

func init() { proto.RegisterFile("test_status.proto", fileDescriptor_3f9a6ab41bff9dae) }

var fileDescriptor_3f9a6ab41bff9dae = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4d, 0x4e, 0x03, 0x31,
	0x0c, 0x85, 0xa1, 0xb4, 0xd3, 0x8e, 0xfb, 0xe7, 0x1a, 0xb8, 0x04, 0x0b, 0x36, 0x9c, 0x20, 0x33,
	0x93, 0x96, 0x68, 0x42, 0x5c, 0xe5, 0x47, 0x15, 0x6c, 0x22, 0x90, 0xba, 0x2e, 0x62, 0xc2, 0x11,
	0xb8, 0x37, 0x4a, 0x8a, 0x44, 0x77, 0xcf, 0xcf, 0xcf, 0x4f, 0x9f, 0x0c, 0x9b, 0x74, 0x1c, 0x52,
	0x1c, 0xd2, 0x7b, 0xfa, 0x1e, 0x1e, 0x3f, 0xbf, 0x4e, 0xe9, 0xf4, 0xf0, 0x33, 0x02, 0xf0, 0xc7,
	0x21, 0xb9, 0x62, 0xd2, 0x12, 0x6a, 0xc3, 0xd1, 0x4a, 0x17, 0xb4, 0xc7, 0x2b, 0x9a, 0xc1, 0x78,
	0x2f, 0x9c, 0xc3, 0x6b, 0xba, 0x03, 0xcc, 0x2a, 0x1e, 0x94, 0x7f, 0x8e, 0xd2, 0x5a, 0xb6, 0x0e,
	0x47, 0x74, 0x0b, 0xeb, 0x7f, 0xd7, 0xf5, 0x6a, 0xef, 0xf0, 0x86, 0xe6, 0x30, 0xb5, 0xc1, 0x18,
	0x65, 0x76, 0x38, 0xa6, 0x7b, 0xd8, 0xb4, 0xc2, 0xcb, 0x1d, 0x5b, 0xf5, 0x26, 0xbb, 0x28, 0x1a,
	0xb6, 0x1e, 0x27, 0x39, 0x13, 0x4c, 0x6f, 0xf8, 0x60, 0xb0, 0x22, 0x80, 0xaa, 0x15, 0xa6, 0x95,
	0x1a, 0xa7, 0x79, 0xd1, 0x68, 0x6e, 0x7b, 0xd9, 0xe1, 0x2c, 0xd3, 0x78, 0xf5, 0x22, 0xbb, 0xc8,
	0xc1, 0x63, 0x9d, 0x19, 0x2e, 0xbb, 0xb6, 0x42, 0x69, 0x04, 0x5a, 0x01, 0x34, 0x41, 0xe9, 0xbf,
	0x79, 0x9e, 0x99, 0x8b, 0x5a, 0x50, 0x0d, 0x93, 0xad, 0x16, 0xfd, 0x2b, 0x2e, 0x4b, 0x13, 0xb3,
	0x3e, 0x67, 0x56, 0x84, 0xb0, 0x38, 0xdf, 0x64, 0x7a, 0xd9, 0xe1, 0xfa, 0xa3, 0x2a, 0xef, 0x78,
	0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xb6, 0x9b, 0x15, 0x23, 0x01, 0x00, 0x00,
}
