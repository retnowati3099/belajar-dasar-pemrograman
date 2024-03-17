# deklarasi variabel beserta tipe datanya
first_name: str = "Retno"

# deklarasi variabel yang umum digunakan
last_name = "Wati"

# output sama, penulisan beda
print("Hello %s %s!" % (first_name, last_name))
print("Hello Retno Wati!")
print(f"Hello {first_name} {last_name}!")
print(f"Hello {first_name} {last_name}!")
print("Hello {} {}!".format(first_name, last_name))
print("Hello {0} {1}!".format(first_name, last_name))
print("Hello {f_name} {l_name}!".format(f_name=first_name, l_name=last_name))
print("Hello " + first_name + " " + last_name + "!")
print("Hello", first_name, last_name, "!")

# deklarasi multi variabel
first, second, third = "satu", "dua", "tiga"
print(first, second, third)