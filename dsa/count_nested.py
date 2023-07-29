def count_nested_levels(nested_comments, target_comment_id, level=1):
    if target_comment_id in nested_comments:
        return level
    for id, comment in nested_comments.items():
        nest_level = count_nested_levels(comment, target_comment_id, level + 1)
        if nest_level != -1:
            return nest_level
    return -1


# don't touch below this line


def test(nested_comments, target_comment_id):
    result = count_nested_levels(nested_comments, target_comment_id)
    print(f"Nested levels of comment {target_comment_id}: {result}")
    print("====================================")


def main():
    test_comments = {
        1: {
            2: {
                3: {},
                4: {5: {}},
            },
            6: {},
            7: {8: {9: {10: {}}}},
        }
    }
    test(test_comments, 1)
    test(test_comments, 2)
    test(test_comments, 3)
    test(test_comments, 5)
    test(test_comments, 10)


main()
