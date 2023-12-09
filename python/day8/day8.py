def parseInput(input):
    return {sp[0]: sp[1].strip()[1:-1].split(", ") for sp in [ln.split(" = ") for ln in input[2:]]}

def prob1(input):
    map = parseInput(input)
    inst = input[0].strip()
    i = 0
    node = "AAA"
    while node != "ZZZ":
        c = inst[i%len(inst)]
        if c == "L":
            node = map[node][0]
        if c == "R":
            node = map[node][1]
        i += 1
    return i

def prob2(input):
    map = parseInput(input)
    inst = input[0].strip()
    i = 0
    nodes = [key for key in map.keys() if key[-1] == "A"]
    while not all([node[-1] == "Z" for node in nodes]):
        c = inst[i%len(inst)]
        if c == "L":
            nodes = [map[node][0] for node in nodes]
        if c == "R":
            nodes = [map[node][1] for node in nodes]
        i += 1
    return i

inp = open("input.txt").readlines()
tst = open("test.txt").readlines()
print("Problem 1 result:")
print(prob1(inp))
print("Problem 2 result:")
print(prob2(inp))