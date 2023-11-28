/**
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
 * that the 6th prime is 13.
 * What is the 10001st prime number?
 */
#include <iomanip>
#include <iostream>

using namespace std;

void print(vector<int> &vc) {
  cout << "\t";
  for (auto i : vc) {
    cout << i << " ";
  }
  cout << endl;
}

class Solution {
public:
  /**
   * The function checks if a given number is prime or not.
   *
   * @param num The parameter "num" is a long integer that we want to check if
   * it is a prime number.
   *
   * @return The function isPrime is returning a boolean value. It returns true
   * if the input number is prime, and false otherwise.
   */
  bool isPrime(long num) {
    int i = 2;
    while (i * i <= num) {
      if (num % i == 0) {
        return false;
      }
      i += 1;
    }
    return true;
  }

  /**
   * The function "xthPrimeNumber" returns the xth prime number.
   *
   * @param x The parameter "x" represents the position of the prime number that
   * you want to find. For example, if you pass in 5 for "x", the function will
   * return the 5th prime number.
   *
   * @return The xth prime number is being returned.
   */
  long xthPrimeNumber(int x) {
    int count = 1;
    long primeNumber = 2;
    while (count <= x) {
      if (isPrime(primeNumber)) {
        count += 1;
      }
      primeNumber += 1;
    }
    return primeNumber - 1;
  }
};

int main() {
  vector<int> primes;
  primes.reserve(10001);

  primes.push_back(2);
  for (int i = 3; primes.size() <= 10000; i += 2) {
    bool isPrime = true;
    for (auto p : primes) {
      if (i % p == 0) {
        isPrime = false;
        break;
      }

      if (p * p > i) {
        break;
      }
    }
    if (isPrime) {
      primes.push_back(i);
    }
  }

  print(primes);

  Solution solution;
  int x;
  cout << "\tEnter the value for x: ";
  cin >> x;
  time_t start, end;
  time(&start);
  // unsync the I/O of C and C++.
  ios_base::sync_with_stdio(false);
  long answer = solution.xthPrimeNumber(x);
  printf("\t%d st/th prime number is %ld\n", x, answer);
  // Recording end time.
  time(&end);

  // Calculating total time taken by the program.
  double time_taken = double(end - start);
  cout << "\tTime taken : " << fixed << time_taken << setprecision(5);
  cout << " sec " << endl;

  printf("\t%d st/th prime number is %d\n", x, primes[x - 1]);
}