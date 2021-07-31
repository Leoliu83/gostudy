#### 字符串模式匹配算法
在一个长字符串中寻找一个指定字符串所在的位置，这种子串的定位操作通常称为字符串的模式匹配。
##### 朴素模式匹配算法（Brute Force）
- 流程如下：
``` flow
st=>start: 开始
ed=>end: 结束
op0=>operation: i = -1
op1=>operation: i += 1
op2=>operation: 从主字符串第 i 位开始，
其每一位与子串的每一位
进行比较
cond1=>condition: 是否完全匹配
op3=>operation: i += 1

st->op0->op1->op2->cond1
cond1(yes)->ed
cond1(no)->op1
```

**简单的说**: 就是对主字符串的每一个字符做为子字符串的开头，与要匹配的字符串进行比对，对主字符串做大外循环，对子字符串做内循环，直到匹配成功为止。

##### KMP模式匹配算法
KMP模式匹配算法的思路是：当某一个字符与主串不匹配时，我们应该知道j指针要移动到哪。
核心：
1. 主串下标只增不减
由于主串下标只增不减，就需要我们正确的计算与i值比对的j值

主串(S)：
| $下标(i)$ |   0   |   1   |   2   |   3   |   4   |   5   |   6   |   7   |   8   |   9   |  10   |  11   |
| :-------: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
|    值     |   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |

子串(T)：
| $下标(j)$ |   0   |   1   |   2   |   3   |   4   |   5   |
| :-------: | :---: | :---: | :---: | :---: | :---: | :---: |
|    值     |   a   |   b   |   c   |   a   |   b   |   x   |


$S_0 = T_0 (i=0,j=0)$ 
$S_1 = T_1 (i=1,j=1)$ 
$S_2 = T_2 (i=2,j=2)$
$S_3 = T_3 (i=3,j=3)$ 
$S_2 = T_2 (i=4,j=4)$ 
$S_5 \neq T_5 (i=5,j=5)$ ← 这里开始注意了，之后$i$值不减，不回跳，$j$值做**适应性**变化
$S_5 = T_2 (i=5,j=2)$  ← 所谓的$KMP$模式匹配就是为了找到$(j=2)$这个值

**接下来就分两种情况说明，在发现匹配不上的时候，为什么比对$(j=2)$而不是其他值**

#### 子串中有相同项
##### 比对流程
- 第1步: 比对$S_0,T_0$

**主串（$S_n$）,子串（$T_n$）** 
|   👇   |       |       |       |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |    b     |    x     |
|   👇   |       |       |       |       |       |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   a   |   b   |   x   |
匹配✔

- 第2步: 比对$S_1,T_1$

**主串（$S_n$）,子串（$T_n$）** 
|       |   👇   |       |       |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |    b     |    x     |
|       |   👇   |       |       |       |       |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   a   |   b   |   x   |
匹配✔

- 第3步: 比对$S_2,T_2$

**主串（$S_n$）,子串（$T_n$）** 
|       |       |   👇   |       |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |    b     |    x     |
|       |       |   👇   |       |       |       |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   a   |   b   |   x   |
匹配✔

- 第4步: 比对$S_3,T_3$

**主串（$S_n$）,子串（$T_n$）** 
|       |       |       |   👇   |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |    b     |    x     |
|       |       |       |   👇   |       |       |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   a   |   b   |   x   |
匹配✔

- 第5步: 比对$S_4,T_4$

**主串（$S_n$）,子串（$T_n$）** 
|       |       |       |       |   👇   |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |    b     |    x     |
|       |       |       |       |   👇   |       |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   a   |   b   |   x   |
匹配 ✔

- 第6步：比对$S_5,T_5$

**主串（$S_n$）,子串（$T_n$）** 
|       |       |       |       |       |   👇   |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |    b     |    x     |
|       |       |       |       |       |   👇   |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   a   |   b   |   x   |
不匹配 ✖

- 第7步：主串游标保持，子串游标回溯

**主串（$S_n$）,子串（$T_n$）** 
|       |       |       |       |       |   👇   |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |    b     |    x     |
|       |       |   👇   |       |       |       |       |       |       |       |          |          |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   a   |   b   |   x   |
|   ●   |   ★   |       |   ●   |   ★   |       |

**INPORTANT:**
这里为什么回溯到$T_2$
因为: $T_0=T_3,T_1=T_4$
其次: $T_3=S_3,T_4=S_4$
所以: $T_0=S_3,T_1=S_4$
因此可以跳过$T_0,T_1$与主字符串的匹配，直接从$T_2$开始与主字符串的$S_5$进行比对。

