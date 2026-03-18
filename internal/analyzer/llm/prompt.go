package llm

import (
	"bytes"
	"text/template"
)

const systemPrompt = `You are Ice Tea, an expert AI application security reviewer.
You are analyzing a potential security vulnerability found by a static analysis tool.

Your job is to:
1. Understand the code snippet and its context.
2. Determine if the vulnerability is a True Positive or False Positive.
3. If it is a True Positive, explain exactly WHY and HOW it can be exploited.
4. Provide a clear, actionable remediation strategy.
5. Provide the fixed code snippet.

Your output MUST be valid JSON matching the following schema:
{
  "is_vulnerable": boolean,
  "confidence": "high" | "medium" | "low",
  "explanation": "Brief explanation of the vulnerability or why it's a false positive",
  "fix": "Step by step instructions on how to fix it (empty if not vulnerable)",
  "fix_code": "The rewritten safe code snippet (empty if not vulnerable)"
}
`

const userPromptTpl = `
Analyze the following finding:

File: {{.File}}
Line: {{.StartLine}}
Rule ID: {{.Rule.ID}} ({{.Rule.Name}})
Vulnerability Type: {{.Rule.Description}}

Code Snippet:
{{.CodeSnippet}}

Analyze the code snippet step-by-step to determine if the vulnerability is real.
Respond ONLY with the requested JSON format.
`

var userPromptTemplate *template.Template

func init() {
	userPromptTemplate = template.Must(template.New("prompt").Parse(userPromptTpl))
}

// buildPrompts generates the system and user prompts for the LLM
func buildPrompts(req AnalysisRequest) (string, string, error) {
	var userPrompt bytes.Buffer
	if err := userPromptTemplate.Execute(&userPrompt, req); err != nil {
		return "", "", err
	}
	return systemPrompt, userPrompt.String(), nil
}
