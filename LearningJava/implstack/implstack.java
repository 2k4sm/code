public class implstack{
    public static void main(String[] args) {
        stack s = new stack();
        s.push(6);
        s.push(3);
        s.push(7);
        s.push(10);
        s.push(8);
        s.push(3);
        s.push(5);

        System.out.println("Top of the stack before deleting any element " + s.top());
        System.out.println("Size of the stack before deleting any element " + s.size());
        System.out.println("The element deleted is " + s.pop());
        System.out.println("Size of the stack after deleting an element " + s.size());
        System.out.println("Top of the stack after deleting an element " + s.top());
    }
}

class stack{
    int size = 1000;
    int arr[] = new int[size];
    int top = -1;

    public void push(int x){
        this.top++;
        this.arr[top] = x;
    }

    public int pop(){
        int delVal = this.arr[top];

        top--;
        return delVal;
    }

    public int top(){
        return this.arr[top];
    }

    public int size(){
        return this.top+1;
    }
}