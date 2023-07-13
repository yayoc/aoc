#include "../utils.h"

class Section
{
public:
    int l;
    int r;
    Section(int a, int b) : l(a), r(b) {}
};

ostream &operator<<(ostream &os, const Section &s)
{
    os << s.l << "-" << s.r;
    return os;
}

bool is_fully_covered(const Section &s1, const Section &s2)
{
    if (s1.l == s2.l)
        return true;

    if (s1.l < s2.l)
    {
        return s2.r <= s1.r;
    }
    return s1.r <= s2.r;
}

bool is_overlap(const Section &s1, const Section &s2)
{
    if (s1.l == s2.l)
        return true;

    if (s1.l < s2.l)
    {
        return s2.l <= s1.r;
    }
    return s1.l <= s2.r;
}

void count(const function<bool(Section &, Section &)> &func)
{
    ifstream ifs = open_file("input.txt");

    int count = 0;
    string line;
    regex pat{R"((\d*)-(\d*),(\d*)-(\d*))"};
    while (getline(ifs, line))
    {
        smatch matches;
        if (regex_search(line, matches, pat))
        {
            Section pair1(stoi(matches[1]), stoi(matches[2]));
            Section pair2(stoi(matches[3]), stoi(matches[4]));
            if (func(pair1, pair2))
            {
                ++count;
            }
        }
    }
    cout << "count=" << count << endl;
}

void part1()
{
    count(is_fully_covered);
}

void part2()
{
    count(is_overlap);
}

int main()
{
    part1();
    part2();
}
