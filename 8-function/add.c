#include "runtime.h"

/**
 不管是C或是汇编实现的函数，其函数名都是以·开头的。还有，C文件中需要包含runtime.h头文件。
 这个原因在该文件中有说明： Go用了特殊寄存器来存放像全局的struct G和struct M。包含这个
 头文件可以让所有链接到Go的C文件都知道这一点，这样编译器可以避免使用这些特定的寄存器作其
 它用途。

 让我们仔细看一下这个C实现的函数。可以看到函数的返回值为空，而参数多了一个，第三个参数实
 际上被作为了返回值使用。其中FLUSH是在pkg/runtime/runtime.h中定义为USED(x)，这个定义是
 Go的C编译器自带的primitive，作用是抑制编译器优化掉对*x的赋值的。
*/
void ·add(uint64 a, uint64 b, uint64 ret) {
    ret = a + b;
    FLUSH(&ret);
}
