#include <iostream>

using namespace std;

int main()
{
    int panjang, lebar;

    cout << "Masukkan panjang tabel perkalian: ";
    cin >> panjang;
    cout << "Masukkan lebar tabel perkalian: ";
    cin >> lebar;

    cout << "*\t";
    for (int i = 1; i <= panjang; i++)
    {
        cout << i << "\t";
    }
    cout << endl;
    for (int i = 1; i <= lebar; i++)
    {
        cout << i << "\t";
        for (int j = 1; j <= panjang; j++)
        {
            cout << (i * j) << "\t";
        }
        cout << endl;
    }

    return 0;
}