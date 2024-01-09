arrLen = int(input("Masukkan panjang array = "))


def arrSum(arrLen):
    array = [None] * arrLen

    total = 0
    for i in range(0, arrLen):
        array[i] = int(input("Elemen ke-" + str(i + 1) + "= "))
        total += array[i]

    return total


print(arrSum(arrLen))
