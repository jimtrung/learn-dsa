function SelectionSort(arr) {
    for (let i = 0; i < arr.length - 1; i++) {
        let key = i;

        for (let j = i + 1; j < arr.length; j++) {
            if (arr[key] > arr[j]) {
                key = j;
            }
        }

        if (key !== i) {
            let tmp = arr[i];
            arr[i] = arr[key];
            arr[key] = tmp;
        }
    }
}

function testSort(arr, expected) {
    SelectionSort(arr)
    console.assert(JSON.stringify(arr) === JSON.stringify(expected),
        `Failed on: ${JSON.stringify(arr)}\nExpected: ${JSON.stringify(expected)}\nGot: ${JSON.stringify(arr)}`
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

