package main

import "fmt"

const maxNum = 60

var dpKS [maxNum][maxNum][maxNum][maxNum]int

func maxLengthWithCombineAAndB(a, b string) int {
	la, lb := len(a), len(b)
	ans := 1
	for p := 0; p <= la; p++ {
		for q := 0; q <= lb; q++ {
			for i := 0; i+p-1 < la; i++ {
				for k := 0; k+q-1 < lb; k++ {
					j := i + p - 1
					m := k + q - 1
					if p+q <= 1 {
						dpKS[i+1][j+1][k+1][m+1] = 1
					} else {
						dpKS[i+1][j+1][k+1][m+1] = 0
						if p > 1 && a[i] == a[j] {
							dpKS[i+1][j+1][k+1][m+1] |= dpKS[i+2][j][k+1][m+1]
						}
						if q > 1 && b[k] == b[m] {
							dpKS[i+1][j+1][k+1][m+1] |= dpKS[i+1][j+1][k+2][m]
						}
						if p > 0 && q > 0 && a[i] == b[m] {
							dpKS[i+1][j+1][k+1][m+1] |= dpKS[i+2][j+1][k+1][m]
						}
						if p > 0 && q > 0 && a[j] == b[k] {
							dpKS[i+1][j+1][k+1][m+1] |= dpKS[i+1][j][k+2][m+1]
						}
					}
					if dpKS[i+1][j+1][k+1][m+1] == 1 {
						ans = max(ans, p+q)
					}
				}
			}
		}
	}
	return ans
}

func main() {
	t := 0
	_, _ = fmt.Scanf("%d", &t)
	var a, b string
	for i := 0; i < t; i++ {
		_, _ = fmt.Scanf("%s", &a)
		_, _ = fmt.Scanf("%s", &b)
		fmt.Println(maxLengthWithCombineAAndB(a, b))
	}
}

/*
#include <algorithm>
#include <iostream>
#include <string>
using namespace std;

const int maxn = 51;

int dp[maxn][maxn][maxn][maxn];

int get(string a, string b) {
    int len1 = a.size();
    int len2 = b.size();
    int ans = -1;
    for (int l1 = 0; l1 <= len1; ++l1) {
        for (int l2 = 0; l2 <= len2; ++l2) {
            for (int i = 0; i + l1 - 1 < len1; ++i) {
                for (int k = 0; k + l2 - 1 < len2; ++k) {
                    int j = i + l1 - 1;
                    int m = k + l2 - 1;
                    if (l1 + l2 <= 1) {
                        dp[i + 1][j + 1][k + 1][m + 1] = 1;
                    } else {
                        dp[i + 1][j + 1][k + 1][m + 1] = 0;
                        if (l1 > 1 && a[i] == a[j]) {
           					dp[i + 1][j + 1][k + 1][m + 1] |= dp[i + 2][j][k + 1][m + 1];
                        }
                        if (l2 > 1 && b[k] == b[m]) {
                			dp[i + 1][j + 1][k + 1][m + 1] |= dp[i + 1][j + 1][k + 2][m];
                        }
                        if (l1 && l2 && a[i] == b[m]) {
							dp[i + 1][j + 1][k + 1][m + 1] |= dp[i + 2][j + 1][k + 1][m];
                        }
                        if (l1 && l2 && a[j] == b[k]) {
              				dp[i + 1][j + 1][k + 1][m + 1] |= dp[i + 1][j][k + 2][m + 1];
                        }
                    }
                    if (dp[i + 1][j + 1][k + 1][m + 1] == 1) {
                        ans = max(ans, l1 + l2);
                    }
                }
            }
        }
    }
    return ans;
}

int main() {
    int t = 0;
    cin >> t;
    string a = "";
    string b = "";
    for (int i = 0; i < t; ++i) {
        cin >> a;
        cin >> b;
        cout << get(a, b) << endl;
    }
}
*/
