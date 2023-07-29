def fib(n):
    # if n == 0:
    #     return 0
    # if n == 1:
    #     return 1
    # return fib(n - 1) + fib(n - 2)
    if n == 0:
        return 0
    if n == 1:
        return 1

    previous = 0
    current = 1

    for _ in range(2, n + 1):
        next = previous + current
        previous = current
        current = next

    return current


# don't touch below this line


def test(n):
    res = fib(n)
    print(f"n: {n}")
    print(f"Fibonacci's {n}th number: {res}")
    print("====================================")


def main():
    print(fib(10))
    print(fib(20))
    print(fib(40))
    print(fib(80))
    print(fib(160))


main()
