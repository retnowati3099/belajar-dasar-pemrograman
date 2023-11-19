#include <iostream>
// #define PHI 3.14

using namespace std;

int main()
{
    const float PHI = 3.14; // deklarasi konstanta PHI
    float radius, luas;     // deklarasi variabel radius

    // nama program
    cout << "MENGHITUNG LUAS LINGKARAN" << endl;
    cout << "===========================" << endl;

    cout << "Masukkan radius: ";
    cin >> radius;

    // proses
    luas = PHI * radius * radius;

    // menampilkan hasil
    cout << "Lingkaran yang memiliki radius sebesar " << radius << " memiliki luas " << luas << endl;

    // memberhentikan program
    return 0;
}