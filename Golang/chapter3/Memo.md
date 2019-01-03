# chapter3 要約

## 実行時間のまとめ

|AlgoName|push()|pop()|add(x)|remove()|get(i)|set(i,x)|
|--------|------|-----|------|--------|------|--------|
|SLList|O(1)|O(1)|O(1)|O(1)|-|-|

|AlgoName|push()|pop()|add(i,x)|remove(i)|get(i)|set(i,x)|
|--------|------|-----|------|--------|------|--------|
|DLList|-|-|O(min{i,n-i})|O(min{i,n-i})|O(min{i,n-i})|O(min{i,n-i})|




## SLList(singly-linked-list)の特徴

- 単方向連結リストである。接続にはポインタを用いる。
- stackとqueueインターフェースを実装する。

## DLList(doubly-linked-list)の特徴

- Listインターフェースを実装する。
- 時間がかかる操作はget_nodeであるから、nodeの参照を別の方法で保持するようにするのが好ましい

----------
SEListはスキップする(2018/01/03)

