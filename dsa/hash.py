class HashMap:
    def insert(self, key, value):
        self.resize()

        index = self.key_to_index(key)
        self.hashmap[index] = (key, value)

    def resize(self):
        if len(self.hashmap) == 0:
            self.hashmap = [None]
            return

        load = self.current_load()
        if load < 0.05:
            return

        size = len(self.hashmap) * 10  # bigger!
        old_hm = self.hashmap
        self.hashmap = [None for i in range(size)]

        for pair in old_hm:
            if pair is not None:
                k, v = pair
                self.insert(k, v)

    def current_load(self):
        if len(self.hashmap) == 0:
            return 1

        filled_entries = sum(1 for p in self.hashmap if p is not None)
        return filled_entries / len(self.hashmap)

    # don't touch below this line

    def __init__(self, size):
        self.hashmap = [None for i in range(size)]

    def key_to_index(self, key):
        sum = 0
        for c in key:
            sum += ord(c)
        return sum % len(self.hashmap)

    def __repr__(self):
        final = ""
        for i, v in enumerate(self.hashmap):
            if v != None:
                final += f" - {str(v)}\n"
        return final


def test(items):
    hm = HashMap(0)
    for item in items:
        key = item[0]
        val = item[1]
        print(f"insert({key}, {val})")
        hm.insert(key, val)
        print(f"Load: {hm.current_load()}")
        print(f"Size: {len(hm.hashmap)}")
        print("-------------------------------------")
    print("=====================================")


def main():
    test(
        [
            ("apple", 1),
            ("banana", 2),
            ("cherry", 3),
            ("mango", 4),
            ("orange", 5),
            ("pear", 6),
            ("plum", 7),
            ("strawberry", 8),
            ("watermelon", 9),
        ]
    )


main()
