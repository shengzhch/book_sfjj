## 二叉查找树

## 平衡二叉查找树（avl树 红黑树）

### B树事实上是一种平衡的多叉查找树，也就是说最多可以开m个叉（m>=2）

![B树](https://images2017.cnblogs.com/blog/758447/201801/758447-20180126165126756-1850778539.png)

### &emsp; &emsp; &emsp; 性质：

  * 每个节点至多可以拥有m棵子树。
  * 根节点，只有至少有2个节点（要么极端情况，就是一棵树就一个根节点，单细胞生物，即是根，也是叶，也是树)。
  * 非根非叶的节点至少有的Ceil(m/2)个子树(Ceil表示向上取整，图中5阶B树，每个节点至少有3个子树，也就是至少有3个叉)。
  * 非叶节点中的信息包括[n,A0,K1,A1,K2,A2,…,Kn,An]，，其中n表示该节点中保存的关键字个数，K为关键字且Ki<Ki+1，A为指向子树根节点的指针。
  * 除根结点之外的结点的关键字的个数n必须满足： [ceil(m / 2)-1]<= n <= m-1（叶子结点也必须满足此条关于关键字数的性质，根结点除外）。
  * 从根到叶子的每一条路径都有相同的长度，也就是说，叶子节在相同的层，叶子节点除了包含了关键字和关键字记录的指针外也有指向其子节点的指针只不过其指针地址都为null。


### 



### &emsp; &emsp; &emsp; 特点：

<ol>
<li>关键字集合分布在整颗树中。</li>
<li>任何一个关键字出现且只出现在一个节点中。</li>
<li>搜索有可能在非叶子节点结束。</li>
<li>其搜索性能等价于在关键字集合内做一次二分查找。</li>
<li>B树在插入删除新的数据记录会破坏B-Tree的性质，因为在插入删除时，需要对树进行一个分裂、合并、转移等操作以保持B-Tree性质。</li>
</ol>


相邻关键字:节点中左子树中最大的关键字以及右子树中最小的关键字