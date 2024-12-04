from sys import argv

def main():
    with open(argv[1]) as f:
        data = f.readlines()
    
    l1, l2 = zip(*[(int(i[0]), int(i[1])) for i in [i.split() for i in data]])
    l1 = sorted(l1)
    l2 = sorted(l2)
         
    sum = 0 
    for i in range(len(l1)):
        sum += abs(l1[i] - l2[i])
        
    print(sum)


if __name__ == "__main__":
    main()