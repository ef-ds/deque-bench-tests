# Performance

Below compares the deque [benchmark tests](BENCHMARK_TESTS.md) results with the other tested queues.

## Running the Tests
In the "testdata" directory, we have included the result of local test runs for all queues. Below uses this run to compare the queues, but it's possible and we highly encourage you to run the tests yourself to help validate the results.

To run the tests locally, clone the deque repo, cd to the deque main directory and run below command.

```
go test -benchmem -timeout 60m -bench=. -run=^$
```

This command will run all tests for all queues locally once. This should be good enouh to give you a sense of the queues performance, but to do a proper comparison, elimating test variations, we recommend you to run the tests as detailed [here](BENCHMARK_TESTS.md) by running the tests with multiple counts, splitting the files with [test-splitter](https://github.com/ef-ds/tools/tree/master/testsplitter) and using the [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) tool to aggregate the results.


## Bottom Line
As a general purpose double-ended queue, deque is the queue that displays the most balanced performance, performing either very competitively or besting all other queues in all the different test scenarios.


## Results
Given the enormous amount of test data, it can be difficult and time consuming to find out the net impact of all the tests, so we generally spend most of the time on the results of the very simple fill tests, which sequentially add and remove N number of items, and the Microservice test, which is a composite test of all other tests.

Below results is for deque [v2.0.0](https://github.com/ef-ds/deque/blob/master/CHANGELOG.md).


### Fill Test Results
deque vs [impl7](https://github.com/christianrpetrin/queue-tests/tree/master/queueimpl7/queueimpl7.go) - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillImpl7Queue.txt
name        old time/op    new time/op    delta
/0-4          40.5ns ± 0%    41.1ns ± 5%     ~     (p=1.000 n=10+10)
/1-4           160ns ± 1%     142ns ± 0%  -11.39%  (p=0.000 n=10+9)
/10-4          553ns ± 2%     748ns ± 1%  +35.24%  (p=0.000 n=10+9)
/100-4        4.15µs ± 1%    4.07µs ± 1%   -1.92%  (p=0.000 n=10+10)
/1000-4       34.6µs ± 2%    35.4µs ± 0%   +2.51%  (p=0.000 n=9+9)
/10000-4       348µs ± 1%     358µs ± 1%   +2.82%  (p=0.000 n=8+9)
/100000-4     3.88ms ± 2%    3.64ms ± 0%   -6.23%  (p=0.000 n=9+9)
/1000000-4    42.7ms ± 2%    43.1ms ± 2%     ~     (p=0.143 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%  -25.00%  (p=0.000 n=10+10)
/1-4            160B ± 0%      112B ± 0%  -30.00%  (p=0.000 n=10+10)
/10-4           432B ± 0%      736B ± 0%  +70.37%  (p=0.000 n=10+10)
/100-4        4.48kB ± 0%    4.26kB ± 0%   -5.00%  (p=0.000 n=10+10)
/1000-4       25.2kB ± 0%    33.2kB ± 0%  +31.98%  (p=0.000 n=10+10)
/10000-4       243kB ± 0%     323kB ± 0%  +33.10%  (p=0.000 n=10+8)
/100000-4     2.42MB ± 0%    3.23MB ± 0%  +33.34%  (p=0.000 n=10+10)
/1000000-4    24.2MB ± 0%    32.3MB ± 0%  +33.34%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            4.00 ± 0%      4.00 ± 0%     ~     (all equal)
/10-4           14.0 ± 0%      17.0 ± 0%  +21.43%  (p=0.000 n=10+10)
/100-4           107 ± 0%       109 ± 0%   +1.87%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.02k ± 0%   +0.99%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.2k ± 0%   +0.79%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      102k ± 0%   +0.78%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.02M ± 0%   +0.78%  (p=0.000 n=10+10)
```

deque vs [list](https://github.com/golang/go/tree/master/src/container/list) - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillListQueue.txt
name        old time/op    new time/op    delta
/0-4          40.5ns ± 0%    44.2ns ±23%    +9.20%  (p=0.000 n=10+10)
/1-4           160ns ± 1%     112ns ± 1%   -30.14%  (p=0.000 n=10+10)
/10-4          553ns ± 2%     736ns ± 0%   +33.08%  (p=0.000 n=10+8)
/100-4        4.15µs ± 1%    6.94µs ± 0%   +67.27%  (p=0.000 n=10+10)
/1000-4       34.6µs ± 2%    68.4µs ± 1%   +97.89%  (p=0.000 n=9+9)
/10000-4       348µs ± 1%     700µs ± 3%  +100.75%  (p=0.000 n=8+9)
/100000-4     3.88ms ± 2%   14.41ms ± 5%  +271.05%  (p=0.000 n=9+10)
/1000000-4    42.7ms ± 2%   128.4ms ±14%  +200.94%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            160B ± 0%      112B ± 0%   -30.00%  (p=0.000 n=10+10)
/10-4           432B ± 0%      688B ± 0%   +59.26%  (p=0.000 n=10+10)
/100-4        4.48kB ± 0%    6.45kB ± 0%   +43.93%  (p=0.000 n=10+10)
/1000-4       25.2kB ± 0%    64.0kB ± 0%  +154.48%  (p=0.000 n=10+10)
/10000-4       243kB ± 0%     640kB ± 0%  +163.91%  (p=0.000 n=10+6)
/100000-4     2.42MB ± 0%    6.40MB ± 0%  +164.43%  (p=0.000 n=10+9)
/1000000-4    24.2MB ± 0%    64.0MB ± 0%  +164.60%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      21.0 ± 0%   +50.00%  (p=0.000 n=10+10)
/100-4           107 ± 0%       201 ± 0%   +87.85%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     2.00k ± 0%   +97.53%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     20.0k ± 0%   +98.36%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      200k ± 0%   +98.44%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     2.00M ± 0%   +98.45%  (p=0.000 n=10+10)
```

deque vs [list](https://github.com/golang/go/tree/master/src/container/list) - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillListStack.txt
name        old time/op    new time/op    delta
/0-4          41.0ns ± 4%    41.2ns ± 1%    +0.47%  (p=0.048 n=9+9)
/1-4           160ns ± 1%     113ns ± 1%   -29.62%  (p=0.000 n=10+8)
/10-4          564ns ± 0%     741ns ± 1%   +31.50%  (p=0.000 n=10+8)
/100-4        4.23µs ± 1%    6.99µs ± 1%   +65.37%  (p=0.000 n=10+10)
/1000-4       34.9µs ± 2%    68.8µs ± 1%   +97.27%  (p=0.000 n=9+9)
/10000-4       347µs ± 1%     694µs ± 1%  +100.11%  (p=0.000 n=9+9)
/100000-4     3.85ms ± 1%   14.20ms ± 6%  +268.45%  (p=0.000 n=10+10)
/1000000-4    42.1ms ± 1%   118.1ms ±11%  +180.59%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            160B ± 0%      112B ± 0%   -30.00%  (p=0.000 n=10+10)
/10-4           432B ± 0%      688B ± 0%   +59.26%  (p=0.000 n=10+10)
/100-4        4.48kB ± 0%    6.45kB ± 0%   +43.93%  (p=0.000 n=10+10)
/1000-4       25.2kB ± 0%    64.0kB ± 0%  +154.48%  (p=0.000 n=10+10)
/10000-4       243kB ± 0%     640kB ± 0%  +163.91%  (p=0.000 n=10+9)
/100000-4     2.42MB ± 0%    6.40MB ± 0%  +164.43%  (p=0.000 n=10+10)
/1000000-4    24.2MB ± 0%    64.0MB ± 0%  +164.60%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      21.0 ± 0%   +50.00%  (p=0.000 n=10+10)
/100-4           107 ± 0%       201 ± 0%   +87.85%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     2.00k ± 0%   +97.53%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     20.0k ± 0%   +98.36%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      200k ± 0%   +98.44%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     2.00M ± 0%   +98.45%  (p=0.000 n=10+10)
```

deque vs [CustomSliceQueue](testdata_test.go) - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillSliceQueue.txt
name        old time/op    new time/op    delta
/0-4          40.5ns ± 0%    40.4ns ± 5%      ~     (p=0.062 n=10+9)
/1-4           160ns ± 1%     104ns ± 1%   -35.28%  (p=0.000 n=10+9)
/10-4          553ns ± 2%     640ns ± 1%   +15.76%  (p=0.000 n=10+8)
/100-4        4.15µs ± 1%    3.98µs ± 1%    -3.92%  (p=0.000 n=10+9)
/1000-4       34.6µs ± 2%    35.2µs ± 2%    +1.77%  (p=0.003 n=9+9)
/10000-4       348µs ± 1%     366µs ± 1%    +5.09%  (p=0.000 n=8+8)
/100000-4     3.88ms ± 2%    7.57ms ± 1%   +94.95%  (p=0.000 n=9+8)
/1000000-4    42.7ms ± 2%    78.5ms ± 1%   +84.04%  (p=0.000 n=10+7)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     24.0B ± 0%   -62.50%  (p=0.000 n=10+10)
/1-4            160B ± 0%       48B ± 0%   -70.00%  (p=0.000 n=10+10)
/10-4           432B ± 0%      432B ± 0%      ~     (all equal)
/100-4        4.48kB ± 0%    3.66kB ± 0%   -18.21%  (p=0.000 n=10+10)
/1000-4       25.2kB ± 0%    41.2kB ± 0%   +63.83%  (p=0.000 n=10+10)
/10000-4       243kB ± 0%     518kB ± 0%  +113.44%  (p=0.000 n=10+9)
/100000-4     2.42MB ± 0%    5.70MB ± 0%  +135.57%  (p=0.000 n=10+9)
/1000000-4    24.2MB ± 0%    57.7MB ± 0%  +138.46%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      16.0 ± 0%   +14.29%  (p=0.000 n=10+10)
/100-4           107 ± 0%       109 ± 0%    +1.87%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%      ~     (all equal)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.62%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs [CustomSliceQueue](testdata_test.go) - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillSliceStack.txt
name        old time/op    new time/op    delta
/0-4          41.0ns ± 4%    39.9ns ± 1%    -2.69%  (p=0.000 n=9+9)
/1-4           160ns ± 1%     101ns ± 0%   -36.74%  (p=0.000 n=10+9)
/10-4          564ns ± 0%     628ns ± 0%   +11.35%  (p=0.000 n=10+10)
/100-4        4.23µs ± 1%    3.85µs ± 1%    -8.94%  (p=0.000 n=10+9)
/1000-4       34.9µs ± 2%    33.6µs ± 1%    -3.60%  (p=0.000 n=9+9)
/10000-4       347µs ± 1%     354µs ± 1%    +1.96%  (p=0.000 n=9+10)
/100000-4     3.85ms ± 1%    7.46ms ± 1%   +93.56%  (p=0.000 n=10+10)
/1000000-4    42.1ms ± 1%    76.3ms ± 3%   +81.26%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     24.0B ± 0%   -62.50%  (p=0.000 n=10+10)
/1-4            160B ± 0%       48B ± 0%   -70.00%  (p=0.000 n=10+10)
/10-4           432B ± 0%      432B ± 0%      ~     (all equal)
/100-4        4.48kB ± 0%    3.66kB ± 0%   -18.21%  (p=0.000 n=10+10)
/1000-4       25.2kB ± 0%    41.2kB ± 0%   +63.83%  (p=0.000 n=10+10)
/10000-4       243kB ± 0%     518kB ± 0%  +113.44%  (p=0.000 n=10+10)
/100000-4     2.42MB ± 0%    5.70MB ± 0%  +135.56%  (p=0.000 n=10+10)
/1000000-4    24.2MB ± 0%    57.7MB ± 0%  +138.46%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      16.0 ± 0%   +14.29%  (p=0.000 n=10+10)
/100-4           107 ± 0%       109 ± 0%    +1.87%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%      ~     (all equal)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.62%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs [phf](https://github.com/phf/go-queue) - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillPhfQueue.txt
name        old time/op    new time/op    delta
/0-4          40.5ns ± 0%    73.5ns ± 1%   +81.52%  (p=0.000 n=10+9)
/1-4           160ns ± 1%     109ns ± 0%   -31.98%  (p=0.000 n=10+9)
/10-4          553ns ± 2%     789ns ± 1%   +42.61%  (p=0.000 n=10+10)
/100-4        4.15µs ± 1%    5.12µs ± 1%   +23.40%  (p=0.000 n=10+10)
/1000-4       34.6µs ± 2%    42.0µs ± 1%   +21.63%  (p=0.000 n=9+9)
/10000-4       348µs ± 1%     449µs ± 1%   +28.71%  (p=0.000 n=8+9)
/100000-4     3.88ms ± 2%    7.62ms ± 1%   +96.23%  (p=0.000 n=9+9)
/1000000-4    42.7ms ± 2%    66.3ms ± 2%   +55.35%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%      ~     (all equal)
/1-4            160B ± 0%       80B ± 0%   -50.00%  (p=0.000 n=10+10)
/10-4           432B ± 0%      832B ± 0%   +92.59%  (p=0.000 n=10+10)
/100-4        4.48kB ± 0%    7.65kB ± 0%   +70.71%  (p=0.000 n=10+10)
/1000-4       25.2kB ± 0%    65.1kB ± 0%  +158.49%  (p=0.000 n=10+10)
/10000-4       243kB ± 0%     946kB ± 0%  +290.20%  (p=0.000 n=10+10)
/100000-4     2.42MB ± 0%    7.89MB ± 0%  +226.05%  (p=0.000 n=10+9)
/1000000-4    24.2MB ± 0%    66.3MB ± 0%  +174.24%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      17.0 ± 0%   +21.43%  (p=0.000 n=10+10)
/100-4           107 ± 0%       113 ± 0%    +5.61%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.02k ± 0%    +0.59%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.56%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+9)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs [phf](https://github.com/phf/go-queue) - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillPhfStack.txt
name        old time/op    new time/op    delta
/0-4          41.0ns ± 4%    72.9ns ± 1%   +78.07%  (p=0.000 n=9+9)
/1-4           160ns ± 1%     107ns ± 0%   -33.01%  (p=0.000 n=10+8)
/10-4          564ns ± 0%     843ns ±21%   +49.50%  (p=0.000 n=10+10)
/100-4        4.23µs ± 1%    5.06µs ± 1%   +19.70%  (p=0.000 n=10+9)
/1000-4       34.9µs ± 2%    42.0µs ± 1%   +20.24%  (p=0.000 n=9+8)
/10000-4       347µs ± 1%     449µs ± 0%   +29.51%  (p=0.000 n=9+6)
/100000-4     3.85ms ± 1%    7.63ms ± 1%   +98.05%  (p=0.000 n=10+10)
/1000000-4    42.1ms ± 1%    65.7ms ± 2%   +56.04%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%      ~     (all equal)
/1-4            160B ± 0%       80B ± 0%   -50.00%  (p=0.000 n=10+10)
/10-4           432B ± 0%      832B ± 0%   +92.59%  (p=0.000 n=10+10)
/100-4        4.48kB ± 0%    7.65kB ± 0%   +70.71%  (p=0.000 n=10+10)
/1000-4       25.2kB ± 0%    65.1kB ± 0%  +158.49%  (p=0.000 n=10+10)
/10000-4       243kB ± 0%     946kB ± 0%  +290.20%  (p=0.000 n=10+10)
/100000-4     2.42MB ± 0%    7.89MB ± 0%  +226.05%  (p=0.000 n=10+10)
/1000000-4    24.2MB ± 0%    66.3MB ± 0%  +174.24%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      17.0 ± 0%   +21.43%  (p=0.000 n=10+10)
/100-4           107 ± 0%       113 ± 0%    +5.61%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.02k ± 0%    +0.59%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.56%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs [gammazero](https://github.com/gammazero/deque) - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillGammazeroQueue.txt
name        old time/op    new time/op    delta
/0-4          40.5ns ± 0%    41.0ns ± 8%      ~     (p=0.562 n=10+9)
/1-4           160ns ± 1%     167ns ± 8%      ~     (p=0.062 n=10+10)
/10-4          553ns ± 2%     477ns ± 5%   -13.74%  (p=0.000 n=10+9)
/100-4        4.15µs ± 1%    5.29µs ±16%   +27.69%  (p=0.000 n=10+10)
/1000-4       34.6µs ± 2%    45.6µs ±13%   +31.90%  (p=0.000 n=9+10)
/10000-4       348µs ± 1%     443µs ± 1%   +27.27%  (p=0.000 n=8+10)
/100000-4     3.88ms ± 2%    7.65ms ± 5%   +97.03%  (p=0.000 n=9+9)
/1000000-4    42.7ms ± 2%    66.5ms ± 2%   +55.79%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            160B ± 0%      320B ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           432B ± 0%      464B ± 0%    +7.41%  (p=0.000 n=10+10)
/100-4        4.48kB ± 0%    7.28kB ± 0%   +62.50%  (p=0.000 n=10+10)
/1000-4       25.2kB ± 0%    64.7kB ± 0%  +157.02%  (p=0.000 n=10+10)
/10000-4       243kB ± 0%     946kB ± 0%  +290.05%  (p=0.000 n=10+10)
/100000-4     2.42MB ± 0%    7.89MB ± 0%  +226.03%  (p=0.000 n=10+10)
/1000000-4    24.2MB ± 0%    66.3MB ± 0%  +174.23%  (p=0.000 n=10+8)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      12.0 ± 0%   -14.29%  (p=0.000 n=10+10)
/100-4           107 ± 0%       108 ± 0%    +0.93%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%    +0.10%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.60%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs [gammazero](https://github.com/gammazero/deque) - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillGammazeroStack.txt
name        old time/op    new time/op    delta
/0-4          41.0ns ± 4%    39.6ns ± 0%    -3.31%  (p=0.000 n=9+10)
/1-4           160ns ± 1%     154ns ± 1%    -3.64%  (p=0.000 n=10+10)
/10-4          564ns ± 0%     468ns ± 1%   -16.94%  (p=0.000 n=10+8)
/100-4        4.23µs ± 1%    4.75µs ± 1%   +12.39%  (p=0.000 n=10+9)
/1000-4       34.9µs ± 2%    41.8µs ± 1%   +19.90%  (p=0.000 n=9+10)
/10000-4       347µs ± 1%     440µs ± 1%   +26.90%  (p=0.000 n=9+10)
/100000-4     3.85ms ± 1%    7.65ms ± 1%   +98.54%  (p=0.000 n=10+9)
/1000000-4    42.1ms ± 1%    65.8ms ± 2%   +56.27%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            160B ± 0%      320B ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           432B ± 0%      464B ± 0%    +7.41%  (p=0.000 n=10+10)
/100-4        4.48kB ± 0%    7.28kB ± 0%   +62.50%  (p=0.000 n=10+10)
/1000-4       25.2kB ± 0%    64.7kB ± 0%  +157.02%  (p=0.000 n=10+10)
/10000-4       243kB ± 0%     946kB ± 0%  +290.04%  (p=0.000 n=10+10)
/100000-4     2.42MB ± 0%    7.89MB ± 0%  +226.03%  (p=0.000 n=10+9)
/1000000-4    24.2MB ± 0%    66.3MB ± 0%  +174.23%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      12.0 ± 0%   -14.29%  (p=0.000 n=10+10)
/100-4           107 ± 0%       108 ± 0%    +0.93%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%    +0.10%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.60%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs [gostl](https://github.com/liyue201/gostl/blob/master/ds/deque/deque.go) - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillGostlQueue.txt
name        old time/op    new time/op    delta
/0-4          40.5ns ± 0%    82.6ns ± 1%  +103.91%  (p=0.000 n=10+8)
/1-4           160ns ± 1%     577ns ± 1%  +260.58%  (p=0.000 n=10+10)
/10-4          553ns ± 2%    1354ns ± 1%  +144.86%  (p=0.000 n=10+10)
/100-4        4.15µs ± 1%    8.85µs ± 1%  +113.35%  (p=0.000 n=10+9)
/1000-4       34.6µs ± 2%    86.6µs ± 2%  +150.59%  (p=0.000 n=9+9)
/10000-4       348µs ± 1%     871µs ± 5%  +150.04%  (p=0.000 n=8+10)
/100000-4     3.88ms ± 2%    9.24ms ± 1%  +137.81%  (p=0.000 n=9+10)
/1000000-4    42.7ms ± 2%    97.4ms ± 2%  +128.32%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     88.0B ± 0%   +37.50%  (p=0.000 n=10+10)
/1-4            160B ± 0%     1216B ± 0%  +660.00%  (p=0.000 n=10+10)
/10-4           432B ± 0%     1360B ± 0%  +214.81%  (p=0.000 n=10+10)
/100-4        4.48kB ± 0%    2.80kB ± 0%   -37.50%  (p=0.000 n=10+10)
/1000-4       25.2kB ± 0%    25.2kB ± 0%    +0.16%  (p=0.000 n=10+10)
/10000-4       243kB ± 0%     250kB ± 0%    +3.06%  (p=0.000 n=10+10)
/100000-4     2.42MB ± 0%    2.48MB ± 0%    +2.67%  (p=0.000 n=10+10)
/1000000-4    24.2MB ± 0%    24.8MB ± 0%    +2.68%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%      9.00 ± 0%  +125.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      18.0 ± 0%   +28.57%  (p=0.000 n=10+10)
/100-4           107 ± 0%       108 ± 0%    +0.93%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.03k ± 0%    +2.07%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.2k ± 0%    +1.15%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      102k ± 0%    +0.84%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.02M ± 0%    +0.78%  (p=0.000 n=10+10)
```

deque vs [gostl](https://github.com/liyue201/gostl/blob/master/ds/deque/deque.go) - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillGostlStack.txt
name        old time/op    new time/op    delta
/0-4          41.0ns ± 4%    83.3ns ± 1%  +103.35%  (p=0.000 n=9+9)
/1-4           160ns ± 1%     582ns ± 1%  +263.34%  (p=0.000 n=10+9)
/10-4          564ns ± 0%    1466ns ± 1%  +160.02%  (p=0.000 n=10+9)
/100-4        4.23µs ± 1%   10.00µs ± 1%  +136.62%  (p=0.000 n=10+10)
/1000-4       34.9µs ± 2%    98.1µs ± 1%  +181.25%  (p=0.000 n=9+8)
/10000-4       347µs ± 1%     985µs ± 1%  +183.85%  (p=0.000 n=9+9)
/100000-4     3.85ms ± 1%   10.44ms ± 0%  +170.91%  (p=0.000 n=10+10)
/1000000-4    42.1ms ± 1%   108.9ms ± 1%  +158.59%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     88.0B ± 0%   +37.50%  (p=0.000 n=10+10)
/1-4            160B ± 0%     1216B ± 0%  +660.00%  (p=0.000 n=10+10)
/10-4           432B ± 0%     1360B ± 0%  +214.81%  (p=0.000 n=10+10)
/100-4        4.48kB ± 0%    2.80kB ± 0%   -37.50%  (p=0.000 n=10+10)
/1000-4       25.2kB ± 0%    25.2kB ± 0%    +0.16%  (p=0.000 n=10+10)
/10000-4       243kB ± 0%     250kB ± 0%    +3.06%  (p=0.000 n=10+10)
/100000-4     2.42MB ± 0%    2.48MB ± 0%    +2.67%  (p=0.000 n=10+9)
/1000000-4    24.2MB ± 0%    24.8MB ± 0%    +2.68%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%      9.00 ± 0%  +125.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      18.0 ± 0%   +28.57%  (p=0.000 n=10+10)
/100-4           107 ± 0%       108 ± 0%    +0.93%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.03k ± 0%    +2.07%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.2k ± 0%    +1.15%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      102k ± 0%    +0.84%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.02M ± 0%    +0.78%  (p=0.000 n=10+10)
```

deque vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillCookiejarQueue.txt
name        old time/op    new time/op    delta
/0-4          40.5ns ± 0%  8242.8ns ± 1%   +20245.06%  (p=0.000 n=10+10)
/1-4           160ns ± 1%    8524ns ± 3%    +5228.35%  (p=0.000 n=10+10)
/10-4          553ns ± 2%    8922ns ± 2%    +1513.02%  (p=0.000 n=10+10)
/100-4        4.15µs ± 1%   12.08µs ± 6%     +191.31%  (p=0.000 n=10+9)
/1000-4       34.6µs ± 2%    39.9µs ± 2%      +15.55%  (p=0.000 n=9+10)
/10000-4       348µs ± 1%     351µs ± 1%       +0.67%  (p=0.036 n=8+9)
/100000-4     3.88ms ± 2%    3.54ms ± 1%       -8.88%  (p=0.000 n=9+9)
/1000000-4    42.7ms ± 2%    41.6ms ± 3%       -2.57%  (p=0.001 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%  65672.0B ± 0%  +102512.50%  (p=0.000 n=10+10)
/1-4            160B ± 0%    65688B ± 0%   +40955.00%  (p=0.000 n=10+10)
/10-4           432B ± 0%    65832B ± 0%   +15138.89%  (p=0.000 n=10+10)
/100-4        4.48kB ± 0%   67.27kB ± 0%    +1401.61%  (p=0.000 n=10+10)
/1000-4       25.2kB ± 0%    81.7kB ± 0%     +224.51%  (p=0.000 n=10+10)
/10000-4       243kB ± 0%     357kB ± 0%      +47.15%  (p=0.000 n=10+10)
/100000-4     2.42MB ± 0%    3.25MB ± 0%      +34.14%  (p=0.000 n=10+9)
/1000000-4    24.2MB ± 0%    32.8MB ± 0%      +35.73%  (p=0.000 n=10+6)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%      4.00 ± 0%         ~     (all equal)
/10-4           14.0 ± 0%      13.0 ± 0%       -7.14%  (p=0.000 n=10+10)
/100-4           107 ± 0%       103 ± 0%       -3.74%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.00k ± 0%       -0.99%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%       -0.75%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%       -0.73%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%       -0.73%  (p=0.000 n=10+10)
```

deque vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillCookiejarStack.txt
name        old time/op    new time/op    delta
/0-4          41.0ns ± 4%  8444.7ns ± 2%   +20513.59%  (p=0.000 n=9+10)
/1-4           160ns ± 1%    8587ns ± 2%    +5258.71%  (p=0.000 n=10+10)
/10-4          564ns ± 0%    9019ns ± 2%    +1499.56%  (p=0.000 n=10+10)
/100-4        4.23µs ± 1%   11.90µs ± 1%     +181.54%  (p=0.000 n=10+10)
/1000-4       34.9µs ± 2%    39.7µs ± 1%      +13.78%  (p=0.000 n=9+10)
/10000-4       347µs ± 1%     353µs ± 0%       +1.73%  (p=0.000 n=9+9)
/100000-4     3.85ms ± 1%    3.59ms ± 1%       -6.77%  (p=0.000 n=10+9)
/1000000-4    42.1ms ± 1%    42.3ms ± 2%         ~     (p=0.156 n=9+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%  65672.0B ± 0%  +102512.50%  (p=0.000 n=10+10)
/1-4            160B ± 0%    65688B ± 0%   +40955.00%  (p=0.000 n=10+10)
/10-4           432B ± 0%    65832B ± 0%   +15138.89%  (p=0.000 n=10+10)
/100-4        4.48kB ± 0%   67.27kB ± 0%    +1401.61%  (p=0.000 n=10+10)
/1000-4       25.2kB ± 0%    81.7kB ± 0%     +224.51%  (p=0.000 n=10+10)
/10000-4       243kB ± 0%     357kB ± 0%      +47.15%  (p=0.000 n=10+9)
/100000-4     2.42MB ± 0%    3.25MB ± 0%      +34.14%  (p=0.000 n=10+10)
/1000000-4    24.2MB ± 0%    32.8MB ± 0%      +35.73%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%      4.00 ± 0%         ~     (all equal)
/10-4           14.0 ± 0%      13.0 ± 0%       -7.14%  (p=0.000 n=10+10)
/100-4           107 ± 0%       103 ± 0%       -3.74%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.00k ± 0%       -0.99%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%       -0.75%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%       -0.73%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%       -0.73%  (p=0.000 n=10+10)
```

### Microservice Test Results
deque vs [impl7](https://github.com/christianrpetrin/queue-tests/tree/master/queueimpl7/queueimpl7.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceImpl7Queue.txt
name        old time/op    new time/op    delta
/0-4          45.9ns ± 0%    43.7ns ± 1%    -4.70%  (p=0.000 n=8+9)
/1-4           431ns ± 1%     650ns ± 1%   +50.89%  (p=0.000 n=9+9)
/10-4         2.70µs ± 1%    4.70µs ± 1%   +73.95%  (p=0.000 n=9+10)
/100-4        23.8µs ± 2%    31.8µs ± 1%   +33.56%  (p=0.000 n=10+10)
/1000-4        235µs ± 4%     303µs ± 1%   +28.88%  (p=0.000 n=10+10)
/10000-4      2.29ms ± 3%    3.03ms ± 1%   +32.27%  (p=0.000 n=10+10)
/100000-4     26.0ms ± 1%    33.4ms ± 1%   +28.83%  (p=0.000 n=10+9)
/1000000-4     275ms ± 2%     348ms ± 0%   +26.87%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            384B ± 0%      432B ± 0%   +12.50%  (p=0.000 n=10+10)
/10-4         1.90kB ± 0%    6.91kB ± 0%  +263.03%  (p=0.000 n=10+10)
/100-4        16.2kB ± 0%    29.6kB ± 0%   +83.18%  (p=0.000 n=10+10)
/1000-4        123kB ± 0%     261kB ± 0%  +111.73%  (p=0.000 n=10+10)
/10000-4      1.28MB ± 0%    2.58MB ± 0%  +101.78%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    25.8MB ± 0%  +100.70%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     258MB ± 0%  +100.60%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            11.0 ± 0%      17.0 ± 0%   +54.55%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%     109.0 ± 0%   +45.33%  (p=0.000 n=10+10)
/100-4           709 ± 0%       927 ± 0%   +30.75%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     9.11k ± 0%   +29.88%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     91.0k ± 0%   +29.65%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      909k ± 0%   +29.62%  (p=0.000 n=10+9)
/1000000-4     7.02M ± 0%     9.09M ± 0%   +29.62%  (p=0.000 n=10+9)
```

deque vs [list](https://github.com/golang/go/tree/master/src/container/list) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceListQueue.txt
name        old time/op    new time/op    delta
/0-4          45.9ns ± 0%    44.7ns ± 1%    -2.56%  (p=0.000 n=8+10)
/1-4           431ns ± 1%     525ns ± 0%   +21.78%  (p=0.000 n=9+9)
/10-4         2.70µs ± 1%    4.90µs ± 0%   +81.40%  (p=0.000 n=9+10)
/100-4        23.8µs ± 2%    47.9µs ± 0%  +101.55%  (p=0.000 n=10+10)
/1000-4        235µs ± 4%     480µs ± 0%  +103.97%  (p=0.000 n=10+9)
/10000-4      2.29ms ± 3%    5.07ms ± 1%  +121.25%  (p=0.000 n=10+9)
/100000-4     26.0ms ± 1%    71.4ms ± 2%  +175.17%  (p=0.000 n=10+10)
/1000000-4     275ms ± 2%     718ms ± 4%  +161.53%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            384B ± 0%      496B ± 0%   +29.17%  (p=0.000 n=10+10)
/10-4         1.90kB ± 0%    4.53kB ± 0%  +137.82%  (p=0.000 n=10+10)
/100-4        16.2kB ± 0%    44.8kB ± 0%  +177.25%  (p=0.000 n=10+10)
/1000-4        123kB ± 0%     448kB ± 0%  +263.49%  (p=0.000 n=10+9)
/10000-4      1.28MB ± 0%    4.48MB ± 0%  +250.56%  (p=0.000 n=10+9)
/100000-4     12.8MB ± 0%    44.8MB ± 0%  +249.14%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     448MB ± 0%  +249.00%  (p=0.000 n=9+9)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            11.0 ± 0%      15.0 ± 0%   +36.36%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%     141.0 ± 0%   +88.00%  (p=0.000 n=10+10)
/100-4           709 ± 0%      1401 ± 0%   +97.60%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%    14.00k ± 0%   +99.59%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%    140.0k ± 0%   +99.56%  (p=0.000 n=10+10)
/100000-4       702k ± 0%     1400k ± 0%   +99.56%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%    14.00M ± 0%   +99.55%  (p=0.000 n=10+7)
```

deque vs [list](https://github.com/golang/go/tree/master/src/container/list) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceListStack.txt
name        old time/op    new time/op    delta
/0-4          45.9ns ± 1%    44.6ns ± 0%    -2.64%  (p=0.000 n=9+8)
/1-4           360ns ± 1%     526ns ± 1%   +46.22%  (p=0.000 n=9+10)
/10-4         2.53µs ± 0%    4.90µs ± 1%   +93.84%  (p=0.000 n=8+10)
/100-4        23.2µs ± 2%    47.9µs ± 1%  +106.29%  (p=0.000 n=10+10)
/1000-4        229µs ± 1%     479µs ± 1%  +109.44%  (p=0.000 n=9+10)
/10000-4      2.38ms ± 1%    5.37ms ± 2%  +125.75%  (p=0.000 n=9+9)
/100000-4     26.2ms ± 4%    79.4ms ± 2%  +203.72%  (p=0.000 n=10+10)
/1000000-4     274ms ± 2%     739ms ± 5%  +170.28%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            256B ± 0%      496B ± 0%   +93.75%  (p=0.000 n=10+10)
/10-4         1.39kB ± 0%    4.53kB ± 0%  +225.29%  (p=0.000 n=10+10)
/100-4        14.1kB ± 0%    44.8kB ± 0%  +218.52%  (p=0.000 n=10+10)
/1000-4        121kB ± 0%     448kB ± 0%  +269.78%  (p=0.000 n=10+8)
/10000-4      1.27MB ± 0%    4.48MB ± 0%  +251.71%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    44.8MB ± 0%  +249.20%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     448MB ± 0%  +249.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            10.0 ± 0%      15.0 ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4           74.0 ± 0%     141.0 ± 0%   +90.54%  (p=0.000 n=10+10)
/100-4           707 ± 0%      1401 ± 0%   +98.16%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%    14.00k ± 0%   +99.64%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%    140.0k ± 0%   +99.57%  (p=0.000 n=10+10)
/100000-4       702k ± 0%     1400k ± 0%   +99.56%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%    14.00M ± 0%   +99.55%  (p=0.000 n=10+9)
```

deque vs [CustomSliceQueue](testdata_test.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceSliceQueue.txt
name        old time/op    new time/op    delta
/0-4          45.9ns ± 0%    43.6ns ± 1%    -4.97%  (p=0.000 n=8+9)
/1-4           431ns ± 1%     513ns ± 1%   +19.03%  (p=0.000 n=9+9)
/10-4         2.70µs ± 1%    3.72µs ± 1%   +37.56%  (p=0.000 n=9+9)
/100-4        23.8µs ± 2%    27.1µs ± 1%   +14.05%  (p=0.000 n=10+10)
/1000-4        235µs ± 4%     297µs ± 1%   +26.16%  (p=0.000 n=10+8)
/10000-4      2.29ms ± 3%    3.26ms ± 3%   +42.31%  (p=0.000 n=10+10)
/100000-4     26.0ms ± 1%    43.6ms ± 1%   +68.10%  (p=0.000 n=10+10)
/1000000-4     275ms ± 2%     487ms ± 3%   +77.45%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     24.0B ± 0%   -62.50%  (p=0.000 n=10+10)
/1-4            384B ± 0%      224B ± 0%   -41.67%  (p=0.000 n=10+10)
/10-4         1.90kB ± 0%    2.16kB ± 0%   +13.45%  (p=0.000 n=10+10)
/100-4        16.2kB ± 0%    21.3kB ± 0%   +31.85%  (p=0.000 n=10+10)
/1000-4        123kB ± 0%     234kB ± 0%   +89.46%  (p=0.000 n=10+10)
/10000-4      1.28MB ± 0%    2.84MB ± 0%  +122.07%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    32.5MB ± 0%  +153.45%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     335MB ± 0%  +161.23%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            11.0 ± 0%      14.0 ± 0%   +27.27%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%     101.0 ± 0%   +34.67%  (p=0.000 n=10+10)
/100-4           709 ± 0%       822 ± 0%   +15.94%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     8.78k ± 0%   +25.17%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     87.8k ± 0%   +25.21%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      886k ± 0%   +26.25%  (p=0.000 n=10+8)
/1000000-4     7.02M ± 0%     8.94M ± 0%   +27.48%  (p=0.000 n=10+10)
```

deque vs [CustomSliceQueue](testdata_test.go) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceSliceStack.txt
name        old time/op    new time/op    delta
/0-4          45.9ns ± 1%    44.2ns ± 1%   -3.56%  (p=0.000 n=9+9)
/1-4           360ns ± 1%     396ns ± 1%  +10.08%  (p=0.000 n=9+10)
/10-4         2.53µs ± 0%    2.46µs ± 1%   -2.84%  (p=0.000 n=8+10)
/100-4        23.2µs ± 2%    20.8µs ± 1%  -10.38%  (p=0.000 n=10+10)
/1000-4        229µs ± 1%     202µs ± 2%  -11.84%  (p=0.000 n=9+9)
/10000-4      2.38ms ± 1%    2.05ms ± 2%  -13.79%  (p=0.000 n=9+9)
/100000-4     26.2ms ± 4%    26.0ms ± 1%     ~     (p=0.720 n=10+9)
/1000000-4     274ms ± 2%     266ms ± 2%   -2.60%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     24.0B ± 0%  -62.50%  (p=0.000 n=10+10)
/1-4            256B ± 0%      192B ± 0%  -25.00%  (p=0.000 n=10+10)
/10-4         1.39kB ± 0%    1.39kB ± 0%     ~     (all equal)
/100-4        14.1kB ± 0%    13.3kB ± 0%   -5.80%  (p=0.000 n=10+10)
/1000-4        121kB ± 0%     137kB ± 0%  +13.26%  (p=0.000 n=10+10)
/10000-4      1.27MB ± 0%    1.48MB ± 0%  +16.00%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    15.3MB ± 0%  +19.27%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     154MB ± 0%  +19.72%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            10.0 ± 0%      11.0 ± 0%  +10.00%  (p=0.000 n=10+10)
/10-4           74.0 ± 0%      76.0 ± 0%   +2.70%  (p=0.000 n=10+10)
/100-4           707 ± 0%       709 ± 0%   +0.28%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.01k ± 0%     ~     (all equal)
/10000-4       70.2k ± 0%     70.0k ± 0%   -0.19%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%   -0.22%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%   -0.22%  (p=0.000 n=10+9)
```

deque vs [phf](https://github.com/phf/go-queue) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroservicePhfQueue.txt
name        old time/op    new time/op    delta
/0-4          45.9ns ± 0%    77.3ns ± 1%   +68.54%  (p=0.000 n=8+9)
/1-4           431ns ± 1%     429ns ± 0%    -0.55%  (p=0.017 n=9+10)
/10-4         2.70µs ± 1%    3.06µs ± 1%   +13.21%  (p=0.000 n=9+10)
/100-4        23.8µs ± 2%    26.2µs ± 1%   +10.02%  (p=0.000 n=10+10)
/1000-4        235µs ± 4%     243µs ± 1%    +3.21%  (p=0.002 n=10+9)
/10000-4      2.29ms ± 3%    2.59ms ± 1%   +12.80%  (p=0.000 n=10+10)
/100000-4     26.0ms ± 1%    30.0ms ± 2%   +15.44%  (p=0.000 n=10+10)
/1000000-4     275ms ± 2%     302ms ± 1%    +9.91%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%      ~     (all equal)
/1-4            384B ± 0%      272B ± 0%   -29.17%  (p=0.000 n=10+10)
/10-4         1.90kB ± 0%    2.18kB ± 0%   +14.29%  (p=0.000 n=10+10)
/100-4        16.2kB ± 0%    23.0kB ± 0%   +42.24%  (p=0.000 n=10+10)
/1000-4        123kB ± 0%     210kB ± 0%   +70.22%  (p=0.000 n=10+9)
/10000-4      1.28MB ± 0%    2.69MB ± 0%  +110.67%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    23.8MB ± 0%   +85.34%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     213MB ± 0%   +65.67%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            11.0 ± 0%      11.0 ± 0%      ~     (all equal)
/10-4           75.0 ± 0%      79.0 ± 0%    +5.33%  (p=0.000 n=10+10)
/100-4           709 ± 0%       721 ± 0%    +1.69%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.03k ± 0%    +0.26%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%    -0.15%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%    -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%    -0.22%  (p=0.000 n=10+10)
```

deque vs [phf](https://github.com/phf/go-queue) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroservicePhfStack.txt
name        old time/op    new time/op    delta
/0-4          45.9ns ± 1%    76.6ns ± 1%   +67.14%  (p=0.000 n=9+9)
/1-4           360ns ± 1%     434ns ± 0%   +20.71%  (p=0.000 n=9+9)
/10-4         2.53µs ± 0%    3.01µs ± 0%   +19.22%  (p=0.000 n=8+8)
/100-4        23.2µs ± 2%    25.9µs ± 1%   +11.56%  (p=0.000 n=10+10)
/1000-4        229µs ± 1%     243µs ± 2%    +6.04%  (p=0.000 n=9+9)
/10000-4      2.38ms ± 1%    2.57ms ± 0%    +8.06%  (p=0.000 n=9+9)
/100000-4     26.2ms ± 4%    29.9ms ± 2%   +14.24%  (p=0.000 n=10+10)
/1000000-4     274ms ± 2%     298ms ± 2%    +8.78%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%      ~     (all equal)
/1-4            256B ± 0%      272B ± 0%    +6.25%  (p=0.000 n=10+10)
/10-4         1.39kB ± 0%    2.18kB ± 0%   +56.32%  (p=0.000 n=10+10)
/100-4        14.1kB ± 0%    23.0kB ± 0%   +63.41%  (p=0.000 n=10+10)
/1000-4        121kB ± 0%     210kB ± 0%   +73.17%  (p=0.000 n=10+8)
/10000-4      1.27MB ± 0%    2.69MB ± 0%  +111.37%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    23.8MB ± 0%   +85.37%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     213MB ± 0%   +65.67%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            10.0 ± 0%      11.0 ± 0%   +10.00%  (p=0.000 n=10+10)
/10-4           74.0 ± 0%      79.0 ± 0%    +6.76%  (p=0.000 n=10+10)
/100-4           707 ± 0%       721 ± 0%    +1.98%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.03k ± 0%    +0.29%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%    -0.15%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%    -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%    -0.22%  (p=0.000 n=10+10)
```

deque vs [gammazero](https://github.com/gammazero/deque) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceGammazeroQueue.txt
name        old time/op    new time/op    delta
/0-4          45.9ns ± 0%    43.6ns ± 0%    -4.87%  (p=0.000 n=8+9)
/1-4           431ns ± 1%     356ns ± 0%   -17.47%  (p=0.000 n=9+9)
/10-4         2.70µs ± 1%    2.45µs ± 1%    -9.36%  (p=0.000 n=9+9)
/100-4        23.8µs ± 2%    25.4µs ± 0%    +6.95%  (p=0.000 n=10+9)
/1000-4        235µs ± 4%     241µs ± 0%    +2.45%  (p=0.002 n=10+10)
/10000-4      2.29ms ± 3%    2.57ms ± 0%   +12.34%  (p=0.000 n=10+10)
/100000-4     26.0ms ± 1%    30.0ms ± 2%   +15.69%  (p=0.000 n=10+10)
/1000000-4     275ms ± 2%     301ms ± 2%    +9.49%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            384B ± 0%      416B ± 0%    +8.33%  (p=0.000 n=10+10)
/10-4         1.90kB ± 0%    1.42kB ± 0%   -25.21%  (p=0.000 n=10+10)
/100-4        16.2kB ± 0%    22.3kB ± 0%   +37.59%  (p=0.000 n=10+10)
/1000-4        123kB ± 0%     209kB ± 0%   +69.61%  (p=0.000 n=10+8)
/10000-4      1.28MB ± 0%    2.69MB ± 0%  +110.62%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    23.8MB ± 0%   +85.34%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     213MB ± 0%   +65.67%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            11.0 ± 0%       9.0 ± 0%   -18.18%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%      72.0 ± 0%    -4.00%  (p=0.000 n=10+10)
/100-4           709 ± 0%       714 ± 0%    +0.71%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.03k ± 0%    +0.16%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%    -0.16%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%    -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%    -0.22%  (p=0.000 n=10+10)
```

deque vs [gammazero](https://github.com/gammazero/deque) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceGammazeroStack.txt
name        old time/op    new time/op    delta
/0-4          45.9ns ± 1%    43.8ns ± 1%    -4.58%  (p=0.000 n=9+10)
/1-4           360ns ± 1%     360ns ± 1%      ~     (p=0.844 n=9+9)
/10-4         2.53µs ± 0%    2.50µs ± 2%    -1.07%  (p=0.027 n=8+9)
/100-4        23.2µs ± 2%    25.6µs ± 0%   +10.34%  (p=0.000 n=10+8)
/1000-4        229µs ± 1%     242µs ± 1%    +5.65%  (p=0.000 n=9+10)
/10000-4      2.38ms ± 1%    2.60ms ± 1%    +9.33%  (p=0.000 n=9+9)
/100000-4     26.2ms ± 4%    30.0ms ± 1%   +14.57%  (p=0.000 n=10+10)
/1000000-4     274ms ± 2%     298ms ± 1%    +9.00%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            256B ± 0%      416B ± 0%   +62.50%  (p=0.000 n=10+10)
/10-4         1.39kB ± 0%    1.42kB ± 0%    +2.30%  (p=0.000 n=10+10)
/100-4        14.1kB ± 0%    22.3kB ± 0%   +58.07%  (p=0.000 n=10+10)
/1000-4        121kB ± 0%     209kB ± 0%   +72.55%  (p=0.000 n=10+10)
/10000-4      1.27MB ± 0%    2.69MB ± 0%  +111.31%  (p=0.000 n=10+9)
/100000-4     12.8MB ± 0%    23.8MB ± 0%   +85.37%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     213MB ± 0%   +65.67%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            10.0 ± 0%       9.0 ± 0%   -10.00%  (p=0.000 n=10+10)
/10-4           74.0 ± 0%      72.0 ± 0%    -2.70%  (p=0.000 n=10+10)
/100-4           707 ± 0%       714 ± 0%    +0.99%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.03k ± 0%    +0.19%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%    -0.16%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%    -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%    -0.22%  (p=0.000 n=10+10)
```

deque vs [gostl](https://github.com/liyue201/gostl/blob/master/ds/deque/deque.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceGostlQueue.txt
name        old time/op    new time/op    delta
/0-4          45.9ns ± 0%    87.9ns ± 2%   +91.81%  (p=0.000 n=8+9)
/1-4           431ns ± 1%    1243ns ± 1%  +188.38%  (p=0.000 n=9+9)
/10-4         2.70µs ± 1%    7.55µs ± 1%  +179.42%  (p=0.000 n=9+10)
/100-4        23.8µs ± 2%    68.3µs ± 1%  +187.27%  (p=0.000 n=10+10)
/1000-4        235µs ± 4%     685µs ± 1%  +191.26%  (p=0.000 n=10+9)
/10000-4      2.29ms ± 3%    6.86ms ± 2%  +199.21%  (p=0.000 n=10+9)
/100000-4     26.0ms ± 1%    71.3ms ± 1%  +174.80%  (p=0.000 n=10+10)
/1000000-4     275ms ± 2%     724ms ± 2%  +163.62%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     88.0B ± 0%   +37.50%  (p=0.000 n=10+10)
/1-4            384B ± 0%     1352B ± 0%  +252.08%  (p=0.000 n=10+10)
/10-4         1.90kB ± 0%    2.60kB ± 0%   +36.55%  (p=0.000 n=10+10)
/100-4        16.2kB ± 0%    14.8kB ± 0%    -8.26%  (p=0.000 n=10+10)
/1000-4        123kB ± 0%     156kB ± 0%   +26.23%  (p=0.000 n=10+10)
/10000-4      1.28MB ± 0%    1.54MB ± 0%   +20.67%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    15.4MB ± 0%   +19.80%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     154MB ± 0%   +19.71%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            11.0 ± 0%      18.0 ± 0%   +63.64%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%     101.0 ± 0%   +34.67%  (p=0.000 n=10+10)
/100-4           709 ± 0%       911 ± 0%   +28.49%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     9.07k ± 0%   +29.32%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     90.4k ± 0%   +28.86%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      903k ± 0%   +28.75%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     9.03M ± 0%   +28.73%  (p=0.000 n=10+10)
```

deque vs [gostl](https://github.com/liyue201/gostl/blob/master/ds/deque/deque.go) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceGostlStack.txt
name        old time/op    new time/op    delta
/0-4          45.9ns ± 1%    87.1ns ± 1%   +89.85%  (p=0.000 n=9+8)
/1-4           360ns ± 1%    1307ns ± 3%  +263.31%  (p=0.000 n=9+9)
/10-4         2.53µs ± 0%    8.37µs ± 1%  +231.01%  (p=0.000 n=8+10)
/100-4        23.2µs ± 2%    76.4µs ± 0%  +228.90%  (p=0.000 n=10+8)
/1000-4        229µs ± 1%     770µs ± 0%  +236.41%  (p=0.000 n=9+10)
/10000-4      2.38ms ± 1%    7.72ms ± 1%  +224.42%  (p=0.000 n=9+10)
/100000-4     26.2ms ± 4%    80.2ms ± 1%  +206.81%  (p=0.000 n=10+10)
/1000000-4     274ms ± 2%     809ms ± 1%  +195.78%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     88.0B ± 0%   +37.50%  (p=0.000 n=10+10)
/1-4            256B ± 0%     1352B ± 0%  +428.12%  (p=0.000 n=10+10)
/10-4         1.39kB ± 0%    2.60kB ± 0%   +86.78%  (p=0.000 n=10+10)
/100-4        14.1kB ± 0%    14.8kB ± 0%    +5.40%  (p=0.000 n=10+10)
/1000-4        121kB ± 0%     157kB ± 0%   +29.37%  (p=0.000 n=10+10)
/10000-4      1.27MB ± 0%    1.54MB ± 0%   +20.99%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    15.4MB ± 0%   +19.81%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     154MB ± 0%   +19.71%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/1-4            10.0 ± 0%      18.0 ± 0%   +80.00%  (p=0.000 n=10+10)
/10-4           74.0 ± 0%     101.0 ± 0%   +36.49%  (p=0.000 n=10+10)
/100-4           707 ± 0%       911 ± 0%   +28.85%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     9.08k ± 0%   +29.46%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     90.4k ± 0%   +28.87%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      903k ± 0%   +28.75%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     9.03M ± 0%   +28.73%  (p=0.000 n=10+9)
```

deque vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceCookiejarQueue.txt
name        old time/op    new time/op    delta
/0-4          45.9ns ± 0%  8574.9ns ± 2%   +18601.56%  (p=0.000 n=8+10)
/1-4           431ns ± 1%    8934ns ± 2%    +1972.99%  (p=0.000 n=9+9)
/10-4         2.70µs ± 1%   11.12µs ± 2%     +311.32%  (p=0.000 n=9+10)
/100-4        23.8µs ± 2%    29.3µs ± 1%      +23.04%  (p=0.000 n=10+10)
/1000-4        235µs ± 4%     224µs ± 1%       -4.81%  (p=0.000 n=10+10)
/10000-4      2.29ms ± 3%    2.12ms ± 0%       -7.50%  (p=0.000 n=10+9)
/100000-4     26.0ms ± 1%    23.6ms ± 1%       -9.05%  (p=0.000 n=10+9)
/1000000-4     275ms ± 2%     239ms ± 1%      -13.03%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%  65672.0B ± 0%  +102512.50%  (p=0.000 n=10+10)
/1-4            384B ± 0%    65784B ± 0%   +17031.25%  (p=0.000 n=10+10)
/10-4         1.90kB ± 0%   66.79kB ± 0%    +3407.98%  (p=0.000 n=10+10)
/100-4        16.2kB ± 0%    76.9kB ± 0%     +375.22%  (p=0.000 n=10+10)
/1000-4        123kB ± 0%     243kB ± 0%      +97.35%  (p=0.000 n=10+9)
/10000-4      1.28MB ± 0%    1.38MB ± 0%       +8.18%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    12.9MB ± 0%       +0.63%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     129MB ± 0%       +0.42%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            11.0 ± 0%      10.0 ± 0%       -9.09%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%      73.0 ± 0%       -2.67%  (p=0.000 n=10+10)
/100-4           709 ± 0%       703 ± 0%       -0.85%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.00k ± 0%       -0.14%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%       -0.21%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%       -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%       -0.22%  (p=0.000 n=10+10)
```

deque vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceCookiejarStack.txt
name        old time/op    new time/op    delta
/0-4          45.9ns ± 1%  8828.7ns ± 3%   +19154.68%  (p=0.000 n=9+10)
/1-4           360ns ± 1%    9156ns ± 2%    +2445.17%  (p=0.000 n=9+9)
/10-4         2.53µs ± 0%   11.10µs ± 1%     +339.02%  (p=0.000 n=8+9)
/100-4        23.2µs ± 2%    29.1µs ± 0%      +25.29%  (p=0.000 n=10+10)
/1000-4        229µs ± 1%     215µs ± 1%       -6.10%  (p=0.000 n=9+9)
/10000-4      2.38ms ± 1%    2.11ms ± 0%      -11.49%  (p=0.000 n=9+9)
/100000-4     26.2ms ± 4%    23.6ms ± 0%       -9.92%  (p=0.000 n=10+8)
/1000000-4     274ms ± 2%     238ms ± 1%      -13.09%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%  65672.0B ± 0%  +102512.50%  (p=0.000 n=10+10)
/1-4            256B ± 0%    65784B ± 0%   +25596.88%  (p=0.000 n=10+10)
/10-4         1.39kB ± 0%   66.79kB ± 0%    +4698.28%  (p=0.000 n=10+10)
/100-4        14.1kB ± 0%    76.9kB ± 0%     +445.97%  (p=0.000 n=10+10)
/1000-4        121kB ± 0%     178kB ± 0%      +46.63%  (p=0.000 n=10+9)
/10000-4      1.27MB ± 0%    1.32MB ± 0%       +3.38%  (p=0.000 n=10+10)
/100000-4     12.8MB ± 0%    12.8MB ± 0%       +0.13%  (p=0.000 n=10+10)
/1000000-4     128MB ± 0%     129MB ± 0%       +0.36%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            10.0 ± 0%      10.0 ± 0%         ~     (all equal)
/10-4           74.0 ± 0%      73.0 ± 0%       -1.35%  (p=0.000 n=10+10)
/100-4           707 ± 0%       703 ± 0%       -0.57%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.00k ± 0%       -0.14%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%       -0.21%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%       -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%       -0.22%  (p=0.000 n=10+10)
```


### Other Test Results
#### deque vs impl7 - FIFO queue
deque vs impl7 - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillImpl7Queue.txt
name       old time/op    new time/op    delta
/1-4         3.72µs ± 1%   10.21µs ± 1%  +174.76%  (p=0.000 n=9+9)
/10-4        35.9µs ± 3%    69.2µs ± 1%   +92.87%  (p=0.000 n=10+10)
/100-4        351µs ± 3%     410µs ± 1%   +16.84%  (p=0.000 n=9+8)
/1000-4      3.43ms ± 2%    3.63ms ± 1%    +5.71%  (p=0.000 n=10+8)
/10000-4     35.1ms ± 4%    36.7ms ± 0%    +4.46%  (p=0.000 n=10+8)
/100000-4     389ms ± 1%     372ms ± 0%    -4.57%  (p=0.000 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    68.8kB ± 0%  +330.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     421kB ± 0%  +163.00%  (p=0.000 n=10+8)
/1000-4      1.60MB ± 0%    3.32MB ± 0%  +107.30%  (p=0.000 n=10+10)
/10000-4     23.1MB ± 0%    32.3MB ± 0%   +39.48%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     323MB ± 0%   +34.00%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       300 ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.60k ± 0%   +60.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.8k ± 0%    +8.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      102k ± 0%    +2.20%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.02M ± 0%    +0.93%  (p=0.000 n=9+9)
/100000-4     10.1M ± 0%     10.2M ± 0%    +0.79%  (p=0.000 n=10+10)
```

deque vs impl7 - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullImpl7Queue.txt
name       old time/op    new time/op    delta
/1-4         3.53µs ± 2%    3.75µs ± 0%    +6.16%  (p=0.000 n=10+10)
/10-4        34.0µs ± 1%    37.3µs ± 1%    +9.72%  (p=0.000 n=9+8)
/100-4        332µs ± 1%     365µs ± 0%    +9.95%  (p=0.000 n=9+9)
/1000-4      3.29ms ± 1%    3.63ms ± 0%   +10.48%  (p=0.000 n=9+9)
/10000-4     35.3ms ± 1%    36.4ms ± 1%    +3.25%  (p=0.000 n=10+10)
/100000-4     388ms ± 1%     369ms ± 1%    -4.82%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    3.23kB ± 0%  +101.56%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    32.2kB ± 0%  +101.56%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     323kB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    3.23MB ± 0%  +101.56%  (p=0.000 n=10+10)
/10000-4     23.1MB ± 0%    32.3MB ± 0%   +39.37%  (p=0.000 n=8+10)
/100000-4     241MB ± 0%     323MB ± 0%   +33.91%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       101 ± 0%    +1.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.01k ± 0%    +1.50%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.2k ± 0%    +1.56%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      102k ± 0%    +1.56%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.02M ± 0%    +0.88%  (p=0.000 n=10+9)
/100000-4     10.1M ± 0%     10.2M ± 0%    +0.79%  (p=0.000 n=10+10)
```

deque vs impl7 - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4           195ns ± 1%     236ns ± 1%   +20.63%  (p=0.000 n=10+9)
/10-4         1.05µs ± 1%    1.51µs ± 1%   +44.58%  (p=0.000 n=9+9)
/100-4        7.35µs ± 1%    7.60µs ± 0%    +3.36%  (p=0.000 n=9+9)
/1000-4       67.2µs ± 1%    69.6µs ± 0%    +3.57%  (p=0.000 n=9+9)
/10000-4       688µs ± 2%     701µs ± 0%    +1.86%  (p=0.021 n=10+8)
/100000-4     7.70ms ± 1%    8.30ms ± 2%    +7.80%  (p=0.000 n=9+10)
/1000000-4    84.1ms ± 2%    88.2ms ± 3%    +4.87%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%      160B ± 0%    -9.09%  (p=0.000 n=10+10)
/10-4         1.10kB ± 0%    2.98kB ± 0%  +169.57%  (p=0.000 n=10+10)
/100-4        6.08kB ± 0%    7.94kB ± 0%   +30.53%  (p=0.000 n=10+10)
/1000-4       43.3kB ± 0%    65.9kB ± 0%   +52.22%  (p=0.000 n=10+10)
/10000-4       405kB ± 0%     647kB ± 0%   +59.94%  (p=0.000 n=10+10)
/100000-4     4.02MB ± 0%    6.45MB ± 0%   +60.48%  (p=0.000 n=10+10)
/1000000-4    40.2MB ± 0%    64.5MB ± 0%   +60.49%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      6.00 ± 0%   +20.00%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      29.0 ± 0%   +16.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       211 ± 0%    +1.93%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.04k ± 0%    +1.19%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.3k ± 0%    +1.18%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      203k ± 0%    +1.17%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.03M ± 0%    +1.17%  (p=0.000 n=10+10)
```

deque vs impl7 - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4          36.2ns ± 1%   102.0ns ± 1%  +181.48%  (p=0.000 n=9+8)
/10-4          368ns ± 1%    1028ns ± 0%  +179.50%  (p=0.000 n=10+10)
/100-4        3.60µs ± 2%   10.15µs ± 1%  +181.99%  (p=0.000 n=8+9)
/1000-4       35.6µs ± 1%   101.7µs ± 1%  +185.95%  (p=0.000 n=10+10)
/10000-4       364µs ± 4%    1014µs ± 0%  +178.63%  (p=0.000 n=10+9)
/100000-4     3.58ms ± 1%   10.18ms ± 1%  +184.24%  (p=0.000 n=9+10)
/1000000-4    35.8ms ± 1%   101.5ms ± 1%  +183.36%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+7)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=9+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      3.00 ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      30.0 ± 0%  +200.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       300 ± 0%  +200.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     3.00k ± 0%  +200.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     30.0k ± 0%  +200.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      300k ± 0%  +200.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     3.00M ± 0%  +200.00%  (p=0.000 n=10+8)
```

deque vs impl7 - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%    36.6ns ± 2%    +8.01%  (p=0.000 n=8+10)
/10-4          344ns ± 1%     369ns ± 1%    +7.44%  (p=0.000 n=10+9)
/100-4        3.34µs ± 1%    3.60µs ± 1%    +7.58%  (p=0.000 n=8+10)
/1000-4       33.9µs ± 4%    35.9µs ± 1%    +5.84%  (p=0.000 n=10+10)
/10000-4       348µs ± 2%     359µs ± 1%    +2.98%  (p=0.000 n=10+10)
/100000-4     3.47ms ± 1%    3.58ms ± 1%    +3.31%  (p=0.000 n=10+9)
/1000000-4    35.0ms ± 2%    36.0ms ± 1%    +2.86%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     32.0B ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      322B ± 0%  +101.25%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    3.23kB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    32.2kB ± 0%  +101.56%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     323kB ± 0%  +101.56%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    3.23MB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    32.3MB ± 0%  +101.56%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%      ~     (all equal)
/100-4           100 ± 0%       101 ± 0%    +1.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     1.01k ± 0%    +1.50%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     10.2k ± 0%    +1.56%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      102k ± 0%    +1.56%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     1.02M ± 0%    +1.56%  (p=0.000 n=10+10)
```

#### deque vs list - FIFO queue
deque vs list - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillListQueue.txt
name       old time/op    new time/op    delta
/1-4         3.72µs ± 1%    7.26µs ± 0%   +95.38%  (p=0.000 n=9+8)
/10-4        35.9µs ± 3%    70.9µs ± 0%   +97.66%  (p=0.000 n=10+10)
/100-4        351µs ± 3%     701µs ± 1%  +100.14%  (p=0.000 n=9+10)
/1000-4      3.43ms ± 2%    6.96ms ± 1%  +102.77%  (p=0.000 n=10+9)
/10000-4     35.1ms ± 4%    70.4ms ± 1%  +100.55%  (p=0.000 n=10+10)
/100000-4     389ms ± 1%    1432ms ± 3%  +267.74%  (p=0.000 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+8)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +299.99%  (p=0.000 n=10+10)
/10000-4     23.1MB ± 0%    64.0MB ± 0%  +176.58%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     640MB ± 0%  +165.74%  (p=0.000 n=10+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     2.00M ± 0%   +98.65%  (p=0.000 n=9+8)
/100000-4     10.1M ± 0%     20.0M ± 0%   +98.47%  (p=0.000 n=10+8)
```

deque vs list - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullListQueue.txt
name       old time/op    new time/op    delta
/1-4         3.53µs ± 2%    8.24µs ± 0%  +133.57%  (p=0.000 n=10+8)
/10-4        34.0µs ± 1%    93.3µs ± 3%  +174.25%  (p=0.000 n=9+10)
/100-4        332µs ± 1%     940µs ± 2%  +182.78%  (p=0.000 n=9+10)
/1000-4      3.29ms ± 1%    9.56ms ± 4%  +190.68%  (p=0.000 n=9+9)
/10000-4     35.3ms ± 1%   107.4ms ± 3%  +204.30%  (p=0.000 n=10+10)
/100000-4     388ms ± 1%    1365ms ± 2%  +252.11%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4     23.1MB ± 0%    64.0MB ± 0%  +176.58%  (p=0.000 n=8+10)
/100000-4     241MB ± 0%     640MB ± 0%  +165.75%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     2.00M ± 0%   +98.65%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     20.0M ± 0%   +98.47%  (p=0.000 n=10+10)
```

deque vs list - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseListQueue.txt
name        old time/op    new time/op    delta
/1-4           195ns ± 1%     183ns ± 1%    -6.54%  (p=0.000 n=10+10)
/10-4         1.05µs ± 1%    1.44µs ± 1%   +37.35%  (p=0.000 n=9+9)
/100-4        7.35µs ± 1%   13.87µs ± 0%   +88.54%  (p=0.000 n=9+9)
/1000-4       67.2µs ± 1%   138.4µs ± 0%  +105.82%  (p=0.000 n=9+10)
/10000-4       688µs ± 2%    1385µs ± 0%  +101.19%  (p=0.000 n=10+8)
/100000-4     7.70ms ± 1%   24.53ms ± 3%  +218.69%  (p=0.000 n=9+9)
/1000000-4    84.1ms ± 2%   238.1ms ±10%  +183.03%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%      176B ± 0%      ~     (all equal)
/10-4         1.10kB ± 0%    1.33kB ± 0%   +20.29%  (p=0.000 n=10+10)
/100-4        6.08kB ± 0%   12.85kB ± 0%  +111.32%  (p=0.000 n=10+10)
/1000-4       43.3kB ± 0%   128.0kB ± 0%  +195.97%  (p=0.000 n=10+10)
/10000-4       405kB ± 0%    1280kB ± 0%  +216.36%  (p=0.000 n=10+10)
/100000-4     4.02MB ± 0%   12.80MB ± 0%  +218.38%  (p=0.000 n=10+10)
/1000000-4    40.2MB ± 0%   128.0MB ± 0%  +218.49%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%      ~     (all equal)
/10-4           25.0 ± 0%      41.0 ± 0%   +64.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       401 ± 0%   +93.72%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     4.00k ± 0%   +98.56%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     40.0k ± 0%   +99.16%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      400k ± 0%   +99.22%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     4.00M ± 0%   +99.22%  (p=0.000 n=10+10)
```

deque vs list - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseListQueue.txt
name        old time/op    new time/op    delta
/1-4          36.2ns ± 1%    70.1ns ± 1%   +93.40%  (p=0.000 n=9+10)
/10-4          368ns ± 1%     705ns ± 1%   +91.59%  (p=0.000 n=10+9)
/100-4        3.60µs ± 2%    7.01µs ± 1%   +94.64%  (p=0.000 n=8+10)
/1000-4       35.6µs ± 1%    69.3µs ± 1%   +94.97%  (p=0.000 n=10+9)
/10000-4       364µs ± 4%     695µs ± 0%   +91.10%  (p=0.000 n=10+9)
/100000-4     3.58ms ± 1%    6.93ms ± 1%   +93.45%  (p=0.000 n=9+9)
/1000000-4    35.8ms ± 1%    69.8ms ± 1%   +94.70%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=9+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

deque vs list - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableListQueue.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%    82.8ns ± 1%  +144.45%  (p=0.000 n=8+10)
/10-4          344ns ± 1%     825ns ± 3%  +139.99%  (p=0.000 n=10+9)
/100-4        3.34µs ± 1%    8.16µs ± 1%  +144.04%  (p=0.000 n=8+10)
/1000-4       33.9µs ± 4%    80.7µs ± 0%  +138.27%  (p=0.000 n=10+9)
/10000-4       348µs ± 2%     812µs ± 0%  +133.01%  (p=0.000 n=10+8)
/100000-4     3.47ms ± 1%    8.05ms ± 0%  +132.18%  (p=0.000 n=10+10)
/1000000-4    35.0ms ± 2%    81.3ms ± 1%  +132.26%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+8)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

#### deque vs list - LIFO stack
deque vs list - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillListStack.txt
name       old time/op    new time/op    delta
/1-4         3.66µs ± 1%    7.23µs ± 1%   +97.23%  (p=0.000 n=9+9)
/10-4        34.1µs ± 1%    70.6µs ± 1%  +107.06%  (p=0.000 n=9+9)
/100-4        335µs ± 1%     694µs ± 0%  +106.84%  (p=0.000 n=9+8)
/1000-4      3.35ms ± 3%    6.91ms ± 1%  +106.34%  (p=0.000 n=10+9)
/10000-4     36.0ms ± 1%    69.9ms ± 3%   +93.99%  (p=0.000 n=9+10)
/100000-4     400ms ± 1%    1449ms ± 6%  +262.42%  (p=0.000 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+8)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +299.99%  (p=0.000 n=10+10)
/10000-4     23.1MB ± 0%    64.0MB ± 0%  +176.74%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     640MB ± 0%  +165.66%  (p=0.000 n=10+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     2.00M ± 0%   +98.65%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     20.0M ± 0%   +98.47%  (p=0.000 n=10+9)
```

deque vs list - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullListStack.txt
name       old time/op    new time/op    delta
/1-4         3.60µs ± 1%    9.86µs ± 4%  +173.50%  (p=0.000 n=9+10)
/10-4        35.2µs ± 2%    97.9µs ± 6%  +177.87%  (p=0.000 n=10+10)
/100-4        339µs ± 5%     960µs ± 3%  +182.82%  (p=0.000 n=10+10)
/1000-4      3.35ms ± 2%    9.49ms ± 2%  +183.03%  (p=0.000 n=9+9)
/10000-4     35.7ms ± 6%   106.9ms ± 2%  +198.98%  (p=0.000 n=10+10)
/100000-4     388ms ± 2%    1383ms ± 9%  +256.05%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=9+9)
/10000-4     23.1MB ± 0%    64.0MB ± 0%  +176.74%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     640MB ± 0%  +165.66%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     2.00M ± 0%   +98.65%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     20.0M ± 0%   +98.47%  (p=0.000 n=10+10)
```

deque vs list - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseListStack.txt
name        old time/op    new time/op    delta
/1-4           196ns ± 1%     179ns ± 1%    -8.69%  (p=0.000 n=10+10)
/10-4          890ns ± 2%    1422ns ± 1%   +59.70%  (p=0.000 n=10+10)
/100-4        7.50µs ± 1%   13.74µs ± 1%   +83.14%  (p=0.000 n=9+10)
/1000-4       68.4µs ± 5%   137.1µs ± 0%  +100.27%  (p=0.000 n=10+10)
/10000-4       712µs ± 2%    1384µs ± 1%   +94.52%  (p=0.000 n=10+9)
/100000-4     7.68ms ± 2%   26.55ms ± 3%  +245.58%  (p=0.000 n=10+10)
/1000000-4    83.9ms ± 2%   251.7ms ±19%  +199.83%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%      176B ± 0%      ~     (all equal)
/10-4           592B ± 0%     1328B ± 0%  +124.32%  (p=0.000 n=10+10)
/100-4        6.08kB ± 0%   12.85kB ± 0%  +111.32%  (p=0.000 n=10+10)
/1000-4       41.2kB ± 0%   128.0kB ± 0%  +211.04%  (p=0.000 n=10+10)
/10000-4       403kB ± 0%    1280kB ± 0%  +218.00%  (p=0.000 n=9+10)
/100000-4     4.02MB ± 0%   12.80MB ± 0%  +218.38%  (p=0.000 n=10+10)
/1000000-4    40.2MB ± 0%   128.0MB ± 0%  +218.49%  (p=0.000 n=9+9)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%      ~     (all equal)
/10-4           24.0 ± 0%      41.0 ± 0%   +70.83%  (p=0.000 n=10+10)
/100-4           207 ± 0%       401 ± 0%   +93.72%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     4.00k ± 0%   +98.76%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     40.0k ± 0%   +99.18%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      400k ± 0%   +99.22%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     4.00M ± 0%   +99.22%  (p=0.000 n=10+10)
```

deque vs list - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreaseListStack.txt
name        old time/op    new time/op    delta
/1-4          36.6ns ± 2%    70.3ns ± 0%   +92.27%  (p=0.000 n=10+9)
/10-4          369ns ± 1%     705ns ± 0%   +90.67%  (p=0.000 n=8+10)
/100-4        3.60µs ± 2%    6.98µs ± 0%   +94.07%  (p=0.000 n=10+10)
/1000-4       36.0µs ± 1%    69.5µs ± 0%   +93.24%  (p=0.000 n=9+9)
/10000-4       362µs ± 5%     699µs ± 2%   +93.23%  (p=0.000 n=10+9)
/100000-4     3.61ms ± 4%    6.95ms ± 0%   +92.56%  (p=0.000 n=10+9)
/1000000-4    35.7ms ± 1%    69.8ms ± 1%   +95.21%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

deque vs list - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableListStack.txt
name        old time/op    new time/op    delta
/1-4          35.4ns ± 2%    97.8ns ± 2%  +176.08%  (p=0.000 n=10+8)
/10-4          362ns ± 4%     977ns ± 3%  +170.28%  (p=0.000 n=10+10)
/100-4        3.51µs ± 2%    9.69µs ± 3%  +175.88%  (p=0.000 n=10+10)
/1000-4       34.7µs ± 1%    97.4µs ± 5%  +180.63%  (p=0.000 n=10+10)
/10000-4       352µs ± 3%     974µs ± 2%  +176.88%  (p=0.000 n=10+9)
/100000-4     3.49ms ± 2%    9.70ms ± 3%  +178.35%  (p=0.000 n=9+10)
/1000000-4    35.0ms ± 2%    98.2ms ± 7%  +180.39%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

#### deque vs [CustomSliceQueue](testdata_test.go) - FIFO queue
deque vs CustomSliceQueue - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillSliceQueue.txt
name       old time/op    new time/op    delta
/1-4         3.72µs ± 1%    6.27µs ± 1%   +68.59%  (p=0.000 n=9+10)
/10-4        35.9µs ± 3%    47.6µs ± 0%   +32.55%  (p=0.000 n=10+9)
/100-4        351µs ± 3%     342µs ± 1%    -2.52%  (p=0.000 n=9+9)
/1000-4      3.43ms ± 2%    3.30ms ± 1%    -3.75%  (p=0.000 n=10+10)
/10000-4     35.1ms ± 4%    36.8ms ± 1%    +4.84%  (p=0.000 n=10+8)
/100000-4     389ms ± 1%     702ms ± 2%   +80.39%  (p=0.000 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    2.40kB ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    32.0kB ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     314kB ± 0%   +96.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    3.39MB ± 0%  +112.00%  (p=0.000 n=10+10)
/10000-4     23.1MB ± 0%    45.7MB ± 0%   +97.62%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     540MB ± 0%  +124.41%  (p=0.000 n=10+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.1k ± 0%    +1.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      100k ± 0%    +0.30%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.62%  (p=0.000 n=9+8)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.76%  (p=0.000 n=10+9)
```

deque vs CustomSliceQueue - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullSliceQueue.txt
name       old time/op    new time/op    delta
/1-4         3.53µs ± 2%    4.32µs ± 4%   +22.48%  (p=0.000 n=10+9)
/10-4        34.0µs ± 1%    40.4µs ± 0%   +18.71%  (p=0.000 n=9+9)
/100-4        332µs ± 1%     402µs ± 0%   +20.96%  (p=0.000 n=9+9)
/1000-4      3.29ms ± 1%    4.15ms ± 0%   +26.26%  (p=0.000 n=9+9)
/10000-4     35.3ms ± 1%    44.6ms ± 1%   +26.29%  (p=0.000 n=10+10)
/100000-4     388ms ± 1%     718ms ± 2%   +85.15%  (p=0.000 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    4.82kB ± 0%  +200.94%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    48.2kB ± 0%  +201.08%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     483kB ± 0%  +201.70%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    5.15MB ± 0%  +221.87%  (p=0.000 n=10+10)
/10000-4     23.1MB ± 0%    51.8MB ± 0%  +124.03%  (p=0.000 n=8+9)
/100000-4     241MB ± 0%     545MB ± 0%  +126.31%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    +0.03%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      100k ± 0%    +0.03%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.65%  (p=0.000 n=10+9)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.76%  (p=0.000 n=10+10)
```

deque vs CustomSliceQueue - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseSliceQueue.txt
name        old time/op    new time/op    delta
/1-4           195ns ± 1%     187ns ± 1%    -4.33%  (p=0.000 n=10+10)
/10-4         1.05µs ± 1%    1.16µs ± 0%   +10.52%  (p=0.000 n=9+9)
/100-4        7.35µs ± 1%    7.54µs ± 1%    +2.55%  (p=0.000 n=9+9)
/1000-4       67.2µs ± 1%    69.1µs ± 1%    +2.83%  (p=0.000 n=9+9)
/10000-4       688µs ± 2%     743µs ± 1%    +7.98%  (p=0.000 n=10+9)
/100000-4     7.70ms ± 1%   13.24ms ± 1%   +72.10%  (p=0.000 n=9+10)
/1000000-4    84.1ms ± 2%   144.2ms ± 3%   +71.47%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%       80B ± 0%   -54.55%  (p=0.000 n=10+10)
/10-4         1.10kB ± 0%    0.77kB ± 0%   -30.43%  (p=0.000 n=10+10)
/100-4        6.08kB ± 0%    6.66kB ± 0%    +9.47%  (p=0.000 n=10+10)
/1000-4       43.3kB ± 0%    78.1kB ± 0%   +80.47%  (p=0.000 n=10+10)
/10000-4       405kB ± 0%     964kB ± 0%  +138.24%  (p=0.000 n=10+10)
/100000-4     4.02MB ± 0%   11.10MB ± 0%  +176.15%  (p=0.000 n=10+10)
/1000000-4    40.2MB ± 0%   113.3MB ± 0%  +181.84%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%      ~     (all equal)
/10-4           25.0 ± 0%      29.0 ± 0%   +16.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       214 ± 0%    +3.38%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.02k ± 0%    +0.30%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.26%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.37%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs CustomSliceQueue - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseSliceQueue.txt
name        old time/op    new time/op    delta
/1-4          36.2ns ± 1%    60.2ns ± 1%   +66.15%  (p=0.000 n=9+9)
/10-4          368ns ± 1%     598ns ± 1%   +62.65%  (p=0.000 n=10+9)
/100-4        3.60µs ± 2%    5.90µs ± 1%   +63.79%  (p=0.000 n=8+10)
/1000-4       35.6µs ± 1%    58.9µs ± 1%   +65.55%  (p=0.000 n=10+9)
/10000-4       364µs ± 4%     589µs ± 1%   +61.76%  (p=0.000 n=10+9)
/100000-4     3.58ms ± 1%    5.87ms ± 0%   +63.75%  (p=0.000 n=9+8)
/1000000-4    35.8ms ± 1%    59.0ms ± 1%   +64.59%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     24.0B ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      240B ± 0%   +50.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    2.40kB ± 0%   +50.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    24.0kB ± 0%   +50.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     240kB ± 0%   +50.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    2.40MB ± 0%   +50.00%  (p=0.000 n=9+10)
/1000000-4    16.0MB ± 0%    24.0MB ± 0%   +50.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      2.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      20.0 ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     2.00M ± 0%  +100.00%  (p=0.000 n=10+10)
```

deque vs CustomSliceQueue - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableSliceQueue.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%    40.3ns ± 4%   +18.85%  (p=0.000 n=8+10)
/10-4          344ns ± 1%     404ns ± 3%   +17.55%  (p=0.000 n=10+9)
/100-4        3.34µs ± 1%    4.04µs ± 6%   +20.76%  (p=0.000 n=8+10)
/1000-4       33.9µs ± 4%    39.2µs ± 1%   +15.66%  (p=0.000 n=10+9)
/10000-4       348µs ± 2%     396µs ± 2%   +13.60%  (p=0.000 n=10+9)
/100000-4     3.47ms ± 1%    3.95ms ± 2%   +13.99%  (p=0.000 n=10+10)
/1000000-4    35.0ms ± 2%    39.1ms ± 1%   +11.77%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     48.0B ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      481B ± 0%  +200.62%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    4.82kB ± 0%  +200.94%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    48.2kB ± 0%  +200.96%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     482kB ± 0%  +200.97%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    4.82MB ± 0%  +200.97%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    48.2MB ± 0%  +200.97%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%      ~     (all equal)
/100-4           100 ± 0%       100 ± 0%      ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    +0.03%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      100k ± 0%    +0.03%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     1.00M ± 0%    +0.03%  (p=0.000 n=10+9)
```

#### deque vs CustomSliceQueue - LIFO stack
deque vs CustomSliceQueue - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillSliceStack.txt
name       old time/op    new time/op    delta
/1-4         3.66µs ± 1%    3.21µs ± 1%  -12.30%  (p=0.000 n=9+9)
/10-4        34.1µs ± 1%    31.3µs ± 1%   -8.35%  (p=0.000 n=9+9)
/100-4        335µs ± 1%     299µs ± 1%  -10.80%  (p=0.000 n=9+10)
/1000-4      3.35ms ± 3%    2.92ms ± 1%  -12.76%  (p=0.000 n=10+8)
/10000-4     36.0ms ± 1%    29.9ms ± 1%  -16.92%  (p=0.000 n=9+9)
/100000-4     400ms ± 1%     362ms ± 1%   -9.39%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   +0.00%  (p=0.000 n=10+9)
/10000-4     23.1MB ± 0%    16.0MB ± 0%  -30.78%  (p=0.000 n=10+6)
/100000-4     241MB ± 0%     161MB ± 0%  -33.02%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+9)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=10+10)
```

deque vs CustomSliceQueue - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullSliceStack.txt
name       old time/op    new time/op    delta
/1-4         3.60µs ± 1%    3.20µs ± 2%  -11.35%  (p=0.000 n=9+9)
/10-4        35.2µs ± 2%    30.6µs ± 2%  -13.18%  (p=0.000 n=10+10)
/100-4        339µs ± 5%     299µs ± 2%  -11.86%  (p=0.000 n=10+9)
/1000-4      3.35ms ± 2%    2.98ms ± 1%  -11.17%  (p=0.000 n=9+9)
/10000-4     35.7ms ± 6%    30.0ms ± 3%  -16.20%  (p=0.000 n=10+10)
/100000-4     388ms ± 2%     357ms ± 3%   -8.09%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   +0.00%  (p=0.000 n=9+10)
/10000-4     23.1MB ± 0%    16.0MB ± 0%  -30.81%  (p=0.000 n=10+9)
/100000-4     241MB ± 0%     160MB ± 0%  -33.58%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=10+10)
```

deque vs CustomSliceQueue - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseSliceStack.txt
name        old time/op    new time/op    delta
/1-4           196ns ± 1%     185ns ± 1%   -5.85%  (p=0.000 n=10+10)
/10-4          890ns ± 2%     950ns ± 2%   +6.65%  (p=0.000 n=10+9)
/100-4        7.50µs ± 1%    6.76µs ± 1%   -9.85%  (p=0.000 n=9+10)
/1000-4       68.4µs ± 5%    62.3µs ± 1%   -8.99%  (p=0.000 n=10+10)
/10000-4       712µs ± 2%     688µs ± 2%   -3.33%  (p=0.000 n=10+10)
/100000-4     7.68ms ± 2%   10.45ms ± 1%  +36.03%  (p=0.000 n=10+9)
/1000000-4    83.9ms ± 2%   111.8ms ± 2%  +33.16%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%       80B ± 0%  -54.55%  (p=0.000 n=10+10)
/10-4           592B ± 0%      592B ± 0%     ~     (all equal)
/100-4        6.08kB ± 0%    5.26kB ± 0%  -13.42%  (p=0.000 n=10+10)
/1000-4       41.2kB ± 0%    57.2kB ± 0%  +39.02%  (p=0.000 n=10+10)
/10000-4       403kB ± 0%     678kB ± 0%  +68.35%  (p=0.000 n=9+10)
/100000-4     4.02MB ± 0%    7.30MB ± 0%  +81.61%  (p=0.000 n=10+10)
/1000000-4    40.2MB ± 0%    73.7MB ± 0%  +83.33%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%     ~     (all equal)
/10-4           24.0 ± 0%      26.0 ± 0%   +8.33%  (p=0.000 n=10+10)
/100-4           207 ± 0%       209 ± 0%   +0.97%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.01k ± 0%     ~     (all equal)
/10000-4       20.1k ± 0%     20.0k ± 0%   -0.31%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%   -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%   -0.39%  (p=0.000 n=10+10)
```

deque vs CustomSliceQueue - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreaseSliceStack.txt
name        old time/op    new time/op    delta
/1-4          36.6ns ± 2%    32.8ns ± 2%  -10.38%  (p=0.000 n=10+10)
/10-4          369ns ± 1%     331ns ± 1%  -10.53%  (p=0.000 n=8+9)
/100-4        3.60µs ± 2%    3.22µs ± 1%  -10.60%  (p=0.000 n=10+9)
/1000-4       36.0µs ± 1%    32.3µs ± 2%  -10.11%  (p=0.000 n=9+10)
/10000-4       362µs ± 5%     320µs ± 1%  -11.50%  (p=0.000 n=10+9)
/100000-4     3.61ms ± 4%    3.23ms ± 1%  -10.58%  (p=0.000 n=10+10)
/1000000-4    35.7ms ± 1%    32.2ms ± 2%   -9.84%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%   -0.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%   -0.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

deque vs CustomSliceQueue - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableSliceStack.txt
name        old time/op    new time/op    delta
/1-4          35.4ns ± 2%    29.4ns ± 1%  -17.15%  (p=0.000 n=10+9)
/10-4          362ns ± 4%     298ns ± 1%  -17.49%  (p=0.000 n=10+10)
/100-4        3.51µs ± 2%    2.90µs ± 1%  -17.52%  (p=0.000 n=10+10)
/1000-4       34.7µs ± 1%    28.8µs ± 1%  -16.98%  (p=0.000 n=10+10)
/10000-4       352µs ± 3%     290µs ± 2%  -17.47%  (p=0.000 n=10+9)
/100000-4     3.49ms ± 2%    2.90ms ± 2%  -16.82%  (p=0.000 n=9+10)
/1000000-4    35.0ms ± 2%    28.9ms ± 1%  -17.63%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (p=0.087 n=10+10)
/100000-4     1.60MB ± 0%    1.60MB ± 0%   +0.00%  (p=0.000 n=10+9)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%   +0.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (p=0.087 n=10+10)
```

#### deque vs phf - FIFO queue
deque vs phf - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillPhfQueue.txt
name       old time/op    new time/op    delta
/1-4         3.72µs ± 1%    3.53µs ± 2%    -4.90%  (p=0.000 n=9+8)
/10-4        35.9µs ± 3%    52.5µs ± 4%   +46.44%  (p=0.000 n=10+8)
/100-4        351µs ± 3%     525µs ± 9%   +49.81%  (p=0.000 n=9+10)
/1000-4      3.43ms ± 2%    4.44ms ± 3%   +29.38%  (p=0.000 n=10+8)
/10000-4     35.1ms ± 4%    45.7ms ± 9%   +30.16%  (p=0.000 n=10+9)
/100000-4     389ms ± 1%     842ms ± 9%  +116.22%  (p=0.000 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    54.4kB ± 0%  +240.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     736kB ± 0%  +360.00%  (p=0.000 n=10+8)
/1000-4      1.60MB ± 0%    6.48MB ± 0%  +304.79%  (p=0.000 n=10+10)
/10000-4     23.1MB ± 0%    94.6MB ± 0%  +308.83%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     789MB ± 0%  +227.66%  (p=0.000 n=10+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.8k ± 0%    +8.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.40%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.46%  (p=0.000 n=9+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.74%  (p=0.000 n=10+9)
```

deque vs phf - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullPhfQueue.txt
name       old time/op    new time/op    delta
/1-4         3.53µs ± 2%    3.57µs ± 0%    +1.24%  (p=0.008 n=10+8)
/10-4        34.0µs ± 1%    35.2µs ± 1%    +3.32%  (p=0.000 n=9+10)
/100-4        332µs ± 1%     343µs ± 1%    +3.18%  (p=0.000 n=9+10)
/1000-4      3.29ms ± 1%    3.42ms ± 1%    +3.91%  (p=0.000 n=9+10)
/10000-4     35.3ms ± 1%    34.3ms ± 1%    -2.80%  (p=0.000 n=10+9)
/100000-4     388ms ± 1%     575ms ± 2%   +48.41%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%      ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%    +0.00%  (p=0.000 n=10+8)
/10000-4     23.1MB ± 0%    16.0MB ± 0%   -30.85%  (p=0.000 n=8+10)
/100000-4     241MB ± 0%     632MB ± 0%  +162.37%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%      ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%      ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.76%  (p=0.000 n=10+9)
```

deque vs phf - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreasePhfQueue.txt
name        old time/op    new time/op    delta
/1-4           195ns ± 1%     202ns ± 1%    +3.31%  (p=0.000 n=10+9)
/10-4         1.05µs ± 1%    1.17µs ± 0%   +12.22%  (p=0.000 n=9+9)
/100-4        7.35µs ± 1%    8.61µs ± 0%   +17.01%  (p=0.000 n=9+10)
/1000-4       67.2µs ± 1%    77.4µs ±10%   +15.18%  (p=0.000 n=9+10)
/10000-4       688µs ± 2%     867µs ± 2%   +26.02%  (p=0.000 n=10+8)
/100000-4     7.70ms ± 1%   10.49ms ± 2%   +36.25%  (p=0.000 n=9+10)
/1000000-4    84.1ms ± 2%   103.1ms ± 2%   +22.51%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%      128B ± 0%   -27.27%  (p=0.000 n=10+10)
/10-4         1.10kB ± 0%    0.99kB ± 0%   -10.14%  (p=0.000 n=10+10)
/100-4        6.08kB ± 0%    9.25kB ± 0%   +52.11%  (p=0.000 n=10+10)
/1000-4       43.3kB ± 0%    81.1kB ± 0%   +87.35%  (p=0.000 n=10+10)
/10000-4       405kB ± 0%    1106kB ± 0%  +173.42%  (p=0.000 n=10+9)
/100000-4     4.02MB ± 0%    9.49MB ± 0%  +136.08%  (p=0.000 n=10+10)
/1000000-4    40.2MB ± 0%    82.3MB ± 0%  +104.86%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%      ~     (all equal)
/10-4           25.0 ± 0%      27.0 ± 0%    +8.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       213 ± 0%    +2.90%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.02k ± 0%    +0.20%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.29%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs phf - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreasePhfQueue.txt
name        old time/op    new time/op    delta
/1-4          36.2ns ± 1%    34.9ns ± 1%  -3.64%  (p=0.000 n=9+9)
/10-4          368ns ± 1%     354ns ± 1%  -3.75%  (p=0.000 n=10+9)
/100-4        3.60µs ± 2%    3.45µs ± 1%  -4.24%  (p=0.000 n=8+10)
/1000-4       35.6µs ± 1%    34.3µs ± 1%  -3.53%  (p=0.000 n=10+10)
/10000-4       364µs ± 4%     344µs ± 1%  -5.54%  (p=0.000 n=10+9)
/100000-4     3.58ms ± 1%    3.44ms ± 1%  -3.97%  (p=0.000 n=9+10)
/1000000-4    35.8ms ± 1%    34.5ms ± 1%  -3.80%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (p=0.151 n=9+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (p=0.897 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

deque vs phf - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStablePhfQueue.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%    34.1ns ± 2%    ~     (p=0.179 n=8+10)
/10-4          344ns ± 1%     343ns ± 1%    ~     (p=0.811 n=10+10)
/100-4        3.34µs ± 1%    3.33µs ± 1%    ~     (p=0.153 n=8+10)
/1000-4       33.9µs ± 4%    33.3µs ± 1%    ~     (p=0.278 n=10+9)
/10000-4       348µs ± 2%     332µs ± 1%  -4.67%  (p=0.000 n=10+10)
/100000-4     3.47ms ± 1%    3.32ms ± 1%  -4.40%  (p=0.000 n=10+9)
/1000000-4    35.0ms ± 2%    33.1ms ± 1%  -5.33%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%  +0.00%  (p=0.006 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%  +0.00%  (p=0.001 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

#### deque vs phf - LIFO stack
deque vs phf - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillPhfStack.txt
name       old time/op    new time/op    delta
/1-4         3.66µs ± 1%    3.58µs ± 6%    -2.25%  (p=0.023 n=9+10)
/10-4        34.1µs ± 1%    58.2µs ±38%   +70.64%  (p=0.000 n=9+10)
/100-4        335µs ± 1%     532µs ±19%   +58.65%  (p=0.000 n=9+10)
/1000-4      3.35ms ± 3%    4.31ms ± 3%   +28.74%  (p=0.000 n=10+10)
/10000-4     36.0ms ± 1%    44.6ms ± 1%   +23.84%  (p=0.000 n=9+10)
/100000-4     400ms ± 1%     775ms ± 1%   +93.88%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    54.4kB ± 0%  +240.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     736kB ± 0%  +360.00%  (p=0.000 n=10+8)
/1000-4      1.60MB ± 0%    6.48MB ± 0%  +304.79%  (p=0.000 n=10+10)
/10000-4     23.1MB ± 0%    94.6MB ± 0%  +309.07%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     789MB ± 0%  +227.55%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.8k ± 0%    +8.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.40%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.46%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.74%  (p=0.000 n=10+10)
```

deque vs phf - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullPhfStack.txt
name       old time/op    new time/op    delta
/1-4         3.60µs ± 1%    3.47µs ± 1%    -3.62%  (p=0.000 n=9+9)
/10-4        35.2µs ± 2%    34.8µs ± 1%    -1.24%  (p=0.002 n=10+9)
/100-4        339µs ± 5%     338µs ± 0%      ~     (p=0.497 n=10+9)
/1000-4      3.35ms ± 2%    3.39ms ± 1%    +1.11%  (p=0.008 n=9+9)
/10000-4     35.7ms ± 6%    34.3ms ± 1%    -4.09%  (p=0.000 n=10+9)
/100000-4     388ms ± 2%     561ms ± 2%   +44.53%  (p=0.000 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%      ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%    +0.00%  (p=0.026 n=9+10)
/10000-4     23.1MB ± 0%    16.0MB ± 0%   -30.81%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     632MB ± 0%  +162.28%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%      ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%      ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.76%  (p=0.000 n=10+10)
```

deque vs phf - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreasePhfStack.txt
name        old time/op    new time/op    delta
/1-4           196ns ± 1%     201ns ± 1%    +2.17%  (p=0.000 n=10+10)
/10-4          890ns ± 2%    1156ns ± 0%   +29.77%  (p=0.000 n=10+9)
/100-4        7.50µs ± 1%    8.62µs ± 1%   +14.86%  (p=0.000 n=9+9)
/1000-4       68.4µs ± 5%    75.3µs ± 1%   +10.00%  (p=0.000 n=10+9)
/10000-4       712µs ± 2%     853µs ± 1%   +19.81%  (p=0.000 n=10+10)
/100000-4     7.68ms ± 2%   10.26ms ± 0%   +33.55%  (p=0.000 n=10+8)
/1000000-4    83.9ms ± 2%    99.9ms ± 2%   +19.01%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%      128B ± 0%   -27.27%  (p=0.000 n=10+10)
/10-4           592B ± 0%      992B ± 0%   +67.57%  (p=0.000 n=10+10)
/100-4        6.08kB ± 0%    9.25kB ± 0%   +52.11%  (p=0.000 n=10+10)
/1000-4       41.2kB ± 0%    81.1kB ± 0%   +96.89%  (p=0.000 n=10+10)
/10000-4       403kB ± 0%    1106kB ± 0%  +174.85%  (p=0.000 n=9+10)
/100000-4     4.02MB ± 0%    9.49MB ± 0%  +136.08%  (p=0.000 n=10+10)
/1000000-4    40.2MB ± 0%    82.3MB ± 0%  +104.86%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%      ~     (all equal)
/10-4           24.0 ± 0%      27.0 ± 0%   +12.50%  (p=0.000 n=10+10)
/100-4           207 ± 0%       213 ± 0%    +2.90%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.02k ± 0%    +0.30%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.28%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs phf - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreasePhfStack.txt
name        old time/op    new time/op    delta
/1-4          36.6ns ± 2%    34.1ns ± 2%  -6.85%  (p=0.000 n=10+9)
/10-4          369ns ± 1%     345ns ± 1%  -6.75%  (p=0.000 n=8+9)
/100-4        3.60µs ± 2%    3.35µs ± 1%  -7.02%  (p=0.000 n=10+10)
/1000-4       36.0µs ± 1%    33.4µs ± 1%  -7.20%  (p=0.000 n=9+10)
/10000-4       362µs ± 5%     335µs ± 1%  -7.44%  (p=0.000 n=10+10)
/100000-4     3.61ms ± 4%    3.35ms ± 1%  -7.32%  (p=0.000 n=10+9)
/1000000-4    35.7ms ± 1%    33.5ms ± 1%  -6.32%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (p=0.157 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%  +0.00%  (p=0.041 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

deque vs phf - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStablePhfStack.txt
name        old time/op    new time/op    delta
/1-4          35.4ns ± 2%    33.4ns ± 1%  -5.82%  (p=0.000 n=10+9)
/10-4          362ns ± 4%     339ns ± 1%  -6.32%  (p=0.000 n=10+10)
/100-4        3.51µs ± 2%    3.30µs ± 0%  -6.09%  (p=0.000 n=10+9)
/1000-4       34.7µs ± 1%    32.8µs ± 1%  -5.42%  (p=0.000 n=10+9)
/10000-4       352µs ± 3%     330µs ± 1%  -6.26%  (p=0.000 n=10+8)
/100000-4     3.49ms ± 2%    3.29ms ± 1%  -5.64%  (p=0.000 n=9+10)
/1000000-4    35.0ms ± 2%    32.9ms ± 1%  -6.00%  (p=0.000 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%  +0.00%  (p=0.001 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%  +0.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

#### deque vs gammazero - FIFO queue
deque vs gammazero - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillGammazeroQueue.txt
name       old time/op    new time/op    delta
/1-4         3.72µs ± 1%    3.55µs ± 1%    -4.61%  (p=0.000 n=9+10)
/10-4        35.9µs ± 3%    34.5µs ± 1%    -3.77%  (p=0.000 n=10+9)
/100-4        351µs ± 3%     465µs ± 1%   +32.71%  (p=0.000 n=9+9)
/1000-4      3.43ms ± 2%    4.21ms ± 2%   +22.74%  (p=0.000 n=10+9)
/10000-4     35.1ms ± 4%    44.6ms ± 1%   +27.11%  (p=0.000 n=10+9)
/100000-4     389ms ± 1%     772ms ± 2%   +98.31%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     698kB ± 0%  +336.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.44MB ± 0%  +302.39%  (p=0.000 n=10+10)
/10000-4     23.1MB ± 0%    94.6MB ± 0%  +308.67%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     789MB ± 0%  +227.64%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.6k ± 0%    +6.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.20%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.48%  (p=0.000 n=9+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.74%  (p=0.000 n=10+10)
```

deque vs gammazero - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullGammazeroQueue.txt
name       old time/op    new time/op    delta
/1-4         3.53µs ± 2%    3.55µs ± 1%    +0.75%  (p=0.037 n=10+10)
/10-4        34.0µs ± 1%    35.3µs ± 1%    +3.66%  (p=0.000 n=9+10)
/100-4        332µs ± 1%     352µs ± 6%    +6.03%  (p=0.000 n=9+10)
/1000-4      3.29ms ± 1%    3.44ms ± 1%    +4.57%  (p=0.000 n=9+10)
/10000-4     35.3ms ± 1%    34.5ms ± 0%    -2.31%  (p=0.000 n=10+9)
/100000-4     388ms ± 1%     579ms ± 1%   +49.44%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%      ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%    +0.00%  (p=0.000 n=10+8)
/10000-4     23.1MB ± 0%    16.0MB ± 0%   -30.85%  (p=0.000 n=8+10)
/100000-4     241MB ± 0%     632MB ± 0%  +162.37%  (p=0.000 n=10+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%      ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%      ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.76%  (p=0.000 n=10+9)
```

deque vs gammazero - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4           195ns ± 1%     189ns ± 1%    -3.35%  (p=0.000 n=10+10)
/10-4         1.05µs ± 1%    0.78µs ± 1%   -25.28%  (p=0.000 n=9+10)
/100-4        7.35µs ± 1%    8.12µs ± 0%   +10.39%  (p=0.000 n=9+9)
/1000-4       67.2µs ± 1%    73.7µs ± 0%    +9.62%  (p=0.000 n=9+10)
/10000-4       688µs ± 2%     827µs ± 1%   +20.18%  (p=0.000 n=10+10)
/100000-4     7.70ms ± 1%   10.37ms ± 2%   +34.78%  (p=0.000 n=9+10)
/1000000-4    84.1ms ± 2%   102.0ms ± 3%   +21.30%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%      336B ± 0%   +90.91%  (p=0.000 n=10+10)
/10-4         1.10kB ± 0%    0.62kB ± 0%   -43.48%  (p=0.000 n=10+10)
/100-4        6.08kB ± 0%    8.88kB ± 0%   +46.05%  (p=0.000 n=10+10)
/1000-4       43.3kB ± 0%    80.7kB ± 0%   +86.50%  (p=0.000 n=10+10)
/10000-4       405kB ± 0%    1106kB ± 0%  +173.33%  (p=0.000 n=10+6)
/100000-4     4.02MB ± 0%    9.49MB ± 0%  +136.08%  (p=0.000 n=10+10)
/1000000-4    40.2MB ± 0%    82.3MB ± 0%  +104.86%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      4.00 ± 0%   -20.00%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      22.0 ± 0%   -12.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       208 ± 0%    +0.48%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.01k ± 0%    -0.05%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.31%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs gammazero - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4          36.2ns ± 1%    34.4ns ± 1%  -5.16%  (p=0.000 n=9+10)
/10-4          368ns ± 1%     347ns ± 0%  -5.57%  (p=0.000 n=10+9)
/100-4        3.60µs ± 2%    3.40µs ± 1%  -5.61%  (p=0.000 n=8+10)
/1000-4       35.6µs ± 1%    34.0µs ± 1%  -4.48%  (p=0.000 n=10+9)
/10000-4       364µs ± 4%     338µs ± 1%  -7.04%  (p=0.000 n=10+10)
/100000-4     3.58ms ± 1%    3.39ms ± 1%  -5.24%  (p=0.000 n=9+9)
/1000000-4    35.8ms ± 1%    34.0ms ± 1%  -5.15%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%  +0.00%  (p=0.001 n=9+9)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (p=0.341 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

deque vs gammazero - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%    33.7ns ± 1%    ~     (p=0.063 n=8+9)
/10-4          344ns ± 1%     341ns ± 1%  -0.81%  (p=0.003 n=10+9)
/100-4        3.34µs ± 1%    3.32µs ± 1%    ~     (p=0.096 n=8+10)
/1000-4       33.9µs ± 4%    33.1µs ± 1%  -2.18%  (p=0.012 n=10+10)
/10000-4       348µs ± 2%     331µs ± 1%  -4.88%  (p=0.000 n=10+9)
/100000-4     3.47ms ± 1%    3.30ms ± 1%  -4.91%  (p=0.000 n=10+9)
/1000000-4    35.0ms ± 2%    33.0ms ± 0%  -5.84%  (p=0.000 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%  +0.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%  +0.00%  (p=0.000 n=10+8)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

#### deque vs gammazero - LIFO stack
deque vs gammazero - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillGammazeroStack.txt
name       old time/op    new time/op    delta
/1-4         3.66µs ± 1%    3.55µs ± 1%    -3.15%  (p=0.000 n=9+10)
/10-4        34.1µs ± 1%    34.4µs ± 1%    +0.95%  (p=0.006 n=9+9)
/100-4        335µs ± 1%     463µs ± 0%   +37.99%  (p=0.000 n=9+8)
/1000-4      3.35ms ± 3%    4.21ms ± 1%   +25.60%  (p=0.000 n=10+10)
/10000-4     36.0ms ± 1%    44.7ms ± 1%   +24.13%  (p=0.000 n=9+10)
/100000-4     400ms ± 1%     777ms ± 2%   +94.25%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     698kB ± 0%  +336.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.44MB ± 0%  +302.39%  (p=0.000 n=10+8)
/10000-4     23.1MB ± 0%    94.6MB ± 0%  +308.90%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     789MB ± 0%  +227.54%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.6k ± 0%    +6.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.20%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.48%  (p=0.000 n=10+8)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.74%  (p=0.000 n=10+10)
```

deque vs gammazero - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullGammazeroStack.txt
name       old time/op    new time/op    delta
/1-4         3.60µs ± 1%    3.47µs ± 1%    -3.66%  (p=0.000 n=9+8)
/10-4        35.2µs ± 2%    34.8µs ± 0%    -1.23%  (p=0.001 n=10+9)
/100-4        339µs ± 5%     339µs ± 1%      ~     (p=0.515 n=10+8)
/1000-4      3.35ms ± 2%    3.39ms ± 1%    +0.96%  (p=0.010 n=9+10)
/10000-4     35.7ms ± 6%    34.4ms ± 1%    -3.79%  (p=0.000 n=10+9)
/100000-4     388ms ± 2%     568ms ± 2%   +46.35%  (p=0.000 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%      ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%    +0.00%  (p=0.000 n=9+10)
/10000-4     23.1MB ± 0%    16.0MB ± 0%   -30.81%  (p=0.000 n=10+9)
/100000-4     241MB ± 0%     632MB ± 0%  +162.28%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%      ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%      ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.76%  (p=0.000 n=10+10)
```

deque vs gammazero - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4           196ns ± 1%     189ns ± 0%    -3.82%  (p=0.000 n=10+9)
/10-4          890ns ± 2%     784ns ± 1%   -11.92%  (p=0.000 n=10+9)
/100-4        7.50µs ± 1%    8.10µs ± 1%    +7.97%  (p=0.000 n=9+10)
/1000-4       68.4µs ± 5%    73.4µs ± 0%    +7.26%  (p=0.000 n=10+8)
/10000-4       712µs ± 2%     819µs ± 1%   +15.02%  (p=0.000 n=10+10)
/100000-4     7.68ms ± 2%   10.57ms ± 3%   +37.66%  (p=0.000 n=10+10)
/1000000-4    83.9ms ± 2%   102.8ms ± 2%   +22.46%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%      336B ± 0%   +90.91%  (p=0.000 n=10+10)
/10-4           592B ± 0%      624B ± 0%    +5.41%  (p=0.000 n=10+10)
/100-4        6.08kB ± 0%    8.88kB ± 0%   +46.05%  (p=0.000 n=10+10)
/1000-4       41.2kB ± 0%    80.7kB ± 0%   +96.00%  (p=0.000 n=10+10)
/10000-4       403kB ± 0%    1106kB ± 0%  +174.76%  (p=0.000 n=9+9)
/100000-4     4.02MB ± 0%    9.49MB ± 0%  +136.08%  (p=0.000 n=10+10)
/1000000-4    40.2MB ± 0%    82.3MB ± 0%  +104.86%  (p=0.000 n=9+9)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      4.00 ± 0%   -20.00%  (p=0.000 n=10+10)
/10-4           24.0 ± 0%      22.0 ± 0%    -8.33%  (p=0.000 n=10+10)
/100-4           207 ± 0%       208 ± 0%    +0.48%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.01k ± 0%    +0.05%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.30%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs gammazero - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreaseGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4          36.6ns ± 2%    34.1ns ± 1%  -6.65%  (p=0.000 n=10+9)
/10-4          369ns ± 1%     346ns ± 0%  -6.40%  (p=0.000 n=8+10)
/100-4        3.60µs ± 2%    3.38µs ± 1%  -6.08%  (p=0.000 n=10+10)
/1000-4       36.0µs ± 1%    33.6µs ± 0%  -6.53%  (p=0.000 n=9+9)
/10000-4       362µs ± 5%     338µs ± 0%  -6.67%  (p=0.000 n=10+9)
/100000-4     3.61ms ± 4%    3.36ms ± 1%  -6.84%  (p=0.000 n=10+9)
/1000000-4    35.7ms ± 1%    33.8ms ± 1%  -5.38%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (p=0.346 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (p=0.183 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

deque vs gammazero - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4          35.4ns ± 2%    33.0ns ± 1%  -6.84%  (p=0.000 n=10+9)
/10-4          362ns ± 4%     334ns ± 0%  -7.56%  (p=0.000 n=10+8)
/100-4        3.51µs ± 2%    3.26µs ± 1%  -7.15%  (p=0.000 n=10+10)
/1000-4       34.7µs ± 1%    32.4µs ± 0%  -6.68%  (p=0.000 n=10+9)
/10000-4       352µs ± 3%     327µs ± 1%  -7.09%  (p=0.000 n=10+9)
/100000-4     3.49ms ± 2%    3.26ms ± 2%  -6.62%  (p=0.000 n=9+9)
/1000000-4    35.0ms ± 2%    32.5ms ± 1%  -7.35%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%  +0.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%  +0.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

#### deque vs Gostl - FIFO queue
deque vs Gostl - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillGostlQueue.txt
name       old time/op    new time/op    delta
/1-4         3.72µs ± 1%   19.76µs ± 1%  +431.84%  (p=0.000 n=9+10)
/10-4        35.9µs ± 3%    97.8µs ± 1%  +172.57%  (p=0.000 n=10+10)
/100-4        351µs ± 3%     870µs ± 1%  +148.09%  (p=0.000 n=9+10)
/1000-4      3.43ms ± 2%    8.71ms ± 0%  +153.65%  (p=0.000 n=10+9)
/10000-4     35.1ms ± 4%    88.0ms ± 2%  +150.52%  (p=0.000 n=10+10)
/100000-4     389ms ± 1%     899ms ± 1%  +130.81%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    4.00kB ± 0%  +150.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    18.4kB ± 0%   +15.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     162kB ± 0%    +1.50%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    2.51MB ± 0%   +56.95%  (p=0.000 n=10+10)
/10000-4     23.1MB ± 0%    25.0MB ± 0%    +7.97%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     248MB ± 0%    +3.17%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       300 ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.2k ± 0%    +2.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      103k ± 0%    +3.10%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.02M ± 0%    +1.27%  (p=0.000 n=9+10)
/100000-4     10.1M ± 0%     10.2M ± 0%    +0.85%  (p=0.000 n=10+10)
```

deque vs Gostl - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullGostlQueue.txt
name       old time/op    new time/op    delta
/1-4         3.53µs ± 2%    8.73µs ± 3%  +147.52%  (p=0.000 n=10+9)
/10-4        34.0µs ± 1%    85.8µs ± 1%  +152.05%  (p=0.000 n=9+8)
/100-4        332µs ± 1%     841µs ± 1%  +152.97%  (p=0.000 n=9+8)
/1000-4      3.29ms ± 1%    8.35ms ± 1%  +154.02%  (p=0.000 n=9+9)
/10000-4     35.3ms ± 1%    87.6ms ± 2%  +148.30%  (p=0.000 n=10+9)
/100000-4     388ms ± 1%     896ms ± 0%  +131.19%  (p=0.000 n=9+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%      ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%      ~     (p=0.325 n=10+10)
/10000-4     23.1MB ± 0%    22.5MB ± 0%    -2.65%  (p=0.000 n=8+10)
/100000-4     241MB ± 0%     246MB ± 0%    +2.12%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%      ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%      ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%    +0.50%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.2M ± 0%    +0.77%  (p=0.000 n=10+10)
```

deque vs Gostl - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseGostlQueue.txt
name        old time/op    new time/op    delta
/1-4           195ns ± 1%     662ns ± 1%  +239.15%  (p=0.000 n=10+10)
/10-4         1.05µs ± 1%    2.20µs ± 1%  +110.66%  (p=0.000 n=9+10)
/100-4        7.35µs ± 1%   16.87µs ± 1%  +129.40%  (p=0.000 n=9+10)
/1000-4       67.2µs ± 1%   168.6µs ± 1%  +150.67%  (p=0.000 n=9+10)
/10000-4       688µs ± 2%    1688µs ± 0%  +145.29%  (p=0.000 n=10+9)
/100000-4     7.70ms ± 1%   18.00ms ± 1%  +133.88%  (p=0.000 n=9+10)
/1000000-4    84.1ms ± 2%   183.8ms ± 2%  +118.49%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%     1232B ± 0%  +600.00%  (p=0.000 n=10+10)
/10-4         1.10kB ± 0%    1.52kB ± 0%   +37.68%  (p=0.000 n=10+10)
/100-4        6.08kB ± 0%    4.40kB ± 0%   -27.63%  (p=0.000 n=10+10)
/1000-4       43.3kB ± 0%    42.4kB ± 0%    -2.03%  (p=0.000 n=10+10)
/10000-4       405kB ± 0%     410kB ± 0%    +1.34%  (p=0.000 n=10+8)
/100000-4     4.02MB ± 0%    4.08MB ± 0%    +1.61%  (p=0.000 n=10+10)
/1000000-4    40.2MB ± 0%    40.8MB ± 0%    +1.61%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%     10.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      28.0 ± 0%   +12.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       208 ± 0%    +0.48%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.04k ± 0%    +1.19%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.2k ± 0%    +0.58%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      202k ± 0%    +0.42%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.02M ± 0%    +0.39%  (p=0.000 n=10+9)
```

deque vs Gostl - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseGostlQueue.txt
name        old time/op    new time/op    delta
/1-4          36.2ns ± 1%   199.9ns ± 0%  +451.81%  (p=0.000 n=9+10)
/10-4          368ns ± 1%    2003ns ± 1%  +444.44%  (p=0.000 n=10+10)
/100-4        3.60µs ± 2%   19.86µs ± 0%  +451.48%  (p=0.000 n=8+8)
/1000-4       35.6µs ± 1%   198.7µs ± 0%  +458.81%  (p=0.000 n=10+10)
/10000-4       364µs ± 4%    1986µs ± 1%  +445.70%  (p=0.000 n=10+10)
/100000-4     3.58ms ± 1%   19.87ms ± 1%  +454.68%  (p=0.000 n=9+10)
/1000000-4    35.8ms ± 1%   198.5ms ± 1%  +453.97%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     40.0B ± 0%  +150.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      400B ± 0%  +150.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    4.00kB ± 0%  +150.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    40.0kB ± 0%  +150.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     400kB ± 0%  +150.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    4.00MB ± 0%  +150.00%  (p=0.000 n=9+10)
/1000000-4    16.0MB ± 0%    40.0MB ± 0%  +150.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      3.00 ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      30.0 ± 0%  +200.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       300 ± 0%  +200.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     3.00k ± 0%  +200.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     30.0k ± 0%  +200.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      300k ± 0%  +200.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     3.00M ± 0%  +200.00%  (p=0.000 n=10+10)
```

deque vs Gostl - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableGostlQueue.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%  1243.0ns ± 1%  +3569.37%  (p=0.000 n=8+9)
/10-4          344ns ± 1%    7496ns ± 0%  +2081.29%  (p=0.000 n=10+8)
/100-4        3.34µs ± 1%   67.76µs ± 1%  +1926.09%  (p=0.000 n=8+10)
/1000-4       33.9µs ± 4%   682.5µs ± 0%  +1914.71%  (p=0.000 n=10+9)
/10000-4       348µs ± 2%    6884µs ± 2%  +1876.43%  (p=0.000 n=10+10)
/100000-4     3.47ms ± 1%   71.43ms ± 1%  +1959.40%  (p=0.000 n=10+10)
/1000000-4    35.0ms ± 2%   723.3ms ± 1%  +1966.27%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%   1352.0B ± 0%  +8350.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%     2600B ± 0%  +1525.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%   14.84kB ± 0%   +827.50%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%   155.6kB ± 0%   +872.50%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%    1542kB ± 0%   +863.85%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%   15.37MB ± 0%   +860.75%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%   153.7MB ± 0%   +860.46%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%     18.00 ± 0%  +1700.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%     101.0 ± 0%   +910.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       911 ± 0%   +811.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     9.07k ± 0%   +807.20%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     90.4k ± 0%   +804.04%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      903k ± 0%   +803.26%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     9.03M ± 0%   +803.14%  (p=0.000 n=10+10)
```

#### deque vs Gostl - LIFO stack
deque vs Gostl - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillGostlStack.txt
name       old time/op    new time/op    delta
/1-4         3.66µs ± 1%   20.10µs ± 1%  +448.44%  (p=0.000 n=9+9)
/10-4        34.1µs ± 1%   107.5µs ± 1%  +215.03%  (p=0.000 n=9+9)
/100-4        335µs ± 1%     959µs ± 0%  +185.84%  (p=0.000 n=9+9)
/1000-4      3.35ms ± 3%    9.88ms ± 1%  +195.03%  (p=0.000 n=10+10)
/10000-4     36.0ms ± 1%    98.5ms ± 1%  +173.41%  (p=0.000 n=9+9)
/100000-4     400ms ± 1%    1022ms ± 1%  +155.54%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    4.00kB ± 0%  +150.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    18.4kB ± 0%   +15.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     162kB ± 0%    +1.50%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    2.51MB ± 0%   +56.95%  (p=0.000 n=10+10)
/10000-4     23.1MB ± 0%    25.0MB ± 0%    +8.03%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     248MB ± 0%    +3.14%  (p=0.000 n=10+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       300 ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.2k ± 0%    +2.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      103k ± 0%    +3.10%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.02M ± 0%    +1.27%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.2M ± 0%    +0.85%  (p=0.000 n=10+9)
```

deque vs Gostl - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullGostlStack.txt
name       old time/op    new time/op    delta
/1-4         3.60µs ± 1%    9.91µs ± 1%  +174.87%  (p=0.000 n=9+10)
/10-4        35.2µs ± 2%    96.6µs ± 1%  +174.09%  (p=0.000 n=10+8)
/100-4        339µs ± 5%     955µs ± 1%  +181.41%  (p=0.000 n=10+9)
/1000-4      3.35ms ± 2%    9.48ms ± 1%  +182.81%  (p=0.000 n=9+9)
/10000-4     35.7ms ± 6%    98.4ms ± 1%  +175.26%  (p=0.000 n=10+9)
/100000-4     388ms ± 2%    1020ms ± 0%  +162.63%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%      ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%      ~     (p=0.111 n=9+8)
/10000-4     23.1MB ± 0%    22.5MB ± 0%    -2.59%  (p=0.000 n=10+9)
/100000-4     241MB ± 0%     246MB ± 0%    +2.08%  (p=0.000 n=10+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%      ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%      ~     (all equal)
/10000-4      1.01M ± 0%     1.01M ± 0%    +0.50%  (p=0.000 n=10+9)
/100000-4     10.1M ± 0%     10.2M ± 0%    +0.77%  (p=0.000 n=10+10)
```

deque vs Gostl - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseGostlStack.txt
name        old time/op    new time/op    delta
/1-4           196ns ± 1%     680ns ± 0%  +245.93%  (p=0.000 n=10+9)
/10-4          890ns ± 2%    2444ns ± 0%  +174.50%  (p=0.000 n=10+10)
/100-4        7.50µs ± 1%   19.39µs ± 0%  +158.44%  (p=0.000 n=9+8)
/1000-4       68.4µs ± 5%   193.4µs ± 1%  +182.51%  (p=0.000 n=10+9)
/10000-4       712µs ± 2%    1936µs ± 1%  +171.97%  (p=0.000 n=10+10)
/100000-4     7.68ms ± 2%   20.30ms ± 1%  +164.31%  (p=0.000 n=10+10)
/1000000-4    83.9ms ± 2%   206.4ms ± 1%  +145.90%  (p=0.000 n=9+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%     1232B ± 0%  +600.00%  (p=0.000 n=10+10)
/10-4           592B ± 0%     1520B ± 0%  +156.76%  (p=0.000 n=10+10)
/100-4        6.08kB ± 0%    4.40kB ± 0%   -27.63%  (p=0.000 n=10+10)
/1000-4       41.2kB ± 0%    41.4kB ± 0%    +0.45%  (p=0.000 n=10+10)
/10000-4       403kB ± 0%     410kB ± 0%    +1.88%  (p=0.000 n=9+10)
/100000-4     4.02MB ± 0%    4.09MB ± 0%    +1.61%  (p=0.000 n=10+9)
/1000000-4    40.2MB ± 0%    40.8MB ± 0%    +1.61%  (p=0.000 n=9+9)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%     10.00 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           24.0 ± 0%      28.0 ± 0%   +16.67%  (p=0.000 n=10+10)
/100-4           207 ± 0%       208 ± 0%    +0.48%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.04k ± 0%    +1.24%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.2k ± 0%    +0.60%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      202k ± 0%    +0.42%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.02M ± 0%    +0.39%  (p=0.000 n=10+9)
```

deque vs Gostl - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreaseGostlStack.txt
name        old time/op    new time/op    delta
/1-4          36.6ns ± 2%   204.7ns ± 1%  +459.80%  (p=0.000 n=10+10)
/10-4          369ns ± 1%    2052ns ± 0%  +455.50%  (p=0.000 n=8+10)
/100-4        3.60µs ± 2%   20.44µs ± 1%  +468.15%  (p=0.000 n=10+10)
/1000-4       36.0µs ± 1%   204.2µs ± 1%  +467.77%  (p=0.000 n=9+10)
/10000-4       362µs ± 5%    2043µs ± 1%  +464.89%  (p=0.000 n=10+9)
/100000-4     3.61ms ± 4%   20.45ms ± 1%  +466.41%  (p=0.000 n=10+9)
/1000000-4    35.7ms ± 1%   203.3ms ± 1%  +468.72%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     40.0B ± 0%  +150.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      400B ± 0%  +150.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    4.00kB ± 0%  +150.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    40.0kB ± 0%  +150.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     400kB ± 0%  +150.00%  (p=0.000 n=10+8)
/100000-4     1.60MB ± 0%    4.00MB ± 0%  +150.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    40.0MB ± 0%  +150.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      3.00 ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      30.0 ± 0%  +200.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       300 ± 0%  +200.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     3.00k ± 0%  +200.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     30.0k ± 0%  +200.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      300k ± 0%  +200.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     3.00M ± 0%  +200.00%  (p=0.000 n=10+9)
```

deque vs Gostl - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableGostlStack.txt
name        old time/op    new time/op    delta
/1-4          35.4ns ± 2%    97.2ns ± 1%  +174.32%  (p=0.000 n=10+9)
/10-4          362ns ± 4%     968ns ± 1%  +167.71%  (p=0.000 n=10+10)
/100-4        3.51µs ± 2%    9.67µs ± 1%  +175.42%  (p=0.000 n=10+10)
/1000-4       34.7µs ± 1%    96.3µs ± 1%  +177.19%  (p=0.000 n=10+9)
/10000-4       352µs ± 3%     963µs ± 1%  +173.81%  (p=0.000 n=10+10)
/100000-4     3.49ms ± 2%    9.59ms ± 1%  +175.15%  (p=0.000 n=9+10)
/1000000-4    35.0ms ± 2%    96.2ms ± 1%  +174.51%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%      ~     (all equal)
/10-4           160B ± 0%      160B ± 0%      ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%      ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%      ~     (p=0.704 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%      ~     (p=0.669 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%      ~     (all equal)
/100-4           100 ± 0%       100 ± 0%      ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%      ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%      ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%      ~     (all equal)
```

#### deque vs cookiejar - FIFO queue
deque vs cookiejar - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillCookiejarQueue.txt
name       old time/op    new time/op    delta
/1-4         3.72µs ± 1%    3.48µs ± 1%   -6.29%  (p=0.000 n=9+10)
/10-4        35.9µs ± 3%    33.3µs ± 1%   -7.25%  (p=0.000 n=10+9)
/100-4        351µs ± 3%     322µs ± 1%   -8.11%  (p=0.000 n=9+9)
/1000-4      3.43ms ± 2%    3.19ms ± 1%   -7.02%  (p=0.000 n=10+10)
/10000-4     35.1ms ± 4%    32.1ms ± 0%   -8.56%  (p=0.000 n=10+10)
/100000-4     389ms ± 1%     356ms ± 1%   -8.43%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%   +0.02%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     160kB ± 0%   +0.02%  (p=0.000 n=10+9)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   +0.02%  (p=0.000 n=10+10)
/10000-4     23.1MB ± 0%    16.0MB ± 0%  -30.82%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     161MB ± 0%  -33.33%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=9+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=10+10)
```

deque vs cookiejar - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullCookiejarQueue.txt
name       old time/op    new time/op    delta
/1-4         3.53µs ± 2%    3.25µs ± 0%   -7.97%  (p=0.000 n=10+9)
/10-4        34.0µs ± 1%    32.6µs ± 2%   -4.18%  (p=0.000 n=9+10)
/100-4        332µs ± 1%     317µs ± 1%   -4.71%  (p=0.000 n=9+9)
/1000-4      3.29ms ± 1%    3.16ms ± 0%   -3.88%  (p=0.000 n=9+10)
/10000-4     35.3ms ± 1%    32.6ms ± 0%   -7.60%  (p=0.000 n=10+9)
/100000-4     388ms ± 1%     351ms ± 2%   -9.36%  (p=0.000 n=9+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   +0.00%  (p=0.010 n=10+10)
/10000-4     23.1MB ± 0%    16.0MB ± 0%  -30.85%  (p=0.000 n=8+10)
/100000-4     241MB ± 0%     160MB ± 0%  -33.56%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=10+10)
```

deque vs cookiejar - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4           195ns ± 1%    9482ns ±11%   +4755.03%  (p=0.000 n=10+8)
/10-4         1.05µs ± 1%    9.85µs ± 1%    +842.50%  (p=0.000 n=9+7)
/100-4        7.35µs ± 1%   14.96µs ± 2%    +103.40%  (p=0.000 n=9+9)
/1000-4       67.2µs ± 1%    69.9µs ± 2%      +3.94%  (p=0.000 n=9+10)
/10000-4       688µs ± 2%     675µs ± 6%        ~     (p=0.143 n=10+10)
/100000-4     7.70ms ± 1%    7.87ms ± 1%      +2.23%  (p=0.000 n=9+8)
/1000000-4    84.1ms ± 2%    78.9ms ± 2%      -6.18%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%    65704B ± 0%  +37231.82%  (p=0.000 n=10+10)
/10-4         1.10kB ± 0%   65.99kB ± 0%   +5877.54%  (p=0.000 n=10+10)
/100-4        6.08kB ± 0%   68.87kB ± 0%   +1032.76%  (p=0.000 n=10+10)
/1000-4       43.3kB ± 0%    97.7kB ± 0%    +125.76%  (p=0.000 n=10+10)
/10000-4       405kB ± 0%     583kB ± 0%     +43.96%  (p=0.000 n=10+10)
/100000-4     4.02MB ± 0%    4.91MB ± 0%     +22.20%  (p=0.000 n=10+10)
/1000000-4    40.2MB ± 0%    48.9MB ± 0%     +21.68%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%        ~     (all equal)
/10-4           25.0 ± 0%      23.0 ± 0%      -8.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       203 ± 0%      -1.93%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.00k ± 0%      -0.60%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%      -0.38%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%      -0.37%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%      -0.36%  (p=0.000 n=10+10)
```

deque vs cookiejar - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4          36.2ns ± 1%    33.0ns ± 1%   -8.84%  (p=0.000 n=9+10)
/10-4          368ns ± 1%     334ns ± 0%   -9.20%  (p=0.000 n=10+9)
/100-4        3.60µs ± 2%    3.26µs ± 1%   -9.35%  (p=0.000 n=8+9)
/1000-4       35.6µs ± 1%    32.5µs ± 1%   -8.52%  (p=0.000 n=10+10)
/10000-4       364µs ± 4%     325µs ± 1%  -10.62%  (p=0.000 n=10+10)
/100000-4     3.58ms ± 1%    3.25ms ± 1%   -9.22%  (p=0.000 n=9+10)
/1000000-4    35.8ms ± 1%    32.5ms ± 0%   -9.30%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%   -0.00%  (p=0.000 n=9+9)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%   -0.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

deque vs cookiejar - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4          33.9ns ± 1%    31.4ns ± 1%   -7.23%  (p=0.000 n=8+10)
/10-4          344ns ± 1%     319ns ± 1%   -7.25%  (p=0.000 n=10+10)
/100-4        3.34µs ± 1%    3.11µs ± 1%   -7.09%  (p=0.000 n=8+9)
/1000-4       33.9µs ± 4%    31.0µs ± 1%   -8.57%  (p=0.000 n=10+9)
/10000-4       348µs ± 2%     309µs ± 1%  -11.23%  (p=0.000 n=10+10)
/100000-4     3.47ms ± 1%    3.08ms ± 1%  -11.06%  (p=0.000 n=10+10)
/1000000-4    35.0ms ± 2%    30.9ms ± 0%  -11.60%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%   +0.00%  (p=0.002 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%   +0.00%  (p=0.015 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

#### deque vs cookiejar - LIFO stack
deque vs cookiejar - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillCookiejarStack.txt
name       old time/op    new time/op    delta
/1-4         3.66µs ± 1%    3.47µs ± 0%   -5.38%  (p=0.000 n=9+8)
/10-4        34.1µs ± 1%    33.2µs ± 1%   -2.62%  (p=0.000 n=9+9)
/100-4        335µs ± 1%     321µs ± 1%   -4.40%  (p=0.000 n=9+10)
/1000-4      3.35ms ± 3%    3.19ms ± 1%   -4.61%  (p=0.000 n=10+10)
/10000-4     36.0ms ± 1%    32.7ms ± 1%   -9.22%  (p=0.000 n=9+9)
/100000-4     400ms ± 1%     366ms ± 1%   -8.53%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%   +0.01%  (p=0.000 n=10+9)
/100-4        160kB ± 0%     160kB ± 0%   +0.01%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   +0.01%  (p=0.000 n=10+9)
/10000-4     23.1MB ± 0%    16.0MB ± 0%  -30.79%  (p=0.000 n=10+9)
/100000-4     241MB ± 0%     161MB ± 0%  -33.36%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=10+10)
```

deque vs cookiejar - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullCookiejarStack.txt
name       old time/op    new time/op    delta
/1-4         3.60µs ± 1%    3.29µs ± 2%   -8.78%  (p=0.000 n=9+10)
/10-4        35.2µs ± 2%    32.6µs ± 1%   -7.35%  (p=0.000 n=10+9)
/100-4        339µs ± 5%     320µs ± 0%   -5.68%  (p=0.000 n=10+9)
/1000-4      3.35ms ± 2%    3.18ms ± 0%   -5.11%  (p=0.000 n=9+10)
/10000-4     35.7ms ± 6%    31.9ms ± 1%  -10.67%  (p=0.000 n=10+9)
/100000-4     388ms ± 2%     356ms ± 2%   -8.26%  (p=0.000 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   +0.00%  (p=0.014 n=9+10)
/10000-4     23.1MB ± 0%    16.0MB ± 0%  -30.81%  (p=0.000 n=10+10)
/100000-4     241MB ± 0%     160MB ± 0%  -33.58%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=10+10)
```

deque vs cookiejar - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4           196ns ± 1%    9014ns ± 1%   +4487.80%  (p=0.000 n=10+10)
/10-4          890ns ± 2%    9886ns ± 2%   +1010.14%  (p=0.000 n=10+10)
/100-4        7.50µs ± 1%   15.03µs ± 2%    +100.39%  (p=0.000 n=9+9)
/1000-4       68.4µs ± 5%    69.6µs ± 0%      +1.65%  (p=0.028 n=10+9)
/10000-4       712µs ± 2%     649µs ± 0%      -8.81%  (p=0.000 n=10+8)
/100000-4     7.68ms ± 2%    7.95ms ± 1%      +3.45%  (p=0.000 n=10+10)
/1000000-4    83.9ms ± 2%    80.1ms ± 2%      -4.57%  (p=0.000 n=9+9)

name        old alloc/op   new alloc/op   delta
/1-4            176B ± 0%    65704B ± 0%  +37231.82%  (p=0.000 n=10+10)
/10-4           592B ± 0%    65992B ± 0%  +11047.30%  (p=0.000 n=10+10)
/100-4        6.08kB ± 0%   68.87kB ± 0%   +1032.76%  (p=0.000 n=10+10)
/1000-4       41.2kB ± 0%    97.7kB ± 0%    +137.25%  (p=0.000 n=10+10)
/10000-4       403kB ± 0%     517kB ± 0%     +28.41%  (p=0.000 n=9+10)
/100000-4     4.02MB ± 0%    4.85MB ± 0%     +20.55%  (p=0.000 n=10+10)
/1000000-4    40.2MB ± 0%    48.8MB ± 0%     +21.50%  (p=0.000 n=9+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%        ~     (all equal)
/10-4           24.0 ± 0%      23.0 ± 0%      -4.17%  (p=0.000 n=10+10)
/100-4           207 ± 0%       203 ± 0%      -1.93%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.00k ± 0%      -0.50%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%      -0.38%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%      -0.37%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%      -0.36%  (p=0.000 n=10+10)
```

deque vs cookiejar - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreaseCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4          36.6ns ± 2%    33.2ns ± 0%  -9.35%  (p=0.000 n=10+10)
/10-4          369ns ± 1%     335ns ± 0%  -9.40%  (p=0.000 n=8+9)
/100-4        3.60µs ± 2%    3.27µs ± 0%  -9.13%  (p=0.000 n=10+10)
/1000-4       36.0µs ± 1%    32.6µs ± 1%  -9.36%  (p=0.000 n=9+10)
/10000-4       362µs ± 5%     328µs ± 2%  -9.43%  (p=0.000 n=10+9)
/100000-4     3.61ms ± 4%    3.37ms ± 6%  -6.63%  (p=0.000 n=10+9)
/1000000-4    35.7ms ± 1%    33.4ms ± 1%  -6.58%  (p=0.000 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%  -0.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%  -0.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

deque vs cookiejar - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4          35.4ns ± 2%    31.3ns ± 1%  -11.69%  (p=0.000 n=10+10)
/10-4          362ns ± 4%     316ns ± 1%  -12.72%  (p=0.000 n=10+9)
/100-4        3.51µs ± 2%    3.07µs ± 1%  -12.46%  (p=0.000 n=10+10)
/1000-4       34.7µs ± 1%    31.2µs ± 2%  -10.23%  (p=0.000 n=10+10)
/10000-4       352µs ± 3%     309µs ± 1%  -12.12%  (p=0.000 n=10+10)
/100000-4     3.49ms ± 2%    3.07ms ± 1%  -12.03%  (p=0.000 n=9+10)
/1000000-4    35.0ms ± 2%    30.7ms ± 1%  -12.31%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (p=0.097 n=10+10)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%   +0.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```
