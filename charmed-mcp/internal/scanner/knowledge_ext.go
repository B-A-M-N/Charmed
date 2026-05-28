// Package scanner — extended knowledge types for primitives, patterns, and constraints-v2.
package scanner

import (
	"os"
	"path/filepath"
)

// Primitive is a behavioral primitive from primitives.yaml
type Primitive struct {
	ID          string   `yaml:"id" json:"id"`
	Category    string   `yaml:"category" json:"category"`
	Description string   `yaml:"description" json:"description"`
	Signals     []string `yaml:"signals" json:"signals"`
	Constraints []string `yaml:"constraints" json:"constraints"`
	Implements  []string `yaml:"implements" json:"implements"`
}

// PrimitivesDoc is the top-level structure of primitives.yaml
type PrimitivesDoc struct {
	Version    int         `yaml:"version" json:"version"`
	Primitives []Primitive `yaml:"primitives" json:"primitives"`
}

// Pattern is an empirical pattern instance
type Pattern struct {
	ID                     string            `yaml:"id" json:"id"`
	Name                   string            `yaml:"name" json:"name"`
	Category               string            `yaml:"category" json:"category"`
	Primitives             []string          `yaml:"primitives" json:"primitives"`
	TopologySignature      string            `yaml:"topology_signature" json:"topology_signature"`
	Confidence             float64           `yaml:"confidence" json:"confidence"`
	ConfidenceFactors      PatternConfidence `yaml:"confidence_factors" json:"confidence_factors"`
	AppliesTo              []string          `yaml:"applies_to" json:"applies_to"`
	GeneralizesTo          []string          `yaml:"generalizes_to" json:"generalizes_to"`
	ConflictsWith          []string          `yaml:"conflicts_with" json:"conflicts_with"`
	SynthesizedConstraints []ConstraintRef   `yaml:"synthesized_constraints" json:"synthesized_constraints"`
	FirstObserved          string            `yaml:"first_observed" json:"first_observed"`
	LastUpdated            string            `yaml:"last_updated" json:"last_updated"`
	DerivedFrom            PatternEvidence   `yaml:"derived_from" json:"derived_from"`
	Structure              PatternStructure  `yaml:"structure" json:"structure"`
}

// PatternConfidence holds confidence scoring factors.
type PatternConfidence struct {
	ObservationCount int     `yaml:"observation_count" json:"observation_count"`
	ConsistencyScore float64 `yaml:"consistency_score" json:"consistency_score"`
	ViolationCount   int     `yaml:"violation_count" json:"violation_count"`
	FixCount         int     `yaml:"fix_count" json:"fix_count"`
}

// PatternEvidence holds evidence references for a pattern.
type PatternEvidence struct {
	Observations []ObservationRef `yaml:"observations" json:"observations"`
	Violations   []ViolationRef   `yaml:"violations" json:"violations"`
	Fixes        []FixRef         `yaml:"fixes" json:"fixes"`
}

// ObservationRef is a reference to an observed implementation.
type ObservationRef struct {
	Repo        string `yaml:"repo" json:"repo"`
	File        string `yaml:"file" json:"file"`
	Line        int    `yaml:"line" json:"line"`
	EvidenceID  string `yaml:"evidence_id" json:"evidence_id"`
	ExtractedAt string `yaml:"extracted_at" json:"extracted_at"`
}

// ViolationRef is a reference to a known violation.
type ViolationRef struct {
	Repo       string `yaml:"repo" json:"repo"`
	IssueURL   string `yaml:"issue_url" json:"issue_url"`
	Symptom    string `yaml:"symptom" json:"symptom"`
	RootCause  string `yaml:"root_cause" json:"root_cause"`
	EvidenceID string `yaml:"evidence_id" json:"evidence_id"`
}

// FixRef is a reference to a commit that applied the pattern.
type FixRef struct {
	Repo             string `yaml:"repo" json:"repo"`
	CommitSHA        string `yaml:"commit_sha" json:"commit_sha"`
	PRURL            string `yaml:"pr_url" json:"pr_url"`
	BeforeEvidenceID string `yaml:"before_evidence_id" json:"before_evidence_id"`
	AfterEvidenceID  string `yaml:"after_evidence_id" json:"after_evidence_id"`
	Improvement      string `yaml:"improvement" json:"improvement"`
}

// ConstraintRef is a synthesized constraint reference.
type ConstraintRef struct {
	ConstraintID       string  `yaml:"constraint_id" json:"constraint_id"`
	Confidence         float64 `yaml:"confidence" json:"confidence"`
	DerivedFromPattern bool    `yaml:"derived_from_pattern" json:"derived_from_pattern"`
}

// PatternStructure describes the topology of a pattern.
type PatternStructure struct {
	Type       string        `yaml:"type" json:"type"`
	Nodes      []PatternNode `yaml:"nodes" json:"nodes"`
	Edges      []PatternEdge `yaml:"edges" json:"edges"`
	Invariants []Invariant   `yaml:"invariants" json:"invariants"`
}

// PatternNode is a primitive usage within a pattern.
type PatternNode struct {
	PrimitiveID string `yaml:"primitive_id" json:"primitive_id"`
	Role        string `yaml:"role" json:"role"`
	Required    bool   `yaml:"required" json:"required"`
	Cardinality string `yaml:"cardinality" json:"cardinality"`
}

// PatternEdge is a relationship between primitives in a pattern.
type PatternEdge struct {
	From     string `yaml:"from" json:"from"`
	To       string `yaml:"to" json:"to"`
	Type     string `yaml:"type" json:"type"`
	Required bool   `yaml:"required" json:"required"`
	Timing   string `yaml:"timing" json:"timing"`
}

