package arrays

func binarySearch(arr []int, x int)  int{
	return binarySearchInternal(arr,0,len(arr),x)
}

func binarySearchInternal(arr []int, low int, high int, x int) int {
	if low > high {
		return -1
	}

	mid := low + (high - low)/2

	if arr[mid] == x {
		return mid
	} else if arr[mid] > x {
		return binarySearchInternal(arr,low,mid-1,x)
	} else {
		return binarySearchInternal(arr,mid+1,high,x)
	}
}
