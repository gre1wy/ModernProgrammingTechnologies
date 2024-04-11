import sys

def bubble_sort(nums: list):
    n = len(nums)
    for i in range(n):
        for j in range(0, n-i-1):
            if nums[j] > nums[j+1]:
                nums[j], nums[j+1] = nums[j+1], nums[j]
    return nums

def data_manip(data: str):
    numbers = data.split()
    converted_numbers = []

    for num in numbers:
        if '.' not in num:
            converted_numbers.append(int(num))
        else:
            converted_numbers.append(float(num))

    sorted_numbers = bubble_sort(converted_numbers)
    return sorted_numbers


def main(data_from_test=None):
    try:
        if data_from_test is None:
            input = sys.stdin.readline() # return string
            data_list = data_manip(input) # string -> list of numbers
        else: 
            data_list = data_from_test
        sorted_numbers = bubble_sort(data_list)
        sys.stdout.write(' '.join(map(str, sorted_numbers))) # shows sorted list
        return 0, sorted_numbers
    except Exception:
        sys.stderr.write(f"Values are invalid\n")
        return 1

if __name__ == "__main__":
    sys.exit(main())