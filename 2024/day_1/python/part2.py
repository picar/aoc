from collections import defaultdict
from sys import argv


def main():
    with open(argv[1]) as f:
        data = f.readlines()
    
    l1, l2 = zip(*[(int(i[0]), int(i[1])) for i in [i.split() for i in data]])
    l2_count = defaultdict(int)
    for i in list(l2):
        l2_count[i] += 1
    sum = 0
    for i in list(l1):
       sum += i * l2_count[i]

    print(sum)


if __name__ == "__main__":
    main()