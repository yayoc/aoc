#include <map>

#include "../utils.h"

enum Element { Rock = 1, Paper = 2, Scissors = 3 };
enum Result { Lost = 0, Draw = 3, Win = 6 };

map<string, Element> elements = {
    {"A", Rock},
    {"X", Rock},
    {"B", Paper},
    {"Y", Paper},
    {"C", Scissors},
    {"Z", Scissors},
};

map<string, Result> results = {
    {"X", Lost},
    {"Y", Draw},
    {"Z", Win},
};

Element counter(Element e) {
    switch (e) {
    case Rock: return Paper;
    case Paper: return Scissors;
    case Scissors: return Rock;
    default: error("unexpected element");
    }
}

Result result(Element you, Element opponent) {
    if (you == opponent) return Draw;
    return (you == counter(opponent)) ? Win : Lost;
}

void part1() {
    ifstream ifs = open_file("input.txt");

    int sum = 0;
    string line;
    while (getline(ifs, line))
    {
        istringstream iss(line);
        string o, y;
        if (!(iss >> o >> y)) { break; } 
        Result r = result(elements.at(y), elements.at(o));
        sum += (r + elements.at(y));
    }
    cout << "sum=" << sum << endl;
}

void part2() {
   ifstream ifs = open_file("input.txt");

    int sum = 0;
    string line;
    while (getline(ifs, line))
    {
        istringstream iss(line);
        string a, b;
        if (!(iss >> a >> b)) { break; } 
        Result r = results.at(b);
        Element opponent = elements.at(a);
        Element you;
        if (r == Draw) {
            you = opponent;
        } else if (r == Win) {
            you = counter(opponent);
        } else if (r == Lost) {
            you = counter(counter(opponent));
        }
        sum += (r + you); 
    }
    cout << "sum=" << sum << endl;
}


int main()
{
    part1(); 
    part2(); 
}
