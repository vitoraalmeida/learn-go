## Em algum ponto:

* Nos impressionamos com programas que possuem grandes quantidades de linhas de código
  Go tenta reduzir a quantidade de minimizar a quantidade de código que precisamos ter nos projetos

* Nos esforçamos para criar grandes abstrações em nosso código fonte
  Nós precisamos de abstrações, mas go propõe camadas finas de abstração

* Nos esquecemos que o hardware é a plataforma
  Não há VMs entre programs go e o hardware
  Go torna fácil colocar simpatia mecânica em ação

* Nos perdemos o entendimento que toda decisão tem consequências

## Questionamentos a se fazer

* É um bom programa
* É eficiente
* É correto
* Termina num prazo correto?
* O que custa

## Aspirar

* Ser um campeão pela qualidade, eficiencia e simplicidade
* Ter um ponto de vista
* Valorizar introspecção e auto-conhecimento


## É importante saber ler código

* Alan Kay - Se a maioria dos programadores estão em falta com entendimento e conhecimento, então o que eles escolhem também estão em falta
* Tom Love (Objective c) - O négócio de software é um dos poucos lugares onde ensinamos pessoas primeiro a escrever antes de ler

## Código legado

* Existem dois tipos de projetos de software: aqueles que falham e aqueles que se tornam legados horríveis - Peter Weinberger (AWK)
* Poucos programas são escritos tendo em mente que serão mantidos e modificados no futuro e por outras pessoas.
* Achamos que códigos ruins são escritos por devs ruins, mas na verdade são escritos por desenvolvedores razoáveis e situações ruins - Sarah Mei


Go permite que façamos mais coisas em 10000 linhas de código que outras linguagens

* Os bugs mais difíceis são aqueles em que seu modelo mental da situação está incorreto, então você não pode entender o problema de fato. - Brian Kernighan
* Todo mundo sabe que debugar é 2x mais difícil que escrever um programa. Então se você não está sendo simples e direto enquanto escreve, como você vai debugar? - BK
* Debuggers não removem bugs, apenas os fazem rodar mais devagar - Desconhecido
* Validar logs, validar modelos mentais e usar o debugger como último recurso.

## Produtividade x Performace

* Go nos deixa ser produtivos e ter a performace que precisamos.
* Rapido suficiente por padrão na maioria dos casos
* Nos deixamos supervalorizar a produtividade em detrimento da performace acreditando que o hardware sempre vai evoluir para compensar
* Go tem a proposta de usuar toda a capacidade do hardware, mas manter a produtividade, pois se não a performance não vai importar se o projeto não ficar pronto

## Correctness vs Performace

Go tem boas ferramentas par saber se nosso programa é rápido o suficiente

Wes Dyer - Faça correto, faça compreensível, concíso e depois faça ser rápido
 
Jaana Dogan - Boa engenharia é menos sobre achar a solução perfeita e mais sobre entender as vantagens e desvantagens e está apto a explicá-las

Otimização deve ser feita, mas não prematuramente - Al Aho (AWK)

## code reviews

* De modo geral, um pedaço de código aloca memória, lê um pedaço de memória e escreve num pedaço de memória

### Integridade

Precisamos ser sérios sobre confiabilidade

Existem 2 forças que estão por trás da integridade:
1. Integridade é sobre cada alocaç~so, leitura e escrita de memória ser precisa, consistente e eficiente. O sistema de tipos é critico para ter certeza que temos esse level micro de integridade
2. Itegridade é sobre cada transformação de dados ser correta, consistente e eficiente. Escrever menos código e lidar com erros de forma adequada é critico para ter certeza de que temos esse level macro de integridade

Go é ums linguagem orietada a dados. todo problema é um problema de transformação de dados
Se não entendemos o dado não entendemos o problema

#### escrever menos código
A média é de 15 a 50 bugs a cada 1000 linhas de código.

#### Error handling

Tratamento de erros em Go é tedioso e é uma coisa boa, pois sempre lidamos com eles e são direto ao ponto

Jaana Dogan - Falha é esperada, falha não é um caso isolado. Crie sistemas que ajudem a identificar falhas. Crie sistemas que ajudam a recuperar de falhas.
Kelsey Hightower - Excelência de produto é a diferença entre algo que apeas funciona sobre certas condições e algo que apenas quebra em certas condições

### Legibilidade

Não escodermos o custo do código - não colocar abstroções acima de abstrações ...
É precisa ser prevísivel o que acontecerá
simplicidade deve ser parte do processo e refatoração deve ser constate

## Performace importa

Motivos para o software não ser rápido

1. Latência externa -> microserviços -> millisegundos de latencias
2. Latência interna -> Coletor de lixo, sincronização e orquestração (concorrência) -> microsegundos
3. Acesso a dados na máquina
4. Eficiência de algorítmos


