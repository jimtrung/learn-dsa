package selectionsort

// Find the min value
func SelectionSort(arr []int) {
    for i := 0; i < len(arr) - 1; i++ {
        key := i;

        for j := i + 1; j < len(arr); j++ {
            if arr[key] > arr[j] {
                key = j;
            }
        }

        if key != i {
            arr[key], arr[i] = arr[i], arr[key]
        }
    }
}
