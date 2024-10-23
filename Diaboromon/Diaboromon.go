package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"

	pb "digimon/grpc-server/proto"
)

// Funcion para leer el input.txt y obtener el dato TD
func leerInput(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, fmt.Errorf("error abriendo el archivo: %v", err)
	}
	defer file.Close()

	var TD int
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) < 3 {
			return 0, fmt.Errorf("formato incorrecto en el archivo INPUT.txt")
		}
		TD, err = strconv.Atoi(parts[2]) // Tercer valor es TD
		if err != nil {
			return 0, fmt.Errorf("error convirtiendo TD: %v", err)
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error leyendo el archivo: %v", err)
	}

	return TD, nil
}

// Funcion para simular los ataques de diaboromon hacia nodotai cada TD segundos
func atacarNodoTai(TD int, tai pb.DiaboromonServiceClient) {
	//Esperar TD segundos antes de realizar el primer ataque
	time.Sleep(time.Duration(TD) * time.Second)

	for {
		// Enviar ataque
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		response, err := tai.EnviarAtaque(ctx, &pb.Ataque{Dano: 10}) // Enviamos 10 puntos de daño

		//Verificamos si se detiene Diaboromon
		senal := response.GetValor()

		if senal == 1 {
			fmt.Println("Diaboromon: Diaboromon ha sido derrotado! NOOOOO")
			return
		}

		if senal == 0 {
			fmt.Println("Diaboromon: Diaboromon ha Vencido a Tai! HIHIJAJA")
			return
		}

		if err != nil {
			log.Printf("Error enviando ataque a NodoTai: %v", err)
		} else {
			fmt.Println("Diaboromon: Diaboromon ha realizado 10 puntos de daño a Tai!")
		}

		cancel()

		time.Sleep(time.Duration(TD) * time.Second)
	}
}

// Funcion principal con toda la funcionalidad de diaboromon
func main() {
	TD, err := leerInput("INPUT.txt")
	if err != nil {
		log.Fatalf("Error leyendo INPUT.txt: %v", err)
	}

	// Conectar a NodoTai (asumiendo que escucha en localhost:50071)
	conn, err := grpc.Dial("dist038:50071", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error conectando con NodoTai: %v", err)
	}
	defer conn.Close()

	tai := pb.NewDiaboromonServiceClient(conn)

	// Iniciar los ataques a NodoTai cada TD segundos
	fmt.Printf("Diaboromon: Iniciando Ataques a nodoTai cada %d segundos\n", TD)
	atacarNodoTai(TD, tai)
}
