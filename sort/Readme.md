//The performance of quicksort depends on the pivot you choose.
//If you always choose the first item as pivot, and the array is already sorted, the big O will be n^2. Why? Because The pivot is always going to be smaller
//compared to the rest of the array, so for each pivot you iterate the whole array. This is the worst-case scenario.
//If you pick the middle element as the pivot, and the array is already sorted, then on every iteraction you are splitting the array in half. That gives an O(nlogn)
//Do I have to choose always the middle element? No, you don't. If you always choose a random element in the array as the pivot, in average the big O will be O(nlogn)