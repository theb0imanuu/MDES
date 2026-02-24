package knowledge

import "mdes-cli/internal/engine"

func LoadKnowledgeBase(es *engine.ExpertSystem) {
	// Load all symptoms
	loadSymptoms(es)
	// Load all diseases
	loadDiseases(es)
	// Generate rules from diseases
	generateRules(es)
}

func loadSymptoms(es *engine.ExpertSystem) {
	symptoms := []engine.Symptom{
		{ID: "fever", Name: "Fever", Description: "Elevated body temperature >38°C", Category: "General", Severity: "Variable"},
		{ID: "high_fever", Name: "High Fever", Description: "Temperature >39°C", Category: "General", Severity: "Severe"},
		{ID: "chills", Name: "Chills", Description: "Feeling cold with shivering", Category: "General", Severity: "Moderate"},
		{ID: "sweating", Name: "Excessive Sweating", Description: "Abnormal perspiration", Category: "General", Severity: "Mild"},
		{ID: "fatigue", Name: "Fatigue", Description: "Extreme tiredness", Category: "General", Severity: "Moderate"},
		{ID: "headache", Name: "Headache", Description: "Pain in head", Category: "Pain", Severity: "Moderate"},
		{ID: "severe_headache", Name: "Severe Headache", Description: "Intense head pain", Category: "Pain", Severity: "Severe"},
		{ID: "chest_pain", Name: "Chest Pain", Description: "Pain in chest", Category: "Pain", Severity: "Severe"},
		{ID: "abdominal_pain", Name: "Abdominal Pain", Description: "Stomach pain", Category: "Pain", Severity: "Moderate"},
		{ID: "cough", Name: "Cough", Description: "Sudden air expulsion", Category: "Respiratory", Severity: "Moderate"},
		{ID: "dry_cough", Name: "Dry Cough", Description: "Cough without mucus", Category: "Respiratory", Severity: "Moderate"},
		{ID: "productive_cough", Name: "Productive Cough", Description: "Cough with mucus", Category: "Respiratory", Severity: "Moderate"},
		{ID: "shortness_of_breath", Name: "Shortness of Breath", Description: "Difficulty breathing", Category: "Respiratory", Severity: "Severe"},
		{ID: "wheezing", Name: "Wheezing", Description: "Whistling breath sound", Category: "Respiratory", Severity: "Moderate"},
		{ID: "nausea", Name: "Nausea", Description: "Feeling of sickness", Category: "Gastrointestinal", Severity: "Moderate"},
		{ID: "vomiting", Name: "Vomiting", Description: "Forceful stomach expulsion", Category: "Gastrointestinal", Severity: "Moderate"},
		{ID: "diarrhea", Name: "Diarrhea", Description: "Loose watery stools", Category: "Gastrointestinal", Severity: "Moderate"},
		{ID: "dizziness", Name: "Dizziness", Description: "Feeling lightheaded", Category: "Neurological", Severity: "Moderate"},
		{ID: "confusion", Name: "Confusion", Description: "Mental disorientation", Category: "Neurological", Severity: "Severe"},
		{ID: "seizures", Name: "Seizures", Description: "Uncontrolled brain activity", Category: "Neurological", Severity: "Critical"},
		{ID: "loss_of_consciousness", Name: "Loss of Consciousness", Description: "Fainting", Category: "Neurological", Severity: "Critical"},
		{ID: "palpitations", Name: "Palpitations", Description: "Irregular heartbeat", Category: "Cardiovascular", Severity: "Moderate"},
		{ID: "rapid_heartbeat", Name: "Rapid Heartbeat", Description: "Heart rate >100 bpm", Category: "Cardiovascular", Severity: "Moderate"},
		{ID: "rash", Name: "Rash", Description: "Skin eruption", Category: "Dermatological", Severity: "Moderate"},
		{ID: "loss_of_taste", Name: "Loss of Taste", Description: "Ageusia", Category: "Viral", Severity: "Moderate"},
		{ID: "loss_of_smell", Name: "Loss of Smell", Description: "Anosmia", Category: "Viral", Severity: "Moderate"},
		{ID: "body_aches", Name: "Body Aches", Description: "Generalized muscle pain", Category: "Viral", Severity: "Moderate"},
		{ID: "night_sweats", Name: "Night Sweats", Description: "Heavy sweating during sleep", Category: "General", Severity: "Moderate"},
		{ID: "weight_loss", Name: "Weight Loss", Description: "Unintentional reduction", Category: "General", Severity: "Moderate"},
		{ID: "loss_of_appetite", Name: "Loss of Appetite", Description: "Reduced desire to eat", Category: "General", Severity: "Mild"},
		{ID: "joint_pain", Name: "Joint Pain", Description: "Pain in joints", Category: "Pain", Severity: "Moderate"},
		{ID: "muscle_pain", Name: "Muscle Pain", Description: "Pain in muscles", Category: "Pain", Severity: "Moderate"},
		{ID: "sore_throat", Name: "Sore Throat", Description: "Pain in throat", Category: "Pain", Severity: "Moderate"},
		{ID: "runny_nose", Name: "Runny Nose", Description: "Nasal discharge", Category: "Respiratory", Severity: "Mild"},
		{ID: "stuffy_nose", Name: "Stuffy Nose", Description: "Nasal congestion", Category: "Respiratory", Severity: "Mild"},
		{ID: "constipation", Name: "Constipation", Description: "Difficulty passing stool", Category: "Gastrointestinal", Severity: "Mild"},
		{ID: "jaundice", Name: "Jaundice", Description: "Yellowing skin/eyes", Category: "Gastrointestinal", Severity: "Severe"},
		{ID: "dehydration", Name: "Dehydration", Description: "Excessive fluid loss", Category: "General", Severity: "Severe"},
		{ID: "lymphadenopathy", Name: "Swollen Lymph Nodes", Description: "Enlarged lymph glands", Category: "Infectious", Severity: "Moderate"},
		{ID: "bullseye_rash", Name: "Bullseye Rash", Description: "Circular expanding rash", Category: "Infectious", Severity: "Severe"},
		{ID: "paroxysmal_fever", Name: "Paroxysmal Fever", Description: "Periodic fever spikes", Category: "Infectious", Severity: "Severe"},
		{ID: "rose_spots", Name: "Rose Spots", Description: "Pink spots on trunk", Category: "Infectious", Severity: "Moderate"},
		{ID: "relative_bradycardia", Name: "Slow Pulse with Fever", Description: "Bradycardia despite fever", Category: "Infectious", Severity: "Moderate"},
		{ID: "hemoptysis", Name: "Coughing Blood", Description: "Blood in sputum", Category: "Respiratory", Severity: "Critical"},
		{ID: "cyanosis", Name: "Blue Lips/Fingertips", Description: "Lack of oxygen", Category: "Respiratory", Severity: "Critical"},
		{ID: "black_stool", Name: "Black Stool", Description: "Tarry black feces", Category: "Gastrointestinal", Severity: "Severe"},
		{ID: "edema", Name: "Swelling", Description: "Fluid retention", Category: "Cardiovascular", Severity: "Moderate"},
		{ID: "anxiety", Name: "Anxiety", Description: "Feeling of worry", Category: "Psychological", Severity: "Moderate"},
		{ID: "depression", Name: "Depression", Description: "Persistent sadness", Category: "Psychological", Severity: "Moderate"},
		{ID: "insomnia", Name: "Insomnia", Description: "Difficulty sleeping", Category: "Psychological", Severity: "Moderate"},
		{ID: "numbness", Name: "Numbness", Description: "Loss of sensation", Category: "Neurological", Severity: "Moderate"},
		{ID: "tingling", Name: "Tingling", Description: "Pins and needles", Category: "Neurological", Severity: "Mild"},
		{ID: "tremors", Name: "Tremors", Description: "Involuntary shaking", Category: "Neurological", Severity: "Moderate"},
		{ID: "slurred_speech", Name: "Slurred Speech", Description: "Difficulty speaking", Category: "Neurological", Severity: "Severe"},
		{ID: "back_pain", Name: "Back Pain", Description: "Pain in back", Category: "Pain", Severity: "Moderate"},
		{ID: "chest_tightness", Name: "Chest Tightness", Description: "Pressure in chest", Category: "Respiratory", Severity: "Moderate"},
		{ID: "rapid_breathing", Name: "Rapid Breathing", Description: "Increased breath rate", Category: "Respiratory", Severity: "Moderate"},
		{ID: "sneezing", Name: "Sneezing", Description: "Sudden air expulsion", Category: "Respiratory", Severity: "Mild"},
		{ID: "bloating", Name: "Bloating", Description: "Feeling of fullness", Category: "Gastrointestinal", Severity: "Mild"},
		{ID: "heartburn", Name: "Heartburn", Description: "Burning chest sensation", Category: "Gastrointestinal", Severity: "Mild"},
		{ID: "vertigo", Name: "Vertigo", Description: "Spinning sensation", Category: "Neurological", Severity: "Moderate"},
		{ID: "memory_loss", Name: "Memory Loss", Description: "Forgetfulness", Category: "Neurological", Severity: "Moderate"},
		{ID: "slow_heartbeat", Name: "Slow Heartbeat", Description: "Heart rate <60 bpm", Category: "Cardiovascular", Severity: "Moderate"},
		{ID: "high_blood_pressure", Name: "High Blood Pressure", Description: "Elevated BP", Category: "Cardiovascular", Severity: "Severe"},
		{ID: "itching", Name: "Itching", Description: "Pruritus", Category: "Dermatological", Severity: "Mild"},
		{ID: "skin_discoloration", Name: "Skin Discoloration", Description: "Abnormal skin color", Category: "Dermatological", Severity: "Moderate"},
		{ID: "bruising", Name: "Easy Bruising", Description: "Unexplained bruises", Category: "Dermatological", Severity: "Moderate"},
		{ID: "petechiae", Name: "Petechiae", Description: "Small red spots", Category: "Dermatological", Severity: "Severe"},
		{ID: "irritability", Name: "Irritability", Description: "Easily annoyed", Category: "Psychological", Severity: "Mild"},
	}

	for _, s := range symptoms {
		es.AddSymptom(s)
	}
}

