public class heap{
    public static void main(String[] args) {
        int[] heapval = {-1,2,1,3,5,6,4,7,10};

        Heap newHeap = new Heap(9);

        newHeap.val = heapval;

        newHeap.deleteMin();

        for (int i = 0; i < newHeap.size;i++){
            System.out.print(newHeap.val[i] + " ");
        }
        System.out.println();
    }
}


class Heap{
    int[] val;
    int size;

    public Heap(int n){
        this.val = new int[n];
        this.size = n;
    }


    public int deleteMin(){
        int min = this.val[0];
        swap(0,this.size - 1);
        this.size--;
        checkOrder();
        return min;

    }

    private void swap(int first, int last){
        int temp = this.val[last];
        this.val[last] = this.val[first];
        this.val[first] = temp; 
    }

    private void checkOrder() {
        int i = 0;

        while (((2 * i) + 1) < this.size && ((2 * i) + 2) < this.size) {
            if (this.val[(2 * i) + 1] > this.val[(2 * i) + 2]) {
                swap(i, (2*i) + 2);
                i = (2 * i) + 2;
            }else{
                swap(i, (2 * i) + 1);
                i = (2 * i) + 1;
            }
        }
    }


}