bil1 = int(input("Inputkan bilangan 1: "))
bil2 = int(input("Inputkan bilangan 2: "))
bil3 = int(input("Inputkan bilangan 3: "))

if bil1 >= bil2 and bil1 >= bil3:
    print("Bilangan terbesar adalah " + str(bil1))
elif bil2 >= bil1 and bil2 >= bil3:
    print("Bilangan terbesar adalah " + str(bil2))
else:
    print("Bilangan terbesar adalah " + str(bil3))
