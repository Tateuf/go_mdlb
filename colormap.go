package main

func give_color() [][3]int {
	colormap := [][3]int{
		{228, 194, 129},
		{70, 217, 46},
		{106, 112, 192},
		{165, 84, 214},
		{130, 151, 133},
		{155, 181, 141},
		{172, 172, 231},
		{117, 210, 40},
		{63, 64, 62},
		{61, 236, 202},
		{47, 105, 51},
		{78, 123, 175},
		{84, 62, 197},
		{224, 171, 247},
		{182, 37, 147},
		{132, 187, 179},
		{215, 236, 218},
		{244, 97, 122},
		{55, 38, 139},
		{127, 145, 219},
		{35, 91, 137},
		{182, 39, 133},
		{113, 122, 254},
		{184, 89, 74},
		{51, 64, 47},
		{62, 213, 85},
		{209, 71, 48},
		{38, 127, 163},
		{78, 44, 190},
		{217, 37, 247},
		{170, 213, 251},
		{75, 27, 206},
		{200, 106, 156},
		{88, 181, 132},
		{168, 113, 99},
		{250, 99, 250},
		{101, 28, 117},
		{89, 24, 135},
		{53, 211, 180},
		{168, 110, 59},
		{129, 143, 198},
		{166, 224, 213},
		{194, 162, 132},
		{26, 122, 41},
		{120, 193, 78},
		{47, 52, 65},
		{148, 80, 174},
		{211, 126, 100},
		{245, 243, 190},
		{213, 111, 227},
		{187, 89, 78},
		{150, 83, 99},
		{82, 56, 125},
		{195, 253, 88},
		{73, 118, 217},
		{105, 161, 201},
		{166, 184, 90},
		{76, 71, 132},
		{175, 39, 25},
		{139, 243, 65},
		{70, 210, 228},
		{200, 122, 245},
		{190, 229, 130},
		{246, 149, 209},
		{85, 208, 71},
		{61, 118, 44},
		{174, 231, 99},
		{193, 79, 124},
		{42, 87, 229},
		{66, 26, 216},
		{163, 56, 122},
		{210, 139, 29},
		{217, 80, 168},
		{240, 128, 89},
		{218, 40, 251},
		{104, 230, 62},
		{132, 104, 159},
		{39, 201, 38},
		{167, 128, 130},
		{70, 181, 154},
		{166, 199, 233},
		{246, 243, 179},
		{89, 204, 142},
		{137, 125, 42},
		{70, 254, 109},
		{239, 209, 163},
		{163, 149, 45},
		{158, 53, 63},
		{96, 97, 34},
		{116, 179, 125},
		{95, 129, 116},
		{63, 49, 122},
		{245, 47, 116},
		{34, 64, 241},
		{221, 116, 211},
		{167, 136, 75},
		{160, 129, 77},
		{170, 209, 81},
		{47, 71, 145},
		{86, 158, 165},
		{71, 127, 234},
		{141, 215, 65},
		{186, 128, 38},
		{241, 236, 89},
		{80, 104, 69},
		{126, 45, 45},
		{149, 93, 211},
		{116, 154, 164},
		{219, 96, 52},
		{183, 166, 129},
		{214, 172, 238},
		{165, 145, 195},
		{234, 238, 108},
		{225, 100, 184},
		{119, 203, 226},
		{214, 221, 170},
		{68, 24, 53},
		{224, 158, 102},
		{207, 174, 169},
		{212, 25, 28},
		{166, 179, 224},
		{119, 28, 212},
		{155, 169, 48},
		{178, 97, 120},
		{249, 137, 213},
		{122, 222, 81},
		{227, 249, 24},
		{73, 87, 131},
		{236, 55, 191},
		{37, 48, 251},
		{206, 24, 57},
		{61, 151, 47},
		{218, 48, 211},
		{48, 98, 85},
		{46, 141, 99},
		{252, 94, 175},
		{91, 24, 82},
		{58, 234, 151},
		{101, 254, 242},
		{74, 208, 159},
		{70, 168, 82},
		{125, 211, 217},
		{212, 171, 194},
		{39, 67, 108},
		{217, 248, 39},
		{220, 39, 132},
		{129, 35, 196},
		{149, 43, 196},
		{45, 158, 26},
		{113, 227, 161},
		{248, 198, 61},
		{121, 205, 43},
		{143, 34, 74},
		{55, 44, 87},
		{188, 144, 217},
		{85, 106, 42},
		{243, 132, 45},
		{229, 196, 86},
		{244, 220, 51},
		{206, 62, 196},
		{26, 86, 126},
		{172, 65, 211},
		{145, 181, 53},
		{68, 24, 61},
		{74, 70, 145},
		{121, 192, 92},
		{74, 94, 148},
		{37, 180, 94},
		{93, 237, 222},
		{216, 231, 110},
		{249, 126, 146},
		{43, 155, 26},
		{161, 76, 189},
		{82, 113, 133},
		{120, 105, 194},
		{28, 216, 120},
		{67, 122, 215},
		{55, 243, 181},
		{159, 122, 175},
		{220, 224, 48},
		{151, 85, 214},
		{52, 67, 133},
		{182, 144, 39},
		{244, 185, 30},
		{109, 137, 97},
		{154, 167, 87},
		{120, 99, 72},
		{103, 148, 226},
		{137, 194, 235},
		{165, 117, 232},
		{238, 92, 165},
		{186, 127, 97},
		{101, 47, 237},
		{145, 111, 222},
		{167, 187, 130},
		{215, 222, 207},
		{50, 208, 218},
		{25, 47, 187},
		{235, 110, 221},
		{155, 81, 25},
		{153, 61, 99},
		{25, 37, 142},
		{110, 58, 149},
		{73, 111, 213},
		{28, 144, 121},
		{77, 223, 225},
		{148, 203, 81},
		{88, 88, 136},
		{39, 145, 218},
		{139, 70, 121},
		{24, 160, 241},
		{148, 36, 166},
		{31, 109, 125},
		{66, 157, 171},
		{32, 59, 83},
		{145, 157, 160},
		{212, 96, 118},
		{34, 112, 205},
		{85, 149, 43},
		{100, 49, 89},
		{229, 242, 40},
		{54, 85, 217},
		{209, 169, 231},
		{192, 138, 158},
		{182, 252, 120},
		{65, 48, 111},
		{172, 38, 161},
		{73, 200, 37},
		{64, 75, 149},
		{57, 135, 208},
		{109, 176, 83},
		{82, 178, 161},
		{185, 226, 77},
		{254, 238, 118},
		{157, 169, 81},
		{200, 140, 59},
		{180, 59, 233},
		{199, 124, 228},
		{171, 29, 181},
		{73, 177, 173},
		{234, 250, 90},
		{88, 125, 214},
		{218, 114, 182},
		{111, 138, 249},
		{87, 129, 90},
		{79, 187, 206},
		{111, 63, 181},
		{74, 170, 168},
		{173, 204, 78},
		{67, 115, 211},
		{27, 128, 30},
		{195, 29, 172},
		{93, 167, 134},
		{97, 50, 108},
		{54, 40, 202},
		{106, 155, 116},
		{39, 248, 142},
		{33, 165, 170},
		{202, 38, 114},
		{60, 24, 224},
		{76, 154, 93},
		{159, 137, 165},
		{160, 25, 54},
		{241, 83, 87},
		{175, 41, 255},
		{53, 106, 168},
		{95, 184, 198},
		{191, 127, 66},
		{161, 188, 83},
		{66, 218, 163},
		{125, 195, 28},
		{241, 233, 251},
		{252, 233, 148},
		{143, 75, 35},
		{180, 175, 236},
		{185, 101, 222},
		{111, 204, 219},
		{79, 27, 118},
		{89, 106, 60},
		{170, 115, 156},
		{132, 141, 31},
		{197, 153, 156},
		{234, 219, 32},
		{184, 138, 197},
		{40, 102, 189},
		{201, 164, 136},
		{94, 219, 222},
		{55, 201, 151},
		{127, 186, 180},
		{167, 73, 60},
		{100, 41, 255},
		{103, 150, 167},
		{110, 134, 93},
		{161, 119, 65},
		{121, 224, 210},
		{161, 179, 92},
		{203, 165, 125},
		{53, 190, 140},
		{88, 152, 212},
		{77, 37, 64},
		{128, 203, 144},
		{201, 250, 184},
		{139, 133, 220},
		{166, 194, 190},
		{195, 232, 190},
		{87, 187, 39},
		{114, 241, 108},
		{155, 209, 112},
		{206, 152, 122},
		{56, 232, 189},
		{208, 197, 119},
		{128, 164, 184},
		{25, 46, 208},
		{210, 169, 105},
		{168, 216, 249},
		{165, 142, 52},
		{140, 93, 175},
		{237, 116, 186},
		{180, 247, 151},
		{207, 147, 197},
		{194, 57, 141},
		{38, 51, 118},
		{175, 184, 58},
		{35, 47, 254},
		{73, 24, 208},
		{31, 167, 150},
		{37, 249, 76},
		{140, 229, 144},
		{137, 121, 195},
		{105, 107, 193},
		{203, 179, 81},
		{145, 231, 115},
		{166, 244, 144},
		{133, 178, 217},
		{194, 98, 111},
		{104, 70, 103},
		{200, 194, 98},
		{55, 89, 94},
		{225, 59, 92},
		{93, 244, 130},
		{47, 56, 77},
		{195, 52, 69},
		{193, 229, 137},
		{168, 249, 200},
		{47, 66, 249},
		{163, 38, 47},
		{127, 204, 135},
		{218, 241, 174},
		{189, 187, 255},
		{66, 42, 245},
		{50, 255, 83},
		{128, 140, 45},
		{119, 149, 154},
		{170, 79, 193},
		{74, 178, 157},
		{209, 191, 236},
		{184, 27, 78},
		{57, 93, 50},
		{241, 81, 127},
		{130, 24, 250},
		{215, 110, 153},
		{129, 44, 90},
		{62, 62, 188},
		{222, 36, 177},
		{100, 173, 211},
		{144, 73, 64},
		{208, 134, 210},
		{102, 24, 226},
		{134, 151, 134},
		{47, 87, 123},
		{177, 253, 64},
		{115, 177, 159},
		{125, 126, 111},
		{44, 134, 161},
		{53, 202, 102},
		{217, 199, 179},
		{29, 72, 111},
		{47, 253, 116},
		{253, 95, 94},
		{124, 28, 128},
		{50, 254, 48},
		{205, 205, 137},
		{182, 125, 157},
		{125, 25, 195},
		{26, 32, 99},
		{137, 252, 34},
		{212, 172, 232},
		{63, 220, 122},
		{47, 90, 58},
		{207, 107, 255},
		{151, 42, 230},
		{60, 186, 162},
		{35, 161, 168},
		{52, 114, 103},
		{183, 98, 31},
		{25, 184, 48},
		{224, 194, 188},
		{147, 124, 52},
		{92, 61, 75},
		{252, 236, 53},
		{41, 107, 187},
		{196, 231, 207},
		{101, 124, 112},
		{253, 35, 58},
		{102, 81, 150},
		{127, 72, 64},
		{55, 56, 239},
		{193, 90, 239},
		{246, 105, 150},
		{231, 104, 149},
		{252, 127, 172},
		{216, 179, 133},
		{239, 220, 204},
		{72, 226, 95},
		{62, 252, 62},
		{183, 216, 200},
		{200, 29, 143},
		{56, 166, 57},
		{194, 51, 176},
		{124, 204, 42},
		{170, 82, 178},
		{133, 187, 143},
		{138, 51, 146},
		{77, 71, 233},
		{118, 59, 130},
		{58, 144, 121},
		{100, 205, 136},
		{217, 251, 214},
		{190, 215, 90},
		{204, 193, 51},
		{81, 198, 84},
		{135, 147, 55},
		{224, 246, 58},
		{243, 173, 40},
		{46, 166, 146},
		{79, 228, 254},
		{133, 52, 55},
		{198, 28, 194},
		{181, 216, 136},
		{224, 227, 82},
		{93, 104, 245},
		{161, 64, 86},
		{134, 217, 41},
		{161, 43, 96},
		{155, 175, 192},
		{75, 131, 171},
		{178, 51, 247},
		{175, 235, 229},
		{143, 150, 91},
		{201, 89, 111},
		{211, 181, 116},
		{185, 98, 180},
		{204, 152, 94},
		{172, 25, 185},
		{123, 94, 100},
		{205, 28, 240},
		{90, 146, 52},
		{217, 229, 241},
		{58, 88, 143},
		{202, 62, 69},
		{193, 79, 237},
		{206, 42, 27},
		{87, 24, 30},
		{178, 235, 75},
		{179, 222, 65},
		{226, 170, 249},
		{233, 113, 117},
		{101, 119, 204},
		{216, 52, 209},
		{218, 193, 174},
		{165, 140, 187},
		{139, 51, 220},
		{91, 179, 204},
		{152, 177, 63},
		{44, 153, 240},
		{98, 35, 34},
		{197, 143, 161},
		{106, 127, 193},
		{63, 116, 91},
		{204, 67, 25},
		{255, 255, 204},
		{174, 255, 154},
		{149, 25, 230},
		{203, 202, 94},
		{51, 151, 43},
		{182, 164, 64},
		{161, 223, 196},
		{154, 153, 226},
		{93, 181, 251},
		{246, 98, 126},
		{229, 159, 194},
		{168, 147, 161},
		{204, 237, 43},
		{121, 201, 122},
		{48, 202, 66},
		{120, 248, 163},
		{127, 86, 211},
		{158, 138, 170},
		{86, 175, 111},
		{173, 54, 25}}
	return colormap
}
