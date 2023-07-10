#include <map>

#include "../utils.h"

const int rock = 1;
const int paper = 2;
const int scissors = 3;

map<string, int> elements = {
    {"A", rock},
    {"X", rock},
    {"B", paper},
    {"Y", paper},
    {"C", scissors},
    {"Z", scissors},
};

const int lost = 0;
const int draw = 3;
const int win = 6;

map<string, int> means = {
    {"X", lost},
    {"Y", draw},
    {"Z", win},
};

// a: you, b: opponent
int score(int a, int b)
{
    if (a == b)
        return draw + a;

    switch (a)
    {
    case rock:
        switch (b)
        {
        case paper:
            return lost + a; 
        case scissors:
            return win + a; 
        }
    case paper:
        switch (b)
        {
        case rock:
            return win + a; 
        case scissors:
            return lost + a; 
        }
    case scissors:
        switch (b)
        {
        case rock:
            return lost + a; 
        case paper:
            return win + a; 
        }
        break;
    }

    error("unexpected args");
}

void part1() {
    ifstream ifs = open_file("input.txt");

    int sum = 0;
    string e;
    string line;
    while (getline(ifs, line))
    {
        istringstream iss(line);
        string o, y;
        if (!(iss >> o >> y)) { break; } 
        sum += score(elements.at(y), elements.at(o));
    }
    cout << "sum=" << sum << endl;
}

// a: opponent, b: mean
int score2(int a, int b) {
    if (b == draw) 
        return draw + a;

    if (b == lost) {
        switch (a) {
            case rock:
                return lost + scissors;
            case scissors:
                return lost + paper;
            case paper:
                return lost + rock;
        }
    }

    if (b == win) {
        switch (a) {
            case rock:
                return win + paper;
            case scissors:
                return win + rock;
            case paper:
                return win + scissors;
        }
    }

    error("unexpected args");
}

void part2() {
   ifstream ifs = open_file("input.txt");

    int sum = 0;
    string e;
    string line;
    while (getline(ifs, line))
    {
        istringstream iss(line);
        string a, b;
        if (!(iss >> a >> b)) { break; } 
        sum += score2(elements.at(a), means.at(b));
    }
    cout << "sum=" << sum << endl;
}


int main()
{
    part1(); 
    part2(); 
}