#### 子串中无相同项
##### 比对流程
- 第1步: 比对$S_0,T_0$

**主串（$S_n$）,子串（$T_n$）** 
|   👇   |       |       |       |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |    b     |    x     |
|   👇   |       |       |       |       |       |       |       |       |       |          |          |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   d   |   e   |   f   |
匹配✔

- 第2步: 比对$S_1,T_1$

**主串（$S_n$）,子串（$T_n$）** 
|       |   👇   |       |       |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |    b     |    x     |
|       |   👇   |       |       |       |       |       |       |       |       |          |          |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   d   |   e   |   f   |
匹配✔

- 第3步: 比对$S_2,T_2$

**主串（$S_n$）,子串（$T_n$）** 
|       |       |   👇   |       |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |    b     |    x     |
|       |       |   👇   |       |       |       |       |       |       |       |          |          |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   d   |   e   |   f   |
匹配✔

- 第4步: 比对$S_3,T_3$

**主串（$S_n$）,子串（$T_n$）** 
|       |       |       |   👇   |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |    b     |    x     |
|       |       |       |   👇   |       |       |       |       |       |       |          |          |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   d   |   e   |   f   |
不匹配✖

- 第5步：主串游标保持，子串游标回溯

**主串（$S_n$）,子串（$T_n$）** 
|       |       |       |   👇   |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |    b     |    x     |
|   👇   |       |       |       |       |       |       |       |       |       |          |          |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   d   |   e   |   f   |

**INPORTANT:**
这里为什么回溯到$T_0$
因为: 子串中没有相同的项，所以子串回溯到0，从$T_0$开始

**求next数组**
什么是next数组，next数组就是，当子串的j位与主串不匹配时，在主串i不回溯的前提下，j回溯到**什么位置**继续与主串进行比较。
当$j=n$时（$n$已知），与主串匹配不上，这时，$j$应该回溯到**什么位置**继续与主串进行比对，这个位置就是 next[j]，并且，这个位置与主串完全无关。

**主串（$S_n$）,子串（$T_n$）**
我们不需要知道主串里面是什么 
|       |       |       |       |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   ?   |   ?   |   ?   |   ?   |   ?   |   ?   |   ?   |   ?   |   ?   |   ?   |    ?     |    ?     |

##### 我们看子串
##### 假设，子串中无重复项时的情况
子字符串是：

**当$j=0$时就匹配不上时**，应该回溯到什么位置呢？
|   ✖   |       |       |       |       |       |
| :---: | :---: | :---: | :---: | :---: | :---: |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   d   |   e   |   f   |
答案是：不需要回溯,因为$T_0$已经是第一个位置了

**当$j<1$都匹配上，$T_1$匹配不上时**，应该回溯到什么位置呢？
|   ✔   |   ✖   |       |       |       |       |
| :---: | :---: | :---: | :---: | :---: | :---: |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   d   |   e   |   f   |
答案是：回溯到$j=0$
因为$T_0$是第一个位置，在子串中没有可比较的项。

**当$j<2$都匹配上，$T_2$匹配不上时**
|   ✔   |   ✔   |   ✖   |       |       |       |
| :---: | :---: | :---: | :---: | :---: | :---: |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   d   |   e   |   f   |
且$j=2$时，即$T_2$与主串匹配不上时，应该回溯到什么位置呢？
答案是：回溯到$j=0$
- **首先**
因为$T_1 \neq T_0$，且$T_1=S_1$，所以 $T_0 \neq S_1$一定成立，所以不需要比对
- **其次**
由于$T_0 \neq T_2$，$T_2 \neq S_2$ 无法推断$T_0$和$S_2$的关系，因此仍然需要比对$T_0$和$S_2$。
- **如果**
子串中$T_1=T_0$，由于$T_1=S_1$，那么$T_0=S_1$一定成立，因此不用比对，只需比对$S_2和T_1$，在这种情况下，就应该回溯到$j=1$，也就是$T_1$
**如下所示：**

|       |   ✔   |   👇   |       |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   a   |   c   |   a   |   b   |   c   |   a   |   b   |   c   |   a   |    b     |    x     |
|       |   ✔   |   👇   |       |       |       |       |       |       |       |          |
|       | $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|       |   a   |   a   |   x   |   x   |   x   |   x   |

**当$j<3$都匹配上，$T_3$匹配不上时** 同上

##### 假设，子串中有重复项时的情况
子字符串是：
$j=0,j<1,j<2$的情况上面都讨论过了，下面看复杂一点的情况


