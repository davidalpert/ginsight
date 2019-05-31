package format

import (
	"fmt"
	"os"
	"sort"

	"github.com/jedib0t/go-pretty/table"

	insight "github.com/davidalpert/ginsight/insight"
	sortStrategies "github.com/davidalpert/ginsight/util/sort"
)

func WriteObjectType(schemaTagType string, schemaTag string, objectType insight.ObjectType) {
	fmt.Printf("\nInsight ObjectType '%d' found for \"%s\"\n\n", objectType.ID, schemaTag)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	WriteObjectTypeHeader(t, schemaTagType)
	WriteObjectTypeRow(t, schemaTag, objectType)
	t.Render()
	fmt.Println()
}

func WriteObjectTypes(schemaTagType string, schemaTag string, objectTypes []insight.ObjectType) {
	fmt.Printf("\nInsight ObjectTypes found for \"%s\"\n\n", schemaTag)

	sort.Sort(sortStrategies.ByObjectTypeName(objectTypes))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	WriteObjectTypeHeader(t, schemaTagType)
	for _, objectType := range objectTypes {
		WriteObjectTypeRow(t, schemaTag, objectType)
	}
	t.Render()

	fmt.Println()
}

func WriteObjectTypeHeader(t table.Writer, schemaTagType string) {
	t.AppendHeader(table.Row{schemaTagType, "Id", "Name", "Description", "ParentObjectTypeID", "Inherited", "Abstract"})
}

func WriteObjectTypeRow(t table.Writer, schemaTag string, objectType insight.ObjectType) {
	t.AppendRow([]interface{}{schemaTag, objectType.ID, objectType.Name, objectType.Description, objectType.ParentObjectTypeID, objectType.Inherited, objectType.AbstractObjectType})
}
