import linkedlist.LinkedList;

public class Main{
    public static void main(String[] args) {
        LinkedList<String> list = new LinkedList<>();
        list.push("hoge");
        list.push("huga");
        list.push("piyo");
        list.push("piyo");
        list.push("piyo");
        list.push("piyo");
        System.out.println(list.toString());
        list.pop();
        System.out.println(list.toString());        
    }
}