def exponential_growth(n, factor, days):
    result = []

    for i in range(days + 1):
        current = n * (factor**i)
        result.append(current)

    return result


# don't touch below this line


def test(n, factor, days):
    growth_sequence = exponential_growth(n, factor, days)
    print(f"- Initial followers: {n}")
    print(f"- Growth factor: {factor}")
    print(f"- Days: {days}")
    print(f"Growth sequence: {growth_sequence}")
    print("=====================================")


def main():
    test(10, 2, 4)
    test(20, 1.5, 6)
    test(30, 3, 3)
    test(40, 10, 10)


main()
