package senml

import (
	"testing"
)

func TestFindByName(t *testing.T) {
	record := Record{Name: "temperature"}
	finder := FindByName("temperature")
	if !finder(record) {
		t.Errorf("FindByName should return true for matching name")
	}
	finder = FindByName("Temperature")
	if !finder(record) {
		t.Errorf("FindByName should be case-insensitive")
	}
	finder = FindByName("humidity")
	if finder(record) {
		t.Errorf("FindByName should return false for non-matching name")
	}
}

func TestFindByNormalizedName(t *testing.T) {
	record := Record{BaseName: "sensor/", Name: "temperature"}
	finder := FindByNormalizedName("sensor/", "temperature")
	if !finder(record) {
		t.Errorf("FindByNormalizedName should return true for matching base name and name")
	}
	finder = FindByNormalizedName("SENSOR/", "TEMPERATURE")
	if !finder(record) {
		t.Errorf("FindByNormalizedName should be case-insensitive")
	}
	finder = FindByNormalizedName("sensor/", "humidity")
	if finder(record) {
		t.Errorf("FindByNormalizedName should return false for non-matching name")
	}
}

func TestFindByUnit(t *testing.T) {
	record := Record{Unit: "C"}
	finder := FindByUnit("C")
	if !finder(record) {
		t.Errorf("FindByUnit should return true for matching unit")
	}
	finder = FindByUnit("F")
	if finder(record) {
		t.Errorf("FindByUnit should return false for non-matching unit")
	}
}

func TestGetRecordWithFindByNormalizedName(t *testing.T) {
	pack := Pack{
		{BaseName: "sensor/", Name: "temperature"},
		{BaseName: "sensor/", Name: "humidity"},
	}

	record, ok := pack.GetRecord(FindByNormalizedName("sensor/", "temperature"))
	if !ok {
		t.Fatalf("Expected to find record with name 'temperature'")
	}
	// After Normalize, BaseName is merged into Name, BaseName is cleared
	if record.BaseName != "" || record.Name != "sensor/temperature" {
		t.Errorf("Got wrong record: got BaseName=%q Name=%q", record.BaseName, record.Name)
	}
}
