#include <iostream>

using namespace std;

int main()
{
    int panjang;
    int sum = 0;

    cout << "Masukkan jumlah angka yang ingin dijumlahkan: ";
    cin >> panjang;
    
    int temp[panjang];

    for (int i = 0; i < panjang; i++)
    {
        cout << "Angka ke-" << i + 1 << ": ";
        cin >> temp[i];
        sum += temp[i];
    }
    cout << sum << endl;
}