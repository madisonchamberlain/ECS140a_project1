package triangle

import "testing"

func TestGetTriangleType(t *testing.T) {
	type Test struct {
		a, b, c  int
		expected triangleType
	}

	var tests = []Test{
		// Unknown = where one side is very large
		{1001, 5, 6, UnknownTriangle},
		{10, 2001, 6, UnknownTriangle},
		{10, 20, 3001, UnknownTriangle},
		// Unknown = where one side is less than/equal to 0
		{0, 5, 6, UnknownTriangle},
		{10, 0, 6, UnknownTriangle},
		{10, 20, 0, UnknownTriangle},
		// Right = a^2 + b^2 = c^2
		{3, 4, 5, RightTriangle},
		// Acute = a^2 + b^2 > c^2
		{3, 4, 2, AcuteTriangle},
		// Invalid = when sides dont connect
		{1, 2, 3, InvalidTriangle},
		// Obtuse = a^2 + b^2 < c^2
		{3, 4, 6, ObtuseTriangle},
	}

	for _, test := range tests {
		actual := getTriangleType(test.a, test.b, test.c)
		if actual != test.expected {
			t.Errorf("getTriangleType(%d, %d, %d)=%v; want %v", test.a, test.b, test.c, actual, test.expected)
		}
	}
}
