def square_digits(num):
    num_str = str(num)

    num_arr = [d for d in num_str]
    result_arr = []

    for n in num_arr:
        result_arr.append(str(int(n) * int(n)))

    result_str = "".join(result_arr)
    return int(result_str)
