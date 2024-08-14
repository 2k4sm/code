import java.util.ArrayList;
import java.util.Random;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

public class msort {

	public static void main(String[] args) throws InterruptedException, ExecutionException {
		ArrayList<Integer> unsortedList = new ArrayList<>();

		Random rand = new Random();
		for (int i = 0; i < 1000; i++) {
			unsortedList.add(rand.nextInt(10000 - 1) + 1);
		}
		System.out.println("Given Array:");
		printArray(unsortedList);

		Sorter sorter = new Sorter(unsortedList);
		ExecutorService es = Executors.newCachedThreadPool();

		Future<ArrayList<Integer>> sortedListFuture = es.submit(sorter);

		// mergeSort(unsortedList, 0, unsortedList.size() - 1);
		System.out.println("\nSorted array:");
		printArray(sortedListFuture.get());
		// printArray(unsortedList);
		es.shutdown();
	}

	public static void printArray(ArrayList<Integer> arr) {
		int n = arr.size();
		for (int i = 0; i < n; i++) {
			System.out.print(arr.get(i) + " ");
		}
		System.out.println();
	}

	public static void mergeSort(ArrayList<Integer> arr, int left, int right) {
		if (left < right) {
			int middle = (left + right) / 2;

			mergeSort(arr, left, middle);
			mergeSort(arr, middle + 1, right);

			merge(arr, left, middle, right);
		}
	}

	public static void merge(ArrayList<Integer> arr, int left, int middle, int right) {
		int n1 = middle - left + 1;
		int n2 = right - middle;

		int[] leftArray = new int[n1];
		int[] rightArray = new int[n2];

		for (int i = 0; i < n1; ++i) {
			leftArray[i] = arr.get(left + i);
		}
		for (int j = 0; j < n2; ++j) {
			rightArray[j] = arr.get(middle + 1 + j);
		}

		int i = 0, j = 0;

		int k = left;
		while (i < n1 && j < n2) {
			if (leftArray[i] <= rightArray[j]) {
				arr.set(k, leftArray[i]);
				i++;
			} else {
				arr.set(k, rightArray[j]);
				j++;
			}
			k++;
		}

		while (i < n1) {
			arr.set(k, leftArray[i]);
			i++;
			k++;
		}

		while (j < n2) {
			arr.set(k, rightArray[j]);
			j++;
			k++;
		}
	}
}

class Sorter implements Callable<ArrayList<Integer>> {

	ArrayList<Integer> listToSort;

	public Sorter(ArrayList<Integer> list) {
		this.listToSort = list;
	}

	@Override
	public ArrayList<Integer> call() throws InterruptedException, ExecutionException {
		if (this.listToSort.size() <= 1) {
			return this.listToSort;
		}

		int mid = this.listToSort.size() / 2;
		ArrayList<Integer> firstHalf = this.getSubArray(this.listToSort, 0, mid - 1);
		ArrayList<Integer> secondHalf =
			this.getSubArray(this.listToSort, mid, this.listToSort.size() - 1);

		ExecutorService es = Executors.newCachedThreadPool();

		Sorter firstSorter = new Sorter(firstHalf);
		Sorter secondSorter = new Sorter(secondHalf);

		Future<ArrayList<Integer>> sortedFirstHalfFuture = es.submit(firstSorter);
		Future<ArrayList<Integer>> sortedSecondHalfFuture = es.submit(secondSorter);

		ArrayList<Integer> sortedFirstHalf = sortedFirstHalfFuture.get();
		ArrayList<Integer> sortedSecondHalf = sortedSecondHalfFuture.get();

		es.shutdown();
		return mergeLists(sortedFirstHalf, sortedSecondHalf);
	}

	public ArrayList<Integer> getSubArray(ArrayList<Integer> listToSort, int sind, int eind) {
		ArrayList<Integer> halfList = new ArrayList<>();

		for (int i = sind; i <= eind; i++) {
			halfList.add(listToSort.get(i));
		}

		return halfList;
	}

	public ArrayList<Integer> mergeLists(
		ArrayList<Integer> firstList,
		ArrayList<Integer> secondList
	) {
		int firstSize = firstList.size();
		int secondSize = secondList.size();
		ArrayList<Integer> mergedList = new ArrayList<>();

		int i = 0;
		int j = 0;

		// Merge elements from both lists until one list is exhausted
		while (i < firstSize && j < secondSize) {
			if (firstList.get(i) <= secondList.get(j)) {
				mergedList.add(firstList.get(i));
				i++;
			} else {
				mergedList.add(secondList.get(j));
				j++;
			}
		}

		// Add remaining elements from firstList (if any)
		while (i < firstSize) {
			mergedList.add(firstList.get(i));
			i++;
		}

		// Add remaining elements from secondList (if any)
		while (j < secondSize) {
			mergedList.add(secondList.get(j));
			j++;
		}

		return mergedList;
	}
}
