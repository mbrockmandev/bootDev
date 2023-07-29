import json


class Trie:
    def words_with_prefix(self, prefix):
        # should return all words in Trie that start with given prefix

        words = []

        curr = self.root
        for c in prefix:
            if c not in curr:
                return []
            curr = curr[c]

        return self.search_level(curr, prefix, words)

    def search_level(self, curr, curr_prefix, words):
        if self.end_symbol in curr:
            words.append(curr_prefix)

        for k in sorted(curr):
            if k != self.end_symbol:
                words = self.search_level(curr[k], curr_prefix + k, words)
        return words

    # don't touch below this line

    def exists(self, word):
        current = self.root
        for letter in word:
            if letter not in current:
                return False
            current = current[letter]
        return self.end_symbol in current

    def add(self, word):
        current = self.root
        for letter in word:
            if letter not in current:
                current[letter] = {}
            current = current[letter]
        current[self.end_symbol] = True

    def __init__(self):
        self.root = {}
        self.end_symbol = "*"


def test(words, prefix):
    trie = Trie()
    for word in words:
        trie.add(word)
    print("Trie:")
    print(json.dumps(trie.root, sort_keys=True, indent=2))
    print("-------------------------------------")
    words = trie.words_with_prefix(prefix)
    print(f"Words with prefix: '{prefix}':")
    for word in sorted(words):
        print(f" - {word}")
    print("=====================================")


def main():
    test(["be", "bad", "back", "bat"], "ba")
    test(["a", "to", "tea", "ted", "ten", "i", "in", "inn"], "te")


main()
