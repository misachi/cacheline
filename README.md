# cacheline

Simple test for the affects of having cacheline aligned data structures implemented in Golang.

Includes data structure with 2 integer counters being updated in separate threads.

One data structure contains the counters on the same cacheline. The other data structure attempts to put the 2 counters on separate cachelines.

The expected result is the second structure should perform better that the setup with 2 counters on the same cacheline
