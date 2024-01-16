public class queuearr {
    public static void main(String[] args) {
        queue q = new queue(6);
        q.push(4);
        q.push(14);
        q.push(24);
        q.push(34);
        System.out.println("The peek of the queue before deleting any element " + q.top());
        System.out.println("The size of the queue before deletion " + q.size());
        System.out.println("The first element to be deleted " + q.pop());
        System.out.println("The peek of the queue after deleting an element " + q.top());
        System.out.println("The size of the queue after deleting an element " + q.size());
    }
}

class queue{
    int rear = -1;
    int front = -1;
    int size;
    int cnt = 0;
    int[] arr = new int[this.size];

    public queue(int size){
        this.size = size;
    }
    public void push(int x){
        if (this.cnt == this.size){
            System.out.println("Queue Full");
            System.exit(0);
        }

        if (this.rear == -1){
            this.rear = 0;
            this.front = 0;
        }else{
            this.rear = (this.rear + 1)%this.size;
        }

        this.arr[this.rear] = x;
        System.out.printf("Pushed %d",x);
        this.cnt++;
    }

    public int pop(){
        if (this.front == -1){
            return -1;
        }

        int removed = this.arr[this.front];

        if (cnt == 1){
            this.front = -1;            
            this.rear = -1;
        }else{
            this.front = (this.front + 1)%this.size;
        }
        this.cnt--;
        return removed;

    }

    public int top(){
        if (this.front == -1){
            return -1;
        }

        return this.arr[this.front];
    }

    public int size(){
        return this.cnt;
    }
}
