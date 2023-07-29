def find_next_square(sq: int) -> int:
    sqrt_sq = sq**0.5

    if sqrt_sq % 1 != 0:
        return -1
    return (sqrt_sq + 1) ** 2


result = find_next_square(121)  # 144
print(result)
result = find_next_square(625)  # 676
print(result)
result = find_next_square(319225)  # 320356
print(result)
result = find_next_square(15241383936)  # 15241630849
print(result)
