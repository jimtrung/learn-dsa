#include <iostream>
using namespace std;

void bubbleSort(int a[], int N) {
    for (int i = 0; i < N; i++) {
        for (int j = 0; j < N - i - 1; j++) {
            if (a[j] > a[j + 1]) {
                int tmp = a[j];
                a[j] = a[j + 1];
                a[j + 1] = tmp;
            }
        }
    }
}

void selectionSort(int a[], int N) {
}

int main() {
    int arr[5] = {4, -1, 3, 6, 9};

    bubbleSort(arr, 5);

    for (int a : arr) {
        cout << a << " ";;
    }

    return 0;
}
