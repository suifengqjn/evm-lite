
#evm概述
--




#### evm

* [1.evm语言虚拟机](#1)
* [2.虚拟机设计组成](#2)
* [3.evm设计目标](#3) 
* [4.evm优缺点](#4)
* [5.IELE设计原理](#5)
* [6.evm和其他语言虚拟机区别](#6)


<h3 id="1">1.evm语言虚拟机</h3>

* evm本质上就是一个简单的语言虚拟机，常见的语言虚拟机有jvm,lvm。语言虚拟机主要实现了隔离技术，就是指在底层实现了环境隔离，它屏蔽了与具体操作系统平台相关的信息，使得程序只在虚拟机上运行的目标代码（字节码），就可以在多种平台上不加修改地运行。
	
* EVM使用了256比特长度的机器码，是一种基于堆栈的虚拟机，用于执行智能合约，并使用了以太坊账户模型（Account Model）来进行价值传输。
* EVM是图灵完备的，由于以太坊系统中引入了Gas的概念，所以原则上在EVM中可执行的计算总量受Gas总量限制。
* EVM是一个隔离的环境，在EVM内部运行的代码不能跟外部有任何联系，比如屏蔽掉了直接调用系统api。


<h3 id="2">2.虚拟机设计组成 </h3>

* 编译器
	* 多语言支持
	* 新设计一门语言
* 虚拟机
	* 指令	opcode算数相关
	* 指令	INSTRUCTION

* 执行机
	* 栈
	* 内存
	* statedb
	* 执行过程
	* 解释器

<h3 id="3">3.evm设计目标</h3>

* 简单：操作码尽可能的少并且低级，数据类型尽可能少，虚拟机的结构尽可能少；
* 结果明确：在VM规范语句中，没有任何可能产生歧义的空间，结果应该是完全确定的。此外，计算步骤应该是精确的，以便可以测量Gas的消耗量；
* 节约空间：EVM组件应尽可能紧凑；
* 预期应用应具备专业化能力：在VM上构建的应用应能处理20字节的地址，以及32位的自定义加密值，拥有用于自定义加密的模数运算、读取区块和交易数据与状态交互等能力；
* 简单安全：为了让VM不被利用，应该能够容易地建立一套Gas消耗成本模型的操作；
* 优化友好：应该易于优化，以便即时编译（JIT）和VM的加速版本能够构建出来。



<h3 id='4'> 4.evm优缺点</h3>

* 优点：  
	* EVM中要么执行智能合约的所有代码，要么一点也不执行，完全不可能只执行其中部分代码，无论是燃料不足还是无效指令，最终状态恢复到交易执行前的checkpoint，这样可以很大程度规避作恶的可能。程序总是从头开始运行，无法跳过Solidity ABI引导代码。  
	* evm与网络，文件系统，进程等隔离  
	* 对外是完全隔离的，甚至不同合约之间也只有有限的访问权限  
	
* 缺点：
	* 机器码长度为256位
	* 缺少标准库
	* 难以调试和测试
	* 不支持浮点数
	* 不可修改的代码
		
<h3 id='5'> 5.IELE设计原理</h3>

* IELE是LLVM的一个变种，专门用于在区块链上执行智能合约
* 与基于栈的EVM不同，IELE是基于寄存器的机器，就像LLVM。它支持无限的寄存器以及无界整数。
* 合约可以通过ABI（应用程序二进制接口）相互作用
* 为所有语言提供统一的gas模型。
* 让编写安全的智能合约更容易。	它避免了使用有界堆栈，而不必担心堆栈或内存泄漏，使智能合约的规范和验证更容易。


<h3 id='6'> 6.evm和其他语言虚拟机区别 </h3>

* evm和jvm，lvm最大的区别就是缺少任务调度和上下文切换，因为这一部分只涉及到性能并不牵扯到执行安全，所以可以将其抽取到外面，比如以太坊其实是用编程语言自带协程来完成即可，而整个虚拟机可以类似理解成一个请求的上下文，这也说明了evm虚拟机其实只需要保证合约执行安全前提下，越精简越好。
* 相比较嵌入式设备上的自建安全语言虚拟机，区块链虚拟机另一大特点是开源，无法通过代码扰乱，重构内存模型（比如地址不连续的栈）等技术实现其安全
* 共同点，相比较lvm虚拟机都可以实现jit加速执行效率，本身分布式架构，这种效率提升又显得并不是很重要


###### 参考链接
**https://github.com/ethereumbook/ethereumbook/blob/develop/evm.asciidoc**
**https://blog.csdn.net/huhaoxuan2010/article/details/80130942**
**http://www.freebuf.com/column/167227.html**  
**https://www.jianshu.com/p/188f6eeb85e3**  
***https://www.jianshu.com/p/e45614c25b62?utm_campaign=maleskine&utm_content=note&utm_medium=seo_notes&utm_source=recommendation***





