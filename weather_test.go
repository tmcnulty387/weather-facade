package weather

import (
	"testing"
)

func TestGetTemp(t *testing.T) {
	// Using coordinates for Irondequoit, NY
	lat := 43.21
	lon := -77.58

	temp, err := GetTemp(lat, lon)

	// Test 1: Did the API call fail?
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	// Test 2: Did we get a realistic temperature?
	// (Unless it's the end of the world, temp should be between -50 and 50 Celsius)
	if temp < -50 || temp > 50 {
		t.Errorf("Temperature %v seems unrealistic for this planet", temp)
	}

	t.Logf("Success! Current temp in Irondequoit is: %.1f°C", temp)
}

func TestGetTemp_InvalidCoords(t *testing.T) {
	// Latitude can't be 999 (it maxes at 90)
	lat := 999.9
	lon := 0.0

	temp, err := GetTemp(lat, lon)

	// We EXPECT an error here because the coordinates are impossible
	if err == nil {
		t.Errorf("Expected an error for latitude 999.9, but got nil and temp: %v", temp)
	} else {
		t.Logf("Correctly caught the error: %v", err)
	}
}
