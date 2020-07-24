import sys

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python3 create_dataset.py ${dataset_num}")
        exit(1)
    
    if not sys.argv[1].isdecimal():
        print("input number for dataset_num")
        exit(1)
    