import time


def find_last_name(names_dict, first_name):
    return names_dict[first_name]


# don't touch below this line


def benchmark(names_dict, first_name):
    start = time.time()
    test(names_dict, first_name)
    end = time.time()

    timeout = 0.05

    if (end - start) < timeout:
        print(f"find_last_name completed in less than {timeout * 1000} milliseconds!")
    else:
        print(
            f"find_last_name took too long ({(end - start) * 1000} milliseconds). Speed it up!"
        )
    print("====================================")


def test(names_dict, first_name):
    res = find_last_name(names_dict, first_name)
    print(f"- first_name: {first_name}")
    print(f"Last name: {res}")
    print("------------------------------------")


def main():
    complexity = 2000000
    names_dict = get_name_dict(complexity)
    benchmark(names_dict, "bree1999999")


def get_name_dict(num):
    names = {}
    for i in range(num):
        names[f"bree{i}"] = f"fuca{i}"
    return names


main()
