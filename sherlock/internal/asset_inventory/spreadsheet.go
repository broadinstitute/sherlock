package asset_inventory

import (
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strings"
)

type Spreadsheet interface {
	ToRows() [][]any
}

var columnReplacementRegex = regexp.MustCompile(`[^a-zA-Z0-9]`)

type spreadsheet struct {
	mapping        spreadsheetColumnMapping
	width          int
	skippedRows    [][]any
	previousAssets map[string]assetRow
	currentAssets  []assetRow
}

type assetRow []any

func NewSpreadsheet(rows [][]any, skip int) (Spreadsheet, error) {
	s := &spreadsheet{}

	// Find and store the index of each column based on the header row
	mappingType := reflect.TypeOf(s.mapping)
columnFindLoop:
	for i := 0; i < mappingType.NumField(); i++ {
		field := mappingType.Field(i)
		columnName := field.Tag.Get("column")
		columnNameToFind := strings.ToLower(columnReplacementRegex.ReplaceAllString(columnName, ""))

		for j := 0; j < len(rows[skip]); j++ {
			if cellAsString, ok := rows[skip][j].(string); !ok {
				return nil, fmt.Errorf("expected column name to be a string but got %T", rows[skip][j])
			} else {
				if strings.Contains(strings.ToLower(columnReplacementRegex.ReplaceAllString(cellAsString, "")), columnNameToFind) {
					s.width = max(s.width, j) // Record that rows should be at least this wide to accommodate this column
					reflect.ValueOf(&s.mapping).Elem().Field(i).SetInt(int64(j))
					continue columnFindLoop
				}
			}
		}

		return nil, fmt.Errorf("could not find a column header containing %q (whitespace-, symbol-, and case-insensitive)", columnName)
	}

	// We looked for the width already, but in case any rows happen to be longer, we'll standardize
	for _, row := range rows {
		s.width = max(s.width, len(row))
	}

	s.skippedRows = rows[:skip+1] // Skipped plus the header row

	// Now we populate our list of existing assets, to persist cells we wouldn't otherwise set.
	// We make sure each is at least width long, so that we can index into it without checking bounds.
	s.previousAssets = make(map[string]assetRow)
	for _, rawRow := range rows[skip+1:] {
		var row assetRow
		if len(rawRow) < s.width {
			row = make(assetRow, s.width)
			copy(row, rawRow)
		} else {
			row = rawRow
		}

		if uniqueAssetIdentifier := s.mapping.getUniqueAssetIdentifier(row); uniqueAssetIdentifier != "" {
			s.previousAssets[uniqueAssetIdentifier] = row
		}
	}

	return s, nil
}

func (s *spreadsheet) ToRows() [][]any {
	rows := make([][]any, 0, len(s.skippedRows)+len(s.currentAssets))
	rows = append(rows, s.skippedRows...)

	slices.SortFunc(s.currentAssets, func(a, b assetRow) int {
		return strings.Compare(s.mapping.getTeam(a), s.mapping.getTeam(b))<<4 +
			strings.Compare(s.mapping.getTerraServiceComponent(a), s.mapping.getTerraServiceComponent(b))<<3 +
			strings.Compare(s.mapping.getGcpProjectOrAzureSubscription(a), s.mapping.getGcpProjectOrAzureSubscription(b))<<2 +
			strings.Compare(s.mapping.getAssetType(a), s.mapping.getAssetType(b))<<1 +
			strings.Compare(s.mapping.getUniqueAssetIdentifier(a), s.mapping.getUniqueAssetIdentifier(b))
	})
	for _, row := range s.currentAssets {
		rows = append(rows, row)
	}

	return rows
}
