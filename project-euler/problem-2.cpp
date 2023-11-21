/**
 * Each new term in the Fibonacci sequence is generated by adding the
 * previous two terms. By starting with 1 and 2, the first 10 terms will be: 1,
 * 2, 3, 5, 8, 13, 21, 34, 55, 89, \dots</p> By considering the terms in the
 * Fibonacci sequence whose values do not exceed four million, find the sum of
 * the even-valued terms.
 */
#include <iostream>
using namespace std;

class Solution {
public:
  /**
   * The function calculates the nth Fibonacci number recursively.
   *
   * @param n The parameter "n" in the above code represents the position of the
   * Fibonacci number to be calculated.
   *
   * @return the nth Fibonacci number.
   */
  int fib(int n) {
    if (n <= 1) {
      return n;
    }
    return fib(n - 1) + fib(n - 2);
  }

  /**
   * The function calculates the sum of even Fibonacci numbers up to a given
   * limit.
   *
   * @param LIMIT The LIMIT parameter represents the maximum value that the sum
   * of even Fibonacci numbers should not exceed.
   *
   * @return the sum of even Fibonacci numbers up to a given limit.
   */
  int sumOfEvenFibs(int LIMIT) {
    int sum = 0;
    int fiboNum = 0;
    int n = 2;
    while (sum < LIMIT) {
      cout << "\t" << n;
      fiboNum = fib(n);
      cout << "\tfiboNum = \t" << fiboNum;
      if (fiboNum % 2 == 0) {
        sum += fiboNum;
      }
      cout << "\tsum = \t" << sum << endl;
      n++;
    }

    cout << "\tValue of n when sum crosses 4 million =\t" << n << endl;
    return sum;
  }
};

int main() {
  Solution solution;
  int answer = solution.sumOfEvenFibs(4000000);
  cout << "\tSum of even fibs less than 4 Million =\t" << answer << endl;
}