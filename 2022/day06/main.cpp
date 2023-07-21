#include "../utils.h"

bool is_unique_chars(const string& s)
{
    unordered_map<char, bool> m;
    for (auto &c : s)
    {
        if (m[c])
        {
            return false;
        }
        else
        {
            m[c] = true;
        }
    }
    return true;
}

int marker(const string& s, int len) {
    int i = 0;
    while(!is_unique_chars(s.substr(i, len))) {
        ++i;
    }
    return i + len;
}


void part1()
{
    ifstream ifs = open_file("input.txt");
    string line;
    ifs >> line;
    cout << "marker=" << marker(line, 4) << endl;
}

void part2()
{
    ifstream ifs = open_file("input.txt");
    string line;
    ifs >> line;
    cout << "marker=" << marker(line, 14) << endl;
}

int main()
{
    part1();
    part2();
}
