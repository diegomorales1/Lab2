package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	randin "math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "digimon/grpc-server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// Configuración y Struct leída desde INPUT.txt
type Config struct {
	PS float64
	TE float64
	TD float64
	CD int
	VI int
}

// Constante para leer el archivo de info.txt
const archivo string = "INFO.txt"

// Mutex para bloquear la lectura/escritura del archivo cuando sea necesario
var mu sync.Mutex

// struct del servidor
type servidor struct {
	pb.UnimplementedPrimarioServiceServer
	done chan struct{}
}

// Mapa para valores de los atributos
var atributoValores = map[string]float64{
	"Vaccine": 3.0,
	"Data":    1.5,
	"Virus":   0.8,
}

// Funcion para mandar notificacion de finalizacion a los demas nodos
func (s *servidor) Finalizar(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	fmt.Println("PrimaryNode: Señal de finalización recibida. PrimaryNode se apaga.")
	close(s.done) // Cierra el canal para terminar el servidor
	return &pb.Empty{}, nil
}

// funcion para conectarse a los datanodes 1 y 2 dependiendo del puerto
func conectarDatanode(puerto string) pb.InfoServiceClient {
	conn, err := grpc.Dial(puerto, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	return pb.NewInfoServiceClient(conn)
}

// Solicita los datos desde el nodo tai para ver la cantidad de digimons sacrificados guardados desde el datanode, devolviendo el CD total acumulado
func (grpcServer *servidor) SolicitarDatos(ctx context.Context, e *pb.Empty) (*pb.NumeroResponse, error) {
	// Leer el archivo info.txt y buscar estados sacrificados
	mu.Lock()
	file, err := os.Open(archivo)
	if err != nil {
		log.Printf("Error al abrir archivo: %v", err)
		mu.Unlock()
	}

	scanner := bufio.NewScanner(file)
	var total float64

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		if len(parts) != 4 {
			log.Printf("Linea invalida: %s", line)
			continue
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Printf("Error al convertir ID: %v", err)
			continue
		}

		numDataNode := parts[1]
		estado := parts[3]

		if estado == "Sacrificado" {
			if numDataNode == "1" {
				// Solicitar atributo al servidor usando el ID
				// (PUERTO)
				c1 := conectarDatanode("dist037:50054")
				// Enviar datos al servidor
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				atributoResponse, err := c1.ObtenerInfo(ctx, &pb.IdRequest{
					Id: int32(id),
				})
				if err != nil {
					log.Printf("Error al obtener atributo: %v", err)
					continue
				}
				atributo := atributoResponse.GetAtributo()
				valor, ok := atributoValores[atributo]

				if !ok {
					log.Printf("Atributo desconocido: %s", atributo)
					continue
				}
				// Sumar el valor al total
				total += valor
			} else {
				// Verificar si el estado es "Sacrificado" en datanode2
				c2 := conectarDatanode("dist039:50055") //Cambiar a PUERTO correspondiente
				// Enviar datos al servidor
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				atributoResponse, err := c2.ObtenerInfo(ctx, &pb.IdRequest{
					Id: int32(id),
				})
				if err != nil {
					log.Printf("Error al obtener atributo: %v", err)
					continue
				}
				atributo := atributoResponse.GetAtributo()
				valor, ok := atributoValores[atributo]
				if !ok {
					log.Printf("Atributo desconocido: %s", atributo)
					continue
				}
				// Sumar el valor al total
				total += valor
			}
		}
	}
	// Verificar si hubo errores durante el escaneo
	if err := scanner.Err(); err != nil {
		log.Printf("Error al leer el archivo: %v", err)
	}
	file.Close()
	mu.Unlock()
	// Enviar el cd total acumulado al NodoTai
	fmt.Printf("PrimaryNode: Solicitud de NodoTai recibida, mensaje enviado: CD acumulado=%f\n", total)
	return &pb.NumeroResponse{
		Numero: total,
	}, nil
}

