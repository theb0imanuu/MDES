package engine

type ExpertSystem struct {
	Symptoms map[string]Symptom
	Diseases map[string]Disease
	Rules    []Rule
}

func NewExpertSystem() *ExpertSystem {
	return &ExpertSystem{
		Symptoms: make(map[string]Symptom),
		Diseases: make(map[string]Disease),
	}
}

func (es *ExpertSystem) AddSymptom(s Symptom) {
	es.Symptoms[s.ID] = s
}

func (es *ExpertSystem) AddDisease(d Disease) {
	es.Diseases[d.ID] = d
}

func (es *ExpertSystem) AddRule(r Rule) {
	es.Rules = append(es.Rules, r)
}
