package test

// TestData 用于管理测试数据
type TestData struct {
	WardIDs    []string
	PatientIDs []string
	BedIDs     []string
	CameraIDs  []string
}

var testData = &TestData{}

// CreateTestWard 创建测试病房数据
func CreateTestWard() (string, error) {
	// TODO: 实现创建测试病房的逻辑
	return "", nil
}

// CreateTestPatient 创建测试病患数据
func CreateTestPatient() (string, error) {
	// TODO: 实现创建测试病患的逻辑
	return "", nil
}

// CleanupTestData 清理测试数据
func CleanupTestData() error {
	// TODO: 实现清理测试数据的逻辑
	return nil
} 