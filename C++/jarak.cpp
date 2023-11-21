#include <iostream>

using namespace std;

int main()
{
    float jarak, kecepatan, waktu;

    cout << "Program Menghitung Jarak Berdasarkan Kecepatan dan Waktu\n";
    cout << "Masukkan kecepatan: ";
    cin >> kecepatan;
    cout << "Masukkan waktu: ";
    cin >> waktu;

    jarak = kecepatan * waktu;

    cout << "Jika kecepatan " << kecepatan << " km/jam dan waktu tempuh " << waktu << " jam, maka jarak yang ditempuh adalah " << jarak << " km.\n";
}