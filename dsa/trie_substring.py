import json


class Trie:
    def find_matches(self, document):
        matches = set()

        for i in range(len(document)):
            level = self.root  # dictionary to hold current level
            for j in range(i, len(document)):
                if document[j] not in level:
                    level = self.root
                    break
                else:
                    level = level[document[j]]
                    if self.end_symbol in level:
                        matches.add(document[i:j+1])
        return matches

    # don't touch below this line

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


def test(words, document):
    trie = Trie()
    for word in words:
        trie.add(word)
    print("Trie:")
    print(json.dumps(trie.root, sort_keys=True, indent=2))
    print("-------------------------------------")
    matches = trie.find_matches(document)
    print(f"Document: '{document}'")
    print("Matches:")
    for match in sorted(matches):
        print(f" - {match}")
    print("=====================================")


def main():
    test(["bad", "baddie", "badguy", "suck"], "the badguy really sucks")
    test(["be", "bad", "back", "bat"], "he is back at bat")


main()
