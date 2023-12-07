def parseInput(input):
    return [(ln[0], int(ln[1])) for ln in [line.split(" ") for line in input]]

def value(hand):
    val = 0
    for i, c in enumerate(hand):
        val += "23456789TJQKA".index(c) * 13**(4-i)
    amt = list({x :hand.count(x) for x in hand}.values())
    five = amt.count(5) == 1
    four = amt.count(4) == 1
    three = amt.count(3) == 1
    onepair = amt.count(2) == 1
    twopair = amt.count(2) == 2
    full = three and onepair
    if five:
        return val + 13**5 * 10
    if four:
        return val + 13**5 * 9
    if full:
        return val + 13**5 * 8
    if three:
        return val + 13**5 * 7
    if twopair:
        return val + 13**5 * 6
    if onepair:
        return val + 13**5 * 5
    return val

def prob1(input):
    s = sorted(parseInput(input), key=lambda x: value(x[0]))
    print(s)
    return sum([(i+1) * s[i][1] for i in range(len(s))])

def prob2(input):
    pass

inp = open("input.txt").readlines()
tst = open("test.txt").readlines()
print("Problem 1 result:")
print(prob1(inp))
print("Problem 2 result:")
print(prob2(inp))