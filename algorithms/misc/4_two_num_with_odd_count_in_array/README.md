You are given an integer array, where all numbers except for TWO numbers appear even number of times. Find out the two numbers which appear odd number of times.

*Solution:*
1. XOR all the n numbers. 
2. Result will be knocked out for all the even pairs as a^a=0 The result now contains only XOR of the two odd out numbers. 
3. Find the first bit position in the result that is 1. Definitely this bit position both the odd numbers have different bit values. i.e. one has a 0 and another has a 1 at this position. Let this position be x 
4. XOR the elements that have 1 at x bit position and XOR the elements that have 0 at x bit position. The two XOR results would give the two odd count numbers.