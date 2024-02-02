/**
 Write a function that takes the binary representation of an unsigned integer
and returns the number of '1' bits it has (also known as the Hamming weight).

Note:

Note that in some languages, such as Java, there is no unsigned integer type. In
this case, the input will be given as a signed integer type. It should not
affect your implementation, as the integer's internal binary representation is
the same, whether it is signed or unsigned. In Java, the compiler represents the
signed integers using 2's complement notation. Therefore, in Example 3, the
input represents the signed integer. -3.


Example 1:

Input: n = 00000000000000000000000000001011
Output: 3
Explanation: The input binary string 00000000000000000000000000001011 has a
total of three '1' bits.

Example 2:

Input: n = 00000000000000000000000010000000
Output: 1
Explanation: The input binary string 00000000000000000000000010000000 has a
total of one '1' bit.

Example 3:

Input: n = 11111111111111111111111111111101
Output: 31
Explanation: The input binary string 11111111111111111111111111111101 has a
total of thirty one '1' bits.


Constraints:

The input must be a binary string of length 32.


Follow up: If this function is called many times, how would you optimize it?
 *
*/
#include <iostream>
using namespace std;

class Solution {
public:
  uint32_t m1 = 0x55555555;  // binary: 0101...
  uint32_t m2 = 0x33333333;  // binary: 00110011.
  uint32_t m4 = 0x0f0f0f0f;  // binary:  4 zeros,  4 ones ...
  uint32_t m8 = 0x00ff00ff;  // binary:  8 zeros,  8 ones ...
  uint32_t m16 = 0x0000ffff; // binary: 16 zeros, 16 ones ...
  int getZeroOneAsPerBits(int numOfBits) {
    switch (numOfBits) {
    case 2:
      return 1431655765;
      break;
    case 4:
      return 252645135;
      break;
    case 8:
      return 16711935;
      break;
    case 16:
      return 65535;
      break;
    default:
      return 2147483647;
    }
  }
  int helper(uint32_t num, int numOfBits) {
    uint32_t n = num;
    cout << "\tnumOfBits => " << numOfBits << endl;
    cout << "\tnum = n => " << n << endl;
    uint32_t n0 = (n >> 0) & getZeroOneAsPerBits(numOfBits);
    uint32_t n1 = (n >> 1) & getZeroOneAsPerBits(numOfBits);
    uint32_t n2 = n0 + n1;
    cout << "\tn0+n1 = " << n2 << endl;
    return n2;
  }
  int hammingWeight(uint32_t n) { n = () }
};

int main() {
  Solution solution;
  string result;

  int num = 27834;
  cout << "\tEntered number is: " << num;
  printf("\tHexadecimal representation of %d is = ", num);
  result = solution.hammingWeight(num);
  cout << result << endl;

  return 0;
}