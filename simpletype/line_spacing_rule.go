package simpletype

type LineSpacingRule string

const (
	LineSpacingRuleAUTO    LineSpacingRule = "auto"    //Automatically Determined Line Height.
	LineSpacingRuleEXACT   LineSpacingRule = "exact"   //Exact Line Height.
	LineSpacingRuleAtLeast LineSpacingRule = "atLeast" //Minimum Line Height.
)
