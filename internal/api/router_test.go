package api

import (
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestAddRoutes(t *testing.T) {
	type args struct {
		g  *gin.Engine
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "implemented",
			args: args{
				g:  gin.Default(),
				db: &gorm.DB{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddRoutes(tt.args.g, tt.args.db)
		})
	}
}
