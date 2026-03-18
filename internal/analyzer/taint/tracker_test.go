package taint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/zakirkun/ice-tea/internal/parser/goparser"
)

func TestTaintPlaceholder(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	tracker := NewTracker(logger.Sugar())

	src := []byte(`package main
func main() {
	user := "admin"
	db.Query("...")
}`)

	gp := goparser.New()
	result, _ := gp.Parse("main.go", src)

	findings := tracker.Analyze(result)
	
	// currently implemented as a placeholder returning nil
	assert.Len(t, findings, 0)
}
