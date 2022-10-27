package repository

import (
	"github.com/jmoiron/sqlx"
	"practice_go/entity"
	"testing"
)

func TestInsertComment(t *testing.T) {
	type args struct {
		db      *sqlx.DB
		comment *entity.Comment
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertComment(tt.args.db, tt.args.comment); (err != nil) != tt.wantErr {
				t.Errorf("InsertComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
