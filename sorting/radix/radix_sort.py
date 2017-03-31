import math


def radix_sort(data, radix=10):
    max_digit = int(math.ceil(math.log(max(data), radix)))

    for digit in range(1, max_digit + 1):
        bucket = [[] for i in range(radix)]
        for val in data:
            bucket[val % radix**digit / radix**(digit-1)].append(val)
        del data[:]
        for each in bucket:
            data.extend(each)


data = [123, 21, 23, 4, 54, 3214, 546, 466, 34, 1]
radix_sort(data, radix=10)
print data
