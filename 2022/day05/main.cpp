#include "../utils.h"

class Crate
{

public:
    char id;
    Crate(char c) : id(c) {}
};

ostream &operator<<(ostream &os, const Crate &c)
{
    os << c.id;
    return os;
}

class Procedure
{
public:
    int from;
    int to;
    int count;
    Procedure(int c, int f, int t) : from(f), to(t), count(c) {}
};

ostream &operator<<(ostream &os, const Procedure &p)
{
    os << "move " << p.count << " from " << p.from << " to " << p.to;
    return os;
}

template <typename T>
stack<T> reverse_stack(stack<T> &origin)
{
    stack<T> reveresed;
    while (!origin.empty())
    {
        reveresed.push(origin.top());
        origin.pop();
    }
    return reveresed;
}

void reverse_crates(vector<stack<Crate>> &crates)
{
    int i = 0;
    for (auto s : crates)
    {
        crates[i] = reverse_stack(s);
        ++i;
    };
}

template <typename T>
void print_stack(stack<T> &s)
{
    while (!s.empty())
    {
        cout << s.top();
        s.pop();
    }
    cout << endl;
}

template <typename T>
void print_top(const vector<stack<T>> &v)
{
    for (auto s : v)
    {
        if (!s.empty())
        {
            cout << s.top();
        }
    }
    cout << endl;
}

void part1()
{
    ifstream ifs = open_file("input.txt");
    bool procedure = false;
    regex pat{R"(move (\d*) from (\d*) to (\d*))"};
    string line;
    vector<stack<Crate>> crates(9);
    vector<Procedure> procedures;
    while (getline(ifs, line))
    {
        if (line.empty())
        {
            procedure = true;
            continue;
        }
        if (procedure)
        {
            smatch matches;
            if (regex_search(line, matches, pat))
            {
                procedures.push_back(Procedure{stoi(matches[1]), stoi(matches[2]), stoi(matches[3])});
            }
        }
        else
        {
            vector<size_t> indexes = {1, 5, 9, 13, 17, 21, 25, 29, 33};
            size_t stack_i = 0;
            for (auto i : indexes)
            {
                if (line[i] != ' ' && !isdigit(line[i]))
                {
                    Crate c{line[i]};
                    crates[stack_i].push(c);
                }
                ++stack_i;
            }
        }
    }
    // stack needs to be reversed
    reverse_crates(crates);

    // run procedures
    for (auto p : procedures)
    {
        for (int i = 0; i < p.count; ++i)
        {
            if (!crates[p.from - 1].empty())
            {

                Crate c = crates[p.from - 1].top();
                crates[p.from - 1].pop();
                crates[p.to - 1].push(c);
            }
            else
            {
                break;
            }
        }
    }

    print_top(crates);
}

void part2()
{
    ifstream ifs = open_file("input.txt");
    bool procedure = false;
    regex pat{R"(move (\d*) from (\d*) to (\d*))"};
    string line;
    vector<stack<Crate>> crates(9);
    vector<Procedure> procedures;
    while (getline(ifs, line))
    {
        if (line.empty())
        {
            procedure = true;
            continue;
        }
        if (procedure)
        {
            smatch matches;
            if (regex_search(line, matches, pat))
            {
                procedures.push_back(Procedure{stoi(matches[1]), stoi(matches[2]), stoi(matches[3])});
            }
        }
        else
        {
            vector<size_t> indexes = {1, 5, 9, 13, 17, 21, 25, 29, 33};
            size_t stack_i = 0;
            for (auto i : indexes)
            {
                if (line[i] != ' ' && !isdigit(line[i]))
                {
                    Crate c{line[i]};
                    crates[stack_i].push(c);
                }
                ++stack_i;
            }
        }
    }
    // stack needs to be reversed
    reverse_crates(crates);

    // run procedures
    for (auto p : procedures)
    {
        stack<Crate> tmp;
        for (int i = 0; i < p.count; ++i)
        {
            if (!crates[p.from - 1].empty())
            {

                Crate c = crates[p.from - 1].top();
                crates[p.from - 1].pop();
                tmp.push(c);
            }
            else
            {
                break;
            }
        }

        while (!tmp.empty())
        {
            crates[p.to - 1].push(tmp.top());
            tmp.pop();
        }
    }

    print_top(crates);
}

int main()
{
    part1();
    part2();
}
