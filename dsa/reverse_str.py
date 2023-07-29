def reverse_string(s):
    if len(s) == 0:
        return s
    return reverse_string(s[1:]) + s[0]


# don't touch below this line


def test(s):
    result = reverse_string(s)
    print(f"Original string: {s}")
    print(f"Reversed string: {result}")
    print("====================================")


def main():
    test("Socialytics")
    test("Hello, World!")
    test("RecursionIsFun")


main()
