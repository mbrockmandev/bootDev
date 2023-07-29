def find_min(nums):
    min = float("inf")
    for num in nums:
        if num < min:
            min = num

    return min


# don't touch below this line


def test(nums):
    res = find_min(nums)
    print(f"Follower counts: {nums}")
    print(f"Lowest follower count: {res}")
    print("----")


def main():
    test([7, 4, 3, 100, 2343243, 343434, 1, 2, 32])
    test([12, 12, 12])
    test([10, 200, 3000, 5000, 4])


main()
