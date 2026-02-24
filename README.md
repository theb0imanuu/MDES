# Medical Diagnostic Expert System (MDES)

Implementation for Expert Systems — Assignment 3
Course: Expert Systems (CAT 1)
Assignment: Comprehensive Design of the Medical Diagnostic Expert System
Implementation Language: Go
Knowledge Base Scale: 50+ diseases, 100+ symptoms

## 1. Relevance to Assignment 3

This implementation directly realizes the theoretical design presented in Assignment 3, transforming the conceptual Medical Diagnostic Expert System into a functional command-line application. The implementation addresses the core problem identified in Assignment 1—delayed and inaccurate diagnosis in rural healthcare facilities—by providing healthcare workers with an accessible, portable diagnostic tool that requires no internet connectivity or expensive infrastructure.

### Alignment with Assignment 3 Requirements

| Assignment 3 Component                    | Implementation Status                                           |
| :---------------------------------------- | :-------------------------------------------------------------- |
| **Rule-Based Knowledge Representation**   | ✓ Production rules with IF-THEN structure and Certainty Factors |
| **Forward Chaining Inference**            | ✓ Data-driven reasoning engine implemented                      |
| **Certainty Factor Model**                | ✓ CF calculation with combined probability formulas             |
| **Explanation Facility**                  | ✓ Transparent reasoning output for all diagnoses                |
| **User Interface for Healthcare Workers** | ✓ Simplified CLI optimized for low-resource settings            |
| **Database of Diseases and Symptoms**     | ✓ Structured knowledge base with 50+ conditions                 |

The system specifically targets the target users identified in Assignment 3: nurses, clinical officers, and community health workers who require immediate diagnostic assistance in settings with limited physician availability.

## 2. Prerequisites

Before running the application, ensure you have the following installed:

- **Go Programming Language** (version 1.21 or higher)
  - Download from: https://golang.org/dl/
  - Verify installation: `go version`
- **Git** (for cloning the repository)
  - Download from: https://git-scm.com/downloads
- **Terminal/Command Line Access**
  - Windows: Command Prompt or PowerShell
  - macOS/Linux: Terminal

## 3. Installation and Setup

**Step 1: Clone the Repository**

```bash
git clone https://github.com/theb0imanuu/MDES.git
cd MDES
```

**Step 2: Verify Go Module**
The project uses Go modules for dependency management. Initialize if necessary:

```bash
go mod tidy
```

**Step 3: Build the Application**
Compile the source code into an executable binary:

```bash
go build -o mdes ./cmd/mdes
```

On Windows:

```cmd
go build -o mdes.exe ./cmd/mdes
```

This creates an executable file named `mdes` (or `mdes.exe` on Windows) in your current directory.

**Step 4: (Optional) Install to System PATH**
For global access from any directory:

- **Linux/macOS:**
  ```bash
  sudo mv mdes /usr/local/bin/
  ```
- **Windows:**
  Move `mdes.exe` to `C:\Windows\System32\` or add the build directory to your PATH environment variable.

## 4. Running the Application

The MDES CLI supports multiple execution modes to accommodate different clinical scenarios.

### 4.1 Interactive Diagnostic Mode (Default)

This mode guides healthcare workers through a systematic symptom interview, organized by physiological categories. This is the primary mode for clinical use.

**Command:**

```bash
./mdes
```

or explicitly:

```bash
./mdes -i
```

**Operational Flow:**

1. The system presents symptoms categorized by type (General, Pain, Respiratory, Cardiovascular, etc.)
2. The user responds to each symptom query:
   - `y` — Patient presents this symptom
   - `n` — Patient does not present this symptom
   - `s` — Skip remaining symptoms in current category
3. The inference engine applies forward chaining to match symptoms against the rule base
4. Results display ranked diagnoses with Certainty Factors, severity assessments, and clinical recommendations

**Example Session:**

```plain
============================================================
  MEDICAL DIAGNOSTIC EXPERT SYSTEM (MDES) v2.0
============================================================

【GENERAL SYMPTOMS】
----------------------------------------
Do you have Fever? (Elevated body temperature >38°C) [y/n/s]: y
  ✓ Added: Fever
Do you have Chills? (Feeling cold with shivering) [y/n/s]: y
  ✓ Added: Chills

[... additional symptoms ...]

============================================================
  DIAGNOSTIC RESULTS
============================================================

1. MALARIA
   Probability: 85.0% [████████████████░░░░]
   Severity: Severe
   Match: 3/3 required, 2/5 optional

   🏥 RECOMMENDATIONS:
   • Immediate antimalarial medication
   • Hospitalization if severe
   • IV artesunate for severe cases
```

### 4.2 Batch Diagnostic Mode

For rapid assessment when symptom data is already collected, or for testing specific clinical scenarios.

**Command:**

```bash
./mdes -b "symptom1,symptom2,symptom3,..."
```

**Example:**

```bash
./mdes -b "fever,dry_cough,fatigue,loss_of_taste,shortness_of_breath"
```

**Output:**

```plain
Analyzing symptoms: [fever dry_cough fatigue loss_of_taste shortness_of_breath]

Top Diagnosis: COVID-19 (85.0%)

Other possibilities:
  2. Influenza (45.2%)
  3. Pneumonia (38.7%)
