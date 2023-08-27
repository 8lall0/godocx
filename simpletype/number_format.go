package simpletype

type NumberFormat string

const (
	NumberFormatDECIMAL                         NumberFormat = "decimal"
	NumberFormatUPPER_ROMAN                     NumberFormat = "upperRoman"                   //Uppercase Roman Numerals
	NumberFormatLOWER_ROMAN                     NumberFormat = "lowerRoman"                   //Lowercase Roman Numerals
	NumberFormatUPPER_LETTER                    NumberFormat = "upperLetter"                  //Uppercase Latin Alphabet
	NumberFormatLOWER_LETTER                    NumberFormat = "lowerLetter"                  //Lowercase Latin Alphabet
	NumberFormatORDINAL                         NumberFormat = "ordinal"                      //Ordinal
	NumberFormatCARDINAL_TEXT                   NumberFormat = "cardinalText"                 //Cardinal Text
	NumberFormatORDINAL_TEXT                    NumberFormat = "ordinalText"                  //Ordinal Text
	NumberFormatHEX                             NumberFormat = "hex"                          //Hexadecimal Numbering
	NumberFormatCHICAGO                         NumberFormat = "chicago"                      //Chicago Manual of Style
	NumberFormatIDEOGRAPH_DIGITAL               NumberFormat = "ideographDigital"             //Ideographs
	NumberFormatJAPANESE_COUNTING               NumberFormat = "japaneseCounting"             //Japanese Counting System
	NumberFormatAIUEO                           NumberFormat = "aiueo"                        //AIUEO Order Hiragana
	NumberFormatIROHA                           NumberFormat = "iroha"                        //Iroha Ordered Katakana
	NumberFormatDECIMAL_FULL_WIDTH              NumberFormat = "decimalFullWidth"             //Double Byte Arabic Numerals
	NumberFormatDECIMAL_HALF_WIDTH              NumberFormat = "decimalHalfWidth"             //Single Byte Arabic Numerals
	NumberFormatJAPANESE_LEGAL                  NumberFormat = "japaneseLegal"                //Japanese Legal Numbering
	NumberFormatJAPANESE_DIGITAL_TEN_THOUSAND   NumberFormat = "japaneseDigitalTenThousand"   //Japanese Digital Ten Thousand Counting System
	NumberFormatDECIMAL_ENCLOSED_CIRCLE         NumberFormat = "decimalEnclosedCircle"        //Decimal Numbers Enclosed in a Circle
	NumberFormatDECIMAL_FULL_WIDTH2             NumberFormat = "decimalFullWidth2"            //Double Byte Arabic Numerals Alternate
	NumberFormatAIUEO_FULL_WIDTH                NumberFormat = "aiueoFullWidth"               //Full-Width AIUEO Order Hiragana
	NumberFormatIROHA_FULL_WIDTH                NumberFormat = "irohaFullWidth"               //Full-Width Iroha Ordered Katakana
	NumberFormatDECIMAL_ZERO                    NumberFormat = "decimalZero"                  //Initial Zero Arabic Numerals
	NumberFormatBULLET                          NumberFormat = "bullet"                       //Bullet
	NumberFormatGANADA                          NumberFormat = "ganada"                       //Korean Ganada Numbering
	NumberFormatCHOSUNG                         NumberFormat = "chosung"                      //Korean Chosung Numbering
	NumberFormatDECIMAL_ENCLOSED_FULL_STOP      NumberFormat = "decimalEnclosedFullstop"      //Decimal Numbers Followed by a Period
	NumberFormatDECIMAL_ENCLOSED_PAREN          NumberFormat = "decimalEnclosedParen"         //Decimal Numbers Enclosed in Parenthesis
	NumberFormatDECIMAL_ENCLOSED_CIRCLE_CHINESE NumberFormat = "decimalEnclosedCircleChinese" //Decimal Numbers Enclosed in a Circle
	NumberFormatIDEOGRAPHENCLOSEDCIRCLE         NumberFormat = "ideographEnclosedCircle"      //Ideographs Enclosed in a Circle
	NumberFormatIDEOGRAPH_TRADITIONAL           NumberFormat = "ideographTraditional"         //Traditional Ideograph Format
	NumberFormatIDEOGRAPH_ZODIAC                NumberFormat = "ideographZodiac"              //Zodiac Ideograph Format
	NumberFormatIDEOGRAPH_ZODIAC_TRADITIONAL    NumberFormat = "ideographZodiacTraditional"   //Traditional Zodiac Ideograph Format
	NumberFormatTAIWANESE_COUNTING              NumberFormat = "taiwaneseCounting"            //Taiwanese Counting System
	NumberFormatIDEOGRAPH_LEGAL_TRADITIONAL     NumberFormat = "ideographLegalTraditional"    //Traditional Legal Ideograph Format
	NumberFormatTAIWANESE_COUNTING_THOUSAND     NumberFormat = "taiwaneseCountingThousand"    //Taiwanese Counting Thousand System
	NumberFormatTAIWANESE_DIGITAL               NumberFormat = "taiwaneseDigital"             //Taiwanese Digital Counting System
	NumberFormatCHINESE_COUNTING                NumberFormat = "chineseCounting"              //Chinese Counting System
	NumberFormatCHINESE_LEGAL_SIMPLIFIED        NumberFormat = "chineseLegalSimplified"       //Chinese Legal Simplified Format
	NumberFormatCHINESE_COUNTING_THOUSAND       NumberFormat = "chineseCountingThousand"      //Chinese Counting Thousand System
	NumberFormatKOREAN_DIGITAL                  NumberFormat = "koreanDigital"                //Korean Digital Counting System
	NumberFormatKOREAN_COUNTING                 NumberFormat = "koreanCounting"               //Korean Counting System
	NumberFormatKOREAN_LEGAL                    NumberFormat = "koreanLegal"                  //Korean Legal Numbering
	NumberFormatKOREAN_DIGITAL2                 NumberFormat = "koreanDigital2"               //Korean Digital Counting System Alternate
	NumberFormatVIETNAMESE_COUNTING             NumberFormat = "vietnameseCounting"           //Vietnamese Numerals
	NumberFormatRUSSIAN_LOWER                   NumberFormat = "russianLower"                 //Lowercase Russian Alphabet
	NumberFormatRUSSIAN_UPPER                   NumberFormat = "russianUpper"                 //Uppercase Russian Alphabet
	NumberFormatNONE                            NumberFormat = "none"                         //No Numbering
	NumberFormatNUMBER_IN_DASH                  NumberFormat = "numberInDash"                 //Number With Dashes
	NumberFormatHEBREW1                         NumberFormat = "hebrew1"                      //Hebrew Numerals
	NumberFormatHEBREW2                         NumberFormat = "hebrew2"                      //Hebrew Alphabet
	NumberFormatARABIC_ALPHA                    NumberFormat = "arabicAlpha"                  //Arabic Alphabet
	NumberFormatARABIC_ABJAD                    NumberFormat = "arabicAbjad"                  //Arabic Abjad Numerals
	NumberFormatHINDI_VOWELS                    NumberFormat = "hindiVowels"                  //Hindi Vowels
	NumberFormatHINDI_CONSONANTS                NumberFormat = "hindiConsonants"              //Hindi Consonants
	NumberFormatHINDI_NUMBERS                   NumberFormat = "hindiNumbers"                 //Hindi Numbers
	NumberFormatHINDI_COUNTING                  NumberFormat = "hindiCounting"                //Hindi Counting System
	NumberFormatTHAI_LETTERS                    NumberFormat = "thaiLetters"                  //Thai Letters
	NumberFormatTHAI_NUMBERS                    NumberFormat = "thaiNumbers"                  //Thai Numerals
	NumberFormatTHAI_COUNTING                   NumberFormat = "thaiCounting"                 //Thai Counting System
)
