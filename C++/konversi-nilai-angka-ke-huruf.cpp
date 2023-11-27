#include <iostream>

using namespace std;

int main()
{
    int nilai;

    cout << "Masukkan nilai kamu: ";
    cin >> nilai;

    if (nilai <= 100 && nilai >= 85)
    {
        cout << "Nilai kamu " << nilai << ", kamu mendapat A.\n";
    }
    else if (nilai >= 70)
    {
        cout << "Nilai kamu " << nilai << ", kamu mendapat B.\n";
    }
    else if (nilai > 40)
    {
        cout << "Nilai kamu " << nilai << ", kamu mendapat C.\n";
    }
    else if (nilai >= 20)
    {
        cout << "Nilai kamu " << nilai << ", kamu mendapat D.\n";
    }
    else
    {
        cout << "Nilai kamu " << nilai << ", kamu mendapat D.\n";
    }
}
