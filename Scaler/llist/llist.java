public class llist{
    public static void main(String[] args) {
        Node n1 = new Node(5);
		insertAtEnd(n1,7);
		insertAtEnd(n1, 16);
        insertAtEnd(n1, 14);
		insertAtEnd(n1, 12);
		insertAtEnd(n1, 18);
		insertAtEnd(n1, 19);
        printLList(n1);
        Node n2 = reverse(n1, 6);
        printLList(n2);

        

    }

    static class Node{
        int val;
        Node next;
    
        public Node(int v){
            this.val = v;
            this.next = null;
        }
    }

    public static void printLList(Node head){
        while (head.next != null){
            System.out.print(head.val+" ");
            head = head.next;
        }
        System.out.print(head.val);
        System.out.println();
    }

    public static Node reverse(Node head,int k){
        Node prev = null;
        Node current = head;
        Node next = head.next;
        int j = 0;
        while (current != null && j < k){
            next = current.next;
            current.next = prev;
            prev = current;
            current = next;
            j++;
        }

        return prev;
    }   

    public static Node insertAtEnd(Node head,int val){
		
		Node newNode = new Node(val);
		while(head.next != null){
			head = head.next;
			
		}

		head.next = newNode;


		return head;
	}

}