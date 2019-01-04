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
	"github.com/ef-ds/deque"
	gammazero "github.com/gammazero/deque"
	juju "github.com/juju/utils/deque"
	phf "github.com/phf/go-queue/queue"
	cookiejar "gopkg.in/karalabe/cookiejar.v2/collections/deque"
)

func BenchmarkSlowIncreaseListQueue(b *testing.B) {
	var l *list.List
	tests.SlowIncrease(
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

func BenchmarkSlowIncreaseListStack(b *testing.B) {
	var l *list.List
	tests.SlowIncrease(
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

func BenchmarkSlowIncreaseSliceQueue(b *testing.B) {
	var q *CustomSliceQueue
	tests.SlowIncrease(
		b,
		func() {
			q = NewCustomSliceQueue()
		},
		func(v interface{}) {
			q.PushBack(v.(*benchmark.TestValue))
		},
		func() (interface{}, bool) {
			return q.PopFront()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkSlowIncreaseSliceStack(b *testing.B) {
	var q *CustomSliceQueue
	tests.SlowIncrease(
		b,
		func() {
			q = NewCustomSliceQueue()
		},
		func(v interface{}) {
			q.PushBack(v.(*benchmark.TestValue))
		},
		func() (interface{}, bool) {
			return q.PopBack()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkSlowIncreaseGammazeroQueue(b *testing.B) {
	var q *gammazero.Deque
	tests.SlowIncrease(
		b,
		func() {
			q = new(gammazero.Deque)
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

func BenchmarkSlowIncreaseGammazeroStack(b *testing.B) {
	var q *gammazero.Deque
	tests.SlowIncrease(
		b,
		func() {
			q = new(gammazero.Deque)
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

func BenchmarkSlowIncreasePhfQueue(b *testing.B) {
	var q *phf.Queue
	tests.SlowIncrease(
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

func BenchmarkSlowIncreasePhfStack(b *testing.B) {
	var q *phf.Queue
	tests.SlowIncrease(
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

func BenchmarkSlowIncreaseCookiejarQueue(b *testing.B) {
	var q *cookiejar.Deque
	tests.SlowIncrease(
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

func BenchmarkSlowIncreaseCookiejarStack(b *testing.B) {
	var q *cookiejar.Deque
	tests.SlowIncrease(
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

func BenchmarkSlowIncreaseJujuQueue(b *testing.B) {
	var q *juju.Deque
	tests.SlowIncrease(
		b,
		func() {
			q = juju.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopFront()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkSlowIncreaseJujuStack(b *testing.B) {
	var q *juju.Deque
	tests.SlowIncrease(
		b,
		func() {
			q = juju.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopBack()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkSlowIncreaseImpl7Queue(b *testing.B) {
	var q *queueimpl7.Queueimpl7
	tests.SlowIncrease(
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

func BenchmarkSlowIncreaseDequeQueue(b *testing.B) {
	var q *deque.Deque
	tests.SlowIncrease(
		b,
		func() {
			q = deque.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopFront()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}

func BenchmarkSlowIncreaseDequeStack(b *testing.B) {
	var q *deque.Deque
	tests.SlowIncrease(
		b,
		func() {
			q = deque.New()
		},
		func(v interface{}) {
			q.PushBack(v)
		},
		func() (interface{}, bool) {
			return q.PopBack()
		},
		func() bool {
			return q.Len() == 0
		},
	)
}
