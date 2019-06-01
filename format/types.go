package format

import (
	"fmt"
	"os"
	"sort"

	"github.com/jedib0t/go-pretty/table"

	insight "github.com/davidalpert/ginsight/insight"
	sortStrategies "github.com/davidalpert/ginsight/util/sort"
)

func WriteObjectType(schemaTagType string, schemaTag string, objectType *insight.ObjectType) {
	fmt.Printf("\nInsight ObjectType '%d' found for \"%s\"\n\n", objectType.ID, schemaTag)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	WriteObjectTypeHeader(t, schemaTagType)
	WriteObjectTypeRow(t, schemaTag, objectType)
	t.Render()
	fmt.Println()
}

func WriteObjectTypes(schemaTagType string, schemaTag string, objectTypes *[]insight.ObjectType) {
	fmt.Printf("\nInsight ObjectTypes found for \"%s\"\n\n", schemaTag)

	sort.Sort(sortStrategies.ByObjectTypeName(*objectTypes))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	WriteObjectTypeHeader(t, schemaTagType)
	for _, objectType := range *objectTypes {
		WriteObjectTypeRow(t, schemaTag, &objectType)
	}
	t.Render()

	fmt.Println()
}

func WriteObjectTypeHeader(t table.Writer, schemaTagType string) {
	t.AppendHeader(table.Row{schemaTagType, "Id", "Name", "Description", "ParentObjectTypeID", "Inherited", "Abstract"})
}

func WriteObjectTypeRow(t table.Writer, schemaTag string, objectType *insight.ObjectType) {
	t.AppendRow([]interface{}{schemaTag, objectType.ID, objectType.Name, objectType.Description, objectType.ParentObjectTypeID, objectType.Inherited, objectType.AbstractObjectType})
}

func WriteObjectSchemasAsTable(schemaList *insight.ObjectSchemaList) {
	schemas := schemaList.Schemas
	sort.Sort(sortStrategies.ByObjectSchemaKey(schemas))

	fmt.Printf("\nInsight Object Schemas found in %s\n\n", insight.DefaultClient().BaseURL)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Key", "Name", "Description", "# of Object Types", "# of Objects"})
	for _, schema := range schemaList.Schemas {
		t.AppendRow([]interface{}{schema.ID, schema.Key, schema.Name, schema.Description, schema.ObjectTypeCount, schema.ObjectCount})
	}
	t.Render()

	fmt.Println()
}

func WriteObjectSchemaDetail(schema *insight.ObjectSchema) {
	fmt.Printf("\nFound schema: %s | %s | %s\n\n", schema.Key, schema.Name, schema.Description)
}
