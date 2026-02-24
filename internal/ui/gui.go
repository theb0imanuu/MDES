package ui

import (
	"fmt"
	"sort"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"mdes-cli/internal/engine"
)

type GUI struct {
	es              *engine.ExpertSystem
	app             fyne.App
	window          fyne.Window
	patientSymptoms []string
}

func NewGUI(es *engine.ExpertSystem) *GUI {
	a := app.New()
	w := a.NewWindow("Medical Diagnostic Expert System (MDES)")
	w.Resize(fyne.NewSize(800, 600))
	return &GUI{
		es:     es,
		app:    a,
		window: w,
	}
}

func (g *GUI) Run() {
	g.showWelcomeScreen()
	g.window.ShowAndRun()
}

func (g *GUI) showWelcomeScreen() {
	title := widget.NewLabelWithStyle("Medical Diagnostic Expert System", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	version := widget.NewLabelWithStyle("v2.0 - Large Knowledge Base Implementation", fyne.TextAlignCenter, fyne.TextStyle{Italic: true})

	instructions := widget.NewRichTextFromMarkdown(`
## Instructions
* Select the symptoms that the patient is currently experiencing.
* Be as accurate as possible for a reliable diagnosis.
* **Disclaimer:** This is a preliminary assessment tool only. Always consult a qualified healthcare provider.
	`)

	startButton := widget.NewButton("Start Diagnosis", func() {
		g.showSymptomSelectionScreen()
	})

	content := container.NewVBox(
		title,
		version,
		widget.NewSeparator(),
		instructions,
		layout.NewSpacer(),
		container.NewHBox(layout.NewSpacer(), startButton, layout.NewSpacer()),
		layout.NewSpacer(),
	)

	g.window.SetContent(container.NewPadded(content))
}

func (g *GUI) showSymptomSelectionScreen() {
	g.patientSymptoms = []string{} // Reset symptoms

	categories := g.getSymptomsByCategory()
	priorityOrder := []string{"General", "Pain", "Respiratory", "Cardiovascular", "Neurological", "Gastrointestinal", "Dermatological", "Infectious", "Viral", "Endocrine", "Autoimmune", "Oncological", "Psychiatric", "Psychological"}

	var items []fyne.CanvasObject
	title := widget.NewLabelWithStyle("Select Patient Symptoms", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	items = append(items, title, widget.NewSeparator())

	checkboxes := make(map[string]*widget.Check)

	for _, category := range priorityOrder {
		symptoms, exists := categories[category]
		if !exists || len(symptoms) == 0 {
			continue
		}

		catLabel := widget.NewLabelWithStyle(strings.ToUpper(category)+" SYMPTOMS", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
		items = append(items, catLabel)

		// Sort symptoms alphabetically within category
		sort.Slice(symptoms, func(i, j int) bool {
			return symptoms[i].Name < symptoms[j].Name
		})

		for _, sym := range symptoms {
			symID := sym.ID // capture for closure
			check := widget.NewCheck(fmt.Sprintf("%s (%s)", sym.Name, sym.Description), nil)
			checkboxes[symID] = check
			items = append(items, check)
		}
		items = append(items, widget.NewSeparator())
	}

	diagnoseButton := widget.NewButton("Run Diagnosis", func() {
		for symID, check := range checkboxes {
			if check.Checked {
				g.patientSymptoms = append(g.patientSymptoms, symID)
			}
		}
		g.showResultsScreen()
	})

	scroll := container.NewScroll(container.NewVBox(items...))

	content := container.NewBorder(
		nil,                                 // top
		container.NewPadded(diagnoseButton), // bottom
		nil,                                 // left
		nil,                                 // right
		scroll,                              // center
	)

	g.window.SetContent(content)
}

func (g *GUI) showResultsScreen() {
	diagnoses := g.es.Diagnose(g.patientSymptoms)

	var items []fyne.CanvasObject
	title := widget.NewLabelWithStyle("Diagnostic Results", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	items = append(items, title, widget.NewSeparator())

	if len(diagnoses) == 0 {
		items = append(items, widget.NewLabel("⚠️ No matching conditions found.\n\nPossible reasons:\n- Symptoms don't match known patterns.\n- Rare condition not in knowledge base.\n- Insufficient symptom information."))
	} else {
		limit := 5
		if len(diagnoses) < limit {
			limit = len(diagnoses)
		}

		for i := 0; i < limit; i++ {
			d := diagnoses[i]
			items = append(items, g.createDiagnosisCard(i+1, d))
		}

		if len(diagnoses) > 1 && (diagnoses[0].Probability-diagnoses[1].Probability) < 0.2 {
			items = append(items, widget.NewLabelWithStyle("⚠️ DIFFERENTIAL DIAGNOSIS ADVISED", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
		}
	}

	items = append(items, widget.NewSeparator(), widget.NewLabelWithStyle("DISCLAIMER: Decision support tool only. Final diagnosis requires clinical examination.", fyne.TextAlignCenter, fyne.TextStyle{Italic: true}))

	restartButton := widget.NewButton("Start New Diagnosis", func() {
		g.showWelcomeScreen()
	})

	scroll := container.NewScroll(container.NewVBox(items...))
	content := container.NewBorder(
		nil,
		container.NewPadded(restartButton),
		nil,
		nil,
		scroll,
	)

	g.window.SetContent(content)
}

func (g *GUI) createDiagnosisCard(rank int, d engine.DiagnosisResult) fyne.CanvasObject {
	header := widget.NewLabelWithStyle(fmt.Sprintf("%d. %s", rank, strings.ToUpper(d.Disease)), fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	probLabel := widget.NewLabel(fmt.Sprintf("Probability: %.1f%%", d.Probability*100))

	// Create a simple styled progress bar
	progress := widget.NewProgressBar()
	progress.SetValue(d.Probability)

	details := widget.NewLabel(fmt.Sprintf("Severity: %s\nMatch: %d/%d required, %d/%d optional", d.Severity, d.MatchedRequired, d.TotalRequired, d.MatchedOptional, d.TotalOptional))

	vbox := container.NewVBox(header, container.NewHBox(probLabel, layout.NewSpacer()), progress, details)

	if rank == 1 && len(d.Recommendations) > 0 {
		recText := "🏥 RECOMMENDATIONS:\n"
		for _, rec := range d.Recommendations {
			recText += "• " + rec + "\n"
		}
		vbox.Add(widget.NewLabel(recText))
	}

	return container.NewPadded(widget.NewCard("", "", vbox))
}

func (g *GUI) getSymptomsByCategory() map[string][]engine.Symptom {
	categories := make(map[string][]engine.Symptom)
	for _, sym := range g.es.Symptoms {
		categories[sym.Category] = append(categories[sym.Category], sym)
	}
	return categories
}
