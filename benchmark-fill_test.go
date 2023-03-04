// Copyright (c) 2018 ef-ds
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package deque_test

import (
	"container/list"
	"testing"

	"github.com/christianrpetrin/queue-tests/queueimpl7"
	"github.com/ef-ds/benchmark"
	deque "github.com/ef-ds/deque/v2"
	gammazero "github.com/gammazero/deque"
	gostl "github.com/liyue201/gostl/ds/deque"
	phf "github.com/phf/go-queue/queue"
	cookiejar "gopkg.in/karalabe/cookiejar.v2/collections/deque"
)

// tests variable is of type benchmark.Tests (https://github.com/ef-ds/benchmark/blob/master/tests.go)
// and is declared in the testdata.go source file.

func BenchmarkFillListQueue(b *testing.B) {
	var l *list.List
	tests.Fill(
		b,
		func() {
			l = list.New()
		},
		func(v interface{}) {
			l.PushBack(v)
		},
		func() (interface{}, bool) {
			return l.Remove(l.Front()), true
		},
		func() bool {
			return l.Front() == nil
		},
	)
}

func BenchmarkFillListStack(b *testing.B) {
	var l *list.List
	tests.Fill(
		b,
		func() {
			l = list.New()
		},
		func(v interface{}) {
			l.PushBack(v)
		},
		func() (interface{}, bool) {
			return l.Remove(l.Back()), true
		},
		func() bool {
			return l.Front() == nil
		},
	)
}

func BenchmarkFillSliceQueue(b *testing.B) {
	var q *CustomSliceQueue
	tests.FillTestObject(
		b,
		func() {
			q = NewCustomSliceQueue()
		},
		func(v *benchmark.TestValue) {
			q.PushBack(v)
		},
		func() (*benchmark.TestValue, bool) {
			return q.PopFront()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkFillSliceStack(b *testing.B) {
	var q *CustomSliceQueue
	tests.FillTestObject(
		b,
		func() {
			q = NewCustomSliceQueue()
		},
		func(v *benchmark.TestValue) {
			q.PushBack(v)
		},
		func() (*benchmark.TestValue, bool) {
			return q.PopBack()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkFillGammazeroQueue(b *testing.B) {
	var q *gammazero.Deque[*benchmark.TestValue]
	tests.FillTestObject(
		b,
		func() {
			q = new(gammazero.Deque[*benchmark.TestValue])
		},
		func(v *benchmark.TestValue) {
			q.PushBack(v)
		},
		func() (*benchmark.TestValue, bool) {
			return q.PopFront(), true
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkFillGammazeroStack(b *testing.B) {
	var q *gammazero.Deque[*benchmark.TestValue]
	tests.FillTestObject(
		b,
		func() {
			q = new(gammazero.Deque[*benchmark.TestValue])
		},
		func(v *benchmark.TestValue) {
			q.PushBack(v)
		},
		func() (*benchmark.TestValue, bool) {
			return q.PopBack(), true
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkFillPhfQueue(b *testing.B) {
	var q *phf.Queue
	tests.Fill(
		b,
		func() {
			q = phf.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopFront(), true
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkFillPhfStack(b *testing.B) {
	var q *phf.Queue
	tests.Fill(
		b,
		func() {
			q = phf.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopBack(), true
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkFillCookiejarQueue(b *testing.B) {
	var q *cookiejar.Deque
	tests.Fill(
		b,
		func() {
			q = cookiejar.New()
		},
		func(v interface{}) {
			q.PushRight(v)
		},
		func() (interface{}, bool) {
			return q.PopLeft(), true
		},
		func() bool {
			return q.Size() == 0
		},
	)
}

func BenchmarkFillCookiejarStack(b *testing.B) {
	var q *cookiejar.Deque
	tests.Fill(
		b,
		func() {
			q = cookiejar.New()
		},
		func(v interface{}) {
			q.PushRight(v)
		},
		func() (interface{}, bool) {
			return q.PopRight(), true
		},
		func() bool {
			return q.Size() == 0
		},
	)
}

func BenchmarkFillGostlQueue(b *testing.B) {
	var q *gostl.Deque[*benchmark.TestValue]
	tests.FillTestObject(
		b,
		func() {
			q = gostl.New[*benchmark.TestValue]()
		},
		func(v *benchmark.TestValue) {
			q.PushBack(v)
		},
		func() (*benchmark.TestValue, bool) {
			return q.PopFront(), true
		},
		func() bool {
			return q.Empty()
		},
	)
}

func BenchmarkFillGostlStack(b *testing.B) {
	q := gostl.New[*benchmark.TestValue]()
	tests.FillTestObject(
		b,
		func() {
			q = gostl.New[*benchmark.TestValue]()
		},
		func(v *benchmark.TestValue) {
			q.PushBack(v)
		},
		func() (*benchmark.TestValue, bool) {
			return q.PopBack(), true
		},
		func() bool {
			return q.Empty()
		},
	)
}

func BenchmarkFillImpl7Queue(b *testing.B) {
	var q *queueimpl7.Queueimpl7
	tests.Fill(
		b,
		func() {
			q = queueimpl7.New()
		},
		func(v interface{}) {
			q.Push(v)
		},
		func() (interface{}, bool) {
			return q.Pop()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkFillDequeQueue(b *testing.B) {
	var q *deque.Deque[*benchmark.TestValue]
	tests.FillTestObject(
		b,
		func() {
			q = deque.New[*benchmark.TestValue]()
		},
		func(v *benchmark.TestValue) {
			q.PushBack(v)
		},
		func() (*benchmark.TestValue, bool) {
			return q.PopFront()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkFillDequeStack(b *testing.B) {
	var q *deque.Deque[*benchmark.TestValue]
	tests.FillTestObject(
		b,
		func() {
			q = deque.New[*benchmark.TestValue]()
		},
		func(v *benchmark.TestValue) {
			q.PushBack(v)
		},
		func() (*benchmark.TestValue, bool) {
			return q.PopBack()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}
