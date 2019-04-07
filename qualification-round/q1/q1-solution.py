#!/bin/env/python3

# input() reads a string with a line of input, stripping the ' ' (newline) at the end.
# This is all you need for most Code Jam problems.
t = int(input()) # read a line with a single integer
for i in range(1, t + 1):
    n = input() # read a list of integers, 2 in this case

    answer = 0
    for j in range(0, len(n)):
        if n[j] == "4":
            answer += 10 ** (len(n) - j - 1)

    # check out .format's specification for more formatting options
    print("Case #{}: {} {}".format(i, int(n) - answer, answer))