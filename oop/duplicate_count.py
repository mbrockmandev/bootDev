def duplicate_count(text):
    seen = set()
    duplicates = set()

    for c in text.lower():
        if c.isalnum():
            if c in seen:
                duplicates.add(c)
            else:
                seen.add(c)

    return len(duplicates)

