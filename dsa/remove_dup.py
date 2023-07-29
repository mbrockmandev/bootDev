def remove_duplicates(nums):
    return sorted(list(set(nums)))


# don't touch below this line


def test(nums):
    result = remove_duplicates(nums)
    print(f"Original list: {nums}")
    print(f"List after removing duplicates: {result}")
    print("====================================")


def main():
    test([1, 2, 3, 4, 4, 5, 6, 7, 7, 7])
    test([10, 10, 20, 30, 30, 30, 40, 50, 50])
    test([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1])


main()
