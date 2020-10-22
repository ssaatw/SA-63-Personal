// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/ssaatw/app/ent/jobtitle"
	"github.com/ssaatw/app/ent/personal"
)

// JobtitleCreate is the builder for creating a Jobtitle entity.
type JobtitleCreate struct {
	config
	mutation *JobtitleMutation
	hooks    []Hook
}

// SetJobtitlename sets the Jobtitlename field.
func (jc *JobtitleCreate) SetJobtitlename(s string) *JobtitleCreate {
	jc.mutation.SetJobtitlename(s)
	return jc
}

// AddPersonalIDs adds the personal edge to Personal by ids.
func (jc *JobtitleCreate) AddPersonalIDs(ids ...int) *JobtitleCreate {
	jc.mutation.AddPersonalIDs(ids...)
	return jc
}

// AddPersonal adds the personal edges to Personal.
func (jc *JobtitleCreate) AddPersonal(p ...*Personal) *JobtitleCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return jc.AddPersonalIDs(ids...)
}

// Mutation returns the JobtitleMutation object of the builder.
func (jc *JobtitleCreate) Mutation() *JobtitleMutation {
	return jc.mutation
}

// Save creates the Jobtitle in the database.
func (jc *JobtitleCreate) Save(ctx context.Context) (*Jobtitle, error) {
	if _, ok := jc.mutation.Jobtitlename(); !ok {
		return nil, &ValidationError{Name: "Jobtitlename", err: errors.New("ent: missing required field \"Jobtitlename\"")}
	}
	var (
		err  error
		node *Jobtitle
	)
	if len(jc.hooks) == 0 {
		node, err = jc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobtitleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			jc.mutation = mutation
			node, err = jc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(jc.hooks) - 1; i >= 0; i-- {
			mut = jc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (jc *JobtitleCreate) SaveX(ctx context.Context) *Jobtitle {
	v, err := jc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jc *JobtitleCreate) sqlSave(ctx context.Context) (*Jobtitle, error) {
	j, _spec := jc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	j.ID = int(id)
	return j, nil
}

func (jc *JobtitleCreate) createSpec() (*Jobtitle, *sqlgraph.CreateSpec) {
	var (
		j     = &Jobtitle{config: jc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: jobtitle.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: jobtitle.FieldID,
			},
		}
	)
	if value, ok := jc.mutation.Jobtitlename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: jobtitle.FieldJobtitlename,
		})
		j.Jobtitlename = value
	}
	if nodes := jc.mutation.PersonalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobtitle.PersonalTable,
			Columns: []string{jobtitle.PersonalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: personal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return j, _spec
}
