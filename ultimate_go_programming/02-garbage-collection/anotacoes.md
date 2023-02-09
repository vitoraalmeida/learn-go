Variável de configuração GOGC = 100
É relacionado com quando a 1ª coleta deve ocorrer

Primeira coleta começa quando a memória heap atinge 4MB quando GOGC = 100

O GC é executado em 3 fases

1. Mark Start ->  Stop the world -> não ocorre código de aplicação executando
2. Marking    -> Coleta acontece ao mesmo tempo em que aplicação é executada
3. Mark Termination -> STW

O tempo que as 3 fases levam para terminar é o tempo que o GC leva executando
Não é determinístico o tempo que levará

O GC tenta minimizar latência e maximizar throughput enquanto limpa a heap

O GC tenta fazer com que a fase de STW leve menos de 100microsegundos

Enquanto o GC está rodando há cerca de 100microsegundos de STW,
Fora dos momentos de STW, marking, o GC usa 25% das goroutines para executar
paralelamente ao código de aplicação.

Para reduzir o número de GC é importante reduzir alocações e o número de goroutines
necessárias 

Então se temos 4 processadores lógicos, 4 SO threads, temos 4 goroutines.
Quando entra na fase 1, todas as goroutines param de executar
Na fase 2, 1 das goroutines é do GC e o restante da aplicação,
Na fase 3, todas as goroutines param novamente

No fim, então todas as goroutines estão disponíveis para a aplicação

Execudar

`GODEBUG=gctrace=1 ./app`

vai fazer com que seja mostrado as informações sobre o GC

O GC vai marcar quais endereços de memória na Heap estão sendo de fato
referenciados na stack. Depois que acha esses endereços, o restante é colocado
a disposição do alocador para ser reutilizado quando o STW acabar.

O foco do GC é reduzir ao máximo o tamanho ocupado da head e não em fazer isso
mais rapidamente.

A próxima GC vai acontecer quando o uso da heap chegar no dobro do valor que 
estava em uso na última

O GC começa o trabalho com uma meta de quanto ele espera ter de uso na heap
depois que ele terminar o trabalho

Se a meta da primeira coleta é 10MB
E o uso atual é 7MB antes de começar o trabalho de marcar o que está vivo (em uso)
Depois de terminar o trabalho de marcar, ele vê que tem 11MB em uso (até pq
as Goroutines de aplicação restantes continuaram executando).
No entanto, dessas 11, apenas 6 estão marcados. Então o restante será disponibilizado
para reutilizar.

A próxima coleta vai começar quando o uso de heap chegar em 2 x 6MB se o GOGC estiver
em 100. Pois é o valor atual + 100% do valor atual.

Quando o resultado está acima da meta, o GC está começando a perder controle
do crescimento de memória. Então ele faz com que uma das Goroutines pare de 
executar código de aplicação para reduzir alocações e auxiliar no GC. Não durante
todo o tempo de GC, apenas uma parcelo 30 ou 40 microsegundo. Mas nesses casos,
durante o GC, temos apenas 50% (ainda no caso de termos 4 Goroutines) de CPU em 
uso para aplicação

Para evitar isso, o GC pode começar a executar antes do valor determinado na 
coleta anterior, para que ele não chegue no ponto de precisar usar uma
goroutine para auxiliar no GC, se perceber que o valor está crescendo muito rapido
ou de forma que ele perca controle nas coletas posteriores

O foco não deve ser escrever um software com 0 alocações (não é possível), mas sim
reduzir o número de alocações. E também devemos lembrar que ocorrem diversas coletas
no decorrer do tempo, então a quantidade de tempo total de latência será afetado pelo GC

Então se cada requisição que fazemos está alocando 4MB, a cada requisição fazemos
uma coleta. Então se conseguirmos reduzir a alocação por request para 2MB, teremos
2 requests antes que uma coleta aconteça.

Então podemos usar ferramnetas de profiling para identificar focos de alocação
para que possamos fazer esforço para reduzir e assim diminuir o número de coletas
