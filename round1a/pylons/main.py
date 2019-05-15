#!/bin/env/python3

# input() reads a string with a line of input, stripping the ' ' (newline) at the end.
# This is all you need for most Code Jam problems.

def find_valid_move(past, pos_x, pos_y, row, col):
    return [(x,y)
            for x in range(1, row)
            for y in range(1, col)
            if x != pos_x
                and y != pos_y
                and x-y != pos_x-pos_y
                and x+y != pos_x+pos_y
                and (x,y) not in past]


def find_sequence(current, row, col):

    # Found the answer - return
    if len(current) == row*col:
        return current

    # Only for first depth
    if len(current) == 0:

        for i in range(1, row+1):
            for j in range(1, col+1):

                # Find the next move posible
                next_move = find_valid_move(current, i, j, row+1, col+1)

                # Unable to find any move return back 1 step
                if len(next_move) == 0:
                    return []

                for item in next_move:
                    current.append(item)
                    #print("going into: ", current)
                    sol = find_sequence(current, row, col)
                    current.remove(item)

                    # We got the solution
                    if len(sol) != 0:
                        return sol
    else:
        next_move = find_valid_move(current, current[-1][0], current[-1][1], row+1, col+1)
        if len(next_move) == 0:
            return []

        for item in next_move:
            current.append(item)
            # print("going into: ", current)
            sol = find_sequence(current, row, col)
            current.remove(item)

            if len(sol) != 0:
                if len(sol) == row*col:
                    return sol
                else:
                    return sol + [item]

    return []


t = int(input()) # read a line with a single integer
for i in range(1, t + 1):

    the_string = input()
    row, col = the_string.split()

    answer = find_sequence([], int(row), int(col))
    if len(answer) == 0:
        print("Case #%d: IMPOSSIBLE" % i)
    else:
        print("Case #%d: POSSIBLE" % i)
        for steps in answer:
            print(steps[0], steps[1])