func loadDiseases(es *engine.ExpertSystem) {
	diseases := []engine.Disease{
		{
			ID: "malaria", Name: "Malaria", Description: "Parasitic disease from mosquitoes", Category: "Infectious",
			Conditions: []string{"fever", "chills", "sweating"},
			Optional:   []string{"headache", "muscle_pain", "fatigue", "nausea", "vomiting", "paroxysmal_fever"},
			CF:         0.90, Exclusions: []string{"rash"}, Severity: "Severe",
			Recommendations: []string{"Immediate antimalarial medication", "Hospitalization if severe", "IV artesunate for severe cases"},
		},
		{
			ID: "typhoid", Name: "Typhoid Fever", Description: "Bacterial infection (Salmonella typhi)", Category: "Infectious",
			Conditions: []string{"fever", "abdominal_pain", "headache"},
			Optional:   []string{"rose_spots", "relative_bradycardia", "constipation", "diarrhea", "fatigue", "loss_of_appetite"},
			CF:         0.85, Exclusions: []string{"cough"}, Severity: "Severe",
			Recommendations: []string{"Antibiotic therapy", "Hydration", "Hospitalization for complications"},
		},
		{
			ID: "dengue", Name: "Dengue Fever", Description: "Mosquito-borne viral infection", Category: "Infectious",
			Conditions: []string{"high_fever", "severe_headache", "joint_pain", "muscle_pain"},
			Optional:   []string{"rash", "nausea", "vomiting", "mild_bleeding"},
			CF:         0.88, Exclusions: []string{"cough"}, Severity: "Severe",
			Recommendations: []string{"Fluid replacement therapy", "Rest", "Paracetamol for fever", "Avoid NSAIDs"},
		},
		{
			ID: "covid19", Name: "COVID-19", Description: "Coronavirus disease 2019", Category: "Viral",
			Conditions: []string{"fever", "dry_cough", "fatigue"},
			Optional:   []string{"loss_of_taste", "loss_of_smell", "shortness_of_breath", "body_aches", "sore_throat", "headache"},
			CF:         0.85, Exclusions: []string{}, Severity: "Severe",
			Recommendations: []string{"Isolation", "Supportive care", "Antivirals if indicated", "Monitor oxygen levels"},
		},
		{
			ID: "influenza", Name: "Influenza", Description: "Viral respiratory infection", Category: "Viral",
			Conditions: []string{"fever", "cough", "sore_throat", "runny_nose"},
			Optional:   []string{"body_aches", "headache", "chills", "fatigue", "weakness"},
			CF:         0.80, Exclusions: []string{"loss_of_taste", "loss_of_smell"}, Severity: "Moderate",
			Recommendations: []string{"Rest", "Hydration", "Antivirals within 48 hours", "Symptomatic treatment"},
		},
		{
			ID: "common_cold", Name: "Common Cold", Description: "Viral upper respiratory infection", Category: "Viral",
			Conditions: []string{"runny_nose", "stuffy_nose", "sneezing", "sore_throat"},
			Optional:   []string{"cough", "mild_fatigue", "mild_headache"},
			CF:         0.75, Exclusions: []string{"high_fever", "severe_headache"}, Severity: "Mild",
			Recommendations: []string{"Rest", "Hydration", "Over-the-counter symptom relief"},
		},
		{
			ID: "tuberculosis", Name: "Tuberculosis", Description: "Bacterial lung infection", Category: "Infectious",
			Conditions: []string{"productive_cough", "night_sweats", "weight_loss", "fever"},
			Optional:   []string{"hemoptysis", "chest_pain", "fatigue", "loss_of_appetite"},
			CF:         0.90, Exclusions: []string{"rash"}, Severity: "Critical",
			Recommendations: []string{"Long-term antibiotic therapy", "Directly observed therapy", "Respiratory isolation"},
		},
		{
			ID: "pneumonia", Name: "Pneumonia", Description: "Lung inflammation from infection", Category: "Respiratory",
			Conditions: []string{"fever", "productive_cough", "shortness_of_breath", "chest_pain"},
			Optional:   []string{"rapid_breathing", "cyanosis", "confusion", "fatigue", "sweating"},
			CF:         0.92, Exclusions: []string{"runny_nose", "sneezing"}, Severity: "Severe",
			Recommendations: []string{"Antibiotics for bacterial", "Oxygen therapy", "IV fluids", "Hospitalization"},
		},
		{
			ID: "asthma", Name: "Asthma", Description: "Chronic inflammatory airway disease", Category: "Respiratory",
			Conditions: []string{"wheezing", "shortness_of_breath", "chest_tightness", "cough"},
			Optional:   []string{"rapid_breathing", "difficulty_speaking", "anxiety"},
			CF:         0.88, Exclusions: []string{"fever"}, Severity: "Severe",
			Recommendations: []string{"Bronchodilators", "Inhaled corticosteroids", "Avoid triggers", "Action plan"},
		},
		{
			ID: "myocardial_infarction", Name: "Heart Attack", Description: "Acute coronary syndrome", Category: "Cardiovascular",
			Conditions: []string{"chest_pain", "shortness_of_breath", "sweating", "nausea"},
			Optional:   []string{"pain_radiating_to_arm", "jaw_pain", "back_pain", "anxiety", "lightheadedness"},
			CF:         0.95, Exclusions: []string{"fever"}, Severity: "Critical",
			Recommendations: []string{"Call emergency services", "Aspirin", "Nitroglycerin", "Angioplasty"},
		},
		{
			ID: "stroke", Name: "Stroke", Description: "Interruption of blood supply to brain", Category: "Neurological",
			Conditions: []string{"sudden_numbness", "confusion", "trouble_speaking", "vision_problems", "severe_headache", "trouble_walking"},
			Optional:   []string{"dizziness", "loss_of_balance", "facial_droop"},
			CF:         0.98, Exclusions: []string{"fever"}, Severity: "Critical",
			Recommendations: []string{"Emergency medical attention", "Thrombolytics if within window", "Rehabilitation"},
		},
		{
			ID: "meningitis", Name: "Meningitis", Description: "Inflammation of brain/spinal cord membranes", Category: "Neurological",
			Conditions: []string{"severe_headache", "fever", "stiff_neck"},
			Optional:   []string{"confusion", "seizures", "sensitivity_to_light", "nausea", "vomiting", "rash"},
			CF:         0.95, Exclusions: []string{}, Severity: "Critical",
			Recommendations: []string{"Emergency hospitalization", "IV antibiotics/antivirals", "Steroids", "Supportive care"},
		},
		{
			ID: "gastroenteritis", Name: "Gastroenteritis", Description: "Stomach and intestine inflammation", Category: "Gastrointestinal",
			Conditions: []string{"diarrhea", "nausea", "vomiting", "abdominal_pain"},
			Optional:   []string{"fever", "headache", "muscle_pain", "dehydration"},
			CF:         0.85, Exclusions: []string{"constipation"}, Severity: "Moderate",
			Recommendations: []string{"Oral rehydration", "Rest", "Bland diet", "Antiemetics if needed"},
		},
		{
			ID: "appendicitis", Name: "Appendicitis", Description: "Inflammation of the appendix", Category: "Gastrointestinal",
			Conditions: []string{"severe_abdominal_pain", "nausea", "vomiting", "fever"},
			Optional:   []string{"loss_of_appetite", "constipation", "diarrhea"},
			CF:         0.90, Exclusions: []string{"cough"}, Severity: "Critical",
			Recommendations: []string{"Emergency surgery", "IV antibiotics", "Pain management"},
		},
		{
			ID: "migraine", Name: "Migraine", Description: "Severe recurring headaches", Category: "Neurological",
			Conditions: []string{"severe_headache", "nausea", "sensitivity_to_light"},
			Optional:   []string{"aura", "visual_disturbances", "dizziness", "vomiting"},
			CF:         0.85, Exclusions: []string{"fever"}, Severity: "Moderate",
			Recommendations: []string{"Triptans", "Pain relievers", "Rest in dark room", "Trigger avoidance"},
		},
		{
			ID: "lyme_disease", Name: "Lyme Disease", Description: "Tick-borne bacterial infection", Category: "Infectious",
			Conditions: []string{"bullseye_rash", "fever", "fatigue"},
			Optional:   []string{"headache", "muscle_pain", "joint_pain", "lymphadenopathy"},
			CF:         0.90, Exclusions: []string{}, Severity: "Moderate",
			Recommendations: []string{"Antibiotics", "Tick removal if present", "Monitor for complications"},
		},
		{
			ID: "hepatitis_a", Name: "Hepatitis A", Description: "Viral liver infection", Category: "Infectious",
			Conditions: []string{"jaundice", "fatigue", "abdominal_pain", "loss_of_appetite"},
			Optional:   []string{"nausea", "vomiting", "fever", "dark_urine"},
			CF:         0.85, Exclusions: []string{"cough"}, Severity: "Moderate",
			Recommendations: []string{"Rest", "Hydration", "Avoid alcohol", "Vaccination for contacts"},
		},
		{
			ID: "diabetes_type2", Name: "Type 2 Diabetes", Description: "Insulin resistance diabetes", Category: "Endocrine",
			Conditions: []string{"increased_thirst", "frequent_urination", "fatigue", "blurred_vision"},
			Optional:   []string{"slow_healing", "frequent_infections", "numbness"},
			CF:         0.85, Exclusions: []string{"fever"}, Severity: "Severe",
			Recommendations: []string{"Metformin", "Lifestyle changes", "Blood sugar monitoring", "Regular screening"},
		},
		{
			ID: "hypertension", Name: "Hypertension", Description: "High blood pressure", Category: "Cardiovascular",
			Conditions: []string{"high_blood_pressure"},
			Optional:   []string{"headache", "dizziness", "chest_pain", "shortness_of_breath"},
			CF:         0.75, Exclusions: []string{"fever"}, Severity: "Severe",
			Recommendations: []string{"Lifestyle modifications", "Antihypertensive medications", "Regular monitoring"},
		},
		{
			ID: "anxiety_disorder", Name: "Generalized Anxiety Disorder", Description: "Chronic anxiety condition", Category: "Psychiatric",
			Conditions: []string{"anxiety", "restlessness", "fatigue", "difficulty_concentrating"},
			Optional:   []string{"irritability", "muscle_tension", "sleep_disturbance"},
			CF:         0.75, Exclusions: []string{"fever"}, Severity: "Moderate",
			Recommendations: []string{"Cognitive behavioral therapy", "SSRIs", "Relaxation techniques", "Exercise"},
		},
		{
			ID: "depression", Name: "Major Depressive Disorder", Description: "Clinical depression", Category: "Psychiatric",
			Conditions: []string{"depression", "loss_of_interest", "sleep_changes", "fatigue"},
			Optional:   []string{"weight_changes", "difficulty_concentrating", "feelings_of_worthlessness"},
			CF:         0.80, Exclusions: []string{"fever", "high_fever"}, Severity: "Moderate",
			Recommendations: []string{"Psychotherapy", "Antidepressants", "Lifestyle modifications", "Crisis support"},
		},
		{
			ID: "epilepsy", Name: "Epilepsy", Description: "Seizure disorder", Category: "Neurological",
			Conditions: []string{"seizures", "temporary_confusion", "staring_spells"},
			Optional:   []string{"uncontrollable_jerking", "loss_of_consciousness", "fear"},
			CF:         0.90, Exclusions: []string{"fever"}, Severity: "Moderate",
			Recommendations: []string{"Antiepileptic drugs", "Ketogenic diet", "Vagus nerve stimulation"},
		},
		{
			ID: "hiv_initial", Name: "Acute HIV Infection", Description: "Early stage HIV infection", Category: "Infectious",
			Conditions: []string{"fever", "sore_throat", "rash_trunk", "lymphadenopathy"},
			Optional:   []string{"muscle_pain", "joint_pain", "headache", "night_sweats"},
			CF:         0.75, Exclusions: []string{}, Severity: "Severe",
			Recommendations: []string{"HIV testing", "Antiretroviral therapy", "Counseling"},
		},
		{
			ID: "cholera", Name: "Cholera", Description: "Bacterial intestinal infection", Category: "Infectious",
			Conditions: []string{"severe_watery_diarrhea", "vomiting", "dehydration"},
			Optional:   []string{"rapid_heartbeat", "low_blood_pressure", "muscle_cramps"},
			CF:         0.95, Exclusions: []string{"fever"}, Severity: "Critical",
			Recommendations: []string{"Immediate rehydration", "Oral rehydration salts", "IV fluids", "Antibiotics"},
		},
	}

	for _, d := range diseases {
		es.AddDisease(d)
	}
}

func generateRules(es *engine.ExpertSystem) {
	for id, disease := range es.Diseases {
		rule := engine.Rule{
			Disease:    id,
			Conditions: disease.Conditions,
			Optional:   disease.Optional,
			Exclusions: disease.Exclusions,
			CF:         disease.CF,
			Weight:     make(map[string]float64),
		}

		for _, cond := range disease.Conditions {
			rule.Weight[cond] = 1.0
		}
		for _, opt := range disease.Optional {
			rule.Weight[opt] = 0.5
		}

		es.AddRule(rule)
	}
}
