print("Program Menghitung Jarak Berdasarkan Kecepatan dan Waktu Tempuh")
kecepatan = float(input("Masukkan kecepatan: "))
waktu = float(input("Masukkan waktu: "))

jarak = kecepatan * waktu
print(
    "Jika kecepatan "
    + str(kecepatan)
    + " km/jam dan waktu tempuh "
    + str(waktu)
    + " jam, maka jarak yang ditempuh adalah "
    + str(jarak)
    + " km."
)
# output = "Jika kecepatan {} km/jam dan waktu tempuh {} jam, maka jarak yang ditempuh adalah {} km."
# print(output.format(kecepatan, waktu, jarak))
