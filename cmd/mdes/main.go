package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"mdes-cli/internal/engine"
	"mdes-cli/internal/knowledge"
	"mdes-cli/internal/ui"
)

var (
	version   = "dev"
	buildTime = "unknown"
)

func main() {
	var (
		interactive = flag.Bool("i", false, "Interactive diagnosis mode")
		batch       = flag.String("b", "", "Batch mode with comma-separated symptoms")
		list        = flag.Bool("l", false, "List all symptoms")
		stats       = flag.Bool("s", false, "Show knowledge base statistics")
		help        = flag.Bool("h", false, "Show help")
		versionFlag = flag.Bool("v", false, "Show version")
	)

	flag.Parse()

	if *versionFlag {
		fmt.Printf("MDES v%s (built %s)\n", version, buildTime)
		os.Exit(0)
	}

	if *help {
		printHelp()
		os.Exit(0)
	}

	// Initialize expert system
	es := engine.NewExpertSystem()
	knowledge.LoadKnowledgeBase(es)

	switch {
	case *interactive:
		ui := ui.NewInteractiveUI(es)
		ui.Run()
	case *stats:
		printStats(es)
	case *list:
		listSymptoms(es)
	case *batch != "":
		runBatchMode(es, *batch)
	default:
		// Default to GUI mode
		gui := ui.NewGUI(es)
		gui.Run()
	}
}

func printHelp() {
	help := `
Medical Diagnostic Expert System (MDES) - CLI

USAGE:
  mdes [OPTIONS]

OPTIONS:
  -i          Interactive CLI diagnosis mode
  -b SYMPTOMS Batch mode with comma-separated symptoms
  -l          List all available symptoms
  -s          Show knowledge base statistics
  -v          Show version
  -h          Show this help message

EXAMPLES:
  # Launch GUI
  mdes

  # Interactive CLI mode
  mdes -i

  # Batch mode with symptoms
  mdes -b "fever,cough,fatigue,loss_of_taste"

  # Show stats
  mdes -s

For more information: https://github.com/yourusername/mdes-cli
`
	fmt.Println(help)
}

func printStats(es *engine.ExpertSystem) {
	fmt.Println("KNOWLEDGE BASE STATISTICS")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("Total Symptoms: %d\n", len(es.Symptoms))
	fmt.Printf("Total Diseases: %d\n", len(es.Diseases))
	fmt.Printf("Total Rules: %d\n", len(es.Rules))

	categories := make(map[string]int)
	for _, d := range es.Diseases {
		categories[d.Category]++
	}
	fmt.Println("\nDiseases by Category:")
	for cat, count := range categories {
		fmt.Printf("  • %s: %d\n", cat, count)
	}
}

func listSymptoms(es *engine.ExpertSystem) {
	fmt.Println("AVAILABLE SYMPTOMS")
	fmt.Println(strings.Repeat("=", 60))

	categories := make(map[string][]engine.Symptom)
	for _, sym := range es.Symptoms {
		categories[sym.Category] = append(categories[sym.Category], sym)
	}

	for cat, syms := range categories {
		fmt.Printf("\n%s:\n", strings.ToUpper(cat))
		for _, s := range syms {
			fmt.Printf("  %-20s %s\n", s.ID, s.Name)
		}
	}
}

func runBatchMode(es *engine.ExpertSystem, symptoms string) {
	symptomList := strings.Split(symptoms, ",")
	for i := range symptomList {
		symptomList[i] = strings.TrimSpace(symptomList[i])
	}

	fmt.Printf("Analyzing symptoms: %v\n", symptomList)

	results := es.Diagnose(symptomList)

	if len(results) == 0 {
		fmt.Println("No diagnosis possible with given symptoms.")
		return
	}

	fmt.Printf("\nTop Diagnosis: %s (%.1f%%)\n", results[0].Disease, results[0].Probability*100)

	if len(results) > 1 {
		fmt.Println("\nOther possibilities:")
		for i := 1; i < len(results) && i < 3; i++ {
			fmt.Printf("  %d. %s (%.1f%%)\n", i+1, results[i].Disease, results[i].Probability*100)
		}
	}
}
