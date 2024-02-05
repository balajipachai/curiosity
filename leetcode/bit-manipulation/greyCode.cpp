/*
An n-bit gray code sequence is a sequence of 2n integers where:

Every integer is in the inclusive range [0, 2n - 1],
The first integer is 0,
An integer appears no more than once in the sequence,
The binary representation of every pair of adjacent integers differs by exactly
one bit, and The binary representation of the first and last integers differs by
exactly one bit. Given an integer n, return any valid n-bit gray code sequence.



Example 1:

Input: n = 2
Output: [0,1,3,2]
Explanation:
The binary representation of [0,1,3,2] is [00,01,11,10].
- 00 and 01 differ by one bit
- 01 and 11 differ by one bit
- 11 and 10 differ by one bit
- 10 and 00 differ by one bit
[0,2,3,1] is also a valid gray code sequence, whose binary representation is
[00,10,11,01].
- 00 and 10 differ by one bit
- 10 and 11 differ by one bit
- 11 and 01 differ by one bit
- 01 and 00 differ by one bit
Example 2:

Input: n = 1
Output: [0,1]


Constraints:

1 <= n <= 16
*/
#include <iostream>
#include <vector>
using namespace std;

/**
 * The function "print" takes a vector of integers as input and prints each
 * element followed by a space.
 *
 * @param vc The parameter "vc" is a reference to a vector of integers.
 */
void print(vector<int> &vc, int n) {
  cout << "\t";
  for (int i = 0; i < n; i++) {
    cout << vc[i] << " ";
  }
  cout << endl;
}

class Solution {
public:
  /**
   * The function `greyCode` generates a sequence of grey code numbers of length
   * `n`.
   *
   * @param n The parameter `n` represents the number of bits in the Grey code
   * sequence.
   *
   * @return a vector of integers, which represents the Grey code sequence of
   * length n.
   */
  vector<int> greyCode(int n) {
    int sequenceLength = pow(2, n);
    vector<int> result;
    for (int i = 0; i < sequenceLength; i++) {
      int greNum = i ^ i >> 1;
      result.push_back(greNum);
    }
    return result;
  }
};

int main() {
  Solution solution;
  int n = 2;
  vector<int> result;
  cout << "\tFor n = " << n << " the grey numbers are:" << endl;
  result = solution.greyCode(n);
  print(result, result.size());

  n = 3;
  cout << "\tFor n = " << n << " the grey numbers are:" << endl;
  result = solution.greyCode(n);
  print(result, result.size());

  n = 4;
  cout << "\tFor n = " << n << " the grey numbers are:" << endl;
  result = solution.greyCode(n);
  print(result, result.size());

  cout << "\tTIME COMPLEXITY: 2^n" << endl;
  return 0;
}
