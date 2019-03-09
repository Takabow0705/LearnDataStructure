package linkedlist;

import java.util.Objects;

public class LinkedList<T>{
    private Node<T> head;
    private Node<T> tail;
    private int n;

    public LinkedList(){
        this.head = null;
        this.tail = null;
        this.n = 0;
    }
    /**
     * <p>
     * データを連結リストに追加します。
     * その後、追加した要素を返却します。
     * </p>
     * 
     * @param T x 
     * @return T x
     */
    public T push(T x){
        Node<T> u = new Node<T>(x);
        u.setNext(this.head);
        this.setHead(u);

        if(n == 0){
            this.setTail(u);
        }
            this.increment();
        }
        return x;
    }

    /**
     * <p>
     * 先頭のデータを返したあと削除を行う。
     * 削除対象データが存在しない場合はnullを返す。
     * </p>
     * 
     * @return T x
     */
    public T pop(){
        if(this.n == 0){
            return null;
        }

        T x = this.head.getData();
        Node<T> u = this.head;
        this.setHead(head.getNext());
        this.decrement();

        if(this.n == 0){
            this.tail = null;
        }
        return x;
    }

    public Node<T> getHead(){
        return this.head;
    }

    public Node<T> getTail(){
        return this.tail;
    }
     /**
      * <p>
      *headに新しいノードをセットする
      *</p>
      *
      *@param Node head
      */
    private void setHead(Node<T> head){
        this.head = head;
    }

    /**
     * <p>
     * tailにノードをセットする。
     * </p>
     * 
     * @param Node<T> tail
     */
    private void setTail(Node<T> tail){
        this.tail = tail;
    }

    /**
     * <p>
     * データ数nを1つ増やす。
     * </p>
     * 
     */
    private void increment(){
        this.n += 1;
    }

    /**
     * <p>
     * データ数nを1つ減らす。
     * </p>
     * 
     */
    private void decrement(){
        this.n -= 1;
    }

    public String toString(){
        String resultForm = "{\n\r %s \n\r}";
        String allDataString = "";
        Node<T> node = this.head;

        for(int i = this.n; i > -1; i--){
            String format = "{data[%d] : %s}\r\n";
            String nodeString = String.format(format,this.n - i,node.toString());
            allDataString += nodeString;
            node = node.getNext();
        }

        return String.format(resultForm, allDataString);
    }

    public int hashCode(){
        return Objects.hash(this.head,this.n,this.tail);
    }

    public boolean equals(Object o){
        if(Objects.isNull(o)){
            return false;
        }
        if(!o.equals(this.getClass())){
            return false; 
        }
        LinkedList<?> list = (LinkedList)o;

        return list.getHead().equals(this.head) && list.getTail().equals(this.tail) && list.n == this.n;
    }
}