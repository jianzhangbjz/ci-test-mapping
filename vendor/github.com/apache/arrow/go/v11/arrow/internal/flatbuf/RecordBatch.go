// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// A data header describing the shared memory layout of a "record" or "row"
/// batch. Some systems call this a "row batch" internally and others a "record
/// batch".
type RecordBatch struct {
	_tab flatbuffers.Table
}

func GetRootAsRecordBatch(buf []byte, offset flatbuffers.UOffsetT) *RecordBatch {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RecordBatch{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *RecordBatch) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RecordBatch) Table() flatbuffers.Table {
	return rcv._tab
}

/// number of records / rows. The arrays in the batch should all have this
/// length
func (rcv *RecordBatch) Length() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// number of records / rows. The arrays in the batch should all have this
/// length
func (rcv *RecordBatch) MutateLength(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

/// Nodes correspond to the pre-ordered flattened logical schema
func (rcv *RecordBatch) Nodes(obj *FieldNode, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 16
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *RecordBatch) NodesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Nodes correspond to the pre-ordered flattened logical schema
/// Buffers correspond to the pre-ordered flattened buffer tree
///
/// The number of buffers appended to this list depends on the schema. For
/// example, most primitive arrays will have 2 buffers, 1 for the validity
/// bitmap and 1 for the values. For struct arrays, there will only be a
/// single buffer for the validity (nulls) bitmap
func (rcv *RecordBatch) Buffers(obj *Buffer, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 16
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *RecordBatch) BuffersLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Buffers correspond to the pre-ordered flattened buffer tree
///
/// The number of buffers appended to this list depends on the schema. For
/// example, most primitive arrays will have 2 buffers, 1 for the validity
/// bitmap and 1 for the values. For struct arrays, there will only be a
/// single buffer for the validity (nulls) bitmap
/// Optional compression of the message body
func (rcv *RecordBatch) Compression(obj *BodyCompression) *BodyCompression {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BodyCompression)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Optional compression of the message body
func RecordBatchStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func RecordBatchAddLength(builder *flatbuffers.Builder, length int64) {
	builder.PrependInt64Slot(0, length, 0)
}
func RecordBatchAddNodes(builder *flatbuffers.Builder, nodes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(nodes), 0)
}
func RecordBatchStartNodesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(16, numElems, 8)
}
func RecordBatchAddBuffers(builder *flatbuffers.Builder, buffers flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(buffers), 0)
}
func RecordBatchStartBuffersVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(16, numElems, 8)
}
func RecordBatchAddCompression(builder *flatbuffers.Builder, compression flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(compression), 0)
}
func RecordBatchEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}