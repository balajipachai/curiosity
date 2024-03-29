/**
 * Each new term in the Fibonacci sequence is generated by adding the
 * previous two terms. By starting with 1 and 2, the first 10 terms will be: 1,
 * 2, 3, 5, 8, 13, 21, 34, 55, 89, \dots</p> By considering the terms in the
 * Fibonacci sequence whose values do not exceed four million, find the sum of
 * the even-valued terms.
 */
#include <iomanip>
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

  /**
   * The function calculates the sum of all even Fibonacci numbers up to a given
   * number.
   *
   * @param n The parameter `n` represents the upper limit or the maximum value
   * of the Fibonacci sequence that we want to consider. The function
   * `sumOfEvenFibs` calculates the sum of all even Fibonacci numbers up to and
   * including `n`.
   *
   * @return the sum of all even Fibonacci numbers less than or equal to the
   * input number `n`.
   */
  unsigned long long sumOfEvenFibs(unsigned long long n) {
    unsigned long long sum = 0;
    unsigned long long a = 1;
    unsigned long long b = 2;

    while (b <= n) {
      if (b % 2 == 0)
        sum += b;
      auto next = a + b;
      a = b;
      b = next;
    }
    return sum;
  }
};

int main() {
  Solution solution;

  time_t start, end;
  time(&start);
  // unsync the I/O of C and C++.
  ios_base::sync_with_stdio(false);

  int answer = solution.sumOfEvenFibs(400000000);
  cout << "\tSum of even fibs less than 4 Million =\t" << answer << endl;
  // Recording end time.
  time(&end);

  // Calculating total time taken by the program.
  double time_taken = double(end - start);
  cout << "Time taken by program 1 is : " << fixed << time_taken
       << setprecision(5);
  cout << " sec " << endl;

  cout << "Enter the value of n\t";
  unsigned long long n;
  cin >> n;
  time(&start);
  answer = solution.sumOfEvenFibs(n);
  cout << "\tSum of even fibs less than 4 Million =\t" << answer << endl;
  time_taken = double(end - start);
  cout << "Time taken by program 2 is : " << fixed << time_taken
       << setprecision(5);
  cout << " sec " << endl;
}