**当$j<3$都匹配上，$T_3$匹配不上时**，应该回溯到什么位置呢？
|   ✔   |   ✔   |   ✔   |   ✖   |       |       |
| :---: | :---: | :---: | :---: | :---: | :---: |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   a   |   c   |   a   |   x   |
答案是：回溯到$j=1$，
由$T_2=T_0$，$T_2=S_2$，可知$T_0=S_2$一定成立，因此不需要比对。
由$T_3 \neq T_1$，$T_3 \neq S_3$,无法推断 **$\underline{T_1}$和$\underline{S_3}$** 的关系，因此仍然需要比对。

**如下所示**
|       |       |   ✔   |   👇   |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   a   |   b   |   a   |   c   |   a   |   x   |   c   |   a   |    b     |    x     |
|       |       |   ✔   |   👇   |       |       |       |       |       |       |          |          |
|   →   |   →   | $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   →   |   →   |   a   |   b   |   a   |   c   |   a   |   x   |

**这里有个地方会有疑问，就是为什么不直接跳到$T_0$和$S_3$比对呢？**

**看下面的例子**
|       |       |       |   👇   |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   a   |   b   |   a   |   c   |   a   |   x   |   c   |   a   |    b     |    x     |
|       |       |       |   👇   |       |       |       |       |       |       |          |          |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   a   |   c   |   a   |   x   |

**按照上面的逻辑，回溯到$T_1$，也就是$T_1$和$S_3$比较，可以看到刚好就匹配上了**
|       |       |   ✔   |   👇   |   ✔   |   ✔   |   ✔   |   ✔   |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   a   |   b   |   a   |   c   |   a   |   x   |   c   |   a   |    b     |    x     |
|       |       |   ✔   |   👇   |   ✔   |   ✔   |   ✔   |   ✔   |       |       |          |          |
|   →   |   →   | $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   →   |   →   |   a   |   b   |   a   |   c   |   a   |   x   |

**如果直接跳到$T_0$和$S_3$比对**
|       |       |       |   ✖   |       |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   a   |   b   |   a   |   c   |   a   |   x   |   c   |   a   |    b     |    x     |
|       |       |       |   ✖   |       |       |       |       |       |       |          |          |
|   →   |   →   |   →   | $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   →   |   →   |   →   |   a   |   b   |   a   |   c   |   a   |   x   |

可以看到匹配不上，错过了能够匹配上的字符串。

所以，结合上面的例子，可以得出，$T_0$和$T_2$相等，那么如果$T_3$不匹配时，$T_0$就需要与原来的$T_2$所在位置保持对齐

**再看下面的情况**
|       |       |       |       |   👇   |       |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   a   |   b   |   a   |   b   |   c   |   x   |   c   |   a   |    b     |    x     |
|       |       |       |       |   👇   |       |       |       |       |       |          |          |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   a   |   b   |   c   |   x   |
|   →   |   →   | $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   →   |   →   |   a   |   b   |   a   |   b   |   c   |   x   |

$T_4$匹配不上，而且$T_0=T_2$，$T_1=T_3$，那么$T_0,T_1$就要与$T_2=T_3$位置对齐。也就是说T子串要回溯到$T_2$位置。
**注意：** 回溯一定发生在主串与子串第一次不匹配时。

**再看一个子串没有相同项的情况:**
|       |       |       |       |       |   👇   |       |       |       |       |          |          |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :------: | :------: |
| $S_0$ | $S_1$ | $S_2$ | $S_3$ | $S_4$ | $S_5$ | $S_6$ | $S_7$ | $S_8$ | $S_9$ | $S_{10}$ | $S_{11}$ |
|   a   |   b   |   c   |   d   |   e   |   b   |   c   |   x   |   c   |   a   |    b     |    x     |
|       |       |       |       |       |   👇   |       |       |       |       |          |          |
| $T_0$ | $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ |
|   a   |   b   |   c   |   d   |   e   |   f   |

当$T_5$匹配不上时，因为$T_0 \neq T_1,T_0 \neq T_2,T_0 \neq T_3,T_0 \neq T_4$，由因为$T_1=S_1,T_2=S_2,T_3=T_3,T_4=S_4$,所以$T_0$一定不等于$T_1,T_2,T_3,T_4$。所以直接$T_0$与$S_5$比较是否相等即可。

