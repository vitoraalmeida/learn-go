package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// waitForResult()
	// fanOut()

	// waitForTask()
	//pooling()

	// Advanced patterns
	// 		fanOutSem()
	// 		boundedWorkPooling()
	// 		drop()

	// Cancellation Pattern
	cancellation()
	// 		cancellationContext()

	// Retry Pattern
	// 		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 		defer cancel()
	// 		retryTimeout(ctx, time.Second, func(ctx context.Context) error { return errors.New("always fail") })

	// Channel Cancellation
	// 		stop := make(chan struct{})
	// 		channelCancellation(stop)
}

func waitForTask() { //chefe de cozinha

	//janela entre cozinha e salão
	ch := make(chan string)

	go func() { //garçon
		//garçon fica a postos esperando a comida aparecer na janela para
		//entregar <-ch (bloqueado)
		p := <-ch
		//quando a comida aparece na janela (p recebe o que veio) ele pode
		//ir entregar. Como ele fica aguardando isso acontecer, está mais
		//focado e percebe o momento antes do chefe.

		fmt.Println("garçon: recebeu:", p)
	}()

	//preparando comida
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	ch <- "comida" //coloca a comida na janela
	//o chefe se da conta que a janela está vazia (sinal de que entregou ao
	//garçon) pouco tempo depois do garçon.
	fmt.Println("chefe: garçon recebeu comida")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

func waitForResult() {
	ch := make(chan string)

	go func() { //empregado
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		//empregado não precisa já sabe o que precisa fazer e envia trabalho
		//quando termina.
		ch <- "paper"
		//ele é certificado de que o chefe ja recebeu depois que o chefe foi
		//notificado que estava pronto
		fmt.Println("employee : sent signal")
	}()

	//chefe esperando o empregado entregar o trabalho
	p := <-ch
	//como ele está esperando, mais focado, antes do empregado perceber que o
	//chefe notou, o chefe nota que já foi entregue
	fmt.Println("manager : recv'd signal :", p)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

func waitForFinished() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		close(ch)
		fmt.Println("Porteiro: aviso de fechamento dado")
	}()

	//não importa o dado (-), precisa saber se veio dado ou não
	//wd = false sem dado (closed)/ true com dado
	_, wd := <-ch
	fmt.Println("Aviso de fechamento recebido. Aberto =", wd) //false
	time.Sleep(time.Second)
}

