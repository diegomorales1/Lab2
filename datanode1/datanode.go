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
	"time"

	pb "digimon/grpc-server/proto"

	"google.golang.org/grpc"
)

// Constante para generar el archivo correspondiente al datanode y guardar los id de los digimons con su estado
const fileName = "DATA_1.txt"

// Struct del servidor para la comunicacion por grpc
type server struct {
	pb.UnimplementedInfoServiceServer
	resetTimer chan bool // Canal para reiniciar el temporizador
}

// MandarInfo recibe los datos de primarynode y los almacena en el archivo info.txt
func (s *server) MandarInfo(ctx context.Context, in *pb.InfoNode) (*pb.Response, error) {
	// Reiniciar temporizador
	s.resetTimer <- true

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("DATANODE1: error abriendo el archivo: %v", err)
	}
	defer file.Close()

	// Aquí se guarda la info en el formato
	data := fmt.Sprintf("%d,%s\n", in.GetId(), in.GetAtributo())

	if _, err := file.WriteString(data); err != nil {
		return nil, fmt.Errorf("DATANODE1: error escribiendo en el archivo: %v", err)
	}

	fmt.Printf("DATANODE1: Información de PrimaryNode recibida, Mensaje recibido: %d, %s\n", in.GetId(), in.GetAtributo())

	return &pb.Response{Message: "DATANODE1: Información guardada correctamente"}, nil
}

// ObtenerInfo lee el archivo info.txt y envía todos los datos almacenados
func (s *server) ObtenerInfo(ctx context.Context, in *pb.IdRequest) (*pb.AtributoResponse, error) {
	// Reiniciar temporizador
	s.resetTimer <- true

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("DATANODE1: error al abrir archivo: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		if len(parts) != 2 {
			continue
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		if int32(id) == in.GetId() {
			// Encuentro la id correspondiente, mando atributo
			fmt.Printf("DATANODE1: Solicitud de PrimaryNode recibida, Mensaje enviado: %s\n", parts[1])
			return &pb.AtributoResponse{
				Atributo: parts[1],
			}, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error al leer archivo: %v", err)
	}

	// Si no se encuentra la ID
	return nil, fmt.Errorf("DATANODE1: ID %d no encontrada", in.GetId())
}

// Funcion principal que contiene toda la funcionalidad de un datanode
func main() {
	// Crear el canal para reiniciar el temporizador
	resetTimer := make(chan bool)

	lis, err := net.Listen("tcp", "0.0.0.0:50054")
	if err != nil {
		log.Fatalf("Fallo al intentar escuchar: %v", err)
	}

	// Crear servidor gRPC
	s := grpc.NewServer()
	serv := &server{resetTimer: resetTimer}
	pb.RegisterInfoServiceServer(s, serv)

	// Crear un temporizador que finaliza el servidor si no recibe mensajes durante 30 segundos
	go func() {
		timer := time.NewTimer(60 * time.Second)
		for {
			select {
			case <-timer.C:
				// Se cumplió el tiempo sin recibir mensajes
				fmt.Println("DATANODE1: No se recibieron mensajes en 30 segundos. Cerrando servidor.")

				err := os.Truncate(fileName, 0)
				if err != nil {
					log.Printf("DATANODE1: Error al vaciar el archivo: %v", err)
				} else {
					log.Printf("DATANODE1: Archivo %s vaciado correctamente", fileName)
				}

				s.GracefulStop() // Finalizar el servidor de manera segura
				return
			case <-resetTimer:
				// Reiniciar el temporizador cuando se reciba una solicitud
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(30 * time.Second)
			}
		}
	}()

	// Iniciar el servidor
	log.Printf("DATANODE1: Servidor escuchando en %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("DATANODE1: Fallo al conectar al server: %v", err)
	}
}
