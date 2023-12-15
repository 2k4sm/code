public class javatest{
	public static void main(String[] Args){
		Node n1 = new Node(5);
		insertAtEnd(n1,7);
		insertAtEnd(n1, 10);
		System.out.println(n1.val);
		System.out.println(n1.next.val);	
		System.out.println(n1.next.next.val);
		Node n2 = insertAtKthPos(n1, 12, 2);
		System.out.println(n2.val);	
		System.out.println(n2.next.val);	
		System.out.println(n2.next.next.val);
	}
	public static Node insertAtStart(Node head,int val){
		Node newHead = new Node(val);

		newHead.next = head;
		head = newHead;
		return head;
	}
	public static Node insertAtEnd(Node head,int val){
		
		Node newNode = new Node(val);
		while(head.next != null){
			head = head.next;
			
		}

		head.next = newNode;


		return head;
	}


	public static Node insertAtKthPos(Node head, int val,int pos){


		Node newNode = new Node(val);
		
		if (pos == 0){
			head = insertAtStart(head, val);
			return head;
		}
		int i = 0;
		while (i != pos-1){
			if (head.next  != null){
				head = head.next;
				i++;
			}else{
				break;
			}

		}

		Node temp = head.next;
		head.next = newNode;
		head.next.next = temp;
		
		return head;
		

	}
}

class Node{
	int val;
	Node next;

	public Node(int v){
		this.val = v;
		this.next = null;
	}
}

