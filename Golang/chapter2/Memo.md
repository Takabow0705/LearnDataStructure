# chapter2 要約

## 実行時間のまとめ

|AlgoName|get(i),set(i,x)|add(i,x),remove(i)|
|--------|---------------|------------------|
|ArrayStack|O(1)|O(n-i)|
|ArrayQueue|O(1)|O(min{i,n-i})|
|DualArrayStack|O(1)|O(min{i,n-i})|
|DualArrayQueue|O(1)|O(n-i)|

## ArrayStackに関するメモ

- resize()の平均実行時間はadd,removeを合計ｍ回行うとき、O(m)となる。

## ArrayQueueに関するメモ

### ArrayQueueの操作手順

1. 配列aに関して、削除対象を指すインデックスをj,キューの要素数をnとする。

2. j,n = 0とする。

3. (1)要素の追加：　a[n+j]に要素を追加。その後nを1増やす。  
(2)要素の削除：　a[j]の要素を削除。その後jを1つ増やし、nを1つ減らす。

- これらの操作を剰余演算を応用して行う。

----------
RootishArrayStackは発展的なので、取り敢えず飛ばす（2019/01/02）
