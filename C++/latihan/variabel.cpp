#include <iostream>

using namespace std;

int main()
{
    // deklarasi variabel beserta tipe datanya
    string firstName = "Retno";
    string lastName;
    lastName = "Wati";

    // output sama, penulisan beda
    cout << "Hallo " << firstName << " " << lastName << "!" << endl;
    cout << "Hallo " << firstName << " " << lastName << "!\n";

    // deklarasi multi variabel
    string first, second, third;
    first = "satu";
    second = "dua";
    third = "tiga";
    string fourth = "empat", fifth = "lima", sixth = "enam";

    cout << first << " " << second << " " << third << " " << fourth << " " << fifth << " " << sixth << endl;

    return 0;
}