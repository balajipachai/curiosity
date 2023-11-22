/**
 * The prime factors of 13195 are 5, 7, 13 and 29.
 * What is the largest prime factor of the number 600851475143?
 */

#include <iostream>

using namespace std;

class Solution {
public:
  /**
   * The function "largestPrimeFactor" finds the largest prime factor of a given
   * number.
   *
   * @param num The parameter "num" represents the number for which we want to
   * find the largest prime factor.
   *
   * @return the largest prime factor of the given number.
   */
  long largestPrimeFactor(long num) {
    long largestPrimeFactor = 0;
    // Number of 2's that divide n
    while (num % 2 == 0)
      num /= 2;

    // n must be odd at this point, hence, we can skip 1 element, thus i+=2
    for (int i = 3; i <= sqrt(num); i += 2) {
      // while i divides n, print i and divide n
      while (num % i == 0) {
        if (i > largestPrimeFactor) {
          largestPrimeFactor = i;
        }
        num /= i;
      }
    }

    // To handle the case when n is a prime number greater than 2
    if (num > 2) {
      if (num > largestPrimeFactor) {
        largestPrimeFactor = num;
      }
    }
    return largestPrimeFactor;
  }
};

int main() {
  int num = 13195;
  cout << "\tPrinting the num: " << num << endl;
  cout << "\tLargest prime factor of above num is:\n";
  Solution solution;
  cout << "\t" << solution.largestPrimeFactor(num) << endl;

  long longNum = 600851475143;
  cout << "\tPrinting the num: " << longNum << endl;
  cout << "\tLargest prime factor of above num is:\n";
  cout << "\t" << solution.largestPrimeFactor(longNum) << endl;
}