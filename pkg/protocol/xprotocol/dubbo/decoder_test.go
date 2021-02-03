package dubbo

import (
	"context"
	"testing"

	"mosn.io/mosn/pkg/types"
	"mosn.io/pkg/buffer"
)

func TestDecodeFramePanic(t *testing.T) {
	data := buffer.NewIoBufferBytes(complexData)
	// decode attachement
	ctx := context.WithValue(context.TODO(), types.ContextKeyListenerName, IngressDubbo)

	_, err := decodeFrame(ctx, data)
	if err != nil {
		t.Logf("recover dubbo decode panic:%s", err)
		return
	}
}

func TestDecodeSkipCheap(t *testing.T) {
	data := buffer.NewIoBufferBytes(benchmarkData)
	// decode attachement
	ctx := context.WithValue(context.TODO(), types.ContextKeyListenerName, "Cheap")

	_, err := decodeFrame(ctx, data)
	if err != nil {
		t.Errorf("dubbo decode [skip cheap] panic:%s", err)
		return
	}
}

func TestDecodeSkip(t *testing.T) {
	data := buffer.NewIoBufferBytes(benchmarkData)
	// decode attachement
	ctx := context.WithValue(context.TODO(), types.ContextKeyListenerName, IngressDubbo)

	_, err := decodeFrame(ctx, data)
	if err != nil {
		t.Errorf("dubbo decode [cheap] panic:%s", err)
		return
	}
}

func BenchmarkDecodeSkipCheap(t *testing.B) {
	ctx := context.WithValue(context.TODO(), types.ContextKeyListenerName, "Cheap")
	for i := 0; i < t.N; i++ {
		data := buffer.NewIoBufferBytes(benchmarkData)
		_, err := decodeFrame(ctx, data)
		if err != nil {
			t.Errorf("recover dubbo decode panic:%s", err)
		}
	}
}

func BenchmarkDecodeSkip(t *testing.B) {
	ctx := context.WithValue(context.TODO(), types.ContextKeyListenerName, IngressDubbo)
	for i := 0; i < t.N; i++ {
		data := buffer.NewIoBufferBytes(benchmarkData)
		_, err := decodeFrame(ctx, data)
		if err != nil {
			t.Errorf("recover dubbo decode panic:%s", err)
		}
	}
}

var complexData = []byte{
	0xda, 0xbb, 0xc2, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x04, 0x45,
	0x05, 0x32, 0x2e, 0x30, 0x2e, 0x32, 0x30, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6d, 0x61, 0x6c,
	0x6c, 0x2e, 0x64, 0x73, 0x66, 0x2e, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x0e, 0x75, 0x6e,
	0x69, 0x74, 0x54, 0x65, 0x73, 0x74, 0x42, 0x79, 0x55, 0x6e, 0x69, 0x74, 0x30, 0x5f, 0x4c, 0x6a,
	0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3b,
	0x4c, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x64, 0x73, 0x66, 0x2f, 0x7a,
	0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2f, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3b, 0x4c, 0x6a,
	0x61, 0x76, 0x61, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x3b, 0x4c, 0x6a,
	0x61, 0x76, 0x61, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x4d, 0x61, 0x70, 0x3b, 0x01, 0x30, 0x43,
	0x30, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x64, 0x73, 0x66, 0x2e,
	0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x9c, 0x0d,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x0c, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x4d, 0x61, 0x70, 0x0b, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x09, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x09, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x08, 0x76,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x04, 0x70, 0x61, 0x72, 0x32, 0x04, 0x70, 0x61, 0x72,
	0x31, 0x60, 0x4e, 0x4e, 0x43, 0x30, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6d, 0x61, 0x6c, 0x6c,
	0x2e, 0x64, 0x73, 0x66, 0x2e, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x99, 0x10, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x48, 0x6f, 0x73,
	0x74, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x0c, 0x62, 0x61, 0x73, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x08, 0x76, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x61, 0x4e, 0x4e, 0x4e, 0x4e, 0x4e, 0x90, 0x0e, 0x42, 0x61, 0x73,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x4e, 0x90, 0x4e, 0x09, 0x31,
	0x32, 0x37, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x04, 0x47, 0x7a, 0x30, 0x31, 0x90, 0x4e, 0x90,
	0x4e, 0x4e, 0x4e, 0x70, 0x13, 0x6a, 0x61, 0x76, 0x61, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x02, 0x78, 0x78, 0x4d, 0x30, 0x26, 0x6a,
	0x61, 0x76, 0x61, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x4d, 0x61, 0x70, 0x02, 0x78, 0x78, 0x02, 0x79, 0x79, 0x5a, 0x02, 0x79, 0x79, 0x4d,
	0x91, 0x02, 0x79, 0x79, 0x02, 0x79, 0x79, 0x5a, 0x5a, 0x48, 0x10, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x4e, 0x0d, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x04, 0x67, 0x7a, 0x30, 0x31, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x13, 0x61, 0x77, 0x71, 0x33, 0x34, 0x33, 0x35, 0x35, 0x65,
	0x67, 0x64, 0x74, 0x36, 0x75, 0x37, 0x69, 0x6b, 0x75, 0x74, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x30, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6d, 0x61, 0x6c, 0x6c, 0x2e,
	0x64, 0x73, 0x66, 0x2e, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x07, 0x70, 0x72, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x4e, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x04, 0x33, 0x30, 0x30, 0x30, 0x0b, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x2e, 0x7a, 0x6f, 0x6e, 0x65, 0x04, 0x47, 0x7a, 0x30, 0x31, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x30, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6d, 0x61, 0x6c, 0x6c, 0x2e, 0x64, 0x73,
	0x66, 0x2e, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x04, 0x67, 0x7a, 0x30, 0x31, 0x0f, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x07, 0x61, 0x70,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x14, 0x7a, 0x6f, 0x6e, 0x65, 0x2d, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x0b, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x14, 0x7a, 0x6f, 0x6e, 0x65, 0x2d, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x0a, 0x64,
	0x6d, 0x61, 0x6c, 0x6c, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x0b, 0x2e, 0x74, 0x6f, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x31, 0x31, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x75,
	0x6e, 0x69, 0x74, 0x04, 0x67, 0x7a, 0x30, 0x31, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x0c, 0x31, 0x30, 0x2e, 0x31, 0x32, 0x2e, 0x32, 0x31,
	0x34, 0x2e, 0x36, 0x33, 0x0c, 0x64, 0x6d, 0x61, 0x6c, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x65, 0x02, 0x65, 0x6e, 0x5a,
}

