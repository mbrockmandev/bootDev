def matchmake(queue, player):
    if player[1] == "leave":
        queue.search_and_remove(player[0])

    if player[1] == "join":
        queue.push(player[0])
        if queue.size() >= 4:
            p1 = queue.pop()
            p2 = queue.pop()
            if p1 is not None and p2 is not None:
                print(f"{p1} matched {p2}")


# don't touch below this line


class Queue:
    def __init__(self):
        self.items = []

    def push(self, item):
        self.items.insert(0, item)

    def pop(self):
        if len(self.items) == 0:
            return None
        temp = self.items[-1]
        del self.items[-1]
        return temp

    def peek(self):
        if len(self.items) == 0:
            return None
        return self.items[-1]

    def size(self):
        return len(self.items)

    def search_and_remove(self, item):
        if item not in self.items:
            return None
        self.items.remove(item)
        return item


def test(players):
    queue = Queue()
    for player in players:
        name = player[0]
        action = player[1]
        if action == "leave":
            print(f"{name} left the queue.")
        if action == "join":
            print(f"{name} joined the queue.")
        matchmake(queue, player)
        print(f"Players in queue: {queue.items}")
        print("-------------------------------------")
    print("=====================================")


def main():
    test(
        [
            ("Alice", "join"),
            ("Bob", "join"),
            ("Charlie", "join"),
            ("David", "join"),
            ("Eve", "join"),
            ("Frank", "join"),
            ("Frank", "leave"),
            ("Grace", "join"),
        ]
    )
    test(
        [
            ("Frank", "join"),
            ("Alice", "join"),
            ("Bob", "join"),
            ("Charlie", "leave"),
            ("David", "join"),
            ("Eve", "join"),
            ("Errol", "join"),
            ("Jake", "join"),
            ("Errol", "leave"),
            ("Errol", "leave"),
            ("Grace", "join"),
        ]
    )


main()
