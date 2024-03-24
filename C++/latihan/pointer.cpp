#include <iostream>

using namespace std;

int main()
{
    // penerapan pointer
    int num1 = 5;
    int *num2 = &num1;

    cout << num1 << endl;
    cout << &num1 << endl;
    cout << num2 << endl;
    cout << *num2 << endl;

    // efek perubahan nilai pointer
    num1 = 10;
    cout << num1 << endl;
    cout << &num1 << endl;
    cout << num2 << endl;
    cout << *num2 << endl;

    *num2 = 15;
    cout << num1 << endl;
    cout << &num1 << endl;
    cout << num2 << endl;
    cout << *num2 << endl;

    num1 = 60;
    num2 = &num1;
    cout << num1 << endl;
    cout << &num1 << endl;
    cout << num2 << endl;
    cout << *num2 << endl;

    //parameter pointer
    
}