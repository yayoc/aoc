#include "../utils.h"

void part1()
{
    ifstream ifs = open_file("input.txt");

    string line;
    int max_val = 0;
    int total = 0;
    while (getline(ifs, line))
    {
        if (line.empty())
        {
            max_val = max(total, max_val);
            total = 0;
            continue;
        }
        total += stoi(line);
    }
    cout << max_val << endl;
}

void part2()
{
    ifstream ifs = open_file("input.txt");

    string line;
    int total = 0;
    vector<int> calories;
    while (getline(ifs, line))
    {
        if (line.empty())
        {
            calories.push_back(total);
            total = 0;
            continue;
        }
        total += stoi(line);
    }

    sort(calories.begin(), calories.end(), greater<int>());

    cout << calories[0] + calories[1] + calories[2] << endl;
}

int main()
{
    part1();
    part2();
}