// Implementación del ResultServiceServer
type resultServer struct {
	pb.UnimplementedResultServiceServer
	mu      sync.Mutex
	results map[string]bool // Para evitar duplicados
	lista   []string        // Lista de Digimons únicos
}

// lista de digimons global para comprobar duplicaciones
var ListaDigimons []string

// llave para desencriptar y encriptar los datos
var key []byte

// Función para encriptar una lista de strings
func encryptList(plainText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Agregar padding al plaintext y completar lo bytes necesarios
	paddedPlainText := pad([]byte(plainText), aes.BlockSize) // Asegúrate de usar el padded text

	ciphertext := make([]byte, aes.BlockSize+len(paddedPlainText))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedPlainText) // Encriptar el texto con padding

	return hex.EncodeToString(ciphertext), nil
}

// Función para agregar padding al texto
func pad(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

// Función para desencriptar un string y devolver el contenido de este
func decryptString(encoded string, key []byte) (string, error) {
	ciphertext, err := hex.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Quitar padding del texto desencriptado
	plainText, err := unpad(ciphertext)
	if err != nil {
		return "", err
	}

	// Retornar el texto desencriptado como un string
	return string(plainText), nil
}

// Función para quitar padding del texto
func unpad(ciphertext []byte) ([]byte, error) {
	length := len(ciphertext)
	unpadding := int(ciphertext[length-1])
	if unpadding > length {
		return nil, fmt.Errorf("padding size error")
	}
	return ciphertext[:(length - unpadding)], nil
}

// Funcion para agregar id unicos a cada digimon en un rango de 5000 ids al azar
func generateUniqueID() int64 {
	return randin.Int63n(5000)
}

// Funcion que dependiendo del nombre del digimon sera guardado en un datanode distinto
func determineDataNode(nombre string) int {
	if nombre == "" {
		return 0
	}
	firstChar := strings.ToUpper(string(nombre[0]))
	if firstChar >= "A" && firstChar <= "M" {
		return 1
	} else if firstChar >= "N" && firstChar <= "Z" {
		return 2
	}
	return 0 // Opcional: manejar caracteres no alfabéticos
}

// Escribe la información del Digimon en INFO.txt
func writeToInfoFile(id int64, dataNode int, nombre, estado string) error {
	// Abrir el archivo en modo append, lo crea si no existe
	file, err := os.OpenFile("INFO.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Formato: ID,NumDataNode,Nombre,Estado
	line := fmt.Sprintf("%d,%d,%s,%s\n", id, dataNode, nombre, estado)
	if _, err := file.WriteString(line); err != nil {
		return err
	}

	return nil
}

// funcion para determinar el sacrificio de un digimon dado por los servidores regionales, ve duplicidad de digimon y escribe un info.txt para cada resultado calculado
// ademas de enviar la informacion correspondiente a cada datanode con el id del digimon y su estado.
func (s *resultServer) SendSacrificioResult(ctx context.Context, result *pb.SacrificioResult) (*pb.Ack, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Clave para desencriptar
	key := []byte("somoslosverdaderosfokinfansdeosi")

	// Desencriptar el mensaje encriptado (retorna un string)
	decryptedString, err := decryptString(result.MensajeEncriptado, key)
	if err != nil {
		log.Printf("Error al desencriptar el mensaje: %v", err)
		return &pb.Ack{Success: false, Message: "Error al desencriptar el mensaje"}, nil
	}
	// Dividir la lista en digimons separados por '#'
	digimonList := strings.Split(decryptedString, "#")

	// Iterar sobre todos los digimons en la lista desencriptada
	for _, digimonStr := range digimonList {
		if digimonStr == "" {
			continue // Ignorar entradas vacías
		}

		// Cada digimonStr tiene el formato "Nombre,Atributo,Estado"
		semipart := strings.Split(digimonStr, ",")
		if len(semipart) != 3 {
			log.Printf("Formato inválido: %s", digimonStr)
			continue
		}
		nombre := semipart[0]
		atributo := semipart[1]
		estado := semipart[2]

		// Formato final del mensaje
		message := fmt.Sprintf("%s,%s,%s", nombre, atributo, estado)

		// Verificar si el nombre del Digimon ya está en la lista global
		if !containsByName(ListaDigimons, nombre) {
			ListaDigimons = append(ListaDigimons, message) // Agregar a la lista global
			log.Printf("PrimaryNode: Nuevo Digimon agregado: %s", message)

			// Generar un ID único
			id := generateUniqueID()

			// Determinar el Data Node
			dataNode := determineDataNode(nombre)

			// Escribir en INFO.txt
			err := writeToInfoFile(id, dataNode, nombre, estado)
			if err != nil {
				log.Printf("Error al escribir en INFO.txt: %v", err)
				continue
			}
			// Enviar datos al data Node 1
			if dataNode == 1 {
				c1 := conectarDatanode("dist037:50054")
				// Enviar datos al servidor
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				response, err := c1.MandarInfo(ctx, &pb.InfoNode{
					Id:       int32(id),
					Atributo: atributo,
				})
				if err != nil {
					log.Fatalf("PrimaryNode: Error al enviar la información al DataNode1: %v", err)
				}
				fmt.Printf("PrimaryNode: Mensaje enviado hacia el DataNode 1: %d, %s\n", id, atributo)
				fmt.Printf("PrimaryNode: Respuesta recibida por DataNode 1: %s\n", response)
			} else {
				c2 := conectarDatanode("dist039:50055")
				// Enviar datos al servidor
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				response, err := c2.MandarInfo(ctx, &pb.InfoNode{
					Id:       int32(id),
					Atributo: atributo,
				})
				if err != nil {
					log.Fatalf("Error al enviar la información al DataNode2: %v", err)
				}
				fmt.Printf("PrimaryNode: Mensaje enviado hacia el DataNode 2: %d, %s\n", id, atributo)
				fmt.Printf("PrimaryNode: Respuesta recibida por DataNode 2: %s\n", response)
			}
		} else {
			log.Printf("PrimaryNode: Digimon duplicado ignorado: %s", message)
		}
	}

	return &pb.Ack{Success: true, Message: "PrimaryNode: Resultados recibidos y procesados"}, nil
}

// Función para verificar si un nombre está en la lista (ve duplicidad)
func containsByName(list []string, name string) bool {
	for _, v := range list {
		// Extraer el nombre del mensaje
		digimonName := strings.Split(v, ",")[0]
		if digimonName == name {
			return true
		}
	}
	return false
}

// funcion para leer el archivo input.txt y devolver un struct formado por cada variable a utilizar
func readInputConfig(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 5 {
			return Config{}, fmt.Errorf("INPUT.txt debe tener 5 valores separados por ','")
		}

		var config Config
		_, err := fmt.Sscanf(line, "%f,%f,%f,%d,%d", &config.PS, &config.TE, &config.TD, &config.CD, &config.VI)
		if err != nil {
			return Config{}, err
		}

		// Validacion de las restricciones
		if config.PS < 0 || config.PS > 1 {
			return Config{}, fmt.Errorf("PS debe estar entre 0 y 1")
		}
		if config.TE <= 0 || config.TE > 5 {
			return Config{}, fmt.Errorf("TE debe estar entre >0 y ≤5")
		}
		if config.TD <= 0 || config.TD > 10 {
			return Config{}, fmt.Errorf("TD debe estar entre >0 y ≤10")
		}

		return config, nil
	}

	if err := scanner.Err(); err != nil {
		return Config{}, err
	}

	return Config{}, fmt.Errorf("INPUT.txt está vacío")
}

func main() {
	// Leer configuración de INPUT.txt
	config, err := readInputConfig("INPUT.txt")
	if err != nil {
		log.Fatalf("Error al leer Input.txt: %v", err)
	}
	// Iniciar el servidor gRPC
	lis, err := net.Listen("tcp", "0.0.0.0:50050")
	if err != nil {
		log.Fatalf("Failed to listen on port 50050: %v", err)
	}

	grpcServer := grpc.NewServer()
	resultSrv := &resultServer{
		results: make(map[string]bool),
		lista:   make([]string, 0),
	}
	pb.RegisterResultServiceServer(grpcServer, resultSrv)

	// Registrar la reflexión para usar herramientas como grpcurl
	reflection.Register(grpcServer)

	// Ejecutar el servidor en una goroutine
	go func() {
		log.Println("PrimaryNode: Servidor principal está escuchando en el puerto 50050")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve ResultService: %v", err)
		}
	}()

	// Esperar un momento para asegurar que el servidor está corriendo
	time.Sleep(time.Second)

	// Configurar conexiones gRPC a los servidores regionales
	conexiones := map[string]string{
		"IslaFile":         "dist037:50051",
		"ContinenteFolder": "dist038:50052",
		"ContinenteServer": "dist039:50053",
	}

	clientes := make(map[string]pb.DigimonServiceClient)
	var wg sync.WaitGroup

	for atributo, addr := range conexiones {
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("No se pudo conectar a %s en %s: %v", atributo, addr, err)
		}
		defer conn.Close()
		clientes[atributo] = pb.NewDigimonServiceClient(conn)
		log.Printf("PrimaryNode: Conectado al nodo %s en %s", atributo, addr)
	}

	// Enviar la notificación de finalización a los nodos
	for atributo, cliente := range clientes {
		wg.Add(1)
		go func(c pb.DigimonServiceClient, attr string) {
			defer wg.Done()
			if err := notifyCompletion(c, strconv.FormatFloat(config.PS, 'E', -1, 64), strconv.FormatFloat(config.TE, 'E', -1, 64)); err != nil {
				log.Printf("Error al notificar a %s: %v", attr, err)
			} else {
				log.Printf("PrimaryNode: Notificación de PS enviada a %s", attr)
			}
		}(cliente, atributo)
	}

	// Esperar a que todas las notificaciones sean enviadas
	wg.Wait()
	// Iniciar el servidor gRPC para escuchar solicitudes de NodoTai
	listener, err := net.Listen("tcp", "0.0.0.0:50060")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
	fmt.Println("PrimaryNode: Esperando solicitudes de NodoTai...")

	grpcServer = grpc.NewServer()
	s := &servidor{done: make(chan struct{})}
	pb.RegisterPrimarioServiceServer(grpcServer, s)
	go func() {
		fmt.Println("PrimaryNode: escuchando en el puerto 50060...")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
		}
	}()
	<-s.done
	grpcServer.GracefulStop() // Apaga el servidor gRPC de manera ordenada
}

// Función para notificar la finalización a un nodo
func notifyCompletion(client pb.DigimonServiceClient, porcentajeSacrificio string, tiempoEspera string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//Aqui hay que encriptar el mensaje para enviarselo a los servidores regionales
	key = []byte("somoslosverdaderosfokinfansdeosi")
	encript_sacrificio, err1 := encryptList(porcentajeSacrificio, key)
	encript_tiempo, err2 := encryptList(tiempoEspera, key)
	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}
	notification := &pb.CompletionInfo{
		SacrificioPorcentaje: encript_sacrificio,
		TiempoEspera:         encript_tiempo,
	}

	ack, err := client.NotifyCompletion(ctx, notification)
	if err != nil {
		return err
	}

	if ack.Success {
		log.Printf("PrimaryNode: Notificación enviada: %s", ack.Message)
	} else {
		log.Printf("Error al enviar notificación: %s", ack.Message)
	}
	return nil
}
