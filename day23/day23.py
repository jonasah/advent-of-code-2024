import networkx as nx
import os


def part1(input_file):
    G = read_graph(input_file)
    sets = [
        nodes
        for nodes in nx.algorithms.clique.enumerate_all_cliques(G)
        if len(nodes) == 3 and any(node.startswith("t") for node in nodes)
    ]
    return len(sets)


def part2(input_file):
    G = read_graph(input_file)
    nodes = max(nx.algorithms.clique.find_cliques(G), key=len)
    nodes.sort()
    return ",".join(nodes)


def read_graph(input_file):
    file_path = os.path.join(os.path.dirname(__file__), input_file)
    return nx.read_adjlist(file_path, delimiter="-")


def verify(expected, actual):
    if expected != actual:
        print(f"Wrong answer! Expected: {expected}. Actual: {actual}")
        exit(1)


verify(7, part1("testinput.txt"))
verify(1467, part1("input.txt"))

verify("co,de,ka,ta", part2("testinput.txt"))
verify("di,gs,jw,kz,md,nc,qp,rp,sa,ss,uk,xk,yn", part2("input.txt"))