var benchmarkData = []byte{
	0xda, 0xbb, 0xc2, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4c,
	0x00, 0x00, 0x01, 0xfb, 0x05, 0x32, 0x2e, 0x36, 0x2e, 0x32, 0x30, 0x30, 0x63, 0x6f, 0x6d, 0x2e,
	0x72, 0x61, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x0e, 0x32, 0x2e, 0x30,
	0x2e, 0x31, 0x2d, 0x76, 0x70, 0x63, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x0d, 0x73, 0x74, 0x6f, 0x70,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4a, 0x6f, 0x62, 0x30, 0x2e, 0x4c, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x61, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x53, 0x74, 0x6f, 0x70, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x4a, 0x6f, 0x62, 0x3b, 0x43, 0x30, 0x2c, 0x63, 0x6f, 0x6d, 0x2e,
	0x72, 0x61, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x4a, 0x6f, 0x62, 0x98, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x06, 0x61, 0x70, 0x70,
	0x4b, 0x65, 0x79, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x08, 0x74, 0x61, 0x6f, 0x62, 0x61, 0x6f,
	0x49, 0x64, 0x60, 0x01, 0x31, 0x01, 0x31, 0x01, 0x31, 0x01, 0x31, 0x01, 0x31, 0x01, 0x31, 0x91,
	0xe1, 0x48, 0x05, 0x61, 0x2e, 0x73, 0x2e, 0x75, 0x0a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x05, 0x61, 0x2e, 0x73, 0x2e, 0x74, 0x10, 0x35, 0x33, 0x35, 0x36, 0x32, 0x35,
	0x38, 0x38, 0x39, 0x38, 0x37, 0x37, 0x31, 0x33, 0x31, 0x34, 0x05, 0x61, 0x2e, 0x73, 0x2e, 0x64,
	0x10, 0x6b, 0x65, 0x66, 0x75, 0x2e, 0x6b, 0x75, 0x61, 0x69, 0x6d, 0x61, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x04, 0x70, 0x61, 0x74, 0x68, 0x30, 0x30, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x79, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x14, 0x63, 0x6f, 0x6d, 0x2d, 0x72, 0x61,
	0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x78, 0x09, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x30, 0x30, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61,
	0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x0e, 0x32, 0x2e, 0x30, 0x2e, 0x31, 0x2d, 0x76, 0x70, 0x63, 0x2d, 0x74, 0x65, 0x73,
	0x74, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x04, 0x35, 0x30, 0x30, 0x30, 0x5a,
}
