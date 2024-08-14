import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Semaphore;

public class store {

	public static void main(String[] args) {
		Store store = new Store(0, 5);
		Semaphore psem = new Semaphore(5);
		Semaphore csem = new Semaphore(0);
		ExecutorService es = Executors.newCachedThreadPool();

		for (int i = 0; i < 100; i++) {
			Producer producer = new Producer(psem, csem, store, "producer :" + i);
			es.execute(producer);
		}

		for (int i = 0; i < 100; i++) {
			Consumer consumer = new Consumer(psem, csem, store, "consumer :" + i);
			es.execute(consumer);
		}

		es.shutdown();
	}
}

class Producer implements Runnable {

	private Store store;
	public String pname;
	public Semaphore psem;
	public Semaphore csem;

	public Producer(Semaphore psem, Semaphore csem, Store store, String pname) {
		this.store = store;
		this.pname = pname;
		this.psem = psem;
		this.csem = csem;
	}

	@Override
	public void run() {
		while (true) {
			try {
				psem.acquire();
			} catch (InterruptedException ex) {
				ex.printStackTrace();
			}
			if (store.currSize < store.getMaxSize()) {
				System.out.println("Being produced by:" + this.pname);
				store.currSize++;
				System.out.println(store.currSize);
			}
			csem.release();
		}
	}
}

class Consumer implements Runnable {

	private Store store;
	public String cname;
	public Semaphore psem;
	public Semaphore csem;

	public Consumer(Semaphore psem, Semaphore csem, Store store, String cname) {
		this.store = store;
		this.cname = cname;
		this.psem = psem;
		this.csem = csem;
	}

	@Override
	public void run() {
		while (true) {
			try {
				csem.acquire();
			} catch (InterruptedException ex) {
				ex.printStackTrace();
			}
			if (store.currSize > 0) {
				System.out.println("Being consumed by:" + this.cname);
				store.currSize--;
				System.out.println(store.currSize);
			}
			psem.release();
		}
	}
}

class Store {

	public int currSize;
	private int maxSize;

	public Store(int currSize, int maxSize) {
		this.currSize = currSize;
		this.maxSize = maxSize;
	}

	public int getMaxSize() {
		return this.maxSize;
	}
}
