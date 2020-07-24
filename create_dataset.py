import sys
import random

NMAX = 50
DAYMAX = 30
# この値が小さいほど-1が返る頻度が多くなる。一旦は適当そうな20%に設定
POSSIBILITY = 20

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python3 create_dataset.py ${dataset_num}")
        exit(1)
    
    if not sys.argv[1].isdecimal():
        print("input number for dataset_num")
        exit(1)
    
    dataset_num = int(sys.argv[1])
    print(dataset_num)
    for i in range(dataset_num):
        n = random.randint(1, NMAX)
        print(n)
        for j in range(n):
            max_day = random.randint(1, DAYMAX)
            # 1人に対して1日は必ず集まれる日を作る
            must_day = random.randint(1, max_day)
            days = list()
            for k in range(1, max_day + 1):
                if k == must_day:
                    days.append(k)
                    continue

                if random.randint(1, 100) <= POSSIBILITY:
                    days.append(k)

            for i, day in enumerate(days):
                if i == len(days) - 1:
                    print(day)
                else:
                    print(day, end=" ")