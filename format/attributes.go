package format

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/jedib0t/go-pretty/table"

	api "github.com/davidalpert/ginsight/api"
	sortStrategies "github.com/davidalpert/ginsight/util/sort"
)

func WriteObjectTypeAttribute(schemaTagType string, schemaTag string, objectType *api.ObjectType, attribute *api.ObjectTypeAttribute) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	WriteObjectTypeAttributeHeader(t, schemaTagType)
	WriteObjectTypeAttributeRow(t, schemaTag, objectType, attribute)
	t.Render()
	fmt.Println()
}

func WriteObjectTypeAttributes(schemaTagType string, schemaTag string, objectType *api.ObjectType, attributes *[]api.ObjectTypeAttribute) {
	sort.Sort(sortStrategies.ByObjectTypeAttributeID(*attributes))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	WriteObjectTypeAttributeHeader(t, schemaTagType)
	for _, attribute := range *attributes {
		WriteObjectTypeAttributeRow(t, schemaTag, objectType, &attribute)
	}
	t.Render()

	fmt.Println()
}

func WriteObjectTypeAttributeHeader(t table.Writer, schemaTagType string) {
	t.AppendHeader(table.Row{schemaTagType, "ObjectType", "Name", "Attr ID", "Description", "Type", "DefaultType", "System", "Editable"})
}

func WriteObjectTypeAttributeRow(t table.Writer, schemaTag string, objectType *api.ObjectType, attribute *api.ObjectTypeAttribute) {
	attributeTypeName := api.AttributeDefaultTypeIDToName(attribute.TypeID)
	var objectTypeIdentifier string

	if objectType == nil {
		objectType = attribute.ObjectType
	}
	if objectType == nil {
		objectTypeIdentifier = "unknown"
	} else {
		objectTypeIdentifier = fmt.Sprintf("%s (%d)", (*objectType).Name, (*objectType).ID) // TODO: refactor to objectType.UniqueName()
	}

	t.AppendRow([]interface{}{schemaTag, objectTypeIdentifier, attribute.Name, attribute.ID, attribute.Description, attributeTypeName, strings.ToLower(attribute.DefaultType.Name), attribute.System, attribute.Editable})
}
