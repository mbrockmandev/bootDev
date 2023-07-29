class Graph:
    def depth_first_search(self, start_vertex):
        visited = set()

        self.depth_first_search_r(visited, start_vertex)

        return sorted(visited)

    def depth_first_search_r(self, visited, current_vertex):
        visited.add(current_vertex)

        sorted_neighbors = sorted(self.graph[current_vertex])

        for neighbor in sorted_neighbors:
            if neighbor not in visited:
                self.depth_first_search_r(visited, neighbor)



        # don't touch below this line

    def __init__(self):
        self.graph = {}

    def add_edge(self, u, v):
        if u in self.graph.keys():
            self.graph[u].add(v)
        else:
            self.graph[u] = set([v])
        if v in self.graph.keys():
            self.graph[v].add(u)
        else:
            self.graph[v] = set([u])

    def __repr__(self):
        result = ""
        for key in self.graph.keys():
            result += f"Vertex: '{key}'\n"
            for v in sorted(self.graph[key]):
                result += f"has an edge leading to --> {v} \n"
        return result


def test(edges_to_add, starting_at):
    graph = Graph()
    for edge in edges_to_add:
        graph.add_edge(edge[0], edge[1])
        print(f"Added edge: {edge}")
    print("-------------------------------------")
    for v in graph.depth_first_search(starting_at):
        print(f"Visiting: {v} \n")
    print("=====================================")


def main():
    test(
        [
            ("New York", "London"),
            ("New York", "Cairo"),
            ("New York", "Tokyo"),
            ("Cairo", "Dubai"),
            ("Dubai", "Kyiv"),
        ],
        "New York",
    )
    test(
        [
            ("New York", "London"),
            ("New York", "Cairo"),
            ("New York", "Tokyo"),
            ("London", "Dubai"),
            ("Cairo", "Kyiv"),
            ("Cairo", "Madrid"),
            ("London", "Madrid"),
            ("Buenos Aires", "New York"),
            ("Tokyo", "Buenos Aires"),
            ("Kyiv", "San Francisco"),
        ],
        "New York",
    )


main()
