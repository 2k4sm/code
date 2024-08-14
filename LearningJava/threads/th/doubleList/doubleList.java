package doubleList;

import java.util.ArrayList;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

public class doubleList {

	public static void main(String[] args) {
		ExecutorService es = Executors.newFixedThreadPool(10);

		ArrayList<Integer> list = new ArrayList<>();

		for (int i = 0; i < 100; i++) {
			list.add(i);
		}

		ArrayListMod task = new ArrayListMod(list);

		Future<ArrayList<Integer>> doubledFuture = es.submit(task);

		ArrayList<Integer> doubled = new ArrayList<>();
		try {
			doubled = doubledFuture.get();
		} catch (InterruptedException e) {
			e.printStackTrace();
		} catch (ExecutionException e) {
			e.printStackTrace();
		}

		printList(doubled);
	}

	public static void printList(ArrayList<Integer> alist) {
		for (int i = 0; i < alist.size(); i++) {
			System.out.println(alist.get(i));
		}
	}
}

class ArrayListMod implements Callable<ArrayList<Integer>> {

	ArrayList<Integer> listToDouble = new ArrayList<>();

	public ArrayListMod(ArrayList<Integer> list) {
		this.listToDouble = list;
	}

	@Override
	public ArrayList<Integer> call() {
		ArrayList<Integer> doubledList = new ArrayList<Integer>();

		for (int i = 0; i < this.listToDouble.size(); i++) {
			doubledList.add(listToDouble.get(i) * 2);
		}

		return doubledList;
	}
}
