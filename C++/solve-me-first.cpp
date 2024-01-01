#include <iostream>

using namespace std;

// function prototype (function declaration)
int solveMeFirst(int a, int b);

int main()
{
    int a, b;
    cout << "a = ";
    cin >> a;
    cout << "b = ";
    cin >> b;

    cout << solveMeFirst(a, b)<<endl; // memanggil fungsi solveMeFirst()
}

// function definition
int solveMeFirst(int a, int b)
{
    return a + b;
}