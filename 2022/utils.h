#include <iostream>
#include <fstream>
#include <string>
#include <algorithm>

using namespace std;

inline void error(const string &s)
{
    throw runtime_error(s);
}

inline ifstream open_file(const string &filename)
{
    ifstream ifs(filename);
    if (!ifs)
    {
        error("can't open " + filename);
    }
    return std::move(ifs);
}