// pooling: You are a manager and you hire a team of employees. None of the new
// employees know what they are expected to do and wait for you to provide work.
// When work is provided to the group, any given employee can take it and you
// don't care who it is. The amount of time you wait for any given employee to
// take your work is unknown because you need a guarantee that the work your
// sending is received by an employee.
func pooling() {
	ch := make(chan string)

	const g = 2
	//g := runtime.GOMAXPROCS(0)

	//cria pool de goroutines
	for e := 0; e < g; e++ {
		go func(emp int) {
			//as goroutines ficam esperando algo vir pela channel.
			//o scheduller vai definir qual goroutine pegara o trabalho
			for p := range ch {
				fmt.Printf("employee %d : recv'd signal : %s\n", emp, p)
			}
			fmt.Printf("employee %d : recv'd shutdown signal\n", emp)
		}(e)
	}

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper"
		//enquando existirem goroutines sem fazer nada, manda pela channel,
		//se todas estiverem ocupadas, fica bloqueado esperando uma ficar livre
		fmt.Println("manager : sent signal :", w)
	}

	close(ch)
	fmt.Println("manager : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// fanOut: You are a manager and you hire one new employee for the exact amount
// of work you have to get done. Each new employee knows immediately what they
// are expected to do and starts their work. You sit waiting for all the results
// of the employees work. The amount of time you wait on the employees is
// unknown because you need a guarantee that all the results sent by employees
// are received by you. No given employee needs an immediate guarantee that you
// received their result.
func fanOut() {
	emps := 20

	//buffered channels devem ser usados quando se tem conhecimento do contexto,
	//dos fatores envolvidos para poder escolher o tamasnho do buffer
	//o buffer não pode ficar cheio, pois se os sinais não puderem ser enviados
	//pode causar blocking
	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func(emp int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
			//quando uma goroutine envia, a outra tem que esperar terminar.
			//há latencia no envio
			fmt.Println("employee : sent signal :", emp)
		}(e)
	}

	for emps > 0 {
		//não há latencia no recebimento. uma vez que foi enviado, o recebimento
		//ocorre na hora
		p := <-ch
		emps--
		fmt.Println(p)
		fmt.Println("manager : recv'd signal :", emps)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// fanOutSem: You are a manager and you hire one new employee for the exact amount
// of work you have to get done. Each new employee knows immediately what they
// are expected to do and starts their work. However, you don't want all the
// employees working at once. You want to limit how many of them are working at
// any given time. You sit waiting for all the results of the employees work.
// The amount of time you wait on the employees is unknown because you need a
// guarantee that all the results sent by employees are received by you. No
// given employee needs an immediate guarantee that you received their result.
func fanOutSem() {
	emps := 20
	ch := make(chan string, emps)

	//numero maximo de goroutines executando no mesmo momento.
	const cap = 5
	//g := runtime.GOMAXPROCS(0)

	//semaforo com capacidade = numero maximo de goroutines
	sem := make(chan bool, cap)

	//cria todas as goroutines
	for e := 0; e < emps; e++ {
		go func(emp int) {
			//se tiver espaço no semaforo, a goroutine envia dados pela chanel
			//fazendo o espaço no semaforo diminuir em 1
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				ch <- "paper"
				fmt.Println("employee : sent signal :", emp)
			}
			//terminou o trabalho, libera espaço
			<-sem
		}(e)
	}

	for emps > 0 {
		p := <-ch
		emps--
		fmt.Println(p)
		fmt.Println("manager : recv'd signal :", emps)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// drop: You are a manager and you hire a new employee. Your new employee
// doesn't know immediately what they are expected to do and waits for
// you to tell them what to do. You prepare the work and send it to them. The
// amount of time they wait is unknown because you need a guarantee that the
// work your sending is received by the employee. You won't wait for the
// employee to take the work if they are not ready to receive it. In that case
// you drop the work on the floor and try again with the next piece of work.
func drop() {
	//buffer de 5
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		//espera por um sinal de send. quando chega, excecuta algo e espera de
		//e espera de novo
		for p := range ch {
			fmt.Println("employee : recv'd signal :", p)
		}
	}()

	const work = 10
	for w := 0; w < work; w++ {
		select {
		//se tiver como enviar, se ainda houver espaço, passa o dado
		case ch <- "paper":
			fmt.Println("manager : sent signal :", w)
		//se não, apenas segue em frente
		default:
			fmt.Println("manager : dropped data :", w)
		}
	}

	close(ch)
	fmt.Println("manager : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// cancellation: You are a manager and you hire a new employee. Your new
// employee knows immediately what they are expected to do and starts their
// work. You sit waiting for the result of the employee's work. The amount
// of time you wait on the employee is unknown because you need a
// guarantee that the result sent by the employee is received by you. Except
// you are not willing to wait forever for the employee to finish their work.
// They have a specified amount of time and if they are not done, you don't
// wait and walk away.

func cancellation() {
	duration := 150 * time.Millisecond

	//se não criar um buffer, depois do tempo ter se expirado, a goroutine main
	//sai do estado de espera(bloqueio) e segue em frente. Mas quando a outra
	//goroutine tiver terminado o trabalho, ela ira enviar o dado pela chanell
	//e nunca vai ter resposta de recebimento, ficando bloqueada para sempre
	ch := make(chan string, 1)

	go func() {
		//vai demorar tempos aleatorios de 0 a 500, provavelmente, alguma hora
		//ira demorar mais que o tempo maximo permitido
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
	}()

	//quando o tempo definido se esgotar, cria uma channel e envia um dado
	tc := time.After(duration)

	select {
	case d := <-ch:
		fmt.Println("work complete", d)

	case t := <-tc:
		fmt.Println("work cancelled: timeout:", t)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

//with context
func cancellationContext() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "paper"
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)

	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// retryTimeout: You need to validate if something can be done with no error
// but it may take time before this is true. You set a retry interval to create
// a delay before you retry the call and you use the context to set a timeout.
func retryTimeout(ctx context.Context, retryInterval time.Duration, check func(ctx context.Context) error) {

	for {
		fmt.Println("perform user check call")
		if err := check(ctx); err == nil {
			fmt.Println("work finished successfully")
			return
		}

		fmt.Println("check if timeout has expired")
		if ctx.Err() != nil {
			fmt.Println("time expired 1 :", ctx.Err())
			return
		}

		fmt.Printf("wait %s before trying again\n", retryInterval)
		t := time.NewTimer(retryInterval)

		select {
		case <-ctx.Done():
			fmt.Println("timed expired 2 :", ctx.Err())
			t.Stop()
			return
		case <-t.C:
			fmt.Println("retry again")
		}
	}
}

// channelCancellation shows how you can take an existing channel being
// used for cancellation and convert that into using a context where
// a context is needed.
func channelCancellation(stop <-chan struct{}) {

	// Create a cancel context for handling the stop signal.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// If a signal is received on the stop channel, cancel the
	// context. This will propagate the cancel into the p.Run
	// function below.
	go func() {
		select {
		case <-stop:
			cancel()
		case <-ctx.Done():
		}
	}()

	// Imagine a function that is performing an I/O operation that is
	// cancellable.
	func(ctx context.Context) error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.ardanlabs.com/blog/index.xml", nil)
		if err != nil {
			return err
		}
		_, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		return nil
	}(ctx)
}

// boundedWorkPooling: You are a manager and you hire a team of employees. None of
// the new employees know what they are expected to do and wait for you to
// provide work. The amount of work that needs to get done is fixed and staged
// ahead of time. Any given employee can take work and you don't care who it is
// or what they take. The amount of time you wait on the employees to finish
// all the work is unknown because you need a guarantee that all the work is
// finished.
func boundedWorkPooling() {
	work := []string{"paper", "paper", "paper", "paper", "paper", 2000: "paper"}

	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for e := 0; e < g; e++ {
		go func(emp int) {
			defer wg.Done()
			for p := range ch {
				fmt.Printf("employee %d : recv'd signal : %s\n", emp, p)
			}
			fmt.Printf("employee %d : recv'd shutdown signal\n", emp)
		}(e)
	}

	for _, wrk := range work {
		ch <- wrk
	}
	close(ch)
	wg.Wait()

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}
