def power_set(input_set):
    # input_length = len(input_set)
    # output = []

    # for i in range(2**input_length):
    #     sub = []

    #     for j in range(input_length):
    #         if (i >> j) & 1:
    #             sub.append(input_set[j])
    #     output.append(sub)
    # return output

    def helper(start, sub):
        result.append(sub[:])

        for i in range(start, len(input_set)):
            helper(i + 1, sub + [input_set[i]])

    result = []
    helper(0, [])
    return result


# don't touch below this line


def test(input_set):
    result = power_set(input_set)
    print(f"Number of subsets of {input_set}: {len(result)}")
    print("====================================")


def main():
    for i in range(1, 22):
        nums = list(range(1, i))
        test(nums)


main()
