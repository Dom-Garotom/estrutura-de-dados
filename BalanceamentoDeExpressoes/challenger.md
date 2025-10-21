## Balanceamento de expressões 

Escreva um programa que verifique se uma expressão matemática com parênteses, colchetes e chaves está balanceada. Uma expressão está balanceada se:   

- Todo parêntese, colchete ou chave aberto possui um fechamento correspondente.
- O fechamento ocorre na ordem correta.   

Por exemplo:

- Expressões balanceadas: (a + b), {[a * (b + c)]}, a + {b - [c * d]}
- Expressões não balanceadas: a + (b, {[a * b], (a + b)}, a + {b - [c * d}

Requisitos:   

- Utilize uma pilha para implementar a solução.
- A entrada será uma string com a expressão matemática.
- A saída deve indicar se a expressão está balanceada ou não.