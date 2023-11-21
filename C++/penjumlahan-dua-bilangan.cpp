#include <iostream>

using namespace std;

int main()
{
    int num1, num2, jumlah;

    cout << "Masukkan bilangan pertama: ";
    cin >> num1;
    cout << "Masukkan bilangan kedua: ";
    cin >> num2;

    jumlah = num1 + num2;

    cout << num1 << " + " << num2 << " = " << jumlah << endl;

    return 0;
}