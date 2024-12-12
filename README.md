# RBMV: ReBalance Based on Memory Variance

RBMV presents a workload re-balancing strategy primarily focused on memory variance. For a detailed understanding of its 
functionality, please refer to the documentation [here](rbmv.en.md), or you may directly examine the implementation 
starting from the `rebalance.go` file.

## Instructions

This repository demonstrates the concept of horizontal shrinkage, which is the most intricate scheme, in a 
straightforward manner. Before designing your own strategy, please consider the following guidelines:

- Outer Control: Adjust parameters of the `ReBalance` function, such as `Node`, `ReBalanceRes`, or introduce additional 
global parameters as needed.
- Expansion Strategy: Customize the `hasBetterMove` function with your own expansion sub-strategies, such as 
Less-Move-Mem, Less-Leader-Count, and others.
- Compatibility with Horizontal Expansion: Modify the `initMem` parameter to alter the process based on specific flags.

By adhering to these guidelines, you can effectively implement and customize your workload re-balancing strategy using RBMV.

## License

This repository is licensed under the [Apache-2.0 License](LICENSE).
