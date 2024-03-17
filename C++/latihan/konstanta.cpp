#include <iostream>
#define PI 3.14 // menggunakan #define

using namespace std;

int main()
{
    // menggunakan keyword const
    const string FIRSTNAME = "Retno";

    // menggunakan keyword constexpr
    constexpr float GRAVITY = 9.8;

    cout << "Phi: " << PI << endl;
    cout << "Hello "<<FIRSTNAME << endl;
    cout << "Gravity: " << GRAVITY << endl;
}

/*
    - Ketika melakukan deklarasi konstanta harus diberi nilai
*/