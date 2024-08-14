public class helloThreads {

	public static void main(String[] args) throws InterruptedException {
		// System.out.println("Hello world From " + Thread.currentThread().getName());
		// HelloWorldPrinter hwp = new HelloWorldPrinter();
		// hwp.start();
		// HelloWorldPrinter hwp2 = new HelloWorldPrinter();
		// hwp2.start();
		//

		for (int i = 1; i <= 100; i++) {
			Print100 ptr = new Print100(i);
			ptr.start();
			Thread.sleep(10);
		}
	}
}

class HelloWorldPrinter extends Thread {

	@Override
	public void run() {
		System.out.println("Hello world " + Thread.currentThread().getName());
	}
}

class Print100 extends Thread {

	int val = -1;

	@Override
	public void run() {
		System.out.println(val + " " + Thread.currentThread().getName());
	}

	public Print100(int val) {
		this.val = val;
	}
}
