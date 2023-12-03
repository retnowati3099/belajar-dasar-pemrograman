panjang = int(input("Masukkan panjang tabel perkalian: "))
lebar = int(input("Masukkan lebar tabel perkalian: "))

print("*\t", end="")
for i in range(1, panjang + 1):
    print(i, "\t", end="")
print()
for i in range(1, lebar + 1):
    print(i, "\t", end="")
    for j in range(1, panjang + 1):
        print(i * j, "\t", end="")
    print()
