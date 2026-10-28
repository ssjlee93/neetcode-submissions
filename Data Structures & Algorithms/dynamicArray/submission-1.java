class DynamicArray {
    int[] a;
    int capacity;
    int len;

    public DynamicArray(int capacity) {
        this.a = new int[capacity];
        this.capacity = capacity;
    }

    public int get(int i) {
        return a[i];
    }

    public void set(int i, int n) {
        a[i] = n;
    }

    public void pushback(int n) {
        if (len == capacity) {
            this.resize();
        }
        this.a[len] = n;
        this.len++;
    }

    public int popback() {
        int popped = a[this.len-1];
        this.len--;
        return popped;
    }

    private void resize() {
        this.capacity = 2 * this.capacity;
        int[] resized = new int[this.capacity];
        for (int i = 0; i < len; i++) {
            resized[i] = a[i];
        }
        this.a = resized;
    }

    public int getSize() {
        return this.len;
    }

    public int getCapacity() {
        return this.capacity;
    }
}
