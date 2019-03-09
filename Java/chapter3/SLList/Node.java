package linkedlist;

import java.util.Objects;
/**
 * <p>
 * 連結リストのブロックを実現するクラスです。
 * </p>
 */
public class Node<T>{

    private Node<T> next;
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

    public String toString(){
        String result = "";
        try{
            result = String.format("{%s}",this.data.toString());
        }catch(Exception e){
            e.printStackTrace();
            String errMsgformat = "%s オブジェクトにはtoString() メソッドが実装されていません。";
            System.out.printf(errMsgformat,this.data.getClass());
        }
        return result;
    }

    public int hashCode(){
        return Objects.hash(this.data,this.next);
    }

    public boolean equals(Object o){
        if(Objects.isNull(o)){
            return false;
        }
        if(!o.equals(this.getClass())){
            return false; 
        }
        Node<?> node = (Node)o;
        return node.getData().equals(this.getData()) && node.getNext().equals(this.getNext());
    }
}

