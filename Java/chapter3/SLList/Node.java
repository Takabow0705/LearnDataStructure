package linkedlist;

/**
 * <p>
 * 連結リストのブロックを実現するクラスです。
 * </p>
 */
public class Node<T>{

    private Node next;
    private T data;

    public Node(T data){
        this.data = data;
        this.next = null;
    }

    /**
     * <p>
     * ノードが保持するデータを返します。
     * </p>
     * 
     * @return data type<T>
     * 
    */
    public T getData(){
        return data;
    }

    /**
     * <p>
     * 次のノードを返す。
     * </p>
     * 
     * @param Node<T> next
     */
    public Node<T> getNext(){
        return this.next;
    }
    /**
     * <p>
     * 次のノードをセットするためのsetter
     * </p>
     * 
     * @param  Node<T> next
     * 
     */
    public void setNext(Node<T> next){
        this.next = next;
    }
}

