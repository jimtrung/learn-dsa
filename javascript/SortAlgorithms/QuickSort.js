function swap(arr, i, j) {
    const tmp = arr[i];
    arr[i] = arr[j];
    arr[j] = tmp;
}

function partition(arr, low, high) {
    const mid = Math.floor((low + high) / 2);

    let pivot = arr[mid];
    swap(arr, mid, high);

    let left = low, right = high - 1;
    while (left <= right) {
        while (left <= right && arr[left] < pivot) {
            left++;
        }
        while (left <= right && arr[right] > pivot) {
            right--;
        }
        if (left <= right) {
            swap(arr, left, right);
            left++;
            right--;
        }
    }
    swap(arr, left, high);

    return left;
}

function QuickSort(arr, low, high) {
    if (low < high) {
        const pivot = partition(arr, low, high);

        QuickSort(arr, low, pivot - 1);
        QuickSort(arr, pivot + 1, high);
    }
}

function testSort(arr, expected) {
    const originArr = [...arr];
    QuickSort(arr, 0, arr.length - 1);
    console.assert(JSON.stringify(arr) === JSON.stringify(expected),
        `Array: ${JSON.stringify(originArr)}\nExpected: ${JSON.stringify(expected)}\nGot: ${JSON.stringify(arr)}\n`
    );
}

// Test Cases
testSort([42, -5, 0, 16, 8, -23, 42, 99, 7, 3, -8, 1, 88, 0, -42],
    [-42, -23, -8, -5, 0, 0, 1, 3, 7, 8, 16, 42, 42, 88, 99]);

testSort([], []); // Edge case: empty array

testSort([5], [5]); // Edge case: single element

testSort([1, 2, 3, 4, 5, 6, 7, 8, 9], [1, 2, 3, 4, 5, 6, 7, 8, 9]); // Already sorted

testSort([9, 8, 7, 6, 5, 4, 3, 2, 1], [1, 2, 3, 4, 5, 6, 7, 8, 9]); // Reverse sorted

testSort([7, 7, 7, 7, 7], [7, 7, 7, 7, 7]); // All same values

console.log("All tests passed if no assertion errors appear!");


