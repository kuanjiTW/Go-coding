package main

// References:
// - https://pkg.go.dev/github.com/firebase/genkit/go

import (
    "context"
    "fmt"
    "log"
    "time"
    "encoding/json"

    "github.com/firebase/genkit/go/ai"
    "github.com/firebase/genkit/go/genkit"
    "github.com/firebase/genkit/go/plugins/googlegenai"
    "github.com/firebase/genkit/go/plugins/ollama"
    "google.golang.org/genai"
)

// Define some model names as we may use different models
const (
    GoogleGemini2_5Flash = "googleai/gemini-2.5-flash"
    OllamaGptOss = "ollama/gpt-oss:latest"
    OllamaLlama3 = "ollama/llama3:latest"
    OllamaLlama3_1 = "ollama/llama3.1:8b"
)

// The is the name of the model to be used
//var model string = GoogleGemini2_5Flash
//var model string = OllamaGptOss
//var model string = OllamaLlama3
var model string = OllamaLlama3_1

// Temperature controls the randomness of text that is generated
// by LLMs during inference.
// - Low temperature: The lower temperature value helps the LLM
//                    to produce more coherent and consistent
//                    text and avoid irrelevant responses
// - High temperature: a high temperature is preferable for creative
//                     outputs or creative tasks such as creative
//                     writing or concept brainstorming
// (https://www.ibm.com/think/topics/llm-temperature)
var temperature float32 = 0.7

// generation configuration, different backends may require
// different structs
var generate_config any

var ch1 chan bool = make(chan bool, 1)
var ch2 chan bool = make(chan bool)

func progress(msg string) {
    fmt.Print(msg)
    for len(ch1) == 0 {
        fmt.Print(".")
        time.Sleep(500 * time.Millisecond)
    }
    fmt.Println()
    <-ch1     // take the item from ch1
    ch2<-true // notify the main thread
}

func main() {
    // A Context carries a deadline, a cancellation signal, and
    // other values across API boundaries.
    //
    // Background returns a non-nil, empty Context. It is never
    // canceled, has no values, and has no deadline
    // (https://pkg.go.dev/context#Background)
    ctx := context.Background()

    // There are plugins for other AI service providers
    google_plugin := googlegenai.GoogleAI{}
    ollama_plugin := ollama.Ollama{
        ServerAddress: "http://127.0.0.1:11434",
    }

    // Initialize Genkit with the two plugins
    //
    // func Init(ctx context.Context, opts ...GenkitOption) *Genkit
    //
    // (https://pkg.go.dev/github.com/firebase/genkit/go@v1.2.0/genkit#Init)
    //
    // See https://pkg.go.dev/github.com/firebase/genkit/go@v1.2.0/genkit#GenkitOption
    // for more GenkitOptions
    g := genkit.Init(ctx,
        genkit.WithPlugins(
            &google_plugin,
            &ollama_plugin,
        ),
    )

    // Define models in ollama so that we can use ai.withModelName to find them
    ollama_plugin.DefineModel(g, ollama.ModelDefinition{Name: "gpt-oss:latest"}, nil)
    ollama_plugin.DefineModel(g, ollama.ModelDefinition{Name: "llama3:latest"}, nil)
    ollama_plugin.DefineModel(g, ollama.ModelDefinition{Name: "llama3.1:8b"}, nil)

    // Setup configuration
    switch model {
    case GoogleGemini2_5Flash:
        generate_config  = &genai.GenerateContentConfig{
            Temperature: genai.Ptr[float32](temperature),
        }
    case OllamaGptOss, OllamaLlama3:
        generate_config = &ai.GenerationCommonConfig{
            Temperature: float64(temperature),
        }
    }

    fmt.Printf("Selected Model: %s\n", model)

    // Run a simple prompt
    go progress("Running SimplePrompt")
    SimplePrompt(ctx, g)

    fmt.Printf("\n\n")

    // Run a prompt with a complex output type
    go progress("Running PromptWithComplexOutputType")
    PromptWithComplexOutputType(ctx, g)
}

func SimplePrompt(ctx context.Context, g *genkit.Genkit) {
    // Generate content using a prompt
    prompt := "Explain why learning the Go programming language with top 3 reasons."
    //prompt := "Explain the concept of recursion in one sentence"
    fmt.Println("== Prompt ==")
    fmt.Println(prompt)

    // func Generate(ctx context.Context, g *Genkit, opts ...ai.GenerateOption) (*ai.ModelResponse, error)
    resp, err := genkit.Generate(ctx, g,
        ai.WithModelName(model),
        ai.WithPrompt(prompt),
        ai.WithConfig(generate_config),
    )
    ch1<-true // notify the progress thread
    <-ch2     // wait for notification
    if err != nil {
        log.Fatalf("Generation failed: %v", err)
    }

    // Print the output
    fmt.Println("== Response ==")
    fmt.Println(resp.Text())
}

func PromptWithComplexOutputType(ctx context.Context, g *genkit.Genkit) {
    type countryData struct {
        Name      string `json:"name"`
        Language  string `json:"language"`
        Habitants int    `json:"habitants"`
    }

    type countries struct {
        Countries []countryData `json:"countries"`
    }

    // Define prompt with output api.
    prompt := genkit.DefinePrompt(
        g, "PromptWithComplexOutputType",
        ai.WithModelName(model),
        // WithOutputType sets the schema and format of the output based
        // on the value provided.
        // Notes:
        // - Not all models support constrained output. For example,
        //   llama3 does not support constrained output.
        // - You may also constrain the inputs by ai.WithInputType.
        ai.WithOutputType(countries{}),
        ai.WithConfig(generate_config),
        // System prompts set the overall framework
        ai.WithSystem("You are a geography teacher. When asked a question about geography."),
        // User prompts are task-specific instructions
        ai.WithPrompt("Give me the 5 biggest countries in the world by habitants and language."),
    )

    // Run the prompt
    resp, err := prompt.Execute(ctx)
    ch1<-true // notify the progress thread
    <-ch2     // wait for notification
    if err != nil {
        log.Fatal(err)
    }

    var c countries
    // Output unmarshals structured JSON output into the provided
    // struct pointer.
    if err = resp.Output(&c); err != nil {
        log.Fatal(err)
    }

    // func MarshalIndent(v any, prefix, indent string) ([]byte, error)
    // - Return the JSON encoding of v.
    // (https://pkg.go.dev/encoding/json#MarshalIndent)
    pretty, err := json.MarshalIndent(c, "", "  ")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(pretty))
}
