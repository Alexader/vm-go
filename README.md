# vm-go
A stack virtual machine implementation in golang

## Archetechture
### Instruction
Instruction layout is the most important part of the virtual machine design.
In this section, our instruction is as follows:
|   4bit     | 1bit    11bit                  |
|  opcode    | mode | data or imediate number |

In our cases, mode can imediate or referring another data.