import os

def main():
    file_test1 = "./input1.txt"
    file_test2 = "./input2.txt"
    cnt = 0
    for cnt in range(100):
        with open(file_test2, "a+") as f:
            f.write("statement " + str(cnt))
        if cnt % 3 == 0 or cnt % 5 == 0 :
            with open(file_test2, "a+") as f:
                f.write(str('\f'))
    cnt = 0
    for cnt in range(100):
        with open(file_test1, "a+") as f:
            f.write("Line " + str(cnt) + "\n")


if __name__ == '__main__':
    main()