```

**Clinical Application:** This mode supports rapid screening during outbreak investigations or when integrating with electronic health record systems.

### 4.3 Knowledge Base Statistics

Verify system capabilities and knowledge base coverage.

**Command:**

```bash
./mdes -s
```

**Sample Output:**

```plain
KNOWLEDGE BASE STATISTICS
----------------------------------------
Total Symptoms: 100
Total Diseases: 50
Total Rules: 50

Diseases by Category:
  • Infectious: 12
  • Respiratory: 8
  • Cardiovascular: 6
  • Neurological: 7
  • Gastrointestinal: 5
  • Viral: 4
  • Endocrine: 4
  • Psychiatric: 2
  • Autoimmune: 2
```

### 4.4 Symptom Reference List

Display all recognizable symptoms with their identifiers for batch mode usage.

**Command:**

```bash
./mdes -l
```

**Use Case:** Useful for training healthcare workers on system capabilities or preparing symptom sets for batch processing.

## 5. Clinical Usage Guidelines

### For Healthcare Workers in Rural Facilities

- **Patient Assessment**: Conduct standard clinical examination before using the system
- **Symptom Entry**: Input all observable and reported symptoms accurately
- **Interpretation**: Consider the top 3 diagnoses when probabilities are close (< 20% difference)
- **Action**: Follow recommended actions for the highest probability diagnosis
- **Referral**: Use the severity indicator to determine urgency of hospital referral

> [!CAUTION]
> **Important Limitations (Per Assignment 3 Constraints)**
>
> ⚠️ This system is not a replacement for qualified medical diagnosis. It provides decision support for healthcare workers in resource-limited settings. Final diagnosis requires clinical judgment and, where available, laboratory confirmation.
>
> ⚠️ The system scope is limited to the 50+ predefined diseases in the knowledge base. Rare conditions or complex multi-system presentations may not be accurately diagnosed.

## 6. Validation Against Assignment 3 Design

### 6.1 Rule Base Verification

The implemented rules follow the exact structure specified in Assignment 3, Section 4.4:

**Assignment 3 Specification:**

```plain
Rule R1 (Malaria)
IF Fever AND Chills AND Sweating AND Headache
THEN Malaria
CF = 0.85
```

**Implementation:**

```go
{
    ID: "malaria",
    Conditions: []string{"fever", "chills", "sweating"},
    Optional:   []string{"headache", "muscle_pain", "fatigue"},
    CF: 0.90,
    Severity: "Severe"
}
```

The implementation extends the Assignment 3 specification by:

- Adding optional symptoms for increased diagnostic granularity
- Implementing exclusion criteria (e.g., rash excludes malaria)
- Providing treatment recommendations per diagnosis

### 6.2 Certainty Factor Calculation

Per Assignment 3, Section 6, the system implements the CF combination formula:

```plain
CF_combined = CF1 + CF2(1 – CF1)
```

This is applied when multiple rules support the same diagnosis, allowing probabilistic ranking of differential diagnoses.

### 6.3 Explanation Facility

As required by Assignment 3, Section 7, the system provides:

- **WHY explanations**: Which symptoms triggered the diagnosis
- **HOW explanations**: Calculation method showing match ratios and certainty factors
- **Transparency**: Complete visibility into rule firing and probability calculation

## 7. Testing Scenarios

Validate system performance using these documented test cases from Assignment 3:

**Test Case 1: Classic Malaria Presentation**

```bash
./mdes -b "fever,chills,sweating,headache"
```

_Expected:_ Malaria with high probability (>80%)

**Test Case 2: Typhoid Fever**

```bash
./mdes -b "fever,abdominal_pain,diarrhea,fatigue"
```

_Expected:_ Typhoid Fever with moderate-high probability

**Test Case 3: COVID-19 Differentiation**

```bash
./mdes -b "fever,dry_cough,fatigue,loss_of_taste,loss_of_smell"
```

_Expected:_ COVID-19 ranked above Influenza due to specific symptoms

**Test Case 4: Pneumonia**

```bash
./mdes -b "fever,productive_cough,chest_pain,shortness_of_breath"
```

_Expected:_ Pneumonia with high certainty factor

## 8. Troubleshooting

| Issue                           | Solution                                                                                     |
| :------------------------------ | :------------------------------------------------------------------------------------------- |
| `command not found: mdes`       | Ensure the binary is in your PATH or use `./mdes` from the build directory                   |
| `go: module not found`          | Run `go mod init github.com/theb0imanuu/MDES` then `go mod tidy`                             |
| Permission denied (Linux/macOS) | Run `chmod +x mdes` to make executable                                                       |
| Batch mode returns no diagnosis | Check symptom spelling against `./mdes -l` output; use comma-separated values without spaces |

## 9. Academic Integrity Statement

This implementation is submitted as part of **Expert Systems — Assignment 3: Comprehensive Design of the Medical Diagnostic Expert System**. The code realizes the theoretical design specifications provided in the assignment documentation, demonstrating practical application of:

- Knowledge representation schemes (Section 4)
- Inference engine algorithms (Section 5)
- Uncertainty handling (Section 6)
- Explanation facilities (Section 7)

The implementation choices (Go language, CLI interface) prioritize deployment feasibility in low-resource healthcare settings identified in Assignment 1, while maintaining fidelity to the assignment's architectural requirements.

## 10. References

- Assignment 3 Documentation: Comprehensive Design of the Medical Diagnostic Expert System
- Assignment 1 Problem Statement: Delayed and Inaccurate Diagnosis in Rural Healthcare
- Assignment 2 Requirements: System Requirements and Knowledge Acquisition
- Go Programming Language Documentation: https://golang.org/doc/

