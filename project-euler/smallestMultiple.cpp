/**
2520 is the smallest number that can be divided by each of the numbers
from 1 to 10 without any remainder.
What is the smallest positive number that isevenly divisible "divisible with no
remainder" by all of the numbers from 1 to 20?
*/
#include <iostream>

using namespace std;

class Solution {
public:
  /**
   * The function calculates the greatest common divisor (GCD) of two unsigned
   * long long integers.
   *
   * @param a The parameter "a" is an unsigned long long integer.
   * @param b The parameter "b" represents the second number for which we want
   * to find the greatest common divisor (GCD).
   *
   * @return the greatest common divisor (gcd) of the two input numbers.
   */
  unsigned long long gcd(unsigned long long a, unsigned long long b) {
    while (a != 0) {
      unsigned long long c = a;
      a = b % a;
      b = c;
    }
    return b;
  }

  /**
   * The function calculates the least common multiple (LCM) of two numbers
   * using their greatest common divisor (GCD).
   *
   * @param a The parameter "a" is an unsigned long long integer.
   * @param b The parameter "b" represents the second number for which we want
   * to find the least common multiple (LCM).
   *
   * @return the least common multiple (LCM) of the two input numbers.
   */
  unsigned long long lcm(unsigned long long a, unsigned long long b) {
    return (a * b) / gcd(a, b);
  }
};

int main() {
  Solution solution;
  unsigned long long answer = 1;
  for (unsigned int i = 2; i <= 20; i++) {
    answer = solution.lcm(answer, i);
  }
  cout << "\tSmallest multiple is " << answer << endl;
  return 0;
}