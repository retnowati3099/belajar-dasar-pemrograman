baris = int(input("Maskkan banyak baris: "))

for i in range(1, baris + 1):
    for j in range(1, i + 1):
        print("* ", end="")
    print()

print("============")

bintang = ""
for i in range(1, baris + 1):
    bintang += "* "
    print(bintang)

print("============")

for i in range(1, baris + 1):
    # spasi
    for j in range(baris - 1, i - 1, -1):
        print(" ", end="")

    # bintang
    for k in range(1, i + 1):
        print("* ", end="")

    print()
