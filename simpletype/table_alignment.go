package simpletype

/**
 * Table Alignment Type.
 *
 * Introduced in ISO/IEC-29500:2008.
 *
 * @since 0.13.0
 */

type tableAlignmentType string

const (
	START  tableAlignmentType = "start"
	CENTER tableAlignmentType = "center"
	END    tableAlignmentType = "end"
)
