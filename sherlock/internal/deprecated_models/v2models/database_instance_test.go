package v2models

import (
	"github.com/broadinstitute/sherlock/go-shared/pkg/testutils"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"gorm.io/gorm"
	"testing"
)

// Test_databaseInstanceSelectorToQuery can't fully test databaseInstanceSelectorToQuery
// similar to Test_chartReleaseSelectorToQuery because these tests run without a database.
func Test_databaseInstanceSelectorToQuery(t *testing.T) {
	type args struct {
		db       *gorm.DB
		selector string
	}
	tests := []struct {
		name    string
		args    args
		want    DatabaseInstance
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{selector: ""},
			wantErr: true,
		},
		{
			name:    "invalid",
			args:    args{selector: "something obviously invalid"},
			wantErr: true,
		},
		{
			name: "valid id",
			args: args{selector: "123"},
			want: DatabaseInstance{Model: gorm.Model{ID: 123}},
		},
		{
			name:    "invalid id",
			args:    args{selector: testutils.StringNumberTooBigForInt},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := databaseInstanceSelectorToQuery(tt.args.db, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("databaseInstanceSelectorToQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_databaseInstanceToSelectors(t *testing.T) {
	type args struct {
		databaseInstance *DatabaseInstance
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "none",
			args: args{databaseInstance: &DatabaseInstance{}},
			want: nil,
		},
		{
			name: "id",
			args: args{databaseInstance: &DatabaseInstance{Model: gorm.Model{ID: 123}}},
			want: []string{"123"},
		},
		{
			name: "chart release id",
			args: args{databaseInstance: &DatabaseInstance{ChartReleaseID: 456}},
			want: []string{"chart-release/456"},
		},
		{
			name: "id and chart release id",
			args: args{databaseInstance: &DatabaseInstance{
				Model:          gorm.Model{ID: 123},
				ChartReleaseID: 456,
			}},
			want: []string{"chart-release/456", "123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := databaseInstanceToSelectors(tt.args.databaseInstance)
			testutils.AssertNoDiff(t, tt.want, got)
		})
	}
}

func Test_validateDatabaseInstance(t *testing.T) {
	type args struct {
		databaseInstance *DatabaseInstance
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "nil",
			args:    args{databaseInstance: nil},
			wantErr: true,
		},
		{
			name: "invalid no chart release",
			args: args{databaseInstance: &DatabaseInstance{
				Platform:        utils.PointerTo("google"),
				GoogleProject:   utils.PointerTo("abc"),
				InstanceName:    utils.PointerTo("ghi"),
				DefaultDatabase: utils.PointerTo("jkl"),
			}},
			wantErr: true,
		},
		{
			name: "invalid bad platform",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID:  123,
				Platform:        utils.PointerTo("foo bar"),
				GoogleProject:   utils.PointerTo("abc"),
				InstanceName:    utils.PointerTo("ghi"),
				DefaultDatabase: utils.PointerTo("jkl"),
			}},
			wantErr: true,
		},
		{
			name: "invalid no platform",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID:  123,
				GoogleProject:   utils.PointerTo("abc"),
				InstanceName:    utils.PointerTo("ghi"),
				DefaultDatabase: utils.PointerTo("jkl"),
			}},
			wantErr: true,
		},
		{
			name: "invalid google no project",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID:  123,
				Platform:        utils.PointerTo("google"),
				InstanceName:    utils.PointerTo("ghi"),
				DefaultDatabase: utils.PointerTo("jkl"),
			}},
			wantErr: true,
		},
		{
			name: "invalid google no instance name",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID:  123,
				Platform:        utils.PointerTo("google"),
				GoogleProject:   utils.PointerTo("abc"),
				DefaultDatabase: utils.PointerTo("jkl"),
			}},
			wantErr: true,
		},
		{
			name: "invalid google no default database",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID: 123,
				Platform:       utils.PointerTo("google"),
				GoogleProject:  utils.PointerTo("abc"),
				InstanceName:   utils.PointerTo("ghi"),
			}},
			wantErr: true,
		},
		{
			name: "valid google",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID:  123,
				Platform:        utils.PointerTo("google"),
				GoogleProject:   utils.PointerTo("abc"),
				InstanceName:    utils.PointerTo("ghi"),
				DefaultDatabase: utils.PointerTo("jkl"),
			}},
			wantErr: false,
		},
		{
			name: "valid maximal google",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID:  123,
				Platform:        utils.PointerTo("google"),
				GoogleProject:   utils.PointerTo("abc"),
				InstanceName:    utils.PointerTo("ghi"),
				DefaultDatabase: utils.PointerTo("jkl"),
			}},
			wantErr: false,
		},
		{
			name: "invalid azure no instance name",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID:  123,
				Platform:        utils.PointerTo("azure"),
				DefaultDatabase: utils.PointerTo("jkl"),
			}},
			wantErr: true,
		},
		{
			name: "invalid azure no default database",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID: 123,
				Platform:       utils.PointerTo("azure"),
				InstanceName:   utils.PointerTo("ghi"),
			}},
			wantErr: true,
		},
		{
			name: "valid azure",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID:  123,
				Platform:        utils.PointerTo("azure"),
				InstanceName:    utils.PointerTo("ghi"),
				DefaultDatabase: utils.PointerTo("jkl"),
			}},
			wantErr: false,
		},
		{
			name: "valid maximal azure",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID:  123,
				Platform:        utils.PointerTo("azure"),
				GoogleProject:   utils.PointerTo("abc"),
				InstanceName:    utils.PointerTo("ghi"),
				DefaultDatabase: utils.PointerTo("jkl"),
			}},
			wantErr: false,
		},
		{
			name: "invalid kubernetes no default database",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID: 123,
				Platform:       utils.PointerTo("kubernetes"),
			}},
			wantErr: true,
		},
		{
			name: "valid kubernetes",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID:  123,
				Platform:        utils.PointerTo("kubernetes"),
				DefaultDatabase: utils.PointerTo("jkl"),
			}},
			wantErr: false,
		},
		{
			name: "valid maximal kubernetes",
			args: args{databaseInstance: &DatabaseInstance{
				ChartReleaseID:  123,
				Platform:        utils.PointerTo("kubernetes"),
				GoogleProject:   utils.PointerTo("abc"),
				InstanceName:    utils.PointerTo("ghi"),
				DefaultDatabase: utils.PointerTo("jkl"),
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateDatabaseInstance(tt.args.databaseInstance); (err != nil) != tt.wantErr {
				t.Errorf("validateDatabaseInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
