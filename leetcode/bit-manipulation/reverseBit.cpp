/**
 * 190. Reverse Bits
Easy
Topics
Companies
Reverse bits of a given 32 bits unsigned integer.

Note:

Note that in some languages, such as Java, there is no unsigned integer type. In
this case, both input and output will be given as a signed integer type. They
should not affect your implementation, as the integer's internal binary
representation is the same, whether it is signed or unsigned. In Java, the
compiler represents the signed integers using 2's complement notation.
Therefore, in Example 2 above, the input represents the signed integer -3 and
the output represents the signed integer -1073741825.


Example 1:

Input: n = 00000010100101000001111010011100
Output:    964176192 (00111001011110000010100101000000)
Explanation: The input binary string 00000010100101000001111010011100 represents
the unsigned integer 43261596, so return 964176192 which its binary
representation is 00111001011110000010100101000000. Example 2:

Input: n = 11111111111111111111111111111101
Output:   3221225471 (10111111111111111111111111111111)
Explanation: The input binary string 11111111111111111111111111111101 represents
the unsigned integer 4294967293, so return 3221225471 which its binary
representation is 10111111111111111111111111111111.


Constraints:

The input must be a binary string of length 32


Follow up: If this function is called many times, how would you optimize it?
*/

#include <iostream>
using namespace std;

class Solution {
public:
  uint32_t reverseBits(uint32_t n) {
    /*
    To reverse bits of an unsigned 32 bit integer
        The bit at position 0 will be at position 31-0
        The bit at position 1 will be at position 31-1
    And so on...
    We have to get the rightmost bit of the integer using n & 1
    We start with n & 1 to get the rightmost bit and << by power which starts
    with 31 Then change n = n >> 1; Move n to the next bit Decrement power by 1
    Do this as long as n != 0
    */
    uint32_t reverse = 0, power = 31;
    /*
    The code is reversing the bits of the given 32-bit unsigned integer `n`.
     */
    while (n != 0) {
      reverse += (n & 1) << power;
      n = n >> 1;
      power -= 1;
    }
    return reverse;
  }
};

int main() {
  Solution solution;
  uint32_t result;

  int num = 27834;
  cout << "\tEntered number is: " << num << endl;
  cout << "\tBinary representation of given number is: "
       << bitset<32>(num).to_string() << endl;
  result = solution.reverseBits(num);
  cout << "\tReversed bits of given number is: "
       << bitset<32>(result).to_string() << endl;

  cout << endl;

  num = 3;
  cout << "\tEntered number is: " << num << endl;
  cout << "\tBinary representation of given number is: "
       << bitset<32>(num).to_string() << endl;
  result = solution.reverseBits(num);
  cout << "\tReversed bits of given number is: "
       << bitset<32>(result).to_string() << endl;

  cout << endl;

  num = -3;
  cout << "\tEntered number is: " << num << endl;
  cout << "\tBinary representation of given number is: "
       << bitset<32>(num).to_string() << endl;
  result = solution.reverseBits(num);
  cout << "\tReversed bits of given number is: "
       << bitset<32>(result).to_string() << endl;

  return 0;
}