package tools

import (
	"fmt"
)

//GdriveViewLink View Link
func GdriveViewLink(fileID string) string {
	return fmt.Sprintf("https://drive.google.com/uc?export=view&id=%s", fileID)
}
