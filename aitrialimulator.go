package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc/status"
)

const systemInstructionsFile = "system-instructions.md"

var sleepTime = struct {
	character time.Duration
	sentence  time.Duration
}{
	character: time.Millisecond * 30,
	sentence:  time.Millisecond * 300,
}

// Streaming output column position.
var col = 0

// getBytes returns the file contents as bytes.
func getBytes(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file bytes %v: %v\n", path, err)
	}
	return bytes
}

// newClient creates a new API client using API_KEY environment variable.
func newClient(ctx context.Context) *genai.Client {
	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		log.Fatalf("Environment variable API_KEY is not set.\n" +
			"To obtain an API key, visit https://aistudio.google.com/, select 'Get API key'.\n")
	}

	// New client, using API key authorization.
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v\n", err)
	}
	return client
}

func main() {
	ctx := context.Background()
	client := newClient(ctx)
	defer client.Close()

	// Configure desired model.
	model := client.GenerativeModel("gemini-2.0-flash-exp")

	// Initialize new chat session.
	session := model.StartChat()

	// Set system instructions.
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(getBytes(systemInstructionsFile))},
		Role:  "system",
	}

	// Introducción al juicio
	introduction := "Te encuentras en una sala de juicio. Eres un abogado defensor, y el juez está esperando que presentes tu defensa en un caso grave."

	// Establecer el historial de la sesión con una introducción.
	session.History = []*genai.Content{{
		Role:  "model",
		Parts: []genai.Part{genai.Text(introduction)},
	}}

	printRuneFormatted('\n')

	// Preguntar sobre el tipo de caso
	caseType := askUser("¿De qué tipo es el caso? (ej., criminal, laboral, familiar, etc.)")
	sendAndPrintResponse(ctx, session, "El caso es de tipo: "+caseType)

	// Preguntar sobre la defensa según el tipo de caso
	defenseOptions := getDefenseOptions(caseType)
	defenseType := askUser(fmt.Sprintf("¿Cuál es la naturaleza de tu defensa en este caso? Elige una de las siguientes: %s", strings.Join(defenseOptions, ", ")))
	sendAndPrintResponse(ctx, session, "Tu defensa es: "+defenseType)

	// Iniciar simulación del juicio
	trialSimulation(ctx, session)
}

// getDefenseOptions devuelve las opciones de defensa según el tipo de caso.
func getDefenseOptions(caseType string) []string {
	switch strings.ToLower(caseType) {
	case "criminal":
		return []string{"legítima defensa", "error de identidad", "falta de pruebas", "prescripción del delito"}
	case "laboral":
		return []string{"despido injustificado", "violación de derechos laborales", "necesidad de reubicación"}
	case "familiar":
		return []string{"custodia compartida", "violencia doméstica", "desacuerdo sobre bienes matrimoniales"}
	default:
		return []string{"defensa desconocida"}
	}
}

// trialSimulation simula el flujo de un juicio.
func trialSimulation(ctx context.Context, session *genai.ChatSession) {
	for {
		// Escenario inicial: El juez solicita que se presente la defensa
		courtState := "El juez te llama a la defensa. ¿Qué dirás en tu declaración inicial?"
		sendAndPrintResponse(ctx, session, courtState)

		// El usuario responde con su declaración inicial
		initialStatement := askUser(">>")
		sendAndPrintResponse(ctx, session, "El abogado presentó la siguiente declaración: "+initialStatement)

		// El fiscal presenta su caso
		prosecutionStatement := "El fiscal presenta su caso, enumerando las pruebas y acusaciones contra tu cliente."
		sendAndPrintResponse(ctx, session, prosecutionStatement)

		// El usuario decide cómo responder a las acusaciones
		action := askUser("¿Cómo responderás a las acusaciones del fiscal? (ej., refutar pruebas, objeción, interrogar testigos)")
		resp := fmt.Sprintf("El abogado defensor decide: %v\n\nDescribe tu siguiente acción.", action)
		sendAndPrintResponse(ctx, session, resp)

		// Continuar con la simulación
		continueTrial(ctx, session)
	}
}

