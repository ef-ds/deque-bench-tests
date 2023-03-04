# Benchmark Tests

## Tests
The benchmark tests are composed of all the tests implemented in the [benchmark package](https://github.com/ef-ds/benchmark) and test sets.


### Test Sets
The two test sets, queue and stack, were designed to test the queues main usage patterns, being used as a FIFO queue and as a LIFO stack.

- Queue: test the data structures using them as a FIFO queue. PushBack() is used to add items to the queue; PopFront() to remove
- Stack: test the data structures using them as a LIFO stack. PushBack() is used to add items to the queue; PopBack() to remove


## Tested Queues

Besides deque, the tests also probe a few high quality open source queue implementations as well as a experimental queue implementation, impl7, alongside the standard list package as well as using simple slice as a deque.

- List based queue: uses the standard [list](https://github.com/golang/go/tree/master/src/container/list) package as a FIFO queue as well as a LIFO stack.
- [CustomSliceQueue](testdata_test.go): uses a simple, dynamically growing slice as its underlying data structure.
- [impl7](https://github.com/christianrpetrin/queue-tests/tree/master/queueimpl7/queueimpl7.go): experimental queue implementation that stores the values in linked slices. This implementation tests the queue performance when performing lazy creation of the internal slice as well as starting with a 1-sized slice, allowing it to grow up to 16 by using the built in append function. Subsequent slices are created with 128 fixed size.
- [phf](https://github.com/phf/go-queue): slice, ring based queue implementation. Interesting to note the author did a pretty good job researching and probing other queue implementations as well.
- [gammazero](https://github.com/gammazero/deque): the deque implemented in this package is also a slice, ring based queue implementation. Supports generics.
- [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go): the deque implemented in this package uses a circular slice of blocks to store the elements. Interesting to note the queue uses a block size of 4096, suggesting it is optimized for large data sets.
[search for "deque"](https://godoc.org/?q=deque) on GoDoc.org shows this deque is by far the most imported of all deques there.
- [gostl](https://github.com/liyue201/gostl/blob/master/ds/deque/deque.go): Deque supports efficient data insertion from the head and tail, random access and iterator access. Supports generics.

We're actively looking for other, high quality queues to add to our tests. Due to the large volume of open source queues available, it is not possible to add all of them to the tests. However, all the new tested ones we're adding to this [issue](https://github.com/ef-ds/deque/issues/10).


### Efficient Data Structures deque vs stack

Efficient Data Structures implements this deque package as well as the [stack](https://github.com/ef-ds/stack) package which can also be used as a LIFO stack.

The stack package is a simplified version of this deque package. When it comes to using the packages as a LIFO stack, the main differences are:

1) Stack is a simpler version of deque that performs better than deque on most, if not all, LIFO stack tests
2) Differently from deque which attempts to reuse slices, stack doesn't implement such logic making it faster, especially for small data sets, but also causes more allocations for larger data sets (but with similar performance)



## Results

The raw results of a local run are stored under the [testdata](testdata) directory.

Refer [here](PERFORMANCE.md) for curated results.


## How To Run

From the package main directory, the tests can be run with below command.

```
go test -benchmem -timeout 60m -bench=. -run=^$
```

To run the test for a single queue, below command can be used.

```
go test -benchmem -timeout 60m -bench="QUEUE_NAME*" -run=^$
```

Replace the QUEUE_NAME with the desired queue such as "List", "Slice", "Gammazero", "Phf", "Cookiejar", "Impl7", "Gostl", "Deque".


To run only a specific test suite, below command can be used.

```
go test -benchmem -timeout 60m -bench="TEST_SUITE_NAME*" -run=^$
```

Replace the TEST_SUITE_NAME with the desired test suite such as "Microservice", "Fill", "Refill", "RefillFull", "SlowIncrease", "SlowDecrease", "Stable".


## Test Variations

It is common to see significant variations in the test numbers with different test runs due to different reasons such as processes running in the hosting computer while the tests are running.

It is recommended to run the test multiple times and compare the aggregated results in order to help eliminate/smooth the test variations.

To run the tests multiple times, use the "go test" count parameter as below.

```
go test -benchmem -count 10 -timeout 600m -bench=. -run=^$
```

As the number of tests and now, test runs as well, is very large, it becomes very difficult to analyze and understand the results. In order to be able to analyze and compare the results between the different queues, the [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) tool can be used to aggregate the test results. But as benchstat was designed to compare the same set of tests, it is necessary to first split all the different tests into separate test files renaming each
test with the same name, so benchstat will be able to match the different tests.

First step is to run the test and output the results in a file. Below command can be used to run all tests 10 times.

```
go test -benchmem -count 10 -timeout 600m -bench=. -run=^$ > testdata/results.txt
```

Next step is to split the "results.txt" file into separate test files. The [test-splitter](https://github.com/ef-ds/tools/tree/master/testsplitter) tool can be used for this purpose. To run the tool, clone the repo and run test-splitter from the "testsplitter" directory as follow.

```
go run *.go --file PATH_TO_RESULTS.TXT
```

Test-splitter should output one file per test name in the tests results file. The file names are named after each test name.

The last step is to run the [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) tool to aggregate and compare the results.

Below are the set of benchstat commands that can be used to compare deque against the other tested queues.

Deque vs impl7
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev2.0.0.txt testdata/BenchmarkMicroserviceImpl7Queue.txt
benchstat testdata/BenchmarkFillDequeQueuev2.0.0.txt testdata/BenchmarkFillImpl7Queue.txt
benchstat testdata/BenchmarkRefillDequeQueuev2.0.0.txt testdata/BenchmarkRefillImpl7Queue.txt
benchstat testdata/BenchmarkRefillFullDequeQueuev2.0.0.txt testdata/BenchmarkRefillFullImpl7Queue.txt
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowIncreaseImpl7Queue.txt
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowDecreaseImpl7Queue.txt
benchstat testdata/BenchmarkStableDequeQueuev2.0.0.txt testdata/BenchmarkStableImpl7Queue.txt
```

Deque vs list
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev2.0.0.txt testdata/BenchmarkMicroserviceListQueue.txt
benchstat testdata/BenchmarkMicroserviceDequeStackv2.0.0.txt testdata/BenchmarkMicroserviceListStack.txt
benchstat testdata/BenchmarkFillDequeQueuev2.0.0.txt testdata/BenchmarkFillListQueue.txt
benchstat testdata/BenchmarkFillDequeStackv2.0.0.txt testdata/BenchmarkFillListStack.txt
benchstat testdata/BenchmarkRefillDequeQueuev2.0.0.txt testdata/BenchmarkRefillListQueue.txt
benchstat testdata/BenchmarkRefillDequeStackv2.0.0.txt testdata/BenchmarkRefillListStack.txt
benchstat testdata/BenchmarkRefillFullDequeQueuev2.0.0.txt testdata/BenchmarkRefillFullListQueue.txt
benchstat testdata/BenchmarkRefillFullDequeStackv2.0.0.txt testdata/BenchmarkRefillFullListStack.txt
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowIncreaseListQueue.txt
benchstat testdata/BenchmarkSlowIncreaseDequeStackv2.0.0.txt testdata/BenchmarkSlowIncreaseListStack.txt
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowDecreaseListQueue.txt
benchstat testdata/BenchmarkSlowDecreaseDequeStackv2.0.0.txt testdata/BenchmarkSlowDecreaseListStack.txt
benchstat testdata/BenchmarkStableDequeQueuev2.0.0.txt testdata/BenchmarkStableListQueue.txt
benchstat testdata/BenchmarkStableDequeStackv2.0.0.txt testdata/BenchmarkStableListStack.txt
```

Deque vs slice
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev2.0.0.txt testdata/BenchmarkMicroserviceSliceQueue.txt
benchstat testdata/BenchmarkMicroserviceDequeStackv2.0.0.txt testdata/BenchmarkMicroserviceSliceStack.txt
benchstat testdata/BenchmarkFillDequeQueuev2.0.0.txt testdata/BenchmarkFillSliceQueue.txt
benchstat testdata/BenchmarkFillDequeStackv2.0.0.txt testdata/BenchmarkFillSliceStack.txt
benchstat testdata/BenchmarkRefillDequeQueuev2.0.0.txt testdata/BenchmarkRefillSliceQueue.txt
benchstat testdata/BenchmarkRefillDequeStackv2.0.0.txt testdata/BenchmarkRefillSliceStack.txt
benchstat testdata/BenchmarkRefillFullDequeQueuev2.0.0.txt testdata/BenchmarkRefillFullSliceQueue.txt
benchstat testdata/BenchmarkRefillFullDequeStackv2.0.0.txt testdata/BenchmarkRefillFullSliceStack.txt
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowIncreaseSliceQueue.txt
benchstat testdata/BenchmarkSlowIncreaseDequeStackv2.0.0.txt testdata/BenchmarkSlowIncreaseSliceStack.txt
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowDecreaseSliceQueue.txt
benchstat testdata/BenchmarkSlowDecreaseDequeStackv2.0.0.txt testdata/BenchmarkSlowDecreaseSliceStack.txt
benchstat testdata/BenchmarkStableDequeQueuev2.0.0.txt testdata/BenchmarkStableSliceQueue.txt
benchstat testdata/BenchmarkStableDequeStackv2.0.0.txt testdata/BenchmarkStableSliceStack.txt
```

Deque vs phf
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev2.0.0.txt testdata/BenchmarkMicroservicePhfQueue.txt
benchstat testdata/BenchmarkMicroserviceDequeStackv2.0.0.txt testdata/BenchmarkMicroservicePhfStack.txt
benchstat testdata/BenchmarkFillDequeQueuev2.0.0.txt testdata/BenchmarkFillPhfQueue.txt
benchstat testdata/BenchmarkFillDequeStackv2.0.0.txt testdata/BenchmarkFillPhfStack.txt
benchstat testdata/BenchmarkRefillDequeQueuev2.0.0.txt testdata/BenchmarkRefillPhfQueue.txt
benchstat testdata/BenchmarkRefillDequeStackv2.0.0.txt testdata/BenchmarkRefillPhfStack.txt
benchstat testdata/BenchmarkRefillFullDequeQueuev2.0.0.txt testdata/BenchmarkRefillFullPhfQueue.txt
benchstat testdata/BenchmarkRefillFullDequeStackv2.0.0.txt testdata/BenchmarkRefillFullPhfStack.txt
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowIncreasePhfQueue.txt
benchstat testdata/BenchmarkSlowIncreaseDequeStackv2.0.0.txt testdata/BenchmarkSlowIncreasePhfStack.txt
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowDecreasePhfQueue.txt
benchstat testdata/BenchmarkSlowDecreaseDequeStackv2.0.0.txt testdata/BenchmarkSlowDecreasePhfStack.txt
benchstat testdata/BenchmarkStableDequeQueuev2.0.0.txt testdata/BenchmarkStablePhfQueue.txt
benchstat testdata/BenchmarkStableDequeStackv2.0.0.txt testdata/BenchmarkStablePhfStack.txt
```

Deque vs gammazero
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev2.0.0.txt testdata/BenchmarkMicroserviceGammazeroQueue.txt
benchstat testdata/BenchmarkMicroserviceDequeStackv2.0.0.txt testdata/BenchmarkMicroserviceGammazeroStack.txt
benchstat testdata/BenchmarkFillDequeQueuev2.0.0.txt testdata/BenchmarkFillGammazeroQueue.txt
benchstat testdata/BenchmarkFillDequeStackv2.0.0.txt testdata/BenchmarkFillGammazeroStack.txt
benchstat testdata/BenchmarkRefillDequeQueuev2.0.0.txt testdata/BenchmarkRefillGammazeroQueue.txt
benchstat testdata/BenchmarkRefillDequeStackv2.0.0.txt testdata/BenchmarkRefillGammazeroStack.txt
benchstat testdata/BenchmarkRefillFullDequeQueuev2.0.0.txt testdata/BenchmarkRefillFullGammazeroQueue.txt
benchstat testdata/BenchmarkRefillFullDequeStackv2.0.0.txt testdata/BenchmarkRefillFullGammazeroStack.txt
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowIncreaseGammazeroQueue.txt
benchstat testdata/BenchmarkSlowIncreaseDequeStackv2.0.0.txt testdata/BenchmarkSlowIncreaseGammazeroStack.txt
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowDecreaseGammazeroQueue.txt
benchstat testdata/BenchmarkSlowDecreaseDequeStackv2.0.0.txt testdata/BenchmarkSlowDecreaseGammazeroStack.txt
benchstat testdata/BenchmarkStableDequeQueuev2.0.0.txt testdata/BenchmarkStableGammazeroQueue.txt
benchstat testdata/BenchmarkStableDequeStackv2.0.0.txt testdata/BenchmarkStableGammazeroStack.txt
```

Deque vs cookiejar
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev2.0.0.txt testdata/BenchmarkMicroserviceCookiejarQueue.txt
benchstat testdata/BenchmarkMicroserviceDequeStackv2.0.0.txt testdata/BenchmarkMicroserviceCookiejarStack.txt
benchstat testdata/BenchmarkFillDequeQueuev2.0.0.txt testdata/BenchmarkFillCookiejarQueue.txt
benchstat testdata/BenchmarkFillDequeStackv2.0.0.txt testdata/BenchmarkFillCookiejarStack.txt
benchstat testdata/BenchmarkRefillDequeQueuev2.0.0.txt testdata/BenchmarkRefillCookiejarQueue.txt
benchstat testdata/BenchmarkRefillDequeStackv2.0.0.txt testdata/BenchmarkRefillCookiejarStack.txt
benchstat testdata/BenchmarkRefillFullDequeQueuev2.0.0.txt testdata/BenchmarkRefillFullCookiejarQueue.txt
benchstat testdata/BenchmarkRefillFullDequeStackv2.0.0.txt testdata/BenchmarkRefillFullCookiejarStack.txt
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowIncreaseCookiejarQueue.txt
benchstat testdata/BenchmarkSlowIncreaseDequeStackv2.0.0.txt testdata/BenchmarkSlowIncreaseCookiejarStack.txt
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowDecreaseCookiejarQueue.txt
benchstat testdata/BenchmarkSlowDecreaseDequeStackv2.0.0.txt testdata/BenchmarkSlowDecreaseCookiejarStack.txt
benchstat testdata/BenchmarkStableDequeQueuev2.0.0.txt testdata/BenchmarkStableCookiejarQueue.txt
benchstat testdata/BenchmarkStableDequeStackv2.0.0.txt testdata/BenchmarkStableCookiejarStack.txt
```

Deque vs gostl
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev2.0.0.txt testdata/BenchmarkMicroserviceGostlQueue.txt
benchstat testdata/BenchmarkMicroserviceDequeStackv2.0.0.txt testdata/BenchmarkMicroserviceGostlStack.txt
benchstat testdata/BenchmarkFillDequeQueuev2.0.0.txt testdata/BenchmarkFillGostlQueue.txt
benchstat testdata/BenchmarkFillDequeStackv2.0.0.txt testdata/BenchmarkFillGostlStack.txt
benchstat testdata/BenchmarkRefillDequeQueuev2.0.0.txt testdata/BenchmarkRefillGostlQueue.txt
benchstat testdata/BenchmarkRefillDequeStackv2.0.0.txt testdata/BenchmarkRefillGostlStack.txt
benchstat testdata/BenchmarkRefillFullDequeQueuev2.0.0.txt testdata/BenchmarkRefillFullGostlQueue.txt
benchstat testdata/BenchmarkRefillFullDequeStackv2.0.0.txt testdata/BenchmarkRefillFullGostlStack.txt
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowIncreaseGostlQueue.txt
benchstat testdata/BenchmarkSlowIncreaseDequeStackv2.0.0.txt testdata/BenchmarkSlowIncreaseGostlStack.txt
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev2.0.0.txt testdata/BenchmarkSlowDecreaseGostlQueue.txt
benchstat testdata/BenchmarkSlowDecreaseDequeStackv2.0.0.txt testdata/BenchmarkSlowDecreaseGostlStack.txt
benchstat testdata/BenchmarkStableDequeQueuev2.0.0.txt testdata/BenchmarkStableGostlQueue.txt
benchstat testdata/BenchmarkStableDequeStackv2.0.0.txt testdata/BenchmarkStableGostlStack.txt
```
