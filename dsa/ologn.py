import time


def binary_search(target, arr):
    lo = 0
    hi = len(arr) - 1
    mid = (lo + hi) // 2

    while lo <= hi:
        if mid < target:
            lo = mid + 1
        else:
            hi = mid - 1
        mid = (lo + hi) // 2
        if arr[mid] == target:
            return True
    return False


# don't touch below this line


def benchmark(names_dict, first_name):
    start = time.time()
    test(names_dict, first_name)
    end = time.time()

    timeout = 0.05

    if (end - start) < timeout:
        print(f"test completed in less than {timeout * 1000} milliseconds!")
    else:
        print(f"test took too long ({(end - start) * 1000} milliseconds). Speed it up!")
    print("====================================")


def test(target, arr):
    res = binary_search(target, arr)
    print(f"- len arr: {len(arr)}")
    print(f"- target: {target}")
    print(f"Result: {res}")
    print("------------------------------------")


def main():
    complexity = 2000000
    nums = get_nums(complexity)
    benchmark(int(complexity * 0.2344), nums)
    benchmark(int(complexity * 2), nums)
    benchmark(int(complexity + 1), nums)
    benchmark(int(complexity * 0.765), nums)


def get_nums(num):
    nums = []
    for i in range(num):
        nums.append(i)
    return nums


main()
