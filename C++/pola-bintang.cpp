#include <iostream>

using namespace std;

int main()
{
    int baris;

    cout << "Masukkan banyak baris: ";
    cin >> baris;

    for (int i = 1; i <= baris; i++)
    {
        for (int j = 1; j <= i; j++)
        {
            cout << "* ";
        }
        cout << "\n";
    }

    cout << "===========\n";

    string bintang = "";
    for (int i = 1; i <= baris; i++)
    {
        bintang += "* ";
        cout << bintang << endl;
    }

    cout << "===========\n";

    for (int i = 1; i <= baris; i++)
    {
        // spasi
        for (int j = (baris - 1); j >= i; j--)
        {
            cout << " ";
        }

        // bintang
        for (int j = 1; j <= i; j++)
        {
            cout << "* ";
        }
        cout << "\n";
    }

    return 0;
}