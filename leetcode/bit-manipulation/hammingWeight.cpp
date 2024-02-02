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

  /**
   * The function calculates the Hamming weight of a 32-bit unsigned integer,
   * which is the number of set bits in its binary representation.
   *
   * @param n The parameter `n` is an unsigned 32-bit integer.
   *
   * @return the number of bits that are set to 1 in the input number `n`.
   */
  int hammingWeight(uint32_t n) {
    n = (n & m1) + ((n >> 1) & m1);
    n = (n & m2) + ((n >> 2) & m2);
    n = (n & m4) + ((n >> 4) & m4);
    n = (n & m8) + ((n >> 8) & m8);
    n = (n & m16) + ((n >> 16) & m16);
    return n;
  }
};

int main() {
  Solution solution;
  int result;

  int num = 27834;
  cout << "\tEntered number is: " << num << endl;
  cout << "\tBinary representation of given number is: "
       << bitset<32>(num).to_string() << endl;
  printf("\tHamming weight (Number of 1s) in %d = ", num);
  result = solution.hammingWeight(num);
  cout << result << endl;

  cout << endl;

  num = 3;
  cout << "\tEntered number is: " << num << endl;
  cout << "\tBinary representation of given number is: "
       << bitset<32>(num).to_string() << endl;
  printf("\tHamming weight (Number of 1s) in %d = ", num);
  result = solution.hammingWeight(num);
  cout << result << endl;

  cout << endl;

  num = -3;
  cout << "\tEntered number is: " << num << endl;
  cout << "\tBinary representation of given number is: "
       << bitset<32>(num).to_string() << endl;
  printf("\tHamming weight (Number of 1s) in %d = ", num);
  result = solution.hammingWeight(num);
  cout << result << endl;

  return 0;
}