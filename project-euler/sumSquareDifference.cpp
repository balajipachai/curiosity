/*
The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385.
The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 55^2 = 3025.
Hence the difference between the sum of the squares of the first ten natural
numbers and the square of the sum is 3025 - 385 = 2640. Find the
difference between the sum of the squares of the first one hundred natural
numbers and the square of the sum.
*/
#include <iostream>

using namespace std;

class Solution {
public:
  /**
   * The function calculates the sum of squares of the first n natural numbers.
   *
   * @param n The parameter "n" represents the number of natural numbers for
   * which we want to calculate the sum of squares.
   *
   * @return the sum of squares of the first n natural numbers.
   */
  unsigned long long sumOfSquares(int n) {
    /**
     * Sum of squares of first n natural numbers is:
     * => n * (n+1) * (2n + 1) / 6
     */
    return ((n * (n + 1) * (2 * n + 1)) / 6);
  }

  /**
   * The function calculates the square of the sum of the first n natural
   * numbers.
   *
   * @param n The parameter "n" represents the number of natural numbers for
   * which we want to calculate the square of the sum.
   *
   * @return the square of the sum of the first n natural numbers.
   */
  unsigned long long squareOfSum(int n) {
    /**
     * Sum of first n natural numbers is:
     * => n * (n + 1) / 2
     */
    return pow((n * (n + 1)) / 2, 2);
  }
};

int main() {
  cout << "\tDifference between the sum of the squares of the first one "
          "hundred natural numbers and the square of the sum"
       << endl;
  Solution solution;
  int sumSquare = solution.sumOfSquares(100);
  int sqaureSum = solution.squareOfSum(100);
  int difference = sqaureSum - sumSquare;
  cout << "\t" << difference << endl;

  cout << "\tEnter the number of test cases" << endl;
  int t;
  cin >> t;
  for (int a0 = 0; a0 < t; a0++) {
    int n;
    cout << "\tEnter the value for n" << endl;
    cin >> n;
    cout << "\tDifference between sum of squares and square of sum = ";
    cout << solution.squareOfSum(n) - solution.sumOfSquares(n) << endl;
  }
  return 0;
}
