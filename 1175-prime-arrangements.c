#include <stdio.h>
#include <math.h>
#include <stdbool.h>

int factorial(int n) {
    const int MOD = ((int)pow(10, 9) + 7);

    int f = 1;
    for (int i = 1; i <= n; i++) {
        f = (int)(((long long)(f % MOD) * (i % MOD)) % MOD);
    }
    return f;
}

int numPrimeArrangements(int n) {
    const int MOD = ((int)pow(10, 9) + 7);

    int primeNums = 0;
    for (int i = 2; i <= n; i++) {
        bool isPrime = true;
        for (int j = 2; j <= sqrt(i); j++) {
            if (i % j == 0) {
                isPrime = false;
                break;
            }
        }
        if (isPrime) {
            primeNums++;
        }
    }

    int result =
        (int)(((long long)factorial(primeNums) * factorial(n - primeNums)) %
              MOD);
    return result;
}

int main() {
    printf("%d\n", numPrimeArrangements(100));
}
