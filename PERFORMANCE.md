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

Below results is for deque [v1.0.3](https://github.com/ef-ds/deque/blob/master/CHANGELOG.md).


### Fill Test Results
deque vs [list](https://github.com/golang/go/tree/master/src/container/list) - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueuev1.0.3.txt testdata/BenchmarkFillListQueue.txt
name        old time/op    new time/op    delta
/0-4          37.4ns ± 1%    39.1ns ± 2%    +4.47%  (p=0.000 n=10+9)
/1-4           171ns ± 1%     107ns ± 1%   -37.73%  (p=0.000 n=10+9)
/10-4          577ns ± 1%     726ns ± 1%   +25.92%  (p=0.000 n=10+10)
/100-4        4.74µs ± 2%    6.74µs ± 1%   +41.98%  (p=0.000 n=9+9)
/1000-4       37.1µs ± 3%    69.1µs ± 1%   +86.23%  (p=0.000 n=10+8)
/10000-4       370µs ± 2%     712µs ± 1%   +92.53%  (p=0.000 n=10+10)
/100000-4     3.87ms ± 0%   20.61ms ±10%  +432.98%  (p=0.000 n=8+10)
/1000000-4    44.1ms ± 1%   148.7ms ± 2%  +236.86%  (p=0.000 n=10+7)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%      112B ± 0%   -41.67%  (p=0.000 n=10+10)
/10-4           592B ± 0%      688B ± 0%   +16.22%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    6.45kB ± 0%   -10.44%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%    64.0kB ± 0%   +88.20%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%     640kB ± 0%   +98.11%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.40MB ± 0%   +98.65%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    64.0MB ± 0%   +98.83%  (p=0.000 n=10+10)

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

deque vs list - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStackv1.0.3.txt testdata/BenchmarkFillListStack.txt
name        old time/op    new time/op    delta
/0-4          37.5ns ± 1%    39.0ns ± 1%    +3.91%  (p=0.000 n=9+9)
/1-4           171ns ± 0%     107ns ± 1%   -37.05%  (p=0.000 n=10+10)
/10-4          579ns ± 1%     725ns ± 1%   +25.27%  (p=0.000 n=10+10)
/100-4        4.75µs ± 2%    6.75µs ± 1%   +42.23%  (p=0.000 n=10+10)
/1000-4       36.7µs ± 2%    68.4µs ± 0%   +86.25%  (p=0.000 n=9+10)
/10000-4       363µs ± 2%     700µs ± 0%   +92.94%  (p=0.000 n=10+9)
/100000-4     3.81ms ± 0%   21.25ms ±10%  +457.23%  (p=0.000 n=8+10)
/1000000-4    43.6ms ± 2%   157.6ms ±10%  +261.42%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%      112B ± 0%   -41.67%  (p=0.000 n=10+10)
/10-4           592B ± 0%      688B ± 0%   +16.22%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    6.45kB ± 0%   -10.44%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%    64.0kB ± 0%   +88.20%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%     640kB ± 0%   +98.11%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.40MB ± 0%   +98.65%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    64.0MB ± 0%   +98.83%  (p=0.000 n=10+8)

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
benchstat testdata/BenchmarkFillDequeQueuev1.0.3.txt testdata/BenchmarkFillSliceQueue.txt
name        old time/op    new time/op    delta
/0-4          37.4ns ± 1%    41.1ns ± 4%    +9.83%  (p=0.000 n=10+10)
/1-4           171ns ± 1%     101ns ± 7%   -41.01%  (p=0.000 n=10+8)
/10-4          577ns ± 1%     662ns ±14%   +14.81%  (p=0.000 n=10+9)
/100-4        4.74µs ± 2%    4.64µs ±30%      ~     (p=0.842 n=9+10)
/1000-4       37.1µs ± 3%    33.1µs ± 3%   -10.72%  (p=0.000 n=10+9)
/10000-4       370µs ± 2%     410µs ± 9%   +10.95%  (p=0.000 n=10+10)
/100000-4     3.87ms ± 0%    9.28ms ±10%  +140.06%  (p=0.000 n=8+9)
/1000000-4    44.1ms ± 1%   115.2ms ±33%  +160.97%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     32.0B ± 0%   -50.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%       56B ± 0%   -70.83%  (p=0.000 n=10+10)
/10-4           592B ± 0%      440B ± 0%   -25.68%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    3.67kB ± 0%   -49.00%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%    32.4kB ± 0%    -4.77%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%     546kB ± 0%   +69.10%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.25MB ± 0%   +94.13%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    61.2MB ± 0%   +90.10%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      16.0 ± 0%   +14.29%  (p=0.000 n=10+10)
/100-4           107 ± 0%       109 ± 0%    +1.87%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%    -0.10%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.61%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs [CustomSliceQueue](testdata_test.go) - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStackv1.0.3.txt testdata/BenchmarkFillSliceStack.txt
name        old time/op    new time/op    delta
/0-4          37.5ns ± 1%    43.4ns ± 3%   +15.63%  (p=0.000 n=9+9)
/1-4           171ns ± 0%     100ns ± 4%   -41.47%  (p=0.000 n=10+9)
/10-4          579ns ± 1%     623ns ± 5%    +7.66%  (p=0.000 n=10+8)
/100-4        4.75µs ± 2%    4.79µs ±37%      ~     (p=0.138 n=10+10)
/1000-4       36.7µs ± 2%    36.0µs ± 7%      ~     (p=0.905 n=9+10)
/10000-4       363µs ± 2%     426µs ± 6%   +17.46%  (p=0.000 n=10+9)
/100000-4     3.81ms ± 0%   10.34ms ±14%  +171.08%  (p=0.000 n=8+10)
/1000000-4    43.6ms ± 2%   116.0ms ± 9%  +166.14%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     32.0B ± 0%   -50.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%       56B ± 0%   -70.83%  (p=0.000 n=10+10)
/10-4           592B ± 0%      440B ± 0%   -25.68%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    3.67kB ± 0%   -49.00%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%    32.4kB ± 0%    -4.77%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%     546kB ± 0%   +69.10%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    6.25MB ± 0%   +94.13%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    61.2MB ± 0%   +90.10%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            4.00 ± 0%      3.00 ± 0%   -25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      16.0 ± 0%   +14.29%  (p=0.000 n=10+10)
/100-4           107 ± 0%       109 ± 0%    +1.87%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.01k ± 0%    -0.10%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.0k ± 0%    -0.61%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      100k ± 0%    -0.75%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.00M ± 0%    -0.77%  (p=0.000 n=10+10)
```

deque vs [impl7](https://github.com/christianrpetrin/queue-tests/tree/master/queueimpl7/queueimpl7.go) - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueuev1.0.3.txt testdata/BenchmarkFillImpl7Queue.txt
name        old time/op    new time/op    delta
/0-4          37.4ns ± 1%    35.9ns ± 3%   -4.08%  (p=0.000 n=10+9)
/1-4           171ns ± 1%     133ns ± 1%  -22.48%  (p=0.000 n=10+10)
/10-4          577ns ± 1%     764ns ± 7%  +32.53%  (p=0.000 n=10+9)
/100-4        4.74µs ± 2%    4.28µs ± 3%   -9.75%  (p=0.000 n=9+9)
/1000-4       37.1µs ± 3%    38.8µs ± 7%   +4.55%  (p=0.001 n=10+10)
/10000-4       370µs ± 2%     388µs ± 5%   +4.94%  (p=0.000 n=10+10)
/100000-4     3.87ms ± 0%    3.95ms ± 2%   +2.09%  (p=0.000 n=8+8)
/1000000-4    44.1ms ± 1%    45.9ms ± 4%   +4.02%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%  -25.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%      112B ± 0%  -41.67%  (p=0.000 n=10+10)
/10-4           592B ± 0%      736B ± 0%  +24.32%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    4.26kB ± 0%  -40.89%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%    33.2kB ± 0%   -2.40%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%     323kB ± 0%   -0.08%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.23MB ± 0%   +0.17%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    32.3MB ± 0%   +0.20%  (p=0.000 n=10+9)

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


deque vs [phf](https://github.com/phf/go-queue) - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueuev1.0.3.txt testdata/BenchmarkFillPhfQueue.txt
name        old time/op    new time/op    delta
/0-4          37.4ns ± 1%    67.5ns ±10%   +80.21%  (p=0.000 n=10+10)
/1-4           171ns ± 1%     107ns ± 5%   -37.36%  (p=0.000 n=10+10)
/10-4          577ns ± 1%     869ns ±16%   +50.70%  (p=0.000 n=10+9)
/100-4        4.74µs ± 2%    6.55µs ±29%   +37.94%  (p=0.000 n=9+10)
/1000-4       37.1µs ± 3%    53.0µs ±14%   +42.83%  (p=0.000 n=10+9)
/10000-4       370µs ± 2%     577µs ±14%   +55.91%  (p=0.000 n=10+9)
/100000-4     3.87ms ± 0%    7.65ms ±10%   +97.72%  (p=0.000 n=8+9)
/1000000-4    44.1ms ± 1%    78.6ms ± 6%   +78.10%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%      ~     (all equal)
/1-4            192B ± 0%       80B ± 0%   -58.33%  (p=0.000 n=10+10)
/10-4           592B ± 0%      832B ± 0%   +40.54%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    7.65kB ± 0%    +6.22%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%    65.1kB ± 0%   +91.16%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%     946kB ± 0%  +192.92%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    7.89MB ± 0%  +144.94%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    66.3MB ± 0%  +106.08%  (p=0.000 n=10+10)

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

deque vs [phf](https://github.com/phf/go-queue) - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStackv1.0.3.txt testdata/BenchmarkFillPhfStack.txt
name        old time/op    new time/op    delta
/0-4          37.5ns ± 1%    69.4ns ± 8%   +84.93%  (p=0.000 n=9+10)
/1-4           171ns ± 0%     105ns ± 5%   -38.57%  (p=0.000 n=10+10)
/10-4          579ns ± 1%     852ns ± 2%   +47.23%  (p=0.000 n=10+9)
/100-4        4.75µs ± 2%    5.85µs ± 5%   +23.35%  (p=0.000 n=10+10)
/1000-4       36.7µs ± 2%    47.5µs ± 5%   +29.40%  (p=0.000 n=9+10)
/10000-4       363µs ± 2%     539µs ± 4%   +48.75%  (p=0.000 n=10+10)
/100000-4     3.81ms ± 0%    7.33ms ± 7%   +92.10%  (p=0.000 n=8+10)
/1000000-4    43.6ms ± 2%    81.9ms ± 6%   +87.93%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%      ~     (all equal)
/1-4            192B ± 0%       80B ± 0%   -58.33%  (p=0.000 n=10+10)
/10-4           592B ± 0%      832B ± 0%   +40.54%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    7.65kB ± 0%    +6.22%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%    65.1kB ± 0%   +91.16%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%     946kB ± 0%  +192.92%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    7.89MB ± 0%  +144.94%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    66.3MB ± 0%  +106.08%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkFillDequeQueuev1.0.3.txt testdata/BenchmarkFillGammazeroQueue.txt
name        old time/op    new time/op    delta
/0-4          37.4ns ± 1%    38.7ns ± 6%      ~     (p=0.136 n=10+10)
/1-4           171ns ± 1%     176ns ± 5%    +2.74%  (p=0.002 n=10+9)
/10-4          577ns ± 1%     510ns ± 7%   -11.64%  (p=0.000 n=10+9)
/100-4        4.74µs ± 2%    5.47µs ±11%   +15.28%  (p=0.000 n=9+10)
/1000-4       37.1µs ± 3%    44.7µs ± 3%   +20.54%  (p=0.000 n=10+9)
/10000-4       370µs ± 2%     509µs ± 2%   +37.55%  (p=0.000 n=10+8)
/100000-4     3.87ms ± 0%    7.22ms ±20%   +86.66%  (p=0.000 n=8+9)
/1000000-4    44.1ms ± 1%    81.8ms ±15%   +85.24%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%      320B ± 0%   +66.67%  (p=0.000 n=10+10)
/10-4           592B ± 0%      464B ± 0%   -21.62%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    7.28kB ± 0%    +1.11%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%    64.7kB ± 0%   +90.08%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%     946kB ± 0%  +192.80%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    7.89MB ± 0%  +144.93%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    66.3MB ± 0%  +106.07%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkFillDequeStackv1.0.3.txt testdata/BenchmarkFillGammazeroStack.txt
name        old time/op    new time/op    delta
/0-4          37.5ns ± 1%    40.3ns ± 7%    +7.49%  (p=0.000 n=9+9)
/1-4           171ns ± 0%     186ns ±10%    +9.09%  (p=0.000 n=10+10)
/10-4          579ns ± 1%     524ns ± 7%    -9.52%  (p=0.000 n=10+10)
/100-4        4.75µs ± 2%    5.47µs ±11%   +15.28%  (p=0.000 n=10+9)
/1000-4       36.7µs ± 2%    46.9µs ± 6%   +27.65%  (p=0.000 n=9+10)
/10000-4       363µs ± 2%     542µs ± 7%   +49.44%  (p=0.000 n=10+10)
/100000-4     3.81ms ± 0%    7.21ms ± 7%   +89.04%  (p=0.000 n=8+9)
/1000000-4    43.6ms ± 2%    80.4ms ± 4%   +84.31%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%      320B ± 0%   +66.67%  (p=0.000 n=10+10)
/10-4           592B ± 0%      464B ± 0%   -21.62%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    7.28kB ± 0%    +1.11%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%    64.7kB ± 0%   +90.08%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%     946kB ± 0%  +192.80%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    7.89MB ± 0%  +144.93%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    66.3MB ± 0%  +106.07%  (p=0.000 n=10+10)

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

deque vs [juju](https://github.com/juju/utils/blob/master/deque/deque.go) - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueuev1.0.3.txt testdata/BenchmarkFillJujuQueue.txt
name        old time/op    new time/op    delta
/0-4          37.4ns ± 1%   376.2ns ± 5%   +904.54%  (p=0.000 n=10+10)
/1-4           171ns ± 1%     400ns ± 1%   +133.57%  (p=0.000 n=10+10)
/10-4          577ns ± 1%     741ns ± 0%    +28.53%  (p=0.000 n=10+9)
/100-4        4.74µs ± 2%    4.57µs ± 0%     -3.62%  (p=0.000 n=9+9)
/1000-4       37.1µs ± 3%    40.7µs ± 0%     +9.62%  (p=0.000 n=10+8)
/10000-4       370µs ± 2%     403µs ± 2%     +9.07%  (p=0.000 n=10+10)
/100000-4     3.87ms ± 0%    4.34ms ± 1%    +12.21%  (p=0.000 n=8+10)
/1000000-4    44.1ms ± 1%    59.1ms ± 4%    +33.95%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%   1184.0B ± 0%  +1750.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%     1200B ± 0%   +525.00%  (p=0.000 n=10+10)
/10-4           592B ± 0%     1344B ± 0%   +127.03%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    4.99kB ± 0%    -30.67%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%    34.8kB ± 0%     +2.40%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%     333kB ± 0%     +3.20%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.33MB ± 0%     +3.22%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    33.3MB ± 0%     +3.30%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      4.00 ± 0%   +300.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%      5.00 ± 0%    +25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      14.0 ± 0%       ~     (all equal)
/100-4           107 ± 0%       110 ± 0%     +2.80%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.05k ± 0%     +3.85%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.5k ± 0%     +3.86%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      105k ± 0%     +3.87%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.05M ± 0%     +3.88%  (p=0.000 n=10+10)
```

deque vs [juju](https://github.com/juju/utils/blob/master/deque/deque.go) - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStackv1.0.3.txt testdata/BenchmarkFillJujuStack.txt
name        old time/op    new time/op    delta
/0-4          37.5ns ± 1%   359.9ns ± 1%   +858.85%  (p=0.000 n=9+9)
/1-4           171ns ± 0%     404ns ± 0%   +136.81%  (p=0.000 n=10+9)
/10-4          579ns ± 1%     737ns ± 0%    +27.33%  (p=0.000 n=10+8)
/100-4        4.75µs ± 2%    4.58µs ± 1%     -3.56%  (p=0.000 n=10+10)
/1000-4       36.7µs ± 2%    40.8µs ± 1%    +11.16%  (p=0.000 n=9+10)
/10000-4       363µs ± 2%     401µs ± 0%    +10.66%  (p=0.000 n=10+10)
/100000-4     3.81ms ± 0%    4.37ms ± 1%    +14.63%  (p=0.000 n=8+9)
/1000000-4    43.6ms ± 2%    58.8ms ± 2%    +34.84%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%   1184.0B ± 0%  +1750.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%     1200B ± 0%   +525.00%  (p=0.000 n=10+10)
/10-4           592B ± 0%     1344B ± 0%   +127.03%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    4.99kB ± 0%    -30.67%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%    34.8kB ± 0%     +2.40%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%     333kB ± 0%     +3.20%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%    3.33MB ± 0%     +3.22%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%    33.3MB ± 0%     +3.30%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      4.00 ± 0%   +300.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%      5.00 ± 0%    +25.00%  (p=0.000 n=10+10)
/10-4           14.0 ± 0%      14.0 ± 0%       ~     (all equal)
/100-4           107 ± 0%       110 ± 0%     +2.80%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%     1.05k ± 0%     +3.85%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%     10.5k ± 0%     +3.86%  (p=0.000 n=10+10)
/100000-4       101k ± 0%      105k ± 0%     +3.87%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%     1.05M ± 0%     +3.88%  (p=0.000 n=10+10)
```

deque vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - FIFO queue - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeQueuev1.0.3.txt testdata/BenchmarkFillCookiejarQueue.txt
name        old time/op    new time/op     delta
/0-4          37.4ns ± 1%  10159.8ns ± 2%   +27028.84%  (p=0.000 n=10+8)
/1-4           171ns ± 1%    10443ns ±10%    +5996.44%  (p=0.000 n=10+10)
/10-4          577ns ± 1%    10646ns ±14%    +1745.64%  (p=0.000 n=10+9)
/100-4        4.74µs ± 2%    13.65µs ± 6%     +187.69%  (p=0.000 n=9+9)
/1000-4       37.1µs ± 3%     42.2µs ±10%      +13.77%  (p=0.000 n=10+10)
/10000-4       370µs ± 2%      342µs ± 6%       -7.50%  (p=0.000 n=10+10)
/100000-4     3.87ms ± 0%     3.87ms ±14%         ~     (p=1.000 n=8+10)
/1000000-4    44.1ms ± 1%     47.5ms ± 8%       +7.51%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op    delta
/0-4           64.0B ± 0%   65680.0B ± 0%  +102525.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%     65696B ± 0%   +34116.67%  (p=0.000 n=10+10)
/10-4           592B ± 0%     65840B ± 0%   +11021.62%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    67.28kB ± 0%     +834.44%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%     81.7kB ± 0%     +140.01%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%      357kB ± 0%      +10.46%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%     3.25MB ± 0%       +0.77%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%     32.8MB ± 0%       +2.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op   delta
/0-4            1.00 ± 0%       3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%       4.00 ± 0%         ~     (all equal)
/10-4           14.0 ± 0%       13.0 ± 0%       -7.14%  (p=0.000 n=10+10)
/100-4           107 ± 0%        103 ± 0%       -3.74%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%      1.00k ± 0%       -0.99%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%      10.0k ± 0%       -0.75%  (p=0.000 n=10+10)
/100000-4       101k ± 0%       100k ± 0%       -0.73%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%      1.00M ± 0%       -0.73%  (p=0.000 n=10+10)
```

deque vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - LIFO stack - [fill tests](benchmark-fill_test.go)
```
benchstat testdata/BenchmarkFillDequeStackv1.0.3.txt testdata/BenchmarkFillCookiejarStack.txt
name        old time/op    new time/op     delta
/0-4          37.5ns ± 1%  10254.3ns ± 3%   +27220.60%  (p=0.000 n=9+9)
/1-4           171ns ± 0%    10843ns ± 7%    +6256.00%  (p=0.000 n=10+9)
/10-4          579ns ± 1%    11329ns ±12%    +1856.72%  (p=0.000 n=10+10)
/100-4        4.75µs ± 2%    13.91µs ± 8%     +193.10%  (p=0.000 n=10+10)
/1000-4       36.7µs ± 2%     43.8µs ± 5%      +19.45%  (p=0.000 n=9+9)
/10000-4       363µs ± 2%      344µs ± 2%       -5.05%  (p=0.000 n=10+10)
/100000-4     3.81ms ± 0%     3.73ms ± 5%       -2.11%  (p=0.034 n=8+10)
/1000000-4    43.6ms ± 2%     46.7ms ± 3%       +7.13%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op    delta
/0-4           64.0B ± 0%   65680.0B ± 0%  +102525.00%  (p=0.000 n=10+10)
/1-4            192B ± 0%     65696B ± 0%   +34116.67%  (p=0.000 n=10+10)
/10-4           592B ± 0%     65840B ± 0%   +11021.62%  (p=0.000 n=10+10)
/100-4        7.20kB ± 0%    67.28kB ± 0%     +834.44%  (p=0.000 n=10+10)
/1000-4       34.0kB ± 0%     81.7kB ± 0%     +140.01%  (p=0.000 n=10+10)
/10000-4       323kB ± 0%      357kB ± 0%      +10.46%  (p=0.000 n=10+10)
/100000-4     3.22MB ± 0%     3.25MB ± 0%       +0.77%  (p=0.000 n=10+10)
/1000000-4    32.2MB ± 0%     32.8MB ± 0%       +2.00%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op   delta
/0-4            1.00 ± 0%       3.00 ± 0%     +200.00%  (p=0.000 n=10+10)
/1-4            4.00 ± 0%       4.00 ± 0%         ~     (all equal)
/10-4           14.0 ± 0%       13.0 ± 0%       -7.14%  (p=0.000 n=10+10)
/100-4           107 ± 0%        103 ± 0%       -3.74%  (p=0.000 n=10+10)
/1000-4        1.01k ± 0%      1.00k ± 0%       -0.99%  (p=0.000 n=10+10)
/10000-4       10.1k ± 0%      10.0k ± 0%       -0.75%  (p=0.000 n=10+10)
/100000-4       101k ± 0%       100k ± 0%       -0.73%  (p=0.000 n=10+10)
/1000000-4     1.01M ± 0%      1.00M ± 0%       -0.73%  (p=0.000 n=10+10)
```


### Microservice Test Results
deque vs [list](https://github.com/golang/go/tree/master/src/container/list) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.3.txt testdata/BenchmarkMicroserviceListQueue.txt
name        old time/op    new time/op    delta
/0-4          38.1ns ± 0%    42.5ns ±11%   +11.55%  (p=0.000 n=8+10)
/1-4           464ns ± 0%     511ns ± 6%   +10.18%  (p=0.000 n=10+9)
/10-4         2.77µs ± 1%    4.81µs ± 2%   +73.57%  (p=0.000 n=10+10)
/100-4        24.2µs ± 0%    46.8µs ± 1%   +93.20%  (p=0.000 n=8+8)
/1000-4        224µs ± 2%     500µs ± 5%  +123.50%  (p=0.000 n=10+8)
/10000-4      2.28ms ± 2%    5.39ms ± 4%  +137.05%  (p=0.000 n=10+9)
/100000-4     24.8ms ± 2%    75.0ms ± 8%  +203.12%  (p=0.000 n=10+10)
/1000000-4     259ms ± 3%     815ms ±13%  +214.53%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            544B ± 0%      496B ± 0%    -8.82%  (p=0.000 n=10+10)
/10-4         2.58kB ± 0%    4.53kB ± 0%   +75.78%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    44.8kB ± 0%  +114.13%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     448kB ± 0%  +233.93%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    4.48MB ± 0%  +212.80%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    44.8MB ± 0%  +210.56%  (p=0.000 n=10+9)
/1000000-4     144MB ± 0%     448MB ± 0%  +210.33%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            11.0 ± 0%      15.0 ± 0%   +36.36%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%     141.0 ± 0%   +88.00%  (p=0.000 n=10+10)
/100-4           709 ± 0%      1401 ± 0%   +97.60%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%    14.00k ± 0%   +99.59%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%    140.0k ± 0%   +99.56%  (p=0.000 n=10+10)
/100000-4       702k ± 0%     1400k ± 0%   +99.56%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%    14.00M ± 0%   +99.55%  (p=0.000 n=10+10)
```

deque vs [list](https://github.com/golang/go/tree/master/src/container/list) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.3.txt testdata/BenchmarkMicroserviceListStack.txt
name        old time/op    new time/op    delta
/0-4          38.1ns ± 1%    43.6ns ± 5%   +14.39%  (p=0.000 n=10+10)
/1-4           356ns ± 0%     510ns ± 4%   +43.43%  (p=0.000 n=8+9)
/10-4         2.59µs ± 7%    4.87µs ± 6%   +88.18%  (p=0.000 n=10+9)
/100-4        23.4µs ± 1%    47.1µs ± 2%  +101.40%  (p=0.000 n=10+8)
/1000-4        220µs ± 0%     488µs ±14%  +121.94%  (p=0.000 n=8+10)
/10000-4      2.27ms ± 1%    5.28ms ± 7%  +132.94%  (p=0.000 n=10+9)
/100000-4     24.9ms ± 2%    75.2ms ± 3%  +202.68%  (p=0.000 n=9+9)
/1000000-4     259ms ± 2%     941ms ±12%  +263.51%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            288B ± 0%      496B ± 0%   +72.22%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%    4.53kB ± 0%  +191.75%  (p=0.000 n=10+10)
/100-4        16.8kB ± 0%    44.8kB ± 0%  +166.95%  (p=0.000 n=10+10)
/1000-4        130kB ± 0%     448kB ± 0%  +244.57%  (p=0.000 n=10+10)
/10000-4      1.42MB ± 0%    4.48MB ± 0%  +214.62%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    44.8MB ± 0%  +210.65%  (p=0.000 n=9+10)
/1000000-4     144MB ± 0%     448MB ± 0%  +210.33%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            10.0 ± 0%      15.0 ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4           74.0 ± 0%     141.0 ± 0%   +90.54%  (p=0.000 n=10+10)
/100-4           707 ± 0%      1401 ± 0%   +98.16%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%    14.00k ± 0%   +99.64%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%    140.0k ± 0%   +99.57%  (p=0.000 n=10+10)
/100000-4       702k ± 0%     1400k ± 0%   +99.56%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%    14.00M ± 0%   +99.55%  (p=0.000 n=10+10)
```

deque vs [CustomSliceQueue](testdata_test.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.3.txt testdata/BenchmarkMicroserviceSliceQueue.txt
name        old time/op    new time/op    delta
/0-4          38.1ns ± 0%    38.5ns ± 1%    +1.15%  (p=0.000 n=8+10)
/1-4           464ns ± 0%     475ns ± 2%    +2.46%  (p=0.000 n=10+9)
/10-4         2.77µs ± 1%    3.54µs ± 2%   +27.82%  (p=0.000 n=10+10)
/100-4        24.2µs ± 0%    27.0µs ± 5%   +11.36%  (p=0.000 n=8+10)
/1000-4        224µs ± 2%     278µs ± 2%   +24.22%  (p=0.000 n=10+10)
/10000-4      2.28ms ± 2%    3.08ms ± 2%   +35.37%  (p=0.000 n=10+10)
/100000-4     24.8ms ± 2%    47.4ms ± 3%   +91.41%  (p=0.000 n=10+10)
/1000000-4     259ms ± 3%     596ms ± 7%  +129.97%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     32.0B ± 0%   -50.00%  (p=0.000 n=10+10)
/1-4            544B ± 0%      232B ± 0%   -57.35%  (p=0.000 n=10+10)
/10-4         2.58kB ± 0%    2.17kB ± 0%   -15.84%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    21.3kB ± 0%    +1.87%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     214kB ± 0%   +59.84%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    2.95MB ± 0%  +105.81%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    33.1MB ± 0%  +129.73%  (p=0.000 n=10+10)
/1000000-4     144MB ± 0%     338MB ± 0%  +134.04%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            11.0 ± 0%      14.0 ± 0%   +27.27%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%     101.0 ± 0%   +34.67%  (p=0.000 n=10+10)
/100-4           709 ± 0%       822 ± 0%   +15.94%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     8.77k ± 0%   +24.96%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     87.8k ± 0%   +25.20%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      875k ± 0%   +24.79%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     8.86M ± 0%   +26.32%  (p=0.000 n=10+10)
```

deque vs [CustomSliceQueue](testdata_test.go) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.3.txt testdata/BenchmarkMicroserviceSliceStack.txt
name        old time/op    new time/op    delta
/0-4          38.1ns ± 1%    38.7ns ± 0%   +1.49%  (p=0.000 n=10+10)
/1-4           356ns ± 0%     366ns ± 2%   +2.74%  (p=0.000 n=8+8)
/10-4         2.59µs ± 7%    2.15µs ± 2%  -16.88%  (p=0.000 n=10+9)
/100-4        23.4µs ± 1%    18.4µs ± 1%  -21.35%  (p=0.000 n=10+10)
/1000-4        220µs ± 0%     177µs ± 1%  -19.77%  (p=0.000 n=8+10)
/10000-4      2.27ms ± 1%    1.88ms ± 2%  -16.79%  (p=0.000 n=10+10)
/100000-4     24.9ms ± 2%    25.7ms ± 4%   +3.51%  (p=0.000 n=9+10)
/1000000-4     259ms ± 2%     283ms ± 5%   +9.28%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     32.0B ± 0%  -50.00%  (p=0.000 n=10+10)
/1-4            288B ± 0%      200B ± 0%  -30.56%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%    1.40kB ± 0%   -9.79%  (p=0.000 n=10+10)
/100-4        16.8kB ± 0%    13.3kB ± 0%  -21.00%  (p=0.000 n=10+10)
/1000-4        130kB ± 0%     128kB ± 0%   -1.25%  (p=0.000 n=10+10)
/10000-4      1.42MB ± 0%    1.51MB ± 0%   +5.78%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    15.9MB ± 0%   +9.94%  (p=0.000 n=9+9)
/1000000-4     144MB ± 0%     157MB ± 0%   +8.88%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            10.0 ± 0%      11.0 ± 0%  +10.00%  (p=0.000 n=10+10)
/10-4           74.0 ± 0%      76.0 ± 0%   +2.70%  (p=0.000 n=10+10)
/100-4           707 ± 0%       709 ± 0%   +0.28%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.01k ± 0%   -0.01%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%   -0.19%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%   -0.22%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%   -0.22%  (p=0.000 n=10+10)
```

deque vs [impl7](https://github.com/christianrpetrin/queue-tests/tree/master/queueimpl7/queueimpl7.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.3.txt testdata/BenchmarkMicroserviceImpl7Queue.txt
name        old time/op    new time/op    delta
/0-4          38.1ns ± 0%    37.5ns ± 1%    -1.52%  (p=0.000 n=8+9)
/1-4           464ns ± 0%     646ns ±10%   +39.38%  (p=0.000 n=10+9)
/10-4         2.77µs ± 1%    4.81µs ± 5%   +73.77%  (p=0.000 n=10+10)
/100-4        24.2µs ± 0%    32.3µs ± 3%   +33.35%  (p=0.000 n=8+10)
/1000-4        224µs ± 2%     313µs ± 5%   +39.90%  (p=0.000 n=10+10)
/10000-4      2.28ms ± 2%    3.17ms ±11%   +39.48%  (p=0.000 n=10+10)
/100000-4     24.8ms ± 2%    33.1ms ± 2%   +33.86%  (p=0.000 n=10+8)
/1000000-4     259ms ± 3%     348ms ± 8%   +34.17%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%   -25.00%  (p=0.000 n=10+10)
/1-4            544B ± 0%      432B ± 0%   -20.59%  (p=0.000 n=10+10)
/10-4         2.58kB ± 0%    6.91kB ± 0%  +168.32%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    29.6kB ± 0%   +41.48%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     261kB ± 0%   +94.51%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    2.58MB ± 0%   +80.05%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    25.8MB ± 0%   +78.52%  (p=0.000 n=10+9)
/1000000-4     144MB ± 0%     258MB ± 0%   +78.37%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/1-4            11.0 ± 0%      17.0 ± 0%   +54.55%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%     109.0 ± 0%   +45.33%  (p=0.000 n=10+10)
/100-4           709 ± 0%       927 ± 0%   +30.75%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     9.11k ± 0%   +29.88%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     91.0k ± 0%   +29.65%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      909k ± 0%   +29.62%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     9.09M ± 0%   +29.62%  (p=0.000 n=10+10)
```

deque vs [phf](https://github.com/phf/go-queue) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.3.txt testdata/BenchmarkMicroservicePhfQueue.txt
name        old time/op    new time/op    delta
/0-4          38.1ns ± 0%    63.1ns ± 0%   +65.58%  (p=0.000 n=8+8)
/1-4           464ns ± 0%     444ns ± 2%    -4.18%  (p=0.000 n=10+10)
/10-4         2.77µs ± 1%    3.39µs ± 7%   +22.25%  (p=0.000 n=10+10)
/100-4        24.2µs ± 0%    29.5µs ± 6%   +21.93%  (p=0.000 n=8+10)
/1000-4        224µs ± 2%     260µs ± 2%   +16.24%  (p=0.000 n=10+9)
/10000-4      2.28ms ± 2%    2.79ms ± 3%   +22.60%  (p=0.000 n=10+10)
/100000-4     24.8ms ± 2%    33.6ms ± 4%   +35.88%  (p=0.000 n=10+10)
/1000000-4     259ms ± 3%     361ms ± 9%   +39.20%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%      ~     (all equal)
/1-4            544B ± 0%      272B ± 0%   -50.00%  (p=0.000 n=10+10)
/10-4         2.58kB ± 0%    2.18kB ± 0%   -15.53%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    23.0kB ± 0%    +9.85%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     210kB ± 0%   +56.38%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    2.69MB ± 0%   +87.98%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    23.8MB ± 0%   +64.86%  (p=0.000 n=10+9)
/1000000-4     144MB ± 0%     213MB ± 0%   +47.31%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.3.txt testdata/BenchmarkMicroservicePhfStack.txt
name        old time/op    new time/op    delta
/0-4          38.1ns ± 1%    63.3ns ± 1%   +65.97%  (p=0.000 n=10+9)
/1-4           356ns ± 0%     458ns ± 9%   +28.78%  (p=0.000 n=8+10)
/10-4         2.59µs ± 7%    3.22µs ± 2%   +24.35%  (p=0.000 n=10+9)
/100-4        23.4µs ± 1%    29.0µs ± 4%   +24.13%  (p=0.000 n=10+10)
/1000-4        220µs ± 0%     265µs ± 5%   +20.24%  (p=0.000 n=8+10)
/10000-4      2.27ms ± 1%    2.90ms ± 3%   +28.07%  (p=0.000 n=10+9)
/100000-4     24.9ms ± 2%    36.0ms ± 3%   +44.75%  (p=0.000 n=9+7)
/1000000-4     259ms ± 2%     359ms ± 7%   +38.78%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     64.0B ± 0%      ~     (all equal)
/1-4            288B ± 0%      272B ± 0%    -5.56%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%    2.18kB ± 0%   +40.21%  (p=0.000 n=10+10)
/100-4        16.8kB ± 0%    23.0kB ± 0%   +36.95%  (p=0.000 n=10+10)
/1000-4        130kB ± 0%     210kB ± 0%   +61.36%  (p=0.000 n=10+10)
/10000-4      1.42MB ± 0%    2.69MB ± 0%   +89.08%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    23.8MB ± 0%   +64.91%  (p=0.000 n=9+10)
/1000000-4     144MB ± 0%     213MB ± 0%   +47.31%  (p=0.000 n=8+9)

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
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.3.txt testdata/BenchmarkMicroserviceGammazeroQueue.txt
name        old time/op    new time/op    delta
/0-4          38.1ns ± 0%    36.6ns ± 0%   -3.94%  (p=0.000 n=8+7)
/1-4           464ns ± 0%     379ns ± 8%  -18.22%  (p=0.000 n=10+10)
/10-4         2.77µs ± 1%    2.57µs ±11%   -7.15%  (p=0.003 n=10+9)
/100-4        24.2µs ± 0%    27.0µs ± 3%  +11.46%  (p=0.000 n=8+10)
/1000-4        224µs ± 2%     251µs ± 2%  +12.17%  (p=0.000 n=10+9)
/10000-4      2.28ms ± 2%    2.75ms ±18%  +20.87%  (p=0.000 n=10+10)
/100000-4     24.8ms ± 2%    33.3ms ±10%  +34.36%  (p=0.000 n=10+9)
/1000000-4     259ms ± 3%     341ms ± 5%  +31.49%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%  -25.00%  (p=0.000 n=10+10)
/1-4            544B ± 0%      416B ± 0%  -23.53%  (p=0.000 n=10+10)
/10-4         2.58kB ± 0%    1.42kB ± 0%  -44.72%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    22.3kB ± 0%   +6.26%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     209kB ± 0%  +55.82%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    2.69MB ± 0%  +87.93%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    23.8MB ± 0%  +64.86%  (p=0.000 n=10+9)
/1000000-4     144MB ± 0%     213MB ± 0%  +47.31%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            11.0 ± 0%       9.0 ± 0%  -18.18%  (p=0.000 n=10+10)
/10-4           75.0 ± 0%      72.0 ± 0%   -4.00%  (p=0.000 n=10+10)
/100-4           709 ± 0%       714 ± 0%   +0.71%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.03k ± 0%   +0.16%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%   -0.16%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%   -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%   -0.22%  (p=0.000 n=10+10)
```

deque vs [gammazero](https://github.com/gammazero/deque) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.3.txt testdata/BenchmarkMicroserviceGammazeroStack.txt
name        old time/op    new time/op    delta
/0-4          38.1ns ± 1%    36.6ns ± 0%   -3.95%  (p=0.000 n=10+9)
/1-4           356ns ± 0%     375ns ± 4%   +5.25%  (p=0.000 n=8+9)
/10-4         2.59µs ± 7%    2.56µs ± 5%     ~     (p=0.183 n=10+10)
/100-4        23.4µs ± 1%    27.6µs ± 7%  +18.02%  (p=0.000 n=10+10)
/1000-4        220µs ± 0%     270µs ± 4%  +22.65%  (p=0.000 n=8+9)
/10000-4      2.27ms ± 1%    2.67ms ± 2%  +17.85%  (p=0.000 n=10+9)
/100000-4     24.9ms ± 2%    33.5ms ±19%  +34.77%  (p=0.000 n=9+10)
/1000000-4     259ms ± 2%     342ms ± 6%  +32.06%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%     48.0B ± 0%  -25.00%  (p=0.000 n=10+10)
/1-4            288B ± 0%      416B ± 0%  +44.44%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%    1.42kB ± 0%   -8.25%  (p=0.000 n=10+10)
/100-4        16.8kB ± 0%    22.3kB ± 0%  +32.48%  (p=0.000 n=10+10)
/1000-4        130kB ± 0%     209kB ± 0%  +60.79%  (p=0.000 n=10+10)
/10000-4      1.42MB ± 0%    2.69MB ± 0%  +89.02%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    23.8MB ± 0%  +64.91%  (p=0.000 n=9+10)
/1000000-4     144MB ± 0%     213MB ± 0%  +47.31%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/1-4            10.0 ± 0%       9.0 ± 0%  -10.00%  (p=0.000 n=10+10)
/10-4           74.0 ± 0%      72.0 ± 0%   -2.70%  (p=0.000 n=10+10)
/100-4           707 ± 0%       714 ± 0%   +0.99%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.03k ± 0%   +0.19%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     70.0k ± 0%   -0.16%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      700k ± 0%   -0.21%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.00M ± 0%   -0.22%  (p=0.000 n=10+10)
```

deque vs [juju](https://github.com/juju/utils/blob/master/deque/deque.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.3.txt testdata/BenchmarkMicroserviceJujuQueue.txt
name        old time/op    new time/op    delta
/0-4          38.1ns ± 0%   357.8ns ± 0%   +839.05%  (p=0.000 n=8+9)
/1-4           464ns ± 0%     607ns ± 1%    +30.97%  (p=0.000 n=10+9)
/10-4         2.77µs ± 1%    3.12µs ± 1%    +12.49%  (p=0.000 n=10+10)
/100-4        24.2µs ± 0%    26.8µs ± 1%    +10.61%  (p=0.000 n=8+10)
/1000-4        224µs ± 2%     265µs ± 1%    +18.40%  (p=0.000 n=10+10)
/10000-4      2.28ms ± 2%    2.68ms ± 1%    +17.92%  (p=0.000 n=10+9)
/100000-4     24.8ms ± 2%    31.0ms ± 3%    +25.18%  (p=0.000 n=10+10)
/1000000-4     259ms ± 3%     340ms ± 4%    +31.13%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%   1184.0B ± 0%  +1750.00%  (p=0.000 n=10+10)
/1-4            544B ± 0%     1296B ± 0%   +138.24%  (p=0.000 n=10+10)
/10-4         2.58kB ± 0%    3.41kB ± 0%    +32.30%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    22.3kB ± 0%     +6.57%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     217kB ± 0%    +61.70%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    2.16MB ± 0%    +50.58%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    21.6MB ± 0%    +49.40%  (p=0.000 n=10+10)
/1000000-4     144MB ± 0%     216MB ± 0%    +49.28%  (p=0.000 n=10+8)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      4.00 ± 0%   +300.00%  (p=0.000 n=10+10)
/1-4            11.0 ± 0%      11.0 ± 0%       ~     (all equal)
/10-4           75.0 ± 0%      77.0 ± 0%     +2.67%  (p=0.000 n=10+10)
/100-4           709 ± 0%       731 ± 0%     +3.10%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.29k ± 0%     +3.86%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     72.8k ± 0%     +3.80%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      728k ± 0%     +3.79%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.28M ± 0%     +3.79%  (p=0.000 n=10+10)
```

deque vs [juju](https://github.com/juju/utils/blob/master/deque/deque.go) - LIFO stack - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.3.txt testdata/BenchmarkMicroserviceJujuStack.txt
name        old time/op    new time/op    delta
/0-4          38.1ns ± 1%   360.0ns ± 0%   +843.89%  (p=0.000 n=10+8)
/1-4           356ns ± 0%     608ns ± 1%    +70.71%  (p=0.000 n=8+10)
/10-4         2.59µs ± 7%    2.75µs ± 1%     +6.23%  (p=0.001 n=10+10)
/100-4        23.4µs ± 1%    26.2µs ± 1%    +12.04%  (p=0.000 n=10+10)
/1000-4        220µs ± 0%     252µs ± 1%    +14.62%  (p=0.000 n=8+10)
/10000-4      2.27ms ± 1%    2.56ms ± 0%    +13.05%  (p=0.000 n=10+10)
/100000-4     24.9ms ± 2%    30.0ms ± 2%    +20.52%  (p=0.000 n=9+10)
/1000000-4     259ms ± 2%     340ms ± 5%    +31.43%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%   1184.0B ± 0%  +1750.00%  (p=0.000 n=10+10)
/1-4            288B ± 0%     1296B ± 0%   +350.00%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%    2.30kB ± 0%    +48.45%  (p=0.000 n=10+10)
/100-4        16.8kB ± 0%    21.2kB ± 0%    +26.29%  (p=0.000 n=10+10)
/1000-4        130kB ± 0%     184kB ± 0%    +41.38%  (p=0.000 n=10+10)
/10000-4      1.42MB ± 0%    1.81MB ± 0%    +27.12%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    18.1MB ± 0%    +25.52%  (p=0.000 n=9+8)
/1000000-4     144MB ± 0%     181MB ± 0%    +25.38%  (p=0.000 n=8+10)

name        old allocs/op  new allocs/op  delta
/0-4            1.00 ± 0%      4.00 ± 0%   +300.00%  (p=0.000 n=10+10)
/1-4            10.0 ± 0%      11.0 ± 0%    +10.00%  (p=0.000 n=10+10)
/10-4           74.0 ± 0%      74.0 ± 0%       ~     (all equal)
/100-4           707 ± 0%       728 ± 0%     +2.97%  (p=0.000 n=10+10)
/1000-4        7.01k ± 0%     7.20k ± 0%     +2.61%  (p=0.000 n=10+10)
/10000-4       70.2k ± 0%     71.9k ± 0%     +2.46%  (p=0.000 n=10+10)
/100000-4       702k ± 0%      719k ± 0%     +2.45%  (p=0.000 n=10+10)
/1000000-4     7.02M ± 0%     7.19M ± 0%     +2.45%  (p=0.000 n=10+10)
```

deque vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - FIFO queue - [microservice tests](benchmark-microservice_test.go)
```
benchstat testdata/BenchmarkMicroserviceDequeQueuev1.0.3.txt testdata/BenchmarkMicroserviceCookiejarQueue.txt
name        old time/op    new time/op    delta
/0-4          38.1ns ± 0%  9813.4ns ± 1%   +25656.96%  (p=0.000 n=8+10)
/1-4           464ns ± 0%   10506ns ± 2%    +2165.22%  (p=0.000 n=10+10)
/10-4         2.77µs ± 1%   12.85µs ± 7%     +363.74%  (p=0.000 n=10+10)
/100-4        24.2µs ± 0%    31.3µs ± 3%      +29.13%  (p=0.000 n=8+10)
/1000-4        224µs ± 2%     226µs ± 7%         ~     (p=0.247 n=10+10)
/10000-4      2.28ms ± 2%    2.08ms ± 4%       -8.73%  (p=0.000 n=10+10)
/100000-4     24.8ms ± 2%    24.3ms ± 4%       -1.70%  (p=0.035 n=10+10)
/1000000-4     259ms ± 3%     242ms ± 4%       -6.54%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%  65680.0B ± 0%  +102525.00%  (p=0.000 n=10+10)
/1-4            544B ± 0%    65792B ± 0%   +11994.12%  (p=0.000 n=10+10)
/10-4         2.58kB ± 0%   66.80kB ± 0%    +2493.17%  (p=0.000 n=10+10)
/100-4        20.9kB ± 0%    76.9kB ± 0%     +267.07%  (p=0.000 n=10+10)
/1000-4        134kB ± 0%     243kB ± 0%      +81.30%  (p=0.000 n=10+10)
/10000-4      1.43MB ± 0%    1.38MB ± 0%       -3.47%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    12.9MB ± 0%      -10.49%  (p=0.000 n=10+10)
/1000000-4     144MB ± 0%     129MB ± 0%      -10.71%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkMicroserviceDequeStackv1.0.3.txt testdata/BenchmarkMicroserviceCookiejarStack.txt
name        old time/op    new time/op    delta
/0-4          38.1ns ± 1%  9834.4ns ± 1%   +25685.00%  (p=0.000 n=10+10)
/1-4           356ns ± 0%   10783ns ± 8%    +2929.97%  (p=0.000 n=8+10)
/10-4         2.59µs ± 7%   13.09µs ± 8%     +405.72%  (p=0.000 n=10+10)
/100-4        23.4µs ± 1%    30.6µs ± 4%      +31.13%  (p=0.000 n=10+8)
/1000-4        220µs ± 0%     220µs ± 6%         ~     (p=0.829 n=8+10)
/10000-4      2.27ms ± 1%    2.07ms ± 5%       -8.79%  (p=0.000 n=10+10)
/100000-4     24.9ms ± 2%    24.0ms ± 5%       -3.35%  (p=0.008 n=9+10)
/1000000-4     259ms ± 2%     250ms ± 5%       -3.29%  (p=0.015 n=10+10)

name        old alloc/op   new alloc/op   delta
/0-4           64.0B ± 0%  65680.0B ± 0%  +102525.00%  (p=0.000 n=10+10)
/1-4            288B ± 0%    65792B ± 0%   +22744.44%  (p=0.000 n=10+10)
/10-4         1.55kB ± 0%   66.80kB ± 0%    +4204.12%  (p=0.000 n=10+10)
/100-4        16.8kB ± 0%    76.9kB ± 0%     +357.62%  (p=0.000 n=10+10)
/1000-4        130kB ± 0%     178kB ± 0%      +36.64%  (p=0.000 n=10+10)
/10000-4      1.42MB ± 0%    1.32MB ± 0%       -7.52%  (p=0.000 n=10+10)
/100000-4     14.4MB ± 0%    12.8MB ± 0%      -10.92%  (p=0.000 n=9+10)
/1000000-4     144MB ± 0%     129MB ± 0%      -10.76%  (p=0.002 n=8+10)

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
#### deque vs [list](https://github.com/golang/go/tree/master/src/container/list) - FIFO queue
deque vs list - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueuev1.0.3.txt testdata/BenchmarkRefillListQueue.txt
name       old time/op    new time/op    delta
/1-4         3.95µs ± 7%    7.36µs ± 4%   +86.09%  (p=0.000 n=10+9)
/10-4        38.9µs ± 8%    74.2µs ±10%   +90.78%  (p=0.000 n=10+10)
/100-4        370µs ± 8%     716µs ± 3%   +93.37%  (p=0.000 n=10+9)
/1000-4      3.85ms ±10%    7.13ms ± 4%   +84.94%  (p=0.000 n=10+8)
/10000-4     41.0ms ±12%    75.9ms ± 9%   +85.28%  (p=0.000 n=10+10)
/100000-4     438ms ± 7%    2066ms ±11%  +371.54%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +299.99%  (p=0.000 n=10+9)
/10000-4     30.1MB ± 0%    64.0MB ± 0%  +112.51%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     640MB ± 0%  +100.12%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     2.00M ± 0%   +98.65%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     20.0M ± 0%   +98.47%  (p=0.000 n=10+10)
```

deque vs list - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.3.txt testdata/BenchmarkRefillFullListQueue.txt
name       old time/op    new time/op    delta
/1-4         3.55µs ± 1%    8.21µs ± 5%  +131.35%  (p=0.000 n=9+10)
/10-4        35.2µs ± 1%    81.5µs ± 4%  +131.70%  (p=0.000 n=10+10)
/100-4        344µs ± 1%     803µs ± 3%  +133.64%  (p=0.000 n=9+10)
/1000-4      3.41ms ± 1%    8.38ms ± 9%  +145.56%  (p=0.000 n=10+10)
/10000-4     40.6ms ±24%    86.6ms ± 7%  +113.04%  (p=0.000 n=9+10)
/100000-4     451ms ±15%    1608ms ± 6%  +256.88%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    64.0MB ± 0%  +112.52%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     640MB ± 0%  +100.12%  (p=0.000 n=9+10)

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
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowIncreaseListQueue.txt
name        old time/op    new time/op    delta
/1-4           241ns ±14%     179ns ± 5%   -25.55%  (p=0.000 n=10+10)
/10-4         1.32µs ±10%    1.44µs ± 4%    +8.85%  (p=0.000 n=10+10)
/100-4        9.07µs ±11%   14.12µs ± 7%   +55.65%  (p=0.000 n=10+9)
/1000-4       78.2µs ±10%   139.3µs ± 5%   +78.06%  (p=0.000 n=10+9)
/10000-4       805µs ± 7%    1473µs ± 5%   +82.92%  (p=0.000 n=9+8)
/100000-4     8.74ms ±12%   22.99ms ± 7%  +163.22%  (p=0.000 n=10+10)
/1000000-4    92.0ms ± 9%   268.3ms ± 8%  +191.71%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      176B ± 0%   -15.38%  (p=0.000 n=10+10)
/10-4         1.78kB ± 0%    1.33kB ± 0%   -25.23%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%   12.85kB ± 0%   +46.00%  (p=0.000 n=10+10)
/1000-4       54.2kB ± 0%   128.0kB ± 0%  +136.36%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%    1280kB ± 0%  +162.73%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%   12.80MB ± 0%  +165.46%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%   128.0MB ± 0%  +165.60%  (p=0.000 n=10+9)

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
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowDecreaseListQueue.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 7%    70.4ns ±12%   +79.88%  (p=0.000 n=9+9)
/10-4          391ns ±11%     703ns ± 4%   +79.64%  (p=0.000 n=10+9)
/100-4        3.71µs ± 8%    6.73µs ± 1%   +81.28%  (p=0.000 n=10+8)
/1000-4       36.1µs ± 7%    70.8µs ± 6%   +96.09%  (p=0.000 n=10+9)
/10000-4       375µs ±10%     696µs ± 3%   +85.42%  (p=0.000 n=10+8)
/100000-4     3.68ms ± 8%    7.25ms ± 6%   +97.18%  (p=0.000 n=10+8)
/1000000-4    37.0ms ± 9%    69.4ms ± 8%   +87.25%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=8+10)

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
benchstat testdata/BenchmarkStableDequeQueuev1.0.3.txt testdata/BenchmarkStableListQueue.txt
name        old time/op    new time/op    delta
/1-4          37.2ns ± 9%    86.9ns ± 6%  +133.90%  (p=0.000 n=10+10)
/10-4          369ns ± 8%     821ns ± 4%  +122.28%  (p=0.000 n=10+9)
/100-4        3.52µs ± 8%    7.76µs ± 1%  +120.10%  (p=0.000 n=10+8)
/1000-4       36.9µs ± 7%    90.0µs ±37%  +144.06%  (p=0.000 n=9+10)
/10000-4       353µs ± 9%     791µs ± 7%  +124.39%  (p=0.000 n=10+9)
/100000-4     3.60ms ± 8%    9.03ms ±26%  +150.64%  (p=0.000 n=10+9)
/1000000-4    35.5ms ± 7%    83.4ms ± 9%  +135.16%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+9)
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

#### deque vs [list](https://github.com/golang/go/tree/master/src/container/list) - LIFO stack
deque vs list - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStackv1.0.3.txt testdata/BenchmarkRefillListStack.txt
name       old time/op    new time/op    delta
/1-4         3.91µs ± 8%    7.28µs ± 5%   +86.21%  (p=0.000 n=10+10)
/10-4        36.6µs ± 8%    71.6µs ± 4%   +95.94%  (p=0.000 n=10+8)
/100-4        372µs ± 9%     718µs ± 6%   +93.22%  (p=0.000 n=10+10)
/1000-4      3.61ms ± 7%    7.04ms ± 5%   +94.70%  (p=0.000 n=10+10)
/10000-4     42.0ms ± 7%    77.3ms ±16%   +83.92%  (p=0.000 n=10+9)
/100000-4     444ms ±14%    2156ms ±11%  +385.49%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +299.99%  (p=0.000 n=10+9)
/10000-4     30.1MB ± 0%    64.0MB ± 0%  +112.69%  (p=0.000 n=8+9)
/100000-4     320MB ± 0%     640MB ± 0%  +100.02%  (p=0.000 n=9+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     2.00M ± 0%   +98.65%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     20.0M ± 0%   +98.47%  (p=0.000 n=9+9)
```

deque vs list - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.3.txt testdata/BenchmarkRefillFullListStack.txt
name       old time/op    new time/op    delta
/1-4         3.85µs ± 9%    8.28µs ± 4%  +115.07%  (p=0.000 n=10+9)
/10-4        38.2µs ± 9%    83.8µs ± 2%  +119.41%  (p=0.000 n=10+9)
/100-4        378µs ± 8%     831µs ± 2%  +119.62%  (p=0.000 n=10+10)
/1000-4      3.76ms ± 8%    8.55ms ±11%  +127.58%  (p=0.000 n=10+9)
/10000-4     42.7ms ±10%    96.3ms ±10%  +125.48%  (p=0.000 n=10+9)
/100000-4     446ms ±11%    1619ms ± 9%  +262.62%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    64.0MB ± 0%  +112.70%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     640MB ± 0%  +100.03%  (p=0.000 n=10+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     2.00k ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     20.0k ± 0%  +100.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      200k ± 0%  +100.00%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     2.00M ± 0%   +98.65%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     20.0M ± 0%   +98.47%  (p=0.000 n=10+8)
```

deque vs list - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.3.txt testdata/BenchmarkSlowIncreaseListStack.txt
name        old time/op    new time/op    delta
/1-4           226ns ±10%     180ns ± 6%   -20.08%  (p=0.000 n=10+10)
/10-4         1.06µs ± 7%    1.46µs ± 9%   +38.60%  (p=0.000 n=10+9)
/100-4        8.71µs ±13%   13.99µs ± 9%   +60.51%  (p=0.000 n=10+9)
/1000-4       75.5µs ± 9%   139.6µs ± 2%   +84.81%  (p=0.000 n=10+8)
/10000-4       731µs ± 6%    1520µs ±11%  +107.86%  (p=0.000 n=10+10)
/100000-4     8.81ms ±10%   27.93ms ± 9%  +216.96%  (p=0.000 n=10+10)
/1000000-4    94.4ms ±11%   357.6ms ±29%  +278.78%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      176B ± 0%   -15.38%  (p=0.000 n=10+10)
/10-4           752B ± 0%     1328B ± 0%   +76.60%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%   12.85kB ± 0%   +46.00%  (p=0.000 n=10+10)
/1000-4       50.0kB ± 0%   128.0kB ± 0%  +155.93%  (p=0.000 n=10+10)
/10000-4       483kB ± 0%    1280kB ± 0%  +164.98%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%   12.80MB ± 0%  +165.46%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%   128.0MB ± 0%  +165.60%  (p=0.000 n=10+9)

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
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.3.txt testdata/BenchmarkSlowDecreaseListStack.txt
name        old time/op    new time/op    delta
/1-4          37.3ns ± 7%    72.6ns ± 3%   +94.76%  (p=0.000 n=10+8)
/10-4          386ns ± 9%     717ns ± 3%   +85.94%  (p=0.000 n=10+10)
/100-4        3.76µs ± 8%    7.01µs ± 4%   +86.22%  (p=0.000 n=10+9)
/1000-4       37.0µs ± 6%    81.4µs ±36%  +120.12%  (p=0.000 n=10+10)
/10000-4       377µs ± 8%     699µs ± 5%   +85.62%  (p=0.000 n=10+8)
/100000-4     3.72ms ± 9%    7.22ms ± 6%   +93.92%  (p=0.000 n=10+8)
/1000000-4    38.1ms ± 9%    71.9ms ± 7%   +88.55%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+8)
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

deque vs list - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStackv1.0.3.txt testdata/BenchmarkStableListStack.txt
name        old time/op    new time/op    delta
/1-4          35.6ns ± 7%    87.1ns ±12%  +144.88%  (p=0.000 n=10+9)
/10-4          362ns ± 8%     920ns ±31%  +153.95%  (p=0.000 n=10+10)
/100-4        3.54µs ± 7%    8.71µs ±17%  +145.98%  (p=0.000 n=10+10)
/1000-4       35.9µs ±10%    87.1µs ±31%  +142.61%  (p=0.000 n=10+9)
/10000-4       360µs ±10%     801µs ± 6%  +122.49%  (p=0.000 n=10+9)
/100000-4     3.53ms ± 6%    8.66ms ±23%  +145.26%  (p=0.000 n=10+9)
/1000000-4    34.4ms ± 7%    85.9ms ±13%  +149.45%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+9)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=8+9)

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
benchstat testdata/BenchmarkRefillDequeQueuev1.0.3.txt testdata/BenchmarkRefillSliceQueue.txt
name       old time/op    new time/op    delta
/1-4         3.95µs ± 7%    6.00µs ± 9%   +51.75%  (p=0.000 n=10+10)
/10-4        38.9µs ± 8%    45.8µs ±11%   +17.78%  (p=0.000 n=10+10)
/100-4        370µs ± 8%     330µs ± 4%   -10.88%  (p=0.000 n=10+10)
/1000-4      3.85ms ±10%    3.77ms ±26%      ~     (p=0.481 n=10+10)
/10000-4     41.0ms ±12%    38.6ms ±13%      ~     (p=0.089 n=10+10)
/100000-4     438ms ± 7%     813ms ±10%   +85.49%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    2.40kB ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    32.0kB ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     314kB ± 0%   +96.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    3.15MB ± 0%   +96.79%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    48.8MB ± 0%   +61.94%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     548MB ± 0%   +71.31%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       200 ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.1k ± 0%    +1.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      100k ± 0%    +0.20%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.62%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.76%  (p=0.000 n=10+10)
```

deque vs CustomSliceQueue - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.3.txt testdata/BenchmarkRefillFullSliceQueue.txt
name       old time/op    new time/op    delta
/1-4         3.55µs ± 1%    4.16µs ±10%   +17.21%  (p=0.000 n=9+9)
/10-4        35.2µs ± 1%    40.4µs ± 7%   +14.85%  (p=0.000 n=10+10)
/100-4        344µs ± 1%     406µs ±27%   +18.30%  (p=0.000 n=9+9)
/1000-4      3.41ms ± 1%    4.11ms ± 5%   +20.52%  (p=0.000 n=10+8)
/10000-4     40.6ms ±24%    40.4ms ± 6%      ~     (p=0.258 n=9+9)
/100000-4     451ms ±15%     929ms ±18%  +106.12%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    4.82kB ± 0%  +200.94%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    48.2kB ± 0%  +201.08%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     483kB ± 0%  +201.70%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    5.15MB ± 0%  +221.87%  (p=0.000 n=10+8)
/10000-4     30.1MB ± 0%    52.0MB ± 0%   +72.82%  (p=0.000 n=10+8)
/100000-4     320MB ± 0%     551MB ± 0%   +72.32%  (p=0.000 n=9+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%    +0.03%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      100k ± 0%    +0.03%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.65%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.76%  (p=0.000 n=10+9)
```

deque vs CustomSliceQueue - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowIncreaseSliceQueue.txt
name        old time/op    new time/op    delta
/1-4           241ns ±14%     160ns ± 7%   -33.42%  (p=0.000 n=10+8)
/10-4         1.32µs ±10%    1.10µs ± 4%   -16.50%  (p=0.000 n=10+9)
/100-4        9.07µs ±11%    7.23µs ± 5%   -20.30%  (p=0.000 n=10+9)
/1000-4       78.2µs ±10%    66.7µs ± 4%   -14.77%  (p=0.000 n=10+9)
/10000-4       805µs ± 7%     729µs ± 3%    -9.50%  (p=0.000 n=9+10)
/100000-4     8.74ms ±12%   15.10ms ±23%   +72.88%  (p=0.000 n=10+9)
/1000000-4    92.0ms ± 9%   169.3ms ± 9%   +84.09%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%       88B ± 0%   -57.69%  (p=0.000 n=10+10)
/10-4         1.78kB ± 0%    0.78kB ± 0%   -56.31%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%    6.66kB ± 0%   -24.27%  (p=0.000 n=10+10)
/1000-4       54.2kB ± 0%    74.0kB ± 0%   +36.58%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%     991kB ± 0%  +103.32%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%   11.42MB ± 0%  +136.75%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%   114.6MB ± 0%  +137.74%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%      ~     (all equal)
/10-4           25.0 ± 0%      29.0 ± 0%   +16.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       214 ± 0%    +3.38%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.02k ± 0%    +0.25%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%    -0.25%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%    -0.37%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%    -0.39%  (p=0.000 n=10+10)
```

deque vs CustomSliceQueue - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowDecreaseSliceQueue.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 7%    56.3ns ± 3%   +43.72%  (p=0.000 n=9+8)
/10-4          391ns ±11%     574ns ± 4%   +46.73%  (p=0.000 n=10+8)
/100-4        3.71µs ± 8%    5.50µs ± 3%   +48.18%  (p=0.000 n=10+9)
/1000-4       36.1µs ± 7%    55.6µs ± 8%   +53.96%  (p=0.000 n=10+9)
/10000-4       375µs ±10%     547µs ± 2%   +45.69%  (p=0.000 n=10+8)
/100000-4     3.68ms ± 8%    6.84ms ±61%   +86.09%  (p=0.000 n=10+10)
/1000000-4    37.0ms ± 9%    56.0ms ± 5%   +51.29%  (p=0.000 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     24.0B ± 0%   +50.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      240B ± 0%   +50.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    2.40kB ± 0%   +50.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    24.0kB ± 0%   +50.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     240kB ± 0%   +50.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    2.40MB ± 0%   +50.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    24.0MB ± 0%   +50.00%  (p=0.000 n=8+10)

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
benchstat testdata/BenchmarkStableDequeQueuev1.0.3.txt testdata/BenchmarkStableSliceQueue.txt
name        old time/op    new time/op    delta
/1-4          37.2ns ± 9%    39.2ns ± 5%      ~     (p=0.097 n=10+8)
/10-4          369ns ± 8%     452ns ±37%   +22.35%  (p=0.000 n=10+10)
/100-4        3.52µs ± 8%    3.77µs ± 8%    +6.99%  (p=0.007 n=10+9)
/1000-4       36.9µs ± 7%    39.7µs ±25%    +7.61%  (p=0.031 n=9+9)
/10000-4       353µs ± 9%     400µs ± 4%   +13.35%  (p=0.000 n=10+9)
/100000-4     3.60ms ± 8%    3.91ms ± 7%    +8.62%  (p=0.000 n=10+9)
/1000000-4    35.5ms ± 7%    39.4ms ± 5%   +11.05%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     48.0B ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      481B ± 0%  +200.62%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    4.82kB ± 0%  +200.94%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    48.2kB ± 0%  +200.96%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     482kB ± 0%  +200.97%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    4.82MB ± 0%  +200.97%  (p=0.000 n=10+8)
/1000000-4    16.0MB ± 0%    48.2MB ± 0%  +200.97%  (p=0.000 n=10+7)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%      ~     (all equal)
/100-4           100 ± 0%       100 ± 0%      ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    +0.03%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      100k ± 0%    +0.03%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     1.00M ± 0%    +0.03%  (p=0.000 n=10+10)
```

#### deque vs [CustomSliceQueue](testdata_test.go) - LIFO stack
deque vs CustomSliceQueue - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStackv1.0.3.txt testdata/BenchmarkRefillSliceStack.txt
name       old time/op    new time/op    delta
/1-4         3.91µs ± 8%    3.01µs ± 5%  -22.99%  (p=0.000 n=10+10)
/10-4        36.6µs ± 8%    31.7µs ±24%  -13.20%  (p=0.035 n=10+10)
/100-4        372µs ± 9%     286µs ±17%  -22.93%  (p=0.000 n=10+9)
/1000-4      3.61ms ± 7%    2.66ms ± 2%  -26.46%  (p=0.000 n=10+9)
/10000-4     42.0ms ± 7%    29.4ms ±16%  -30.04%  (p=0.000 n=10+9)
/100000-4     444ms ±14%     343ms ± 3%  -22.70%  (p=0.000 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%   -0.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   -0.00%  (p=0.000 n=10+9)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.80%  (p=0.000 n=8+9)
/100000-4     320MB ± 0%     162MB ± 0%  -49.51%  (p=0.000 n=9+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=9+8)
```

deque vs CustomSliceQueue - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.3.txt testdata/BenchmarkRefillFullSliceStack.txt
name       old time/op    new time/op    delta
/1-4         3.85µs ± 9%    3.11µs ± 9%  -19.18%  (p=0.000 n=10+9)
/10-4        38.2µs ± 9%    29.6µs ± 9%  -22.57%  (p=0.000 n=10+9)
/100-4        378µs ± 8%     288µs ± 6%  -23.78%  (p=0.000 n=10+9)
/1000-4      3.76ms ± 8%    2.96ms ±12%  -21.07%  (p=0.000 n=10+9)
/10000-4     42.7ms ±10%    29.3ms ± 7%  -31.46%  (p=0.000 n=10+10)
/100000-4     446ms ±11%     336ms ± 2%  -24.67%  (p=0.000 n=10+8)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.83%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     160MB ± 0%  -49.99%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.3.txt testdata/BenchmarkSlowIncreaseSliceStack.txt
name        old time/op    new time/op    delta
/1-4           226ns ±10%     158ns ± 5%  -30.10%  (p=0.000 n=10+10)
/10-4         1.06µs ± 7%    0.89µs ± 6%  -15.74%  (p=0.000 n=10+10)
/100-4        8.71µs ±13%    6.52µs ± 2%  -25.22%  (p=0.000 n=10+8)
/1000-4       75.5µs ± 9%    65.0µs ±25%     ~     (p=0.075 n=10+10)
/10000-4       731µs ± 6%     639µs ± 3%  -12.62%  (p=0.000 n=10+9)
/100000-4     8.81ms ±10%   11.38ms ± 5%  +29.09%  (p=0.000 n=10+9)
/1000000-4    94.4ms ±11%   138.6ms ± 8%  +46.78%  (p=0.000 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%       88B ± 0%  -57.69%  (p=0.000 n=10+10)
/10-4           752B ± 0%      600B ± 0%  -20.21%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%    5.27kB ± 0%  -40.09%  (p=0.000 n=10+10)
/1000-4       50.0kB ± 0%    48.4kB ± 0%   -3.25%  (p=0.000 n=10+10)
/10000-4       483kB ± 0%     706kB ± 0%  +46.22%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    7.85MB ± 0%  +62.89%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    77.2MB ± 0%  +60.17%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      5.00 ± 0%     ~     (all equal)
/10-4           24.0 ± 0%      26.0 ± 0%   +8.33%  (p=0.000 n=10+10)
/100-4           207 ± 0%       209 ± 0%   +0.97%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.01k ± 0%   -0.05%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.0k ± 0%   -0.31%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      200k ± 0%   -0.38%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.00M ± 0%   -0.39%  (p=0.000 n=10+10)
```

deque vs CustomSliceQueue - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.3.txt testdata/BenchmarkSlowDecreaseSliceStack.txt
name        old time/op    new time/op    delta
/1-4          37.3ns ± 7%    32.3ns ± 6%  -13.20%  (p=0.000 n=10+8)
/10-4          386ns ± 9%     307ns ± 2%  -20.28%  (p=0.000 n=10+9)
/100-4        3.76µs ± 8%    3.04µs ± 4%  -19.31%  (p=0.000 n=10+10)
/1000-4       37.0µs ± 6%    31.8µs ±11%  -14.15%  (p=0.000 n=10+9)
/10000-4       377µs ± 8%     317µs ± 6%  -15.79%  (p=0.000 n=10+10)
/100000-4     3.72ms ± 9%    3.08ms ± 5%  -17.14%  (p=0.000 n=10+10)
/1000000-4    38.1ms ± 9%    30.8ms ± 5%  -19.15%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

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
benchstat testdata/BenchmarkStableDequeStackv1.0.3.txt testdata/BenchmarkStableSliceStack.txt
name        old time/op    new time/op    delta
/1-4          35.6ns ± 7%    27.4ns ± 5%  -23.07%  (p=0.000 n=10+9)
/10-4          362ns ± 8%     282ns ± 6%  -22.19%  (p=0.000 n=10+8)
/100-4        3.54µs ± 7%    2.70µs ± 3%  -23.85%  (p=0.000 n=10+8)
/1000-4       35.9µs ±10%    26.8µs ± 6%  -25.28%  (p=0.000 n=10+8)
/10000-4       360µs ±10%     265µs ± 6%  -26.39%  (p=0.000 n=10+9)
/100000-4     3.53ms ± 6%    2.64ms ± 4%  -25.20%  (p=0.000 n=10+8)
/1000000-4    34.4ms ± 7%    26.9ms ± 7%  -21.81%  (p=0.000 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

#### deque vs [impl7](https://github.com/christianrpetrin/queue-tests/tree/master/queueimpl7/queueimpl7.go) - FIFO queue
deque vs impl7 - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueuev1.0.3.txt testdata/BenchmarkRefillImpl7Queue.txt
name       old time/op    new time/op    delta
/1-4         3.95µs ± 7%   10.03µs ± 6%  +153.77%  (p=0.000 n=10+10)
/10-4        38.9µs ± 8%    74.7µs ± 5%   +92.18%  (p=0.000 n=10+10)
/100-4        370µs ± 8%     442µs ± 4%   +19.23%  (p=0.000 n=10+10)
/1000-4      3.85ms ±10%    3.97ms ± 3%      ~     (p=0.243 n=10+9)
/10000-4     41.0ms ±12%    39.1ms ± 6%      ~     (p=0.315 n=10+10)
/100000-4     438ms ± 7%     421ms ±24%      ~     (p=0.079 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    68.8kB ± 0%  +330.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     421kB ± 0%  +163.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    3.32MB ± 0%  +107.29%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    32.3MB ± 0%    +7.17%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     323MB ± 0%    +0.90%  (p=0.000 n=9+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       300 ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.60k ± 0%   +60.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.8k ± 0%    +8.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      102k ± 0%    +2.20%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.02M ± 0%    +0.93%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.2M ± 0%    +0.79%  (p=0.000 n=10+10)
```

deque vs impl7 - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.3.txt testdata/BenchmarkRefillFullImpl7Queue.txt
name       old time/op    new time/op    delta
/1-4         3.55µs ± 1%    4.24µs ± 4%   +19.45%  (p=0.000 n=9+10)
/10-4        35.2µs ± 1%    43.3µs ±24%   +23.14%  (p=0.000 n=10+10)
/100-4        344µs ± 1%     398µs ± 5%   +15.81%  (p=0.000 n=9+10)
/1000-4      3.41ms ± 1%    4.01ms ± 7%   +17.55%  (p=0.000 n=10+10)
/10000-4     40.6ms ±24%    39.9ms ± 7%      ~     (p=0.258 n=9+9)
/100000-4     451ms ±15%     412ms ±15%    -8.65%  (p=0.011 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    3.23kB ± 0%  +101.56%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    32.2kB ± 0%  +101.56%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     322kB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    3.23MB ± 0%  +101.56%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    32.3MB ± 0%    +7.09%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     322MB ± 0%    +0.84%  (p=0.000 n=9+8)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       101 ± 0%    +1.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.01k ± 0%    +1.50%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.2k ± 0%    +1.56%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      102k ± 0%    +1.56%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.02M ± 0%    +0.88%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.2M ± 0%    +0.79%  (p=0.000 n=10+10)
```

deque vs impl7 - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowIncreaseImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4           241ns ±14%     230ns ± 3%     ~     (p=0.138 n=10+10)
/10-4         1.32µs ±10%    1.61µs ± 6%  +21.75%  (p=0.000 n=10+10)
/100-4        9.07µs ±11%    7.87µs ± 2%  -13.30%  (p=0.000 n=10+9)
/1000-4       78.2µs ±10%    71.5µs ± 1%   -8.59%  (p=0.008 n=10+9)
/10000-4       805µs ± 7%     723µs ± 3%  -10.27%  (p=0.000 n=9+9)
/100000-4     8.74ms ±12%    8.50ms ± 9%     ~     (p=0.579 n=10+10)
/1000000-4    92.0ms ± 9%    93.4ms ± 4%     ~     (p=0.631 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      160B ± 0%  -23.08%  (p=0.000 n=10+10)
/10-4         1.78kB ± 0%    2.98kB ± 0%  +67.57%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%    7.94kB ± 0%   -9.82%  (p=0.000 n=10+10)
/1000-4       54.2kB ± 0%    65.9kB ± 0%  +21.56%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%     647kB ± 0%  +32.82%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    6.45MB ± 0%  +33.80%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    64.5MB ± 0%  +33.84%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      6.00 ± 0%  +20.00%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      29.0 ± 0%  +16.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       211 ± 0%   +1.93%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.04k ± 0%   +1.19%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.3k ± 0%   +1.18%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      203k ± 0%   +1.17%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.03M ± 0%   +1.17%  (p=0.000 n=10+10)
```

deque vs impl7 - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowDecreaseImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 7%    99.0ns ±12%  +152.79%  (p=0.000 n=9+10)
/10-4          391ns ±11%    1013ns ± 8%  +158.91%  (p=0.000 n=10+10)
/100-4        3.71µs ± 8%    9.67µs ± 5%  +160.70%  (p=0.000 n=10+9)
/1000-4       36.1µs ± 7%    93.4µs ± 3%  +158.71%  (p=0.000 n=10+8)
/10000-4       375µs ±10%     983µs ±12%  +161.89%  (p=0.000 n=10+10)
/100000-4     3.68ms ± 8%   10.70ms ±22%  +191.20%  (p=0.000 n=10+10)
/1000000-4    37.0ms ± 9%    92.9ms ± 6%  +150.68%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     64.0B ± 0%  +300.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      640B ± 0%  +300.00%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    6.40kB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    64.0kB ± 0%  +300.00%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     640kB ± 0%  +300.00%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    6.40MB ± 0%  +300.00%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    64.0MB ± 0%  +300.00%  (p=0.000 n=8+9)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      3.00 ± 0%  +200.00%  (p=0.000 n=10+10)
/10-4           10.0 ± 0%      30.0 ± 0%  +200.00%  (p=0.000 n=10+10)
/100-4           100 ± 0%       300 ± 0%  +200.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     3.00k ± 0%  +200.00%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     30.0k ± 0%  +200.00%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      300k ± 0%  +200.00%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     3.00M ± 0%  +200.00%  (p=0.000 n=10+10)
```

deque vs impl7 - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueuev1.0.3.txt testdata/BenchmarkStableImpl7Queue.txt
name        old time/op    new time/op    delta
/1-4          37.2ns ± 9%    38.1ns ± 5%      ~     (p=0.646 n=10+9)
/10-4          369ns ± 8%     391ns ± 7%    +5.96%  (p=0.022 n=10+10)
/100-4        3.52µs ± 8%    4.13µs ±19%   +17.30%  (p=0.000 n=10+10)
/1000-4       36.9µs ± 7%    39.6µs ±12%    +7.55%  (p=0.008 n=9+10)
/10000-4       353µs ± 9%     385µs ± 4%    +9.28%  (p=0.000 n=10+10)
/100000-4     3.60ms ± 8%    3.85ms ±11%    +6.81%  (p=0.019 n=10+10)
/1000000-4    35.5ms ± 7%    38.3ms ± 6%    +7.89%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     32.0B ± 0%  +100.00%  (p=0.000 n=10+10)
/10-4           160B ± 0%      322B ± 0%  +101.25%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    3.23kB ± 0%  +101.56%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    32.2kB ± 0%  +101.56%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     322kB ± 0%  +101.56%  (p=0.000 n=10+10)
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

#### deque vs [phf](https://github.com/phf/go-queue) - FIFO queue
deque vs phf - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueuev1.0.3.txt testdata/BenchmarkRefillPhfQueue.txt
name       old time/op    new time/op    delta
/1-4         3.95µs ± 7%    3.87µs ± 1%      ~     (p=0.143 n=10+10)
/10-4        38.9µs ± 8%    57.1µs ± 2%   +46.91%  (p=0.000 n=10+10)
/100-4        370µs ± 8%     581µs ± 6%   +56.81%  (p=0.000 n=10+10)
/1000-4      3.85ms ±10%    4.79ms ± 4%   +24.33%  (p=0.000 n=10+10)
/10000-4     41.0ms ±12%    54.1ms ± 3%   +32.02%  (p=0.000 n=10+10)
/100000-4     438ms ± 7%     721ms ± 3%   +64.59%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    54.4kB ± 0%  +240.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     736kB ± 0%  +360.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    6.48MB ± 0%  +304.79%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    94.6MB ± 0%  +214.13%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     789MB ± 0%  +146.74%  (p=0.000 n=9+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.8k ± 0%    +8.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.40%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.46%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.74%  (p=0.000 n=10+9)
```

deque vs phf - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.3.txt testdata/BenchmarkRefillFullPhfQueue.txt
name       old time/op    new time/op    delta
/1-4         3.55µs ± 1%    3.94µs ± 2%  +11.02%  (p=0.000 n=9+9)
/10-4        35.2µs ± 1%    39.2µs ± 3%  +11.37%  (p=0.000 n=10+9)
/100-4        344µs ± 1%     384µs ± 5%  +11.84%  (p=0.000 n=9+10)
/1000-4      3.41ms ± 1%    3.86ms ± 8%  +13.01%  (p=0.000 n=10+10)
/10000-4     40.6ms ±24%    38.8ms ± 2%     ~     (p=0.481 n=9+8)
/100000-4     451ms ±15%     707ms ± 2%  +56.81%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.87%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     632MB ± 0%  +97.58%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.76%  (p=0.000 n=10+10)
```

deque vs phf - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowIncreasePhfQueue.txt
name        old time/op    new time/op    delta
/1-4           241ns ±14%     215ns ± 7%   -10.68%  (p=0.022 n=10+10)
/10-4         1.32µs ±10%    1.36µs ± 7%      ~     (p=0.353 n=10+10)
/100-4        9.07µs ±11%    9.66µs ± 5%    +6.43%  (p=0.015 n=10+10)
/1000-4       78.2µs ±10%    83.0µs ± 2%    +6.07%  (p=0.028 n=10+9)
/10000-4       805µs ± 7%     956µs ± 2%   +18.78%  (p=0.000 n=9+10)
/100000-4     8.74ms ±12%   11.87ms ± 6%   +35.90%  (p=0.000 n=10+10)
/1000000-4    92.0ms ± 9%   129.3ms ± 4%   +40.52%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      128B ± 0%   -38.46%  (p=0.000 n=10+10)
/10-4         1.78kB ± 0%    0.99kB ± 0%   -44.14%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%    9.25kB ± 0%    +5.09%  (p=0.000 n=10+10)
/1000-4       54.2kB ± 0%    81.1kB ± 0%   +49.62%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%    1106kB ± 0%  +127.07%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    9.49MB ± 0%   +96.84%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    82.3MB ± 0%   +70.84%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowDecreasePhfQueue.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 7%    37.1ns ± 4%  -5.33%  (p=0.003 n=9+10)
/10-4          391ns ±11%     372ns ± 2%    ~     (p=0.234 n=10+9)
/100-4        3.71µs ± 8%    3.62µs ± 3%    ~     (p=0.190 n=10+10)
/1000-4       36.1µs ± 7%    36.4µs ± 3%    ~     (p=0.481 n=10+10)
/10000-4       375µs ±10%     368µs ± 2%    ~     (p=0.684 n=10+10)
/100000-4     3.68ms ± 8%    3.70ms ± 4%    ~     (p=0.853 n=10+10)
/1000000-4    37.0ms ± 9%    37.4ms ± 3%    ~     (p=1.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

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
benchstat testdata/BenchmarkStableDequeQueuev1.0.3.txt testdata/BenchmarkStablePhfQueue.txt
name        old time/op    new time/op    delta
/1-4          37.2ns ± 9%    37.4ns ± 3%    ~     (p=0.780 n=10+10)
/10-4          369ns ± 8%     379ns ± 5%    ~     (p=0.483 n=10+9)
/100-4        3.52µs ± 8%    3.67µs ± 2%    ~     (p=0.063 n=10+10)
/1000-4       36.9µs ± 7%    36.2µs ± 3%    ~     (p=0.079 n=9+10)
/10000-4       353µs ± 9%     371µs ± 4%    ~     (p=0.052 n=10+10)
/100000-4     3.60ms ± 8%    3.71ms ± 3%    ~     (p=0.280 n=10+10)
/1000000-4    35.5ms ± 7%    37.5ms ± 4%  +5.83%  (p=0.007 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

#### deque vs [phf](https://github.com/phf/go-queue) - LIFO stack
deque vs phf - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStackv1.0.3.txt testdata/BenchmarkRefillPhfStack.txt
name       old time/op    new time/op    delta
/1-4         3.91µs ± 8%    3.93µs ± 3%      ~     (p=0.853 n=10+10)
/10-4        36.6µs ± 8%    57.8µs ± 3%   +58.14%  (p=0.000 n=10+9)
/100-4        372µs ± 9%     557µs ± 2%   +49.87%  (p=0.000 n=10+10)
/1000-4      3.61ms ± 7%    4.72ms ± 2%   +30.60%  (p=0.000 n=10+10)
/10000-4     42.0ms ± 7%    54.4ms ± 3%   +29.36%  (p=0.000 n=10+10)
/100000-4     444ms ±14%     750ms ±10%   +68.91%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    54.4kB ± 0%  +240.00%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     736kB ± 0%  +360.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.48MB ± 0%  +304.79%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    94.6MB ± 0%  +214.40%  (p=0.000 n=8+9)
/100000-4     320MB ± 0%     789MB ± 0%  +146.62%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.20k ± 0%   +20.00%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.8k ± 0%    +8.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.40%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.46%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.74%  (p=0.000 n=9+10)
```

deque vs phf - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.3.txt testdata/BenchmarkRefillFullPhfStack.txt
name       old time/op    new time/op    delta
/1-4         3.85µs ± 9%    3.87µs ± 6%     ~     (p=0.859 n=10+9)
/10-4        38.2µs ± 9%    37.8µs ± 2%     ~     (p=0.684 n=10+10)
/100-4        378µs ± 8%     372µs ± 4%     ~     (p=0.280 n=10+10)
/1000-4      3.76ms ± 8%    3.69ms ± 2%     ~     (p=0.353 n=10+10)
/10000-4     42.7ms ±10%    38.3ms ± 3%  -10.25%  (p=0.000 n=10+10)
/100000-4     446ms ±11%     733ms ± 8%  +64.19%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.83%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     632MB ± 0%  +97.48%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.76%  (p=0.000 n=10+10)
```

deque vs phf - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.3.txt testdata/BenchmarkSlowIncreasePhfStack.txt
name        old time/op    new time/op    delta
/1-4           226ns ±10%     212ns ± 2%    -6.18%  (p=0.001 n=10+9)
/10-4         1.06µs ± 7%    1.27µs ± 4%   +20.58%  (p=0.000 n=10+10)
/100-4        8.71µs ±13%    9.49µs ± 6%    +8.86%  (p=0.002 n=10+10)
/1000-4       75.5µs ± 9%    81.5µs ± 2%    +7.85%  (p=0.000 n=10+10)
/10000-4       731µs ± 6%     943µs ± 3%   +28.97%  (p=0.000 n=10+10)
/100000-4     8.81ms ±10%   11.64ms ± 5%   +32.11%  (p=0.000 n=10+9)
/1000000-4    94.4ms ±11%   138.1ms ±13%   +46.30%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      128B ± 0%   -38.46%  (p=0.000 n=10+10)
/10-4           752B ± 0%      992B ± 0%   +31.91%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%    9.25kB ± 0%    +5.09%  (p=0.000 n=10+10)
/1000-4       50.0kB ± 0%    81.1kB ± 0%   +62.01%  (p=0.000 n=10+10)
/10000-4       483kB ± 0%    1106kB ± 0%  +129.02%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    9.49MB ± 0%   +96.84%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    82.3MB ± 0%   +70.84%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.3.txt testdata/BenchmarkSlowDecreasePhfStack.txt
name        old time/op    new time/op    delta
/1-4          37.3ns ± 7%    36.7ns ± 3%   ~     (p=0.674 n=10+9)
/10-4          386ns ± 9%     387ns ± 4%   ~     (p=0.723 n=10+10)
/100-4        3.76µs ± 8%    3.79µs ± 5%   ~     (p=0.869 n=10+10)
/1000-4       37.0µs ± 6%    37.4µs ± 4%   ~     (p=0.739 n=10+10)
/10000-4       377µs ± 8%     378µs ± 4%   ~     (p=0.842 n=10+9)
/100000-4     3.72ms ± 9%    3.77ms ± 6%   ~     (p=0.447 n=10+9)
/1000000-4    38.1ms ± 9%    37.1ms ± 1%   ~     (p=0.055 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%   ~     (all equal)
/10-4           160B ± 0%      160B ± 0%   ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%   ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%   ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%   ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%   ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%   ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%   ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%   ~     (all equal)
/100-4           100 ± 0%       100 ± 0%   ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%   ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%   ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%   ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%   ~     (all equal)
```

deque vs phf - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStackv1.0.3.txt testdata/BenchmarkStablePhfStack.txt
name        old time/op    new time/op    delta
/1-4          35.6ns ± 7%    39.2ns ±13%  +10.26%  (p=0.004 n=10+10)
/10-4          362ns ± 8%     370ns ± 8%     ~     (p=0.684 n=10+10)
/100-4        3.54µs ± 7%    3.84µs ± 6%   +8.47%  (p=0.002 n=10+10)
/1000-4       35.9µs ±10%    37.0µs ± 5%     ~     (p=0.190 n=10+10)
/10000-4       360µs ±10%     365µs ± 6%     ~     (p=0.353 n=10+10)
/100000-4     3.53ms ± 6%    3.57ms ± 8%     ~     (p=0.796 n=10+10)
/1000000-4    34.4ms ± 7%    35.5ms ± 2%   +3.01%  (p=0.035 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

#### deque vs [gammazero](https://github.com/gammazero/deque) - FIFO queue
deque vs gammazero - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueuev1.0.3.txt testdata/BenchmarkRefillGammazeroQueue.txt
name       old time/op    new time/op    delta
/1-4         3.95µs ± 7%    3.85µs ± 1%      ~     (p=0.138 n=10+10)
/10-4        38.9µs ± 8%    36.3µs ± 1%    -6.72%  (p=0.001 n=10+10)
/100-4        370µs ± 8%     511µs ± 2%   +37.99%  (p=0.000 n=10+10)
/1000-4      3.85ms ±10%    4.50ms ± 1%   +16.76%  (p=0.000 n=10+10)
/10000-4     41.0ms ±12%    51.8ms ± 2%   +26.37%  (p=0.000 n=10+10)
/100000-4     438ms ± 7%     695ms ± 1%   +58.53%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     698kB ± 0%  +336.00%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    6.44MB ± 0%  +302.39%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    94.6MB ± 0%  +214.00%  (p=0.000 n=10+8)
/100000-4     320MB ± 0%     789MB ± 0%  +146.73%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.6k ± 0%    +6.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.20%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.48%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.74%  (p=0.000 n=10+10)
```

deque vs gammazero - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.3.txt testdata/BenchmarkRefillFullGammazeroQueue.txt
name       old time/op    new time/op    delta
/1-4         3.55µs ± 1%    3.64µs ± 2%   +2.53%  (p=0.000 n=9+10)
/10-4        35.2µs ± 1%    38.1µs ± 9%   +8.29%  (p=0.000 n=10+10)
/100-4        344µs ± 1%     365µs ± 2%   +6.16%  (p=0.000 n=9+10)
/1000-4      3.41ms ± 1%    3.58ms ± 4%   +4.90%  (p=0.000 n=10+10)
/10000-4     40.6ms ±24%    37.1ms ± 3%   -8.80%  (p=0.004 n=9+9)
/100000-4     451ms ±15%     818ms ±34%  +81.51%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.87%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     632MB ± 0%  +97.58%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.76%  (p=0.000 n=10+10)
```

deque vs gammazero - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowIncreaseGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4           241ns ±14%     221ns ± 4%      ~     (p=0.068 n=10+9)
/10-4         1.32µs ±10%    0.85µs ± 5%   -35.97%  (p=0.000 n=10+9)
/100-4        9.07µs ±11%    8.91µs ± 8%      ~     (p=0.515 n=10+8)
/1000-4       78.2µs ±10%    80.4µs ± 4%      ~     (p=0.529 n=10+10)
/10000-4       805µs ± 7%     924µs ± 7%   +14.69%  (p=0.000 n=9+9)
/100000-4     8.74ms ±12%   10.97ms ± 5%   +25.59%  (p=0.000 n=10+9)
/1000000-4    92.0ms ± 9%   122.9ms ± 9%   +33.65%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      336B ± 0%   +61.54%  (p=0.000 n=10+10)
/10-4         1.78kB ± 0%    0.62kB ± 0%   -64.86%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%    8.88kB ± 0%    +0.91%  (p=0.000 n=10+10)
/1000-4       54.2kB ± 0%    80.7kB ± 0%   +48.94%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%    1106kB ± 0%  +127.00%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    9.49MB ± 0%   +96.84%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    82.3MB ± 0%   +70.84%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowDecreaseGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 7%    35.1ns ± 1%  -10.31%  (p=0.000 n=9+10)
/10-4          391ns ±11%     355ns ± 1%   -9.23%  (p=0.001 n=10+10)
/100-4        3.71µs ± 8%    3.65µs ±16%     ~     (p=0.388 n=10+9)
/1000-4       36.1µs ± 7%    36.4µs ± 3%     ~     (p=0.481 n=10+10)
/10000-4       375µs ±10%     363µs ± 4%     ~     (p=0.356 n=10+9)
/100000-4     3.68ms ± 8%    3.61ms ± 4%     ~     (p=0.393 n=10+10)
/1000000-4    37.0ms ± 9%    35.0ms ± 3%     ~     (p=0.122 n=10+8)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

deque vs gammazero - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueuev1.0.3.txt testdata/BenchmarkStableGammazeroQueue.txt
name        old time/op    new time/op    delta
/1-4          37.2ns ± 9%    35.5ns ± 3%    ~     (p=0.210 n=10+10)
/10-4          369ns ± 8%     370ns ± 3%    ~     (p=0.889 n=10+9)
/100-4        3.52µs ± 8%    3.77µs ± 8%  +7.02%  (p=0.022 n=10+9)
/1000-4       36.9µs ± 7%    36.6µs ± 6%    ~     (p=0.489 n=9+9)
/10000-4       353µs ± 9%     357µs ± 4%    ~     (p=0.497 n=10+9)
/100000-4     3.60ms ± 8%    3.59ms ± 4%    ~     (p=0.842 n=10+9)
/1000000-4    35.5ms ± 7%    38.8ms ±25%    ~     (p=0.247 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

#### deque vs [gammazero](https://github.com/gammazero/deque) - LIFO stack
deque vs gammazero - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStackv1.0.3.txt testdata/BenchmarkRefillGammazeroStack.txt
name       old time/op    new time/op    delta
/1-4         3.91µs ± 8%    3.67µs ± 1%    -6.04%  (p=0.022 n=10+10)
/10-4        36.6µs ± 8%    36.0µs ± 2%      ~     (p=0.971 n=10+10)
/100-4        372µs ± 9%     506µs ± 1%   +36.13%  (p=0.000 n=10+10)
/1000-4      3.61ms ± 7%    4.45ms ± 1%   +23.02%  (p=0.000 n=10+10)
/10000-4     42.0ms ± 7%    51.1ms ± 1%   +21.54%  (p=0.000 n=10+10)
/100000-4     444ms ±14%     694ms ± 2%   +56.25%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     698kB ± 0%  +336.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    6.44MB ± 0%  +302.39%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    94.6MB ± 0%  +214.27%  (p=0.000 n=8+10)
/100000-4     320MB ± 0%     789MB ± 0%  +146.61%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.6k ± 0%    +6.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      101k ± 0%    +1.20%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.00M ± 0%    -0.48%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%    -0.74%  (p=0.000 n=9+10)
```

deque vs gammazero - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.3.txt testdata/BenchmarkRefillFullGammazeroStack.txt
name       old time/op    new time/op    delta
/1-4         3.85µs ± 9%    3.77µs ± 7%     ~     (p=0.243 n=10+9)
/10-4        38.2µs ± 9%    38.6µs ± 5%     ~     (p=0.356 n=10+9)
/100-4        378µs ± 8%     365µs ± 3%     ~     (p=0.143 n=10+10)
/1000-4      3.76ms ± 8%    3.57ms ± 1%     ~     (p=0.156 n=10+9)
/10000-4     42.7ms ±10%    37.5ms ± 7%  -12.14%  (p=0.000 n=10+9)
/100000-4     446ms ±11%     691ms ± 3%  +54.74%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.83%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     632MB ± 0%  +97.48%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.76%  (p=0.000 n=10+10)
```

deque vs gammazero - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.3.txt testdata/BenchmarkSlowIncreaseGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4           226ns ±10%     216ns ± 9%      ~     (p=0.251 n=10+9)
/10-4         1.06µs ± 7%    0.85µs ± 6%   -19.91%  (p=0.000 n=10+10)
/100-4        8.71µs ±13%    8.69µs ± 5%      ~     (p=0.971 n=10+10)
/1000-4       75.5µs ± 9%    80.1µs ± 5%    +5.99%  (p=0.008 n=10+9)
/10000-4       731µs ± 6%     944µs ± 6%   +29.11%  (p=0.000 n=10+10)
/100000-4     8.81ms ±10%   11.52ms ± 8%   +30.66%  (p=0.000 n=10+9)
/1000000-4    94.4ms ±11%   124.7ms ± 3%   +32.07%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%      336B ± 0%   +61.54%  (p=0.000 n=10+10)
/10-4           752B ± 0%      624B ± 0%   -17.02%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%    8.88kB ± 0%    +0.91%  (p=0.000 n=10+10)
/1000-4       50.0kB ± 0%    80.7kB ± 0%   +61.27%  (p=0.000 n=10+10)
/10000-4       483kB ± 0%    1106kB ± 0%  +128.94%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    9.49MB ± 0%   +96.84%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    82.3MB ± 0%   +70.84%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.3.txt testdata/BenchmarkSlowDecreaseGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4          37.3ns ± 7%    39.0ns ±24%   ~     (p=0.400 n=10+9)
/10-4          386ns ± 9%     430ns ±34%   ~     (p=0.101 n=10+10)
/100-4        3.76µs ± 8%    3.61µs ± 3%   ~     (p=0.165 n=10+10)
/1000-4       37.0µs ± 6%    36.9µs ± 4%   ~     (p=0.842 n=10+9)
/10000-4       377µs ± 8%     371µs ± 3%   ~     (p=0.497 n=10+9)
/100000-4     3.72ms ± 9%    3.70ms ± 4%   ~     (p=0.853 n=10+10)
/1000000-4    38.1ms ± 9%    37.0ms ± 3%   ~     (p=0.105 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%   ~     (all equal)
/10-4           160B ± 0%      160B ± 0%   ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%   ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%   ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%   ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%   ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%   ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%   ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%   ~     (all equal)
/100-4           100 ± 0%       100 ± 0%   ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%   ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%   ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%   ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%   ~     (all equal)
```

deque vs gammazero - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStackv1.0.3.txt testdata/BenchmarkStableGammazeroStack.txt
name        old time/op    new time/op    delta
/1-4          35.6ns ± 7%    37.2ns ± 7%   ~     (p=0.079 n=10+10)
/10-4          362ns ± 8%     358ns ± 3%   ~     (p=0.343 n=10+10)
/100-4        3.54µs ± 7%    3.55µs ± 7%   ~     (p=0.842 n=10+9)
/1000-4       35.9µs ±10%    35.5µs ± 4%   ~     (p=0.280 n=10+10)
/10000-4       360µs ±10%     351µs ± 4%   ~     (p=0.393 n=10+10)
/100000-4     3.53ms ± 6%    3.67ms ±13%   ~     (p=0.113 n=10+9)
/1000000-4    34.4ms ± 7%    35.5ms ± 5%   ~     (p=0.165 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%   ~     (all equal)
/10-4           160B ± 0%      160B ± 0%   ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%   ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%   ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%   ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%   ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%   ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%   ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%   ~     (all equal)
/100-4           100 ± 0%       100 ± 0%   ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%   ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%   ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%   ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%   ~     (all equal)
```

#### deque vs [juju](https://github.com/juju/utils/blob/master/deque/deque.go) - FIFO queue
deque vs juju - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueuev1.0.3.txt testdata/BenchmarkRefillJujuQueue.txt
name       old time/op    new time/op    delta
/1-4         3.95µs ± 7%    4.07µs ± 6%      ~     (p=0.325 n=10+10)
/10-4        38.9µs ± 8%    45.4µs ± 9%   +16.74%  (p=0.000 n=10+10)
/100-4        370µs ± 8%     464µs ±11%   +25.36%  (p=0.000 n=10+10)
/1000-4      3.85ms ±10%    4.53ms ±10%   +17.50%  (p=0.000 n=10+10)
/10000-4     41.0ms ±12%    49.5ms ±12%   +20.88%  (p=0.000 n=10+10)
/100000-4     438ms ± 7%     493ms ±10%   +12.43%  (p=0.003 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    29.8kB ± 0%   +86.25%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     326kB ± 0%  +103.50%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    3.31MB ± 0%  +106.94%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    33.2MB ± 0%   +10.31%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     332MB ± 0%    +3.95%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.04k ± 0%    +3.70%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.4k ± 0%    +4.50%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      105k ± 0%    +4.65%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.05M ± 0%    +3.97%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.5M ± 0%    +3.88%  (p=0.000 n=10+10)
```

deque vs juju - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.3.txt testdata/BenchmarkRefillFullJujuQueue.txt
name       old time/op    new time/op    delta
/1-4         3.55µs ± 1%    4.31µs ± 1%   +21.37%  (p=0.000 n=9+10)
/10-4        35.2µs ± 1%    43.7µs ± 1%   +24.14%  (p=0.000 n=10+10)
/100-4        344µs ± 1%     421µs ± 1%   +22.47%  (p=0.000 n=9+10)
/1000-4      3.41ms ± 1%    4.21ms ± 1%   +23.25%  (p=0.000 n=10+10)
/10000-4     40.6ms ±24%    41.5ms ± 1%      ~     (p=0.258 n=9+9)
/100000-4     451ms ±15%     428ms ± 1%      ~     (p=0.053 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    3.33kB ± 0%  +107.81%  (p=0.000 n=10+10)
/10-4        16.0kB ± 0%    33.2kB ± 0%  +107.81%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     332kB ± 0%  +107.81%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    3.33MB ± 0%  +107.81%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    33.2MB ± 0%   +10.41%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     332MB ± 0%    +3.97%  (p=0.000 n=9+9)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       104 ± 0%    +4.00%  (p=0.000 n=10+10)
/10-4         1.00k ± 0%     1.05k ± 0%    +4.60%  (p=0.000 n=10+10)
/100-4        10.0k ± 0%     10.5k ± 0%    +4.68%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      105k ± 0%    +4.69%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.05M ± 0%    +3.98%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.5M ± 0%    +3.89%  (p=0.000 n=10+10)
```

deque vs juju - FIFO queue - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowIncreaseJujuQueue.txt
name        old time/op    new time/op    delta
/1-4           241ns ±14%     549ns ±19%  +128.00%  (p=0.000 n=10+10)
/10-4         1.32µs ±10%    1.22µs ±10%    -7.38%  (p=0.043 n=10+10)
/100-4        9.07µs ±11%    9.57µs ±17%      ~     (p=0.165 n=10+10)
/1000-4       78.2µs ±10%    93.7µs ±15%   +19.82%  (p=0.002 n=10+10)
/10000-4       805µs ± 7%     878µs ±11%    +9.01%  (p=0.004 n=9+10)
/100000-4     8.74ms ±12%   10.67ms ± 6%   +22.10%  (p=0.000 n=10+9)
/1000000-4    92.0ms ± 9%   118.5ms ±13%   +28.80%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%     1216B ± 0%  +484.62%  (p=0.000 n=10+10)
/10-4         1.78kB ± 0%    1.50kB ± 0%   -15.32%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%    7.70kB ± 0%   -12.55%  (p=0.000 n=10+10)
/1000-4       54.2kB ± 0%    67.4kB ± 0%   +24.42%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%     666kB ± 0%   +36.62%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    6.65MB ± 0%   +37.94%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    66.5MB ± 0%   +37.99%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      6.00 ± 0%   +20.00%  (p=0.000 n=10+10)
/10-4           25.0 ± 0%      24.0 ± 0%    -4.00%  (p=0.000 n=10+10)
/100-4           207 ± 0%       213 ± 0%    +2.90%  (p=0.000 n=10+10)
/1000-4        2.02k ± 0%     2.10k ± 0%    +4.07%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.9k ± 0%    +4.26%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      209k ± 0%    +4.28%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.09M ± 0%    +4.28%  (p=0.000 n=10+10)
```

deque vs juju - FIFO queue - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowDecreaseJujuQueue.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 7%    38.9ns ± 8%    ~     (p=0.951 n=9+10)
/10-4          391ns ±11%     394ns ± 9%    ~     (p=0.812 n=10+10)
/100-4        3.71µs ± 8%    3.95µs ± 8%    ~     (p=0.075 n=10+10)
/1000-4       36.1µs ± 7%    38.6µs ± 9%  +6.83%  (p=0.019 n=10+10)
/10000-4       375µs ±10%     406µs ±14%  +8.06%  (p=0.023 n=10+10)
/100000-4     3.68ms ± 8%    3.83ms ± 8%    ~     (p=0.247 n=10+10)
/1000000-4    37.0ms ± 9%    38.6ms ±14%    ~     (p=0.315 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%    ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%    ~     (all equal)
/100-4           100 ± 0%       100 ± 0%    ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%    ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%    ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%    ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%    ~     (all equal)
```

deque vs juju - FIFO queue - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeQueuev1.0.3.txt testdata/BenchmarkStableJujuQueue.txt
name        old time/op    new time/op    delta
/1-4          37.2ns ± 9%    46.0ns ±11%   +23.65%  (p=0.000 n=10+10)
/10-4          369ns ± 8%     488ns ±15%   +32.29%  (p=0.000 n=10+10)
/100-4        3.52µs ± 8%    4.70µs ±11%   +33.26%  (p=0.000 n=10+10)
/1000-4       36.9µs ± 7%    45.1µs ±11%   +22.42%  (p=0.000 n=9+10)
/10000-4       353µs ± 9%     453µs ±10%   +28.33%  (p=0.000 n=10+10)
/100000-4     3.60ms ± 8%    4.59ms ±12%   +27.41%  (p=0.000 n=10+10)
/1000000-4    35.5ms ± 7%    45.1ms ± 9%   +27.24%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     33.0B ± 0%  +106.25%  (p=0.000 n=10+10)
/10-4           160B ± 0%      332B ± 0%  +107.50%  (p=0.000 n=10+10)
/100-4        1.60kB ± 0%    3.33kB ± 0%  +107.81%  (p=0.000 n=10+10)
/1000-4       16.0kB ± 0%    33.2kB ± 0%  +107.81%  (p=0.000 n=10+10)
/10000-4       160kB ± 0%     332kB ± 0%  +107.81%  (p=0.000 n=10+10)
/100000-4     1.60MB ± 0%    3.33MB ± 0%  +107.81%  (p=0.000 n=10+10)
/1000000-4    16.0MB ± 0%    33.2MB ± 0%  +107.81%  (p=0.000 n=10+10)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%      ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%      ~     (all equal)
/100-4           100 ± 0%       104 ± 0%    +4.00%  (p=0.000 n=10+10)
/1000-4        1.00k ± 0%     1.05k ± 0%    +4.60%  (p=0.000 n=10+10)
/10000-4       10.0k ± 0%     10.5k ± 0%    +4.68%  (p=0.000 n=10+10)
/100000-4       100k ± 0%      105k ± 0%    +4.69%  (p=0.000 n=10+10)
/1000000-4     1.00M ± 0%     1.05M ± 0%    +4.69%  (p=0.000 n=10+10)
```

#### deque vs [juju](https://github.com/juju/utils/blob/master/deque/deque.go) - LIFO stack

deque vs juju - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStackv1.0.3.txt testdata/BenchmarkRefillJujuStack.txt
name       old time/op    new time/op    delta
/1-4         3.91µs ± 8%    4.21µs ± 9%    +7.77%  (p=0.019 n=10+10)
/10-4        36.6µs ± 8%    40.1µs ± 8%    +9.65%  (p=0.002 n=10+10)
/100-4        372µs ± 9%     504µs ±14%   +35.57%  (p=0.000 n=10+10)
/1000-4      3.61ms ± 7%    4.57ms ±10%   +26.43%  (p=0.000 n=10+10)
/10000-4     42.0ms ± 7%    48.2ms ±13%   +14.58%  (p=0.001 n=10+10)
/100000-4     444ms ±14%     490ms ±11%      ~     (p=0.063 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     381kB ± 0%  +138.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    3.37MB ± 0%  +110.40%  (p=0.000 n=10+9)
/10000-4     30.1MB ± 0%    33.2MB ± 0%   +10.41%  (p=0.000 n=8+10)
/100000-4     320MB ± 0%     332MB ± 0%    +3.90%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.6k ± 0%    +6.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      105k ± 0%    +4.80%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.05M ± 0%    +3.97%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.5M ± 0%    +3.88%  (p=0.000 n=9+10)
```

deque vs juju - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.3.txt testdata/BenchmarkRefillFullJujuStack.txt
name       old time/op    new time/op    delta
/1-4         3.85µs ± 9%    4.46µs ±19%   +15.74%  (p=0.009 n=10+10)
/10-4        38.2µs ± 9%    40.6µs ± 8%    +6.37%  (p=0.019 n=10+10)
/100-4        378µs ± 8%     517µs ± 4%   +36.60%  (p=0.000 n=10+8)
/1000-4      3.76ms ± 8%    4.71ms ±10%   +25.39%  (p=0.000 n=10+10)
/10000-4     42.7ms ±10%    48.9ms ±12%   +14.52%  (p=0.000 n=10+10)
/100000-4     446ms ±11%     505ms ±14%   +13.18%  (p=0.015 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%      ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%      ~     (all equal)
/100-4        160kB ± 0%     381kB ± 0%  +138.00%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    3.37MB ± 0%  +110.40%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    33.2MB ± 0%   +10.41%  (p=0.000 n=10+9)
/100000-4     320MB ± 0%     333MB ± 0%    +3.94%  (p=0.000 n=10+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%      ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%      ~     (all equal)
/100-4        10.0k ± 0%     10.6k ± 0%    +6.00%  (p=0.000 n=10+10)
/1000-4        100k ± 0%      105k ± 0%    +4.80%  (p=0.000 n=10+10)
/10000-4      1.01M ± 0%     1.05M ± 0%    +3.97%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.5M ± 0%    +3.89%  (p=0.000 n=10+10)
```

deque vs juju - LIFO stack - [slow increase tests](benchmark-slow-increase_test.go)
```
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.3.txt testdata/BenchmarkSlowIncreaseJujuStack.txt
name        old time/op    new time/op    delta
/1-4           226ns ±10%     546ns ±21%  +142.02%  (p=0.000 n=10+10)
/10-4         1.06µs ± 7%    1.19µs ±14%   +12.92%  (p=0.017 n=10+10)
/100-4        8.71µs ±13%    9.92µs ±12%   +13.86%  (p=0.002 n=10+10)
/1000-4       75.5µs ± 9%    92.3µs ±15%   +22.14%  (p=0.000 n=10+10)
/10000-4       731µs ± 6%     882µs ± 9%   +20.59%  (p=0.000 n=10+10)
/100000-4     8.81ms ±10%   10.95ms ±12%   +24.28%  (p=0.000 n=10+10)
/1000000-4    94.4ms ±11%   129.9ms ±10%   +37.59%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%     1216B ± 0%  +484.62%  (p=0.000 n=10+10)
/10-4           752B ± 0%     1504B ± 0%  +100.00%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%    8.80kB ± 0%      ~     (all equal)
/1000-4       50.0kB ± 0%    68.5kB ± 0%   +36.94%  (p=0.000 n=10+10)
/10000-4       483kB ± 0%     666kB ± 0%   +37.79%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    6.65MB ± 0%   +37.94%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    66.5MB ± 0%   +37.99%  (p=0.000 n=10+9)

name        old allocs/op  new allocs/op  delta
/1-4            5.00 ± 0%      6.00 ± 0%   +20.00%  (p=0.000 n=10+10)
/10-4           24.0 ± 0%      24.0 ± 0%      ~     (all equal)
/100-4           207 ± 0%       216 ± 0%    +4.35%  (p=0.000 n=10+10)
/1000-4        2.01k ± 0%     2.10k ± 0%    +4.32%  (p=0.000 n=10+10)
/10000-4       20.1k ± 0%     20.9k ± 0%    +4.27%  (p=0.000 n=10+10)
/100000-4       201k ± 0%      209k ± 0%    +4.28%  (p=0.000 n=10+10)
/1000000-4     2.01M ± 0%     2.09M ± 0%    +4.28%  (p=0.000 n=10+10)
```

deque vs juju - LIFO stack - [slow decrease tests](benchmark-slow-decrease_test.go)
```
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.3.txt testdata/BenchmarkSlowDecreaseJujuStack.txt
name        old time/op    new time/op    delta
/1-4          37.3ns ± 7%    39.0ns ± 8%   ~     (p=0.093 n=10+10)
/10-4          386ns ± 9%     412ns ±14%   ~     (p=0.066 n=10+10)
/100-4        3.76µs ± 8%    3.87µs ± 7%   ~     (p=0.247 n=10+10)
/1000-4       37.0µs ± 6%    37.7µs ± 8%   ~     (p=0.631 n=10+10)
/10000-4       377µs ± 8%     379µs ± 8%   ~     (p=0.739 n=10+10)
/100000-4     3.72ms ± 9%    3.88ms ± 9%   ~     (p=0.143 n=10+10)
/1000000-4    38.1ms ± 9%    38.6ms ± 8%   ~     (p=0.481 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%   ~     (all equal)
/10-4           160B ± 0%      160B ± 0%   ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%   ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%   ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%   ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%   ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%   ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%   ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%   ~     (all equal)
/100-4           100 ± 0%       100 ± 0%   ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%   ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%   ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%   ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%   ~     (all equal)
```

deque vs juju - LIFO stack - [stable tests](benchmark-stable_test.go)
```
benchstat testdata/BenchmarkStableDequeStackv1.0.3.txt testdata/BenchmarkStableJujuStack.txt
name        old time/op    new time/op    delta
/1-4          35.6ns ± 7%    38.1ns ± 9%   +7.23%  (p=0.020 n=10+10)
/10-4          362ns ± 8%     406ns ±13%  +12.06%  (p=0.006 n=10+10)
/100-4        3.54µs ± 7%    3.81µs ± 8%   +7.49%  (p=0.001 n=10+10)
/1000-4       35.9µs ±10%    39.4µs ±13%   +9.81%  (p=0.011 n=10+10)
/10000-4       360µs ±10%     388µs ±15%   +7.83%  (p=0.029 n=10+10)
/100000-4     3.53ms ± 6%    3.80ms ±10%   +7.78%  (p=0.023 n=10+10)
/1000000-4    34.4ms ± 7%    37.3ms ± 9%   +8.25%  (p=0.009 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

#### deque vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - FIFO queue
deque vs cookiejar - FIFO queue - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeQueuev1.0.3.txt testdata/BenchmarkRefillCookiejarQueue.txt
name       old time/op    new time/op    delta
/1-4         3.95µs ± 7%    3.51µs ± 5%  -11.36%  (p=0.000 n=10+10)
/10-4        38.9µs ± 8%    32.7µs ± 6%  -15.82%  (p=0.000 n=10+10)
/100-4        370µs ± 8%     311µs ± 2%  -16.05%  (p=0.000 n=10+10)
/1000-4      3.85ms ±10%    3.05ms ± 3%  -20.75%  (p=0.000 n=10+10)
/10000-4     41.0ms ±12%    31.3ms ± 4%  -23.56%  (p=0.000 n=10+10)
/100000-4     438ms ± 7%     365ms ± 7%  -16.64%  (p=0.000 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%   +0.01%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     160kB ± 0%   +0.02%  (p=0.000 n=9+10)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   +0.01%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.86%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     161MB ± 0%  -49.79%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=10+10)
```

deque vs cookiejar - FIFO queue - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeQueuev1.0.3.txt testdata/BenchmarkRefillFullCookiejarQueue.txt
name       old time/op    new time/op    delta
/1-4         3.55µs ± 1%    3.13µs ± 3%  -11.78%  (p=0.000 n=9+9)
/10-4        35.2µs ± 1%    32.2µs ± 7%   -8.38%  (p=0.000 n=10+10)
/100-4        344µs ± 1%     316µs ± 3%   -7.92%  (p=0.000 n=9+9)
/1000-4      3.41ms ± 1%    3.13ms ± 0%   -8.29%  (p=0.000 n=10+7)
/10000-4     40.6ms ±24%    31.6ms ± 3%  -22.26%  (p=0.000 n=9+10)
/100000-4     451ms ±15%     366ms ± 7%  -18.68%  (p=0.000 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.87%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     160MB ± 0%  -49.97%  (p=0.000 n=9+10)

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
benchstat testdata/BenchmarkSlowIncreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowIncreaseCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4           241ns ±14%   10208ns ± 2%   +4140.92%  (p=0.000 n=10+9)
/10-4         1.32µs ±10%   10.86µs ± 2%    +721.39%  (p=0.000 n=10+10)
/100-4        9.07µs ±11%   16.48µs ± 4%     +81.64%  (p=0.000 n=10+10)
/1000-4       78.2µs ±10%    69.0µs ± 2%     -11.79%  (p=0.000 n=10+10)
/10000-4       805µs ± 7%     679µs ±10%     -15.67%  (p=0.000 n=9+10)
/100000-4     8.74ms ±12%    8.12ms ± 5%      -7.09%  (p=0.023 n=10+10)
/1000000-4    92.0ms ± 9%    83.7ms ± 3%      -9.05%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%    65712B ± 0%  +31492.31%  (p=0.000 n=10+10)
/10-4         1.78kB ± 0%   66.00kB ± 0%   +3616.22%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%   68.88kB ± 0%    +682.73%  (p=0.000 n=10+10)
/1000-4       54.2kB ± 0%    97.7kB ± 0%     +80.30%  (p=0.000 n=10+10)
/10000-4       487kB ± 0%     583kB ± 0%     +19.56%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    4.91MB ± 0%      +1.89%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    48.9MB ± 0%      +1.47%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkSlowDecreaseDequeQueuev1.0.3.txt testdata/BenchmarkSlowDecreaseCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4          39.2ns ± 7%    36.1ns ±12%   -7.83%  (p=0.023 n=9+10)
/10-4          391ns ±11%     341ns ± 2%  -12.78%  (p=0.000 n=10+9)
/100-4        3.71µs ± 8%    3.32µs ± 2%  -10.49%  (p=0.000 n=10+10)
/1000-4       36.1µs ± 7%    33.4µs ± 2%   -7.51%  (p=0.000 n=10+10)
/10000-4       375µs ±10%     332µs ± 1%  -11.61%  (p=0.000 n=10+9)
/100000-4     3.68ms ± 8%    3.30ms ± 2%  -10.17%  (p=0.000 n=10+9)
/1000000-4    37.0ms ± 9%    33.5ms ± 6%   -9.51%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

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
benchstat testdata/BenchmarkStableDequeQueuev1.0.3.txt testdata/BenchmarkStableCookiejarQueue.txt
name        old time/op    new time/op    delta
/1-4          37.2ns ± 9%    30.8ns ± 3%  -17.19%  (p=0.000 n=10+10)
/10-4          369ns ± 8%     309ns ± 4%  -16.33%  (p=0.000 n=10+10)
/100-4        3.52µs ± 8%    3.03µs ± 2%  -13.99%  (p=0.000 n=10+10)
/1000-4       36.9µs ± 7%    30.1µs ± 3%  -18.37%  (p=0.000 n=9+9)
/10000-4       353µs ± 9%     311µs ±12%  -11.70%  (p=0.000 n=10+10)
/100000-4     3.60ms ± 8%    3.05ms ± 4%  -15.21%  (p=0.000 n=10+10)
/1000000-4    35.5ms ± 7%    30.3ms ± 5%  -14.57%  (p=0.000 n=10+9)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```

#### deque vs [cookiejar](https://github.com/karalabe/cookiejar/blob/master/collections/deque/deque.go) - LIFO stack
deque vs cookiejar - LIFO stack - [refill tests](benchmark-refill_test.go)
```
benchstat testdata/BenchmarkRefillDequeStackv1.0.3.txt testdata/BenchmarkRefillCookiejarStack.txt
name       old time/op    new time/op    delta
/1-4         3.91µs ± 8%    3.63µs ± 6%   -7.22%  (p=0.002 n=10+10)
/10-4        36.6µs ± 8%    33.9µs ± 5%   -7.31%  (p=0.002 n=10+10)
/100-4        372µs ± 9%     330µs ±11%  -11.18%  (p=0.000 n=10+10)
/1000-4      3.61ms ± 7%    3.38ms ± 6%   -6.46%  (p=0.009 n=10+8)
/10000-4     42.0ms ± 7%    36.5ms ±16%  -13.14%  (p=0.000 n=10+10)
/100000-4     444ms ±14%     400ms ±11%   -9.87%  (p=0.011 n=10+10)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%   +0.01%  (p=0.000 n=10+10)
/100-4        160kB ± 0%     160kB ± 0%   +0.01%  (p=0.000 n=10+10)
/1000-4      1.60MB ± 0%    1.60MB ± 0%   +0.01%  (p=0.000 n=10+10)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.81%  (p=0.000 n=8+9)
/100000-4     320MB ± 0%     161MB ± 0%  -49.82%  (p=0.000 n=9+10)

name       old allocs/op  new allocs/op  delta
/1-4            100 ± 0%       100 ± 0%     ~     (all equal)
/10-4         1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/100-4        10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/1000-4        100k ± 0%      100k ± 0%     ~     (all equal)
/10000-4      1.01M ± 0%     1.00M ± 0%   -0.68%  (p=0.000 n=10+10)
/100000-4     10.1M ± 0%     10.0M ± 0%   -0.77%  (p=0.000 n=9+10)
```

deque vs cookiejar - LIFO stack - [refill full tests](benchmark-refill-full_test.go)
```
benchstat testdata/BenchmarkRefillFullDequeStackv1.0.3.txt testdata/BenchmarkRefillFullCookiejarStack.txt
name       old time/op    new time/op    delta
/1-4         3.85µs ± 9%    3.26µs ± 7%  -15.37%  (p=0.000 n=10+10)
/10-4        38.2µs ± 9%    33.1µs ± 9%  -13.40%  (p=0.000 n=10+10)
/100-4        378µs ± 8%     305µs ± 2%  -19.36%  (p=0.000 n=10+8)
/1000-4      3.76ms ± 8%    3.18ms ± 5%  -15.40%  (p=0.000 n=10+9)
/10000-4     42.7ms ±10%    31.5ms ± 6%  -26.36%  (p=0.000 n=10+9)
/100000-4     446ms ±11%     354ms ± 3%  -20.74%  (p=0.000 n=10+9)

name       old alloc/op   new alloc/op   delta
/1-4         1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/10-4        16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/100-4        160kB ± 0%     160kB ± 0%     ~     (all equal)
/1000-4      1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/10000-4     30.1MB ± 0%    16.0MB ± 0%  -46.83%  (p=0.000 n=10+10)
/100000-4     320MB ± 0%     160MB ± 0%  -49.99%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkSlowIncreaseDequeStackv1.0.3.txt testdata/BenchmarkSlowIncreaseCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4           226ns ±10%   10375ns ± 3%   +4498.80%  (p=0.000 n=10+10)
/10-4         1.06µs ± 7%   11.12µs ± 3%    +952.26%  (p=0.000 n=10+10)
/100-4        8.71µs ±13%   16.70µs ± 4%     +91.59%  (p=0.000 n=10+9)
/1000-4       75.5µs ± 9%    71.1µs ± 3%      -5.90%  (p=0.007 n=10+10)
/10000-4       731µs ± 6%     643µs ± 2%     -12.02%  (p=0.000 n=10+10)
/100000-4     8.81ms ±10%    8.24ms ± 5%      -6.48%  (p=0.043 n=10+10)
/1000000-4    94.4ms ±11%    89.0ms ± 3%        ~     (p=0.143 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4            208B ± 0%    65712B ± 0%  +31492.31%  (p=0.000 n=10+10)
/10-4           752B ± 0%    66000B ± 0%   +8676.60%  (p=0.000 n=10+10)
/100-4        8.80kB ± 0%   68.88kB ± 0%    +682.73%  (p=0.000 n=10+10)
/1000-4       50.0kB ± 0%    97.7kB ± 0%     +95.24%  (p=0.000 n=10+10)
/10000-4       483kB ± 0%     517kB ± 0%      +7.00%  (p=0.000 n=10+10)
/100000-4     4.82MB ± 0%    4.85MB ± 0%      +0.51%  (p=0.000 n=10+10)
/1000000-4    48.2MB ± 0%    48.8MB ± 0%      +1.32%  (p=0.000 n=10+10)

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
benchstat testdata/BenchmarkSlowDecreaseDequeStackv1.0.3.txt testdata/BenchmarkSlowDecreaseCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4          37.3ns ± 7%    35.9ns ±10%    ~     (p=0.079 n=10+10)
/10-4          386ns ± 9%     378ns ±17%    ~     (p=0.423 n=10+10)
/100-4        3.76µs ± 8%    3.63µs ±13%    ~     (p=0.109 n=10+10)
/1000-4       37.0µs ± 6%    34.6µs ± 2%  -6.53%  (p=0.000 n=10+8)
/10000-4       377µs ± 8%     344µs ± 1%  -8.78%  (p=0.000 n=10+9)
/100000-4     3.72ms ± 9%    3.50ms ± 5%  -5.95%  (p=0.015 n=10+10)
/1000000-4    38.1ms ± 9%    34.4ms ± 2%  -9.89%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%    ~     (all equal)
/10-4           160B ± 0%      160B ± 0%    ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%    ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%    ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%    ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%    ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%    ~     (all equal)

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
benchstat testdata/BenchmarkStableDequeStackv1.0.3.txt testdata/BenchmarkStableCookiejarStack.txt
name        old time/op    new time/op    delta
/1-4          35.6ns ± 7%    30.8ns ± 4%  -13.44%  (p=0.000 n=10+10)
/10-4          362ns ± 8%     304ns ± 2%  -16.11%  (p=0.000 n=10+8)
/100-4        3.54µs ± 7%    2.99µs ± 2%  -15.67%  (p=0.000 n=10+9)
/1000-4       35.9µs ±10%    30.5µs ± 9%  -14.92%  (p=0.000 n=10+10)
/10000-4       360µs ±10%     300µs ± 4%  -16.76%  (p=0.000 n=10+10)
/100000-4     3.53ms ± 6%    2.94ms ± 2%  -16.84%  (p=0.000 n=10+10)
/1000000-4    34.4ms ± 7%    30.1ms ± 3%  -12.55%  (p=0.000 n=10+10)

name        old alloc/op   new alloc/op   delta
/1-4           16.0B ± 0%     16.0B ± 0%     ~     (all equal)
/10-4           160B ± 0%      160B ± 0%     ~     (all equal)
/100-4        1.60kB ± 0%    1.60kB ± 0%     ~     (all equal)
/1000-4       16.0kB ± 0%    16.0kB ± 0%     ~     (all equal)
/10000-4       160kB ± 0%     160kB ± 0%     ~     (all equal)
/100000-4     1.60MB ± 0%    1.60MB ± 0%     ~     (all equal)
/1000000-4    16.0MB ± 0%    16.0MB ± 0%     ~     (all equal)

name        old allocs/op  new allocs/op  delta
/1-4            1.00 ± 0%      1.00 ± 0%     ~     (all equal)
/10-4           10.0 ± 0%      10.0 ± 0%     ~     (all equal)
/100-4           100 ± 0%       100 ± 0%     ~     (all equal)
/1000-4        1.00k ± 0%     1.00k ± 0%     ~     (all equal)
/10000-4       10.0k ± 0%     10.0k ± 0%     ~     (all equal)
/100000-4       100k ± 0%      100k ± 0%     ~     (all equal)
/1000000-4     1.00M ± 0%     1.00M ± 0%     ~     (all equal)
```