// continueTrial simula el siguiente paso del juicio.
func continueTrial(ctx context.Context, session *genai.ChatSession) {
	// El juez solicita testigos
	witnessQuestion := "El juez te solicita que presentes tu primer testigo. ¿Quién será?"
	sendAndPrintResponse(ctx, session, witnessQuestion)

	// El usuario presenta a su testigo
	witnessName := askUser(">>")
	sendAndPrintResponse(ctx, session, "El abogado presenta al testigo: "+witnessName)

	// El testigo es interrogado
	witnessStatement := "El testigo declara bajo juramento. ¿Cómo lo interrogarás?"
	sendAndPrintResponse(ctx, session, witnessStatement)

	// El usuario decide cómo interrogar al testigo
	questionToWitness := askUser(">>")
	sendAndPrintResponse(ctx, session, "El abogado pregunta al testigo: "+questionToWitness)

	// El testigo responde
	testimonyResponse := "El testigo responde, y ahora el fiscal tiene la oportunidad de contrainterrogar."
	sendAndPrintResponse(ctx, session, testimonyResponse)

	// Preguntar qué hará el abogado después
	action := askUser("¿Qué hará el abogado después del contrainterrogatorio?")
	resp := fmt.Sprintf("El abogado decide: %v\n\nDescribe la respuesta de la corte.", action)
	sendAndPrintResponse(ctx, session, resp)

	// Continuar con la simulación
}

// askUser solicita una entrada del usuario.
func askUser(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		printStringFormatted(fmt.Sprintf("%v ", prompt))
		action, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error leyendo la entrada: %v\n", err)
		}
		action = strings.TrimSpace(action)
		if len(action) == 0 {
			continue
		}
		return action
	}
}

// sendAndPrintResponse envía un mensaje al modelo y muestra la respuesta.
func sendAndPrintResponse(ctx context.Context, session *genai.ChatSession, text string) {
	it := session.SendMessageStream(ctx, genai.Text(text))
	printRuneFormatted('\n')
	printRuneFormatted('\n')

	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			printStringFormatted("\n\nYou feel a jolt of electricity as you realize you're being unplugged from the matrix.\n\n")
			log.Printf("Error sending message: err=%v\n", err)

			var ae *apierror.APIError
			if errors.As(err, &ae) {
				log.Printf("ae.Reason(): %v\n", ae.Reason())
				log.Printf("ae.Details().Help.GetLinks(): %v\n", ae.Details().Help.GetLinks())
			}

			if s, ok := status.FromError(err); ok {
				log.Printf("s.Message: %v\n", s.Message())
				for _, d := range s.Proto().Details {
					log.Printf("- Details: %v\n", d)
				}
			}
			os.Exit(1)
		}
		for _, cand := range resp.Candidates {
			streamPartialResponse(cand.Content.Parts)
		}
	}
	printRuneFormatted('\n')
}

// streamPartialResponse imprime la respuesta parcial.
func streamPartialResponse(parts []genai.Part) {
	for _, part := range parts {
		printStringFormatted(fmt.Sprintf("%v", part))
	}
}

// printStringFormatted imprime el texto y lo formatea, con retrasos para efecto.
func printStringFormatted(text string) {
	for _, c := range text {
		printRuneFormatted(c)
	}
}

// printRuneFormatted imprime el carácter y lo formatea, con retrasos para efecto.
func printRuneFormatted(c rune) {
	switch c {
	case '.':
		fmt.Print(string(c))
		col++
		time.Sleep(sleepTime.sentence)
	case '\n':
		fmt.Print(string(c))
		col = 0
	case ' ':
		if col == 0 {
			// Do nothing.
		} else if col > 80 {
			fmt.Print("\n")
			col = 0
		} else {
			fmt.Print(string(c))
			col++
		}
	default:
		fmt.Print(string(c))
		col++
	}
	time.Sleep(sleepTime.character)
}
