# 「みんなのデータ構造」の演習レポジトリ

このレポジトリは以下の書籍の実装コードを個人的な演習を兼ねて保管するためのものです。　
- [みんなのデータ構造](https://www.lambdanote.com/collections/frontpage/products/opendatastructures) 

- ディレクトリは各実装言語ごとに分割され、更にChapterごとにディレクトリが分割されています。
- 記号のNotationはなるべく書籍と一致させるようにしてあります。

## 用語解説

### 1.Queue Deque Stack interface
以下の操作を持つデータ構造のこと。
- add(x)　：データxをQueueに追加する
- remove() ：直前に追加されたyをQueueから削除し、yを返す。

LIFO Queue => Stack  
FIFO Queue => Queue

Stackの場合、add(x),remove()をそれぞれpush(x),pop()と呼んで区別する。

### 2.List interface

以下の操作を持つデータ構造のこと

- size(): リストの長さnを返す。
- get(i): x_iを返す。
- set(i,x): x_iの値をxにする。
- add(i,x): x_iをi番目として追加し、その後x_i ... x_nを後ろにずらす。
- remove(i): x_iを削除し、x_(i+1) ... x_(n-1)を前にずらす。

### 3.USet interface
以下の操作をもつデータ構造である。

- size(): 要素数nを返す。
- add(): 要素xが集合になければ集合に追加
- remove(): 集合からxを追加する。
- find(): 集合にxが入っていればそれをみつける。なければnullを返す。

USetは順序がない

### 4.SSet interface

- USetに順序関係が入ったもの


