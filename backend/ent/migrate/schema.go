// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "departmentname", Type: field.TypeString, Unique: true},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:        "departments",
		Columns:     DepartmentsColumns,
		PrimaryKey:  []*schema.Column{DepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// JobtitlesColumns holds the columns for the "jobtitles" table.
	JobtitlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "jobtitlename", Type: field.TypeString, Unique: true},
	}
	// JobtitlesTable holds the schema information for the "jobtitles" table.
	JobtitlesTable = &schema.Table{
		Name:        "jobtitles",
		Columns:     JobtitlesColumns,
		PrimaryKey:  []*schema.Column{JobtitlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PersonalsColumns holds the columns for the "personals" table.
	PersonalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "personal_name", Type: field.TypeString},
		{Name: "personal_mail", Type: field.TypeString},
		{Name: "personal_phone", Type: field.TypeString},
		{Name: "added", Type: field.TypeTime},
		{Name: "department_id", Type: field.TypeInt, Nullable: true},
		{Name: "jobtitle_id", Type: field.TypeInt, Nullable: true},
		{Name: "systemmember_id", Type: field.TypeInt, Nullable: true},
	}
	// PersonalsTable holds the schema information for the "personals" table.
	PersonalsTable = &schema.Table{
		Name:       "personals",
		Columns:    PersonalsColumns,
		PrimaryKey: []*schema.Column{PersonalsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "personals_departments_personal",
				Columns: []*schema.Column{PersonalsColumns[5]},

				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "personals_jobtitles_personal",
				Columns: []*schema.Column{PersonalsColumns[6]},

				RefColumns: []*schema.Column{JobtitlesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "personals_systemmembers_personal",
				Columns: []*schema.Column{PersonalsColumns[7]},

				RefColumns: []*schema.Column{SystemmembersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SystemmembersColumns holds the columns for the "systemmembers" table.
	SystemmembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "mail", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// SystemmembersTable holds the schema information for the "systemmembers" table.
	SystemmembersTable = &schema.Table{
		Name:        "systemmembers",
		Columns:     SystemmembersColumns,
		PrimaryKey:  []*schema.Column{SystemmembersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DepartmentsTable,
		JobtitlesTable,
		PersonalsTable,
		SystemmembersTable,
	}
)

func init() {
	PersonalsTable.ForeignKeys[0].RefTable = DepartmentsTable
	PersonalsTable.ForeignKeys[1].RefTable = JobtitlesTable
	PersonalsTable.ForeignKeys[2].RefTable = SystemmembersTable
}
