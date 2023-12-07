def parseInput(lines):
    return [[int(el) for el in line.split(" ")[1:] if el != ""] for line in lines]


def prob1(input):
    times, results = parseInput(input)
    r = 1
    for i, time in enumerate(times):
        s = 0
        for vel in range(1, time):
            dst = (time - vel)*vel
            if dst > results[i]:
                s += 1
        r *= s
    return r

def prob2(input):
    input = ["A " + "".join(line.split(" ")[1:]).strip() for line in input]
    return prob1(input)

inp = open("input.txt").readlines()
tst = open("test.txt").readlines()
print("Problem 1 result:")
print(prob1(inp))
print("Problem 2 result:")
print(prob2(inp))