sum = 0
panjang = int(input("Masukkan jumlah angka yang ingin dijumlahkan: "))

temp = [None] * panjang
print(len(temp))
for i in range(0, panjang):
    temp[i] = int(input("Angka ke-" + str(i + 1) + ": "))
    sum += temp[i]
print(sum)
