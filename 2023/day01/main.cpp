#include "../utils.h"

void part1()
{
    ifstream ifs = open_file("input.txt");

    string line;
    int sum = 0;
    while (getline(ifs, line))
    {
        string first, last;
        for (auto c: line) 
        {
            if (isdigit(c)) {
                first = c;
                break;
            }
        }

        for (auto it = line.rbegin(); it != line.rend(); ++it) 
        {
            if (isdigit(*it)) {
                last = *it;
                break;
            }
        }

        sum += stoi(first + last);
    }
    cout << sum << endl;
}

void part2()
{
    ifstream ifs = open_file("input.txt");

    string line;
    int sum = 0;
    unordered_map<string, int> nums = {
        {"0", 0}, {"1", 1}, {"2", 2}, {"3", 3}, {"4", 4},
        {"5", 5}, {"6", 6}, {"7", 7}, {"8", 8}, {"9", 9},
        {"zero", 0}, {"one", 1}, {"two", 2}, {"three", 3}, {"four", 4},
        {"five", 5}, {"six", 6}, {"seven", 7}, {"eight", 8}, {"nine", 9}
    };

    while (getline(ifs, line))
    {
        if (line.empty()) {
            break;
        }

        int first, last;
        int first_pos = line.length();
        int last_pos = 0;

        for (const auto& pair : nums) {
            size_t f_pos = line.find(pair.first);
            size_t l_pos = line.rfind(pair.first);

            if (f_pos != string::npos && f_pos < first_pos) {
                first_pos = f_pos;
                first = pair.second;
            }
            if (l_pos != string::npos && l_pos > last_pos) {
                last_pos = l_pos;
                last = pair.second;
            }
        }
        sum += (first * 10 + last);
    }
    cout << sum << endl;
}



int main() 
{
    // part1();
    part2();
}

