package main

// Log levels, kept as constants so repeated literals stay below goconst's threshold.
const (
	logInfo  = "info"
	logWarn  = "warn"
	logError = "error"
)

// Structured log field keys.
const (
	flowIDField = "flow_id"
	targetField = "target"
	errField    = "err"
)

// Nuclei template tags selected by coverage categories.
const (
	tagCVE       = "cve"
	tagExposure  = "exposure"
	tagPanel     = "panel"
	tagFile      = "file"
	tagMisconfig = "misconfig"
	tagTech      = "tech"
	tagSQLi      = "sqli"
	tagRCE       = "rce"
	tagCMDi      = "cmdi"
	tagSSRF      = "ssrf"
	tagRedirect  = "redirect"
	tagXSS       = "xss"
	tagSSTI      = "ssti"
	tagXXE       = "xxe"
	tagCRLF      = "crlf"

	// Display names that differ from their tag value.
	nameExposures    = "exposures"
	nameCVEInjection = "cve-injection"
)
