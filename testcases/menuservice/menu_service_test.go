package menuservice

import (
	menu "DishTV/menuservice"
	//"fmt"
	"os"
	"testing"
)

func TestServiceMenu(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"Testing",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			menu.ServiceMenu()
			os.Exit(0)
		})
	}
}
