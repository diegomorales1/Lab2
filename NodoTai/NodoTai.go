package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"

	pb "digimon/grpc-server/proto"
)

// Estructura para almacenar los valores de CD y VI
type TaiNode struct {
	CD  float64
	VI  int
	TD  int
	FIN int
}

var mu sync.Mutex // Mutex para modificar VI

// Funcion para leer el input.txt y formar un struct por TaiNode para rellenar cada variable de este ultimo
func leerInput(fileName string) (*TaiNode, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error abriendo el archivo: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tai TaiNode

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ",")
		if len(values) != 5 {
			continue // Salta líneas que no tienen el formato correcto
		}

		tai.TD, err = strconv.Atoi(values[2]) // TD
		if err != nil {
			return nil, fmt.Errorf("error convirtiendo TD: %v", err)
		}

		tai.CD, err = strconv.ParseFloat(values[3], 64) // CD
		if err != nil {
			return nil, fmt.Errorf("error convirtiendo CD: %v", err)
		}

		tai.VI, err = strconv.Atoi(values[4]) // VI
		if err != nil {
			return nil, fmt.Errorf("error convirtiendo VI: %v", err)
		}

		break // Solo necesitamos una línea, asumimos que solo hay una válida
	}

	tai.FIN = 2

	return &tai, nil
}

// Funcion para conectar con el PrimaryNode
func conectarConPrimario() pb.PrimarioServiceClient {
	conn, err := grpc.Dial("dist040:50060", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al cliente: %v", err)
	}
	return pb.NewPrimarioServiceClient(conn)
}

// Consola para NodoTai para solicitar datos a PrimaryNode
func solicitarDatos(tai *TaiNode, cliente pb.PrimarioServiceClient, done chan struct{}) {
	//Time sleep para que salga print despues de conexiones
	time.Sleep(1 * time.Second)
	fmt.Printf("NodoTai: Número CD necesario para ganar: %f\n", tai.CD)
	for {
		// Proteger variable de la vida antes de printear o verificar
		mu.Lock()
		if tai.VI <= 0 {
			mu.Unlock()
			fmt.Println("NodoTai: Tai ha sido derrotado... NodoTai finaliza.")
			close(done)
			return
		}
		fmt.Printf("NodoTai: Vida actual Tai: %d\n", tai.VI)
		mu.Unlock()

		// Solicitar en consola los datos al nodo primario
		fmt.Print("NodoTai: ¿Desea solicitar datos al nodoPrimario? (Escribir: y): ")
		var opcion string
		fmt.Scanln(&opcion)

		if opcion == "y" {
			// Solicita los datos al cliente
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			// Suponiendo que hay una función en el cliente que regresa un número flotante
			response, err := cliente.SolicitarDatos(ctx, &pb.Empty{})
			if err != nil {
				log.Printf("Error solicitando datos al cliente: %v", err)
				continue
			}

			numeroRecibido := response.GetNumero()
			fmt.Printf("NodoTai: Número CD recibido del PrimaryNode: %f\n", numeroRecibido)

			mu.Lock()
			// Comparar con el valor CD
			if numeroRecibido >= tai.CD {
				fmt.Println("NodoTai: Tai ha vencido a Diaboromon!")
				tai.FIN = 1
				mu.Unlock()
				//Sleep para que Diaboron reciba señal de que perdio
				time.Sleep(time.Duration(tai.TD+10) * time.Second)
				close(done)
				return
			} else {
				fmt.Println("NodoTai: Los datos obtenidos son menores que CD. Tai recibe 10 de daño.")
				tai.VI -= 10
			}
			mu.Unlock()
		}
		//Para evitar preguntar tan rapido
		time.Sleep(1 * time.Second)
	}
}

// Servicio que implementa para recibir el daño de Diaboromon
type NodoTaiServer struct {
	pb.UnimplementedDiaboromonServiceServer
	tai  *TaiNode
	done chan struct{}
}

// Hebra para recibir ataques de Diaboromon
func (s *NodoTaiServer) EnviarAtaque(ctx context.Context, ataque *pb.Ataque) (*pb.Senal, error) {
	// Se bloquea el acceso a la vida
	mu.Lock()
	defer mu.Unlock()

	//Printea esto solo cuando Tai o Diaboron sigan peleando
	if s.tai.FIN == 2 {
		fmt.Printf("\nNodoTai: Diaboromon ha infligido %d puntos de daño\n", ataque.Dano)
		s.tai.VI -= int(ataque.Dano)
		fmt.Printf("NodoTai: Vida restante de NodoTai: %d\n", s.tai.VI)
	}

	if s.tai.VI <= 0 {
		fmt.Println("NodoTai: Tai ha sido derrotado... NodoTai finaliza.")
		close(s.done)
		return &pb.Senal{Valor: int32(0)}, nil
	}

	if s.tai.FIN == 1 {
		//fmt.Println("Diaboromon ha sido derrotado! NodoTai finaliza.")
		//El print deberia realizarse en Diaboromon
		close(s.done)
		return &pb.Senal{Valor: int32(1)}, nil
	}

	fmt.Print("NodoTai: ¿Desea solicitar datos al nodoPrimario? (Escribir: y): ")
	return &pb.Senal{Valor: int32(2)}, nil
}

func main() {
	tai, err := leerInput("INPUT.txt")
	if err != nil {
		log.Fatalf("Error leyendo INPUT.txt: %v", err)
	}

	// Canal para terminar ejecucion de NodoTai
	done := make(chan struct{})

	// Inicia el servidor gRPC para recibir ataques de Diaboromon
	go func() {
		lis, err := net.Listen("tcp", "0.0.0.0:50071") // NodoTai escuchará en el puerto 50061
		if err != nil {
			log.Fatalf("Error al intentar escuchar: %v", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterDiaboromonServiceServer(grpcServer, &NodoTaiServer{tai: tai, done: done})

		log.Println("NodoTai: NodoTai escuchando ataques de Diaboromon en el puerto 50071...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
		}
	}()

	cliente := conectarConPrimario()

	go solicitarDatos(tai, cliente, done)

	// Bloqueo hasta que el canal "done" se cierre, lo que significa que Tai ha sido derrotado o ha ganado
	<-done
	// Conectarse a PrimaryNode y mandar una señal de finalizacion para cerrar todo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = cliente.Finalizar(ctx, &pb.Empty{})
	if err != nil {
		log.Printf("Error al enviar la señal de finalización a NodoPrimario: %v", err)
	}
	time.Sleep(20 * time.Second)
	fmt.Println("NodoTai: NodoTai finalizado y todas las hebras se han cerrado.")
}
