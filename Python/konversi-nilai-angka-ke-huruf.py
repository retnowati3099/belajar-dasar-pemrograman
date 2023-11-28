nilai = int(input("Masukkan nilai kamu: "))

if nilai <= 100 and nilai >= 85:
    print("Nilai kamu " + str(nilai) + ", kamu mendapat A.")
elif nilai >= 70:
    print("Nilai kamu " + str(nilai) + ", kamu mendapat B.")
elif nilai >= 40:
    print("Nilai kamu " + str(nilai) + ", kamu mendapat C.")
elif nilai >= 20:
    print("Nilai kamu " + str(nilai) + ", kamu mendapat D.")
else:
    print("Nilai kamu " + str(nilai) + ", kamu mendapat E.")
