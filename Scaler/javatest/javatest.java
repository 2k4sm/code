public class javatest{
	public static void main(String[] args) {
		Stack newstack = new Stack(7);

		newstack.push(10);
		newstack.push(12);
		newstack.push(13);
		newstack.push(14);
		newstack.push(15);
		newstack.push(16);
		newstack.push(17);
		newstack.pop();
		//newstack.push(1);
		System.out.println(newstack.isFull());


	}

}




class Stack{
	int[] arr;
	int capacity;
	int size;

	public Stack(int capacity){
		this.capacity = capacity;
		this.arr = new int[capacity];
		this.size = 0;
	}


	public void push(int val){

		if (this.size < this.capacity){
			this.arr[size] = val;
			size++;
		}else{
			System.out.println("No capacity left to push");
		}
	}

	public int pop(){
		if (this.size > 0){
			capacity--;
		}else{
			return -1;
		}
		return this.arr[capacity - 1];
	}

	public int size(){
		return this.size;
	}

	public boolean isFull(){
		if (this.size == this.capacity){
			return true;
		}
		return false;
	}
}

class Node{
	Node next;
	Node prev;
	int val;
}

class linkedStack{

}
