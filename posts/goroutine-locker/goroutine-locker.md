[PostLink] = pinning-goroutines-into-specific-cores
[PostTitle] =  Pinning goroutines into specific cores
[Brief] = An experimental package for pinning goroutines to specific CPU cores
[Language] = en
[Tag] = go
[Visible] = true

### **Introduction**

Efficient CPU core and thread management are pivotal in concurrent programming for high performance and scalability. This blog post examines the **`goroutine-locker`** repository, highlighting its utility in scenarios where pinning threads to specific cores is beneficial, akin to strategies employed in high-performance databases like ScyllaDB. Since we cannot 

### **The Shard-Per-Core Concept in ScyllaDB**

ScyllaDB, a leading NoSQL database, utilizes a shard-per-core architecture. Each CPU core is assigned a shard/thread, facilitating lock-free data structure development. This approach enhances CPU cache usage and scalability, as cores operate independently, reducing race conditions and mutex lock dependencies. 

### **Goroutine-Locker: Enhancing Go's Concurrency**

The **`goroutine-locker`** repository adopts a similar principle, binding goroutines to specific CPU cores. This is particularly advantageous for:

1. **Maximizing CPU Cache Efficiency**: Consistent execution on the same core improves cache hit rates, reducing latency from cache misses.
2. **Minimizing Context Switching**: Core-bound threads reduce context switching overhead, enhancing performance in CPU-intensive tasks.

### **How Goroutine-Locker Works: A Step-by-Step Guide**

1. **Initialization**: The process begins with the **`NewCoreManager`** function in **`cpu.go`**, which initializes the **`CoreManager`**. This manager creates a **`Core`** for each CPU core on the system, with each **`Core`** having its task queue and goroutine counter.
2. **Task Assignment**: Tasks are assigned to cores, with each task being queued up in the respective **`Core`**'s task queue.
3. **Goroutine Creation and Binding**: For each task, a goroutine is created. The **`lock_thread`** function from **`cpuaffinity.c`** is then called to bind this goroutine to the designated CPU core. This binding is crucial for ensuring that the goroutine consistently runs on its assigned core.
4. **Executing Tasks**: The goroutines execute their assigned tasks while being confined to their specific cores, optimizing CPU cache usage and reducing context switching.

### **Benchmark Analysis**

### With goroutine-locker

- **Total Duration**: 2.741806793 seconds
- **Core Utilization**:
    - CPU 0: 7 kernel threads, 250,000 goroutines
    - CPU 1: 5 kernel threads, 250,000 goroutines
    - CPU 2: 6 kernel threads, 250,000 goroutines
    - CPU 3: 0 kernel threads, 0 goroutines
- **Total Kernel Thread Count**: 18

### Without goroutine-locker

- **Total Duration**: 246.083667 milliseconds
- **Total Kernel Thread Count**: 6

### **Analysis and Implications**

The benchmarks reveal a significant difference in execution time and kernel thread count when using the **`goroutine-locker`** package. While the total duration is longer with the package, it demonstrates a more balanced and dedicated use of CPU cores. This can be particularly beneficial in scenarios where consistent and isolated CPU performance is critical.

However, it's important to note that the increased duration and higher kernel thread count might not always be advantageous, especially for applications that are not bottlenecked by CPU cache misses or context switching.

### **Go's Scheduler**

Go's runtime includes a sophisticated scheduler (M:N scheduler) designed to efficiently manage goroutines. This scheduler dynamically distributes goroutines across multiple OS threads (M) in a way that is typically abstracted away from the developer. The scheduler's primary goals are to keep these threads busy and to balance the load across them, while also considering factors like I/O blocking and CPU usage. 

### **Interaction with Go's Scheduler**

We interact with Go's runtime scheduler in a way that is not typical for standard Go applications. Understanding this interaction is crucial to evaluate its compatibility and potential conflicts with Go's scheduler.

By setting CPU affinity at the thread level (using CGo to interact with OS-specific thread management), the package further constrains the scheduler. It effectively overrides the scheduler's ability to move these threads across different CPU cores, pinning them to specific cores.

### **Potential Conflicts and Compatibility Issues**

1. **Reduced Flexibility**: The Go scheduler is designed to optimize performance by dynamically managing goroutines. By locking goroutines to specific threads and cores, **`goroutine-locker`** reduces this flexibility, which can lead to suboptimal utilization of system resources in certain scenarios.
2. **Scheduler Overhead**: The Go scheduler might incur additional overhead trying to manage these locked threads, potentially leading to performance degradation, especially in systems with a large number of goroutines and cores.
3. **Context Switching**: While the package aims to reduce context switching at the CPU level, it may inadvertently increase context switching at the OS thread level, as the Go scheduler attempts to work around the locked threads.
4. **Work-stealing**: Locking goroutines to specific cores means they cannot be moved or stolen by other threads. This static binding goes against the dynamic scheduling nature of Go's scheduler. If certain cores are heavily loaded while others are idle, the work-stealing scheduler cannot redistribute the workload. This can lead to inefficient CPU utilization, where some cores are overburdened, and others are underutilized.

### **Conclusion**

This repository is partially compatible with Go's scheduler in the sense that it can function alongside it, but it introduces a paradigm that is quite different from the scheduler's intended use. This approach can be beneficial in specific use cases where dedicated CPU core utilization is paramount. However, it can also lead to conflicts with the Go scheduler's natural behavior, potentially affecting the overall efficiency and performance of the application. Developers should carefully consider these aspects when deciding to use such an approach in their Go applications.

Source code: https://github.com/emrekasg/goroutine-locker