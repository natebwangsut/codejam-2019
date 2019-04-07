#!/bin/env/python3

def mirror(x):
    if x == "E":
        return "S"
    else:
        return "E"

# input() reads a string with a line of input, stripping the ' ' (newline) at the end.
# This is all you need for most Code Jam problems.
t = int(input()) # read a line with a single integer
for i in range(1, t + 1):

    n = input() # read a list of integers, 2 in this case
    path = input() # read the path

    answer = ""
    buffer = 0
    firstchar = ""

    #
    for j in range(0, len(path)):
        answer += mirror(path[j])

    # check out .format's specification for more formatting options
    print("Case #{}: {}".format(i, answer))