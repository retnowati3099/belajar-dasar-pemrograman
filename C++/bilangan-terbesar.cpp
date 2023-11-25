#include <iostream>

using namespace std;

int main()
{
    int bil1, bil2, bil3;

    cout << "Inputkan bilangan 1: ";
    cin >> bil1;
    cout << "Inputkan bilangan 2: ";
    cin >> bil2;
    cout << "Inputkan bilangan 3: ";
    cin >> bil3;

    if (bil1 >= bil2 && bil1 >= bil3)
    {
        cout << "Bilangan terbesar adalah " << bil1 << endl;
    }
    else if (bil2 >= bil1 && bil2 >= bil3)
    {
        cout << "Bilangan terbesar adalah " << bil2 << endl;
    }
    else
    {
        cout << "Bilangan terbesar adalah " << bil3 << endl;
    }
}