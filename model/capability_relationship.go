package model

// Relationship maps to a DTDL relationship type
type Relationship struct {
	NamedEntity
	maxMultiplicity int
	minMultiplicity int
	target string
	writable bool
}

// MaxMultiplicity returns the maximum no. of instances of a target allowed for the relationship.
func (r *Relationship) MaxMultiplicity() int {
	return r.maxMultiplicity
}

// MinMultiplicity returns the minimum no. of instances of a target allowed for the relationship.
func (r *Relationship) MinMultiplicity() int {
	return r.minMultiplicity
}

// Target returns the target of the relationship
func (r *Relationship) Target() string {
	return r.target
}

// Writable returns whether the relationship is writable
func (r *Relationship) Writable() bool {
	return r.writable
}

func parseRelationship(i map[string]interface{}, t entityTracker) *Relationship {
	r := &Relationship{
		NamedEntity: parseNamedEntity(i),
	}

	if maxm, ok := i["maxMultiplicity"].(int); ok {
		r.maxMultiplicity = int(maxm)
	}

	if minm, ok := i["minMultiplicity"].(int); ok {
		r.minMultiplicity = int(minm)
	}

	if w, ok := i["writable"].(bool); ok {
		r.writable = w
	}

	if tar, ok := i["target"].(string); ok {
		r.target = tar
	}

	t.Add(r)
	return r
}