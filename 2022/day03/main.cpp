#include "../utils.h"

string common_chars(const string &str1, const string &str2)
{
    unordered_set<char> set1(str1.begin(), str1.end());
    unordered_set<char> set2(str2.begin(), str2.end());

    string common = "";
    for (auto c : set1)
    {
        if (set2.find(c) != set2.end())
        {
            common += c;
        }
    }
    return common;
}

string common_chars(const string &str1, const string &str2, const string &str3)
{
    string common = common_chars(str1, str2);
    return common_chars(common, str3);
}

int priority(const char &c)
{
    return isupper(c) ? static_cast<int>(c) - 38 : static_cast<int>(c) - 96;
}

class Priority
{
public:
    int operator()(int sum, char c) const { return sum + priority(c); }
};

void part1()
{
    ifstream ifs = open_file("input.txt");
    string line;
    int sum = 0;
    while (getline(ifs, line))
    {
        size_t len = line.length();
        string first = line.substr(0, len / 2);
        string second = line.substr(len / 2, len / 2);
        string common = common_chars(first, second);
        Priority p;
        sum += accumulate(common.begin(), common.end(), 0, p);
    }
    cout << "sum=" << sum << endl;
}

void part2()
{
    ifstream ifs = open_file("input.txt");
    string line;
    int i = 0;
    int sum = 0;
    string first, second, third;
    while (getline(ifs, line))
    {
        ++i;
        if (i % 3 == 1)
        {
            first = line;
        }
        else if (i % 3 == 2)
        {
            second = line;
        }
        else if (i % 3 == 0)
        {
            third = line;
            string common = common_chars(first, second, third);
            Priority p;
            sum += accumulate(common.begin(), common.end(), 0, p);
        }
    }
    cout << "sum=" << sum << endl;
}

int main()
{
    part1();
    part2();
}