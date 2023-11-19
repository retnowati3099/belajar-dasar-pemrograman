#include <iostream>

using namespace std;

int main()
{
    // deklarasi variabel
    int bilangan;

    cout << "Masukkan bilangan: ";
    cin >> bilangan;

    // proses
    if (bilangan % 2 == 0)
    {
        cout << bilangan << " merupakan bilangan genap.\n";
    }
    else
    {
        cout << bilangan << " merupakan bilangan ganjil.\n";
    }

    // memberhentikan program
    return 0;
}