**书上的公式如下**
$$ next(j)=\left\{
\begin{aligned}
&0,当j=1时 \\
&Max\{k|1<k<j, 且'P_1 \dotsm P_{k-1}'='P_{j-(k-1)} \dotsm P_{j-1}'\} \\
&1, 其他情况
\end{aligned}
\right.
$$

$j=0,j=1$都没有问题 **（书中的下标从1开始）**
中间的$Max\{k|1<k<j, 且'P_1 \dotsm P_{k-1}'='P_{j-k+1} \dotsm P_{j-1}'\}$如何解释：
简单来说就是要找到子串中最长相同字符串

假设有如下子字符串：
|   1   |   2   |   3   |   4   |   5   |   6   |   7   |   8   |   9   |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ | $T_6$ | $T_7$ | $T_8$ | $T_9$ |
|   a   |   b   |   a   |   b   |   a   |   a   |   a   |   b   |   a   |


**当子串$j=6$位置与主串不匹配时，找$\{T_1 \dotsm T_5\}$子串中的相同字符串**
只有如下一种：
|   1   |   2   |   3   |   4   |   5   |   6   |   7   |   8   |   9   |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ | $T_6$ | $T_7$ | $T_8$ | $T_9$ |
|   a   |   b   |   a   |   b   |   a   |   a   |   a   |   b   |   a   |
|   ★   |   ★   |   ★   |       |       |       |       |       |       |
|       |       |   ●   |   ●   |   ●   |       |       |       |       |

最长相同字符串是$'T_1T_2T_3'='T_3T_4T_5'$

此时 $k-1=3,j-1=6$，得$k=4,j=7$,因此回溯位置就是$k=4$，也就是$T_4$

**为甚么下面这个不是**
|   1   |   2   |   3   |   4   |   5   |   6   |   7   |   8   |   9   |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ | $T_6$ | $T_7$ | $T_8$ | $T_9$ |
|   a   |   b   |   a   |   b   |   a   |   a   |   a   |   b   |   a   |
|   ★   |   ★   |       |       |       |       |       |       |       |
|       |       |   ●   |   ●   |       |       |       |       |       |

因为最长字符串必须是：
$\{'P_1 \dotsm P_{k-1}'='P_{j-k+1} \dotsm \color{red}{P_{j-1}}'\}$
匹配字符串必须匹配到$P_{j-1}$,在这里 $'T_1T_2'='T_3\color{red}{T_4}'$中$\color{red}{4}$不满足$\color{red}{j-1}$这个条件。

**为什么不是下面这种**

|   1   |   2   |   3   |   4   |   5   |   6   |   7   |   8   |   9   |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ | $T_6$ | $T_7$ | $T_8$ | $T_9$ |
|   a   |   b   |   a   |   b   |   a   |   a   |   a   |   b   |   a   |
|       |   ★   |   ★   |       |       |       |       |       |       |
|       |       |       |   ●   |   ●   |       |       |       |       |
因为最长字符串必须是：
$\{'$$\color{red}{P_1}$$ \dotsm P_{k-1}'='P_{j-k+1} \dotsm P_{j-1}'\}$
匹配字符串必须从$\color{red}{P_{1}}$开始匹配,在这里 $'\color{red}{T_2}$$T_3'='T_4T_5'$中$2$不满足从$\color{red}{1}$开始这个条件。

**再看看当子串$j=7$位置与主串不匹配时，找$\{T_1 \dotsm T_6\}$子串中的相同字符串**
只有如下一种：
|   1   |   2   |   3   |   4   |   5   |   6   |   7   |   8   |   9   |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ | $T_6$ | $T_7$ | $T_8$ | $T_9$ |
|   a   |   b   |   a   |   b   |   a   |   a   |   a   |   b   |   a   |
|   ★   |       |       |       |       |       |       |       |       |
|       |       |       |       |       |   ●   |       |       |       |

**再看看当子串$j=8$位置与主串不匹配时，找$\{T_1 \dotsm T_7\}$子串中的相同字符串**
只有如下一种：
|   1   |   2   |   3   |   4   |   5   |   6   |   7   |   8   |   9   |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ | $T_6$ | $T_7$ | $T_8$ | $T_9$ |
|   a   |   b   |   a   |   b   |   a   |   a   |   a   |   b   |   a   |
|   ★   |       |       |       |       |       |       |       |       |
|       |       |       |       |       |       |   ●   |       |       |

**再看看当子串$j=9$位置与主串不匹配时，找$\{T_1 \dotsm T_8\}$子串中的相同字符串**
只有如下一种：
|   1   |   2   |   3   |   4   |   5   |   6   |   7   |   8   |   9   |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| $T_1$ | $T_2$ | $T_3$ | $T_4$ | $T_5$ | $T_6$ | $T_7$ | $T_8$ | $T_9$ |
|   a   |   b   |   a   |   b   |   a   |   a   |   a   |   b   |   a   |
|   ★   |   ★   |       |       |       |       |       |       |       |
|       |       |       |       |       |       |   ●   |   ●   |       |
