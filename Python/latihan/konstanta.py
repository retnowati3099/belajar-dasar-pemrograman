from typing import Final
from pconst import const
import constant

PI: Final = 3.14
print("Phi: %.2f" % (PI))

const.FIRSTNAME = "Retno"
print(f"Hallo {const.FIRSTNAME}!")

print("Gravity:", constant.GRAVITY)


"""
    - Deklarasi konstanta di Python dilakukan menggunakan bantuan tipe class bernama typing.Final
    - Nama konstanta harus dituliskan dalam huruf besar(UPPER_CASE)
"""
