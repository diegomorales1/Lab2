package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	cryptoRand "crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	pb "digimon/grpc-server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Implementación de DigimonServiceServer
type server struct {
	pb.UnimplementedDigimonServiceServer
	ps       float64
	digimons []Digimon
}

// Digimon struct para leer de DIGIMONS.txt
type Digimon struct {
	Nombre   string
	Atributo string
}

// Función para agregar padding al texto
func pad(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
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

// Función para encriptar una lista de strings
func encryptList(plainText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Agregar padding al plaintext
	paddedPlainText := pad([]byte(plainText), aes.BlockSize) // Asegúrate de usar el padded text

	ciphertext := make([]byte, aes.BlockSize+len(paddedPlainText))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(cryptoRand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedPlainText) // Encriptar el texto con padding

	return hex.EncodeToString(ciphertext), nil
}

// Función para desencriptar un string
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

// Método para recibir la notificación de finalización
func (s *server) NotifyCompletion(ctx context.Context, info *pb.CompletionInfo) (*pb.Ack, error) {
	key := []byte("somoslosverdaderosfokinfansdeosi")
	desencript_sacrificio, err1 := decryptString(info.SacrificioPorcentaje, key)
	desencript_tiempo, err2 := decryptString(info.TiempoEspera, key)

	sacrificio, err3 := strconv.ParseFloat(desencript_sacrificio, 64)
	tiempo, err4 := strconv.ParseFloat(desencript_tiempo, 64) // Tiempo de TE segundos

	if (err3 != nil) && (err1 != nil) && (err2 != nil) && (err4 != nil) {
		panic(err3)
	}

	s.ps = sacrificio
	// Leer DIGIMONS.txt
	digimons, err := readDigimons("DIGIMONS.txt")
	if err != nil {
		log.Printf("Error al leer DIGIMONS.txt: %v", err)
		return &pb.Ack{Success: false, Message: "Error al leer DIGIMONS.txt"}, nil
	}
	s.digimons = digimons
	// Seleccionar 6 Digimons al azar (primera iteración)
	selectedDigimons := selectRandomDigimons(s.digimons, 6)
	if len(selectedDigimons) == 0 {
		log.Println("ContinenteFolder: ContinenteFolder no tiene Digimons para sacrificar.")
		return &pb.Ack{Success: false, Message: "ContinenteFolder: No hay Digimons para sacrificar"}, nil
	}

	// Decidir sacrificio basado en PS y enviar los primeros 6 Digimons
	sendDigimonBatch(selectedDigimons, key, s.ps)

	// Enviar un Digimon cada TE segundos después de los primeros 6
	go func() {
		ticker := time.NewTicker(time.Duration(tiempo) * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if len(s.digimons) == 0 {
				log.Println("ContinenteFolder: ContinenteFolder se quedó sin Digimons para enviar.")
				return
			}

			// Seleccionar y enviar un solo Digimon
			selectedDigimons = selectRandomDigimons(s.digimons, 1)
			sendDigimonBatch(selectedDigimons, key, s.ps)
		}
	}()

	return &pb.Ack{Success: true, Message: "ContinenteFolder: Sacrificio procesado y resultados enviados"}, nil
}

// Función para enviar un lote de Digimons (ya sea 6 o 1 dependiendo del caso a probar)
func sendDigimonBatch(digimons []Digimon, key []byte, sacri float64) {
	var probandouwu string
	for _, digimon := range digimons {
		sacrificio := decideSacrifice(sacri)
		estado := "No-sacrificado"
		if sacrificio {
			estado = "Sacrificado"
		}
		n := digimon.Nombre
		a := digimon.Atributo
		e := estado
		probandouwu = probandouwu + n + "," + a + "," + e + "#"
		fmt.Printf("ContinenteFolder: Estado enviado: %s, %s\n", n, e)
	}

	encrypted, err := encryptList(probandouwu, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("ContinenteFolder: Mensaje encriptado enviado a PrimaryNode:", encrypted)

	// Enviar los resultados al servidor principal
	err = sendResults("dist040:50050", encrypted)
	if err != nil {
		log.Printf("Error al enviar resultados al servidor principal: %v", err)
	}
}

// Función para leer DIGIMONS.txt
func readDigimons(filePath string) ([]Digimon, error) {
	var digimons []Digimon

	file, err := os.Open(filePath)
	if err != nil {
		return digimons, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			log.Printf("Línea inválida en DIGIMONS.txt: %s", line)
			continue
		}
		nombre := strings.TrimSpace(parts[0])
		atributo := strings.TrimSpace(parts[1])
		digimons = append(digimons, Digimon{Nombre: nombre, Atributo: atributo})
	}

	if err := scanner.Err(); err != nil {
		return digimons, err
	}

	return digimons, nil
}

// Función para seleccionar aleatoriamente n Digimons
func selectRandomDigimons(digimons []Digimon, n int) []Digimon {
	if len(digimons) <= n {
		return digimons
	}

	selected := make([]Digimon, 0, n)
	perm := rand.Perm(len(digimons))
	for i := 0; i < n; i++ {
		selected = append(selected, digimons[perm[i]])
	}
	return selected
}

// Función para decidir si sacrificar basado en PS
func decideSacrifice(ps float64) bool {
	return rand.Float64() <= ps
}

// Función para enviar resultados al Primary Node
func sendResults(serverAddr, resultados string) error {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("no se pudo conectar al servidor principal: %v", err)
	}
	defer conn.Close()

	client := pb.NewResultServiceClient(conn)

	sacResult := &pb.SacrificioResult{
		MensajeEncriptado: resultados,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	ack, err := client.SendSacrificioResult(ctx, sacResult)
	if err != nil {
		log.Fatalf("ContinenteFolder: sin conexion con PrimaryNode, finalizando ejecucion")
	}

	if ack.Success {
		log.Printf("Servidor respondió: %s", ack.Message)
	} else {
		log.Printf("Servidor informó error: %s", ack.Message)
	}

	return nil
}

// Funcion principal que llama toda la funcionalidad necesaria como servidor regional
func main() {
	rand.Seed(time.Now().UnixNano())
	lis, err := net.Listen("tcp", "0.0.0.0:50052") // Puerto para ContinenteFolder
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDigimonServiceServer(grpcServer, &server{
		digimons: make([]Digimon, 0),
	})

	log.Println("ContinenteFolder: ContinenteFolder está escuchando en el puerto 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve DigimonService: %v", err)
	}
}
