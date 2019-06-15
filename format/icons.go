package format

import (
	"fmt"
	"net/url"
	"os"
	"sort"

	"github.com/jedib0t/go-pretty/table"

	api "github.com/davidalpert/ginsight/api"
	sortStrategies "github.com/davidalpert/ginsight/util/sort"
)

func WriteIcons(schemaTagType string, schemaTag string, icons *[]api.ObjectIcon) {
	fmt.Printf("\nInsight Icons found for \"%s\"\n\n", schemaTag)

	sort.Sort(sortStrategies.ByObjectIconID(*icons))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	WriteObjectIconHeader(t, schemaTagType)
	for _, icon := range *icons {
		WriteObjectIconRow(t, schemaTag, &icon)
	}
	t.Render()

	fmt.Println()
}

func WriteObjectIconHeader(t table.Writer, schemaTagType string) {
	t.AppendHeader(table.Row{
		schemaTagType,
		"Id",
		"Name",
		"Removable",
		"Url48",
	})
}

func WriteObjectIconRow(t table.Writer, schemaTag string, icon *api.ObjectIcon) {
	t.AppendRow([]interface{}{
		schemaTag,
		icon.ID,
		icon.Name,
		icon.Removable,
		stripQueryParam(icon.Url48, "uuid"),
	})
}

func stripQueryParam(inURL string, stripKey string) string {
	u, err := url.Parse(inURL)
	if err != nil {
		// TODO: log or handle error, in the meanwhile just return the original
		return inURL
	}
	q := u.Query()
	q.Del(stripKey)
	u.RawQuery = q.Encode()
	return u.String()
}
