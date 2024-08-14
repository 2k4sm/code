import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;

public class subadd {

	public static void main(String[] args) throws InterruptedException, ExecutionException {
		Value refval = new Value(0);
		System.out.printf("Expected Value: %s\n", refval.value);
		Adder add = new Adder(refval);
		Subter sub = new Subter(refval);

		ExecutorService es = Executors.newFixedThreadPool(2);
		Future<Void> addvoid = es.submit(add);
		Future<Void> subvoid = es.submit(sub);

		addvoid.get();
		subvoid.get();

		System.out.printf("Got Value: %s\n", refval.value);
		es.shutdown();
	}
}

class Value {

	int value;

	public Value(int val) {
		this.value = val;
	}
}

class Adder implements Callable<Void> {

	Value val;

	public Adder(Value val) {
		this.val = val;
	}

	@Override
	public Void call() {
		for (int i = 1; i <= 10000; i++) {
			this.val.value += i;
		}
		return null;
	}
}

class Subter implements Callable<Void> {

	Value val;

	public Subter(Value val) {
		this.val = val;
	}

	@Override
	public Void call() {
		for (int i = 1; i <= 10000; i++) {
			this.val.value -= i;
		}
		return null;
	}
}
