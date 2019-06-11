package format

import (
	"fmt"
	"os"
	"sort"

	"github.com/jedib0t/go-pretty/table"

	api "github.com/davidalpert/ginsight/api"
	sortStrategies "github.com/davidalpert/ginsight/util/sort"
)

func WriteObjectType(schemaTagType string, schemaTag string, objectType *api.ObjectType) {
	fmt.Printf("\nInsight ObjectType '%d' found for \"%s\"\n\n", objectType.ID, schemaTag)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	WriteObjectTypeHeader(t, schemaTagType)
	WriteObjectTypeRow(t, schemaTag, objectType)
	t.Render()
	fmt.Println()
}

func WriteObjectTypes(schemaTagType string, schemaTag string, objectTypes *[]api.ObjectType) {
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

func WriteObjectTypeRow(t table.Writer, schemaTag string, objectType *api.ObjectType) {
	t.AppendRow([]interface{}{schemaTag, objectType.ID, objectType.Name, objectType.Description, objectType.ParentObjectTypeID, objectType.Inherited, objectType.AbstractObjectType})
}

func WriteObjectSchemasAsTable(schemaList *api.ObjectSchemaList) {
	schemas := schemaList.Schemas
	sort.Sort(sortStrategies.ByObjectSchemaKey(schemas))

	fmt.Printf("\nInsight Object Schemas found in %s\n\n", api.DefaultClient().BaseURL)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Key", "Name", "Description", "# of Object Types", "# of Objects"})
	for _, schema := range schemaList.Schemas {
		t.AppendRow([]interface{}{schema.ID, schema.Key, schema.Name, schema.Description, schema.ObjectTypeCount, schema.ObjectCount})
	}
	t.Render()

	fmt.Println()
}

func WriteObjectSchemaDetail(schema *api.ObjectSchema) {
	fmt.Printf("\nFound schema: %s | %s | %s\n\n", schema.Key, schema.Name, schema.Description)
}