// Invariant is an invariant that the pattern must uphold.
type Invariant struct {
	Description string   `yaml:"description" json:"description"`
	ViolatedBy  []string `yaml:"violated_by" json:"violated_by"`
	EnforcedBy  []string `yaml:"enforced_by" json:"enforced_by"`
}

// ConstraintV2 is the enhanced constraint from constraints-v2.yaml
type ConstraintV2 struct {
	ID                 string               `yaml:"id" json:"id"`
	Severity           string               `yaml:"severity" json:"severity"`
	Category           string               `yaml:"category" json:"category"`
	Type               string               `yaml:"type" json:"type"` // "canonical" | "empirical"
	Confidence         float64              `yaml:"confidence" json:"confidence"`
	Description        string               `yaml:"description" json:"description"`
	Rationale          string               `yaml:"rationale" json:"rationale"`
	PrimitivesViolated []string             `yaml:"primitives_violated" json:"primitives_violated"`
	Message            string               `yaml:"message" json:"message"`
	Fix                ConstraintV2Fix      `yaml:"fix" json:"fix"`
	Evidence           ConstraintV2Evidence `yaml:"evidence" json:"evidence"`
	Match              ConstraintV2Match    `yaml:"match" json:"match"`
}

// ConstraintV2Fix describes how to fix a constraint violation.
type ConstraintV2Fix struct {
	Suggestion     string `yaml:"suggestion" json:"suggestion"`
	Pattern        string `yaml:"pattern" json:"pattern"`
	RelatedPattern string `yaml:"related_pattern" json:"related_pattern"`
}

// ConstraintV2Evidence holds evidence for an empirical constraint.
type ConstraintV2Evidence struct {
	Canonical          bool                        `yaml:"canonical" json:"canonical"`
	RuntimeRequirement string                      `yaml:"runtime_requirement" json:"runtime_requirement"`
	DerivedFromPattern string                      `yaml:"derived_from_pattern" json:"derived_from_pattern"`
	ConfidenceFactors  ConstraintConfidenceFactors `yaml:"confidence_factors" json:"confidence_factors"`
	Observations       []string                    `yaml:"observations" json:"observations"`
	ConfidenceBasis    string                      `yaml:"confidence_basis" json:"confidence_basis"`
}

// ConstraintConfidenceFactors holds numeric confidence factors.
type ConstraintConfidenceFactors struct {
	ObservationCount int `yaml:"observation_count" json:"observation_count"`
	ViolationCount   int `yaml:"violation_count" json:"violation_count"`
	FixCount         int `yaml:"fix_count" json:"fix_count"`
}

// ConstraintV2Match describes what code pattern triggers a constraint.
type ConstraintV2Match struct {
	Function    string      `yaml:"function" json:"function"`
	Pattern     string      `yaml:"pattern" json:"pattern"`
	Context     string      `yaml:"context" json:"context"`
	ContainsAny []MatchItem `yaml:"contains_any" json:"contains_any"`
	Patterns    []string    `yaml:"patterns" json:"patterns"`
	Without     string      `yaml:"without" json:"without"`
	Scope       string      `yaml:"scope" json:"scope"`
	Imports     string      `yaml:"imports" json:"imports"`
	Has         string      `yaml:"has" json:"has"`
	LOCGt       int         `yaml:"loc_greater_than" json:"loc_greater_than"`
}

// MatchItem is a pattern + label pair used in contains_any.
type MatchItem struct {
	Pattern string `yaml:"pattern" json:"pattern"`
	Label   string `yaml:"label" json:"label"`
}

// ── Loading functions ──

func (s *Scanner) loadExtendedKnowledge(knowledgeDir string) {
	if data, err := os.ReadFile(filepath.Join(knowledgeDir, "primitives.yaml")); err == nil {
		s.loadPrimitives(data)
	}
	if data, err := os.ReadFile(filepath.Join(knowledgeDir, "patterns.yaml")); err == nil {
		s.loadPatterns(data)
	}
	if data, err := os.ReadFile(filepath.Join(knowledgeDir, "constraints-v2.yaml")); err == nil {
		s.loadConstraintsV2(data)
	}
}

func (s *Scanner) loadPrimitives(data []byte) {
	var doc PrimitivesDoc
	if err := unmarshalYAML(string(data), &doc); err == nil && len(doc.Primitives) > 0 {
		s.primitives = doc.Primitives
	}
}

func (s *Scanner) loadPatterns(data []byte) {
	var patterns []Pattern
	for _, doc := range splitYAMLDocuments(string(data)) {
		var p Pattern
		if err := unmarshalYAML(doc, &p); err == nil && p.ID != "" {
			patterns = append(patterns, p)
		}
	}
	s.patterns = patterns
}

func (s *Scanner) loadConstraintsV2(data []byte) {
	var constraints []ConstraintV2
	for _, doc := range splitYAMLDocuments(string(data)) {
		var c ConstraintV2
		if err := unmarshalYAML(doc, &c); err == nil && c.ID != "" {
			constraints = append(constraints, c)
		}
	}
	s.constraintsV2 = constraints
}

// ── Accessors ──

// GetPrimitives returns the loaded behavioral primitives.
func (s *Scanner) GetPrimitives() []Primitive { return s.primitives }

// GetPatterns returns the loaded empirical patterns.
func (s *Scanner) GetPatterns() []Pattern { return s.patterns }

// GetConstraintsV2 returns the loaded enhanced constraints.
func (s *Scanner) GetConstraintsV2() []ConstraintV2 { return s.constraintsV2 }
