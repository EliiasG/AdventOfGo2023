def extrap(nums):
    return [nums[i+1] - nums[i] for i in range(len(nums)-1)]

def prob1(input):
    sm = 0
    for ln in input:
        nums = list(map(int, ln.split()))
        while any([n != 0 for n in nums]):
            sm += nums[-1]
            nums = extrap(nums)
    return sm

def backextrap(nums):
    if all([n == 0 for n in nums]):
        return 0
    return nums[0]-backextrap(extrap(nums))

def prob2(input):
    sm = 0
    for ln in input:
        nums = list(map(int, ln.split()))
        sm += backextrap(nums)
    return sm

inp = open("input.txt").readlines()
tst = open("test.txt").readlines()
print("Problem 1 result:")
print(prob1(inp))
print("Problem 2 result:")
print(prob2(inp))