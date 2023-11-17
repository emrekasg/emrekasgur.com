[PostLink] = cpu-architecture
[PostTitle] = CPU Architecture
[Brief] = In this post, I will try to explain the CPU architecture in a simple way.
[Language] = en
[Tag] = computer-science
[Visible] = true

Note: The main reason of writing a blog for me, is to write down what I’ve learned so far while trying to understand a topic. 

### What does a CPU consist of?

An ordinary CPU generally consists of cores, cache memory, buses, clock and FPU, memory management unit, random number generator, ALU, CU and more…

### Arithmetic Logic Unit (ALU)

The ALU is responsible for performing all arithmetic and logical operations within the CPU. This includes basic operations like addition, subtraction, multiplication, division, and bitwise operations (AND, OR, XOR, NOT).

**In a core or integrated with CPU?**: Each core in a CPU has its own ALU.

### Control Unit ( CU )

The CU acts as the conductor of the CPU, directing the operation of the processor. It fetches instructions from memory, decodes them, and executes them by directing the coordinated operations of the ALU, registers, and other components. If a program requires data to be fetched from memory, it's the CU that sends the necessary signals to retrieve this data.

**In a core or CPU?** : Each core in a CPU has its own CU.

### Registers

Registers are small, fast storage locations directly within the CPU. They are used to hold temporary data and instructions that are being executed. When performing a calculation, the ALU uses data from these registers, ensuring rapid execution.

**In a core or CPU?** : Each core contains its own set of registers.

### Cache memory

Cache is a small-sized type of volatile computer memory that provides high-speed data access to a processor and stores frequently used computer programs, applications, and data. There are three levels of cache: L1, L2, and L3, each with varying speed and storage capacity. L1 is the fastest but has the least capacity, while L3 is slower but has more capacity. The efficiency of cache memory significantly impacts the speed of data access and, consequently, the overall performance of a computer program.

**In a core or CPU?** : The situation with cache memory is a bit more complex. Modern CPUs often have a hierarchy of cache memory, including L1, L2, and sometimes L3 caches. The L1 cache is usually core-specific. The L2 cache might be either core-specific or shared between a pair of cores, depending on the CPU's architecture. The L3 cache, if present, is typically larger and slower, and is often shared across all cores in the CPU.

### Buses

Buses are subsystems that transfer data within a computer. They are critical for the CPU to communicate with other parts of the computer. There are three main types:

- **Data Bus**: Transfers data between the CPU, memory, and other hardware components.
- **Address Bus**: Carries the addresses of data (but not the data itself) to and from the CPU.
- **Control Bus**: Transmits command and control signals from the CPU.

**In a core or CPU?** : Buses generally span across the entire CPU and are not limited to a single core. In multicore processors, there are internal buses within the CPU that facilitate communication between cores, as well as with other components like cache memory and the external motherboard.

### Clock

The CPU clock is the oscillator that sets the tempo for the processor. It generates a series of electrical pulses at regular intervals, which are used to synchronize the operations of the CPU's components.

In multicore processors, each core has its own clock. These cores can operate at the same or different clock speeds. Some multicore CPUs feature dynamic adjustment of clock speeds (known as dynamic frequency scaling or CPU throttling) to balance performance and energy efficiency.

**In a core or CPU?** 

Modern CPUs often have a central clock that operates across the entire CPU. However, each core may have the ability to independently adjust its operating frequency based on power management or performance requirements. This feature is known as dynamic frequency scaling or CPU throttling.

## Instruction Set Architecture

The Instruction Set Architecture (ISA) is a fundamental aspect of CPU design, acting as the bridge between software and hardware. It defines the set of instructions that the CPU can execute, the data types, the registers, the memory architecture, and the input/output model of a computer. Essentially, the ISA tells programmers how to write software that the CPU can understand and execute.

## **Instruction** pipelining

CPU pipelining is a technique used in modern processors to improve throughput - the number of instructions that can be executed over a period of time.

### Stages of CPU Pipelining

1. **Fetch**: The CPU retrieves an instruction from memory.
2. **Decode**: The instruction is decoded to understand what action is required.
3. **Execute**: The CPU performs the instruction's action.
4. **Memory Access**: If needed, the CPU accesses the memory for data.
5. **Write Back**: The results of the instruction are written back to the CPU register or memory.

By having each of these stages operate concurrently with different instructions, the CPU can process multiple instructions simultaneously, greatly improving efficiency and speed.

## Branch prediction

Branch prediction is a technique used in CPU design to improve the flow in the instruction pipeline. They contain a dedicated hardware unit called a branch prediction unit (BPU). Many programs use many conditional statements (like **`if-else`** or loops), where the flow of execution depends on certain conditions. When a CPU encounters a branch (a point where the program can follow one of two paths), it must decide which path to follow before the condition is fully evaluated. 

Branch prediction involves predicting which path the program will take before this is known for certain, allowing the CPU to continue executing instructions without waiting. If the prediction is correct, the CPU avoids stalling and saves valuable time. If the prediction is wrong, the CPU must discard the wrongly executed instructions and redirect to the correct path, which can be costly in terms of performance.

Since, I don’t think I need to know more about branch prediction, I won’t research about it 😄

— To be continued about threads